pragma solidity ^0.4.18;

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

/**
 * @title Pausable
 * @dev Base contract which allows children to implement an emergency stop mechanism.
 */
contract Pausable is Ownable {
  event Pause();
  event Unpause();

  bool public paused = false;


  /**
   * @dev Modifier to make a function callable only when the contract is not paused.
   */
  modifier whenNotPaused() {
    require(!paused);
    _;
  }

  /**
   * @dev Modifier to make a function callable only when the contract is paused.
   */
  modifier whenPaused() {
    require(paused);
    _;
  }

  /**
   * @dev called by the owner to pause, triggers stopped state
   */
  function pause() onlyOwner whenNotPaused public {
    paused = true;
    Pause();
  }

  /**
   * @dev called by the owner to unpause, returns to normal state
   */
  function unpause() onlyOwner whenPaused public {
    paused = false;
    Unpause();
  }
}



/**
 * @title SafeMath
 * @dev Math operations with safety checks that throw on error
 */
library SafeMath {
  function mul(uint256 a, uint256 b) internal pure returns (uint256) {
    if (a == 0) {
      return 0;
    }
    uint256 c = a * b;
    assert(c / a == b);
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
 * @title DropletToken
 */
contract DropletToken is StandardToken, Pausable {
    
    string public constant name = "Droplet Token";
    string public constant symbol = "DPLT";
    uint32 public constant decimals = 18;
    
    /**
     * @dev Give all tokens to msg.sender.
     */
    function DropletToken(uint _totalSupply) public {
        require (_totalSupply > 0);
        totalSupply = _totalSupply;
        balances[msg.sender] = _totalSupply;
    }

    function transfer(address _to, uint _value) public whenNotPaused returns (bool) {
        return super.transfer(_to, _value);
    }

    function transferFrom(address _from, address _to, uint _value) public whenNotPaused returns (bool) {
        return super.transferFrom(_from, _to, _value);
    }
}

/**
 * @title DropletTokenCrowdSale
 * @title crowdsale contract.
 */
contract DropletCrowdSale is Pausable {
    using SafeMath for uint256;

    address beneficiaryAddress;

    // token instance
    DropletToken public token;

    uint256 public maxTokensAmount;
    uint256 public issuedTokensAmount = 0;
    uint256 public tokenRate; // token per 1 ETH

    uint256 public endDate;

    bool public isFinished = false;
    bool public isUnlocked = false;

    // buffer for claimable tokens
    mapping(address => uint256) public tokens;
    mapping(uint32 => address) public tokenReceivers;
    uint32 public receiversCount = 0;

    /**
    * Events for token purchase logging
    */
    event TokenBought(address indexed _buyer, uint256 _tokens, uint256 _amount);
    event TokenAdded(address indexed _receiver, uint256 _tokens, uint256 _equivalentAmount);
    event TokenToppedUp(address indexed _receiver, uint256 _tokens, uint256 _equivalentAmount);
    event TokenSubtracted(address indexed _receiver, uint256 _tokens, uint256 _equivalentAmount);
    event TokenSent(address indexed _receiver, uint256 _tokens);

    modifier inProgress() {
        require (!isFinished);
        require (issuedTokensAmount < maxTokensAmount);
        require (now <= endDate);
        _;
    }

    function DropletCrowdSale(
        address _tokenAddress,
        address _beneficiaryAddress,
        uint256 _tokenRate,
        uint256 _maxTokensAmount,
        uint256 _endDate
    ) public {
        token = DropletToken(_tokenAddress);
        beneficiaryAddress = _beneficiaryAddress;

        tokenRate = _tokenRate;
        maxTokensAmount = _maxTokensAmount;

        endDate = _endDate;
    }

    /*
     * @dev Set new Droplet token exchange rate.
     */
    function setTokenRate(uint256 _tokenRate) public onlyOwner {
        require (_tokenRate > 0);
        tokenRate = _tokenRate;
    }

    /*
     * @dev Set new Droplet token sale end date.
     */
    function setEndDate(uint256 _endDate) public onlyOwner {
        endDate = _endDate;
    }

    /**
     * Buy Droplet. Tokens will be stored in contract until claim stage
     */
    function buy() public payable inProgress whenNotPaused {
        uint256 payAmount = msg.value;
        uint256 returnAmount = 0;

        // calculate token amount to be transfered to investor
        uint256 tokensAmount = tokenRate.mul(payAmount);

        if (issuedTokensAmount + tokensAmount > maxTokensAmount) {
            tokensAmount = maxTokensAmount.sub(issuedTokensAmount);
            payAmount = tokensAmount.div(tokenRate);
            returnAmount = msg.value.sub(payAmount);
        }

        issuedTokensAmount = issuedTokensAmount.add(tokensAmount);
        require (issuedTokensAmount <= maxTokensAmount);

        storeTokens(msg.sender, tokensAmount);
        TokenBought(msg.sender, tokensAmount, payAmount);

        beneficiaryAddress.transfer(payAmount);

        if (returnAmount > 0) {
            msg.sender.transfer(returnAmount);
        }
    }

    /**
     * Add Tokens for owner by ETH. Tokens will be stored in contract until claim stage
     */
    function add(address _receiver, uint256 _equivalentEthAmount) public onlyOwner inProgress whenNotPaused {
        uint256 tokensAmount = tokenRate.mul(_equivalentEthAmount);
        issuedTokensAmount = issuedTokensAmount.add(tokensAmount);

        storeTokens(_receiver, tokensAmount);
        TokenAdded(_receiver, tokensAmount, _equivalentEthAmount);
    }

    /**
     * Topup Tokens for owner. Tokens will be stored in contract until claim stage
     */
    function topUp(address _receiver, uint256 _equivalentEthAmount) public onlyOwner whenNotPaused {
        uint256 tokensAmount = tokenRate.mul(_equivalentEthAmount);
        issuedTokensAmount = issuedTokensAmount.add(tokensAmount);

        storeTokens(_receiver, tokensAmount);
        TokenToppedUp(_receiver, tokensAmount, _equivalentEthAmount);
    }

    /**
     * Reduce bought token amount. Emergency use only
     */
    function sub(address _receiver, uint256 _equivalentEthAmount) public onlyOwner whenNotPaused {
        uint256 tokensAmount = tokenRate.mul(_equivalentEthAmount);

        require (tokens[_receiver] >= tokensAmount);

        tokens[_receiver] = tokens[_receiver].sub(tokensAmount);
        issuedTokensAmount = issuedTokensAmount.sub(tokensAmount);

        TokenSubtracted(_receiver, tokensAmount, _equivalentEthAmount);
    }

    /**
     * Internal method for storing tokens in contract until claim stage
     */
    function storeTokens(address _receiver, uint256 _tokensAmount) internal whenNotPaused {
        if (tokens[_receiver] == 0) {
            tokenReceivers[receiversCount] = _receiver;
            receiversCount++;
        }
        tokens[_receiver] = tokens[_receiver].add(_tokensAmount);
    }

    /**
     * Claim all bought tokens. Available tokens will be sent to transaction sender address if unlocked
     */
    function claim() public whenNotPaused {
        claimFor(msg.sender);
    }

    /**
     * Claim all bought tokens for specific address
     */
    function claimOne(address _receiver) public onlyOwner whenNotPaused {
        claimFor(_receiver);
    }

    /**
     * Claim all bought tokens for all addresses
     */
    function claimAll() public onlyOwner whenNotPaused {
        for (uint32 i = 0; i < receiversCount; i++) {
            address receiver = tokenReceivers[i];
            if (isUnlocked && tokens[receiver] > 0) {
                claimFor(receiver);
            }
        }
    }

    /**
     * Internal method for claiming tokens for specific address
     */
    function claimFor(address _receiver) internal whenNotPaused {
        require(isUnlocked);
        require(tokens[_receiver] > 0);

        uint256 tokensToSend = tokens[_receiver];
        tokens[_receiver] = 0;

        require(token.transferFrom(owner, _receiver, tokensToSend));
        TokenSent(_receiver, tokensToSend);
    }

    function unLockTokens() public onlyOwner whenNotPaused {
        isUnlocked = true;
    }

    function lockTokens() public onlyOwner whenNotPaused {
        isUnlocked = false;
    }

    /**
     * Finish Sale.
     */
    function finish() public onlyOwner {
        require (!isFinished);
        isFinished = true;
    }

    function getReceiversCount() public constant onlyOwner returns (uint32) {
        return receiversCount;
    }

    function getReceiver(uint32 i) public constant onlyOwner returns (address) {
        return tokenReceivers[i];
    }

    /**
     * Buy token. Tokens will be stored in contract until claim stage
     */
    function() external payable {
        buy();
    }
}