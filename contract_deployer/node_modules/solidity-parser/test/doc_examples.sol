// Examples taken from the Solidity documentation online.

// for pragma version numbers, see https://docs.npmjs.com/misc/semver#versions
pragma solidity 0.4.0;
pragma solidity v0.4.0; // like npm
pragma solidity ^0.4.0;
pragma solidity >= 0.4.0;
pragma solidity <= 0.4.0;
pragma solidity < 0.4.0;
pragma solidity > 0.4.0;
pragma solidity != 0.4.0;
pragma solidity >=0.4.0 <0.4.8; // from https://github.com/ethereum/solidity/releases/tag/v0.4.0

import "SomeFile.sol";
import "SomeFile.sol" as SomeOtherFile;
import * as SomeSymbol from "AnotherFile.sol";
import {symbol1 as alias, symbol2} from "File.sol";

contract c {
  function c()
  {
      val1 = 1 wei;    // 1
      val2 = 1 szabo;  // 1 * 10 ** 12
      val3 = 1 finney; // 1 * 10 ** 15
      val4 = 1 ether;  // 1 * 10 ** 18
 }
  uint256 val1;
  uint256 val2;
  uint256 val3;
  uint256 val4;
}

contract test {
    enum ActionChoices { GoLeft, GoRight, GoStraight, SitStill };
    function test()
    {
        choices = ActionChoices.GoStraight;
    }
    function getChoice() returns (uint d)
    {
        d = uint256(choices);
    }
    ActionChoices choices;
}

contract Base {
    function Base(uint i)
    {
        m_i = i;
    }
    uint public m_i;
}
contract Derived is Base(0) {
    function Derived(uint i) Base(i) {}
}

contract C {
  uint248 x; // 31 bytes: slot 0, offset 0
  uint16 y; // 2 bytes: slot 1, offset 0 (does not fit in slot 0)
  uint240 z; // 30 bytes: slot 1, offset 2 bytes
  uint8 a; // 1 byte: slot 2, offset 0 bytes
  struct S {
    uint8 a; // 1 byte, slot +0, offset 0 bytes
    uint256 b; // 32 bytes, slot +1, offset 0 bytes (does not fit)
  }
  S structData; // 2 slots, slot 3, offset 0 bytes (does not really apply)
  uint8 alpha; // 1 byte, slot 4 (start new slot after struct)
  uint16[3] beta; // 3*16 bytes, slots 5+6 (start new slot for array)
  uint8 gamma; // 1 byte, slot 7 (start new slot after array)
}

contract test {
  function f(uint x, uint y) returns (uint z) {
    var c = x + 3;
    var b = 7 + (c * (8 - 7)) - x;
    return -(-b | 0);
  }
}

contract test {
  function f(uint x, uint y) returns (uint z) {
    return 10;
  }
}

contract c {
  function () returns (uint) { return g(8); }
  function g(uint pos) internal returns (uint) { setData(pos, 8); return getData(pos); }
  function setData(uint pos, uint value) internal { data[pos] = value; }
  function getData(uint pos) internal { return data[pos]; }
  mapping(uint => uint) data;
}

contract Sharer {
    function sendHalf(address addr) returns (uint balance) {
        if (!addr.send(msg.value/2))
            throw; // also reverts the transfer to Sharer
        return address(this).balance;
    }
}

/// @dev Models a modifiable and iterable set of uint values.
library IntegerSet
{
  struct data
  {
    /// Mapping item => index (or zero if not present)
    mapping(uint => uint) index;
    /// Items by index (index 0 is invalid), items with index[item] == 0 are invalid.
    uint[] items;
    /// Number of stored items.
    uint size;
  }
  function insert(data storage self, uint value) returns (bool alreadyPresent)
  {
    uint index = self.index[value];
    if (index > 0)
      return true;
    else
    {
      if (self.items.length == 0) self.items.length = 1;
      index = self.items.length++;
      self.items[index] = value;
      self.index[value] = index;
      self.size++;
      return false;
    }
  }
  function remove(data storage self, uint value) returns (bool success)
  {
    uint index = self.index[value];
    if (index == 0)
      return false;
    delete self.index[value];
    delete self.items[index];
    self.size --;
  }
  function contains(data storage self, uint value) returns (bool)
  {
    return self.index[value] > 0;
  }
  function iterate_start(data storage self) returns (uint index)
  {
    return iterate_advance(self, 0);
  }
  function iterate_valid(data storage self, uint index) returns (bool)
  {
    return index < self.items.length;
  }
  function iterate_advance(data storage self, uint index) returns (uint r_index)
  {
    index++;
    while (iterate_valid(self, index) && self.index[self.items[index]] == index)
      index++;
    return index;
  }
  function iterate_get(data storage self, uint index) returns (uint value)
  {
      return self.items[index];
  }
}

/// How to use it:
contract User
{
  /// Just a struct holding our data.
  IntegerSet.data data;
  /// Insert something
  function insert(uint v) returns (uint size)
  {
    /// Sends `data` via reference, so IntegerSet can modify it.
    IntegerSet.insert(data, v);
    /// We can access members of the struct - but we should take care not to mess with them.
    return data.size;
  }
  /// Computes the sum of all stored data.
  function sum() returns (uint s)
  {
    for (var i = IntegerSet.iterate_start(data); IntegerSet.iterate_valid(data, i); i = IntegerSet.iterate_advance(data, i))
      s += IntegerSet.iterate_get(data, i);
  }
}

// This broke it at one point (namely the modifiers).
contract DualIndex {
  mapping(uint => mapping(uint => uint)) data;
  address public admin;

  modifier restricted { if (msg.sender == admin) _ }

  function DualIndex() {
    admin = msg.sender;
  }

  function set(uint key1, uint key2, uint value) restricted {
    uint[2][4] memory defaults; // "memory" broke things at one time.
    data[key1][key2] = value;
  }

  function transfer_ownership(address _admin) restricted {
    admin = _admin;
  }

  function lookup(uint key1, uint key2) returns(uint) {
    return data[key1][key2];
  }
}

contract A {

}

contract B {

}

contract C is A, B {

}

contract TestPrivate
{
  uint private value;
}

contract TestInternal
{
  uint internal value;
}

contract FromSolparse is A, B, TestPrivate, TestInternal {
  function() {
    uint a = 6 ** 9;
    var (x) = 100;
    uint y = 2 days
  }
}

contract CommentedOutFunction {
  // FYI: This empty function, as well as the commented
  // out function below (bad code) is important to this test.
  function() {

  }

  // function something()
  //  uint x = 10;
  // }
}

library UsingExampleLibrary {
  function sum(uint[] storage self) returns (uint s) {
    for (uint i = 0; i < self.length; i++)
      s += self[i];
  }
}

contract UsingExampleContract {
  using UsingExampleLibrary for uint[];
}

contract NewStuff {
  function someFunction() payable {
    string a = hex"ab1248fe";
  }
}
