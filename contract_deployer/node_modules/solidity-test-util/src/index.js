const solidityTestUtil = {

  prepareValue: value => value && value.toNumber != undefined ? value.toNumber().toString() : value,

  prepareArray: arr => {
    if (!Array.isArray(arr)) throw new Error('Argument for function "prepareArray" must be array, ' + typeof(arr) + ' given');
    return arr.map(value => solidityTestUtil.prepareValue(value));
  },

  prepareObject: obj =>
    Object.keys(obj).reduce(
      (preparedObj, key) => Object.assign(preparedObj, {
        [key]: solidityTestUtil.prepareValue(obj[key])
      }), {}
    ),

  getEventLog: eventListener =>
    new Promise(
      (resolve, reject) => eventListener.get(
        (error, log) => error ? reject(error) : resolve(log)
      )),


  assertEvent: (result, eventName, args) =>
    assert.deepEqual(
      solidityTestUtil.prepareObject(result.logs.find(log => log.event == eventName).args),
      args
    ),

  assertJump: (error, message = '') => {
    assert.isAbove(error.message.search('invalid'), -1, message + ': error must be returned');
  },

  assertThrow: async (callback, message = '') => {
    var error;
    try {
      await callback();
    } catch (err) {
      error = err;
    }

    if (error) solidityTestUtil.assertJump(error, message);
    else  assert.notEqual(error, undefined, 'Error need to be thrown: ' + message);
  },


  evmIncreaseTime: (seconds) => new Promise((resolve, reject) =>
    web3.currentProvider.sendAsync({
      jsonrpc: "2.0",
      method: "evm_increaseTime",
      params: [seconds],
      id: new Date().getTime()
    }, (error, result) => error ? reject(error) : resolve(result.result))),


  EMPTY_ADDRESS: '0x0000000000000000000000000000000000000000'
};

module.exports = solidityTestUtil;