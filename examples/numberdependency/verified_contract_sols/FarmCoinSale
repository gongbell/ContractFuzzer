pragma solidity ^0.4.16;

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

contract Ownable {
  address public owner;


  event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

 mapping (address => uint) public pendingWithdrawals;
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



function withdraw() {
        uint amount = pendingWithdrawals[msg.sender];
        pendingWithdrawals[msg.sender] = 0;
        msg.sender.transfer(amount);
    }

}
/**
 * @title Claimable
 * @dev Extension for the Ownable contract, where the ownership needs to be claimed.
 * This allows the new owner to accept the transfer.
 */

contract AirDrop is Ownable {

  Token token;

  event TransferredToken(address indexed to, uint256 value);
  event FailedTransfer(address indexed to, uint256 value);

  modifier whenDropIsActive() {
    assert(isActive());

    _;
  }
address public creator;
  function AirDrop () {
      address _tokenAddr = creator; //here pass address of your token
      token = Token(_tokenAddr);
  }

  function isActive() constant returns (bool) {
    return (
        tokensAvailable() > 0 // Tokens must be available to send
    );
  }
  //below function can be used when you want to send every recipeint with different number of tokens
  function sendTokens(address[] dests, uint256[] values) whenDropIsActive onlyOwner external {
    uint256 i = 0;
    while (i < dests.length) {
        uint256 toSend = values[i] ;
        sendInternally(dests[i] , toSend, values[i]);
        i++;
    }
  }

  // this function can be used when you want to send same number of tokens to all the recipients
  function sendTokensSingleValue(address[] dests, uint256 value) whenDropIsActive onlyOwner external {
    uint256 i = 0;
    uint256 toSend = value;
    while (i < dests.length) {
        sendInternally(dests[i] , toSend, value);
        i++;
    }
  }  

  function sendInternally(address recipient, uint256 tokensToSend, uint256 valueToPresent) internal {
    if(recipient == address(0)) return;

    if(tokensAvailable() >= tokensToSend) {
      token.transfer(recipient, tokensToSend);
      TransferredToken(recipient, valueToPresent);
    } else {
      FailedTransfer(recipient, valueToPresent); 
    }
  }   


  function tokensAvailable() constant returns (uint256) {
    return token.balanceOf(this);
  }

}
contract Claimable is Ownable {
    address public pendingOwner;

    /**
     * @dev Modifier throws if called by any account other than the pendingOwner.
     */
    modifier onlyPendingOwner() {
        require(msg.sender == pendingOwner);
        _;
    }

    /**
     * @dev Allows the current owner to set the pendingOwner address.
     * @param newOwner The address to transfer ownership to.
     */
    function transferOwnership(address newOwner) onlyOwner public {
        pendingOwner = newOwner;
    }

    /**
     * @dev Allows the pendingOwner address to finalize the transfer.
     */
    function claimOwnership() onlyPendingOwner public {
        OwnershipTransferred(owner, pendingOwner);
        owner = pendingOwner;
        pendingOwner = address(0);
    }


}


contract EtherToFARM is Ownable {
 using SafeMath for uint;
 using SafeMath for uint256;


uint256 public totalSupply;// total no of tokens in supply
uint remaining;
uint price;

mapping (address => uint) investors; //it maps no of FarmCoin given to each address

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

function transfer(address _to, uint256 _value) returns (bool success) {}

function ()  payable {// called when ether is send

    uint256 remaining;
    uint256 totalSupply;
    uint price;
    assert(remaining < totalSupply);
    uint FarmCoin = div(msg.value,price); // calculate no of FarmCoin to be issued depending on the price and ether send
    assert(FarmCoin < sub(totalSupply,remaining)); //FarmCoin available should be greater than the one to be issued
    add(investors[msg.sender],FarmCoin);
    remaining = add(remaining, FarmCoin);
    transfer(msg.sender, FarmCoin);
}

function setPrice(uint _price)
{ //  price need to be set maually as it cannot be done via ethereum network
    uint price;
    price = _price;
}
}

