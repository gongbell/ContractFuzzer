#! /local/bin/bnode

import * as fs from "fs";
import {
    ContractUtils,
    Contract_DIR,
    web3,
    defaultAmountParamsWithValue,
    defaultAmountParamsWithoutValue,
} from "../utils/ContractUtils.js";
import {
    defaultAccount
} from "../utils/Account.js";
import {
    DBHelper
} from "../utils/DBHelper.js"
import {
    config
} from "dotenv";
// let argv = process.argv;
// let argc = argv.length;
// print("argc:", argc);
// print("argv:", argv);
const util = new ContractUtils();
const fread = fs.readFileSync;
const fopen = fs.openSync;
const fclose = fs.closeSync;
const fwrite = fs.writeFileSync;
const print = console.log;
const error = console.error;
const listdir = fs.readdirSync;
const BalanceOf = web3.eth.getBalance;
const BIN_DIR = "verified_contract_bins";
const BIN_SUFFIX = ".bin";
const ABI_DIR = "verified_contract_abis";
const ABI_SUFFIX = ".abi";

function deJSON(file) {
    let f = fopen(file, "r");
    let data = fread(f, {
        encoding: "utf8"
    });
    // print(data);
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
    data = "0x" + data;
    return data;
}

function deAbi(workplace, name) {
    let file = workplace + "/" + ABI_DIR + "/" + name + ABI_SUFFIX;
    let f = fopen(file, "r");
    let data = fread(f, {
        encoding: "utf8"
    });
    data = JSON.parse(data);
    // print(data);
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
        // print("has no parameters");
        await util.DeployWithoutParams(contract, aParams);
    } else {
        // print("has parameters");
        let paramArr = values;
        // print(paramArr);
        await util.DeployWithParams(contract, paramArr, aParams);
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
            //print(contract)
            await deployAcontract(workplace, contract);
            contract.deployed = 1;
        } catch (err) {
            print("Error Profile");
            print("contract:", contract.name);
            print(err.toString().split("\n")[0]);
            contract.deployed = 0;
        }
    }
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
    for (let item of items) {
        files[index++] = dir + "/" + item;
    }
    return files;
}

function executeOuterConfigs() {
    // let dir = "./DownloadContracts/config";
    let dir = "/home/liuye/smart_contracts/verified_contract_config";
    let configs = getConfigs(dir);
    // let config_dir = "/home/liuye/smart_contracts/verified_contract_config";
    // let data = fread("./deployer/undeployedContract.list", "utf-8");
    // let contracts = data.split("\n");
    // let configs = [];
    // for (let contract of contracts) {
    //     configs.push(config_dir + "/" + contract + ".json");
    // }
    // return;
    let js_arr = [];
    const batchMaxSize = 20;
    const batchInterval = 30 * 1000;
    let times = 0;
    // for (let config of configs) {
    //     print(config)
    //     let js = deJSON(config);
    //     let newjsData = await deploy(js);
    //     writeJSONtoFile(config, newjsData);
    //     // break;
    // }
    configs = configs.slice(0, configs.length - 1);
    // print(configs)
    for (let config of configs) {
        let js = deJSON(config);
        js_arr.push(js);
        if (js_arr.length > batchMaxSize) {
            setTimeout(deployBatch, times * batchInterval, js_arr);
            times++;
            js_arr = [];
        }
    }
    setTimeout(deployBatch, times * batchInterval, js_arr);
}

async function executeSingleConfig() {
    // let file = "./DownloadContracts/config/YesNo.abi";
    let file = "./deployer/MyCaller3.json";
    let js = deJSON(file);
    // print(js);
    let newjsData = await deploy(js);
    writeJSONtoFile(file, newjsData);
}
async function executeDefault() {
    let file = "./deployer/default.json";
    let js = deJSON(file);
    // print(js);
    let newjsData = await deploy(js);
    writeJSONtoFile(file, newjsData);
}
async function deployKitties() {
    let file = "./contract/config/KittyCore.json";
    let js = deJSON(file);
    let newjsData = await deploy(js);
    writeJSONtoFile(file, newjsData);
}
async function main() {
    // print("main");
    //  await executeDefault()
    if (__option__ == "__single__") {
        // await executeSingleConfig();
        await deployKitties();
    } else {
        // await executeOuterConfigs();
        executeOuterConfigs()
    }

}

