/**
 * babel-preset-power-assert
 *   Babel preset for all power-assert plugins
 * 
 * https://github.com/twada/babel-preset-power-assert
 *
 * Copyright (c) 2016 Takuto Wada
 * Licensed under the MIT license.
 *   http://twada.mit-license.org/
 */
'use strict';

module.exports = {
    plugins: [
        require('babel-plugin-empower-assert'),
        require('babel-plugin-espower')
    ]
};
