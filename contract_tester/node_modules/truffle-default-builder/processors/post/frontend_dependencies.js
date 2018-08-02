var path = require("path");

module.exports = function(contents, file, options, process, callback) {

  var includes = [
    path.join(__dirname, "../../frontend/bluebird.js"),
    path.join(__dirname, "../../frontend/web3.min.js")
  ];

  process(includes, function(err, processed) {
    if (err != null) {
      callback(err);
      return;
    }

    callback(null, processed + "\n\n" + contents);
  });
};
