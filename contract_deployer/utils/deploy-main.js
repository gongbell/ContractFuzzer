#! /local/bin/bnode

import * as fs from "fs";
import {
    ContractUtils,
    Contract_DIR,
    web3,
    defaultAmountParamsWithValue,
    defaultAmountParamsWithoutValue,
    dotenv,
} from "../utils/ContractUtils.js";
import {
    defaultAccount
} from "../utils/Account.js";

dotenv.config();

const BIN_DIR = process.env.BIN_SUB_DIR;
const BIN_SUFFIX = process.env.BIN_SUFFIX;
const ABI_DIR = process.env.ABI_SUB_DIR;
const ABI_SUFFIX = process.env.ABI_SUFFIX;
const CONFIG_PATH = process.env.CONFIG_PATH;

const util = new ContractUtils();

const fread = fs.readFileSync;
const fopen = fs.openSync;
const fclose = fs.closeSync;
const fwrite = fs.writeFileSync;
const listdir = fs.readdirSync;
const print = console.log;
const error = console.error;
const BalanceOf = web3.eth.getBalance;
function deJSON(file) {
    let f = fopen(file, "r");
    let data = fread(f, {
        encoding: "utf8"
    });
    let js = JSON.parse(data);
    fclose(f);
    return js;
}
function writeJSONtoFile(file, jsData) {
    let f = fopen(file, "w");
    fwrite(f, jsData, {
        encoding: "utf8"
    });
    fclose(f);
}

function deBin(workplace, name) {
    let file = workplace + "/" + BIN_DIR + "/" + name + BIN_SUFFIX;
    let f = fopen(file, "r");
    let data = fread(f, {
        encoding: "utf8"
    });
    fclose(f);
    data = "0x" + data;
    return data;
}

function deAbi(workplace, name) {
    let file = workplace + "/" + ABI_DIR + "/" + name + ABI_SUFFIX;
    let f = fopen(file, "r");
    let data = fread(f, {
        encoding: "utf8"
    });
    fclose(f);
    data = JSON.parse(data);
    return data;
}

function getJsonObjLength(jsonObj) {
    var Length = 0;
    for (var item in jsonObj) {
        Length++;
    }
    return Length;
}
async function deployAcontract(workplace, contract) {
    let name = contract.name;
    let param_Values = contract.param_Values;
    if (param_Values == "none") {
        param_Values = {};
    }
    let from = contract.from;
    let gas = contract.gas;
    let value = contract.value;
    let payable = contract.payable;
    let values = contract.values;
    let abi = deAbi(workplace, name);
    let bin = deBin(workplace, name);
    contract = {
        name: name,
        abi: abi,
        bin: bin
    };
    let aParams;
    // gas = 80000;
    if (payable == true) {
        aParams = {
            from: from,
            value: value,
            gas: gas
        };
    } else {
        aParams = {
            from: from,
            gas: gas
        };
    }

    let len = getJsonObjLength(param_Values)
    if (len == 0) {
        return  await util.DeployWithoutParams(contract, aParams);
    } else {
        let paramArr = values;
        return await util.DeployWithParams(contract, paramArr, aParams);
    }
}

async function deploy(js) {
    let home = js.home;
    let value = js.value;
    let gas = js.gas;
    let contracts = js.contracts;
    let from = js.from;
    if (from == undefined || from == "none") {
        from = web3.eth.accounts[0];
    }
    for (let index in contracts) {
        let contract = contracts[index];
        if (contract.deployed != undefined && contract.deployed == 1)
            continue;
        let workplace;
        if (contract.home == undefined || contract.home == "none")
            contract.home = home;
        if (contract.childhome == undefined || contract.childhome == "none")
            workplace = contract.home;
        else
            workplace = contract.home + contract.childhome;
        if (contract.gas == undefined || contract.gas == "none")
            contract.gas = gas;
        if (contract.value == undefined || contract.value == "none")
            contract.value = value;
        if (contract.from == undefined || contract.from == "none") {
            contract.from = from;
        }
        try {
            // console.log(workplace,contract);
            let instance = await deployAcontract(workplace, contract);
            contract.deployed = 1;
            // console.log(instance);
            contract["address"] = instance.address;
        } catch (err) {
            print("Error Profile");
            print("contract:", contract.name);
            print(err.toString().split("\n")[0]);
            contract.deployed = 0;
        }
    }
    console.log(js);
    writeJSONtoFile(js.file_path,JSON.stringify(js));
    return JSON.stringify(js);
}

function deployBatch(js_arr) {
    for (let js of js_arr) {
        deploy(js);
    }
}
function getConfigs(dir) {
    let items = listdir(dir);
    let files = [];
    let index = 0;
    for (let i=0;i<items.length;i++){
        let item =  items[i];
        files[index++] = dir + "/" + item;
    }
    return files;
}

function executeOuterConfigs() {
    let dir = CONFIG_PATH;
    let configs = getConfigs(dir);
    let js_arr = [];
    const batchMaxSize = 20;
    const batchInterval = 30 * 1000;
    let times = 0;
    // configs = configs.slice(0, configs.length - 1);
     for (let config of configs) {
        let js = deJSON(config);
        js.file_path = config;
        js_arr.push(js);
        if (js_arr.length > batchMaxSize) {
            setTimeout(deployBatch, times * batchInterval, js_arr);
            times++;
            js_arr = [];
        }
    }
    console.log(js_arr);
    setTimeout(deployBatch, times * batchInterval, js_arr);
}

function main() {
    executeOuterConfigs()
}

let __name__ = "__main__";
if (__name__ == "__main__") {
    main();
}
