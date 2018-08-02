const Web3 = require("web3");
const dotenv = require("dotenv");
dotenv.config();

const HttpRpcAddr = process.env.GethHttpRpcAddr;
const Provider = new Web3.providers.HttpProvider(HttpRpcAddr);
export const web3 = new Web3(Provider);