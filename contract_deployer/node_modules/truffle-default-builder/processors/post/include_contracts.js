var path = require("path");
var temp = require("temp").track();
var browserify = require('browserify');
var Pudding = require("ether-pudding");
var async = require("async");
var fs = require("fs");

var digest = function(source_directory, callback) {
  var self = this;
  var digest = "module.exports = {\n"

  Pudding.contractFiles(source_directory, function(err, files) {
    if (err) return callback(err);

    async.eachSeries(files, function(file, finished) {
      Pudding.requireFile(file, function(err, contract) {
        if (err) return finished(err);

        if (process.platform === "win32") {
          file = file.split("\\").join("\\\\")
        }

        digest += "  \"" + contract.contract_name + "\": require(\"" + file + "\"),\n";

        finished();
      });
    }, function(err) {
      if (err) return callback(err);

      digest += "};"
      callback(null, digest);
    });
  });
};

var digestSource = function(options, callback) {
  digest(options.contracts_build_directory, function(err, source) {
    if (err) return callback(err);

    temp.open('contract-digest-', function(err, info) {
      if (err) return callback(err);

      fs.writeFile(info.path, source, {encoding: "utf8"}, function(err) {
        if (err) return callback(err);

        var b = browserify({
          baseDir: options.working_directory,
          standalone: "__contracts__",
          paths: [
            path.resolve(path.join(__dirname, "../", "../", "node_modules")),
            path.resolve(path.join(__dirname, "../", "../", "../"))
          ]
        });
        b.add(info.path);

        var bundleErrorFound = false;

        b.bundle(function(err, buf) {
          // browserify calls the callback more than once on error...
          if (err) {
            if (!bundleErrorFound) {
              bundleErrorFound = true;
              return callback(err);
            } else {
              return;
            }
          } else {
            callback(null, buf);
          }
        });
      });
    });
  });
};

module.exports = function(contents, file, options, process, callback) {
  digestSource(options, function(err, source) {
    if (err) return callback(err);

    callback(null, source.toString() + "\n\n" + contents);
  });
};
