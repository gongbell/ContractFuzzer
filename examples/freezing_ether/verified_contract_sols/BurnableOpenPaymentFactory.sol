//A BurnableOpenPaymet is instantiated with a specified payer and a commitThreshold.
//The recipient not set when the contract is instantiated.

//All behavior of the contract is directed by the payer, but
//the payer can never directly recover the payment unless he becomes the recipient.

//Anyone can become the recipient by contributing the commitThreshold.
//The recipient cannot change once it's been set.

//The payer can at any time choose to burn or release to the recipient any amount of funds.

pragma solidity ^0.4.1;

contract BurnableOpenPayment {
    address public payer;
    address public recipient;
    address public burnAddress = 0xdead;
    uint public commitThreshold;
    
    modifier onlyPayer() {
        if (msg.sender != payer) throw;
        _;
    }
    
    modifier onlyWithRecipient() {
        if (recipient == address(0x0)) throw;
        _;
    }
    
    //Only allowing the payer to fund the contract ensures that the contract's
    //balance is at most as difficult to predict or interpret as the payer.
    //If the payer is another smart contract or script-based, balance changes
    //could reliably indicate certain intentions, judgements or states of the payer.
    
    
    function BurnableOpenPayment(address _payer, uint _commitThreshold) payable {
        payer = _payer;
        commitThreshold = _commitThreshold;
    }
    
    
    
    
    
    
    
    
    
    
    
    
}

contract BurnableOpenPaymentFactory {
    function newBurnableOpenPayment(address payer, uint commitThreshold) payable returns (address) {
        //pass along any ether to the constructor
        return (new BurnableOpenPayment).value(msg.value)(payer, commitThreshold);
    }
}