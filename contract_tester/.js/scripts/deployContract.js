#! /usr/local/bin/bnode
"use strict";

var myDeployContract = function () {
    var _ref = _asyncToGenerator(regeneratorRuntime.mark(function _callee(contract, abi, bin) {
        var contract_bin, contract_abi, contractJSON, vParams, aParams, initialParams, Contract, _utils, _utils2;

        return regeneratorRuntime.wrap(function _callee$(_context) {
            while (1) {
                switch (_context.prev = _context.next) {
                    case 0:
                        contract_bin = void 0;
                        contract_abi = void 0;
                        _context.prev = 2;
                        _context.next = 5;
                        return fread(bin);

                    case 5:
                        contract_bin = _context.sent;
                        _context.next = 8;
                        return fread(abi);

                    case 8:
                        contract_abi = _context.sent;
                        _context.next = 14;
                        break;

                    case 11:
                        _context.prev = 11;
                        _context.t0 = _context["catch"](2);
                        return _context.abrupt("return");

                    case 14:
                        contractJSON = new Object();

                        contractJSON.abi = JSON.parse(contract_abi);
                        contractJSON.unlinked_binary = contract_bin;
                        contractJSON.contract_name = contract;

                        vParams = {
                            from: _Account.defaultAccount,
                            gas: _Account.defaultGas,
                            value: _Account.defaultValue
                        };
                        aParams = {
                            from: _Account.defaultAccount,
                            gas: _Account.defaultGas
                        };
                        initialParams = {};
                        Contract = void 0;
                        _context.prev = 22;
                        _utils = new _ContractUtils.ContractUtils();
                        _context.next = 26;
                        return _utils.deploy_json(contractJSON, initialParams, vParams);

                    case 26:
                        Contract = _context.sent;

                        if (Contract != undefined) console.log(Contract.address + Contract.abi);
                        _context.next = 43;
                        break;

                    case 30:
                        _context.prev = 30;
                        _context.t1 = _context["catch"](22);
                        _context.prev = 32;
                        _utils2 = new _ContractUtils.ContractUtils();
                        _context.next = 36;
                        return _utils2.deploy_json(contractJSON, initialParams, aParams);

                    case 36:
                        Contract = _context.sent;

                        if (Contract != undefined) console.log(Contract.address + Contract.abi);
                        _context.next = 43;
                        break;

                    case 40:
                        _context.prev = 40;
                        _context.t2 = _context["catch"](32);

                        console.log("error:%s", _context.t2);

                    case 43:
                        return _context.abrupt("return", Contract);

                    case 44:
                    case "end":
                        return _context.stop();
                }
            }
        }, _callee, this, [[2, 11], [22, 30], [32, 40]]);
    }));

    return function myDeployContract(_x, _x2, _x3) {
        return _ref.apply(this, arguments);
    };
}();

