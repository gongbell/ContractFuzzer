/**
 *Submitted for verification at Etherscan.io on 2017-12-08
*/

pragma solidity ^0.4.18;

/**
 * @title ERC20Basic
 * @dev Simpler version of ERC20 interface
 * @dev see https://github.com/ethereum/EIPs/issues/179
 */
contract ERC20Basic {
  uint256 public totalSupply;
  function balanceOf(address who) public view returns (uint256);
  function transfer(address to, uint256 value) public returns (bool);
  event Transfer(address indexed from, address indexed to, uint256 value);
}

/**
 * @title ERC20 interface
 * @dev see https://github.com/ethereum/EIPs/issues/20
 */
contract ERC20 is ERC20Basic {
  function allowance(address owner, address spender) public view returns (uint256);
  function transferFrom(address from, address to, uint256 value) public returns (bool);
  function approve(address spender, uint256 value) public returns (bool);
  event Approval(address indexed owner, address indexed spender, uint256 value);
}

/**
 * @title SafeMath
 * @dev Math operations with safety checks that throw on error
 */
library SafeMath {
  function mul(uint256 a, uint256 b) internal pure returns (uint256) {
    uint256 c = a * b;
    assert(a == 0 || c / a == b);
    return c;
  }

  function div(uint256 a, uint256 b) internal pure returns (uint256) {
    // assert(b > 0); // Solidity automatically throws when dividing by 0
    uint256 c = a / b;
    // assert(a == b * c + a % b); // There is no case in which this doesn't hold
    return c;
  }

  function sub(uint256 a, uint256 b) internal pure returns (uint256) {
    assert(b <= a);
    return a - b;
  }

  function add(uint256 a, uint256 b) internal pure returns (uint256) {
    uint256 c = a + b;
    assert(c >= a);
    return c;
  }
}

/**
 * @title Basic token
 * @dev Basic version of StandardToken, with no allowances.
 */
contract BasicToken is ERC20Basic {
  using SafeMath for uint256;

  mapping(address => uint256) balances;

  /**
  * @dev transfer token for a specified address
  * @param _to The address to transfer to.
  * @param _value The amount to be transferred.
  */
  function transfer(address _to, uint256 _value) public returns (bool) {
    require(_to != address(0));
    require(_value <= balances[msg.sender]);

    // SafeMath.sub will throw if there is not enough balance.
    balances[msg.sender] = balances[msg.sender].sub(_value);
    balances[_to] = balances[_to].add(_value);
    Transfer(msg.sender, _to, _value);
    return true;
  }

  /**
  * @dev Gets the balance of the specified address.
  * @param _owner The address to query the the balance of.
  * @return An uint256 representing the amount owned by the passed address.
  */
  function balanceOf(address _owner) public view returns (uint256 balance) {
    return balances[_owner];
  }

}

/**
 * @title Standard ERC20 token
 *
 * @dev Implementation of the basic standard token.
 * @dev https://github.com/ethereum/EIPs/issues/20
 * @dev Based on code by FirstBlood: https://github.com/Firstbloodio/token/blob/master/smart_contract/FirstBloodToken.sol
 */
