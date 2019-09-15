pragma solidity ^0.4.17;


/********************************* Oraclize API ************************************************/
// The Oraclize API has been put into the same file to make Etherscan verification easier.
// FOR THE TOKEN CONTRACTS SCROLL DOWN TO ABOUT LINE 1000


// <ORACLIZE_API>
/*
Copyright (c) 2015-2016 Oraclize SRL
Copyright (c) 2016 Oraclize LTD



Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:



The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.



THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

contract OraclizeI {
    address public cbAddress;
    function query(uint _timestamp, string _datasource, string _arg) payable returns (bytes32 _id);
    function query_withGasLimit(uint _timestamp, string _datasource, string _arg, uint _gaslimit) payable returns (bytes32 _id);
    function query2(uint _timestamp, string _datasource, string _arg1, string _arg2) payable returns (bytes32 _id);
    function query2_withGasLimit(uint _timestamp, string _datasource, string _arg1, string _arg2, uint _gaslimit) payable returns (bytes32 _id);
    function queryN(uint _timestamp, string _datasource, bytes _argN) payable returns (bytes32 _id);
    function queryN_withGasLimit(uint _timestamp, string _datasource, bytes _argN, uint _gaslimit) payable returns (bytes32 _id);
    function getPrice(string _datasource) returns (uint _dsprice);
    function getPrice(string _datasource, uint gaslimit) returns (uint _dsprice);
    function useCoupon(string _coupon);
    function setProofType(byte _proofType);
    function setConfig(bytes32 _config);
    function setCustomGasPrice(uint _gasPrice);
    function randomDS_getSessionPubKeyHash() returns(bytes32);
}
contract OraclizeAddrResolverI {
    function getAddress() returns (address _addr);
}
contract usingOraclize {
    uint constant day = 60*60*24;
    uint constant week = 60*60*24*7;
    uint constant month = 60*60*24*30;
    byte constant proofType_NONE = 0x00;
    byte constant proofType_TLSNotary = 0x10;
    byte constant proofType_Android = 0x20;
    byte constant proofType_Ledger = 0x30;
    byte constant proofType_Native = 0xF0;
    byte constant proofStorage_IPFS = 0x01;
    uint8 constant networkID_auto = 0;
    uint8 constant networkID_mainnet = 1;
    uint8 constant networkID_testnet = 2;
    uint8 constant networkID_morden = 2;
    uint8 constant networkID_consensys = 161;

    OraclizeAddrResolverI OAR;

    OraclizeI oraclize;
    modifier oraclizeAPI {
        if((address(OAR)==0)||(getCodeSize(address(OAR))==0)) oraclize_setNetwork(networkID_auto);
        oraclize = OraclizeI(OAR.getAddress());
        _;
    }
    modifier coupon(string code){
        oraclize = OraclizeI(OAR.getAddress());
        oraclize.useCoupon(code);
        _;
    }

    function oraclize_setNetwork(uint8 networkID) internal returns(bool){
        if (getCodeSize(0x1d3B2638a7cC9f2CB3D298A3DA7a90B67E5506ed)>0){ //mainnet
            OAR = OraclizeAddrResolverI(0x1d3B2638a7cC9f2CB3D298A3DA7a90B67E5506ed);
            oraclize_setNetworkName("eth_mainnet");
            return true;
        }
        if (getCodeSize(0xc03A2615D5efaf5F49F60B7BB6583eaec212fdf1)>0){ //ropsten testnet
            OAR = OraclizeAddrResolverI(0xc03A2615D5efaf5F49F60B7BB6583eaec212fdf1);
            oraclize_setNetworkName("eth_ropsten3");
            return true;
        }
        if (getCodeSize(0xB7A07BcF2Ba2f2703b24C0691b5278999C59AC7e)>0){ //kovan testnet
            OAR = OraclizeAddrResolverI(0xB7A07BcF2Ba2f2703b24C0691b5278999C59AC7e);
            oraclize_setNetworkName("eth_kovan");
            return true;
        }
        if (getCodeSize(0x146500cfd35B22E4A392Fe0aDc06De1a1368Ed48)>0){ //rinkeby testnet
            OAR = OraclizeAddrResolverI(0x146500cfd35B22E4A392Fe0aDc06De1a1368Ed48);
            oraclize_setNetworkName("eth_rinkeby");
            return true;
        }
        if (getCodeSize(0x6f485C8BF6fc43eA212E93BBF8ce046C7f1cb475)>0){ //ethereum-bridge
            OAR = OraclizeAddrResolverI(0x6f485C8BF6fc43eA212E93BBF8ce046C7f1cb475);
            return true;
        }
        if (getCodeSize(0x20e12A1F859B3FeaE5Fb2A0A32C18F5a65555bBF)>0){ //ether.camp ide
            OAR = OraclizeAddrResolverI(0x20e12A1F859B3FeaE5Fb2A0A32C18F5a65555bBF);
            return true;
        }
        if (getCodeSize(0x51efaF4c8B3C9AfBD5aB9F4bbC82784Ab6ef8fAA)>0){ //browser-solidity
            OAR = OraclizeAddrResolverI(0x51efaF4c8B3C9AfBD5aB9F4bbC82784Ab6ef8fAA);
            return true;
        }
        return false;
    }

    function __callback(bytes32 myid, string result) {
        __callback(myid, result, new bytes(0));
    }
    function __callback(bytes32 myid, string result, bytes proof) {
    }

    function oraclize_useCoupon(string code) oraclizeAPI internal {
        oraclize.useCoupon(code);
    }

    function oraclize_getPrice(string datasource) oraclizeAPI internal returns (uint){
        return oraclize.getPrice(datasource);
    }

    function oraclize_getPrice(string datasource, uint gaslimit) oraclizeAPI internal returns (uint){
        return oraclize.getPrice(datasource, gaslimit);
    }

    function oraclize_query(string datasource, string arg) oraclizeAPI internal returns (bytes32 id){
        uint price = oraclize.getPrice(datasource);
        if (price > 1 ether + tx.gasprice*200000) return 0; // unexpectedly high price
        return oraclize.query.value(price)(0, datasource, arg);
    }
    function oraclize_query(uint timestamp, string datasource, string arg) oraclizeAPI internal returns (bytes32 id){
        uint price = oraclize.getPrice(datasource);
        if (price > 1 ether + tx.gasprice*200000) return 0; // unexpectedly high price
        return oraclize.query.value(price)(timestamp, datasource, arg);
    }
    function oraclize_query(uint timestamp, string datasource, string arg, uint gaslimit) oraclizeAPI internal returns (bytes32 id){
        uint price = oraclize.getPrice(datasource, gaslimit);
        if (price > 1 ether + tx.gasprice*gaslimit) return 0; // unexpectedly high price
        return oraclize.query_withGasLimit.value(price)(timestamp, datasource, arg, gaslimit);
    }
    function oraclize_query(string datasource, string arg, uint gaslimit) oraclizeAPI internal returns (bytes32 id){
        uint price = oraclize.getPrice(datasource, gaslimit);
        if (price > 1 ether + tx.gasprice*gaslimit) return 0; // unexpectedly high price
        return oraclize.query_withGasLimit.value(price)(0, datasource, arg, gaslimit);
    }
    function oraclize_query(string datasource, string arg1, string arg2) oraclizeAPI internal returns (bytes32 id){
        uint price = oraclize.getPrice(datasource);
        if (price > 1 ether + tx.gasprice*200000) return 0; // unexpectedly high price
        return oraclize.query2.value(price)(0, datasource, arg1, arg2);
    }
    function oraclize_query(uint timestamp, string datasource, string arg1, string arg2) oraclizeAPI internal returns (bytes32 id){
        uint price = oraclize.getPrice(datasource);
        if (price > 1 ether + tx.gasprice*200000) return 0; // unexpectedly high price
        return oraclize.query2.value(price)(timestamp, datasource, arg1, arg2);
    }
    function oraclize_query(uint timestamp, string datasource, string arg1, string arg2, uint gaslimit) oraclizeAPI internal returns (bytes32 id){
        uint price = oraclize.getPrice(datasource, gaslimit);
        if (price > 1 ether + tx.gasprice*gaslimit) return 0; // unexpectedly high price
        return oraclize.query2_withGasLimit.value(price)(timestamp, datasource, arg1, arg2, gaslimit);
    }
    function oraclize_query(string datasource, string arg1, string arg2, uint gaslimit) oraclizeAPI internal returns (bytes32 id){
        uint price = oraclize.getPrice(datasource, gaslimit);
        if (price > 1 ether + tx.gasprice*gaslimit) return 0; // unexpectedly high price
        return oraclize.query2_withGasLimit.value(price)(0, datasource, arg1, arg2, gaslimit);
    }
    function oraclize_query(string datasource, string[] argN) oraclizeAPI internal returns (bytes32 id){
        uint price = oraclize.getPrice(datasource);
        if (price > 1 ether + tx.gasprice*200000) return 0; // unexpectedly high price
        bytes memory args = stra2cbor(argN);
        return oraclize.queryN.value(price)(0, datasource, args);
    }
    function oraclize_query(uint timestamp, string datasource, string[] argN) oraclizeAPI internal returns (bytes32 id){
        uint price = oraclize.getPrice(datasource);
        if (price > 1 ether + tx.gasprice*200000) return 0; // unexpectedly high price
        bytes memory args = stra2cbor(argN);
        return oraclize.queryN.value(price)(timestamp, datasource, args);
    }
    function oraclize_query(uint timestamp, string datasource, string[] argN, uint gaslimit) oraclizeAPI internal returns (bytes32 id){
        uint price = oraclize.getPrice(datasource, gaslimit);
        if (price > 1 ether + tx.gasprice*gaslimit) return 0; // unexpectedly high price
        bytes memory args = stra2cbor(argN);
        return oraclize.queryN_withGasLimit.value(price)(timestamp, datasource, args, gaslimit);
    }
    function oraclize_query(string datasource, string[] argN, uint gaslimit) oraclizeAPI internal returns (bytes32 id){
        uint price = oraclize.getPrice(datasource, gaslimit);
        if (price > 1 ether + tx.gasprice*gaslimit) return 0; // unexpectedly high price
        bytes memory args = stra2cbor(argN);
        return oraclize.queryN_withGasLimit.value(price)(0, datasource, args, gaslimit);
    }
    function oraclize_query(string datasource, string[1] args) oraclizeAPI internal returns (bytes32 id) {
        string[] memory dynargs = new string[](1);
        dynargs[0] = args[0];
        return oraclize_query(datasource, dynargs);
    }
    function oraclize_query(uint timestamp, string datasource, string[1] args) oraclizeAPI internal returns (bytes32 id) {
        string[] memory dynargs = new string[](1);
        dynargs[0] = args[0];
        return oraclize_query(timestamp, datasource, dynargs);
    }
    function oraclize_query(uint timestamp, string datasource, string[1] args, uint gaslimit) oraclizeAPI internal returns (bytes32 id) {
        string[] memory dynargs = new string[](1);
        dynargs[0] = args[0];
        return oraclize_query(timestamp, datasource, dynargs, gaslimit);
    }
    function oraclize_query(string datasource, string[1] args, uint gaslimit) oraclizeAPI internal returns (bytes32 id) {
        string[] memory dynargs = new string[](1);
        dynargs[0] = args[0];
        return oraclize_query(datasource, dynargs, gaslimit);
    }

    function oraclize_query(string datasource, string[2] args) oraclizeAPI internal returns (bytes32 id) {
        string[] memory dynargs = new string[](2);
        dynargs[0] = args[0];
        dynargs[1] = args[1];
        return oraclize_query(datasource, dynargs);
    }
    function oraclize_query(uint timestamp, string datasource, string[2] args) oraclizeAPI internal returns (bytes32 id) {
        string[] memory dynargs = new string[](2);
        dynargs[0] = args[0];
        dynargs[1] = args[1];
        return oraclize_query(timestamp, datasource, dynargs);
    }
    function oraclize_query(uint timestamp, string datasource, string[2] args, uint gaslimit) oraclizeAPI internal returns (bytes32 id) {
        string[] memory dynargs = new string[](2);
        dynargs[0] = args[0];
        dynargs[1] = args[1];
        return oraclize_query(timestamp, datasource, dynargs, gaslimit);
    }
    function oraclize_query(string datasource, string[2] args, uint gaslimit) oraclizeAPI internal returns (bytes32 id) {
        string[] memory dynargs = new string[](2);
        dynargs[0] = args[0];
        dynargs[1] = args[1];
        return oraclize_query(datasource, dynargs, gaslimit);
    }
    function oraclize_query(string datasource, string[3] args) oraclizeAPI internal returns (bytes32 id) {
        string[] memory dynargs = new string[](3);
        dynargs[0] = args[0];
        dynargs[1] = args[1];
        dynargs[2] = args[2];
        return oraclize_query(datasource, dynargs);
    }
    function oraclize_query(uint timestamp, string datasource, string[3] args) oraclizeAPI internal returns (bytes32 id) {
        string[] memory dynargs = new string[](3);
        dynargs[0] = args[0];
        dynargs[1] = args[1];
        dynargs[2] = args[2];
        return oraclize_query(timestamp, datasource, dynargs);
    }
    function oraclize_query(uint timestamp, string datasource, string[3] args, uint gaslimit) oraclizeAPI internal returns (bytes32 id) {
        string[] memory dynargs = new string[](3);
        dynargs[0] = args[0];
        dynargs[1] = args[1];
        dynargs[2] = args[2];
        return oraclize_query(timestamp, datasource, dynargs, gaslimit);
    }
    function oraclize_query(string datasource, string[3] args, uint gaslimit) oraclizeAPI internal returns (bytes32 id) {
        string[] memory dynargs = new string[](3);
        dynargs[0] = args[0];
        dynargs[1] = args[1];
        dynargs[2] = args[2];
        return oraclize_query(datasource, dynargs, gaslimit);
    }

    function oraclize_query(string datasource, string[4] args) oraclizeAPI internal returns (bytes32 id) {
        string[] memory dynargs = new string[](4);
        dynargs[0] = args[0];
        dynargs[1] = args[1];
        dynargs[2] = args[2];
        dynargs[3] = args[3];
        return oraclize_query(datasource, dynargs);
    }
    function oraclize_query(uint timestamp, string datasource, string[4] args) oraclizeAPI internal returns (bytes32 id) {
        string[] memory dynargs = new string[](4);
        dynargs[0] = args[0];
        dynargs[1] = args[1];
        dynargs[2] = args[2];
        dynargs[3] = args[3];
        return oraclize_query(timestamp, datasource, dynargs);
    }
    function oraclize_query(uint timestamp, string datasource, string[4] args, uint gaslimit) oraclizeAPI internal returns (bytes32 id) {
        string[] memory dynargs = new string[](4);
        dynargs[0] = args[0];
        dynargs[1] = args[1];
        dynargs[2] = args[2];
        dynargs[3] = args[3];
        return oraclize_query(timestamp, datasource, dynargs, gaslimit);
    }
    function oraclize_query(string datasource, string[4] args, uint gaslimit) oraclizeAPI internal returns (bytes32 id) {
        string[] memory dynargs = new string[](4);
        dynargs[0] = args[0];
        dynargs[1] = args[1];
        dynargs[2] = args[2];
        dynargs[3] = args[3];
        return oraclize_query(datasource, dynargs, gaslimit);
    }
    function oraclize_query(string datasource, string[5] args) oraclizeAPI internal returns (bytes32 id) {
        string[] memory dynargs = new string[](5);
        dynargs[0] = args[0];
        dynargs[1] = args[1];
        dynargs[2] = args[2];
        dynargs[3] = args[3];
        dynargs[4] = args[4];
        return oraclize_query(datasource, dynargs);
    }
    function oraclize_query(uint timestamp, string datasource, string[5] args) oraclizeAPI internal returns (bytes32 id) {
        string[] memory dynargs = new string[](5);
        dynargs[0] = args[0];
        dynargs[1] = args[1];
        dynargs[2] = args[2];
        dynargs[3] = args[3];
        dynargs[4] = args[4];
        return oraclize_query(timestamp, datasource, dynargs);
    }
    function oraclize_query(uint timestamp, string datasource, string[5] args, uint gaslimit) oraclizeAPI internal returns (bytes32 id) {
        string[] memory dynargs = new string[](5);
        dynargs[0] = args[0];
        dynargs[1] = args[1];
        dynargs[2] = args[2];
        dynargs[3] = args[3];
        dynargs[4] = args[4];
        return oraclize_query(timestamp, datasource, dynargs, gaslimit);
    }
    function oraclize_query(string datasource, string[5] args, uint gaslimit) oraclizeAPI internal returns (bytes32 id) {
        string[] memory dynargs = new string[](5);
        dynargs[0] = args[0];
        dynargs[1] = args[1];
        dynargs[2] = args[2];
        dynargs[3] = args[3];
        dynargs[4] = args[4];
        return oraclize_query(datasource, dynargs, gaslimit);
    }
    function oraclize_query(string datasource, bytes[] argN) oraclizeAPI internal returns (bytes32 id){
        uint price = oraclize.getPrice(datasource);
        if (price > 1 ether + tx.gasprice*200000) return 0; // unexpectedly high price
        bytes memory args = ba2cbor(argN);
        return oraclize.queryN.value(price)(0, datasource, args);
    }
    function oraclize_query(uint timestamp, string datasource, bytes[] argN) oraclizeAPI internal returns (bytes32 id){
        uint price = oraclize.getPrice(datasource);
        if (price > 1 ether + tx.gasprice*200000) return 0; // unexpectedly high price
        bytes memory args = ba2cbor(argN);
        return oraclize.queryN.value(price)(timestamp, datasource, args);
    }
    function oraclize_query(uint timestamp, string datasource, bytes[] argN, uint gaslimit) oraclizeAPI internal returns (bytes32 id){
        uint price = oraclize.getPrice(datasource, gaslimit);
        if (price > 1 ether + tx.gasprice*gaslimit) return 0; // unexpectedly high price
        bytes memory args = ba2cbor(argN);
        return oraclize.queryN_withGasLimit.value(price)(timestamp, datasource, args, gaslimit);
    }
    function oraclize_query(string datasource, bytes[] argN, uint gaslimit) oraclizeAPI internal returns (bytes32 id){
        uint price = oraclize.getPrice(datasource, gaslimit);
        if (price > 1 ether + tx.gasprice*gaslimit) return 0; // unexpectedly high price
        bytes memory args = ba2cbor(argN);
        return oraclize.queryN_withGasLimit.value(price)(0, datasource, args, gaslimit);
    }
    function oraclize_query(string datasource, bytes[1] args) oraclizeAPI internal returns (bytes32 id) {
        bytes[] memory dynargs = new bytes[](1);
        dynargs[0] = args[0];
        return oraclize_query(datasource, dynargs);
    }
    function oraclize_query(uint timestamp, string datasource, bytes[1] args) oraclizeAPI internal returns (bytes32 id) {
        bytes[] memory dynargs = new bytes[](1);
        dynargs[0] = args[0];
        return oraclize_query(timestamp, datasource, dynargs);
    }
    function oraclize_query(uint timestamp, string datasource, bytes[1] args, uint gaslimit) oraclizeAPI internal returns (bytes32 id) {
        bytes[] memory dynargs = new bytes[](1);
        dynargs[0] = args[0];
        return oraclize_query(timestamp, datasource, dynargs, gaslimit);
    }
    function oraclize_query(string datasource, bytes[1] args, uint gaslimit) oraclizeAPI internal returns (bytes32 id) {
        bytes[] memory dynargs = new bytes[](1);
        dynargs[0] = args[0];
        return oraclize_query(datasource, dynargs, gaslimit);
    }

    function oraclize_query(string datasource, bytes[2] args) oraclizeAPI internal returns (bytes32 id) {
        bytes[] memory dynargs = new bytes[](2);
        dynargs[0] = args[0];
        dynargs[1] = args[1];
        return oraclize_query(datasource, dynargs);
    }
    function oraclize_query(uint timestamp, string datasource, bytes[2] args) oraclizeAPI internal returns (bytes32 id) {
        bytes[] memory dynargs = new bytes[](2);
        dynargs[0] = args[0];
        dynargs[1] = args[1];
        return oraclize_query(timestamp, datasource, dynargs);
    }
    function oraclize_query(uint timestamp, string datasource, bytes[2] args, uint gaslimit) oraclizeAPI internal returns (bytes32 id) {
        bytes[] memory dynargs = new bytes[](2);
        dynargs[0] = args[0];
        dynargs[1] = args[1];
        return oraclize_query(timestamp, datasource, dynargs, gaslimit);
    }
    function oraclize_query(string datasource, bytes[2] args, uint gaslimit) oraclizeAPI internal returns (bytes32 id) {
        bytes[] memory dynargs = new bytes[](2);
        dynargs[0] = args[0];
        dynargs[1] = args[1];
        return oraclize_query(datasource, dynargs, gaslimit);
    }
    function oraclize_query(string datasource, bytes[3] args) oraclizeAPI internal returns (bytes32 id) {
        bytes[] memory dynargs = new bytes[](3);
        dynargs[0] = args[0];
        dynargs[1] = args[1];
        dynargs[2] = args[2];
        return oraclize_query(datasource, dynargs);
    }
    function oraclize_query(uint timestamp, string datasource, bytes[3] args) oraclizeAPI internal returns (bytes32 id) {
        bytes[] memory dynargs = new bytes[](3);
        dynargs[0] = args[0];
        dynargs[1] = args[1];
        dynargs[2] = args[2];
        return oraclize_query(timestamp, datasource, dynargs);
    }
    function oraclize_query(uint timestamp, string datasource, bytes[3] args, uint gaslimit) oraclizeAPI internal returns (bytes32 id) {
        bytes[] memory dynargs = new bytes[](3);
        dynargs[0] = args[0];
        dynargs[1] = args[1];
        dynargs[2] = args[2];
        return oraclize_query(timestamp, datasource, dynargs, gaslimit);
    }
    function oraclize_query(string datasource, bytes[3] args, uint gaslimit) oraclizeAPI internal returns (bytes32 id) {
        bytes[] memory dynargs = new bytes[](3);
        dynargs[0] = args[0];
        dynargs[1] = args[1];
        dynargs[2] = args[2];
        return oraclize_query(datasource, dynargs, gaslimit);
    }

    function oraclize_query(string datasource, bytes[4] args) oraclizeAPI internal returns (bytes32 id) {
        bytes[] memory dynargs = new bytes[](4);
        dynargs[0] = args[0];
        dynargs[1] = args[1];
        dynargs[2] = args[2];
        dynargs[3] = args[3];
        return oraclize_query(datasource, dynargs);
    }
    function oraclize_query(uint timestamp, string datasource, bytes[4] args) oraclizeAPI internal returns (bytes32 id) {
        bytes[] memory dynargs = new bytes[](4);
        dynargs[0] = args[0];
        dynargs[1] = args[1];
        dynargs[2] = args[2];
        dynargs[3] = args[3];
        return oraclize_query(timestamp, datasource, dynargs);
    }
    function oraclize_query(uint timestamp, string datasource, bytes[4] args, uint gaslimit) oraclizeAPI internal returns (bytes32 id) {
        bytes[] memory dynargs = new bytes[](4);
        dynargs[0] = args[0];
        dynargs[1] = args[1];
        dynargs[2] = args[2];
        dynargs[3] = args[3];
        return oraclize_query(timestamp, datasource, dynargs, gaslimit);
    }
    function oraclize_query(string datasource, bytes[4] args, uint gaslimit) oraclizeAPI internal returns (bytes32 id) {
        bytes[] memory dynargs = new bytes[](4);
        dynargs[0] = args[0];
        dynargs[1] = args[1];
        dynargs[2] = args[2];
        dynargs[3] = args[3];
        return oraclize_query(datasource, dynargs, gaslimit);
    }
    function oraclize_query(string datasource, bytes[5] args) oraclizeAPI internal returns (bytes32 id) {
        bytes[] memory dynargs = new bytes[](5);
        dynargs[0] = args[0];
        dynargs[1] = args[1];
        dynargs[2] = args[2];
        dynargs[3] = args[3];
        dynargs[4] = args[4];
        return oraclize_query(datasource, dynargs);
    }
    function oraclize_query(uint timestamp, string datasource, bytes[5] args) oraclizeAPI internal returns (bytes32 id) {
        bytes[] memory dynargs = new bytes[](5);
        dynargs[0] = args[0];
        dynargs[1] = args[1];
        dynargs[2] = args[2];
        dynargs[3] = args[3];
        dynargs[4] = args[4];
        return oraclize_query(timestamp, datasource, dynargs);
    }
    function oraclize_query(uint timestamp, string datasource, bytes[5] args, uint gaslimit) oraclizeAPI internal returns (bytes32 id) {
        bytes[] memory dynargs = new bytes[](5);
        dynargs[0] = args[0];
        dynargs[1] = args[1];
        dynargs[2] = args[2];
        dynargs[3] = args[3];
        dynargs[4] = args[4];
        return oraclize_query(timestamp, datasource, dynargs, gaslimit);
    }
    function oraclize_query(string datasource, bytes[5] args, uint gaslimit) oraclizeAPI internal returns (bytes32 id) {
        bytes[] memory dynargs = new bytes[](5);
        dynargs[0] = args[0];
        dynargs[1] = args[1];
        dynargs[2] = args[2];
        dynargs[3] = args[3];
        dynargs[4] = args[4];
        return oraclize_query(datasource, dynargs, gaslimit);
    }

    function oraclize_cbAddress() oraclizeAPI internal returns (address){
        return oraclize.cbAddress();
    }
    function oraclize_setProof(byte proofP) oraclizeAPI internal {
        return oraclize.setProofType(proofP);
    }
    function oraclize_setCustomGasPrice(uint gasPrice) oraclizeAPI internal {
        return oraclize.setCustomGasPrice(gasPrice);
    }
    function oraclize_setConfig(bytes32 config) oraclizeAPI internal {
        return oraclize.setConfig(config);
    }

    function oraclize_randomDS_getSessionPubKeyHash() oraclizeAPI internal returns (bytes32){
        return oraclize.randomDS_getSessionPubKeyHash();
    }

    function getCodeSize(address _addr) constant internal returns(uint _size) {
        assembly {
            _size := extcodesize(_addr)
        }
    }

    function parseAddr(string _a) internal returns (address){
        bytes memory tmp = bytes(_a);
        uint160 iaddr = 0;
        uint160 b1;
        uint160 b2;
        for (uint i=2; i<2+2*20; i+=2){
            iaddr *= 256;
            b1 = uint160(tmp[i]);
            b2 = uint160(tmp[i+1]);
            if ((b1 >= 97)&&(b1 <= 102)) b1 -= 87;
            else if ((b1 >= 65)&&(b1 <= 70)) b1 -= 55;
            else if ((b1 >= 48)&&(b1 <= 57)) b1 -= 48;
            if ((b2 >= 97)&&(b2 <= 102)) b2 -= 87;
            else if ((b2 >= 65)&&(b2 <= 70)) b2 -= 55;
            else if ((b2 >= 48)&&(b2 <= 57)) b2 -= 48;
            iaddr += (b1*16+b2);
        }
        return address(iaddr);
    }

    function strCompare(string _a, string _b) internal returns (int) {
        bytes memory a = bytes(_a);
        bytes memory b = bytes(_b);
        uint minLength = a.length;
        if (b.length < minLength) minLength = b.length;
        for (uint i = 0; i < minLength; i ++)
            if (a[i] < b[i])
                return -1;
            else if (a[i] > b[i])
                return 1;
        if (a.length < b.length)
            return -1;
        else if (a.length > b.length)
            return 1;
        else
            return 0;
    }

    function indexOf(string _haystack, string _needle) internal returns (int) {
        bytes memory h = bytes(_haystack);
        bytes memory n = bytes(_needle);
        if(h.length < 1 || n.length < 1 || (n.length > h.length))
            return -1;
        else if(h.length > (2**128 -1))
            return -1;
        else
        {
            uint subindex = 0;
            for (uint i = 0; i < h.length; i ++)
            {
                if (h[i] == n[0])
                {
                    subindex = 1;
                    while(subindex < n.length && (i + subindex) < h.length && h[i + subindex] == n[subindex])
                    {
                        subindex++;
                    }
                    if(subindex == n.length)
                        return int(i);
                }
            }
            return -1;
        }
    }

    function strConcat(string _a, string _b, string _c, string _d, string _e) internal returns (string) {
        bytes memory _ba = bytes(_a);
        bytes memory _bb = bytes(_b);
        bytes memory _bc = bytes(_c);
        bytes memory _bd = bytes(_d);
        bytes memory _be = bytes(_e);
        string memory abcde = new string(_ba.length + _bb.length + _bc.length + _bd.length + _be.length);
        bytes memory babcde = bytes(abcde);
        uint k = 0;
        for (uint i = 0; i < _ba.length; i++) babcde[k++] = _ba[i];
        for (i = 0; i < _bb.length; i++) babcde[k++] = _bb[i];
        for (i = 0; i < _bc.length; i++) babcde[k++] = _bc[i];
        for (i = 0; i < _bd.length; i++) babcde[k++] = _bd[i];
        for (i = 0; i < _be.length; i++) babcde[k++] = _be[i];
        return string(babcde);
    }

    function strConcat(string _a, string _b, string _c, string _d) internal returns (string) {
        return strConcat(_a, _b, _c, _d, "");
    }

    function strConcat(string _a, string _b, string _c) internal returns (string) {
        return strConcat(_a, _b, _c, "", "");
    }

    function strConcat(string _a, string _b) internal returns (string) {
        return strConcat(_a, _b, "", "", "");
    }

    // parseInt
    function parseInt(string _a) internal returns (uint) {
        return parseInt(_a, 0);
    }

    // parseInt(parseFloat*10^_b)
    function parseInt(string _a, uint _b) internal returns (uint) {
        bytes memory bresult = bytes(_a);
        uint mint = 0;
        bool decimals = false;
        for (uint i=0; i<bresult.length; i++){
            if ((bresult[i] >= 48)&&(bresult[i] <= 57)){
                if (decimals){
                   if (_b == 0) break;
                    else _b--;
                }
                mint *= 10;
                mint += uint(bresult[i]) - 48;
            } else if (bresult[i] == 46) decimals = true;
        }
        if (_b > 0) mint *= 10**_b;
        return mint;
    }

    function uint2str(uint i) internal returns (string){
        if (i == 0) return "0";
        uint j = i;
        uint len;
        while (j != 0){
            len++;
            j /= 10;
        }
        bytes memory bstr = new bytes(len);
        uint k = len - 1;
        while (i != 0){
            bstr[k--] = byte(48 + i % 10);
            i /= 10;
        }
        return string(bstr);
    }

    function stra2cbor(string[] arr) internal returns (bytes) {
            uint arrlen = arr.length;

            // get correct cbor output length
            uint outputlen = 0;
            bytes[] memory elemArray = new bytes[](arrlen);
            for (uint i = 0; i < arrlen; i++) {
                elemArray[i] = (bytes(arr[i]));
                outputlen += elemArray[i].length + (elemArray[i].length - 1)/23 + 3; //+3 accounts for paired identifier types
            }
            uint ctr = 0;
            uint cborlen = arrlen + 0x80;
            outputlen += byte(cborlen).length;
            bytes memory res = new bytes(outputlen);

            while (byte(cborlen).length > ctr) {
                res[ctr] = byte(cborlen)[ctr];
                ctr++;
            }
            for (i = 0; i < arrlen; i++) {
                res[ctr] = 0x5F;
                ctr++;
                for (uint x = 0; x < elemArray[i].length; x++) {
                    // if there's a bug with larger strings, this may be the culprit
                    if (x % 23 == 0) {
                        uint elemcborlen = elemArray[i].length - x >= 24 ? 23 : elemArray[i].length - x;
                        elemcborlen += 0x40;
                        uint lctr = ctr;
                        while (byte(elemcborlen).length > ctr - lctr) {
                            res[ctr] = byte(elemcborlen)[ctr - lctr];
                            ctr++;
                        }
                    }
                    res[ctr] = elemArray[i][x];
                    ctr++;
                }
                res[ctr] = 0xFF;
                ctr++;
            }
            return res;
        }

    function ba2cbor(bytes[] arr) internal returns (bytes) {
            uint arrlen = arr.length;

            // get correct cbor output length
            uint outputlen = 0;
            bytes[] memory elemArray = new bytes[](arrlen);
            for (uint i = 0; i < arrlen; i++) {
                elemArray[i] = (bytes(arr[i]));
                outputlen += elemArray[i].length + (elemArray[i].length - 1)/23 + 3; //+3 accounts for paired identifier types
            }
            uint ctr = 0;
            uint cborlen = arrlen + 0x80;
            outputlen += byte(cborlen).length;
            bytes memory res = new bytes(outputlen);

            while (byte(cborlen).length > ctr) {
                res[ctr] = byte(cborlen)[ctr];
                ctr++;
            }
            for (i = 0; i < arrlen; i++) {
                res[ctr] = 0x5F;
                ctr++;
                for (uint x = 0; x < elemArray[i].length; x++) {
                    // if there's a bug with larger strings, this may be the culprit
                    if (x % 23 == 0) {
                        uint elemcborlen = elemArray[i].length - x >= 24 ? 23 : elemArray[i].length - x;
                        elemcborlen += 0x40;
                        uint lctr = ctr;
                        while (byte(elemcborlen).length > ctr - lctr) {
                            res[ctr] = byte(elemcborlen)[ctr - lctr];
                            ctr++;
                        }
                    }
                    res[ctr] = elemArray[i][x];
                    ctr++;
                }
                res[ctr] = 0xFF;
                ctr++;
            }
            return res;
        }


    string oraclize_network_name;
    function oraclize_setNetworkName(string _network_name) internal {
        oraclize_network_name = _network_name;
    }

    function oraclize_getNetworkName() internal returns (string) {
        return oraclize_network_name;
    }

    function oraclize_newRandomDSQuery(uint _delay, uint _nbytes, uint _customGasLimit) internal returns (bytes32){
        if ((_nbytes == 0)||(_nbytes > 32)) throw;
        bytes memory nbytes = new bytes(1);
        nbytes[0] = byte(_nbytes);
        bytes memory unonce = new bytes(32);
        bytes memory sessionKeyHash = new bytes(32);
        bytes32 sessionKeyHash_bytes32 = oraclize_randomDS_getSessionPubKeyHash();
        assembly {
            mstore(unonce, 0x20)
            mstore(add(unonce, 0x20), xor(blockhash(sub(number, 1)), xor(coinbase, timestamp)))
            mstore(sessionKeyHash, 0x20)
            mstore(add(sessionKeyHash, 0x20), sessionKeyHash_bytes32)
        }
        bytes[3] memory args = [unonce, nbytes, sessionKeyHash];
        bytes32 queryId = oraclize_query(_delay, "random", args, _customGasLimit);
        oraclize_randomDS_setCommitment(queryId, sha3(bytes8(_delay), args[1], sha256(args[0]), args[2]));
        return queryId;
    }

    function oraclize_randomDS_setCommitment(bytes32 queryId, bytes32 commitment) internal {
        oraclize_randomDS_args[queryId] = commitment;
    }

    mapping(bytes32=>bytes32) oraclize_randomDS_args;
    mapping(bytes32=>bool) oraclize_randomDS_sessionKeysHashVerified;

    function verifySig(bytes32 tosignh, bytes dersig, bytes pubkey) internal returns (bool){
        bool sigok;
        address signer;

        bytes32 sigr;
        bytes32 sigs;

        bytes memory sigr_ = new bytes(32);
        uint offset = 4+(uint(dersig[3]) - 0x20);
        sigr_ = copyBytes(dersig, offset, 32, sigr_, 0);
        bytes memory sigs_ = new bytes(32);
        offset += 32 + 2;
        sigs_ = copyBytes(dersig, offset+(uint(dersig[offset-1]) - 0x20), 32, sigs_, 0);

        assembly {
            sigr := mload(add(sigr_, 32))
            sigs := mload(add(sigs_, 32))
        }


        (sigok, signer) = safer_ecrecover(tosignh, 27, sigr, sigs);
        if (address(sha3(pubkey)) == signer) return true;
        else {
            (sigok, signer) = safer_ecrecover(tosignh, 28, sigr, sigs);
            return (address(sha3(pubkey)) == signer);
        }
    }

    function oraclize_randomDS_proofVerify__sessionKeyValidity(bytes proof, uint sig2offset) internal returns (bool) {
        bool sigok;

        // Step 6: verify the attestation signature, APPKEY1 must sign the sessionKey from the correct ledger app (CODEHASH)
        bytes memory sig2 = new bytes(uint(proof[sig2offset+1])+2);
        copyBytes(proof, sig2offset, sig2.length, sig2, 0);

        bytes memory appkey1_pubkey = new bytes(64);
        copyBytes(proof, 3+1, 64, appkey1_pubkey, 0);

        bytes memory tosign2 = new bytes(1+65+32);
        tosign2[0] = 1; //role
        copyBytes(proof, sig2offset-65, 65, tosign2, 1);
        bytes memory CODEHASH = hex"fd94fa71bc0ba10d39d464d0d8f465efeef0a2764e3887fcc9df41ded20f505c";
        copyBytes(CODEHASH, 0, 32, tosign2, 1+65);
        sigok = verifySig(sha256(tosign2), sig2, appkey1_pubkey);

        if (sigok == false) return false;


        // Step 7: verify the APPKEY1 provenance (must be signed by Ledger)
        bytes memory LEDGERKEY = hex"7fb956469c5c9b89840d55b43537e66a98dd4811ea0a27224272c2e5622911e8537a2f8e86a46baec82864e98dd01e9ccc2f8bc5dfc9cbe5a91a290498dd96e4";

        bytes memory tosign3 = new bytes(1+65);
        tosign3[0] = 0xFE;
        copyBytes(proof, 3, 65, tosign3, 1);

        bytes memory sig3 = new bytes(uint(proof[3+65+1])+2);
        copyBytes(proof, 3+65, sig3.length, sig3, 0);

        sigok = verifySig(sha256(tosign3), sig3, LEDGERKEY);

        return sigok;
    }

    modifier oraclize_randomDS_proofVerify(bytes32 _queryId, string _result, bytes _proof) {
        // Step 1: the prefix has to match 'LP\x01' (Ledger Proof version 1)
        if ((_proof[0] != "L")||(_proof[1] != "P")||(_proof[2] != 1)) throw;

        bool proofVerified = oraclize_randomDS_proofVerify__main(_proof, _queryId, bytes(_result), oraclize_getNetworkName());
        if (proofVerified == false) throw;

        _;
    }

    function matchBytes32Prefix(bytes32 content, bytes prefix) internal returns (bool){
        bool match_ = true;

        for (var i=0; i<prefix.length; i++){
            if (content[i] != prefix[i]) match_ = false;
        }

        return match_;
    }

    function oraclize_randomDS_proofVerify__main(bytes proof, bytes32 queryId, bytes result, string context_name) internal returns (bool){
        bool checkok;


        // Step 2: the unique keyhash has to match with the sha256 of (context name + queryId)
        uint ledgerProofLength = 3+65+(uint(proof[3+65+1])+2)+32;
        bytes memory keyhash = new bytes(32);
        copyBytes(proof, ledgerProofLength, 32, keyhash, 0);
        checkok = (sha3(keyhash) == sha3(sha256(context_name, queryId)));
        if (checkok == false) return false;

        bytes memory sig1 = new bytes(uint(proof[ledgerProofLength+(32+8+1+32)+1])+2);
        copyBytes(proof, ledgerProofLength+(32+8+1+32), sig1.length, sig1, 0);


        // Step 3: we assume sig1 is valid (it will be verified during step 5) and we verify if 'result' is the prefix of sha256(sig1)
        checkok = matchBytes32Prefix(sha256(sig1), result);
        if (checkok == false) return false;


        // Step 4: commitment match verification, sha3(delay, nbytes, unonce, sessionKeyHash) == commitment in storage.
        // This is to verify that the computed args match with the ones specified in the query.
        bytes memory commitmentSlice1 = new bytes(8+1+32);
        copyBytes(proof, ledgerProofLength+32, 8+1+32, commitmentSlice1, 0);

        bytes memory sessionPubkey = new bytes(64);
        uint sig2offset = ledgerProofLength+32+(8+1+32)+sig1.length+65;
        copyBytes(proof, sig2offset-64, 64, sessionPubkey, 0);

        bytes32 sessionPubkeyHash = sha256(sessionPubkey);
        if (oraclize_randomDS_args[queryId] == sha3(commitmentSlice1, sessionPubkeyHash)){ //unonce, nbytes and sessionKeyHash match
            delete oraclize_randomDS_args[queryId];
        } else return false;


        // Step 5: validity verification for sig1 (keyhash and args signed with the sessionKey)
        bytes memory tosign1 = new bytes(32+8+1+32);
        copyBytes(proof, ledgerProofLength, 32+8+1+32, tosign1, 0);
        checkok = verifySig(sha256(tosign1), sig1, sessionPubkey);
        if (checkok == false) return false;

        // verify if sessionPubkeyHash was verified already, if not.. let's do it!
        if (oraclize_randomDS_sessionKeysHashVerified[sessionPubkeyHash] == false){
            oraclize_randomDS_sessionKeysHashVerified[sessionPubkeyHash] = oraclize_randomDS_proofVerify__sessionKeyValidity(proof, sig2offset);
        }

        return oraclize_randomDS_sessionKeysHashVerified[sessionPubkeyHash];
    }


    // the following function has been written by Alex Beregszaszi (@axic), use it under the terms of the MIT license
    function copyBytes(bytes from, uint fromOffset, uint length, bytes to, uint toOffset) internal returns (bytes) {
        uint minLength = length + toOffset;

        if (to.length < minLength) {
            // Buffer too small
            throw; // Should be a better way?
        }

        // NOTE: the offset 32 is added to skip the `size` field of both bytes variables
        uint i = 32 + fromOffset;
        uint j = 32 + toOffset;

        while (i < (32 + fromOffset + length)) {
            assembly {
                let tmp := mload(add(from, i))
                mstore(add(to, j), tmp)
            }
            i += 32;
            j += 32;
        }

        return to;
    }

    // the following function has been written by Alex Beregszaszi (@axic), use it under the terms of the MIT license
    // Duplicate Solidity's ecrecover, but catching the CALL return value
    function safer_ecrecover(bytes32 hash, uint8 v, bytes32 r, bytes32 s) internal returns (bool, address) {
        // We do our own memory management here. Solidity uses memory offset
        // 0x40 to store the current end of memory. We write past it (as
        // writes are memory extensions), but don't update the offset so
        // Solidity will reuse it. The memory used here is only needed for
        // this context.

        // FIXME: inline assembly can't access return values
        bool ret;
        address addr;

        assembly {
            let size := mload(0x40)
            mstore(size, hash)
            mstore(add(size, 32), v)
            mstore(add(size, 64), r)
            mstore(add(size, 96), s)

            // NOTE: we can reuse the request memory because we deal with
            //       the return code
            ret := call(3000, 1, 0, size, 128, size, 32)
            addr := mload(size)
        }

        return (ret, addr);
    }

    // the following function has been written by Alex Beregszaszi (@axic), use it under the terms of the MIT license
    function ecrecovery(bytes32 hash, bytes sig) internal returns (bool, address) {
        bytes32 r;
        bytes32 s;
        uint8 v;

        if (sig.length != 65)
          return (false, 0);

        // The signature format is a compact form of:
        //   {bytes32 r}{bytes32 s}{uint8 v}
        // Compact means, uint8 is not padded to 32 bytes.
        assembly {
            r := mload(add(sig, 32))
            s := mload(add(sig, 64))

            // Here we are loading the last 32 bytes. We exploit the fact that
            // 'mload' will pad with zeroes if we overread.
            // There is no 'mload8' to do this, but that would be nicer.
            v := byte(0, mload(add(sig, 96)))

            // Alternative solution:
            // 'byte' is not working due to the Solidity parser, so lets
            // use the second best option, 'and'
            // v := and(mload(add(sig, 65)), 255)
        }

        // albeit non-transactional signatures are not specified by the YP, one would expect it
        // to match the YP range of [27, 28]
        //
        // geth uses [0, 1] and some clients have followed. This might change, see:
        //  https://github.com/ethereum/go-ethereum/issues/2053
        if (v < 27)
          v += 27;

        if (v != 27 && v != 28)
            return (false, 0);

        return safer_ecrecover(hash, v, r, s);
    }

}
// </ORACLIZE_API>




/*************************************** Trump Token Contracts *********************************/

