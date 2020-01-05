pragma solidity ^0.4.17;
/*
	REA Lottery Wheel Contract

	The constructor sets last_hash to some initial value.
	Every call to spin() will increase the round_count by one and
	a put a new "random" hash into the storage map "hashes".
	spin() accepts an argument which can be used to introdue more "randomness".

	The community can participate by sending small amounts of Eth (no matter the value)
	to the smart contract. The value sent together with timestamp and blocknumber increase
	the "randomness".

	The outcome of round <n> can be retrived via call to get_hash(<n>).  

	WARNING and DISCLAIMER: 
	We fully understand the fact that Ethereum Smart Contracts
	by design of Ethereum Blockchain and Solidity language work
	in a determenistic and predictable way. 

	The block number and the timestamp are not random variables in
	a mathematical sense. Even worse, the interested miners can 
	affect the outcome by not including the contract transaction
	in a current block if they are not happy about the outcome 
	(since miners in theory know the outcome of every contract transaction
	before the transaction is included in a block). 

	2017 Pavel Metelitsyn

*/

contract REALotteryWheel{
    
    uint16 public round_count = 0;
    bytes32 public last_hash;
    address public controller;
    
    mapping (uint16 => bytes32) public hashes;
    
    function REALotteryWheel() public {
        controller = msg.sender;
        last_hash = keccak256(block.number, now);    
    }
    
    function spin(bytes32 s) public {
        if(controller != msg.sender) revert();
        round_count = round_count + 1;
        last_hash = keccak256(block.number,now,s);
        hashes[round_count] = last_hash;
        
    }
    
    function get_hash (uint16 i) constant returns (bytes32){
        return hashes[i];
    }
    
    function () payable {
        spin(bytes32(msg.value));
    }
    
}