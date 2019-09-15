pragma solidity ^0.4.18;

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
 * The Bitkom Token (BTT) has a fixed supply and restricts the ability
 * to transfer tokens until the owner has called the enableTransfer()
 * function.
 *
 * The owner can associate the token with a token sale contract. In that
 * case, the token balance is moved to the token sale contract, which
 * in turn can transfer its tokens to contributors to the sale.
 */
contract BitkomToken is StandardToken, Ownable {

    // Constants
    string  public constant name = "Bitkom Token";
    string  public constant symbol = "BTT";
    uint8   public constant decimals = 18;
    uint256 public constant INITIAL_SUPPLY = 50000000 * 1 ether;
    uint256 public constant CROWDSALE_ALLOWANCE =  33500000 * 1 ether;
    uint256 public constant TEAM_ALLOWANCE =  16500000 * 1 ether;

    // Properties
    uint256 public crowdsaleAllowance;               // the number of tokens available for crowdsales
    uint256 public teamAllowance;               // the number of tokens available for the administrator
    address public crowdsaleAddr;                    // the address of a crowdsale currently selling this token
    address public teamAddr;                    // the address of the team account
    bool    public transferEnabled = false;     // indicates if transferring tokens is enabled or not

    // Modifiers
    modifier onlyWhenTransferEnabled() {
        if (!transferEnabled) {
            require(msg.sender == teamAddr || msg.sender == crowdsaleAddr);
        }
        _;
    }

    // Events
    event Burn(address indexed burner, uint256 value);

    /**
     * The listed addresses are not valid recipients of tokens.
     *
     * 0x0           - the zero address is not valid
     * this          - the contract itself should not receive tokens
     * owner         - the owner has all the initial tokens, but cannot receive any back
     * teamAddr      - the team has an allowance of tokens to transfer, but does not receive any
     * crowdsaleAddr      - the sale has an allowance of tokens to transfer, but does not receive any
     */
    modifier validDestination(address _to) {
        require(_to != address(0x0));
        require(_to != address(this));
        require(_to != owner);
        require(_to != address(teamAddr));
        require(_to != address(crowdsaleAddr));
        _;
    }

    /**
     * Constructor - instantiates token supply and allocates balanace of
     * to the owner (msg.sender).
     */
    function BitkomToken(address _team) public {
        // the owner is a custodian of tokens that can
        // give an allowance of tokens for crowdsales
        // or to the admin, but cannot itself transfer
        // tokens; hence, this requirement
        require(msg.sender != _team);

        totalSupply = INITIAL_SUPPLY;
        crowdsaleAllowance = CROWDSALE_ALLOWANCE;
        teamAllowance = TEAM_ALLOWANCE;

        // mint all tokens
        balances[msg.sender] = totalSupply;
        Transfer(address(0x0), msg.sender, totalSupply);

        teamAddr = _team;
        approve(teamAddr, teamAllowance);
    }

    /**
     * Associates this token with a current crowdsale, giving the crowdsale
     * an allowance of tokens from the crowdsale supply. This gives the
     * crowdsale the ability to call transferFrom to transfer tokens to
     * whomever has purchased them.
     *
     * Note that if _amountForSale is 0, then it is assumed that the full
     * remaining crowdsale supply is made available to the crowdsale.
     *
     * @param _crowdsaleAddr The address of a crowdsale contract that will sell this token
     * @param _amountForSale The supply of tokens provided to the crowdsale
     */
    function setCrowdsale(address _crowdsaleAddr, uint256 _amountForSale) external onlyOwner {
        require(!transferEnabled);
        require(_amountForSale <= crowdsaleAllowance);

        // if 0, then full available crowdsale supply is assumed
        uint amount = (_amountForSale == 0) ? crowdsaleAllowance : _amountForSale;

        // Clear allowance of old, and set allowance of new
        approve(crowdsaleAddr, 0);
        approve(_crowdsaleAddr, amount);

        crowdsaleAddr = _crowdsaleAddr;
    }

    /**
     * Enables the ability of anyone to transfer their tokens. This can
     * only be called by the token owner. Once enabled, it is not
     * possible to disable transfers.
     */
    function enableTransfer() external onlyOwner {
        transferEnabled = true;
        approve(crowdsaleAddr, 0);
        approve(teamAddr, 0);
        crowdsaleAllowance = 0;
        teamAllowance = 0;
    }

    /**
     * Overrides ERC20 transfer function with modifier that prevents the
     * ability to transfer tokens until after transfers have been enabled.
     */
    function transfer(address _to, uint256 _value) public onlyWhenTransferEnabled validDestination(_to) returns (bool) {
        return super.transfer(_to, _value);
    }

    /**
     * Overrides ERC20 transferFrom function with modifier that prevents the
     * ability to transfer tokens until after transfers have been enabled.
     */
    function transferFrom(address _from, address _to, uint256 _value) public onlyWhenTransferEnabled validDestination(_to) returns (bool) {
        bool result = super.transferFrom(_from, _to, _value);
        if (result) {
            if (msg.sender == crowdsaleAddr)
                crowdsaleAllowance = crowdsaleAllowance.sub(_value);
            if (msg.sender == teamAddr)
                teamAllowance = teamAllowance.sub(_value);
        }
        return result;
    }

    /**
     * @dev Burns a specific amount of tokens if msg.sender == owner
     * or transferEnabled == true
     * @param _value The amount of token to be burned.
     */
    function burn(uint256 _value) public {
        require(_value > 0);
        require(_value <= balances[msg.sender]);
        require(transferEnabled || msg.sender == owner);
        // no need to require value <= totalSupply, since that would imply the
        // sender's balance is greater than the totalSupply, which *should* be an assertion failure

        address burner = msg.sender;
        balances[burner] = balances[burner].sub(_value);
        totalSupply = totalSupply.sub(_value);
        Burn(burner, _value);
        Transfer(msg.sender, address(0x0), _value);
    }
}