/* This is the full TrumpImpeachmentToken contract.

Copyright 2017 by the holders of this Ethereum address:
0x454dDD95B0Bc3D30224bdA5639F29d7fa16CFa0b


This contract is based on substantial amount of inheritance, each new subclass
adding one particular peace of functionality.

The inheritance graph looks like this:


            ERC20Interface
                 |
                 |
            ERC20Implementation
                 |
                 |
            LittleSisterToken
               /          \
              /            \
             /          BigSisterToken
            /                 |
           |                  |
   PriceIncreasingLST     PriceIncreasingToken                 usingOraclize
            |                                \                      |
            |                                \                      |
            |                           TimedEvidenceToken  TrumpOralce   StingOps
            |                                         \          /       /
            |                                         \         /       /
    TumpFullTermToken                             TrumpImpeachmentToken



 ERC20Token: Sending and receiving Tokens between holders

 Little/BigSisterToken: Buying and Selling tokens from and to contract

 PriceIncreasingToken/PriceIncreasingLST: Increase buyPrice over time
                                          to make the token more expensive

TimedEvidenceToken: Checks for winning conditions and suspends and
                    activates buying and selling

TrumpOracle: Use oraclize to answer who is President of the United States?

StringOps: String comparison, check if one string ends with another

TrumpImpeachmentToken: Combines TimedEvidenceToken and Oraclize

TumpFullTermToken: LittleSister of TrumpImpeachmentToken

*/


