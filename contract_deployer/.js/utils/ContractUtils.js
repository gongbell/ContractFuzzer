"use strict";

Object.defineProperty(exports, "__esModule", {
    value: true
});
exports.ContractUtils = exports.web3 = exports.truffle_Contract = exports.solc = exports.Web3 = exports.Promise = exports.assert = exports.dotenv = undefined;

var _createClass = function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; }();

var _fs = require("fs");

var fs = _interopRequireWildcard(_fs);

var _DBHelper = require("./DBHelper.js");

function _interopRequireWildcard(obj) { if (obj && obj.__esModule) { return obj; } else { var newObj = {}; if (obj != null) { for (var key in obj) { if (Object.prototype.hasOwnProperty.call(obj, key)) newObj[key] = obj[key]; } } newObj.default = obj; return newObj; } }

function _toConsumableArray(arr) { if (Array.isArray(arr)) { for (var i = 0, arr2 = Array(arr.length); i < arr.length; i++) { arr2[i] = arr[i]; } return arr2; } else { return Array.from(arr); } }

function _asyncToGenerator(fn) { return function () { var gen = fn.apply(this, arguments); return new Promise(function (resolve, reject) { function step(key, arg) { try { var info = gen[key](arg); var value = info.value; } catch (error) { reject(error); return; } if (info.done) { resolve(value); } else { return Promise.resolve(value).then(function (value) { step("next", value); }, function (err) { step("throw", err); }); } } return step("next"); }); }; }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

var dotenv = exports.dotenv = require("dotenv");
var assert = exports.assert = require("assert");
var Promise = exports.Promise = require("bluebird");
var Web3 = exports.Web3 = require("web3");
var solc = exports.solc = require("solc");
var truffle_Contract = exports.truffle_Contract = require("truffle-contract");
dotenv.config();

var WorkPlace = process.env.WorkPlace;
var ABI_DIR = process.env.ABI;

var writeFilePs = Promise.promisify(fs.writeFile);
var HttpRpcAddr = process.env.GethHttpRpcAddr;
var Provider = new Web3.providers.HttpProvider(HttpRpcAddr);
var web3 = exports.web3 = new Web3(Provider);

var accounts = web3.eth.accounts;
var defaultAccount = web3.eth.accounts[0];
var defaultGas = 8000000;
var defaultValue = 100000;