var myDeploy = function () {
    var _ref2 = _asyncToGenerator(regeneratorRuntime.mark(function _callee2() {
        var abidir, bindir, root, pa, _iteratorNormalCompletion, _didIteratorError, _iteratorError, _iterator, _step, ele, info, pattern, reg, contract, abi, bin;

        return regeneratorRuntime.wrap(function _callee2$(_context2) {
            while (1) {
                switch (_context2.prev = _context2.next) {
                    case 0:
                        abidir = program.abidir;
                        bindir = program.bindir;
                        root = path.join(program.contractdir);
                        pa = fs.readdirSync(root);
                        _iteratorNormalCompletion = true;
                        _didIteratorError = false;
                        _iteratorError = undefined;
                        _context2.prev = 7;
                        _iterator = pa[Symbol.iterator]();

                    case 9:
                        if (_iteratorNormalCompletion = (_step = _iterator.next()).done) {
                            _context2.next = 33;
                            break;
                        }

                        ele = _step.value;
                        info = fs.statSync(root + "/" + ele);

                        if (!info.isDirectory()) {
                            _context2.next = 16;
                            break;
                        }

                        console.log("ele: " + ele);
                        _context2.next = 30;
                        break;

                    case 16:
                        pattern = '\\w+(?=\\' + program.contractsuffix + ')';
                        _context2.prev = 17;
                        reg = new RegExp(pattern, "g");
                        contract = reg.exec(ele.toString())[0];
                        abi = abidir + "/" + contract + program.abisuffix;
                        bin = bindir + "/" + contract + program.binsuffix;

                        console.log("abi:%j", abi);
                        _context2.next = 25;
                        return myDeployContract(contract, abi, bin);

                    case 25:
                        _context2.next = 30;
                        break;

                    case 27:
                        _context2.prev = 27;
                        _context2.t0 = _context2["catch"](17);

                        console.log("%s", _context2.t0.toString());

                    case 30:
                        _iteratorNormalCompletion = true;
                        _context2.next = 9;
                        break;

                    case 33:
                        _context2.next = 39;
                        break;

                    case 35:
                        _context2.prev = 35;
                        _context2.t1 = _context2["catch"](7);
                        _didIteratorError = true;
                        _iteratorError = _context2.t1;

                    case 39:
                        _context2.prev = 39;
                        _context2.prev = 40;

                        if (!_iteratorNormalCompletion && _iterator.return) {
                            _iterator.return();
                        }

                    case 42:
                        _context2.prev = 42;

                        if (!_didIteratorError) {
                            _context2.next = 45;
                            break;
                        }

                        throw _iteratorError;

                    case 45:
                        return _context2.finish(42);

                    case 46:
                        return _context2.finish(39);

                    case 47:
                        //let contract = "Corporation";
                        // let contract = "ARK";
                        //let abi = abidir + "/" + contract + program.abisuffix;
                        // let bin = bindir + "/" + contract + program.binsuffix;
                        //await myDeployContract(contract, abi, bin);
                        process.exit(0);

                    case 48:
                    case "end":
                        return _context2.stop();
                }
            }
        }, _callee2, this, [[7, 35, 39, 47], [17, 27], [40,, 42, 46]]);
    }));

    return function myDeploy() {
        return _ref2.apply(this, arguments);
    };
}();

var deployContract = function () {
    var _ref3 = _asyncToGenerator(regeneratorRuntime.mark(function _callee3(contract) {
        var contract_bin, contract_abi, contractJSON, vParams, initialParams, Contract;
        return regeneratorRuntime.wrap(function _callee3$(_context3) {
            while (1) {
                switch (_context3.prev = _context3.next) {
                    case 0:
                        _context3.next = 2;
                        return fread(WorkPlace + "/caller.bin/" + contract + ".bin");

                    case 2:
                        contract_bin = _context3.sent;
                        _context3.next = 5;
                        return fread(WorkPlace + "/caller.abi/" + contract + "_abi.json");

                    case 5:
                        contract_abi = _context3.sent;
                        contractJSON = new Object();

                        contractJSON.abi = JSON.parse(contract_abi);
                        contractJSON.unlinked_binary = contract_bin;
                        contractJSON.contract_name = contract;

                        vParams = {
                            from: _Account.defaultAccount,
                            gas: _Account.defaultGas,
                            value: _Account.defaultValue
                        };
                        initialParams = {};
                        _context3.next = 14;
                        return utils.deploy_json(contractJSON, initialParams, vParams);

                    case 14:
                        Contract = _context3.sent;

                        if (Contract != undefined) {
                            console.log(Contract.address + Contract.abi);
                        }
                        return _context3.abrupt("return", Contract);

                    case 17:
                    case "end":
                        return _context3.stop();
                }
            }
        }, _callee3, this);
    }));

    return function deployContract(_x4) {
        return _ref3.apply(this, arguments);
    };
}();