contract StandardToken is ERC20, BasicToken {

  mapping (address => mapping (address => uint256)) internal allowed;


  /**
   * @dev Transfer tokens from one address to another
   * @param _from address The address which you want to send tokens from
   * @param _to address The address which you want to transfer to
   * @param _value uint256 the amount of tokens to be transferred
   */
  function transferFrom(address _from, address _to, uint256 _value) public returns (bool) {
    require(_to != address(0));
    require(_value <= balances[_from]);
    require(_value <= allowed[_from][msg.sender]);

    balances[_from] = balances[_from].sub(_value);
    balances[_to] = balances[_to].add(_value);
    allowed[_from][msg.sender] = allowed[_from][msg.sender].sub(_value);
    Transfer(_from, _to, _value);
    return true;
  }

  /**
   * @dev Approve the passed address to spend the specified amount of tokens on behalf of msg.sender.
   *
   * Beware that changing an allowance with this method brings the risk that someone may use both the old
   * and the new allowance by unfortunate transaction ordering. One possible solution to mitigate this
   * race condition is to first reduce the spender's allowance to 0 and set the desired value afterwards:
   * https://github.com/ethereum/EIPs/issues/20#issuecomment-263524729
   * @param _spender The address which will spend the funds.
   * @param _value The amount of tokens to be spent.
   */
  function approve(address _spender, uint256 _value) public returns (bool) {
    allowed[msg.sender][_spender] = _value;
    Approval(msg.sender, _spender, _value);
    return true;
  }

  /**
   * @dev Function to check the amount of tokens that an owner allowed to a spender.
   * @param _owner address The address which owns the funds.
   * @param _spender address The address which will spend the funds.
   * @return A uint256 specifying the amount of tokens still available for the spender.
   */
  function allowance(address _owner, address _spender) public view returns (uint256) {
    return allowed[_owner][_spender];
  }

  /**
   * approve should be called when allowed[_spender] == 0. To increment
   * allowed value is better to use this function to avoid 2 calls (and wait until
   * the first transaction is mined)
   * From MonolithDAO Token.sol
   */
  function increaseApproval(address _spender, uint _addedValue) public returns (bool) {
    allowed[msg.sender][_spender] = allowed[msg.sender][_spender].add(_addedValue);
    Approval(msg.sender, _spender, allowed[msg.sender][_spender]);
    return true;
  }

  function decreaseApproval(address _spender, uint _subtractedValue) public returns (bool) {
    uint oldValue = allowed[msg.sender][_spender];
    if (_subtractedValue > oldValue) {
      allowed[msg.sender][_spender] = 0;
    } else {
      allowed[msg.sender][_spender] = oldValue.sub(_subtractedValue);
    }
    Approval(msg.sender, _spender, allowed[msg.sender][_spender]);
    return true;
  }

}

/**
 * @title Ownable
 * @dev The Ownable contract has an owner address, and provides basic authorization control
 * functions, this simplifies the implementation of "user permissions".
 */
contract Ownable {
  address public owner;


  event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);


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
  function transferOwnership(address newOwner) onlyOwner public {
    require(newOwner != address(0));
    OwnershipTransferred(owner, newOwner);
    owner = newOwner;
  }

}

contract TaskFairToken is StandardToken, Ownable {	

  event Mint(address indexed to, uint256 amount);

  event MintFinished();
    
  string public constant name = "Task Fair Token";
   
  string public constant symbol = "TFT";
    
  uint32 public constant decimals = 18;

  bool public mintingFinished = false;
 
  address public saleAgent;

  modifier notLocked() {
    require(msg.sender == owner || msg.sender == saleAgent || mintingFinished);
    _;
  }

  function transfer(address _to, uint256 _value) public notLocked returns (bool) {
    return super.transfer(_to, _value);
  }

  function transferFrom(address from, address to, uint256 value) public notLocked returns (bool) {
    return super.transferFrom(from, to, value);
  }

  function setSaleAgent(address newSaleAgent) public {
    require(saleAgent == msg.sender || owner == msg.sender);
    saleAgent = newSaleAgent;
  }

  function mint(address _to, uint256 _amount) public returns (bool) {
    require(!mintingFinished);
    require(msg.sender == saleAgent);
    totalSupply = totalSupply.add(_amount);
    balances[_to] = balances[_to].add(_amount);
    Mint(_to, _amount);
    Transfer(address(0), _to, _amount);
    return true;
  }

  function finishMinting() public returns (bool) {
    require(!mintingFinished);
    require(msg.sender == owner || msg.sender == saleAgent);
    mintingFinished = true;
    MintFinished();
    return true;
  }

}

