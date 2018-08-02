"use strict";

var listFiles = function () {
  var _ref = _asyncToGenerator(regeneratorRuntime.mark(function _callee(DIR) {
    var files, _iteratorNormalCompletion, _didIteratorError, _iteratorError, _iterator, _step, file;

    return regeneratorRuntime.wrap(function _callee$(_context) {
      while (1) {
        switch (_context.prev = _context.next) {
          case 0:
            _context.next = 2;
            return readdir(DIR);

          case 2:
            files = _context.sent;
            _iteratorNormalCompletion = true;
            _didIteratorError = false;
            _iteratorError = undefined;
            _context.prev = 6;

            for (_iterator = files[Symbol.iterator](); !(_iteratorNormalCompletion = (_step = _iterator.next()).done); _iteratorNormalCompletion = true) {
              file = _step.value;

              console.log(file);
            }_context.next = 14;
            break;

          case 10:
            _context.prev = 10;
            _context.t0 = _context["catch"](6);
            _didIteratorError = true;
            _iteratorError = _context.t0;

          case 14:
            _context.prev = 14;
            _context.prev = 15;

            if (!_iteratorNormalCompletion && _iterator.return) {
              _iterator.return();
            }

          case 17:
            _context.prev = 17;

            if (!_didIteratorError) {
              _context.next = 20;
              break;
            }

            throw _iteratorError;

          case 20:
            return _context.finish(17);

          case 21:
            return _context.finish(14);

          case 22:
          case "end":
            return _context.stop();
        }
      }
    }, _callee, this, [[6, 10, 14, 22], [15,, 17, 21]]);
  }));

  return function listFiles(_x) {
    return _ref.apply(this, arguments);
  };
}();

var deploy = function () {
  var _ref2 = _asyncToGenerator(regeneratorRuntime.mark(function _callee2(DIR) {
    var files, _iteratorNormalCompletion2, _didIteratorError2, _iteratorError2, _iterator2, _step2, file, contractJSON, instance;

    return regeneratorRuntime.wrap(function _callee2$(_context2) {
      while (1) {
        switch (_context2.prev = _context2.next) {
          case 0:
            _context2.next = 2;
            return readdir(DIR);

          case 2:
            files = _context2.sent;
            _iteratorNormalCompletion2 = true;
            _didIteratorError2 = false;
            _iteratorError2 = undefined;
            _context2.prev = 6;
            _iterator2 = files[Symbol.iterator]();

          case 8:
            if (_iteratorNormalCompletion2 = (_step2 = _iterator2.next()).done) {
              _context2.next = 20;
              break;
            }

            file = _step2.value;
            _context2.next = 12;
            return fread(DIR + "/" + file);

          case 12:
            contractJSON = _context2.sent;

            contractJSON = JSON.parse(contractJSON);
            _context2.next = 16;
            return utils.deploy_json(contractJSON);

          case 16:
            instance = _context2.sent;

          case 17:
            _iteratorNormalCompletion2 = true;
            _context2.next = 8;
            break;

          case 20:
            _context2.next = 26;
            break;

          case 22:
            _context2.prev = 22;
            _context2.t0 = _context2["catch"](6);
            _didIteratorError2 = true;
            _iteratorError2 = _context2.t0;

          case 26:
            _context2.prev = 26;
            _context2.prev = 27;

            if (!_iteratorNormalCompletion2 && _iterator2.return) {
              _iterator2.return();
            }

          case 29:
            _context2.prev = 29;

            if (!_didIteratorError2) {
              _context2.next = 32;
              break;
            }

            throw _iteratorError2;

          case 32:
            return _context2.finish(29);

          case 33:
            return _context2.finish(26);

          case 34:
          case "end":
            return _context2.stop();
        }
      }
    }, _callee2, this, [[6, 22, 26, 34], [27,, 29, 33]]);
  }));

  return function deploy(_x2) {
    return _ref2.apply(this, arguments);
  };
}();
//deploy(JSON_DIR);


var deployFile = function () {
  var _ref3 = _asyncToGenerator(regeneratorRuntime.mark(function _callee3(file, initalParams, amountParams) {
    var contractJSON, instance;
    return regeneratorRuntime.wrap(function _callee3$(_context3) {
      while (1) {
        switch (_context3.prev = _context3.next) {
          case 0:
            _context3.next = 2;
            return fread(file);

          case 2:
            contractJSON = _context3.sent;

            contractJSON = JSON.parse(contractJSON);
            _context3.next = 6;
            return utils.deploy_json(contractJSON, initalParams, amountParams);

          case 6:
            instance = _context3.sent;

          case 7:
          case "end":
            return _context3.stop();
        }
      }
    }, _callee3, this);
  }));

  return function deployFile(_x3, _x4, _x5) {
    return _ref3.apply(this, arguments);
  };
}();

var _fs = require("fs");

var fs = _interopRequireWildcard(_fs);

var _ContractUtils = require("../utils/ContractUtils.js");

var _Account = require("../utils/Account.js");

function _interopRequireWildcard(obj) { if (obj && obj.__esModule) { return obj; } else { var newObj = {}; if (obj != null) { for (var key in obj) { if (Object.prototype.hasOwnProperty.call(obj, key)) newObj[key] = obj[key]; } } newObj.default = obj; return newObj; } }

function _asyncToGenerator(fn) { return function () { var gen = fn.apply(this, arguments); return new _ContractUtils.Promise(function (resolve, reject) { function step(key, arg) { try { var info = gen[key](arg); var value = info.value; } catch (error) { reject(error); return; } if (info.done) { resolve(value); } else { return _ContractUtils.Promise.resolve(value).then(function (value) { step("next", value); }, function (err) { step("throw", err); }); } } return step("next"); }); }; }

_ContractUtils.dotenv.config();
var WorkPlace = process.env.WorkPlace;
var JSON_DIR = process.env.Contract_JSON;

var utils = new _ContractUtils.ContractUtils();
var readdir = _ContractUtils.Promise.promisify(fs.readdir);
var fread = _ContractUtils.Promise.promisify(fs.readFile);

var initialParams = {
  minimumQuorumForProposals: 3,
  minutesForDebate: 5,
  marginOfVotesForMajority: 2,
  congressLeader: _Account.accounts[0]
};
var amountParams = {
  from: _Account.defaultAccount,
  gas: _Account.defaultGas,
  value: _Account.defaultValue
};

deployFile(JSON_DIR + "/" + "Congress.json", initialParams, amountParams);
//deployFile(JSON_DIR+"/"+"SendBalance_V3.json");
//deployFile(JSON_DIR+"/"+"SendBalance_V2.json");