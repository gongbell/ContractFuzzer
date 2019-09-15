pragma solidity ^0.4.16;



/**
 * @title ERC20Basic
 * @dev Simpler version of ERC20 interface
 * @dev see https://github.com
 */
contract ERC20Basic {
  uint256 public totalSupply;
  function balanceOf(address who) constant returns (uint256);
  function transfer(address to, uint256 value) returns (bool);
  event Transfer(address indexed from, address indexed to, uint256 value);
}

/**
 * @title ERC20 interface
 * @dev see https://github.com
 */
contract ERC20 is ERC20Basic {
  function allowance(address owner, address spender) constant returns (uint256);
  function transferFrom(address from, address to, uint256 value) returns (bool);
  function approve(address spender, uint256 value) returns (bool);
  event Approval(address indexed owner, address indexed spender, uint256 value);
}

/**
 * @title SafeMath
 * @dev Math operations with safety checks that throw on error
 */
library SafeMath {
    
  function mul(uint256 a, uint256 b) internal constant returns (uint256) {
    uint256 c = a * b;
    assert(a == 0 || c / a == b);
    return c;
  }

  function div(uint256 a, uint256 b) internal constant returns (uint256) {
    // assert(b > 0); // Solidity automatically throws when dividing by 0
    uint256 c = a / b;
    // assert(a == b * c + a % b); // There is no case in which this doesn't hold
    return c;
  }

  function sub(uint256 a, uint256 b) internal constant returns (uint256) {
    assert(b <= a);
    return a - b;
  }

  function add(uint256 a, uint256 b) internal constant returns (uint256) {
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


  modifier onlyPayloadSize(uint size) {
    require(msg.data.length >= size + 4);
    _;
  }

  /**
  * @dev transfer token for a specified address
  * @param _to The address to transfer to.
  * @param _value The amount to be transferred.
  */

  function transfer(address _to, uint256 _value) onlyPayloadSize(32*2) returns (bool) {
    require(_to != address(0));
    require(_value <= balances[msg.sender]);
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
  function balanceOf(address _owner) constant returns (uint256 balance) {
    return balances[_owner];
  }

}

/**
 * @title Standard ERC20 token
 *
 * @dev Implementation of the basic standard token.
 * @dev https://github.com/ethereum/EIPs/issues/20
 * @dev Based on code by FirstBlood: https://github.com
 */
contract StandardToken is ERC20, BasicToken {

  mapping (address => mapping (address => uint256)) allowed;

  /**
   * @dev Transfer tokens from one address to another
   * @param _from address The address which you want to send tokens from
   * @param _to address The address which you want to transfer to
   * @param _value uint256 the amout of tokens to be transfered
   */
  function transferFrom(address _from, address _to, uint256 _value) returns (bool) {
    var _allowance = allowed[_from][msg.sender];

    // Check is not needed because sub(_allowance, _value) will already throw if this condition is not met
    require (_value <= _allowance);
    require(_to != address(0));
    require(_value <= balances[_from]);

    balances[_to] = balances[_to].add(_value);
    balances[_from] = balances[_from].sub(_value);
    allowed[_from][msg.sender] = _allowance.sub(_value);
    Transfer(_from, _to, _value);
    return true;
  }

  /**
   * @dev Aprove the passed address to spend the specified amount of tokens on behalf of msg.sender.
   * @param _spender The address which will spend the funds.
   * @param _value The amount of tokens to be spent.
   */
  function approve(address _spender, uint256 _value) returns (bool) {

    // To change the approve amount you first have to reduce the addresses`
    //  allowance to zero by calling `approve(_spender, 0)` if it is not
    //  already 0 to mitigate the race condition described here:
    //  https://github.com/
    require((_value == 0) || (allowed[msg.sender][_spender] == 0));

    allowed[msg.sender][_spender] = _value;
    Approval(msg.sender, _spender, _value);
    return true;
  }

  /**
   * @dev Function to check the amount of tokens that an owner allowed to a spender.
   * @param _owner address The address which owns the funds.
   * @param _spender address The address which will spend the funds.
   * @return A uint256 specifing the amount of tokens still available for the spender.
   */
  function allowance(address _owner, address _spender) constant returns (uint256 remaining) {
    return allowed[_owner][_spender];
  }

}

/**
 * @title Ownable
 * @dev The Ownable contract has an owner address, and provides basic authorization control
 * functions, this simplifies the implementation of "user permissions".
 */
contract Ownable {
    
  address public owner;
  address public candidate;



  /**
   * @dev The Ownable constructor sets the original `owner` of the contract to the sender
   * account.
   */
  function Ownable() {
    owner = msg.sender;
  }

  /**
   * @dev Throws if called by any account other than the owner.
   */
  modifier onlyOwner() {
    require(msg.sender == owner);
    _;
  }
  
  function changeOwner(address _owner) onlyOwner {
    candidate = _owner;      
  }
  
  function confirmOwner() public {
    require(candidate == msg.sender); 
    owner = candidate;   
  }
  

  /**
   * @dev Allows the current owner to transfer control of the contract to a newOwner.
   * @param newOwner The address to transfer ownership to.
   */
  /**
   *function transferOwnership(address newOwner) onlyOwner {
   * require(newOwner != address(0));      
   * owner = newOwner;
   *}
   */
}

/**
 * @title Mintable token
 * @dev Simple ERC20 Token example, with mintable token creation
 * @dev Issue: * https://github.com
 * Based on code by TokenMarketNet: https://github.com
 */

contract MintableToken is StandardToken, Ownable {
    
  event Mint(address indexed to, uint256 amount);
  
  event MintFinished();

  bool public mintingFinished = false;
  uint256 public lastTotalSupply = 0;

  address public saleAgent = 0;



  modifier canMint() {
    require(!mintingFinished);
    _;
  }


  function setSaleAgent(address newSaleAgent) public {
    require(msg.sender == saleAgent || msg.sender == owner);
    saleAgent = newSaleAgent;
  }

  /**
   * @dev Function to mint tokens
   * @param _to The address that will recieve the minted tokens.
   * @param _amount The amount of tokens to mint.
   * @return A boolean that indicates if the operation was successful.
   */
  function mint(address _to, uint256 _amount) canMint returns (bool) {
    require(msg.sender == saleAgent); 
    totalSupply = totalSupply.add(_amount);
    balances[_to] = balances[_to].add(_amount);
    Mint(_to, _amount);
    Transfer(address(0), _to, _amount);
    return true;
  }

  /**
   * @dev Function to stop minting new tokens.
   * @return True if the operation was successful.
   */

  function finishMinting() returns (bool) {
    require((msg.sender == saleAgent) || (msg.sender == owner)); 
    lastTotalSupply = totalSupply;
    mintingFinished = true;
    MintFinished();
    return mintingFinished;
  }
  function startMinting()  returns (bool) {
    require((msg.sender == saleAgent) || (msg.sender == owner)); 
    mintingFinished = false;
    return mintingFinished;
  }
  
}

contract BetOnCryptToken is MintableToken {
    
    string public constant name = "BetOnCrypt_Token";
    
    string public constant symbol = "BEC";
    
    uint32 public constant decimals = 18;
    
}

contract CrowdsalePre is Ownable {
    
    using SafeMath for uint;
    
    enum State { Active, Refunding, Close }
    
    address public multisig;

    uint public restrictedPercent;

    address public restricted;

    BetOnCryptToken public token; 
    
    uint public start;
    
    uint public period;

    uint public hardcap;

    uint public rateboc;

    uint public softcap;

    uint public minboc;

    bool is_finishmining;

    State public state;
    
    uint public first;
    uint public second;
    uint public third; 
    uint public fourth; 
    uint public fifth;

    mapping(address => uint) public balances;
    uint public indexBalance;
    
    event Closed();
    event RefundsEnabled();
    event Refunded(address indexed beneficiary, uint256 weiAmount);  


    function CrowdsalePre(address _tokencoin) {
        token = BetOnCryptToken(_tokencoin);
        is_finishmining = false;
        state = State.Active;
    }

 

    function setParams(address _multisig, address _restricted, uint _period, uint _start, uint _rateboc, uint _minboc, uint _softcap, uint _hardcap, uint _restrictedPercent, uint _first, uint _second,  uint _third, uint _fourth, uint _fifth) onlyOwner {
      multisig = _multisig; 
      restricted = _restricted;    
      start = _start;      
      period = _period;      
      minboc = _minboc.mul(1000000000000000000);      
      rateboc = _rateboc.mul(1000000000000000000);      
      softcap = _softcap.mul(1000000000000000000);      
      hardcap = _hardcap.mul(1000000000000000000);
      if (_restrictedPercent > 0 && _restrictedPercent < 50){
        restrictedPercent = _restrictedPercent;      
      } 
      else{
        restrictedPercent = 40;
      }
      first  = _first;
      second = _second;
      third  = _third; 
      fourth = _fourth; 
      fifth  = _fifth;
    }


    modifier saleIsOn() {
    	require((now > start) && (now < (start + period * 1 days)));
    	_;
    }

    modifier isUnderHardcap() {
        require(multisig.balance <= hardcap);
        _;
    }

    modifier isUnderRefunds() {
        require((this.balance < softcap) && (now > (start + period * 1 days)));
        _;
    }

    function finishMinting() onlyOwner {
        require(this.balance >= softcap);
        multisig.transfer(this.balance);
        uint issuedTokenSupply = token.totalSupply();
        uint restrictedTokens = issuedTokenSupply.mul(restrictedPercent).div(100 - restrictedPercent);
        token.mint(restricted, restrictedTokens);
        token.finishMinting();
        is_finishmining = true;
    }


    function destroyCrowdsale() onlyOwner {
      require(state == State.Close);
      selfdestruct(owner);
    }

    function closeCrowdsale() onlyOwner {
      require(state == State.Active);
      require(now > (start + (period * 1 days)));
      require(this.balance >= softcap);
      state = State.Close;
      if (is_finishmining == false){
        finishMinting();
      }
      Closed();
    }

    function enableRefunds() onlyOwner isUnderRefunds public {
      require(state == State.Active);
      state = State.Refunding;
      RefundsEnabled();
    }

    function refund() isUnderRefunds public {
      require(state == State.Refunding);
      uint value = 0;
      value = balances[msg.sender]; 
      balances[msg.sender] = 0; 
      if (indexBalance > 0) {
         indexBalance --;
      }
      if (indexBalance == 0) {
        state = State.Close;
      }
      msg.sender.transfer(value); 
      Refunded(msg.sender, value);
    }


    function createTokens() isUnderHardcap saleIsOn payable {
        require(msg.sender != address(0));
        require(state == State.Active);
        uint tokens = rateboc.mul(msg.value).div(1 ether);
        require(tokens > minboc);
        uint bonusTokens = 0;
        if(now < (start + 6 days)) {
          bonusTokens = tokens.mul(first).div(100);
        } else if(now >= (start +  6 days) && now < (start + 12 days)) {
          bonusTokens = tokens.mul(second).div(100);
        } else if(now >= (start + 12 days) && now < (start + 18 days)) {
          bonusTokens = tokens.mul(third).div(100);
        } else if(now >= (start + 18 days) && now < (start + 24 days)) {
          bonusTokens = tokens.mul(fourth).div(100);
        } else if(now >= (start + 24 days)) {
          bonusTokens = tokens.mul(fifth).div(100);
        }
        tokens += bonusTokens;
        token.mint(this, tokens);
        token.transfer(msg.sender, tokens);
        balances[msg.sender] = balances[msg.sender].add(msg.value);
        indexBalance ++;
        
        
    }

    function sendTokens(address beneficiary, uint _tokens) onlyOwner public {
      uint value = _tokens.mul(1000000000000000000);
      token.mint(this, value);
      token.transfer(beneficiary, value);
    }

    function() external payable {
      createTokens();
    }
}