contract StagedCrowdsale is Ownable {

  using SafeMath for uint;

  uint public price;

  struct Stage {
    uint period;
    uint hardCap;
    uint discount;
    uint invested;
    uint closed;
  }

  uint public constant STAGES_PERCENT_RATE = 100;

  uint public start;

  uint public totalPeriod;

  uint public totalHardCap;
 
  uint public invested;

  Stage[] public stages;

  function stagesCount() public constant returns(uint) {
    return stages.length;
  }

  function setStart(uint newStart) public onlyOwner {
    start = newStart;
  }

  function setPrice(uint newPrice) public onlyOwner {
    price = newPrice;
  }

  function addStage(uint period, uint hardCap, uint discount) public onlyOwner {
    require(period > 0 && hardCap > 0);
    stages.push(Stage(period, hardCap, discount, 0, 0));
    totalPeriod = totalPeriod.add(period);
    totalHardCap = totalHardCap.add(hardCap);
  }

  function removeStage(uint8 number) public onlyOwner {
    require(number >=0 && number < stages.length);

    Stage storage stage = stages[number];
    totalHardCap = totalHardCap.sub(stage.hardCap);    
    totalPeriod = totalPeriod.sub(stage.period);

    delete stages[number];

    for (uint i = number; i < stages.length - 1; i++) {
      stages[i] = stages[i+1];
    }

    stages.length--;
  }

  function changeStage(uint8 number, uint period, uint hardCap, uint discount) public onlyOwner {
    require(number >= 0 && number < stages.length);

    Stage storage stage = stages[number];

    totalHardCap = totalHardCap.sub(stage.hardCap);    
    totalPeriod = totalPeriod.sub(stage.period);    

    stage.hardCap = hardCap;
    stage.period = period;
    stage.discount = discount;

    totalHardCap = totalHardCap.add(hardCap);    
    totalPeriod = totalPeriod.add(period);    
  }

  function insertStage(uint8 numberAfter, uint period, uint hardCap, uint discount) public onlyOwner {
    require(numberAfter < stages.length);


    totalPeriod = totalPeriod.add(period);
    totalHardCap = totalHardCap.add(hardCap);

    stages.length++;

    for (uint i = stages.length - 2; i > numberAfter; i--) {
      stages[i + 1] = stages[i];
    }

    stages[numberAfter + 1] = Stage(period, hardCap, discount, 0, 0);
  }

  function clearStages() public onlyOwner {
    for (uint i = 0; i < stages.length; i++) {
      delete stages[i];
    }
    stages.length -= stages.length;
    totalPeriod = 0;
    totalHardCap = 0;
  }

  function lastSaleDate() public constant returns(uint) {
    require(stages.length > 0);
    uint lastDate = start;
    for(uint i=0; i < stages.length; i++) {
      if(stages[i].invested >= stages[i].hardCap) {
        lastDate = stages[i].closed;
      } else {
        lastDate = lastDate.add(stages[i].period * 1 days);
      }
    }
    return lastDate;
  }

  function currentStage() public constant returns(uint) {
    require(now >= start);
    uint previousDate = start;
    for(uint i=0; i < stages.length; i++) {
      if(stages[i].invested < stages[i].hardCap) {
        if(now >= previousDate && now < previousDate + stages[i].period * 1 days) {
          return i;
        }
        previousDate = previousDate.add(stages[i].period * 1 days);
      } else {
        previousDate = stages[i].closed;
      }
    }
    revert();
  }

  function updateStageWithInvested(uint stageIndex, uint investedInWei) internal {
    invested = invested.add(investedInWei);
    Stage storage stage = stages[stageIndex];
    stage.invested = stage.invested.add(investedInWei);
    if(stage.invested >= stage.hardCap) {
      stage.closed = now;
    }
  }


}


