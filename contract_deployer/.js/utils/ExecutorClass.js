"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.Executor = undefined;

var _createClass = function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; }(); //import "babel-register";


var testCase = function () {
  var _ref3 = _asyncToGenerator(regeneratorRuntime.mark(function _callee3(contract_name_) {
    var exec, next;
    return regeneratorRuntime.wrap(function _callee3$(_context3) {
      while (1) {
        switch (_context3.prev = _context3.next) {
          case 0:
            exec = new Executor();
            _context3.next = 3;
            return exec.init(contract_name_);

          case 3:
            _context3.next = 5;
            return exec.next();

          case 5:
            next = _context3.sent;

          case 6:
            if (!(next.done != true)) {
              _context3.next = 12;
              break;
            }

            _context3.next = 9;
            return exec.next();

          case 9:
            next = _context3.sent;
            _context3.next = 6;
            break;

          case 12:
          case "end":
            return _context3.stop();
        }
      }
    }, _callee3, this);
  }));

  return function testCase(_x2) {
    return _ref3.apply(this, arguments);
  };
}();

var test = function () {
  var _ref4 = _asyncToGenerator(regeneratorRuntime.mark(function _callee4(contract_name_) {
    var contractReader, instance, ret;
    return regeneratorRuntime.wrap(function _callee4$(_context4) {
      while (1) {
        switch (_context4.prev = _context4.next) {
          case 0:
            contractReader = new _ContractReaderClass.ContractReader();
            _context4.next = 3;
            return contractReader.contract(contract_name_);

          case 3:
            instance = _context4.sent;

            //await instance.addMember(accounts[0], "Name for default account",aParams);
            //let newP = await instance.newProposal(...Object.values(proposalParams),vParams);
            //let len = instance.proposals.length;
            //let ret = await instance.members("0x2");

            instance.members(2).then(function (ret) {
              console.log(ret);
            });
            //console.log(ret);
            _context4.next = 7;
            return instance.members(1);

          case 7:
            ret = _context4.sent;


            console.log(ret);
            _context4.next = 11;
            return instance.members(0x2);

          case 11:
            ret = _context4.sent;

            console.log(ret);
            return _context4.abrupt("return", instance);

          case 14:
          case "end":
            return _context4.stop();
        }
      }
    }, _callee4, this);
  }));

  return function test(_x3) {
    return _ref4.apply(this, arguments);
  };
}();

var _FuzzReaderClass = require("./FuzzReaderClass.js");

var _ContractReaderClass = require("./ContractReaderClass.js");

var _Account = require("./Account.js");

function _toConsumableArray(arr) { if (Array.isArray(arr)) { for (var i = 0, arr2 = Array(arr.length); i < arr.length; i++) { arr2[i] = arr[i]; } return arr2; } else { return Array.from(arr); } }

function _asyncToGenerator(fn) { return function () { var gen = fn.apply(this, arguments); return new Promise(function (resolve, reject) { function step(key, arg) { try { var info = gen[key](arg); var value = info.value; } catch (error) { reject(error); return; } if (info.done) { resolve(value); } else { return Promise.resolve(value).then(function (value) { step("next", value); }, function (err) { step("throw", err); }); } } return step("next"); }); }; }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

console.log(_Account.accounts);

/*let reader = new FuzzReader("../fuzzSource/Congress.fuzz.backup");
function Test(){
  reader.read();
  reader.nextFunction();
  let func = reader.nextFunction();
  let ret = func.nextTestCase();
  while(ret !=undefined){
      ret = func.nextTestCase();
  }
}
Test();*/
var aParams = { from: _Account.defaultAccount,
  gas: _Account.defaultGas
};
var vParams = { from: _Account.defaultAccount,
  gas: _Account.defaultGas,
  value: _Account.defaultValue
};

var proposalParams = {
  beneficiary: _Account.accounts[0],
  etherAmount: 1,
  JobDescription: 'Some job description',
  transactionBytecode: _Account.web3.sha3('some content')
};

