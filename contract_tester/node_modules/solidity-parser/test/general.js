var SolidityParser = require('../index.js');

describe("Parser", function() {
  it("parses documentation examples without throwing an error", function(done) {
    var result = SolidityParser.parseFile("./test/doc_examples.sol", true);
    //console.log(JSON.stringify(result.body.filter(function(i) {return i.type == "ImportStatement"}), null, 2));
    done();
  });

  it("parses documentation examples using imports parser without throwing an error", function(done) {
    var result = SolidityParser.parseFile("./test/doc_examples.sol", "imports", true);
    //console.log(JSON.stringify(result, null, 2));
    done();
  });
});

describe("Built Parser", function() {
  it("parses documentation examples without throwing an error", function(done) {
    var result = SolidityParser.parseFile("./test/doc_examples.sol", false);
    //console.log(JSON.stringify(result.body.filter(function(i) {return i.type == "ImportStatement"}), null, 2));
    done();
  });

  it("parses documentation examples using imports parser without throwing an error", function(done) {
    var result = SolidityParser.parseFile("./test/doc_examples.sol", "imports", false);
    //console.log(JSON.stringify(result, null, 2));
    done();
  });
});