contract CommonCrowdsale is StagedCrowdsale {

  uint public constant PERCENT_RATE = 1000;

  uint public minInvestedLimit;

  uint public minted;

  address public directMintAgent;
  
  address public wallet;

  address public devWallet;

  address public devTokensWallet;

  address public securityWallet;

  address public foundersTokensWallet;

  address public bountyTokensWallet;

  address public growthTokensWallet;

  address public advisorsTokensWallet;

  address public securityTokensWallet;

  uint public devPercent;

  uint public securityPercent;

  uint public bountyTokensPercent;

  uint public devTokensPercent;

  uint public advisorsTokensPercent;

  uint public foundersTokensPercent;

  uint public growthTokensPercent;

  uint public securityTokensPercent;

  TaskFairToken public token;

  modifier canMint(uint value) {
    require(now >= start && value >= minInvestedLimit);
    _;
  }

  modifier onlyDirectMintAgentOrOwner() {
    require(directMintAgent == msg.sender || owner == msg.sender);
    _;
  }

  function setMinInvestedLimit(uint newMinInvestedLimit) public onlyOwner {
    minInvestedLimit = newMinInvestedLimit;
  }

  function setDevPercent(uint newDevPercent) public onlyOwner { 
    devPercent = newDevPercent;
  }

  function setSecurityPercent(uint newSecurityPercent) public onlyOwner { 
    securityPercent = newSecurityPercent;
  }

  function setBountyTokensPercent(uint newBountyTokensPercent) public onlyOwner { 
    bountyTokensPercent = newBountyTokensPercent;
  }

  function setGrowthTokensPercent(uint newGrowthTokensPercent) public onlyOwner { 
    growthTokensPercent = newGrowthTokensPercent;
  }

  function setFoundersTokensPercent(uint newFoundersTokensPercent) public onlyOwner { 
    foundersTokensPercent = newFoundersTokensPercent;
  }

  function setAdvisorsTokensPercent(uint newAdvisorsTokensPercent) public onlyOwner { 
    advisorsTokensPercent = newAdvisorsTokensPercent;
  }

  function setDevTokensPercent(uint newDevTokensPercent) public onlyOwner { 
    devTokensPercent = newDevTokensPercent;
  }

  function setSecurityTokensPercent(uint newSecurityTokensPercent) public onlyOwner { 
    securityTokensPercent = newSecurityTokensPercent;
  }

  function setFoundersTokensWallet(address newFoundersTokensWallet) public onlyOwner { 
    foundersTokensWallet = newFoundersTokensWallet;
  }

  function setGrowthTokensWallet(address newGrowthTokensWallet) public onlyOwner { 
    growthTokensWallet = newGrowthTokensWallet;
  }

  function setBountyTokensWallet(address newBountyTokensWallet) public onlyOwner { 
    bountyTokensWallet = newBountyTokensWallet;
  }

  function setAdvisorsTokensWallet(address newAdvisorsTokensWallet) public onlyOwner { 
    advisorsTokensWallet = newAdvisorsTokensWallet;
  }

  function setDevTokensWallet(address newDevTokensWallet) public onlyOwner { 
    devTokensWallet = newDevTokensWallet;
  }

  function setSecurityTokensWallet(address newSecurityTokensWallet) public onlyOwner { 
    securityTokensWallet = newSecurityTokensWallet;
  }

  function setWallet(address newWallet) public onlyOwner { 
    wallet = newWallet;
  }

  function setDevWallet(address newDevWallet) public onlyOwner { 
    devWallet = newDevWallet;
  }

  function setSecurityWallet(address newSecurityWallet) public onlyOwner { 
    securityWallet = newSecurityWallet;
  }

  function setDirectMintAgent(address newDirectMintAgent) public onlyOwner {
    directMintAgent = newDirectMintAgent;
  }

  function directMint(address to, uint investedWei) public onlyDirectMintAgentOrOwner canMint(investedWei) {
    calculateAndTransferTokens(to, investedWei);
  }

  function setStart(uint newStart) public onlyOwner { 
    start = newStart;
  }

  function setToken(address newToken) public onlyOwner { 
    token = TaskFairToken(newToken);
  }

  function mintExtendedTokens() internal {
    uint extendedTokensPercent = bountyTokensPercent.add(devTokensPercent).add(advisorsTokensPercent).add(foundersTokensPercent).add(growthTokensPercent).add(securityTokensPercent);
    uint allTokens = minted.mul(PERCENT_RATE).div(PERCENT_RATE.sub(extendedTokensPercent));

    uint bountyTokens = allTokens.mul(bountyTokensPercent).div(PERCENT_RATE);
    mintAndSendTokens(bountyTokensWallet, bountyTokens);

    uint advisorsTokens = allTokens.mul(advisorsTokensPercent).div(PERCENT_RATE);
    mintAndSendTokens(advisorsTokensWallet, advisorsTokens);

    uint foundersTokens = allTokens.mul(foundersTokensPercent).div(PERCENT_RATE);
    mintAndSendTokens(foundersTokensWallet, foundersTokens);

    uint growthTokens = allTokens.mul(growthTokensPercent).div(PERCENT_RATE);
    mintAndSendTokens(growthTokensWallet, growthTokens);

    uint devTokens = allTokens.mul(devTokensPercent).div(PERCENT_RATE);
    mintAndSendTokens(devTokensWallet, devTokens);

    uint secuirtyTokens = allTokens.mul(securityTokensPercent).div(PERCENT_RATE);
    mintAndSendTokens(securityTokensWallet, secuirtyTokens);
  }

  function mintAndSendTokens(address to, uint amount) internal {
    token.mint(to, amount);
    minted = minted.add(amount);
  }

  function calculateAndTransferTokens(address to, uint investedInWei) internal {
    uint stageIndex = currentStage();
    Stage storage stage = stages[stageIndex];

    // calculate tokens
    uint tokens = investedInWei.mul(price).mul(STAGES_PERCENT_RATE).div(STAGES_PERCENT_RATE.sub(stage.discount)).div(1 ether);
    
    // transfer tokens
    mintAndSendTokens(to, tokens);

    updateStageWithInvested(stageIndex, investedInWei);
  }

  function createTokens() public payable;

  function() external payable {
    createTokens();
  }

  function retrieveTokens(address anotherToken) public onlyOwner {
    ERC20 alienToken = ERC20(anotherToken);
    alienToken.transfer(wallet, alienToken.balanceOf(this));
  }

}

