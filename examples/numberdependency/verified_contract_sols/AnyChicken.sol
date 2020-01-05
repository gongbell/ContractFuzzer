pragma solidity ^0.4.16;

contract AnyChicken {

    address public owner;
	address public bigChicken;
	uint public bigAmount;
	uint public lastBlock;
	
	function AnyChicken() public payable {
		owner = msg.sender;
		bigChicken = msg.sender;
		bigAmount = msg.value;
		lastBlock = block.number;
	}
	
	function () public payable {
		if (block.number <= lastBlock + 1000) {
			require(msg.value > bigAmount);
			bigChicken = msg.sender;
			bigAmount = msg.value;
			lastBlock = block.number;
			owner.transfer(msg.value/100);
		}
		else {
			require(msg.sender == bigChicken);
			bigChicken.transfer(this.balance);
		}
	}
}