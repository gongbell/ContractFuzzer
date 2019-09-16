/**
 *Submitted for verification at Etherscan.io on 2018-01-20
*/

pragma solidity ^0.4.19;

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
 * @title ERC20Basic
 * @dev Simpler version of ERC20 interface
 * @dev see https://github.com/ethereum/EIPs/issues/179
 */
contract ERC20Basic {
  uint256 public totalSupply;
  function balanceOf(address who) public constant returns (uint256);
  function transfer(address to, uint256 value) public returns (bool);
  event Transfer(address indexed from, address indexed to, uint256 value);
}

/**
 * @title ERC20 interface
 * @dev see https://github.com/ethereum/EIPs/issues/20
 */
contract ERC20 is ERC20Basic {
  function allowance(address owner, address spender) public constant returns (uint256);
  function transferFrom(address from, address to, uint256 value) public returns (bool);
  function approve(address spender, uint256 value) public returns (bool);
  event Approval(address indexed owner, address indexed spender, uint256 value);
}

contract Pausable is ERC20Basic {

    uint public constant startPreICO = 1516525200;
    uint public constant endPreICO = startPreICO + 30 days;
    
    uint public constant startICOStage1 = 1520931600;
    uint public constant endICOStage1 = startICOStage1 + 15 days;
    
    uint public constant startICOStage2 = endICOStage1;
    uint public constant endICOStage2 = startICOStage2 + 15 days;
    
    uint public constant startICOStage3 = endICOStage2;
    uint public constant endICOStage3 = startICOStage3 + 15 days;
    
    uint public constant startICOStage4 = endICOStage3;
    uint public constant endICOStage4 = startICOStage4 + 15 days;

  /**
   * @dev modifier to allow actions only when the contract IS not paused
   */
  modifier whenNotPaused() {
    require(now < startPreICO || now > endICOStage4);
    _;
  }

}

/**
 * @title Basic token
 * @dev Basic version of StandardToken, with no allowances.
 */