var deploy = function () {
    var _ref4 = _asyncToGenerator(regeneratorRuntime.mark(function _callee4() {
        var contracts, _iteratorNormalCompletion2, _didIteratorError2, _iteratorError2, _iterator2, _step2, contract;

        return regeneratorRuntime.wrap(function _callee4$(_context4) {
            while (1) {
                switch (_context4.prev = _context4.next) {
                    case 0:
                        console.log("argv:" + argv);
                        console.log("argc:" + argc);
                        contracts = argv.slice(2);

                        console.log("contracts:" + contracts);

                        _iteratorNormalCompletion2 = true;
                        _didIteratorError2 = false;
                        _iteratorError2 = undefined;
                        _context4.prev = 7;
                        _iterator2 = contracts[Symbol.iterator]();

                    case 9:
                        if (_iteratorNormalCompletion2 = (_step2 = _iterator2.next()).done) {
                            _context4.next = 17;
                            break;
                        }

                        contract = _step2.value;

                        console.log(contract);
                        _context4.next = 14;
                        return deployContract(contract);

                    case 14:
                        _iteratorNormalCompletion2 = true;
                        _context4.next = 9;
                        break;

                    case 17:
                        _context4.next = 23;
                        break;

                    case 19:
                        _context4.prev = 19;
                        _context4.t0 = _context4["catch"](7);
                        _didIteratorError2 = true;
                        _iteratorError2 = _context4.t0;

                    case 23:
                        _context4.prev = 23;
                        _context4.prev = 24;

                        if (!_iteratorNormalCompletion2 && _iterator2.return) {
                            _iterator2.return();
                        }

                    case 26:
                        _context4.prev = 26;

                        if (!_didIteratorError2) {
                            _context4.next = 29;
                            break;
                        }

                        throw _iteratorError2;

                    case 29:
                        return _context4.finish(26);

                    case 30:
                        return _context4.finish(23);

                    case 31:
                    case "end":
                        return _context4.stop();
                }
            }
        }, _callee4, this, [[7, 19, 23, 31], [24,, 26, 30]]);
    }));

    return function deploy() {
        return _ref4.apply(this, arguments);
    };
}();
//deploy();


var readDirSync = function () {
    var _ref5 = _asyncToGenerator(regeneratorRuntime.mark(function _callee5(path) {
        var pa, _iteratorNormalCompletion3, _didIteratorError3, _iteratorError3, _iterator3, _step3, ele, info, reg, contract;

        return regeneratorRuntime.wrap(function _callee5$(_context5) {
            while (1) {
                switch (_context5.prev = _context5.next) {
                    case 0:
                        pa = fs.readdirSync(path);
                        // pa.forEach(function (ele, index) {

                        _iteratorNormalCompletion3 = true;
                        _didIteratorError3 = false;
                        _iteratorError3 = undefined;
                        _context5.prev = 4;
                        _iterator3 = pa[Symbol.iterator]();

                    case 6:
                        if (_iteratorNormalCompletion3 = (_step3 = _iterator3.next()).done) {
                            _context5.next = 21;
                            break;
                        }

                        ele = _step3.value;
                        info = fs.statSync(path + "/" + ele);

                        if (!info.isDirectory()) {
                            _context5.next = 13;
                            break;
                        }

                        console.log("dir: " + ele);
                        //readDirSync(path + "/" + ele);
                        _context5.next = 18;
                        break;

                    case 13:
                        //console.log(ele);
                        reg = /\w+(?=\.sol)/g;
                        contract = ele.match(reg)[0];

                        console.log(contract);
                        _context5.next = 18;
                        return deployContract(contract);

                    case 18:
                        _iteratorNormalCompletion3 = true;
                        _context5.next = 6;
                        break;

                    case 21:
                        _context5.next = 27;
                        break;

                    case 23:
                        _context5.prev = 23;
                        _context5.t0 = _context5["catch"](4);
                        _didIteratorError3 = true;
                        _iteratorError3 = _context5.t0;

                    case 27:
                        _context5.prev = 27;
                        _context5.prev = 28;

                        if (!_iteratorNormalCompletion3 && _iterator3.return) {
                            _iterator3.return();
                        }

                    case 30:
                        _context5.prev = 30;

                        if (!_didIteratorError3) {
                            _context5.next = 33;
                            break;
                        }

                        throw _iteratorError3;

                    case 33:
                        return _context5.finish(30);

                    case 34:
                        return _context5.finish(27);

                    case 35:
                    case "end":
                        return _context5.stop();
                }
            }
        }, _callee5, this, [[4, 23, 27, 35], [28,, 30, 34]]);
    }));

    return function readDirSync(_x5) {
        return _ref5.apply(this, arguments);
    };
}();

