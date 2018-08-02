#!/usr/bin/env node
var argv = require("yargs").argv;
var SolidityParser = require("./index.js");


var result;

if (argv.e) {
  result = SolidityParser.parse(argv.e || argv.expression);
} else {
  SolidityParser.parseFile(argv.f || argv.file || argv._[0]);
}

console.log(JSON.stringify(result, null, 2));
