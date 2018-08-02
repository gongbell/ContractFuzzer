"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.FuzzReader = exports.Function = undefined;

var _createClass = function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; }();

require("babel-register");

var _fs = require("fs");

var fs = _interopRequireWildcard(_fs);

function _interopRequireWildcard(obj) { if (obj && obj.__esModule) { return obj; } else { var newObj = {}; if (obj != null) { for (var key in obj) { if (Object.prototype.hasOwnProperty.call(obj, key)) newObj[key] = obj[key]; } } newObj.default = obj; return newObj; } }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

var fwrite = fs.writeFileSync;
var fread = fs.readFileSync;

var dotenv = require("dotenv");
dotenv.config();
var WorkPlace = process.env.WorkPlace;
var FUZZ_DIR = process.env.FUZZ;
var FUZZ_suffix = process.env.FUZZ_suffix;
var ABI_DIR = process.env.ABI;
var ABI_suffix = process.env.ABI_suffix;

var Function = exports.Function = function () {
  function Function(contract_name_, fun_name_, argNames_, inputArray_, len_, payable_) {
    _classCallCheck(this, Function);

    this.contract_name = contract_name_;
    this.argNames = argNames_;
    this.fun_name = fun_name_;
    this.inputArray = inputArray_;
    this.len = len_;
    this.payable = payable_;
    if (this.inputArray == null) this.inputArray = new Array();
    this.iter = this.inputArray[Symbol.iterator]();
  }

  _createClass(Function, [{
    key: "nextTestCase",
    value: function nextTestCase() {
      var next = this.iter.next();
      if (next.done) {
        return { value: null, done: true };
      } else {
        var Params = new Object();

        var input = next.value;
        var exp = "";
        for (var i = 0; i < this.len; i++) {
          exp = exp + "Params." + this.argNames[i] + "=\"" + input[i] + "\";";
        }
        eval(exp);

        var expression = void 0;
        var val_gas_params = new Object();
        if (this.payable) {
          var default_value = 10000000;
          var default_gas = 8000000;
          val_gas_params.value = default_value;
          val_gas_params.gas = default_gas;
          //val_gas_params.from = web3.defaultAccount;
        } else {
          var _default_value = 0;
          var _default_gas = 0;
          val_gas_params.value = _default_value;
          val_gas_params.gas = _default_gas;
        }
        expression = this.contract_name + "." + this.fun_name + "(...Object.values(Params),val_gas_params);";
        // let expression = this.contract_name+"."+this.fun+"("+ret+");";
        // console.log(expression);
        // eval("let ret = await "+expression);

        //  return expression;
        var ret = { value: Params, done: false };
        return ret;
      }
    }
  }]);

  return Function;
}();

function test(a, b) {
  console.log(a + b);
}