var deploy2 = function () {
    var _ref6 = _asyncToGenerator(regeneratorRuntime.mark(function _callee6() {
        var dirname, root;
        return regeneratorRuntime.wrap(function _callee6$(_context6) {
            while (1) {
                switch (_context6.prev = _context6.next) {
                    case 0:
                        dirname = WorkPlace + "/caller/waitting";
                        root = path.join(dirname);
                        _context6.next = 4;
                        return readDirSync(root);

                    case 4:
                        process.exit(0);

                    case 5:
                    case "end":
                        return _context6.stop();
                }
            }
        }, _callee6, this);
    }));

    return function deploy2() {
        return _ref6.apply(this, arguments);
    };
}();
//deploy2();


var _unixTimestamp = require("unix-timestamp");

var timestamp = _interopRequireWildcard(_unixTimestamp);

var _fs = require("fs");

var fs = _interopRequireWildcard(_fs);

var _path = require("path");

var path = _interopRequireWildcard(_path);

var _ContractUtils = require("../utils/ContractUtils.js");

var _Account = require("../utils/Account.js");

function _interopRequireWildcard(obj) { if (obj && obj.__esModule) { return obj; } else { var newObj = {}; if (obj != null) { for (var key in obj) { if (Object.prototype.hasOwnProperty.call(obj, key)) newObj[key] = obj[key]; } } newObj.default = obj; return newObj; } }

function _asyncToGenerator(fn) { return function () { var gen = fn.apply(this, arguments); return new _ContractUtils.Promise(function (resolve, reject) { function step(key, arg) { try { var info = gen[key](arg); var value = info.value; } catch (error) { reject(error); return; } if (info.done) { resolve(value); } else { return _ContractUtils.Promise.resolve(value).then(function (value) { step("next", value); }, function (err) { step("throw", err); }); } } return step("next"); }); }; }

_ContractUtils.dotenv.config();

var getBalance = _Account.web3.eth.getBalance;
var WorkPlace = process.env.WorkPlace;

var BIN_DIR = process.env.BIN;
var BIN_Suffix = process.env.BIN_suffix;
var ABI_DIR = process.env.ABI;
var ABI_Suffix = process.env.ABI_suffix;
var CONTRACT_DIR = process.env.Contract;
var CONTRACT_Suffix = process.env.Contract_suffix;
var program = require("commander");

program.version("0.0.1").option("-c, --contractdir <path>", "set CONTRACT_DIR(" + CONTRACT_DIR + ")", CONTRACT_DIR).option("-b, --bindir <path>", "set BIN_DIR(" + BIN_DIR + ")", BIN_DIR).option("-s, --binsuffix <suffix>", "set bin suffix(" + BIN_Suffix + ")", BIN_Suffix).option("-a, --abidir <path>", "set ABI_DIR(" + ABI_DIR + ")", ABI_DIR).option("-x, --abisuffix <suffix>", "set abi suffix(" + ABI_Suffix + ")", ABI_Suffix).option("-f, --contractsuffix <suffix>", "set contract suffix(" + CONTRACT_Suffix + ")", CONTRACT_Suffix).parse(process.argv);
var argv = process.argv;
var argc = argv.length;
console.log("argv:%j", argv);
console.log("bindir:%j", program.bindir);
console.log("binsuffix:%j", program.binsuffix);
console.log("abidir:%j", program.abidir);
console.log("abisuffix:%j", program.abisuffix);
console.log("contractdir:%j", program.contractdir);
var utils = new _ContractUtils.ContractUtils();
var fread = _ContractUtils.Promise.promisify(fs.readFile);

myDeploy();
