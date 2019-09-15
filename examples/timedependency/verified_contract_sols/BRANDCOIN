pragma solidity ^0.4.19;

/**
 * @title ERC20Basic
 * @dev Simpler version of ERC20 interface
 * @dev see https://github.com/ethereum/EIPs/issues/179
 */
contract ERC20Basic {
  function totalSupply() public view returns (uint256);
  function balanceOf(address who) public view returns (uint256);
  function transfer(address to, uint256 value) public returns (bool);
  event Transfer(address indexed from, address indexed to, uint256 value);
}

/**
 * @title SafeMath
 * @dev Math operations with safety checks that throw on error
 */
library SafeMath {

  /**
  * @dev Multiplies two numbers, throws on overflow.
  */
  function mul(uint256 a, uint256 b) internal pure returns (uint256) {
    if (a == 0) {
      return 0;
    }
    uint256 c = a * b;
    assert(c / a == b);
    return c;
  }

  /**
  * @dev Integer division of two numbers, truncating the quotient.
  */
  function div(uint256 a, uint256 b) internal pure returns (uint256) {
    // assert(b > 0); // Solidity automatically throws when dividing by 0
    uint256 c = a / b;
    // assert(a == b * c + a % b); // There is no case in which this doesn't hold
    return c;
  }

  /**
  * @dev Substracts two numbers, throws on overflow (i.e. if subtrahend is greater than minuend).
  */
  function sub(uint256 a, uint256 b) internal pure returns (uint256) {
    assert(b <= a);
    return a - b;
  }

  /**
  * @dev Adds two numbers, throws on overflow.
  */
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

  uint256 totalSupply_;

  /**
  * @dev total number of tokens in existence
  */
  function totalSupply() public view returns (uint256) {
    return totalSupply_;
  }

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
 * @title Burnable Token
 * @dev Token that can be irreversibly burned (destroyed).
 */
contract BurnableToken is BasicToken {

  event Burn(address indexed burner, uint256 value);

  /**
   * @dev Burns a specific amount of tokens.
   * @param _value The amount of token to be burned.
   */
  function burn(uint256 _value) public {
    require(_value <= balances[msg.sender]);
    // no need to require value <= totalSupply, since that would imply the
    // sender's balance is greater than the totalSupply, which *should* be an assertion failure

    address burner = msg.sender;
    balances[burner] = balances[burner].sub(_value);
    totalSupply_ = totalSupply_.sub(_value);
    Burn(burner, _value);
    Transfer(burner, address(0), _value);
  }
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
   * @dev Increase the amount of tokens that an owner allowed to a spender.
   *
   * approve should be called when allowed[_spender] == 0. To increment
   * allowed value is better to use this function to avoid 2 calls (and wait until
   * the first transaction is mined)
   * From MonolithDAO Token.sol
   * @param _spender The address which will spend the funds.
   * @param _addedValue The amount of tokens to increase the allowance by.
   */
  function increaseApproval(address _spender, uint _addedValue) public returns (bool) {
    allowed[msg.sender][_spender] = allowed[msg.sender][_spender].add(_addedValue);
    Approval(msg.sender, _spender, allowed[msg.sender][_spender]);
    return true;
  }

  /**
   * @dev Decrease the amount of tokens that an owner allowed to a spender.
   *
   * approve should be called when allowed[_spender] == 0. To decrement
   * allowed value is better to use this function to avoid 2 calls (and wait until
   * the first transaction is mined)
   * From MonolithDAO Token.sol
   * @param _spender The address which will spend the funds.
   * @param _subtractedValue The amount of tokens to decrease the allowance by.
   */
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
  function transferOwnership(address newOwner) public onlyOwner {
    require(newOwner != address(0));
    OwnershipTransferred(owner, newOwner);
    owner = newOwner;
  }

}

contract BRANDCOIN is StandardToken, BurnableToken, Ownable
{
    // ERC20 token parameters
    string public constant name = "BRANDCOIN";
    string public constant symbol = "BRA";
    uint256 public constant decimals = 18;
    
    // Crowdsale base price
    uint256 public ETH_per_BRA = 0.00024261 ether;
    
    // 15 april - 30 april: 43% bonus for purchases of at least 1000 BRA
    uint256 private first_period_start_date = 1523750400;
    uint256 private constant first_period_bonus_percentage = 43;
    uint256 private constant first_period_bonus_minimum_purchased_BRA = 1000 * (uint256(10) ** decimals);
    
    // 1 may - 7 may: 15% bonus
    uint256 private second_period_start_date = 1525132800;
    uint256 private constant second_period_bonus_percentage = 15;
    
    // 8 may - 14 may: 10% bonus
    uint256 private third_period_start_date = 1525737600;
    uint256 private constant third_period_bonus_percentage = 10;
    
    // 15 may - 21 may: 6% bonus
    uint256 private fourth_period_start_date = 1526342400;
    uint256 private constant fourth_period_bonus_percentage = 6;
    
    // 22 may - 31 may: 3% bonus
    uint256 private fifth_period_start_date = 1526947200;
    uint256 private constant fifth_period_bonus_percentage = 3;
    
    // End of ICO: 1 june
    uint256 private crowdsale_end_timestamp = 1527811200;
    
    // The target of the crowdsale is 8000000 BRANDCOIN's.
    // If the crowdsale has finished, and the target has not been reached,
    // all crowdsale participants will be able to call refund() and get their
    // ETH back. The refundMany() function can be used to refund multiple
    // participants in one transaction.
    uint256 public constant crowdsaleTargetBRA = 8000000 * (uint256(10) ** decimals);
    
    
    // Keep track of all participants, how much they bought and how much they spent.
    address[] public allParticipants;
    mapping(address => uint256) public participantToEtherSpent;
    mapping(address => uint256) public participantToBRAbought;
    
    
    function crowdsaleTargetReached() public view returns (bool)
    {
        return amountOfBRAsold() >= crowdsaleTargetBRA;
    }
    
    function crowdsaleStarted() public view returns (bool)
    {
        return now >= first_period_start_date;
    }
    
    function crowdsaleFinished() public view returns (bool)
    {
        return now >= crowdsale_end_timestamp;
    }
    
    function amountOfParticipants() external view returns (uint256)
    {
        return allParticipants.length;
    }
    
    function amountOfBRAsold() public view returns (uint256)
    {
        return totalSupply_ / 2 - balances[this];
    }
    
    // If the crowdsale target has not been reached, or the crowdsale has not finished,
    // don't allow the transfer of tokens purchased in the crowdsale.
    function transfer(address _to, uint256 _amount) public returns (bool)
    {
        if (!crowdsaleTargetReached() || !crowdsaleFinished())
        {
            require(balances[msg.sender] - participantToBRAbought[msg.sender] >= _amount);
        }
        
        return super.transfer(_to, _amount);
    }
    function transferFrom(address _from, address _to, uint256 _amount) public returns (bool)
    {
        if (!crowdsaleTargetReached() || !crowdsaleFinished())
        {
            require(balances[_from] - participantToBRAbought[_from] >= _amount);
        }
        
        return super.transferFrom(_from, _to, _amount);
    }
    
    address public founderWallet = 0x6bC5aa2B9eb4aa5b6170Dafce4482efF56184ADd;
    address public teamWallet = 0xb054D33607fC07e55469c81ABcB1553B92914E9e;
    address public bountyAffiliateWallet = 0x9460bc2bB546B640060E0268Ba8C392b0A0D6330;
    address public earlyBackersWallet = 0x4681B5c67ae0632c57ee206e1f9c2Ca58D6Af34c;
    address public reserveWallet = 0x4d70B2aCaE5e6558A9f5d55E672E93916Ba5c7aE;
    
    // Constructor function
    function BRANDCOIN() public
    {
        totalSupply_ = 1650000000 * (uint256(10) ** decimals);
        balances[this] = totalSupply_;
        Transfer(0x0, this, totalSupply_);
    }
    
    bool private distributedInitialFunds = false;
    function distributeInitialFunds() public onlyOwner
    {
        require(!distributedInitialFunds);
        distributedInitialFunds = true;
        this.transfer(founderWallet, totalSupply_*15/100);
        this.transfer(earlyBackersWallet, totalSupply_*5/100);
        this.transfer(teamWallet, totalSupply_*15/100);
        this.transfer(bountyAffiliateWallet, totalSupply_*5/100);
        this.transfer(reserveWallet, totalSupply_*10/100);
    }
    
    function destroyUnsoldTokens() external
    {
        require(crowdsaleStarted() && crowdsaleFinished());
        
        this.burn(balances[this]);
    }
    
    // If someone sends ETH to the contract address,
    // assume that they are trying to buy tokens.
    function () payable external
    {
        buyTokens();
    }
    
    function buyTokens() payable public
    {
        uint256 amountOfBRApurchased = msg.value * (uint256(10)**decimals) / ETH_per_BRA;
        
        // Only allow buying tokens if the ICO has started, and has not finished
        require(crowdsaleStarted());
        require(!crowdsaleFinished());
        
        // If the pre-ICO hasn't started yet, cancel the transaction
        if (now < first_period_start_date)
        {
            revert();
        }
        
        else if (now >= first_period_start_date && now < second_period_start_date)
        {
            if (amountOfBRApurchased >= first_period_bonus_minimum_purchased_BRA)
            {
                amountOfBRApurchased = amountOfBRApurchased * (100 + first_period_bonus_percentage) / 100;
            }
        }
        
        else if (now >= second_period_start_date && now < third_period_start_date)
        {
            amountOfBRApurchased = amountOfBRApurchased * (100 + second_period_bonus_percentage) / 100;
        }
        
        else if (now >= third_period_start_date && now < fourth_period_start_date)
        {
            amountOfBRApurchased = amountOfBRApurchased * (100 + third_period_bonus_percentage) / 100;
        }
        
        else if (now >= fourth_period_start_date && now < fifth_period_start_date)
        {
            amountOfBRApurchased = amountOfBRApurchased * (100 + fourth_period_bonus_percentage) / 100;
        }
        
        else if (now >= fifth_period_start_date && now < crowdsale_end_timestamp)
        {
            amountOfBRApurchased = amountOfBRApurchased * (100 + fifth_period_bonus_percentage) / 100;
        }
        
        // If we are passed the final sale, cancel the transaction.
        else
        {
            revert();
        }
        
        // Send the purchased tokens to the buyer
        this.transfer(msg.sender, amountOfBRApurchased);
        
        // Track statistics
        if (participantToEtherSpent[msg.sender] == 0)
        {
            allParticipants.push(msg.sender);
        }
        participantToBRAbought[msg.sender] += amountOfBRApurchased;
        participantToEtherSpent[msg.sender] += msg.value;
    }
    
    function refund() external
    {
        // If the crowdsale has not started yet, don't allow refund
        require(crowdsaleStarted());
        
        // If the crowdsale has not finished yet, don't allow refund
        require(crowdsaleFinished());
        
        // If the target was reached, don't allow refund
        require(!crowdsaleTargetReached());
        
        _refundParticipant(msg.sender);
    }
    
    function refundMany(uint256 _startIndex, uint256 _endIndex) external
    {
        // If the crowdsale has not started yet, don't allow refund
        require(crowdsaleStarted());
        
        // If the crowdsale has not finished yet, don't allow refund
        require(crowdsaleFinished());
        
        // If the target was reached, don't allow refund
        require(!crowdsaleTargetReached());
        
        for (uint256 i=_startIndex; i<=_endIndex && i<allParticipants.length; i++)
        {
            _refundParticipant(allParticipants[i]);
        }
    }
    
    function _refundParticipant(address _participant) internal
    {
        if (participantToEtherSpent[_participant] > 0)
        {
            // Return the BRA they bought to this contract
            uint256 refundBRA = participantToBRAbought[_participant];
            participantToBRAbought[_participant] = 0;
            balances[_participant] -= refundBRA;
            balances[this] += refundBRA;
            Transfer(_participant, this, refundBRA);
            
            // Return the ETH they spent to buy them
            uint256 refundETH = participantToEtherSpent[_participant];
            participantToEtherSpent[_participant] = 0;
            _participant.transfer(refundETH);
        }
    }
    
    function ownerWithdrawETH() external onlyOwner
    {
        // Only allow the owner to withdraw if the crowdsale target has been reached
        require(crowdsaleTargetReached());
        owner.transfer(this.balance);
    }
    
    // As long as the crowdsale has not started yet, the owner can change the base price
    function setPrice(uint256 _ETH_PER_BRA) external onlyOwner
    {
        require(!crowdsaleStarted());
        ETH_per_BRA = _ETH_PER_BRA;
    }
}