// Contract with just one string function endswith
contract StringOps {

    // returns true if string _a ends with string _b
    function stringEndsWith(string _a, string _b) internal returns (bool) {
        // convert to bytes because strings are rather powerless in Solidity as of 0.4.x
        bytes memory a = bytes(_a);
        bytes memory b = bytes(_b);

        // in case a is shorter than b, a most certainly cannot end with b
        if (a.length < b.length){
            return false;
        }

        // Difference in length between a and b, also offset for a
        uint length_diff = a.length - b.length;

        for (uint i = 0; i < b.length; i ++)
            // Check letter by letter, with an offset by length_diff
            if (a[i + length_diff] != b[i]){
                return false;
            }
        return true;
    }

}


// ERC Token standard #20 Interface
// https://github.com/ethereum/EIPs/issues/20
contract ERC20Interface {

    // Get the total token supply
    function totalSupply() public constant returns (uint256 supply);

    // Get the account balance of another account with address _owner
    function balanceOf(address _owner) public constant returns (uint256 balance);

    // Send _value amount of tokens to address _to
    function transfer(address _to, uint256 _value) public returns (bool success);

    // Send _value amount of tokens from address _from to address _to
    function transferFrom(address _from, address _to, uint256 _value) public returns (bool success);

    // Allow _spender to withdraw from your account, multiple times, up to the _value amount.
    // If this function is called again it overwrites the current allowance with _value.
    // this function is required for some DEX functionality
    function approve(address _spender, uint256 _value) public returns (bool success);

    // Returns the amount which _spender is still allowed to withdraw from _owner
    function allowance(address _owner, address _spender) public constant returns (uint256 remaining);

    // Triggered when tokens are transferred.
    event Transfer(address indexed _from, address indexed _to, uint256 _value);

    // Triggered whenever approve(address _spender, uint256 _value) is called.
    event Approval(address indexed _owner, address indexed _spender, uint256 _value);
}