var FuzzReader = exports.FuzzReader = function () {
  function FuzzReader(contract_name_) {
    _classCallCheck(this, FuzzReader);

    this.fuzzPath = FUZZ_DIR + "/" + contract_name_ + FUZZ_suffix;
    this._read();
  }

  _createClass(FuzzReader, [{
    key: "_read",
    value: function _read() {
      var fuzzSource = JSON.parse(fread(this.fuzzPath));
      this.contract_name = fuzzSource.contract_name;
      this.fuzzFunsJSON = fuzzSource.funs_fuzzing;
      //console.log(this.fuzzFunsJSON);
      this.iter = this.fuzzFunsJSON[Symbol.iterator]();
    }
  }, {
    key: "_cartProduct1",
    value: function _cartProduct1(set) {
      console.log("_cartProduct1");
      var retSet = new Array();
      var iter = set[Symbol.iterator]();
      var next1 = iter.next();
      //  console.log(next1);
      while (!next1.done) {
        var arr = new Array(1);
        arr[0] = next1.value;
        retSet.push(arr);
        next1 = iter.next();
      }
      return retSet;
    }
  }, {
    key: "_cartProduct2",
    value: function _cartProduct2(set1, set2) {
      console.log("_cartProduct2");
      var retSet = new Array();

      var iter1 = set1[Symbol.iterator]();
      var next1 = iter1.next();
      while (!next1.done) {
        var iter2 = set2[Symbol.iterator]();
        var next2 = iter2.next();
        while (!next2.done) {
          var arr = new Array();
          var arr_value1 = next1.value;
          var value2 = next2.value;
          //  console.log("value1:"+arr_value1);
          //  console.log("value2:"+value2);
          arr = arr.concat(arr_value1);
          arr.push(value2);
          retSet.push(arr);
          //console.log("arr:"+arr);
          next2 = iter2.next();
        }
        next1 = iter1.next();
      }
      return retSet;
    }
  }, {
    key: "nextFunction",
    value: function nextFunction() {
      var next = this.iter.next();
      if (next.done) {
        return { value: null, done: true };
      } else {
        var value = next.value;
        this.fun = value.function;
        this.len = value.argc;
        this.inputs = value.inputs;
        this.payable = value.payable;
        this.argNames = new Array();
        if (null == this.inputs) {
          this.inputs = new Array();
        }
        var _iteratorNormalCompletion = true;
        var _didIteratorError = false;
        var _iteratorError = undefined;

        try {
          for (var _iterator = this.inputs[Symbol.iterator](), _step; !(_iteratorNormalCompletion = (_step = _iterator.next()).done); _iteratorNormalCompletion = true) {
            var filed = _step.value;

            this.argNames.push(filed.name);
          }
        } catch (err) {
          _didIteratorError = true;
          _iteratorError = err;
        } finally {
          try {
            if (!_iteratorNormalCompletion && _iterator.return) {
              _iterator.return();
            }
          } finally {
            if (_didIteratorError) {
              throw _iteratorError;
            }
          }
        }

        this.outputs = value.output; //value["outputs"]
        //console.log( {"fun_name":this.fun,"len":this.len,"inputs":this.inputs,"outputs":this.outputs});
        this.inputArray = this._getInput(this.inputs);
        //  console.log(this.inputArray);
        var func = new Function(this.contract_name, this.fun, this.argNames, this.inputArray, this.len, this.payable);
        var ret = { value: func, done: false };
        return ret;
      }
    }
  }, {
    key: "_getInput",
    value: function _getInput(inputs) {
      var fuzzArray = new Array(); //new Array(total_len)
      var len = inputs.length;
      if (inputs.length == 0) {
        return new Array();
      }
      //    console.log("len:"+len);
      var fuzzingVectors = new Array();
      var total_len = 1;
      var _iteratorNormalCompletion2 = true;
      var _didIteratorError2 = false;
      var _iteratorError2 = undefined;

      try {
        for (var _iterator2 = this.inputs[Symbol.iterator](), _step2; !(_iteratorNormalCompletion2 = (_step2 = _iterator2.next()).done); _iteratorNormalCompletion2 = true) {
          var typeJSON = _step2.value;

          var fuzzing = typeJSON.fuzzing;
          total_len *= fuzzing.length;
          fuzzingVectors.push(fuzzing);
        }
      } catch (err) {
        _didIteratorError2 = true;
        _iteratorError2 = err;
      } finally {
        try {
          if (!_iteratorNormalCompletion2 && _iterator2.return) {
            _iterator2.return();
          }
        } finally {
          if (_didIteratorError2) {
            throw _iteratorError2;
          }
        }
      }

      var firstFuzzing = fuzzingVectors[0];
      var leftSet = this._cartProduct1(firstFuzzing);

      for (var i = 1; i < len; i++) {
        leftSet = this._cartProduct2(leftSet, fuzzingVectors[i]);
      }
      fuzzArray = leftSet;
      return fuzzArray;
    }
  }]);

  return FuzzReader;
}();
/******************/
//Case1
/*
const reader = new FuzzReader("../fuzzSource/Congress.fuzz.backup");
reader.read();
reader.nextFunction();
let func = reader.nextFunction();
let ret = func.nextTestCase();
while(ret !=undefined){
    ret = func.nextTestCase();
}*/