var Executor = exports.Executor = function () {
  function Executor() {
    _classCallCheck(this, Executor);

    this.contractReader = new _ContractReaderClass.ContractReader();
  }

  _createClass(Executor, [{
    key: "init",
    value: function () {
      var _ref = _asyncToGenerator(regeneratorRuntime.mark(function _callee(contract_name_) {
        return regeneratorRuntime.wrap(function _callee$(_context) {
          while (1) {
            switch (_context.prev = _context.next) {
              case 0:
                this.fuzzReader = new _FuzzReaderClass.FuzzReader(contract_name_);
                _context.next = 3;
                return this.contractReader.contract(contract_name_);

              case 3:
                this.instance = _context.sent;

              case 4:
              case "end":
                return _context.stop();
            }
          }
        }, _callee, this);
      }));

      function init(_x) {
        return _ref.apply(this, arguments);
      }

      return init;
    }()
  }, {
    key: "next",
    value: function () {
      var _ref2 = _asyncToGenerator(regeneratorRuntime.mark(function _callee2() {
        var nextFun, fun, contract_fun, nextCase, inputs, ret, _inputs, _ret;

        return regeneratorRuntime.wrap(function _callee2$(_context2) {
          while (1) {
            switch (_context2.prev = _context2.next) {
              case 0:
                nextFun = this.fuzzReader.nextFunction();

                if (!(nextFun.done == true)) {
                  _context2.next = 3;
                  break;
                }

                return _context2.abrupt("return", { value: null, done: true });

              case 3:
                fun = nextFun.value;

                console.log(fun);
                contract_fun = this.instance[fun.fun_name];
                nextCase = fun.nextTestCase();

                if (!(nextCase.done == true)) {
                  _context2.next = 31;
                  break;
                }

                inputs = nextCase.value;
                ret = void 0;
                _context2.prev = 10;

                if (!(fun.payable === 'true')) {
                  _context2.next = 18;
                  break;
                }

                _context2.next = 14;
                return contract_fun(vParams);

              case 14:
                ret = _context2.sent;

                console.log("function:" + fun.fun_name + "  vParams:" + JSON.stringify(vParams));
                _context2.next = 22;
                break;

              case 18:
                _context2.next = 20;
                return contract_fun(aParams);

              case 20:
                ret = _context2.sent;

                console.log("function:" + fun.fun_name + "  aParams:" + JSON.stringify(aParams));

              case 22:
                console.log(ret);
                _context2.next = 31;
                break;

              case 25:
                _context2.prev = 25;
                _context2.t0 = _context2["catch"](10);

                console.error("****failed****");
                console.error(_context2.t0);
                if (fun.payable === 'true') {
                  console.error("function:" + fun.fun_name + "  vParams:" + JSON.stringify(vParams));
                } else {
                  console.error("function:" + fun.fun_name + "  aParams:" + JSON.stringify(aParams));
                }
                console.error("***************");

              case 31:
                if (!(nextCase.done != true)) {
                  _context2.next = 60;
                  break;
                }

                _inputs = nextCase.value;
                _ret = void 0;
                _context2.prev = 34;

                if (!(fun.payable === 'true')) {
                  _context2.next = 42;
                  break;
                }

                _context2.next = 38;
                return contract_fun.apply(undefined, _toConsumableArray(Object.values(_inputs)).concat([vParams]));

              case 38:
                _ret = _context2.sent;

                console.log("function:" + fun.fun_name + "  input:" + Object.values(_inputs) + "  vParams:" + JSON.stringify(vParams));
                _context2.next = 46;
                break;

              case 42:
                _context2.next = 44;
                return contract_fun.apply(undefined, _toConsumableArray(Object.values(_inputs)).concat([aParams]));

              case 44:
                _ret = _context2.sent;

                console.log("function:" + fun.fun_name + "  input:" + Object.values(_inputs) + "  aParams:" + JSON.stringify(aParams));

              case 46:

                console.log(_ret);
                _context2.next = 55;
                break;

              case 49:
                _context2.prev = 49;
                _context2.t1 = _context2["catch"](34);

                console.error("****failed****");
                console.error(_context2.t1);
                if (fun.payable === 'true') {
                  console.error("function:" + fun.fun_name + "  input:" + Object.values(_inputs) + "  vParams:" + JSON.stringify(vParams));
                } else {
                  console.error("function:" + fun.fun_name + "  input:" + Object.values(_inputs) + "  aParams:" + JSON.stringify(aParams));
                }
                console.error("***************");

              case 55:
                _context2.prev = 55;

                nextCase = fun.nextTestCase();
                return _context2.finish(55);

              case 58:
                _context2.next = 31;
                break;

              case 60:
                return _context2.abrupt("return", { value: fun.fun_name, done: false });

              case 61:
              case "end":
                return _context2.stop();
            }
          }
        }, _callee2, this, [[10, 25], [34, 49, 55, 58]]);
      }));

      function next() {
        return _ref2.apply(this, arguments);
      }

      return next;
    }()
  }]);

  return Executor;
}();

var contract_name = "SendBalance_V3";
testCase(contract_name);
//test(contract_name)
//test case