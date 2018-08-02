"use strict";

Object.defineProperty(exports, "__esModule", {
   value: true
});
exports.ContractReader = undefined;

var _createClass = function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; }();
/*var db = require("../modules/ContractDBConn.js");
var conn = new db();
var autoTool = require("../modules/AutoContractTool.js");
var at = new autoTool();*/


require("babel-register");

var _ContractUtils = require("./ContractUtils.js");

function _asyncToGenerator(fn) { return function () { var gen = fn.apply(this, arguments); return new Promise(function (resolve, reject) { function step(key, arg) { try { var info = gen[key](arg); var value = info.value; } catch (error) { reject(error); return; } if (info.done) { resolve(value); } else { return Promise.resolve(value).then(function (value) { step("next", value); }, function (err) { step("throw", err); }); } } return step("next"); }); }; }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

var ContractReader = exports.ContractReader = function () {
   function ContractReader() {
      _classCallCheck(this, ContractReader);

      this.utils = new _ContractUtils.ContractUtils();
   }

   _createClass(ContractReader, [{
      key: "contract",
      value: function () {
         var _ref = _asyncToGenerator(regeneratorRuntime.mark(function _callee(contract_name_) {
            return regeneratorRuntime.wrap(function _callee$(_context) {
               while (1) {
                  switch (_context.prev = _context.next) {
                     case 0:
                        _context.next = 2;
                        return this.utils.contract(contract_name_);

                     case 2:
                        this.contract = _context.sent;
                        return _context.abrupt("return", this.contract);

                     case 4:
                     case "end":
                        return _context.stop();
                  }
               }
            }, _callee, this);
         }));

         function contract(_x) {
            return _ref.apply(this, arguments);
         }

         return contract;
      }()
   }]);

   return ContractReader;
}();

/*
//case
let reader = new AbiReader("SendBalance_V2");
reader.contract();
reader.contract();
*/