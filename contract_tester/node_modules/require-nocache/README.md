# require-nocache

A module for non-caching `require()` calls. Useful for avoiding server
restarts when developing; probably shouldn't ever be used in production.

## Installation

    npm install require-nocache

### Example

example.js

    var reload = require('require-nocache')(module)
    
    setInterval(
      function(){
        console.log(reload('./data'))
      },
      100
    )
    
data.js

    module.exports = Math.random()

Output

    0.5634244652464986
    0.551380671793595
    0.12054721242748201
    0.44393923226743937
