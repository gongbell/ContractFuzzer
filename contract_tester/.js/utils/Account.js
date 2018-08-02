"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
var Web3 = require("web3");
var dotenv = require("dotenv");
dotenv.config();

var HttpRpcAddr = process.env.GethHttpRpcAddr;
var Provider = new Web3.providers.HttpProvider(HttpRpcAddr);
var web3 = exports.web3 = new Web3(Provider);
var accounts = exports.accounts = web3.eth.accounts;
var defaultAccount = exports.defaultAccount = web3.eth.accounts[0];
var defaultGas = exports.defaultGas = 8000000;
var defaultValue = exports.defaultValue = 1000000;