// Implementation of the most intricate parts of the ERC20Interface that
// allows to send tokens around
contract ERC20Token is ERC20Interface{

    // The three letter symbol to define the token, should be overwritten in subclass
    string public constant symbol = "TBA";

    // Name of token should be overwritten in child
    string public constant name = "TBA";

    // The number of decimals, 6 should be enough (srsly who needs 18??)
    uint8 public constant decimals = 6;

    // With 6 decimals, a single unit is 10**6
    uint256 public constant unit = 1000000;

    // Balances for each account
    mapping(address => uint256) balances;

    // Owner of account approves the transfer of amount to another account
    mapping(address => mapping (address => uint256)) allowed;

    // What is the balance of a particular account?
    function balanceOf(address _owner) public constant returns (uint256) {
        return balances[_owner];
    }

    // Transfer the balance from owner's account to another account
    function transfer(address _to, uint256 _amount) public returns (bool) {
        if (balances[msg.sender] >= _amount && _amount > 0
                && balances[_to] + _amount > balances[_to]) {
            balances[msg.sender] -= _amount;
            balances[_to] += _amount;
            Transfer(msg.sender, _to, _amount);
            return true;
        } else {
            return false;
        }
    }

    // Send _value amount of tokens from address _from to address _to
    // The transferFrom method is used for a withdraw workflow, allowing contracts to send
    // tokens on your behalf, for example to "deposit" to a contract address and/or to charge
    // fees in sub-currencies; the command should fail unless the _from account has
    // deliberately authorized the sender of the message via some mechanism; we propose
    // these standardized APIs for approval:
    function transferFrom(
        address _from,
        address _to,
        uint256 _amount
    ) public returns (bool) {
        if (balances[_from] >= _amount
            && allowed[_from][msg.sender] >= _amount && _amount > 0
                && balances[_to] + _amount > balances[_to]) {
            balances[_from] -= _amount;
            allowed[_from][msg.sender] -= _amount;
            balances[_to] += _amount;
            Transfer(_from, _to, _amount);
            return true;
        } else {
            return false;
        }
    }

    // Allow _spender to withdraw from your account, multiple times, up to the _value amount.
    // If this function is called again it overwrites the current allowance with _value.
    function approve(address _spender, uint256 _amount) public returns (bool) {
        allowed[msg.sender][_spender] = _amount;
        Approval(msg.sender, _spender, _amount);
        return true;
    }

    // Function to specify how much _spender is allowed to transfer on _owner's behalf
    function allowance(address _owner, address _spender) public constant returns (uint256) {
        return allowed[_owner][_spender];
    }

}


