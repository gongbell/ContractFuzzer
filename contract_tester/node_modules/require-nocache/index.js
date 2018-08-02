var Module = require('module');

module.exports = function (module) {
  return function(path) {
    delete require.cache[Module._resolveFilename(path, module)];
    return module.require(path);
  };
}