contract BasicToken is Pausable {
  using SafeMath for uint256;

  mapping(address => uint256) balances;

  /**
  * @dev transfer token for a specified address
  * @param _to The address to transfer to.
  * @param _value The amount to be transferred.
  */
  function transfer(address _to, uint256 _value) public whenNotPaused returns (bool) {
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
  function balanceOf(address _owner) public constant returns (uint256 balance) {
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
  function transferFrom(address _from, address _to, uint256 _value) public whenNotPaused returns (bool) {
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
  function allowance(address _owner, address _spender) public constant returns (uint256 remaining) {
    return allowed[_owner][_spender];
  }

  /**
   * approve should be called when allowed[_spender] == 0. To increment
   * allowed value is better to use this function to avoid 2 calls (and wait until
   * the first transaction is mined)
   * From MonolithDAO Token.sol
   */
  function increaseApproval (address _spender, uint _addedValue) public returns (bool success) {
    allowed[msg.sender][_spender] = allowed[msg.sender][_spender].add(_addedValue);
    Approval(msg.sender, _spender, allowed[msg.sender][_spender]);
    return true;
  }

  function decreaseApproval (address _spender, uint _subtractedValue) public returns (bool success) {
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

  /**
   * @dev Allows the current owner to transfer control of the contract to a newOwner.
   * @param newOwner The address to transfer ownership to.
   */
  function transferOwnership(address newOwner) onlyOwner {
    require(newOwner != address(0));
    owner = newOwner;
  }
  
}

contract Gelios is Ownable, StandardToken {
    using SafeMath for uint256;

    string public constant name = "Gelios Token";
    string public constant symbol = "GLS";
    uint256 public constant decimals = 18;

    uint256 public constant INITIAL_SUPPLY = 16808824 ether;
    address public tokenWallet;
    address public multiSig;

    uint256 public tokenRate = 1000; // tokens per 1 ether

    function Gelios(address _tokenWallet, address _multiSig) {
        tokenWallet = _tokenWallet;
        multiSig = _multiSig;
        totalSupply = INITIAL_SUPPLY;
        balances[_tokenWallet] = INITIAL_SUPPLY;
    }

    function () payable public {
        require(now >= startPreICO);
        buyTokens(msg.value);
    }

    function buyTokensBonus(address bonusAddress) public payable {
        require(now >= startPreICO && now < endICOStage4);
        if (bonusAddress != 0x0 && msg.sender != bonusAddress) {
            uint bonus = msg.value.mul(tokenRate).div(100).mul(5);
            if(buyTokens(msg.value)) {
               sendTokensRef(bonusAddress, bonus);
            }
        }
    }

    uint preIcoCap = 1300000 ether;
    uint icoStage1Cap = 600000 ether;
    uint icoStage2Cap = 862500 ether;
    uint icoStage3Cap = 810000 ether;
    uint icoStage4Cap = 5000000 ether;
    
    struct Stats {
        uint preICO;
        uint preICOETHRaised;
        
        uint ICOStage1;
        uint ICOStage1ETHRaised;
        
        uint ICOStage2;
        uint ICOStage2ETHRaised;
        
        uint ICOStage3;
        uint ICOStage3ETHRaised;
        
        uint ICOStage4;
        uint ICOStage4ETHRaised;
        
        uint RefBonusese;
    }
    
    event Burn(address indexed burner, uint256 value);
    
    Stats public stats;
    uint public burnAmount = preIcoCap;
    bool[] public burnStage = [true, true, true, true];

    function buyTokens(uint amount) private returns (bool){
        // PreICO - 30% 1516525200 01/21/2018 @ 9:00am (UTC) 30 days 1300000
        // Ico 1 - 20% 1520931600 03/13/2018 @ 9:00am (UTC) cap or 15 days 600000
        // ico 2 - 15% cap or 15 days  862500
        // ico 3 - 8% cap or 15 days 810000
        // ico 4 - 0% cap or 15 days 5000000
        
        uint tokens = amount.mul(tokenRate);
        if(now >= startPreICO && now < endPreICO && stats.preICO < preIcoCap) {
            tokens = tokens.add(tokens.div(100).mul(30));
            tokens = safeSend(tokens, preIcoCap.sub(stats.preICO));
            stats.preICO = stats.preICO.add(tokens);
            stats.preICOETHRaised = stats.preICOETHRaised.add(amount);
            burnAmount = burnAmount.sub(tokens);
            
            return true;
        } else if (now >= startICOStage1 && now < endICOStage1 && stats.ICOStage1 < icoStage1Cap) {
            if (burnAmount > 0 && burnStage[0]) {
                burnTokens();
                burnStage[0] = false;
                burnAmount = icoStage1Cap;
            }
            
            tokens = tokens.add(tokens.div(100).mul(20));
            tokens = safeSend(tokens, icoStage1Cap.sub(stats.ICOStage1));
            stats.ICOStage1 = stats.ICOStage1.add(tokens);
            stats.ICOStage1ETHRaised = stats.ICOStage1ETHRaised.add(amount);
            burnAmount = burnAmount.sub(tokens);

            return true;
        } else if ( now < endICOStage2 && stats.ICOStage2 < icoStage2Cap ) {
            if (burnAmount > 0 && burnStage[1]) {
                burnTokens();
                burnStage[1] = false;
                burnAmount = icoStage2Cap;
            }
            
            tokens = tokens.add(tokens.div(100).mul(15));
            tokens = safeSend(tokens, icoStage2Cap.sub(stats.ICOStage2));
            stats.ICOStage2 = stats.ICOStage2.add(tokens);
            stats.ICOStage2ETHRaised = stats.ICOStage2ETHRaised.add(amount);
            burnAmount = burnAmount.sub(tokens);
            
            return true;
        } else if ( now < endICOStage3 && stats.ICOStage3 < icoStage3Cap ) {
            if (burnAmount > 0 && burnStage[2]) {
                burnTokens();
                burnStage[2] = false;
                burnAmount = icoStage3Cap;
            }
            
            tokens = tokens.add(tokens.div(100).mul(8));
            tokens = safeSend(tokens, icoStage3Cap.sub(stats.ICOStage3));
            stats.ICOStage3 = stats.ICOStage3.add(tokens);
            stats.ICOStage3ETHRaised = stats.ICOStage3ETHRaised.add(amount);
            burnAmount = burnAmount.sub(tokens);
            
            return true;
        } else if ( now < endICOStage4 && stats.ICOStage4 < icoStage4Cap ) {
            if (burnAmount > 0 && burnStage[3]) {
                burnTokens();
                burnStage[3] = false;
                burnAmount = icoStage4Cap;
            }
            
            tokens = safeSend(tokens, icoStage4Cap.sub(stats.ICOStage4));
            stats.ICOStage4 = stats.ICOStage4.add(tokens);
            stats.ICOStage4ETHRaised = stats.ICOStage4ETHRaised.add(amount);
            burnAmount = burnAmount.sub(tokens);
            
            return true;
        } else if (now > endICOStage4 && burnAmount > 0) {
            burnTokens();
            msg.sender.transfer(msg.value);
            burnAmount = 0;
        } else {
            revert();
        }
    }
    
    /**
     * Burn tokens which are not sold on previous stage
     **/
    function burnTokens() private {
        balances[tokenWallet] = balances[tokenWallet].sub(burnAmount);
        totalSupply = totalSupply.sub(burnAmount);
        Burn(tokenWallet, burnAmount);
    }

    /**
     * Check last token on sale
     **/
    function safeSend(uint tokens, uint stageLimmit) private returns(uint) {
        if (stageLimmit < tokens) {
            uint toReturn = tokenRate.mul(tokens.sub(stageLimmit));
            sendTokens(msg.sender, stageLimmit);
            msg.sender.transfer(toReturn);
            return stageLimmit;
        } else {
            sendTokens(msg.sender, tokens);
            return tokens;
        }
    }

    /**
     * Low-level function for tokens transfer
     **/
    function sendTokens(address _to, uint tokens) private {
        balances[tokenWallet] = balances[tokenWallet].sub(tokens);
        balances[_to] += tokens;
        Transfer(tokenWallet, _to, tokens);
        multiSig.transfer(msg.value);
    }
    
    /**
     * Burn tokens which are not sold on previous stage
     **/    
    function sendTokensRef(address _to, uint tokens) private {
        balances[tokenWallet] = balances[tokenWallet].sub(tokens);
        balances[_to] += tokens;
        Transfer(tokenWallet, _to, tokens);
        stats.RefBonusese += tokens; 
    }
    
    /**
     * Update token rate manually
     **/
    function updateTokenRate(uint newRate) onlyOwner public {
        tokenRate = newRate;
    }
    
}
