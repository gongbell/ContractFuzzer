var async = require("async");
var fs = require("fs");
var path = require("path");
var colors = require("colors");
var Promise = require("bluebird");
var mkdirp = Promise.promisify(require("mkdirp"));
var copy = require("./lib/copy");

function DefaultBuilder(config, build_or_dist, extra_processors) {
  this.config = config;
  this.key = build_or_dist;
  this.processors = {
    // These processors do nothing, but are registered to reduce warnings.
    ".html": "./processors/null.js",
    ".js": "./processors/null.js",
    ".css": "./processors/null.js",
    ".json": "./processors/null.js",

    // Helpful default processors.
    ".coffee": "./processors/coffee.js",
    ".scss": "./processors/scss.js",

    // Babel-related processors.
    ".es6": "./processors/babel.js",
    ".es": "./processors/babel.js",
    ".jsx": "./processors/babel.js",

    // Named processors for post-processing.
    "null": "./processors/null.js", // does nothing; useful in some edge cases.
    "uglify": "./processors/post/uglify.js",
    "frontend-dependencies": "./processors/post/frontend_dependencies.js",
    "bootstrap": "./processors/post/bootstrap.js",
    "include-contracts": "./processors/post/include_contracts.js"
  };
  this.contexts = {
    build: {
      defaults: {
        "post-process": {
          "app.js": [
            "bootstrap",
            "include-contracts",
            "frontend-dependencies"
          ]
        }
      }
    },
    dist: {
      defaults: {
        "post-process": {
          "app.js": [
            "bootstrap",
            "include-contracts",
            "frontend-dependencies",
            "uglify"
          ]
        }
      }
    }
  };
  this.context = this.contexts[this.key];

  this.setup(extra_processors);
};

DefaultBuilder.prototype.setup = function(extra_processors) {
  extra_processors = extra_processors || {};

  // Evaluate build targets, making the configuration conform, adding
  // default post processing, if any.
  var targets = Object.keys(this.config);
  for (var i = 0; i < targets.length; i++) {
    var target = targets[i];
    var options = this.config[target];
    if (typeof options == "string") options = [options];
    if (options instanceof Array) {
      options = {
        files: options,
        "post-process": {
          build: [],
          dist: []
        }
      }
    }

    if (options["post-process"] == null) {
      options["post-process"] = {build: [], dist: []};
    }

    // If an array was passed, use the same post processing in both contexts.
    if (options["post-process"] instanceof Array) {
      var new_post_process = {
        build: options["post-process"],
        dist: options["post-process"]
      }
      options["post-process"] = new_post_process;
    }

    // Check for default post processing for this target,
    // and add it if the target hasn't specified any post processing.
    if (this.context.defaults["post-process"][target] != null && options["post-process"][this.key].length == 0) {
      options["post-process"][this.key] = this.context.defaults["post-process"][target];
    }

    this.config[target] = options;
  }

  // Use full paths for default processors.
  var extensions = Object.keys(this.processors);
  for (var i = 0; i < extensions.length; i++) {
    var extension = extensions[i];
    var full_path = path.join(__dirname, this.processors[extension]);
    this.processors[extension] = full_path;
  }

  // Add extra processors to the processors list.
  var extensions = Object.keys(extra_processors);
  for (var i = 0; i < extensions.length; i++) {
    var extension = extensions[i];
    var full_path = extra_processors[extension];
    extension = extension.toLowerCase();
    this.processors[extension] = full_path;
  }
};

DefaultBuilder.prototype.build = function(options, callback) {
  this.working_directory = options.working_directory;
  this.destination_directory = options.destination_directory;
  this.source_directory = path.join(options.working_directory, "app");

  this.options = options;

  this.process_all_targets(callback);
};

