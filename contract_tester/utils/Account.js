const Web3 = require("web3");
const dotenv = require("dotenv");
dotenv.config();

const HttpRpcAddr = process.env.GethHttpRpcAddr;
const Provider = new Web3.providers.HttpProvider(HttpRpcAddr);
export const web3 = new Web3(Provider);
export const accounts = web3.eth.accounts;
export const defaultAccount = web3.eth.accounts[0];
web3.personal.unlockAccount(defaultAccount, "123456", 200 * 60 * 60);
export const Account1 = web3.eth.accounts[1];
web3.personal.unlockAccount(Account1, "123456", 200 * 60 * 60);
//export const defaultGas = 8000000;
export const defaultGas = 50000000000; //web3.toWei("50000000000 ", "wei");
export const defaultValue = 1000000; // web3.toWei("1000000", "wei");