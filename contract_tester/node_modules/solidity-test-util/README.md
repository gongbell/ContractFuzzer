# Solidity-test-util

A collection of utility functions for testing ethereum contracts with [truffle framework](https://github.com/Consensys/truffle) and [testrpc](https://github.com/ethereumjs/testrpc).

Examples can be found on [Github solidity-test-example project](https://github.com/vitiko/solidity-test-example/blob/master/test/CongressWithTestUtil.js) 

## getEventLog
 
**Parameters**

-   `Object` Contract events subscriber,  [see web3 doc](https://github.com/ethereum/wiki/wiki/JavaScript-API#contract-events)

**Returns**
 
 - `Array`  Events
 
 
## assertThrow

**Parameters**

-   `Callback` with  contract method call, that should throw exception

**Returns**
 
 - `Boolean` exception throwed


## evmIncreaseTime

Jump testrpc forward in time 

**Parameters**

-   `Number` amount of time to increase in seconds

**Returns**
 
 - `Number`  total time adjustment, in seconds.
 
 
 
 ## prepareValue 
   
 Convert `BigNumber` value to `Number` using toNumber() method
   
 **Parameters**
   
 -  `Mixed` 
   
 **Returns**
   
 - `Mixed` 
  
 ## prepareArray 
 
  Convert `BigNumber` array values to `Number` using toNumber() method
    
  **Parameters**
    
  - `Array` 
    
  **Returns**
    
  - `Array` 
   
 ## prepareObject
 
  Convert `BigNumber` object values to `Number` using toNumber() method
    
  **Parameters**
    
  -   `Object` 
    
  **Returns**
    
  - `Object` 