contract PayToken is EtherToFARM {
 function() public payable{
         if(msg.sender!=owner)
       giveReward(msg.sender,msg.value);
}

 function giveReward(address _payer,uint _payment) public payable returns (bool _success){
        uint tokenamount = _payment / price;
        return transfer(_payer,tokenamount);
    }     
}

contract Token is EtherToFARM, PayToken {

    /// @param _owner The address from which the balance will be retrieved
    /// @return The balance
    function balanceOf(address _owner) constant returns (uint256 balance) {}

    /// @notice send `_value` token to `_to` from `msg.sender`
    /// @param _to The address of the recipient
    /// @param _value The amount of token to be transferred
    /// @return Whether the transfer was successful or not
    function transfer(address _to, uint256 _value) returns (bool success) {}

    /// @notice send `_value` token to `_to` from `_from` on the condition it is approved by `_from`
    /// @param _from The address of the sender
    /// @param _to The address of the recipient
    /// @param _value The amount of token to be transferred
    /// @return Whether the transfer was successful or not
    function transferFrom(address _from, address _to, uint256 _value) returns (bool success) {}

    /// @notice `msg.sender` approves `_addr` to spend `_value` tokens
    /// @param _spender The address of the account able to transfer the tokens
    /// @param _value The amount of wei to be approved for transfer
    /// @return Whether the approval was successful or not
    function approve(address _spender, uint256 _value) returns (bool success) {}

    /// @param _owner The address of the account owning tokens
    /// @param _spender The address of the account able to transfer the tokens
    /// @return Amount of remaining tokens allowed to spent
    function allowance(address _owner, address _spender) constant returns (uint256 remaining) {}

    event Transfer(address indexed _from, address indexed _to, uint256 _value);
    event Approval(address indexed _owner, address indexed _spender, uint256 _value);
    
}


contract StandardToken is Token {

    function transfer(address _to, uint256 _value) returns (bool success) {
        //Default assumes totalSupply can't be over max (2^256 - 1).
        //If your token leaves out totalSupply and can issue more tokens as time goes on, you need to check if it doesn't wrap.
        //Replace the if with this one instead.
        //if (balances[msg.sender] >= _value && balances[_to] + _value > balances[_to]) {
        if (balances[msg.sender] >= _value && _value > 0) {
            balances[msg.sender] -= _value;
            balances[_to] += _value;
            Transfer(msg.sender, _to, _value);
            return true;
        } else { return false; }
    }

   
uint constant MAX_UINT = 2**256 - 1;

/// @dev ERC20 transferFrom, modified such that an allowance of MAX_UINT represents an unlimited allowance.
/// @param _from Address to transfer from.
/// @param _to Address to transfer to.
/// @param _value Amount to transfer.
/// @return Success of transfer.
function transferFrom(address _from, address _to, uint _value)
    public
    returns (bool)
{
    uint allowance = allowed[_from][msg.sender];
    require(balances[_from] >= _value
            && allowance >= _value
            && balances[_to] + _value >= balances[_to]);
    balances[_to] += _value;
    balances[_from] -= _value;
    if (allowance < MAX_UINT) {
        allowed[_from][msg.sender] -= _value;
    }
    Transfer(_from, _to, _value);
    return true;
}

    function balanceOf(address _owner) constant returns (uint256 balance) {
        return balances[_owner];
    }

    function approve(address _spender, uint256 _value) returns (bool success) {
        allowed[msg.sender][_spender] = _value;
        Approval(msg.sender, _spender, _value);
        return true;
    }

    function allowance(address _owner, address _spender) constant returns (uint256 remaining) {
      return allowed[_owner][_spender];
    }

    mapping (address => uint256) balances;
    mapping (address => mapping (address => uint256)) allowed;
    uint256 public totalSupply;
}



