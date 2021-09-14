import subprocess
import os
import json
import shutil
import logging

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

def compute_percentage(explored_contracts, known_contracts):
    total = len(known_contracts)
    for c in explored_contracts:
        if c in known_contracts:
            known_contracts.remove(c)
    return (total - len(known_contracts)) / total

def main():
    logging.info("Initializing benchmark")

    logging.info("Reading \"benchmark.json\" file")
    fbugs = open(f'{os.getcwd()}/benchmark.json', 'r')
    bugs = json.load(fbugs)
    fbugs.close()


    results = {}
    for key in bugs:
        logging.info(f"Testing {key} vulnerability")
        clean_results_folder(key)

        logging.info(f"Executing {key} contracts")
        result = subprocess.call([f'{os.getcwd()}/run.sh', '--contract_dir', f'{os.getcwd()}/examples/{key}'])
        if (result != 0):
            print(f'An error occurred while running {key} vulnerability')
            continue

        logging.info(f"Get vulnerable contracts")
        explored_contracts = get_explored_vulnerabilities(key, bugs[key]["results_file"])
        perc = compute_percentage(explored_contracts, bugs[key]["known_contracts"])

        results[key] = {
            "total": len(bugs[key]["known_contracts"]),
            "explored_contracts": len(explored_contracts),
            "percentage": perc,
        }

    logging.info("Writing results")
    f = open(f'{os.getcwd()}/results.json', 'w')
    f.write(json.dumps(results))
    f.close()

    logging.info("Finishing benchmark")

if __name__ == "__main__":
    main()