/*

The LittleSister allows selling and buying of tokens while paying a vig to the owner.

However, the LittleSisterToken is feeling insecure alone and needs to be guided by her
BigSistersToken that tells her how much stuff is worth.

Both contracts are useless unless they know of each other via registerSister of the
BigSisterToken.

*/
contract LittleSisterToken is ERC20Token{

    // The vig sent to the owner for each buying transaction
    // e.g. vig = 40 -> amount / 40 -> 2.5% are deducted
    uint8 public vig;

    // Maximum Supply of tokens, same as for the BigSisterToken
    // Note that totalSupply_BigSister + totalSupply_LittleSister <= maxSupply
    uint256 public maxSupply;

    // currently circulating Supply bounded by maxSupply (see above)
    uint256 totalSupply_;

    // Maximum date contract lives on its own, afterwards contract can be destroyed
    // Should be a picked far, far in the future to allow Token holders enough time to redeem their
    // winnings
    uint public validityDate;

    // Maximum date unitl buying is still possible
    uint public maxBuyingDate;

    // Address of the (big) sister of the contract
    address public sister;

    // Owners of this contract, awesome people!
    address public owner;

    // Address to send vig to
    address public vigAddress;

    // If contract has lost, i.e. if so, the other contract has won (obviously)
    bool public lost;

    // Constructor only sets owner and totalSupply, other properties are either
    // taken from the big sister or set during registering
    function LittleSisterToken() {
        owner = msg.sender;
        vigAddress = msg.sender;
        totalSupply_ = 0;
        lost = false;
    }

    // fallback function, note only the big sister can send funds to this contract
    // without buying, other people just buy the tokens
    function () payable {
        if (msg.sender != sister){
            buy();
        }
    }

    // Changes the vig address
    function setVigAddress(address _newVigAddress) public {
        require(msg.sender == owner);
        vigAddress = _newVigAddress;
    }

    // What is total circulating supply?
    function totalSupply() public constant returns (uint256) {
        return totalSupply_;
    }

    // function to link big and little sister, this one can only be
    // called by the big sister and only once
    // Once linked contracts CANNOT be unlinked or replaced!
    function registerSister(address _sister) public returns (bool) {
        BigSisterToken sisterContract;

        // take care that registering can only be called once and the sister
        // cannot be changed afterwards
        require(sister == address(0));
        // can only be called from the big sister, linking can be invoked there by the owner
        require(_sister != address(this));

        // Check that the big sister is who she claims to be
        sisterContract = BigSisterToken(_sister);
        require(sisterContract.owner() == owner);
        require(sisterContract.sister() == address(this));

        // she can be trusted and is accepted as the new big sister
        sister = _sister;

        // maxSupply, vig and validity and max buying date should be equal to the big sister's
        maxSupply = sisterContract.maxSupply();
        vig = sisterContract.vig();
        validityDate = sisterContract.validityDate();
        maxBuyingDate = sisterContract.maxBuyingDate();

        return true;
    }

    // The selling price is also determined by the big sister
    // returns an uint256 sell price
    function sellPrice() public constant returns (uint256){
        BigSisterToken sisterContract = BigSisterToken(sister);
        return sisterContract.sellPrice();
    }

    // Likewise the big sister dictates the buying price
    function buyPrice() public constant returns (uint256){
        BigSisterToken sisterContract = BigSisterToken(sister);
        return sisterContract.buyPrice();
    }

    // Sends all the funds to the big sister, can only be invoked by her
    function sendFunds() public returns(bool){
        require(msg.sender == sister);
        require(msg.sender.send(this.balance));
        // in this case the sister won and this contract lost
        lost = true;
        return true;
    }

    // Kills the contract and sends funds to owner
    // However, only if validityDate has been surpassed (should be chosen far, far in the future)
    function releaseContract() public {
        require(now > validityDate);
        require(msg.sender == sister);
        selfdestruct(owner);
    }

    // What is the value of this token in case this token wins?
    // Value is balance of this plus balance of sister account per token
    function winningValue() public constant returns (uint256 value){
        LittleSisterToken sisterContract = LittleSisterToken(sister);

        if (totalSupply_ > 0){
            value = (this.balance + sisterContract.balance) / totalSupply_;
        }

        return value;
    }

    // Supply that still can be bought
    function availableSupply() public constant returns (uint256 supply){
        LittleSisterToken sisterContract;
        uint256 sisterSupply;

        if ((buyPrice()) == 0 || (now > maxBuyingDate)){
            // there can only be supply in case buying is allowed
            return 0;
        } else {

            sisterContract = LittleSisterToken(sister);
            sisterSupply = sisterContract.totalSupply();
            // This is the most important assertion ever ;-)
            assert(maxSupply >= totalSupply_ + sisterSupply);
            supply = maxSupply - totalSupply_ - sisterSupply;
            return supply;
        }
    }

    // Buy tokens from the contract depending on how much money has been sent
    function buy() public payable returns (uint amount){
        uint256 house;
        uint256 supply;
        uint256 price;

        // check that buying is still allowed
        require(now <= maxBuyingDate);
        // get the price
        price = buyPrice();
        // Stuff can only be sold if the price is nonzero
        // i.e. setting the price to 0 can suspend the Token sale
        require(price > 0);
        house = msg.value / vig;
        // C'mon you gotta pay the house!
        require(house > 0);
        amount = msg.value / price;
        supply = availableSupply();
        require(amount <= supply);
        // Send the commission to the vigAddress of this contract
        vigAddress.transfer(house);
        totalSupply_ += amount;
        // Increase the sender's balance by the appropriate amount
        balances[msg.sender] += amount;
        Transfer(this, msg.sender, amount);
        return amount;
    }

    // Users can redeem their cash and sell their tokens back to the contract
    // Note that this usually requires that either the little or the big sister
    // has transferred all their funds to the other contract and set sell price nonzero
    function sell(uint _amount) public returns (uint256 revenue){
        // Tokens can only be sold back to the contract if the price is nonzero
        // i.e. the sellPrice can also act as a switch, usually it is off by default
        // and only turned on by some event (i.e. one of the two tokens winning)
        require(sellPrice() > 0);
        require(!lost);
        // check if the user has sufficient amount of tokens
        require(balances[msg.sender] >= _amount);
        revenue = _amount * sellPrice();
        require(this.balance >= revenue);
        // first deduct tokens from user
        balances[msg.sender] -= _amount;
        totalSupply_ -= _amount;
        // finally transfer the ether, this withdrawal pattern prevents re-entry attacks
        msg.sender.transfer(revenue);
        Transfer(msg.sender, this, _amount);
        return revenue;
    }

    // Sells all tokens of a user
    function sellAll() public returns (uint256){
        uint256 amount = balances[msg.sender];
        return sell(amount);
    }

}


