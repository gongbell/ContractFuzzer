[![power-assert][power-assert-banner]][power-assert-url]

[![NPM version][npm-image]][npm-url]
[![License][license-image]][license-url]


`babel-preset-power-assert` is a [Babel](https://babeljs.io/) preset for all [power-assert](https://github.com/power-assert-js/power-assert) plugins.


INSTALL
---------------------------------------

```
$ npm install --save-dev babel-preset-power-assert
```


HOW TO USE
---------------------------------------

### via [.babelrc](http://babeljs.io/docs/usage/babelrc/) (Recommended)

```javascript
{
  "presets": [
    "babel-preset-power-assert"
  ]
}
```

### via [Babel CLI](http://babeljs.io/docs/usage/cli/)

```
$ babel --presets babel-preset-power-assert /path/to/src/target.js > /path/to/build/target.js
```

### via [Babel API](http://babeljs.io/docs/usage/api/)

```javascript
var babel = require('babel-core');
var jsCode = fs.readFileSync('/path/to/src/target.js');
var transformed = babel.transform(jsCode, {
    presets: ['babel-preset-power-assert']
});
console.log(transformed.code);
```


AUTHOR
---------------------------------------
* [Takuto Wada](https://github.com/twada)


LICENSE
---------------------------------------
Licensed under the [MIT](http://twada.mit-license.org/) license.


[power-assert-url]: https://github.com/power-assert-js/power-assert
[power-assert-banner]: https://raw.githubusercontent.com/power-assert-js/power-assert-js-logo/master/banner/banner-official-fullcolor.png

[npm-url]: https://npmjs.org/package/babel-preset-power-assert
[npm-image]: https://badge.fury.io/js/babel-preset-power-assert.svg

[license-url]: http://twada.mit-license.org/
[license-image]: https://img.shields.io/badge/license-MIT-brightgreen.svg
