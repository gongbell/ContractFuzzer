# unix-timestamp

Tiny library to create and manipulate Unix timestamps in Javascript. (A Unix timestamp is the number of seconds elapsed since Unix epoch time, i.e. January 1 1970 00:00 UTC.)

[![NPM version](https://badge.fury.io/js/unix-timestamp.png)](http://badge.fury.io/js/unix-timestamp)

## Usage

Install with `npm install unix-timestamp`, then:

- `timestamp.now([delta])` gives the current time, optionally applying a delta (see below)
- `timestamp.fromDate(dateOrString)` gives the time from a Javascript Date object or an ISO 8601 date string
- `timestamp.toDate(time)` correspondingly gives the date from a timestamp
- `timestamp.add(time, delta)` applies a delta to the given time
- `timestamp.duration(delta)` gives the delta timestamp for the given delta string

A delta can be either a number (unit: seconds) or a string with format `[+|-] [{years}y] [{months}M] [{weeks}w] [{days}d] [{hours}h] [{minutes}m] [{seconds}s] [{milliseconds}ms]` (for example `-30s`). The actual values (in seconds) used for each unit of time are accessible in properties `Millisecond`, `Second`, `Minute`, `Hour`, `Day`, `Week`, `Month` (i.e. mean gregorian month) and `Year`.

By default timestamps include decimals (fractions of a second). You can set the lib to round all returned timestamps to the second with `timestamp.round = true`.


## Tests

Install dev dependencies with `npm install`, then `npm test`.

## License

[Revised BSD license](https://github.com/pryv/documents/blob/master/license-bsd-revised.md)