async function testGaslessSend() {
    let c = "C";
    let d1 = "D1";
    let d2 = "D2";
    let C = await util.contract(c);
    let D1 = await util.contract(d1);
    let D2 = await util.contract(d2);
    let n = 50000;
    let receipt = await C.pay(n, D1.address, defaultAmountParamsWithoutValue);
    receipt = await C.pay(n, D2.address, defaultAmountParamsWithoutValue);
}
async function testExceptionDisorder() {
    let dbUtil = new DBHelper();
    let alice = "Alice";
    let bob = "BobWatch"
    let Alices = await util.contracts(alice);
    let Bobs = await util.contracts(bob);
    let Alice = Alices[0];
    let Bob = Bobs[0];
    print(0, "Alice", Alice.address);
    print(0, "Bob", Bob.address);
    // print(Bob.abi);
    let xpre = await Bob.x(defaultAmountParamsWithoutValue);
    let receipt = await Bob.pong(Alice.address, defaultAmountParamsWithValue);
    let xnext = await Bob.x(defaultAmountParamsWithoutValue);
    print(xpre, xnext);

    bob = "Bob2";
    Bobs = await util.contracts(bob);
    Bob = Bobs[0];
    print(0, "Alice", Alice.address);
    print(0, "Bob", Bob.address);
    receipt = await Bob.setX(1, defaultAmountParamsWithoutValue);
    xpre = await Bob.x(defaultAmountParamsWithoutValue);
    receipt = await Bob.pongProblem(Alice.address, defaultAmountParamsWithValue);
    xnext = await Bob.x(defaultAmountParamsWithoutValue);
    print(xpre, xnext);
}
async function testContractWithInitialParams() {
    let name = "YesNo";
    let contracts = await util.contracts(name);
    for (let contract of contracts) {
        let factHash = await contract.factHash(defaultAmountParamsWithoutValue);
        if (factHash != "0000000000000000000000000000000000000000000000000000000000000160") {
            print("not equal", factHash);
        } else {
            print("equal", factHash);
        }
        let fee = await contract.fee(defaultAmountParamsWithoutValue);
        if (fee.toString() != "608") {
            print("not equal", fee);
        } else {
            print("equal", fee);
        }
        let url = await contract.url(defaultAmountParamsWithoutValue);
        if (url.toString() != "Mayweather Yes") {
            print("not equal", url);
        } else {
            print("equal", url);
        }
        let ethaddr = await contract.ethAddr(defaultAmountParamsWithoutValue);
        if (ethaddr.toString() != "0x1a0") {
            print("not equal", ethaddr);
        } else {
            print("equal", ethaddr);
        }
    }
}
async function case1() {
    let name = "MyCaller2";
    let contract = await util.contract(name);
    let testeeName = "InternalContract";
    let testee = await util.contract(testeeName);
    print(defaultAccount, contract.address, testee.address);
    try {
        defaultAmountParamsWithoutValue.gas = 800000;
        // let contract_addr = "0xfe55080a36877b7edc9a57cccd85697aaf64f7a3";
        // let msg_data = "0x11bac8e5000000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000b";
        let contract_addr = testee.address;
        let msg_data = "0x93ed3714";
        let receipt = await contract.MyCallWithoutValue(contract_addr, msg_data, defaultAmountParamsWithoutValue);
        print(receipt);
        // receipt = await contract.getContractAddr(defaultAmountParamsWithoutValue);
        // print(receipt);
        // receipt = await contract.getCallMsgData(defaultAmountParamsWithoutValue);
        // print(receipt);
        receipt = await contract.call_contract_addr(defaultAmountParamsWithoutValue);
        print(receipt);
        receipt = await contract.call_msg_data(defaultAmountParamsWithoutValue);
        print(receipt);
        receipt = await contract.count(defaultAmountParamsWithoutValue);
        print(receipt);
        receipt = await testee.internal_sign(defaultAmountParamsWithoutValue);
        print(receipt);
    } catch (err) {
        print("Error profile");
        print(err.toString().split("\n")[0]);
        print(err);
    }
}
async function case2() {
    let name = "DirectCaller";
    let contract = await util.contract(name);
    let testeeName = "DirectCallee";
    let testee = await util.contract(testeeName);
    print(defaultAccount, contract.address, testee.address);
    try {
        defaultAmountParamsWithoutValue.gas = 800000;
        let testee_addr = testee.address;

        let receipt = await contract.test(testee_addr, defaultAmountParamsWithoutValue);
        print(receipt);
        let count = await contract.count(defaultAmountParamsWithoutValue);
        print(count);
        print("\n");

        receipt = await contract.testThrow(testee_addr, defaultAmountParamsWithoutValue);
        print(receipt);
        let isTestThrow = await contract.isTestThrow(defaultAmountParamsWithoutValue);
        print(isTestThrow);
        print("\n");

        receipt = await contract.testRet10(testee_addr, defaultAmountParamsWithoutValue);
        print(receipt);
        let retCount = await contract.retCount(defaultAmountParamsWithoutValue);
        let isTestRet = await contract.istestRet(defaultAmountParamsWithoutValue);
        print(retCount, isTestRet);
        print("\n");

        receipt = await contract.testRetThrow(testee_addr, defaultAmountParamsWithoutValue);
        print(receipt);
        retCount = await contract.retCount(defaultAmountParamsWithoutValue);
        isTestRet = await contract.istestRet(defaultAmountParamsWithoutValue);
        print(retCount, isTestRet);


    } catch (err) {
        print("Error profile");
        print(err.toString().split("\n")[0]);
        print(err);
    }
}
async function case3() {
    let name = "MyCaller2";
    let contract = await util.contract(name);
    let testeeName = "DirectCallee";
    let testee = await util.contract(testeeName);
    print(defaultAccount, contract.address, testee.address);
    try {
        defaultAmountParamsWithoutValue.gas = 800000;
        // let contract_addr = "0xfe55080a36877b7edc9a57cccd85697aaf64f7a3";
        // let msg_data = "0x11bac8e5000000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000b";
        let count = await testee.retCount(defaultAmountParamsWithoutValue);
        print(count);
        print("\n");

        let contract_addr = testee.address;
        let msg_data = "0xe56a0fed0000000000000000000000000000000000000000000000000000000000000001";
        let receipt = await contract.MyCallWithoutValue(contract_addr, msg_data, defaultAmountParamsWithoutValue);
        print(receipt);
        receipt = await contract.call_contract_addr(defaultAmountParamsWithoutValue);
        print(receipt);
        receipt = await contract.call_msg_data(defaultAmountParamsWithoutValue);
        print(receipt);
        count = await testee.retCount(defaultAmountParamsWithoutValue);
        print(count);
        print("\n");

        msg_data = "0xe56a0fed0000000000000000000000000000000000000000000000000000000000000000";
        receipt = await contract.MyCallWithoutValue(contract_addr, msg_data, defaultAmountParamsWithoutValue);
        print(receipt);
        receipt = await contract.call_contract_addr(defaultAmountParamsWithoutValue);
        print(receipt);
        receipt = await contract.call_msg_data(defaultAmountParamsWithoutValue);
        print(receipt);
        count = await testee.retCount(defaultAmountParamsWithoutValue);
        print(count);
        print("\n");

    } catch (err) {
        print("Error profile");
        print(err.toString().split("\n")[0]);
        print(err);
    }
}
async function case4() {
    let name = "DirectCaller3";
    let contract = await util.contract(name);
    let testeeName = "DirectCallee";
    let testee = await util.contract(testeeName);
    print(defaultAccount, contract.address, testee.address);
    try {
        defaultAmountParamsWithoutValue.gas = 800000;
        let testee_addr = testee.address;

        let count = await contract.count(defaultAmountParamsWithoutValue);
        print(count);
        let receipt = await contract.testRet10(testee_addr, defaultAmountParamsWithoutValue);
        print(receipt);
        count = await contract.count(defaultAmountParamsWithoutValue);
        print(count);
        print("\n");

        count = await contract.countTestThrow(defaultAmountParamsWithoutValue);
        print(count);
        receipt = await contract.testRetThrow(testee_addr, defaultAmountParamsWithoutValue);
        print(receipt);
        count = await contract.countTestThrow(defaultAmountParamsWithoutValue);
        print(count);
        print("\n");

    } catch (err) {
        print("Error profile");
        print(err.toString().split("\n")[0]);
        print(err);
    }
}