/**
 * The BitkomSale smart contract is used for selling BitkomToken
 * tokens (BTT). It does so by converting ETH received into a quantity of
 * tokens that are transferred to the contributor via the ERC20-compatible
 * transferFrom() function.
 */
contract BitkomSale is Pausable {

    using SafeMath for uint256;

    uint public constant RATE = 2500;       // constant for converting ETH to BTT
    uint public constant GAS_LIMIT_IN_WEI = 50000000000 wei;
    uint public constant MIN_CONTRIBUTION = 0.1 * 1 ether;  // lower bound on amount a contributor can send (in wei)
    uint public constant TOKEN_HARDCAP = 33500000 * 1 ether;  // lower bound on amount a contributor can send (in wei)

    bool public fundingCapReached = false;  // funding cap has been reached
    bool public tokenHardcapReached = false;  // TOKEN_HARDCAP has been reached
    bool public softcapReached = false;     // softcap has been reached
    bool public saleClosed = false;         // crowdsale is closed or not
    bool private rentrancy_lock = false;    // prevent certain functions from recursize calls

    uint public fundingCap;                 // upper bound on amount that can be raised (in wei)
    uint256 public soldTokens = 0;          // amount of sold tokens 
    uint256 public softCapInTokens = 1600000 * 1 ether;   // softcap in tokens for project launch

    uint public weiRaised;                  // amount of raised wei
    uint public weiRefunded;                  // amount of refunded wei

    uint public startTime;                  // UNIX timestamp for start of sale
    uint public deadline;                   // UNIX timestamp for end (deadline) of sale

    address public beneficiary;             // The beneficiary is the future recipient of the funds

    BitkomToken public tokenReward;     // The token being sold

    mapping (address => uint256) public balanceOf;   // tracks the amount of wei contributed by address during all sales
    mapping (address => bool) refunded; // tracks the status of refunding for each address

    // Events
    event CapReached(address _beneficiary, uint _weiRaised);
    event SoftcapReached(address _beneficiary, uint _weiRaised);
    event FundTransfer(address _backer, uint _amount, bool _isContribution);
    event Refunded(address indexed holder, uint256 amount);

    // Modifiers
    modifier beforeDeadline()   { require (currentTime() < deadline); _; }
    modifier afterDeadline()    { require (currentTime() >= deadline); _; }
    modifier afterStartTime()   { require (currentTime() >= startTime); _; }
    modifier saleNotClosed()    { require (!saleClosed); _; }
    modifier softCapRaised()    { require (softcapReached); _; }

    modifier nonReentrant() {
        require(!rentrancy_lock);
        rentrancy_lock = true;
        _;
        rentrancy_lock = false;
    }

    /**
     * Constructor for a crowdsale of BitkomToken tokens.
     *
     * @param ifSuccessfulSendTo            the beneficiary of the fund
     * @param fundingCapInEthers            the cap (maximum) size of the fund
     * @param start                         the start time (UNIX timestamp)
     * @param durationInDays                the duration of the crowdsale in days
     * @param addressOfTokenUsedAsReward    address of the token being sold
     */
    function BitkomSale(
        address ifSuccessfulSendTo,
        uint256 fundingCapInEthers,
        uint256 start,
        uint256 durationInDays,
        address addressOfTokenUsedAsReward
    ) public 
    {      
        require(ifSuccessfulSendTo != address(0) && ifSuccessfulSendTo != address(this));
        require(addressOfTokenUsedAsReward != address(0) && addressOfTokenUsedAsReward != address(this));
        require(durationInDays > 0);
        beneficiary = ifSuccessfulSendTo;
        fundingCap = fundingCapInEthers * 1 ether;
        startTime = start;
        deadline = start + (durationInDays * 1 days);
        tokenReward = BitkomToken(addressOfTokenUsedAsReward);
    }


    function () public payable {
        buy();
    }


    function buy()
        payable
        public
        whenNotPaused
        beforeDeadline
        afterStartTime
        saleNotClosed
        nonReentrant
    {
        uint amount = msg.value;
        require(amount >= MIN_CONTRIBUTION);

        weiRaised = weiRaised.add(amount);

        //require(weiRaised <= fundingCap);
        // if we overflow the fundingCap, transfer the overflow amount
        if (weiRaised > fundingCap) {
            uint overflow = weiRaised.sub(fundingCap);
            amount = amount.sub(overflow);
            weiRaised = fundingCap;
            // transfer overflow back to the user
            msg.sender.transfer(overflow);
        }

        // Calculate bonus for user
        uint256 bonus = calculateBonus();

        // Calculate amount of tokens for user
        uint256 tokensAmountForUser = (amount.mul(RATE)).mul(bonus);

        soldTokens = soldTokens.add(tokensAmountForUser);

        // 1 - вернуть лишние бабки и отдать токены чтоб их было впритык
        if (soldTokens > TOKEN_HARDCAP) {
            uint256 overflowInTokens = soldTokens.sub(TOKEN_HARDCAP);
            uint256 overflowInWei = (overflowInTokens.div(bonus)).div(RATE);
            amount = amount.sub(overflowInWei);
            weiRaised = weiRaised.sub(overflowInWei);
            // transfer overflow back to the user
            msg.sender.transfer(overflowInWei);

            // update amount of tokens
            tokensAmountForUser = tokensAmountForUser.sub(overflowInTokens);

            soldTokens = TOKEN_HARDCAP;
        }


        // Update the sender's balance of wei contributed and the total amount raised
        balanceOf[msg.sender] = balanceOf[msg.sender].add(amount);

        // Transfer the tokens from the crowdsale supply to the sender
        if (!tokenReward.transferFrom(tokenReward.owner(), msg.sender, tokensAmountForUser)) {
            revert();
        }

        FundTransfer(msg.sender, amount, true);

        if (soldTokens >= softCapInTokens && !softcapReached) {
            softcapReached = true;
            SoftcapReached(beneficiary, weiRaised);
        }

        checkCaps();
    }

    /**
     * The owner can terminate the crowdsale at any time.
     */
    function terminate() external onlyOwner {
        saleClosed = true;
    }


    /**
     * The owner can call this function to withdraw the funds that
     * have been sent to this contract. The funds will be sent to
     * the beneficiary specified when the crowdsale was created.
     * ONLY IF SOLDTOKENS >= SOFTCAPINTOKENS !!!!!!!!!!!!!!!!!!
     */
    function ownerSafeWithdrawal() external onlyOwner softCapRaised nonReentrant {
        uint balanceToSend = this.balance;
        beneficiary.transfer(balanceToSend);
        FundTransfer(beneficiary, balanceToSend, false);
    }

    /**
     * Checks if the funding cap or TOKEN_HARDCAP has been reached. 
     * If it has, then the CapReached event is triggered.
     */
    function checkCaps() internal {
        if (weiRaised == fundingCap) {
            // Check if the funding cap has been reached
            fundingCapReached = true;
            saleClosed = true;
            CapReached(beneficiary, weiRaised);
        }
        if (soldTokens == TOKEN_HARDCAP) {
            // Check if the funding cap has been reached
            tokenHardcapReached = true;
            saleClosed = true;
            CapReached(beneficiary, weiRaised);
        }
    }

    /**
     * Returns the current time.
    */
    function currentTime() internal constant returns (uint _currentTime) {
        return now;
    }

    /**
     * Returns the bonus value.
    */
    function calculateBonus() internal constant returns (uint) {
        if (soldTokens >= 0 && soldTokens <= 10000000 * 1 ether) {
            return 4;
        } else if (soldTokens > 10000000 * 1 ether && soldTokens <= 20000000 * 1 ether) {
            return 3;
        } else if (soldTokens > 20000000 * 1 ether && soldTokens <= 30000000 * 1 ether) {
            return 2;
        } else {
            return 1;
        }
    }

    function refund() external afterDeadline {
        require(!softcapReached);
        require(refunded[msg.sender] == false);

        uint256 balance = this.balanceOf(msg.sender);
        require(balance > 0);

        uint refund = balance;
        if (refund > this.balance) {
            refund = this.balance;
        }

        if (!msg.sender.send(refund)) {
            revert();
        }
        refunded[msg.sender] = true;
        weiRefunded = weiRefunded.add(refund);
        Refunded(msg.sender, refund);
    }
}