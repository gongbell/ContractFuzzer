#! /local/bin/bnode

import * as fs from "fs";
import {
    ContractUtils,
    defaultAmountParamsWithValue,
    defaultAmountParamsWithoutValue,
    Promise,
    defaultAccount,
    Account1,
    accounts,
    defaultGas
} from "./ContractUtils.js";
import { type } from "os";
const truffle_Contract = require("truffle-contract");

const assert = require("assert");
const Web3 = require("web3");
let web3;
let Provider;

const fread = fs.readFileSync;
const fopen = fs.openFileSync;
const fwrite = fs.writeFileSync;
const print = console.log;

//SET Account and Value in each Transaction;
let ACCOUNT;
let VALUE;
let BYAGENT;
if (!Array.prototype.shuffle) {
    Array.prototype.shuffle = function () {
        for (var j, x, i = this.length; i; j = parseInt(Math.random() * i), x = this[--i], this[i] = this[j], this[j] = x);
        return this;
    };
}

function sendBatchTransaction(transactions) {
    const sendTransaction = Promise.promisify(web3.eth.sendTransaction);
    for (let transaction of transactions) {
        sendTransaction(transaction).catch(function(err){
            //do nothing but output err msg
            console.log(err);            
        });
    }
}
function MyCallWithValueBatch(args){
    for (let arg of args) {
        arg.Caller.MyCallWithValue(arg.contract_addr, arg.msg_data,{from:defaultAccount, value:arg.value,  gas:defaultAmountParamsWithValue.gas}).catch(function(err){
          console.log(err);
        });
    }
}

async function getAgent(){
    let name = "Agent";
    let address = "0xe930e50b62af818dbc955f345f9a3a3108f7a70d";
    let abi = JSON.parse('[{"constant":true,"inputs":[],"name":"count","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"hasValue","outputs":[{"name":"","type":"bool"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"sendFailedCount","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"contract_addr","type":"address"}],"name":"MySend","outputs":[],"payable":true,"type":"function"},{"constant":false,"inputs":[{"name":"contract_addr","type":"address"},{"name":"msg_data","type":"bytes"}],"name":"MyCallWithoutValue","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"call_contract_addr","outputs":[{"name":"","type":"address"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"sendCount","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"contract_addr","type":"address"},{"name":"msg_data","type":"bytes"}],"name":"MyCallWithValue","outputs":[],"payable":true,"type":"function"},{"constant":true,"inputs":[],"name":"call_msg_data","outputs":[{"name":"","type":"bytes"}],"payable":false,"type":"function"},{"constant":false,"inputs":[],"name":"getContractAddr","outputs":[{"name":"addr","type":"address"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"turnoff","outputs":[{"name":"","type":"bool"}],"payable":false,"type":"function"},{"constant":false,"inputs":[],"name":"getCallMsgData","outputs":[{"name":"msg_data","type":"bytes"}],"payable":false,"type":"function"},{"inputs":[],"type":"constructor"},{"payable":true,"type":"fallback"}]');
    let code = "606060405260006000600050556001600360006101000a81548160ff021916908302179055506000600360016101000a81548160ff02191690830217905550600060046000505560006005600050555b5b610a508061005e6000396000f3606060405236156100b6576000357c01000000000000000000000000000000000000000000000000000000009004806306661abd146101de57806315140bd1146102065780633f948cac1461023057806348cccce9146102585780635aa945a4146102705780636b66ae0e146102d45780636ed65dae14610312578063789d1c5c1461033a57806383a64c1e146103995780639e455939146104195780639eeb30e614610457578063d3e204d714610481576100b6565b6101dc5b600360009054906101000a900460ff16156101bf5760006000818150548092919060010191905055506000600360006101000a81548160ff02191690830217905550600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166002600050604051808280546001816001161561010002031660029004801561019f5780601f106101745761010080835404028352916020019161019f565b820191906000526020600020905b81548152906001019060200180831161018257829003601f168201915b50509150506000604051808303816000866161da5a03f1915050506101d9565b6001600360006101000a81548160ff021916908302179055505b5b565b005b34610002576101f06004805050610501565b6040518082815260200191505060405180910390f35b3461000257610218600480505061050a565b60405180821515815260200191505060405180910390f35b3461000257610242600480505061051d565b6040518082815260200191505060405180910390f35b61026e6004808035906020019091905050610526565b005b34610002576102d26004808035906020019091908035906020019082018035906020019191908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050909091905050610591565b005b34610002576102e66004805050610706565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3461000257610324600480505061072c565b6040518082815260200191505060405180910390f35b6103976004808035906020019091908035906020019082018035906020019191908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050909091905050610735565b005b34610002576103ab60048050506108b1565b60405180806020018281038252838181518152602001915080519060200190808383829060006004602084601f0104600302600f01f150905090810190601f16801561040b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b346100025761042b6004805050610952565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34610002576104696004805050610981565b60405180821515815260200191505060405180910390f35b34610002576104936004805050610994565b60405180806020018281038252838181518152602001915080519060200190808383829060006004602084601f0104600302600f01f150905090810190601f1680156104f35780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60006000505481565b600360019054906101000a900460ff1681565b60056000505481565b60046000818150548092919060010191905055508073ffffffffffffffffffffffffffffffffffffffff166108fc349081150290604051809050600060405180830381858888f19350505050151561058d5760056000818150548092919060010191905055505b5b50565b6000600360016101000a81548160ff0219169083021790555081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908302179055508060026000509080519060200190828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061062457805160ff1916838001178555610655565b82800160010185558215610655579182015b82811115610654578251826000505591602001919060010190610636565b5b5090506106809190610662565b8082111561067c5760008181506000905550600101610662565b5090565b50508173ffffffffffffffffffffffffffffffffffffffff1681604051808280519060200190808383829060006004602084601f0104600302600f01f150905090810190601f1680156106e75780820380516001836020036101000a031916815260200191505b509150506000604051808303816000866161da5a03f1915050505b5050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60046000505481565b60006001600360016101000a81548160ff0219169083021790555034905082600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908302179055508160026000509080519060200190828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106107cd57805160ff19168380011785556107fe565b828001600101855582156107fe579182015b828111156107fd5782518260005055916020019190600101906107df565b5b509050610829919061080b565b80821115610825576000818150600090555060010161080b565b5090565b50508273ffffffffffffffffffffffffffffffffffffffff168183604051808280519060200190808383829060006004602084601f0104600302600f01f150905090810190601f1680156108915780820380516001836020036101000a031916815260200191505b5091505060006040518083038185876185025a03f192505050505b505050565b60026000508054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561094a5780601f1061091f5761010080835404028352916020019161094a565b820191906000526020600020905b81548152906001019060200180831161092d57829003601f168201915b505050505081565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905061097e565b90565b600360009054906101000a900460ff1681565b602060405190810160405280600081526020015060026000508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a415780601f10610a1657610100808354040283529160200191610a41565b820191906000526020600020905b815481529060010190602001808311610a2457829003601f168201915b50505050509050610a4d565b9056";
    let MyContract = truffle_Contract({
        contract_name: name,
        abi: abi,
        unlinked_binary: code,
        network_id: 1900,
        address: address,
        default_network: 1900
    });
    MyContract.setProvider(Provider);
    let Caller3 = await MyContract.deployed();
    return Caller3;
}
function getOwner(){
    return defaultAccount;
}
function getNormal(){
    return Account1;
}
let Agent;
let Owner;
let Normal;
let Robin=[0,10000000,10000000000];
const Robin_Index = 3
let Robin_no = 0;
function go(address,msg_group){
    let argsAgent = [];
    let args = [];
    let value = VALUE;
    for (let index=0;index<msg_group.length;index++){
        Robin_no++;
        value = Robin[Robin_no%Robin_Index];
        argsAgent.push({
            Caller: Agent,
            contract_addr: address,
            msg_data: msg_group[index],
            value:value
        });
        value = Robin[(Robin_no%Robin_Index+1)%Robin_Index];
        args.push({
            from: Owner,
            to: address,
            value: value,
            gas: defaultGas,
            data:  msg_group[index]
        });
        value = Robin[(Robin_no%Robin_Index+1)%Robin_Index];
        args.push({
            from: Normal,
            to: address,
            value: value,
            gas: defaultGas,
            data:  msg_group[index]
        });
    }
    MyCallWithValueBatch(argsAgent);
    sendBatchTransaction(args);
}

