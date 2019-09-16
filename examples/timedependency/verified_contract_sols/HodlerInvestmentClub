/**
 *Submitted for verification at Etherscan.io on 2017-12-25
*/

pragma solidity ^0.4.18;
/**
 * Holds funds for a year.  Send to or deposit directly to this contract.
 * Each new acccount is initialized with a 1 year hold period, and is only 
 * retrievable from the designated address after the set hold time.
*/
contract HodlerInvestmentClub {
    uint public hodl_interval= 1 years;
    uint public m_hodlers = 1;
    
    struct Hodler {
        uint value;
        uint time;
    }
    
    mapping(address => Hodler) public hodlers;
    
    modifier onlyHodler {
        require(hodlers[msg.sender].value > 0);
        _;
    }
    
    /* Constructor */
    function HodlerInvestmentClub() payable public {
        if (msg.value > 0)  {
            hodlers[msg.sender].value = msg.value;
            hodlers[msg.sender].time = now + hodl_interval;
        }
    }
    
    // join the club!
    // make a deposit to another account if it exists 
    // or initialize a deposit for a new account
    function deposit(address _to) payable public {
        require(msg.value > 0);
        if (_to == 0) _to = msg.sender;
        // if a new member, init a hodl time
        if (hodlers[_to].time == 0) {
            hodlers[_to].time = now + hodl_interval;
            m_hodlers++;
        } 
        hodlers[_to].value += msg.value;
    }
    
    // withdrawal can only occur after deposit time is exceeded
    function withdraw() public onlyHodler {
        require(hodlers[msg.sender].time <= now);
        uint256 value = hodlers[msg.sender].value;
        delete hodlers[msg.sender];
        m_hodlers--;
        require(msg.sender.send(value));
    }
    
    // join the club!
    // simple deposit and hold time set for msg.sender
    function() payable public {
        require(msg.value > 0);
        hodlers[msg.sender].value += msg.value;
        // init for first deposit
        if (hodlers[msg.sender].time == 0) {
            hodlers[msg.sender].time = now + hodl_interval;
            m_hodlers++;
        }
    }

}
