var babel = require('babel-core');
var es2015 = require("babel-preset-es2015");
var syntax_jsx = require("babel-plugin-syntax-jsx");
var transform_jsx = require("babel-plugin-transform-react-jsx")
var stagetwo = require("babel-preset-stage-2")

module.exports = function(contents, file, options, process, callback) {
  try {
    var code = babel.transform(contents, {
      filename: file,
      compact: false,
      presets: [
        es2015,
        //react,
        stagetwo
      ],
      plugins: [
        syntax_jsx,
        transform_jsx
      ],
      ast: false,
    }).code;
    callback(null, code);
  } catch(e) {
    callback(e);
  }
};