/*

This is the big sister of the LittleSisterToken, it governs the little sister.
It knows about the sell and buying prices.

*/
contract BigSisterToken is LittleSisterToken {

    // The buying price, 0 means buying is suspended
    uint256 buyPrice_;

    // The sell price, 0 suspends selling
    uint256 sellPrice_;

    // Constructor
    function BigSisterToken(uint256 _maxSupply, uint256 _buyPrice, uint256 _sellPrice, uint8 _vig,
            uint _buyingHorizon, uint _maximumLifetime) {
        maxSupply = _maxSupply;
        vig = _vig;
        buyPrice_ = _buyPrice;
        sellPrice_ = _sellPrice;
        // Maximum lifetime is not a date, but a time interval
        validityDate = now + _maximumLifetime;
        // Horizon after which buying is definitely suspended
        maxBuyingDate = now + _buyingHorizon;
    }

    function buyPrice() public constant returns (uint256){
        return buyPrice_;
    }

    function sellPrice() public constant returns (uint256){
        return sellPrice_;
    }

    // Destroys the contract as well as the little sister.
    // Important _maximumLifetime should be set generously to give token holders
    // plenty fo time to withdraw their money in case they win
    // only the owner can invoke this command
    function releaseContract() public {
        LittleSisterToken sisterContract = LittleSisterToken(sister);

        require(now > validityDate);
        require(msg.sender == owner);
        // also destroy the little sister
        sisterContract.releaseContract();
        selfdestruct(owner);
    }

    // Links a little and a big sister together, can only be invoked by owner.
    // IMPORTANT contracts can only be registered once, linking CANNOT be changed or modified.
    function registerSister(address _sister) public returns (bool) {
        LittleSisterToken sisterContract;

        // make sure that contract registration works only once
        require(sister == address(0));
        // can onyl be called by owner
        require(msg.sender == owner);
        require(_sister != address(this));
        sisterContract = LittleSisterToken(_sister);
        require(sisterContract.sister() == address(0));
        sister = _sister;
        // finally also register this big sister with its new little sister
        require(sisterContract.registerSister(address(this)));
        return true;
    }

}