async function case5() {
    let name = "MyCaller3";
    let caller = await util.contract(name);
    let cbalance = await BalanceOf(caller.address);
    print("\x1B[33mBefore SendEther\x1B[0m");
    print("cbalance:%s", cbalance);
    let callee = "0xC908B1770FddbE522Dd8c808BCb823F1e0B6561d";
    let ebalance = await BalanceOf(callee);
    print("ebalance:%s", ebalance);
    let count = await caller.sendFailedCount(defaultAmountParamsWithoutValue);
    print("%s", count);

    await caller.MySend(callee, defaultAmountParamsWithValue);

    count = await caller.sendFailedCount(defaultAmountParamsWithoutValue);
    print("\x1B[33mAfter SendEther\x1B[0m");
    print("%s", count);
    ebalance = await BalanceOf(callee);
    print("ebalance:%s", ebalance);
    cbalance = await BalanceOf(caller.address);
    print("cbalance:%s", cbalance);
}
async function case6() {
    print("\x1B[36mcase6\x1B[0m")
    let name = "MyCaller3";
    let caller = await util.contract(name);
    let cbalance = await BalanceOf(caller.address);
    print("\x1B[33mBefore SendEther\x1B[0m");
    print("cbalance:%s", cbalance);
    let callee = "0xC908B1770FddbE522Dd8c808BCb823F1e0B6561d";
    let ebalance = await BalanceOf(callee);
    print("ebalance:%s", ebalance);
    let count = await caller.sendFailedCount(defaultAmountParamsWithoutValue);
    print("%s", count);
    let msg_data = "0x01";
    await caller.MyCallWithoutValue(callee, msg_data, defaultAmountParamsWithoutValue);

    count = await caller.sendFailedCount(defaultAmountParamsWithoutValue);
    print("\x1B[33mAfter SendEther\x1B[0m");
    print("%s", count);
    ebalance = await BalanceOf(callee);
    print("ebalance:%s", ebalance);
    cbalance = await BalanceOf(caller.address);
    print("cbalance:%s", cbalance);
}
async function case7() {
    let name = "MyCaller3";
    let caller = await util.contract(name);
    let cbalance = await BalanceOf(caller.address);
    print("\x1B[33mBefore SendEther\x1B[0m");
    print("cbalance:%s", cbalance);
    let callee = "0xC908B1770FddbE522Dd8c808BCb823F1e0B6561d";
    let ebalance = await BalanceOf(callee);
    print("ebalance:%s", ebalance);
    let count = await caller.sendFailedCount(defaultAmountParamsWithoutValue);
    print("%s", count);
    let msg_data = "0x01";
    await caller.MyCallWithValue(callee, msg_data, defaultAmountParamsWithValue);

    count = await caller.sendFailedCount(defaultAmountParamsWithoutValue);
    print("\x1B[33mAfter SendEther\x1B[0m");
    print("%s", count);
    ebalance = await BalanceOf(callee);
    print("ebalance:%s", ebalance);
    cbalance = await BalanceOf(caller.address);
    print("cbalance:%s", cbalance);
}
async function testContract() {
    // await case1();
    // await case2();
    // await case3();
    // await case4();
    await case5();
    // await case6();
    // await case7();
}
async function test() {
    // await testExceptionDisorder()
    // await testGaslessSend()
    // await testContractWithInitialParams();
    await testContract();
}
let __option__ = "__multiple__"; //__single__/__multiple__
let __name__ = "__main__";

let myArgs = process.argv.slice(2);

if (__name__ == "__main__") {
    main();
}
if (__name__ == "__test__") {
    test();
}