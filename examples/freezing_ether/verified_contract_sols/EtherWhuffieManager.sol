pragma solidity ^0.4.0;
contract EtherWhuffieManager {
    struct Status {
        uint positiveWhuffies;
        uint negativeWhuffies;
    }
    
    mapping(address => Status) public stats;
    
    event PositiveWhuffiesSent(address indexed _from, address indexed _to, uint whuffies, string message);
    event NegativeWhuffiesSent(address indexed _from, address indexed _to, uint whuffies, string message);
    
    ///////////////////////////////////////////////////////////
    
    function sendPositiveWhuffies(address _to, string message) public payable {
            stats[_to].positiveWhuffies += msg.value;            
            PositiveWhuffiesSent(msg.sender, _to, msg.value, message);
    }
    
    function sendNegativeWhuffies(address _to, string message) public payable {
            stats[_to].negativeWhuffies += msg.value;            
            NegativeWhuffiesSent(msg.sender, _to, msg.value, message);
    }
}