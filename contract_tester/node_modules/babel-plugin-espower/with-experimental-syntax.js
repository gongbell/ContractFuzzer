/**
 * Bridge module for babel-preset-power-assert to enable `embedAst` option by default.
 * 
 * NOTE: this is an internal & interim module and will be removed from next major version,
 *   since `embedAst` will be true by default in next major.
 */
'use strict';

var createEspowerPlugin = require('./create');

module.exports = function (babel) {
    return createEspowerPlugin(babel, { embedAst: true });
};
