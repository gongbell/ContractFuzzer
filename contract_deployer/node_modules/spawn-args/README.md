spawn-args
==========

![Build status](https://api.travis-ci.org/binocarlos/spawn-args.png)

Turn a string of command line options into an array for child_process.spawn

## install

```
$ npm install spawn-args
```

## usage

```js
var spawnargs = require('spawn-args');

var args = spawnargs('-port 80 --title "this is a title"');

/*

	[
		'-port',
		'80',
		'--title',
		'"this is a title"'
	]
	
*/
```

## license

MIT