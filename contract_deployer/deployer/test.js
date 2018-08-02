#! /local/bin/bnode

import {
    web3
} from "../utils/ContractUtils.js";
import {
    defaultAccount
} from "../utils/Account";
let print = console.log;
// let obj = {
//     "factHash_": "0000000000000000000000000000000000000000000000000000000000000160",
//     "ethAddr_": "0x1a0",
//     "url_": "Mayweather Yes",
//     "feeAccount_": "0x220",
//     "fee_": 608
// };
// let arr = Object.values(obj);
// print(arr);
// print(...arr);
// printObj(...arr);

// function printObj(factHash, ethAddr, url, feeAccount, fee) {
//     print(factHash, ";", ethAddr, url, ";", feeAccount, fee);
// }
// let bytes = "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855";
// let val = web3.fromAscii("20160528");
// print(val);
// val = web3.toAscii(val);
// print(val);
// val = web3.toAscii(bytes);
// print(val);

// let arr = [
//     ["0xf550f6154a1dbd9986b113f6c920c5b3aa52c8d819c8634e50435df70192e265", "0x7a155e107edc9ddb83d5295e57c76915c78c80519c3cd65221e3ff26e56fd4b0", "0x1ccb68b366acc415fe8ed4fc045c7d2c8233acb05656484c5457bd6798becaab", "0x1d64c63d7faf6fd09fff3472eb8e445995f28c8f27875383104e7c8d2c6ff6de", "0x45476abf6bd2df4c9f91c8d91dd00ce4201fca866fc6ddec76bf73cd5f93c168"],
//     ["0x6ad2eda6b9474b63c8a23d3635d0882e26747a18b916c01a2d1b53185fd49524", "0x6af140f0eb5caef52c52883f37d50e05e90d622871ab8bab72eba8eb4bd8b749", "0x269e0795f4c1cd6672a84a0934e2a9bffb834c5f9ee1485ffb210667d25f8a3f", "0xb5842a2d785c8c82f82fc1280ee499b00e55274ad2d59bf4e799544d09e30293", "0x8ea8d05aed27088ceb74791e3a86613558c699ad0f8ecba2b503b0c73826a353"]
// ]
// // print(arr);
// print(...arr);

// function show(arr1, arr2) {
//     print("arr1:", arr1);
//     print("arr2:", arr2);
// }
// show(...arr)

// let funs = ["getAddress()"];
// for (let fun of funs) {
//     let hash = web3.sha3(fun);
//     print(hash);
// }
var gasPrice = web3.eth.gasPrice;
console.log(gasPrice.toString(10));
var gasLimit;
gasLimit = web3.eth.getBlock(0).gasLimit;
console.log(gasLimit);
gasLimit = web3.eth.getBlock("latest").gasLimit;
console.log(gasLimit);
var block = web3.eth.getBlock("latest");
block = web3.eth.getBlock(block.number - 1);
console.log(block);
block = web3.eth.getBlock(block.number - 1);
console.log(block);
block = web3.eth.getBlock(block.number - 1);
console.log(block);
block = web3.eth.getBlock(block.number - 1);
console.log(block);
block = web3.eth.getBlock(block.number - 1);
console.log(block);
// for (var i = 0; i < 100; i++) {
//     setTimeout(function (i) {
//         console.log(i);
//     }, (i + 1) * 1000, i);
// }
var Balance = web3.eth.getBalance(defaultAccount);
console.log(Balance);