//name this contract whatever you'd like
contract FarmCoin is StandardToken {

   
    /* Public variables of the token */

    /*
    NOTE:
    The following variables are OPTIONAL vanities. One does not have to include them.
    They allow one to customise the token contract & in no way influences the core functionality.
    Some wallets/interfaces might not even bother to look at this information.
    */
    string public name = 'WorldFarmCoin';                   //fancy name: eg Simon Bucks
    uint8 public decimals = 0;                //How many decimals to show. ie. There could 1000 base units with 3 decimals. Meaning 0.980 SBX = 980 base units. It's like comparing 1 wei to 1 ether.
    string public symbol = 'WFARM';                 //An identifier: eg SBX
    string public version = 'H1.0';       //human 0.1 standard. Just an arbitrary versioning scheme.

//
// CHANGE THESE VALUES FOR YOUR TOKEN
//

//make sure this function name matches the contract name above. So if you're token is called TutorialToken, make sure the //contract name above is also TutorialToken instead of ERC20Token

    function FarmCoin(
        ) {
        balances[msg.sender] = 5000000;               // Give the creator all initial tokens (100000 for example)
        totalSupply = 5000000;                        // Update total supply (100000 for example)
        name = "WorldFarmCoin";                                   // Set the name for display purposes
        decimals = 0;                            // Amount of decimals for display purposes
        symbol = "WFARM";                               // Set the symbol for display purposes
    }

    /* Approves and then calls the receiving contract */
    function approveAndCall(address _spender, uint256 _value, bytes _extraData) returns (bool success) {
        allowed[msg.sender][_spender] = _value;
        Approval(msg.sender, _spender, _value);

        //call the receiveApproval function on the contract you want to be notified. This crafts the function signature manually so one doesn't have to include a contract in here just for this.
        //receiveApproval(address _from, uint256 _value, address _tokenContract, bytes _extraData)
        //it is assumed that when does this that the call *should* succeed, otherwise one would use vanilla approve instead.
        if(!_spender.call(bytes4(bytes32(sha3("receiveApproval(address,uint256,address,bytes)"))), msg.sender, _value, this, _extraData)) { revert; }
        return true;
    }
}


