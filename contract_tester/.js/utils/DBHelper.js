"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});

var _createClass = function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; }();

function _asyncToGenerator(fn) { return function () { var gen = fn.apply(this, arguments); return new Promise(function (resolve, reject) { function step(key, arg) { try { var info = gen[key](arg); var value = info.value; } catch (error) { reject(error); return; } if (info.done) { resolve(value); } else { return Promise.resolve(value).then(function (value) { step("next", value); }, function (err) { step("throw", err); }); } } return step("next"); }); }; }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

require("babel-register");
require('dotenv').config();
var assert = require("assert");
var mysql = require("mysql");

var DB_USER = process.env.DB_USER;
var PASSWORD = process.env.PASSWORD;

var DB_HOST = process.env.DB_HOST;
var DB_NAME = process.env.DB_NAME;

var PORT = process.env.PORT;
var LIMIT = process.env.DB_LIMIT;
var pool = mysql.createPool({
  connectionLimit: LIMIT,
  host: DB_HOST,
  database: DB_NAME,
  user: DB_USER,
  password: PASSWORD,
  port: PORT
});

var Promise = require("bluebird");
var queryAsync = Promise.promisify(pool.query, { context: pool, multiArgs: true });

var DBHelper = exports.DBHelper = function () {
  function DBHelper() {
    _classCallCheck(this, DBHelper);
  }

  _createClass(DBHelper, [{
    key: "isConnected",
    value: function () {
      var _ref = _asyncToGenerator(regeneratorRuntime.mark(function _callee() {
        var getConnection, conn;
        return regeneratorRuntime.wrap(function _callee$(_context) {
          while (1) {
            switch (_context.prev = _context.next) {
              case 0:
                getConnection = Promise.promisify(pool.getConnection, { context: pool });
                _context.next = 3;
                return getConnection();

              case 3:
                conn = _context.sent;


                assert.ok(conn, "[ERROR]: connect failed!");
                return _context.abrupt("return", true);

              case 6:
              case "end":
                return _context.stop();
            }
          }
        }, _callee, this);
      }));

      function isConnected() {
        return _ref.apply(this, arguments);
      }

      return isConnected;
    }()
  }, {
    key: "_query",
    value: function () {
      var _ref2 = _asyncToGenerator(regeneratorRuntime.mark(function _callee2() {
        var res,
            _args2 = arguments;
        return regeneratorRuntime.wrap(function _callee2$(_context2) {
          while (1) {
            switch (_context2.prev = _context2.next) {
              case 0:
                _context2.next = 2;
                return queryAsync.apply(this, _args2);

              case 2:
                res = _context2.sent;
                return _context2.abrupt("return", res[0]);

              case 4:
              case "end":
                return _context2.stop();
            }
          }
        }, _callee2, this);
      }));

      function _query() {
        return _ref2.apply(this, arguments);
      }

      return _query;
    }()
  }, {
    key: "log",
    value: function () {
      var _ref3 = _asyncToGenerator(regeneratorRuntime.mark(function _callee3(name, error, type) {
        var query, post, res;
        return regeneratorRuntime.wrap(function _callee3$(_context3) {
          while (1) {
            switch (_context3.prev = _context3.next) {
              case 0:
                query = "insert into log Set ?";
                post = { "name": name, "error": escape(error.toString()), "type": type };
                _context3.next = 4;
                return this._query(query, post);

              case 4:
                res = _context3.sent;

                //    console.log(res);
                assert.ok(res.length);
                return _context3.abrupt("return", res);

              case 7:
              case "end":
                return _context3.stop();
            }
          }
        }, _callee3, this);
      }));

      function log(_x, _x2, _x3) {
        return _ref3.apply(this, arguments);
      }

      return log;
    }()
  }, {
    key: "regContract",
    value: function () {
      var _ref4 = _asyncToGenerator(regeneratorRuntime.mark(function _callee4(name, address, txhash, abi, code, owner) {
        var query, post, res;
        return regeneratorRuntime.wrap(function _callee4$(_context4) {
          while (1) {
            switch (_context4.prev = _context4.next) {
              case 0:
                query = "insert into register Set ?";
                post = { "name": name,
                  "address": address,
                  "txhash": txhash,
                  "abi": escape(JSON.stringify(abi)),
                  "code": escape(JSON.stringify(code)),
                  "owner": owner };
                _context4.next = 4;
                return this._query(query, post);

              case 4:
                res = _context4.sent;

                console.log(res.affectedRows);
                assert.ok(res.affectedRows, "[ERROR]:insert  " + name + " into register failed");
                return _context4.abrupt("return", res);

              case 8:
              case "end":
                return _context4.stop();
            }
          }
        }, _callee4, this);
      }));

      function regContract(_x4, _x5, _x6, _x7, _x8, _x9) {
        return _ref4.apply(this, arguments);
      }

      return regContract;
    }()
  }, {
    key: "contracts",
    value: function () {
      var _ref5 = _asyncToGenerator(regeneratorRuntime.mark(function _callee5(name) {
        var query, post, res;
        return regeneratorRuntime.wrap(function _callee5$(_context5) {
          while (1) {
            switch (_context5.prev = _context5.next) {
              case 0:
                //  let query ="select name,abi,address, from register where ?";
                query = "select * from register where ?";
                post = { "name": name };
                _context5.next = 4;
                return this._query(query, post);

              case 4:
                res = _context5.sent;

                //  console.log(res);
                assert.ok(res.length, "[ERROR]:contract " + name + "not exist in db!");
                return _context5.abrupt("return", res);

              case 7:
              case "end":
                return _context5.stop();
            }
          }
        }, _callee5, this);
      }));

      function contracts(_x10) {
        return _ref5.apply(this, arguments);
      }

      return contracts;
    }()
  }]);

  return DBHelper;
}();

//let helper = new DBHelper();
//helper.isConnected();