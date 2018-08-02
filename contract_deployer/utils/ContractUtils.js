import * as fs from "fs";

export const dotenv = require("dotenv");
export const assert = require("assert");
export const Promise = require("bluebird");
export const Web3 = require("web3");
export const solc = require("solc");
export const truffle_Contract = require("truffle-contract");

dotenv.config();


const writeFilePs = Promise.promisify(fs.writeFile);
const HttpRpcAddr = process.env.GethHttpRpcAddr;
console.log(HttpRpcAddr)
export const Provider = new Web3.providers.HttpProvider(HttpRpcAddr);
export const web3 = new Web3(Provider);

const accounts = web3.eth.accounts;
const defaultAccount = web3.eth.accounts[0];
const defaultGas = 800000;
const defaultValue = 100000;
export const defaultAmountParamsWithValue = {
    from: defaultAccount,
    value: defaultValue,
    gas: defaultGas
};
export const defaultAmountParamsWithoutValue = {
    from: defaultAccount,
    gas: defaultGas
};
export
class ContractUtils {
    Password = "123456";
    GasPrice = 8000000;
    constructor() {
    }
    async _read(file) {
        let readFile = Promise.promisify(fs.readFile);
        let res = await readFile(file);
        return res.toString();
    }
    async DeployWithParams(contract, iParamsArr, aParams) {
        aParams.value = web3.toBigNumber(aParams.value);
        aParams.gas = web3.toBigNumber(aParams.gas);
        const owner = this.defaultAccount;
        let contract_name_ = contract.name;
        let abi_ = contract.abi;
        let code_ = contract.bin;

        let MyContract = truffle_Contract({
            contract_name: contract_name_,
            abi: abi_,
            unlinked_binary: code_,
            default_network: 1900
        });
        MyContract.setProvider(Provider);
        let instance = await MyContract.new(...iParamsArr, aParams);
        return instance;
    }
    async DeployWithoutParams(contract, aParams) {
        aParams.value = web3.toBigNumber(aParams.value);
        aParams.gas = web3.toBigNumber(aParams.gas);
        const owner = this.defaultAccount;
        let contract_name_ = contract.name;
        let abi_ = contract.abi;
        let code_ = contract.bin;

        let MyContract = truffle_Contract({
            contract_name: contract_name_,
            abi: abi_,
            network_id: 1900,
            unlinked_binary: code_,
            default_network: 1900
        });
        MyContract.setProvider(Provider);
        let instance = await MyContract.new(aParams);
        return instance;
    }
    async _instance(contract) {
        let name_ = contract.name;
        let abi_ = JSON.parse(unescape(contract.abi));
        let code_ = unescape(contract.code);
        let address_ = contract.address;
        let MyContract = truffle_Contract({
            contract_name: name_,
            abi: abi_,
            unlinked_binary: code_,
            network_id: 1900,
            address: address_,
            default_network: 1900
        });
        MyContract.setProvider(Provider);
        let instance = await MyContract.deployed();
        return instance;
    }
    show() {
        console.log("hello ContractUtils");
    }
}
//Use Case
//const utils = new ContractUtils();
//utils.test_promise();
//utils._read("./Contracts/SendBalance_V3.sol");