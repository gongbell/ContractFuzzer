/**
 * Tiny library to create and manipulate Unix timestamps
 * (i.e. defined as the number of seconds since Unix epoch time).
 */

var timestamp = module.exports = {};

// constants

timestamp.Millisecond = 0.001;
timestamp.Second = 1;
timestamp.Minute = 60;
timestamp.Hour = 60 * timestamp.Minute;
timestamp.Day = 24 * timestamp.Hour;
timestamp.Week = 7 * timestamp.Day;
/**
 * = mean Gregorian month
 */
timestamp.Month = 30.436875 * timestamp.Day;
timestamp.Year = 12 * timestamp.Month;

var DeltaRegExp = new RegExp('^\\s*' +
    '([-+]?)\\s*' +
    [ 'y', 'M', 'w', 'd', 'h', 'm', 's', 'ms' ]
        .map(function (t) { return '(?:(\\d+)\\s*' + t + ')?'; })
        .join('\\s*') +
    '\\s*$');

var outputFn = dontRound;
function dontRound(time) { return time; }
function round(time) { return Math.round(time); }
/**
 * Set to `true` to round all returned timestamps to the second. Defaults to `false`.
 */
Object.defineProperty(timestamp, 'round', {
  get: function () { return outputFn; },
  set: function (value) { outputFn = value ? round : dontRound; }
});


/**
 * Gets the current time as Unix timestamp.
 * Optionally applying a given delta specified as either a human-readable string or a number of
 * seconds.
 *
 * @param {String|Number} delta The optional delta time to apply
 * @returns {Number} The corresponding timestamp
 */
timestamp.now = function (delta) {
  var now = Date.now() / 1000;
  return outputFn(delta ? timestamp.add(now, delta) : now);
};

/**
 * Applies the given delta to the given timestamp.
 * The delta is specified as either a human-readable string or a number of
 * seconds.
 *
 * @param {Number} time The original timestamp
 * @param {String|Number} delta The delta time to apply
 * @returns {Number} The result timestamp
 */
timestamp.add = function (time, delta) {
  if (! isNumber(time)) {
    throw new Error('Time must be a number');
  }
  if (isString(delta)) {
    var matches = DeltaRegExp.exec(delta);
    if (! matches) {
      throw new Error('Expected delta string format: [+|-] [{years}y] [{months}M] [{weeks}w] ' +
          '[{days}d] [{hours}h] [{minutes}m] [{seconds}s] [{milliseconds}ms]');
    }
    delta = (matches[1] === '-' ? -1 : 1) * (
        (matches[2] || 0) * timestamp.Year +
        (matches[3] || 0) * timestamp.Month +
        (matches[4] || 0) * timestamp.Week +
        (matches[5] || 0) * timestamp.Day +
        (matches[6] || 0) * timestamp.Hour +
        (matches[7] || 0) * timestamp.Minute +
        (matches[8] || 0) * timestamp.Second +
        (matches[9] || 0) * timestamp.Millisecond
    );
  } else if (! isNumber(delta)) {
    throw new Error('Delta must be either a string or a number');
  }
  return outputFn(time + delta);
};

/**
 * Gets the delta timestamp for the given delta string.
 * (Alias for .add() using a time of zero.)
 *
 * @param {String|Number} delta The delta time for the duration
 * @returns {Number} The result time delta
 */
timestamp.duration = function (delta) {
  return timestamp.add(0, delta);
};

/**
 * Gets the Unix timestamp for the given date object or string.
 *
 * @param {Date|String} date A date object or an ISO 8601 date string
 * @returns {Number} The corresponding timestamp
 */
timestamp.fromDate = function (date) {
  if (isString(date)) {
    date = new Date(date);
  } else if (! isDate(date)) {
    throw new Error('Expected either a string or a date');
  }
  return outputFn(date.getTime() / 1000);
};

/**
 * Gets the date for the given Unix timestamp.
 *
 * @param {Number} time A timestamp
 * @returns {Date} The corresponding date
 */
timestamp.toDate = function (time) {
  if (! isNumber(time)) {
    throw new Error('Expected a number');
  }
  return new Date(time * 1000);
};

function isString(value) {
  return typeof value === 'string' || Object.prototype.toString.call(value) === '[object String]';
}

function isNumber(value) {
  return typeof value === 'number' || Object.prototype.toString.call(value) === '[object Number]';
}

function isDate(value) {
  return Object.prototype.toString.call(value) === '[object Date]';
}
