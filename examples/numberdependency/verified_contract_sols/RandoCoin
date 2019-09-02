/**
 *Submitted for verification at Etherscan.io on 2018-02-04
*/

pragma solidity ^0.4.18;

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
  * @dev Substracts two numbers, throws on overflow (i.e. if subtrahend is
greater than minuend).
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
 * @title ERC20Basic
 * @dev Simpler version of ERC20 interface
 * @dev see https://github.com/ethereum/EIPs/issues/179
 */
contract ERC20Basic {
  function balanceOf(address who) public view returns (uint256);
  function transfer(address to, uint256 value) public returns (bool);
  event Transfer(address indexed from, address indexed to, uint256 value);
}

/**
 * @title ERC20 interface
 * @dev see https://github.com/ethereum/EIPs/issues/20
 */
contract ERC20 is ERC20Basic {
  function allowance(address owner, address spender) public view returns
(uint256);
  function transferFrom(address from, address to, uint256 value) public
returns (bool);
  function approve(address spender, uint256 value) public returns (bool);
  event Approval(address indexed owner, address indexed spender, uint256
value);
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
 * @dev Based on code by FirstBlood:
https://github.com/Firstbloodio/token/blob/master/smart_contract/FirstBloodToken.sol
 */
contract StandardToken is ERC20, BasicToken {

  mapping (address => mapping (address => uint256)) internal allowed;


  /**
   * @dev Transfer tokens from one address to another
   * @param _from address The address which you want to send tokens from
   * @param _to address The address which you want to transfer to
   * @param _value uint256 the amount of tokens to be transferred
   */
  function transferFrom(address _from, address _to, uint256 _value) public
returns (bool) {
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
   * @dev Approve the passed address to spend the specified amount of
tokens on behalf of msg.sender.
   *
   * Beware that changing an allowance with this method brings the risk
that someone may use both the old
   * and the new allowance by unfortunate transaction ordering. One
possible solution to mitigate this
   * race condition is to first reduce the spender's allowance to 0 and set
the desired value afterwards:
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
   * @dev Function to check the amount of tokens that an owner allowed to a
spender.
   * @param _owner address The address which owns the funds.
   * @param _spender address The address which will spend the funds.
   * @return A uint256 specifying the amount of tokens still available for
the spender.
   */
  function allowance(address _owner, address _spender) public view returns
(uint256) {
    return allowed[_owner][_spender];
  }

  /**
   * @dev Increase the amount of tokens that an owner allowed to a spender.
   *
   * approve should be called when allowed[_spender] == 0. To increment
   * allowed value is better to use this function to avoid 2 calls (and
wait until
   * the first transaction is mined)
   * From MonolithDAO Token.sol
   * @param _spender The address which will spend the funds.
   * @param _addedValue The amount of tokens to increase the allowance by.
   */
  function increaseApproval(address _spender, uint _addedValue) public
returns (bool) {
    allowed[msg.sender][_spender] =
allowed[msg.sender][_spender].add(_addedValue);
    Approval(msg.sender, _spender, allowed[msg.sender][_spender]);
    return true;
  }

  /**
   * @dev Decrease the amount of tokens that an owner allowed to a spender.
   *
   * approve should be called when allowed[_spender] == 0. To decrement
   * allowed value is better to use this function to avoid 2 calls (and
wait until
   * the first transaction is mined)
   * From MonolithDAO Token.sol
   * @param _spender The address which will spend the funds.
   * @param _subtractedValue The amount of tokens to decrease the allowance
by.
   */
  function decreaseApproval(address _spender, uint _subtractedValue) public
returns (bool) {
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

contract RandoCoin is StandardToken {
    using SafeMath for uint256;
    
    // Standard token variables
    // Initial supply is 100MM RAND
    uint256 public totalSupply = (100000000) * 1000;
    string public name = "RandoCoin";
    string public symbol = "RAND";
    uint8 public decimals = 3;
    uint BLOCK_WAIT_TIME = 30;
    uint INIT_BLOCK_WAIT = 250;
    
    // Dev variables
    address owner;
    uint public buyPrice;
    uint public sellPrice;
    uint public priceChangeBlock;
    uint public oldPriceChangeBlock;
    bool isInitialized = false;
    
    // PRICE VARIABLES -- all prices are in wei per rando
    // 1000 rando =  1 RAND
    // Prices will change randomly in the range
    // between 0.00001 and 0.01 ETH per rand
    // Which is between $0.01 and $10
    // Initial price $5 per RAND
    // That means that the first price change has a 50/50
    // chance of going up or down.
    uint public PRICE_MIN = 0.00000001 ether;
    uint public PRICE_MAX = 0.00001 ether;
    uint public PRICE_MID = 0.000005 ether;
    
    // If anyone wants to write a bot...
    event BuyPriceChanged(uint newBuyPrice);
    event SellPriceChanged(uint newSellPrice);

    function RandoCoin() public payable {
        owner = msg.sender;
        // No premining!
        // The contract holds the whole balance
        balances[this] = totalSupply;
        
        // These numbers don't matter, they will be overriden when init() is called
        // Which will kick off the contract
        priceChangeBlock = block.number + INIT_BLOCK_WAIT;
        oldPriceChangeBlock = block.number;
        buyPrice = PRICE_MID;
        sellPrice = PRICE_MID;
    }
    
    // Can only be called once
    // This kicks off the initial 1 hour timer
    // So I can time it with a social media post
    function init() public {
        require(msg.sender == owner);
        require(!isInitialized);
        
        // Initial prices in wei per rando
        buyPrice = PRICE_MID;
        sellPrice = PRICE_MID;
        
        // First time change is roughly 1 hr (250 blocks)
        // This gives more time for people to invest in the initial price
        oldPriceChangeBlock = block.number;
        priceChangeBlock = block.number + INIT_BLOCK_WAIT;
        isInitialized = true;
    }
    
    function buy() public requireNotExpired requireCooldown payable returns (uint amount){
        amount = msg.value / buyPrice;
        require(balances[this] >= amount);
        balances[msg.sender] = balances[msg.sender].add(amount);
        balances[this] = balances[this].sub(amount);
        
        Transfer(this, msg.sender, amount);
        return amount;
    }
    
    function sell(uint amount) public requireNotExpired requireCooldown returns (uint revenue){
        require(balances[msg.sender] >= amount);
        balances[this] += amount;
        balances[msg.sender] -= amount;

        revenue = amount.mul(sellPrice);
        msg.sender.transfer(revenue);
        
        Transfer(msg.sender, this, amount);
        return revenue;
    }
    
    // Change the price if possible
    // Get rewarded with 1 RAND
    function maybeChangePrice() public {
        // We actually need two block hashes, one for buy price, one for sell
        // We will use ppriceChangeBlock and priceChangeBlock + 1, so we need
        // to wait for 1 more block
        // This will create a 1 block period where you cannot buy/sell or
        // change the price, sorry!
        require(block.number > priceChangeBlock + 1);
        
        // Block is too far away to get hash, restart timer
        // Sorry, no reward here. At this point the contract
        // is probably dead anyway.
        if (block.number - priceChangeBlock > 250) {
            waitMoreTime();
            return;
        }
        
        // I know this isn't good but
        // Open challenge if a miner can break this
        sellPrice = shittyRand(0);
        buyPrice = shittyRand(1);
        
        // Set minimum prices to avoid miniscule amounts
        if (sellPrice < PRICE_MIN) {
            sellPrice = PRICE_MIN;
        }
        
        if (buyPrice < PRICE_MIN) {
            buyPrice = PRICE_MIN;
        }
        
        BuyPriceChanged(buyPrice);
        SellPriceChanged(sellPrice);

        oldPriceChangeBlock = priceChangeBlock;
        priceChangeBlock = block.number + BLOCK_WAIT_TIME;
        
        // Reward the person who refreshed priceChangeBlock 0.1 RAND
        uint reward = 100;
        if (balances[this] > reward) {
            balances[msg.sender] = balances[msg.sender].add(reward);
            balances[this] = balances[this].sub(reward);
        }
    }
    
    // You don't want someone to be able to change the price and then
    // Execute buy and sell in the same block, they could potentially
    // game the system (I think..), so freeze buying for 2 blocks after a price change.
    modifier requireCooldown() {
        // This should always be true..
        if (block.number >= oldPriceChangeBlock) {
            require(block.number - priceChangeBlock > 2);
        }
        _;
    }
    
    modifier requireNotExpired() {
        require(block.number < priceChangeBlock);
        _;
    }
    
    // Wait more time without changing the price
    // Used only when the blockhash is too far away
    // If we didn't do this, and instead picked a block within 256
    // Someone could game the system and wait to call the function 
    // until a block which gave favorable prices.
    function waitMoreTime() internal {
        priceChangeBlock = block.number + BLOCK_WAIT_TIME;
    }
    
    // Requires block to be 256 away
    function shittyRand(uint seed) public returns(uint) {
        uint randomSeed = uint(block.blockhash(priceChangeBlock + seed));
        return randomSeed % PRICE_MAX;
    }
    
    function getBlockNumber() public returns(uint) {
        return block.number;
    }

}
