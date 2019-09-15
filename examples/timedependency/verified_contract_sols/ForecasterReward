/**
 *  ForecasterReward.sol v1.1.0
 * 
 *  Bilal Arif - https://twitter.com/furusiyya_
 *  Draglet GbmH
 */

pragma solidity 0.4.19;

library SafeMath {
  function mul(uint256 a, uint256 b) pure internal returns (uint256) {
    uint256 c = a * b;
    assert(a == 0 || c / a == b);
    return c;
  }

  function div(uint256 a, uint256 b) pure internal returns (uint256) {
    uint256 c = a / b;
    return c;
  }

  function sub(uint256 a, uint256 b) pure internal returns (uint256) {
    assert(b <= a);
    return a - b;
  }

  function add(uint256 a, uint256 b) pure internal returns (uint256) {
    uint256 c = a + b;
    assert(c >= a);
    return c;
  }

  function max64(uint64 a, uint64 b) pure internal returns (uint64) {
    return a >= b ? a : b;
  }

  function min64(uint64 a, uint64 b) pure internal returns (uint64) {
    return a < b ? a : b;
  }

  function max256(uint256 a, uint256 b) pure internal returns (uint256) {
    return a >= b ? a : b;
  }

  function min256(uint256 a, uint256 b) pure internal returns (uint256) {
    return a < b ? a : b;
  }

}

/**
 * @title Ownable
 * @dev The Ownable contract has an owner address, and provides basic authorization control
 * functions, this simplifies the implementation of "user permissions".
 */
contract Ownable {
  address public owner;


  /**
   * @dev The Ownable constructor sets the original `owner` of the contract to the sender
   * account.
   */
  function Ownable() public {
    owner = msg.sender;
  }


  /**
   * @dev Throws if called by any account other than the owner.
   */
  modifier onlyOwner() {
    require(msg.sender == owner);
    _;
  }


  /**
   * @dev Allows the current owner to transfer control of the contract to a newOwner.
   * @param newOwner The address to transfer ownership to.
   */
  function transferOwnership(address newOwner) public onlyOwner {
    require(newOwner != address(0));      
    owner = newOwner;
  }

}

/*
 * Haltable
 *
 * Abstract contract that allows children to implement an
 * emergency stop mechanism. Differs from Pausable by causing a throw when in halt mode.
 *
 *
 * Originally envisioned in FirstBlood ICO contract.
 */
contract Haltable is Ownable {
  bool public halted;

  modifier stopInEmergency {
    assert(!halted);
    _;
  }

  modifier onlyInEmergency {
    require(halted);
    _;
  }

  // called by the owner on emergency, triggers stopped state
  function halt() external onlyOwner {
    halted = true;
  }

  // called by the owner on end of emergency, returns to normal state
  function unhalt() external onlyOwner onlyInEmergency {
    halted = false;
  }

}

contract ForecasterReward is Haltable {

  using SafeMath for uint;

  /* the starting time of the crowdsale */
  uint private startsAt;

  /* the ending time of the crowdsale */
  uint private endsAt;

  /* How many wei of funding we have received so far */
  uint private weiRaised = 0;

  /* How many distinct addresses have invested */
  uint private investorCount = 0;
  
  /* How many total investments have been made */
  uint private totalInvestments = 0;
  
  /* Address of pre-ico contract*/
  address private multisig;
 

  /** How much ETH each address has invested to this crowdsale */
  mapping (address => uint256) public investedAmountOf;

  
  /** State machine
   *
   * - Prefunding: We have not passed start time yet
   * - Funding: Active crowdsale
   * - Closed: Funding is closed.
   */
  enum State{PreFunding, Funding, Closed}

  // A new investment was made
  event Invested(uint index, address indexed investor, uint weiAmount);

  // Funds transfer to other address
  event Transfer(address indexed receiver, uint weiAmount);

  // Crowdsale end time has been changed
  event EndsAtChanged(uint endTimestamp);

  function ForecasterReward() public
  {

    owner = 0xed4C73Ad76D90715d648797Acd29A8529ED511A0;
    multisig = 0xDadF84E3adFc746e005D55aB427C1a8B1cc9cBA5;
    
    startsAt = 1515600000;
    endsAt = 1516118400;
  }

  /**
   * Allow investor to just send in money
   */
  function() nonZero payable public{
    buy(msg.sender);
  }

  /**
   * Make an investment.
   *
   * Crowdsale must be running for one to invest.
   * We must have not pressed the emergency brake.
   *
   * @param receiver The Ethereum address who have invested
   *
   */
  function buy(address receiver) stopInEmergency inState(State.Funding) nonZero public payable{
    require(receiver != 0x00);
    
    uint weiAmount = msg.value;
   
    if(investedAmountOf[receiver] == 0) {
      // A new investor
      investorCount++;
    }

    // count all investments
    totalInvestments++;

    // Update investor
    investedAmountOf[receiver] = investedAmountOf[receiver].add(weiAmount);
    
    // Up total accumulated fudns
    weiRaised = weiRaised.add(weiAmount);
    
    // Pocket the money
    if(!distributeFunds()) revert();
    
    // Tell us invest was success
    Invested(totalInvestments, receiver, weiAmount);
  }

 
  /**
   * @return multisig Address of Multisig Wallet contract
   */
  function multisigAddress() public constant returns(address){
      return multisig;
  }
  
  /**
   * @return startDate Crowdsale opening date
   */
  function fundingStartAt() public constant returns(uint ){
      return startsAt;
  }
  
  /**
   * @return endDate Crowdsale closing date
   */
  function fundingEndsAt() public constant returns(uint){
      return endsAt;
  }
  
  /**
   * @return investors Total of distinct investors
   */
  function distinctInvestors() public constant returns(uint){
      return investorCount;
  }
  
  /**
   * @return investments Crowdsale closing date
   */
  function investments() public constant returns(uint){
      return totalInvestments;
  }
  
  
  /**
   * Send out contributions imediately
   */
  function distributeFunds() private returns(bool){
        
    Transfer(multisig,this.balance);
    
    if(!multisig.send(this.balance)){
      return false;
    }
    
    return true;
  }
  
  /**
   * Allow crowdsale owner to close early or extend the crowdsale.
   *
   * This is useful e.g. for a manual soft cap implementation:
   * - after X amount is reached determine manual closing
   *
   * This may put the crowdsale to an invalid state,
   * but we trust owners know what they are doing.
   *
   */
  function setEndsAt(uint _endsAt) public onlyOwner {
    
    // Don't change past
    require(_endsAt > now);

    endsAt = _endsAt;
    EndsAtChanged(_endsAt);
  }

  /**
   * @return total of amount of wie collected by the contract 
   */
  function fundingRaised() public constant returns (uint){
    return weiRaised;
  }
  
  
  /**
   * Crowdfund state machine management.
   *
   * We make it a function and do not assign the result to a variable, so there is no chance of the variable being stale.
   */
  function getState() public constant returns (State) {
    if (now < startsAt) return State.PreFunding;
    else if (now <= endsAt) return State.Funding;
    else if (now > endsAt) return State.Closed;
  }

  /** Interface marker. */
  function isCrowdsale() public pure returns (bool) {
    return true;
  }

  //
  // Modifiers
  //
  /** Modifier allowing execution only if the crowdsale is currently running.  */
  modifier inState(State state) {
    require(getState() == state);
    _;
  }

  /** Modifier allowing execution only if received value is greater than zero */
  modifier nonZero(){
    require(msg.value > 0);
    _;
  }
}