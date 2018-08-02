"use strict";

var deployDaoCreatorContract = function () {
    var _ref = _asyncToGenerator(regeneratorRuntime.mark(function _callee() {
        var creator_bin, creator_abi, creatorJSON, amountParams, initialParams, creatorContract;
        return regeneratorRuntime.wrap(function _callee$(_context) {
            while (1) {
                switch (_context.prev = _context.next) {
                    case 0:
                        _context.next = 2;
                        return fread(WorkPlace + "/bin/DAO_Creator.bin");

                    case 2:
                        creator_bin = _context.sent;
                        _context.next = 5;
                        return fread(WorkPlace + "/abi/DAO_Creator.abi");

                    case 5:
                        creator_abi = _context.sent;
                        creatorJSON = new Object();

                        creatorJSON.abi = JSON.parse(creator_abi);
                        creatorJSON.unlinked_binary = creator_bin;
                        creatorJSON.contract_name = "DAO_Creator";

                        amountParams = {
                            from: _Account.defaultAccount,
                            gas: _Account.defaultGas
                        };
                        initialParams = {};
                        _context.next = 14;
                        return utils.deploy_json(creatorJSON, initialParams, amountParams);

                    case 14:
                        creatorContract = _context.sent;
                        return _context.abrupt("return", creatorContract);

                    case 16:
                    case "end":
                        return _context.stop();
                }
            }
        }, _callee, this);
    }));

    return function deployDaoCreatorContract() {
        return _ref.apply(this, arguments);
    };
}();

var deployDaoContract = function () {
    var _ref2 = _asyncToGenerator(regeneratorRuntime.mark(function _callee2() {
        var creatorContract, dao_bin, dao_abi, daoJSON, proposalDeposit, minTokensToCreate, curator, closingTime, initialParams, amountParams, daoContract;
        return regeneratorRuntime.wrap(function _callee2$(_context2) {
            while (1) {
                switch (_context2.prev = _context2.next) {
                    case 0:
                        _context2.next = 2;
                        return deployDaoCreatorContract();

                    case 2:
                        creatorContract = _context2.sent;

                        console.log(creatorContract.address + creatorContract.abi);
                        _context2.next = 6;
                        return fread(WorkPlace + "/bin/DAO.bin");

                    case 6:
                        dao_bin = _context2.sent;
                        _context2.next = 9;
                        return fread(WorkPlace + "/abi/DAO.abi");

                    case 9:
                        dao_abi = _context2.sent;
                        daoJSON = new Object();

                        daoJSON.abi = JSON.parse(dao_abi);
                        daoJSON.unlinked_binary = dao_bin;
                        daoJSON.contract_name = "DAO";
                        proposalDeposit = 30;
                        minTokensToCreate = 100;
                        curator = _Account.defaultAccount;
                        closingTime = timestamp.add(timestamp.now(), "+10m");
                        initialParams = {
                            _curator: curator,
                            _daoCreator: creatorContract.address,
                            _proposalDeposit: _ContractUtils.web3.toWei(proposalDeposit, "ether"),
                            _minTokensToCreate: _ContractUtils.web3.toWei(minTokensToCreate, "ether"),
                            _closingTime: closingTime,
                            _privateCreation: 0
                        };
                        amountParams = {
                            from: _Account.defaultAccount,
                            gas: _Account.defaultGas
                        };
                        _context2.next = 22;
                        return utils.deploy_json(daoJSON, initialParams, amountParams);

                    case 22:
                        daoContract = _context2.sent;

                        console.log(daoContract.address + daoContract.abi);
                        return _context2.abrupt("return", daoContract);

                    case 25:
                    case "end":
                        return _context2.stop();
                }
            }
        }, _callee2, this);
    }));

    return function deployDaoContract() {
        return _ref2.apply(this, arguments);
    };
}();

var _unixTimestamp = require("unix-timestamp");

var timestamp = _interopRequireWildcard(_unixTimestamp);

var _fs = require("fs");

var fs = _interopRequireWildcard(_fs);

var _ContractUtils = require("../utils/ContractUtils.js");

var _Account = require("../utils/Account.js");

function _interopRequireWildcard(obj) { if (obj && obj.__esModule) { return obj; } else { var newObj = {}; if (obj != null) { for (var key in obj) { if (Object.prototype.hasOwnProperty.call(obj, key)) newObj[key] = obj[key]; } } newObj.default = obj; return newObj; } }

function _asyncToGenerator(fn) { return function () { var gen = fn.apply(this, arguments); return new _ContractUtils.Promise(function (resolve, reject) { function step(key, arg) { try { var info = gen[key](arg); var value = info.value; } catch (error) { reject(error); return; } if (info.done) { resolve(value); } else { return _ContractUtils.Promise.resolve(value).then(function (value) { step("next", value); }, function (err) { step("throw", err); }); } } return step("next"); }); }; }

_ContractUtils.dotenv.config();
var WorkPlace = process.env.WorkPlace;

var utils = new _ContractUtils.ContractUtils();
var fread = _ContractUtils.Promise.promisify(fs.readFile);

deployDaoContract();