function RunnerMonitor(){
    const port = 8088;
    const http = require('http');
    const url = require('url');
    try{
    http.createServer(function (request, response) {
        console.log(request.url);
        let obj = url.parse(request.url,true);
        console.log(obj.query);
        let address = obj.query["address"];
        let msg_group = obj.query["msg"];

        if (address!=undefined || msg_group!=undefined){
            if (!(msg_group instanceof Array))
                msg_group = [msg_group]
            go(address,msg_group);
        }
        response.writeHead(200, {'Content-Type': 'text/plain'});
        // 发送响应数据 "Hello World"
	    response.end('Hello World\n');
    }).listen(port);
    }catch(err){
        console.log(err);
    }
    // 终端打印如下信息
    console.log('Monitor running at http://127.0.0.1:8088/');
}
function parse_cmd() {
    let args = process.argv.slice(2,process.argv.length);
    console.log(args);
    if (args.length==6){
        let i=0;
        while(i<6){
           if (args[i].indexOf("--gethrpcport")==0){
            let httpRpcAddr = args[i+1]; 
            console.log(httpRpcAddr);
            Provider = new Web3.providers.HttpProvider(httpRpcAddr);
            web3 = new Web3(Provider);
            web3.personal.unlockAccount(defaultAccount, "123456", 200 * 60 * 60);
            web3.personal.unlockAccount(Account1, "123456", 200 * 60 * 60);
            // console.log(web3);    
          }
          if (args[i].indexOf("--account")==0){
            ACCOUNT = accounts[parseInt(args[i+1])]; 
          }
          if (args[i].indexOf("--value")==0){
            VALUE = parseInt(args[i+1]); 
          }
          if (args[i].indexOf("--Agent")==0){
            BYAGENT = args[i+1]; 
          }
          i += 2;
        }
    }else{
        process.exit(-1);
    }
}
async function Running(){
    parse_cmd();
    Agent = await getAgent();
    Owner = getOwner();
    Normal = getNormal();
    RunnerMonitor();
}
Running();