contract FarmCoinSale is FarmCoin {
 using SafeMath for uint256;
    uint256 public maxMintable;
    uint256 public totalMinted;
    uint256 totalTokens;
    uint256 public decimals = 0;
    uint public endBlock;
    uint public startBlock;
    uint256 public exchangeRate;
    
    uint public startTime;
    bool public isFunding;
    address public ETHWallet;
    uint256 public heldTotal;

    bool private configSet;
    address public creator;

    mapping (address => uint256) public heldTokens;
    mapping (address => uint) public heldTimeline;

    event Contribution(address from, uint256 amount);
    event ReleaseTokens(address from, uint256 amount);

// start and end dates where crowdsale is allowed (both inclusive)
  uint256 constant public START = 1517461200000; // +new Date(2018, 2, 1) / 1000
  uint256 constant public END = 1522555200000; // +new Date(2018, 4, 1) / 1000

// @return the rate in FARM per 1 ETH according to the time of the tx and the FARM pricing program.
    // @Override
  function getRate() constant returns (uint256 rate) {
    if      (now < START)            return rate = 1190476190476200; // presale, 40% bonus
    else if (now <= START +  6 days) return rate = 1234567900000000 ; // day 1 to 6, 35% bonus
    else if (now <= START + 13 days) return rate = 1282051300000000 ; // day 7 to 13, 30% bonus
    else if (now <= START + 20 days) return rate = 1333333300000000 ; // day 14 to 20, 25% bonus
    else if (now <= START + 28 days) return rate = 1388888900000000 ; // day 21 to 28, 20% bonus
    return rate = 1666666700000000; // no bonus
 
  }
  

 mapping (address => uint256) balance;
 mapping (address => mapping (address => uint256)) allowed;

 
    function buy() payable returns (bool success) {
	if (!isFunding) {return true;} 
	else {
	var buyPrice = getRate();
	buyPrice;
	uint amount = msg.value / buyPrice;                
        totalTokens += amount;                          
        balance[msg.sender] += amount;                   
        Transfer(this, msg.sender, amount); 
	return true; }            
    }

    function fund (uint256 amount) onlyOwner {
        if (!msg.sender.send(amount)) {                      		
          revert;                                         
        }           
    }

    function () payable {
    var buyPrice = getRate();
	buyPrice;
	uint amount = msg.value / buyPrice;                
        totalTokens += amount;                          
        balance[msg.sender] += amount;                   
        Transfer(this, msg.sender, amount); 
	 }               
    
    function FarmCoinSale() {
        startBlock = block.number;
        maxMintable = 5000000; // 3 million max sellable (18 decimals)
        ETHWallet = 0x3b444fC8c2C45DCa5e6610E49dC54423c5Dcd86E;
        isFunding = true;
        
        creator = msg.sender;
        createHeldCoins();
        startTime = 1517461200000;
        var buyPrice = getRate();
	    buyPrice;
        }

 
    // setup function to be ran only 1 time
    // setup token address
    // setup end Block number
    function setup(address TOKEN, uint endBlockTime) {
        require(!configSet);
        endBlock = endBlockTime;
        configSet = true;
    }

    function closeSale() external {
      require(msg.sender==creator);
      isFunding = false;
    }

    // CONTRIBUTE FUNCTION
    // converts ETH to TOKEN and sends new TOKEN to the sender
    function contribute() external payable {
        require(msg.value>0);
        require(isFunding);
        require(block.number <= endBlock);
        uint256 amount = msg.value * exchangeRate;
        uint256 total = totalMinted + amount;
        require(total<=maxMintable);
        totalMinted += total;
        ETHWallet.transfer(msg.value);
        Contribution(msg.sender, amount);
    }

    function deposit() payable {
      create(msg.sender);
    }
    function register(address sender) payable {
    }
  
    function create(address _beneficiary) payable{
    uint256 amount = msg.value;
    /// 
    }

    function withdraw() {
    require ( msg.sender == owner );
    msg.sender.transfer(this.balance);
}
    // update the ETH/COIN rate
    function updateRate(uint256 rate) external {
        require(msg.sender==creator);
        require(isFunding);
        exchangeRate = rate;
    }

    // change creator address
    function changeCreator(address _creator) external {
        require(msg.sender==creator);
        creator = _creator;
    }

    // change transfer status for WorldFarmCoin token
    function changeTransferStats(bool _allowed) external {
        require(msg.sender==creator);
     }

    // internal function that allocates a specific amount of ATYX at a specific block number.
    // only ran 1 time on initialization
    function createHeldCoins() internal {
        // TOTAL SUPPLY = 5,000,000
        createHoldToken(msg.sender, 0);
        createHoldToken(0xd9710D829fa7c36E025011b801664009E4e7c69D, 1000000);
        createHoldToken(0xd9710D829fa7c36E025011b801664009E4e7c69D, 1000000);
    }

    // function to create held tokens for developer
    function createHoldToken(address _to, uint256 amount) internal {
        heldTokens[_to] = amount;
        heldTimeline[_to] = block.number + 0;
        heldTotal += amount;
        totalMinted += heldTotal;
    }

    // function to release held tokens for developers
    function releaseHeldCoins() external {
        uint256 held = heldTokens[msg.sender];
        uint heldBlock = heldTimeline[msg.sender];
        require(!isFunding);
        require(held >= 0);
        require(block.number >= heldBlock);
        heldTokens[msg.sender] = 0;
        heldTimeline[msg.sender] = 0;
        ReleaseTokens(msg.sender, held);
    }

}