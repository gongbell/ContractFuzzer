pragma solidity ^0.4.6;

// ## Matthew - a contract for increasing "whaleth"
// README: https://github.com/rolandkofler/matthew
// MIT LICENSE 2016 Roland Kofler, thanks to Crul for testing

contract Matthew {
    address owner;
    address whale;
    uint256 blockheight;
    uint256 period = 18; //180 blocks ~ 42 min, 300 blocks ~ 1h 10 min;
    uint constant DELTA = 0.1 ether;
    uint constant WINNERTAX_PRECENT = 10;
    bool mustBeDestroyed = false;
    uint newPeriod = 5;
    
    event MatthewWon(string msg, address winner, uint value,  uint blocknumber);
    event StakeIncreased(string msg, address staker, uint value, uint blocknumber);
    
    function Matthew(){
        owner = msg.sender;
        setFacts();
    }
    
    function setFacts() private {
        period = newPeriod;
        blockheight = block.number;
        whale = msg.sender;
    }
    
    /// The rich get richer, the whale get whaler
    function () payable{
    
        if (block.number - period >= blockheight){ // time is over, Matthew won
            bool isSuccess=false; //mutex against recursion attack
            var nextStake = this.balance * WINNERTAX_PRECENT/100;  // leave some money for the next round
            if (isSuccess == false) //check against recursion attack
                isSuccess = whale.send(this.balance - nextStake); // pay out the stake
            MatthewWon("Matthew won", whale, this.balance, block.number);
            setFacts();//reset the game
            if (mustBeDestroyed) selfdestruct(whale); 
            return;
            
        }else{ // top the stake
            if (msg.value < this.balance + DELTA) throw; // you must rise the stake by Delta
            bool isOtherSuccess = msg.sender.send(this.balance); // give back the old stake
            setFacts(); //reset the game
            StakeIncreased("stake increased", whale, this.balance, blockheight);
        }
    }
    
    // better safe than sorry
    function destroyWhenRoundOver() onlyOwner{
        mustBeDestroyed = true;
    }
    
    // next round we set a new staking perioud
    function setNewPeriod(uint _newPeriod) onlyOwner{
        newPeriod = _newPeriod;
    }
    
    function getPeriod() constant returns (uint){
        period;
    }
    
    //how long until a Matthew wins?
    function getBlocksTillMatthew() public constant returns(uint){
        if (blockheight + period > block.number)
            return blockheight + period - block.number;
        else
            return 0;
    }
    
    modifier onlyOwner(){
        if (msg.sender != owner) throw;
        _;
    }
}