contract PreTGE is CommonCrowdsale {
  
  uint public softcap;
  
  bool public refundOn;

  bool public softcapAchieved;

  address public nextSaleAgent;

  mapping (address => uint) public balances;

  event RefundsEnabled();

  event SoftcapReached();

  event Refunded(address indexed beneficiary, uint256 weiAmount);

  function PreTGE() public {
    setMinInvestedLimit(1000000000000000000);  
    setPrice(4000000000000000000000);
    setBountyTokensPercent(50);
    setAdvisorsTokensPercent(20);
    setDevTokensPercent(30);
    setFoundersTokensPercent(50);
    setGrowthTokensPercent(300);
    setSecurityTokensPercent(5);
    setDevPercent(20);
    setSecurityPercent(10);
    
    // fix in prod
    setSoftcap(40000000000000000000);
    
    addStage(7, 570000000000000000000, 40);
    addStage(7, 1400000000000000000000, 30);
    addStage(7, 2570000000000000000000, 20);
    
    setStart(1512867600);
    setWallet(0x73598a82559f3566Ecf93aab415323668124191C);
    setBountyTokensWallet(0x1C59BD0658DA5f357926D38083286A7E25Cd6f97);
    setDevTokensWallet(0xad3Df84A21d508Ad1E782956badeBE8725a9A447);
    setAdvisorsTokensWallet(0x17D34009D6e16Ae35dCfF3840d9eeC832d75FeA6);
    setFoundersTokensWallet(0xd63c6c4977B80a2042aA71bEd548e32A856e9481);
    setGrowthTokensWallet(0x9518ea93647DC3B198d3B04AD229977d8485fA1A);
    setDevWallet(0xad3Df84A21d508Ad1E782956badeBE8725a9A447);
    setSecurityTokensWallet(0x6Ea796DA599827ba871BE76fAF1948e45Bce4628);
    setSecurityWallet(0xfA4b94A9Ab8b5Ae3a1fd10aCE18724Bf1EC8CB07);
  }

  function setNextSaleAgent(address newNextSaleAgent) public onlyOwner {
    nextSaleAgent = newNextSaleAgent;
  }

  function setSoftcap(uint newSoftcap) public onlyOwner {
    softcap = newSoftcap;
  }

  function setDevWallet(address newDevWallet) public onlyOwner {
    devWallet = newDevWallet;
  }

  function refund() public {
    require(now > start && refundOn && balances[msg.sender] > 0);
    uint value = balances[msg.sender];
    balances[msg.sender] = 0;
    msg.sender.transfer(value);
    Refunded(msg.sender, value);
  } 

  function createTokens() public payable canMint(msg.value) {
    balances[msg.sender] = balances[msg.sender].add(msg.value);
    calculateAndTransferTokens(msg.sender, msg.value);
  } 

  function calculateAndTransferTokens(address to, uint investorWei) internal {
    super.calculateAndTransferTokens(to, investorWei);
    if(!softcapAchieved && invested >= softcap) {
      softcapAchieved = true;      
      SoftcapReached();
    }
  }

  function widthraw() public onlyOwner {
    require(softcapAchieved);
    uint devWei = this.balance.mul(devPercent).div(PERCENT_RATE);
    devWallet.transfer(devWei);
    uint securityWei = this.balance.mul(securityPercent).div(PERCENT_RATE);
    securityWallet.transfer(securityWei);
    wallet.transfer(this.balance);
  } 

  function finishMinting() public onlyOwner {
    if(!softcapAchieved) {
      refundOn = true;      
      token.finishMinting();
      RefundsEnabled();
    } else {
      widthraw();
      mintExtendedTokens();
      token.setSaleAgent(nextSaleAgent);
    }    
  }

}