var ContractUtils = exports.ContractUtils = function () {
    function ContractUtils() {
        _classCallCheck(this, ContractUtils);

        this.Password = "123456";
        this.GasPrice = 8000000;

        this.dbHelper = new _DBHelper.DBHelper();
    }

    _createClass(ContractUtils, [{
        key: "_read",
        value: function () {
            var _ref = _asyncToGenerator(regeneratorRuntime.mark(function _callee(file) {
                var readFile, res;
                return regeneratorRuntime.wrap(function _callee$(_context) {
                    while (1) {
                        switch (_context.prev = _context.next) {
                            case 0:
                                readFile = Promise.promisify(fs.readFile);
                                _context.next = 3;
                                return readFile(file);

                            case 3:
                                res = _context.sent;
                                return _context.abrupt("return", res.toString());

                            case 5:
                            case "end":
                                return _context.stop();
                        }
                    }
                }, _callee, this);
            }));

            function _read(_x) {
                return _ref.apply(this, arguments);
            }

            return _read;
        }()
    }, {
        key: "_compile",
        value: function () {
            var _ref2 = _asyncToGenerator(regeneratorRuntime.mark(function _callee2(buffer) {
                var solidity, contract;
                return regeneratorRuntime.wrap(function _callee2$(_context2) {
                    while (1) {
                        switch (_context2.prev = _context2.next) {
                            case 0:
                                if (typeof buffer != "string") buffer = buffer.toString();
                                solidity = Promise.promisify(web3.eth.compile.solidity);
                                _context2.next = 4;
                                return solidity(buffer);

                            case 4:
                                contract = _context2.sent;
                                return _context2.abrupt("return", contract);

                            case 6:
                            case "end":
                                return _context2.stop();
                        }
                    }
                }, _callee2, this);
            }));

            function _compile(_x2) {
                return _ref2.apply(this, arguments);
            }

            return _compile;
        }()
    }, {
        key: "compile",
        value: function () {
            var _ref3 = _asyncToGenerator(regeneratorRuntime.mark(function _callee3(solidityFile) {
                var buffer, contract;
                return regeneratorRuntime.wrap(function _callee3$(_context3) {
                    while (1) {
                        switch (_context3.prev = _context3.next) {
                            case 0:
                                _context3.next = 2;
                                return this._read(solidityFile);

                            case 2:
                                buffer = _context3.sent;
                                _context3.next = 5;
                                return this._compile(buffer);

                            case 5:
                                contract = _context3.sent;
                                return _context3.abrupt("return", contract);

                            case 7:
                            case "end":
                                return _context3.stop();
                        }
                    }
                }, _callee3, this);
            }));

            function compile(_x3) {
                return _ref3.apply(this, arguments);
            }

            return compile;
        }()
    }, {
        key: "compiles",
        value: function () {
            var _ref4 = _asyncToGenerator(regeneratorRuntime.mark(function _callee4(solidityFiles) {
                var contracts, _iteratorNormalCompletion, _didIteratorError, _iteratorError, _iterator, _step, file, contract, _contracts, _contract;

                return regeneratorRuntime.wrap(function _callee4$(_context4) {
                    while (1) {
                        switch (_context4.prev = _context4.next) {
                            case 0:
                                if (!Array.isArray(solidityFiles)) {
                                    _context4.next = 33;
                                    break;
                                }

                                contracts = new Array();
                                _iteratorNormalCompletion = true;
                                _didIteratorError = false;
                                _iteratorError = undefined;
                                _context4.prev = 5;
                                _iterator = solidityFiles[Symbol.iterator]();

                            case 7:
                                if (_iteratorNormalCompletion = (_step = _iterator.next()).done) {
                                    _context4.next = 16;
                                    break;
                                }

                                file = _step.value;
                                _context4.next = 11;
                                return this.compiler(file);

                            case 11:
                                contract = _context4.sent;

                                contracts.push(contract);

                            case 13:
                                _iteratorNormalCompletion = true;
                                _context4.next = 7;
                                break;

                            case 16:
                                _context4.next = 22;
                                break;

                            case 18:
                                _context4.prev = 18;
                                _context4.t0 = _context4["catch"](5);
                                _didIteratorError = true;
                                _iteratorError = _context4.t0;

                            case 22:
                                _context4.prev = 22;
                                _context4.prev = 23;

                                if (!_iteratorNormalCompletion && _iterator.return) {
                                    _iterator.return();
                                }

                            case 25:
                                _context4.prev = 25;

                                if (!_didIteratorError) {
                                    _context4.next = 28;
                                    break;
                                }

                                throw _iteratorError;

                            case 28:
                                return _context4.finish(25);

                            case 29:
                                return _context4.finish(22);

                            case 30:
                                return _context4.abrupt("return", contracts);

                            case 33:
                                _contracts = new Array();
                                _context4.next = 36;
                                return this.compile(solidityFiles);

                            case 36:
                                _contract = _context4.sent;

                                assert.ok(_contract, "contract undefined");

                                _contracts.push(_contract);
                                return _context4.abrupt("return", _contracts);

                            case 40:
                            case "end":
                                return _context4.stop();
                        }
                    }
                }, _callee4, this, [[5, 18, 22, 30], [23,, 25, 29]]);
            }));

            function compiles(_x4) {
                return _ref4.apply(this, arguments);
            }

            return compiles;
        }()
    }, {
        key: "deploy",
        value: function () {
            var _ref5 = _asyncToGenerator(regeneratorRuntime.mark(function _callee5(contract, initalParams, amountParams) {
                var owner, aParams, iParams, contract_name_, abi_, code_, name, MyContract, instance, res, abiJSON;
                return regeneratorRuntime.wrap(function _callee5$(_context5) {
                    while (1) {
                        switch (_context5.prev = _context5.next) {
                            case 0:
                                owner = this.defaultAccount;
                                aParams = void 0;
                                iParams = void 0;

                                if ("undefined" == typeof initalParams) {
                                    iParams = new Array();
                                } else {
                                    iParams = initalParams;
                                }
                                if ("undefined" == typeof amountParams) {
                                    aParams = {
                                        from: defaultAccount,
                                        //value:web3.toWei(2,"ether"),
                                        gas: defaultGas
                                    };
                                    //console.log(aParams);
                                } else {
                                    aParams = amountParams;
                                }
                                contract_name_ = void 0;
                                abi_ = void 0;
                                code_ = void 0;

                                for (name in contract) {
                                    contract_name_ = name;
                                    abi_ = contract[name].info.abiDefinition;
                                    code_ = contract[name].code;
                                }
                                MyContract = truffle_Contract({
                                    contract_name: contract_name_,
                                    abi: abi_,
                                    unlinked_binary: code_,
                                    default_network: 1900
                                });

                                MyContract.setProvider(Provider);
                                _context5.next = 13;
                                return MyContract.new.apply(MyContract, _toConsumableArray(Object.values(iParams)).concat([aParams]));

                            case 13:
                                instance = _context5.sent;
                                _context5.next = 16;
                                return this.dbHelper.regContract(contract_name_, instance.address, instance.transactionHash, abi_, code_, owner);

                            case 16:
                                res = _context5.sent;

                                console.log(res);
                                abiJSON = {
                                    abi: abi_
                                };
                                _context5.next = 21;
                                return writeFilePs(ABI_DIR + "/" + contract_name_ + ".json", JSON.stringify(abiJSON));

                            case 21:
                                console.log(res);
                                return _context5.abrupt("return", instance);

                            case 23:
                            case "end":
                                return _context5.stop();
                        }
                    }
                }, _callee5, this);
            }));

            function deploy(_x5, _x6, _x7) {
                return _ref5.apply(this, arguments);
            }

            return deploy;
        }()
    }, {
        key: "compileAndDeploy",
        value: function () {
            var _ref6 = _asyncToGenerator(regeneratorRuntime.mark(function _callee6(solidityFiles) {
                var contracts, entities, _iteratorNormalCompletion2, _didIteratorError2, _iteratorError2, _iterator2, _step2, contract, entity;

                return regeneratorRuntime.wrap(function _callee6$(_context6) {
                    while (1) {
                        switch (_context6.prev = _context6.next) {
                            case 0:
                                _context6.next = 2;
                                return this.compiles(solidityFiles);

                            case 2:
                                contracts = _context6.sent;
                                entities = new Array();
                                _iteratorNormalCompletion2 = true;
                                _didIteratorError2 = false;
                                _iteratorError2 = undefined;
                                _context6.prev = 7;
                                _iterator2 = contracts[Symbol.iterator]();

                            case 9:
                                if (_iteratorNormalCompletion2 = (_step2 = _iterator2.next()).done) {
                                    _context6.next = 18;
                                    break;
                                }

                                contract = _step2.value;
                                _context6.next = 13;
                                return this.deploy(contract);

                            case 13:
                                entity = _context6.sent;

                                entities.push(entity);

                            case 15:
                                _iteratorNormalCompletion2 = true;
                                _context6.next = 9;
                                break;

                            case 18:
                                _context6.next = 24;
                                break;

                            case 20:
                                _context6.prev = 20;
                                _context6.t0 = _context6["catch"](7);
                                _didIteratorError2 = true;
                                _iteratorError2 = _context6.t0;

                            case 24:
                                _context6.prev = 24;
                                _context6.prev = 25;

                                if (!_iteratorNormalCompletion2 && _iterator2.return) {
                                    _iterator2.return();
                                }

                            case 27:
                                _context6.prev = 27;

                                if (!_didIteratorError2) {
                                    _context6.next = 30;
                                    break;
                                }

                                throw _iteratorError2;

                            case 30:
                                return _context6.finish(27);

                            case 31:
                                return _context6.finish(24);

                            case 32:
                                return _context6.abrupt("return", entities);

                            case 33:
                            case "end":
                                return _context6.stop();
                        }
                    }
                }, _callee6, this, [[7, 20, 24, 32], [25,, 27, 31]]);
            }));

            function compileAndDeploy(_x8) {
                return _ref6.apply(this, arguments);
            }

            return compileAndDeploy;
        }()
    }, {
        key: "_instance",
        value: function () {
            var _ref7 = _asyncToGenerator(regeneratorRuntime.mark(function _callee7(contract) {
                var name_, abi_, code_, address_, MyContract, instance;
                return regeneratorRuntime.wrap(function _callee7$(_context7) {
                    while (1) {
                        switch (_context7.prev = _context7.next) {
                            case 0:
                                name_ = contract.name;
                                abi_ = JSON.parse(unescape(contract.abi));
                                code_ = unescape(contract.code);
                                address_ = contract.address;
                                MyContract = truffle_Contract({
                                    contract_name: name_,
                                    abi: abi_,
                                    unlinked_binary: code_,
                                    network_id: 1900,
                                    address: address_
                                });

                                MyContract.setProvider(Provider);
                                _context7.next = 8;
                                return MyContract.deployed();

                            case 8:
                                instance = _context7.sent;
                                return _context7.abrupt("return", instance);

                            case 10:
                            case "end":
                                return _context7.stop();
                        }
                    }
                }, _callee7, this);
            }));

            function _instance(_x9) {
                return _ref7.apply(this, arguments);
            }

            return _instance;
        }()
    }, {
        key: "contract",
        value: function () {
            var _ref8 = _asyncToGenerator(regeneratorRuntime.mark(function _callee8(name) {
                var contracts, contract, instance;
                return regeneratorRuntime.wrap(function _callee8$(_context8) {
                    while (1) {
                        switch (_context8.prev = _context8.next) {
                            case 0:
                                _context8.next = 2;
                                return this.dbHelper.contracts(name);

                            case 2:
                                contracts = _context8.sent;

                                assert.ok(contracts, "[ERROR]:contract " + name + " not exist in db!");
                                contract = contracts[0];
                                // console.log(contract);

                                _context8.next = 7;
                                return this._instance(contract);

                            case 7:
                                instance = _context8.sent;
                                return _context8.abrupt("return", instance);

                            case 9:
                            case "end":
                                return _context8.stop();
                        }
                    }
                }, _callee8, this);
            }));

            function contract(_x10) {
                return _ref8.apply(this, arguments);
            }

            return contract;
        }()
    }, {
        key: "contracts",
        value: function () {
            var _ref9 = _asyncToGenerator(regeneratorRuntime.mark(function _callee9(name) {
                var contracts, instances, _iteratorNormalCompletion3, _didIteratorError3, _iteratorError3, _iterator3, _step3, contract, instance;

                return regeneratorRuntime.wrap(function _callee9$(_context9) {
                    while (1) {
                        switch (_context9.prev = _context9.next) {
                            case 0:
                                console.log("contracts()");
                                _context9.next = 3;
                                return this.dbHelper.contracts(name);

                            case 3:
                                contracts = _context9.sent;

                                assert.ok(contracts, "[ERROR]:contract " + name + " not exist in db!");
                                instances = new Array();
                                _iteratorNormalCompletion3 = true;
                                _didIteratorError3 = false;
                                _iteratorError3 = undefined;
                                _context9.prev = 9;
                                _iterator3 = contracts[Symbol.iterator]();

                            case 11:
                                if (_iteratorNormalCompletion3 = (_step3 = _iterator3.next()).done) {
                                    _context9.next = 20;
                                    break;
                                }

                                contract = _step3.value;
                                _context9.next = 15;
                                return this._instance(contract);

                            case 15:
                                instance = _context9.sent;

                                instances.push(instance);

                            case 17:
                                _iteratorNormalCompletion3 = true;
                                _context9.next = 11;
                                break;

                            case 20:
                                _context9.next = 26;
                                break;

                            case 22:
                                _context9.prev = 22;
                                _context9.t0 = _context9["catch"](9);
                                _didIteratorError3 = true;
                                _iteratorError3 = _context9.t0;

                            case 26:
                                _context9.prev = 26;
                                _context9.prev = 27;

                                if (!_iteratorNormalCompletion3 && _iterator3.return) {
                                    _iterator3.return();
                                }

                            case 29:
                                _context9.prev = 29;

                                if (!_didIteratorError3) {
                                    _context9.next = 32;
                                    break;
                                }

                                throw _iteratorError3;

                            case 32:
                                return _context9.finish(29);

                            case 33:
                                return _context9.finish(26);

                            case 34:
                                return _context9.abrupt("return", instances);

                            case 35:
                            case "end":
                                return _context9.stop();
                        }
                    }
                }, _callee9, this, [[9, 22, 26, 34], [27,, 29, 33]]);
            }));

            function contracts(_x11) {
                return _ref9.apply(this, arguments);
            }

            return contracts;
        }()
    }, {
        key: "show",
        value: function show() {
            console.log("hello ContractUtils");
        }
    }, {
        key: "test_promise",
        value: function () {
            var _ref10 = _asyncToGenerator(regeneratorRuntime.mark(function _callee10() {
                var buffer, contract;
                return regeneratorRuntime.wrap(function _callee10$(_context10) {
                    while (1) {
                        switch (_context10.prev = _context10.next) {
                            case 0:
                                _context10.next = 2;
                                return this._read("/home/liuye/repositories/FuzzExecutor/FuzzExecutor/nodejs/Contracts/SendBalance_V2.sol");

                            case 2:
                                buffer = _context10.sent;
                                _context10.next = 5;
                                return this._compile(buffer);

                            case 5:
                                contract = _context10.sent;

                                console.log(contract);

                            case 7:
                            case "end":
                                return _context10.stop();
                        }
                    }
                }, _callee10, this);
            }));

            function test_promise() {
                return _ref10.apply(this, arguments);
            }

            return test_promise;
        }()
    }, {
        key: "deploy_json",
        value: function () {
            var _ref11 = _asyncToGenerator(regeneratorRuntime.mark(function _callee11(contractJSON, initialParams, amountParams) {
                var owner, aParams, iParams, contract_name_, abi_, code_, MyContract, instance, res;
                return regeneratorRuntime.wrap(function _callee11$(_context11) {
                    while (1) {
                        switch (_context11.prev = _context11.next) {
                            case 0:
                                owner = this.defaultAccount;
                                aParams = void 0;
                                iParams = void 0;

                                if ("undefined" == typeof initialParams || initialParams == null) {
                                    iParams = new Array();
                                } else {
                                    iParams = initialParams;
                                }
                                if ("undefined" == typeof amountParams || amountParams == null) {
                                    aParams = {
                                        from: defaultAccount,
                                        //value:web3.toWei(2,"ether"),
                                        gas: defaultGas
                                    };
                                    //console.log(aParams);
                                } else {
                                    aParams = amountParams;
                                }
                                contract_name_ = void 0;
                                abi_ = void 0;
                                code_ = void 0;

                                contract_name_ = contractJSON.contract_name;
                                abi_ = contractJSON.abi;
                                code_ = contractJSON.unlinked_binary;
                                MyContract = truffle_Contract({
                                    contract_name: contract_name_,
                                    abi: abi_,
                                    unlinked_binary: code_,
                                    default_network: 1900
                                });

                                MyContract.setProvider(Provider);
                                _context11.next = 15;
                                return MyContract.new.apply(MyContract, _toConsumableArray(Object.values(iParams)).concat([aParams]));

                            case 15:
                                instance = _context11.sent;
                                _context11.next = 18;
                                return this.dbHelper.regContract(contract_name_, instance.address, instance.transactionHash, abi_, code_, owner);

                            case 18:
                                res = _context11.sent;
                                return _context11.abrupt("return", instance);

                            case 20:
                            case "end":
                                return _context11.stop();
                        }
                    }
                }, _callee11, this);
            }));

            function deploy_json(_x12, _x13, _x14) {
                return _ref11.apply(this, arguments);
            }

            return deploy_json;
        }()
    }]);

    return ContractUtils;
}();
//Use Case
//const utils = new ContractUtils();
//utils.test_promise();
//utils._read("./Contracts/SendBalance_V3.sol");