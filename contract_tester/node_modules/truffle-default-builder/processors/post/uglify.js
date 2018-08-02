var UglifyJS = require("uglify-js");
var fs = require("fs");
var temp = require("temp");

module.exports = function(contents, file, options, process, callback) {
  try {
    var code = UglifyJS.minify(contents, {fromString: true}).code;
    callback(null, code);
  } catch(ex) {
    // Error handling taken from:
    // https://github.com/mishoo/UglifyJS2/commit/a89d23331802cabb690ac817a5a7c93144a4f9c9
    if (ex instanceof UglifyJS.JS_Parse_Error) {
      // Turn it into a string so it's seen as a build error.

      var tempfile = temp.path({prefix: "output-", suffix: '.js'});

      fs.writeFileSync(tempfile, contents, {encoding: "utf8"});

      ex = "UglifyJS parse error at " + file + ":" + ex.line + "," + ex.col + "\n\n" + ex.message + "\n\n" + "Output file written to: " + tempfile + "\n\n" + "Note that some ES6 keywords like `let` cause UglifyJS to fail.\nSee bug report here: https://github.com/mishoo/UglifyJS2/issues/448";
    }
    callback(ex);
  }
};
