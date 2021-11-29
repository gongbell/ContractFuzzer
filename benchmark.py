import subprocess
import os
import json
import shutil
import logging
import time

def delete_folder_content(folder):
    for filename in os.listdir(folder):
        file_path = os.path.join(folder, filename)
        try:
            if os.path.isfile(file_path) or os.path.islink(file_path):
                os.unlink(file_path)
            elif os.path.isdir(file_path):
                shutil.rmtree(file_path)
        except Exception as e:
            print('Failed to delete %s. Reason: %s' % (file_path, e))

def clean_results_folder(vbl):
    delete_folder_content(f"{os.getcwd()}/examples/{vbl}/fuzzer/reporter")
    os.mkdir(f"{os.getcwd()}/examples/{vbl}/fuzzer/reporter/bug")


def get_explored_vulnerabilities(vbl, results_file):
    contracts = []
    f = open(f"{os.getcwd()}/examples/{vbl}/fuzzer/reporter/bug/{results_file}.list", "r")
    lines = f.readlines()
    for line in lines:
        contracts.append(line.strip())
    temp = set(contracts)
    f.close()
    return list(temp)

def get_tested_contracts(vbl):
    contracts = []
    f = open(f"{os.getcwd()}/examples/{vbl}/fuzzer/config/contracts.list", "r")
    lines = f.readlines()
    for line in lines:
        contracts.append(line.strip())
    temp = set(contracts)
    f.close()
    return list(temp)

def compute_percentage(explored_contracts, known_contracts):
    total = len(known_contracts)
    for c in explored_contracts:
        if c in known_contracts:
            known_contracts.remove(c)
    return (total - len(known_contracts)) / total

def get_elapsed_times():
    times = {}
    f = open("/tmp/c_elapsed.csv", "r")
    lines = f.readlines()
    for line in lines:
        values = line.split(",")
        if len(values) != 2:
            continue
        contract = values[0].strip()
        elapsed_time = values[1].strip()
        times[contract] = elapsed_time
    f.close()
    return times

def get_contract_size(vulnerability, contract_name):
    size = -1
    if os.path.isfile(f"{os.getcwd()}/examples/{vulnerability}/verified_contract_bins/{contract_name}.bin"):
        size = os.path.getsize(f"{os.getcwd()}/examples/{vulnerability}/verified_contract_bins/{contract_name}.bin") / 2
    if os.path.isfile(f'{os.getcwd()}/examples/{vulnerability}/verified_contract_bins/{contract_name}.bin-runtime'):
        size = os.path.getsize(f"{os.getcwd()}/examples/{vulnerability}/verified_contract_bins/{contract_name}.bin-runtime") / 2
    
    return size

def get_num_functions(vulnerability, contract_name):
    num = -1
    if os.path.isfile(f"{os.getcwd()}/examples/{vulnerability}/verified_contract_abis/{contract_name}.abi"):
        f = open(f"{os.getcwd()}/examples/{vulnerability}/verified_contract_abis/{contract_name}.abi", "r")
        abi = json.load(f)
        num = len([f for f in abi if f["type"] == "function"])
        f.close()
    return num

def get_execution_result(v, contract_name, explored_contracts, times):
    return {
        "name": contract_name,
        "elapsed_time": times[contract_name] if contract_name in times else -1,
        "vulnerable": contract_name in explored_contracts,
        "size": get_contract_size(v, contract_name),
        "num_functions": get_num_functions(v, contract_name)
    }


def main(reps = 1):
    logging.info("Initializing benchmark")

    logging.info("Reading \"benchmark.json\" file")
    fbugs = open(f'{os.getcwd()}/benchmark.json', 'r')
    bugs = json.load(fbugs)
    fbugs.close()

    results = []
    for _ in range(reps):
        execution = {}
        for key in bugs:
            logging.info(f"Testing {key} vulnerability")
            clean_results_folder(key)

            logging.info(f"Executing {key} contracts")
            start = time.time()
            os.environ["CONTRACT_DIR"] = f'{os.getcwd()}/examples/{key}'
            result = subprocess.call([f'{os.getcwd()}/fuzzer_run.sh'])
            elapsed_time = time.time() - start
            if (result != 0):
                print(f'An error occurred while running {key} vulnerability')
                continue

            logging.info(f"Get vulnerable contracts")
            tested_contracts = get_tested_contracts(key)
            explored_contracts = get_explored_vulnerabilities(key, bugs[key])
            perc = compute_percentage(explored_contracts, tested_contracts.copy())
            times = get_elapsed_times()

            execution[key] = {
                "contracts_total": len(tested_contracts),
                "contracts_flagged": len(explored_contracts),
                "percentage": perc,
                "elapsed_time": elapsed_time,
                "contracts": [get_execution_result(key, c, explored_contracts, times) for c in tested_contracts]
            }
        results.append(execution)

    logging.info("Writing results")
    f = open('/output/results.json', 'w')
    f.write(json.dumps(results))
    f.close()

    logging.info("Finishing benchmark")

if __name__ == "__main__":
    main(5)