DefaultBuilder.prototype.process_file = function(file, callback) {
  var self = this;
  var extension = path.extname(file).toLowerCase();
  var processor_path = this.processors[extension];
  this.expect(processor_path, "specified \"" + extension + "\" processor", "Check your app config.");
  var processor = require(processor_path);

  if (processor == null) {
    var display_name = "." + file.replace(this.working_directory, "");
    console.log(colors.yellow("Warning: Couldn't find processor for " + display_name + ". Including as is."));
    processor = this.processors["null"];
  }

  fs.readFile(file, {encoding: "utf8"}, function(err, contents) {
    if (err != null) {
      callback(err);
      return;
    }

    processor(contents, file, self.options, self.process_files.bind(self), callback);
  });
};

DefaultBuilder.prototype.process_files = function(files, base_path, separator, callback) {
  var self = this;

  if (typeof base_path == "function") {
    separator = base_path;
    base_path = null;
  }

  if (typeof separator == "function") {
    callback = separator;
    separator = "\n\n";
  }

  if (typeof files == "string") files = [files];

  async.reduce(files, "", function(memo, file, iterator_callback) {
    var full_path = file;
    if (base_path != null) {
      full_path = path.join(base_path, file);
    }

    if (!self.expect(full_path, iterator_callback)) {
      return;
    }

    self.process_file(full_path, function(err, processed) {
      if (err != null) {
        console.log("");
        console.log(colors.red("Error in " + file));
        console.log("");
        iterator_callback(err);
        return;
      }

      iterator_callback(null, memo + separator + processed);
    });
  }, callback);
};

DefaultBuilder.prototype.process_directory = function(target, callback) {
  var value = this.config[target];
  var destination_directory = path.join(this.destination_directory, target);
  var source_directory = path.join(this.source_directory, value.files[0]);

  if (!this.expect(source_directory, "source directory for target " + target, "Check app configuration.", callback)) {
    return;
  }

  mkdirp(destination_directory).then(function() {
    copy(source_directory, destination_directory, callback);
  }).catch(callback)
};

DefaultBuilder.prototype.process_target = function(target, callback) {
  var self = this;

  // Is this a directory?
  if (target[target.length - 1] == "/") {
    this.process_directory(target, callback);
    return;
  }

  var files = this.config[target].files;
  var post_processing = this.config[target]["post-process"][this.key];
  var target_file = path.join(this.destination_directory, target);

  this.process_files(files, this.source_directory, function(err, processed) {
    if (err != null) {
      callback(err);
      return;
    }

    // Now do post processing.
    async.reduce(post_processing, processed, function(memo, processor_name, post_processor_finished) {
      var post_processor_path = self.processors[processor_name];
      self.expect(post_processor_path, "specified post processor \"" + processor_name + "\"", "Check your app config.");
      var post_processor = require(post_processor_path);

      if (!post_processor) {
        post_processor_finished(new Error("Cannot find processor named '" + processor_name + "' during post-processing. Check app configuration."));
        return;
      }

      if (typeof post_processor != "function") {
        post_processor_finished(new Error("Couldn't load custom processor '" + processor_name + "'; processor function not correctly exported."));
        return;
      }

      post_processor(memo, target_file, self.options, self.process_files.bind(self), post_processor_finished);
    }, function(err, final_post_processed) {
      if (err != null) {
        callback(err);
        return;
      }

      mkdirp(path.dirname(target_file)).then(function() {
        fs.writeFile(target_file, final_post_processed, {encoding: 'utf8'}, callback);
      }).catch(callback);
    });
  });
};

DefaultBuilder.prototype.process_all_targets = function(callback) {
  var self = this;
  async.eachSeries(Object.keys(this.config), function(target, finished_with_target) {
    self.process_target(target, finished_with_target);
  }, callback);
};

DefaultBuilder.prototype.expect = function(expected_path, description, extra, callback) {
  if (typeof description == "function") {
    callback = description;
    description = "file";
    extra = "";
  }

  if (typeof extra == "function") {
    callback = description;
    extra = "";
  }

  if (!fs.existsSync(expected_path)) {
    var display_path = expected_path.replace(this.working_directory, "./");
    var error = "Couldn't find " + description + " at " + display_path + ". " + extra;

    if (callback != null) {
      callback(error);
      return false;
    } else {
      throw error;
    }
  }
  return true;
};

module.exports = DefaultBuilder;
