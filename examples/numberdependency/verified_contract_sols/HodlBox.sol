pragma solidity ^0.4.10;
/*
      No more panic sells.
      Force yourself to hodl them eths with HodlBox!
*/

contract HodlBox {







  uint public hodlTillBlock;
  address public hodler;
  uint public hodling;
  bool public withdrawn;

  event HodlReleased(bool _isReleased);
  event Hodling(bool _isCreated);

  function HodlBox(uint _blocks) payable {
    hodler = msg.sender;
    hodling = msg.value;
    hodlTillBlock = block.number + _blocks;
    withdrawn = false;
    Hodling(true);
  }

  function deposit() payable {


    
hodling += msg.value;
  }

  function releaseTheHodl() {

    // Only the contract creator can release funds from their HodlBox,
    // and only after the defined number of blocks has passed.

    
if (msg.sender != hodler) {
throw;}else { }


    
if (block.number < hodlTillBlock) {
throw;}else { }


    
if (withdrawn) {
throw;}else { }


    
if (hodling <= 0) {
throw;}else { }


    
withdrawn = true;

    
hodling = 0;

    // Send event to notifiy UI

    
HodlReleased(true);


    
selfdestruct(hodler);
  }

  // constant functions do not mutate state
  function hodlCountdown()  returns (uint) {


    
var hodlCount = hodlTillBlock - block.number;

    
if (block.number >= hodlTillBlock) {

      
return 0;
    }else { }


    
return hodlCount;
  }

  function isDeholdable()  returns (bool) {


    
if (block.number < hodlTillBlock) {

      
return false;
    } else {

      
return true;
    }
  }

}