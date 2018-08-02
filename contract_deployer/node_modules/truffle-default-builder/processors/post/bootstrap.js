var fs = require("fs");

module.exports = function(contents, file, options, process, callback) {

  var contract_names = options.contracts.map(function(contract) {return contract.contract_name;});

  // Note: __contracts__ is provided by frontend_dependencies.js

  contents = "\
//// TRUFFLE BOOTSTRAP                                          \n\n\
Object.keys(__contracts__).forEach(function(contract_name) {    \n\n\
  window[contract_name] = __contracts__[contract_name];         \n\n\
});                                                             \n\n\
window.addEventListener('load', function() {                    \n\n\
                                                                \n\n\
  // Supports Mist, and other wallets that provide 'web3'.      \n\n\
  if (typeof web3 !== 'undefined') {                            \n\n\
    // Use the Mist/wallet provider.                            \n\n\
    window.web3 = new Web3(web3.currentProvider);               \n\n\
  } else {                                                      \n\n\
    // Use the provider from the config.                        \n\n\
    window.web3 = new Web3(new Web3.providers.HttpProvider('http://" + options.rpc.host + ":" + options.rpc.port + "')); \n\n\
  }                                                             \n\n\
                                                                \n\n\
  [" + contract_names + "].forEach(function(contract) {         \n\n\
    contract.setProvider(window.web3.currentProvider);          \n\n\
  });                                                           \n\n\
});                                                              \n\n\
//// END TRUFFLE BOOTSTRAP                                      \n\n\ " + contents;

  callback(null, contents);
};
