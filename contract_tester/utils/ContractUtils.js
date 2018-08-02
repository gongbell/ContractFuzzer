import * as fs from "fs";
export const Promise = require("bluebird");
export const assert = require("assert");
export const Web3 = require("web3");

const fread = fs.readFileSync;
const writeFilePs = Promise.promisify(fs.writeFile);

export const accounts = [
    "0x2b71cc952c8e3dfe97a696cf5c5b29f8a07de3d8",
    "0xa31a0f4653f62aca35b6e986743c8f4fc6c8f38f",
    "0x6d62f53305d3c247cd856a8a4eaf65518a7030cf",
    "0x04c8862a82faf3fb90b73768c50dc7f23d7d26bd",
    "0x271eab2a8058d1e3a45941f49d5671bb1cee8ca1"
];
export const defaultAccount = accounts[0];
export const Account1 = accounts[1];
export const defaultGas = 800000;
export const defaultValue = 100000;
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
    setProvider(provider) {
        this.Provider = provider;
        this.web3 = new Web3(this.Provider);
    }
    createProvider(httpRpcAddr) {
        return new Web3.providers.HttpProvider(httpRpcAddr);
    }
 
    show() {
        console.log("hello ContractUtils");
    }
    async sendTransactionWithValue(_from, _to, _value, _msgdata) {
        let sendTransaction = Promise.promisify(this.web3.eth.sendTransaction);
        await sendTransaction({
            from: _from,
            to: _to,
            value: _value,
            gas: defaultGas,
            data: _msgdata
        });
    }　
    async sendTransaction(_from, _to, _msgdata) {
        let sendTransaction = Promise.promisify(this.web3.eth.sendTransaction);
        await sendTransaction({
            from: _from,
            to: _to,
            value: 0,
            gas: defaultGas,
            data: _msgdata
        });
    }
}
let util =  new ContractUtils();