/*

This token allows for price increases after a certain amount of purchases.
Note that this functionality is actually provided by its big sister called
PriceIncreasingToken.

*/
contract PriceIncreasingLittleSisterToken is LittleSisterToken{

    // fallback function, note only the big sister can send funds to this contract
    // without buying, other people just buy the tokens
    function () payable {
        if (msg.sender != sister){
            buy();
        }
    }

    // Buy tokens from the contract, as in parent class, but
    // ask the sister if price needs to be increased for future purchases
    function buy() public payable returns (uint amount){
        PriceIncreasingToken sisterContract = PriceIncreasingToken(sister);

        // buy tokens
        amount = super.buy();
        // notify the sister about the sold amount and maybe update prices
        sisterContract.sisterCheckPrice(amount);
        return amount;
    }

}


/*

This is the big sister for price increases.
Needs to be linked to a PriceIncreasingLittleSisterToken.
Note that price increase always happens after someone bought tokens, never before.

*/
contract PriceIncreasingToken is BigSisterToken{

    // Helper variable to store number of tokens purchases since last price increase
    uint256 public currentBatch;

    // Number of tokens that need to be bought to trigger a price increase
    uint256 public thresholdAmount;

    // Increase of price if thresholdAmount is met
    uint256 public priceIncrease;

    function PriceIncreasingToken(uint256 _maxSupply, uint256 _buyPrice, uint256 sellPrice_, uint8 _vig,
        uint256 _thresholdAmount, uint256 _priceIncrease, uint _buyingHorizon, uint _maximumLifetime)
            BigSisterToken( _maxSupply, _buyPrice, sellPrice_, _vig, _buyingHorizon, _maximumLifetime){
        currentBatch = 0;
        thresholdAmount = _thresholdAmount;
        priceIncrease = _priceIncrease;
    }

    // fallback function, note only the big sister can send funds to this contract
    // without buying, other people just buy the tokens
    function () payable {
        if (msg.sender != sister){
            buy();
        }
    }

    // Buy tokens from the contract and checks price and potentially increases
    // the price afterwards
    function buy() public payable returns (uint amount){
        amount = super.buy();
        // check for price increase
        _checkPrice(amount);
        return amount;
    }

    // Function that can only be called by the sister to check and increase prices
    // This is needed because the _checkPrice function is internal and cannot
    // be called by the sister directly
    function sisterCheckPrice(uint256 amount) public{
        require(msg.sender == sister);
        _checkPrice(amount);
    }

    // internal function to check and maybe increase the buying price
    // every time the thresholdAmount is met, the price is increased
    function _checkPrice(uint256 amount) internal{
        currentBatch += amount;
        if (currentBatch >= thresholdAmount){
            buyPrice_ += priceIncrease;
            // it is important to subtract the thresholdAmount
            // and not set currentBatch = 0 instead. This ensures that
            // a huge order crossing multiple thresholds will trigger price
            // increases in quick successions
            currentBatch -= thresholdAmount;
        }
    }

}


/*

This contract adds the ability to check for evidence to enable the contract's winning.
There are two conditions, either a certain fixed date has passed (dateSisterWins) and,
in this case, the litte sister wins, or some evidence is gathered proving the opposite.

How this evidence is gathered needs to be implemented in the subclass. Note that
this evidence needs to be acquired a couple of times to be certain. Moreover, in between
different findings of repeating evidence there needs to pass a certain evidenceInterval.

*/

contract TimedEvidenceToken is PriceIncreasingToken{

    // fixed date that specifies when the little sister will win (unix timestamp)
    uint public dateSisterWins;

    // amount of evidence found
    uint8 public foundEvidence;

    // amount of evidence required for victory,
    // the big sister wins if foundEvidence >= requiredEvidence
    uint8 public requiredEvidence;

    // Time interval between two consecutive evidence checks
    uint public evidenceInterval;

    // Last time evidence was checked
    uint public lastEvidenceCheck;

    // Helper variable to store the last price.
    // If the contract needs to suspend buying immediately in case
    // of evidence, this can restore buying if evidence cannot be found subsequently
    uint256 lastBuyPrice;

    function TimedEvidenceToken(uint256 _maxSupply, uint256 _buyPrice, uint8 _vig,
        uint256 _thresholdAmount, uint256 _priceIncrease, uint _evidenceInterval,
        uint8 _requiredEvidence, uint _dateSisterWins,
        uint _buyingHorizon, uint _maximumLifetime)
        PriceIncreasingToken(_maxSupply, _buyPrice, 0,_vig,
        _thresholdAmount, _priceIncrease, _buyingHorizon, _maximumLifetime){
    evidenceInterval = _evidenceInterval;
    lastEvidenceCheck = 0;
    foundEvidence = 0;
    lastBuyPrice = 0;
    dateSisterWins = _dateSisterWins;
    requiredEvidence = _requiredEvidence;
    }

    // Checks and initiates payout, i.e. users can sell their tokens back to the contract
    function checkForPayout() public returns(bool){
        LittleSisterToken sisterContract = LittleSisterToken(sister);

        // if we already sell, payout is definitely true
        if (sellPrice_ > 0) return true;

        // Check if the fixed date has passed:
        if (now > dateSisterWins){
            // Sister wins and this token becomes worthless
            require(sisterContract.send(this.balance));
            // this contract lost
            lost = true;
            // buying is no more possible
            buyPrice_ = 0;
            // now set the sell price
            sellPrice_ = sisterContract.balance / sisterContract.totalSupply();
            return true;
        }

        // Check if enough evidence was found
        if (foundEvidence >= requiredEvidence){
            // Trump is impeached this token is gaining in value!
            // Sister lost and needs to send her funds
            require(sisterContract.sendFunds());
            // buying is no more possible
            buyPrice_ = 0;
            // now set the sell price
            sellPrice_ = this.balance / totalSupply();
            return true;

        }

        // If both cases are not true, then unfortunately, we cannot pay you (yet)
        return false;
    }

    // internal function needs to be called by the evidence gathering implementation
    // in subclass
    function _accumulateEvidence(bool evidence) internal{
        // make sure that enough time has passed since the last evidence check
        require(now > lastEvidenceCheck + evidenceInterval);
        lastEvidenceCheck = now;

        if (evidence){
            if (buyPrice_ > 0){
                // suspend buying as soon as there is evidence
                lastBuyPrice = buyPrice_;
                buyPrice_ = 0;
            }
            // increase evidence
            foundEvidence += 1;
        } else {
            // resume buying in case foundEvidence is reduced.
            // Note that resuming buying needs to find one more evidence against than pro.
            // This should stop an attacker to reduce evidence to 0 and subsequently buy
            // tokens quickly
            if ((lastBuyPrice > 0) && (foundEvidence == 0)){
                buyPrice_ = lastBuyPrice;
                lastBuyPrice = 0;
            }

            // if evidence is not found consecutively it is decreased
            // buying can only be enabled again if found evidence is 0 for a consecutive interval
            if (foundEvidence > 0) foundEvidence -= 1;
        }
    }

}


// Dummy LittleSisterContract that just specifies the names
// Has a TrumpImpeachmentToken big sister
contract TrumpFullTermToken is PriceIncreasingLittleSisterToken{

    string public constant symbol = "TFT";

    string public constant name = "Trump Full Term Token";

}


//Implements the Oraclize Usage to query for the President of the United States
contract TrumpOracle is usingOraclize{

    // Keeps query ids to make sure the callback is valid
    mapping(bytes32=>bool) validIds;

    // logs an oraclize query
    event newOraclizeQuery(string description);

    // logs the query result
    event newOraclizeResult(bytes32 id, string result);

    // question asked to oracle
    string public constant question = "President of the United States";

    // price of (last) Oraclize query
    uint public oraclizePrice;

    // callback function used by oraclize to provide the result
    function __callback(bytes32 _queryId, string result) public {
        // make sure the request comes from oraclize
        require(msg.sender == oraclize_cbAddress());
        // and matches a previous query
        require(validIds[_queryId]);
        delete validIds[_queryId];
        newOraclizeResult(_queryId, result);
    }

    // Can be called by users and token holders to check if Trump is still president
    function requestEvidence() public payable {
        if (getOraclizePrice() > msg.value) {
            // note that oraclize deducts the query costs from the contract, so the
            // users should compensate for that!
            newOraclizeQuery("Oraclize query was NOT sent, please add some ETH to cover for the query fee");
            // send money back to user and revert transaction
            revert();
        } else {
            newOraclizeQuery("Oraclize query was sent, standing by for the answer...");
            // THE MOST IMPORTANT QUERY!
            bytes32 queryId = oraclize_query("WolframAlpha", question);
            // Keep track of id for callback check
            validIds[queryId] = true;
        }
    }

    // returns the price of oracle call in wei
    function getOraclizePrice() public returns (uint) {
        oraclizePrice = oraclize_getPrice("WolframAlpha");
        return oraclizePrice;
    }

}


// The big sister contract of TrumpFullTermToken
contract TrumpImpeachmentToken is TrumpOracle, TimedEvidenceToken, StringOps{

    // 3 letter symbol
    string public constant symbol = "TIT";

    // Name of token
    string public constant name = "Trump Impeachment Token";

    // Last String returned by oracle query
    string public lastEvidence;

    // expected answer
    string public constant answer = "Trump";

    function TrumpImpeachmentToken()
             TimedEvidenceToken(2000000 * unit, // amount
                                69 finney / unit, // buy price 0.069 Ether
                                40, // vig => 2.5%
                                6600 * unit, //threshold amount
                                100 szabo / unit, // price increase
                                2 days, //evidence interval
                                3, // required evidence
                                1611014400, // date sister wins,
                                            //one day before Trump`s term ends
                                222 days, // buying horizon where buying tokens is allowed
                                7 years //maximum Lifetime from deployment on
                                        // after which contract can be destroyed
             ){
            lastEvidence = "N/A";
            // for very fast confirmations set the gas Price to 30GWei
            oraclize_setCustomGasPrice(30000000000);
        }

    // callback used by oraclize
    function __callback(bytes32 _queryId, string result) public{
        bool evidence;

        super.__callback(_queryId, result);
        require(bytes(result).length > 0);
        lastEvidence = result;
        evidence = !stringEndsWith(result, answer);
        // accumulates evidence over time, also checks for sufficient interval between
        // evidence gatherings
        _accumulateEvidence(evidence);
    }

    // can be called by user to request evidence from oraclize
    function requestEvidence() payable{
        // check if enough time has past before querying oraclize
        require(now > lastEvidenceCheck + evidenceInterval);
        super.requestEvidence();
    }

    // We as owners can (only) increase or keep the price in case there is a gas price surge
    // within the range of 30 to 300 GWei.
    // This cannot do any harm (well except making the confirmations a bit expensive)
    // but it can be helpful in case of surging prices
    function setOraclizeGasPrice(uint gasPrice){
        require(msg.sender == owner);
        // must be greater or equal to the original gasPrice 30 GWei
        require(gasPrice >= 30000000000);
        // and lets keep it within reasonable bounds maximum is 300 GWei, i.e. 10 fold
        // this should be a reasonable range, for reacting to price surges and
        // impeachment token holders not fearing to pay a too high price for confirmations
        require(gasPrice <= 300000000000);
        oraclize_setCustomGasPrice(gasPrice);
    }

}