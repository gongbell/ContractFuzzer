pragma solidity ^0.4.18;

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

contract Token {
    /* This is a slight change to the ERC20 base standard.
    function totalSupply() constant returns (uint256 supply);
    is replaced with:
    uint256 public totalSupply;
    This automatically creates a getter function for the totalSupply.
    This is moved to the base contract since public getter functions are not
    currently recognised as an implementation of the matching abstract
    function by the compiler.
    */
    /// total amount of tokens
    uint256 public totalSupply;

    /// @param _owner The address from which the balance will be retrieved
    /// @return The balance
    function balanceOf(address _owner) public view returns (uint256 balance);

    /// @notice send `_value` token to `_to` from `msg.sender`
    /// @param _to The address of the recipient
    /// @param _value The amount of token to be transferred
    /// @return Whether the transfer was successful or not
    function transfer(address _to, uint256 _value) public returns (bool success);

    /// @notice send `_value` token to `_to` from `_from` on the condition it is approved by `_from`
    /// @param _from The address of the sender
    /// @param _to The address of the recipient
    /// @param _value The amount of token to be transferred
    /// @return Whether the transfer was successful or not
    function transferFrom(address _from, address _to, uint256 _value) public returns (bool success);

    /// @notice `msg.sender` approves `_spender` to spend `_value` tokens
    /// @param _spender The address of the account able to transfer the tokens
    /// @param _value The amount of tokens to be approved for transfer
    /// @return Whether the approval was successful or not
    function approve(address _spender, uint256 _value) public returns (bool success);

    /// @param _owner The address of the account owning tokens
    /// @param _spender The address of the account able to transfer the tokens
    /// @return Amount of remaining tokens allowed to spent
    function allowance(address _owner, address _spender) public view returns (uint256 remaining);

    event Transfer(address indexed _from, address indexed _to, uint256 _value);
    event Approval(address indexed _owner, address indexed _spender, uint256 _value);
}

contract PrayerCoinToken is Token {

    uint256 constant MAX_UINT256 = 2**256 - 1;

    function transfer(address _to, uint256 _value) public returns (bool success) {
        //Default assumes totalSupply can't be over max (2^256 - 1).
        //If your token leaves out totalSupply and can issue more tokens as time goes on, you need to check if it doesn't wrap.
        //Replace the if with this one instead.
        //require(balances[msg.sender] >= _value && balances[_to] + _value > balances[_to]);
        require(balances[msg.sender] >= _value);
        balances[msg.sender] -= _value;
        balances[_to] += _value;
        Transfer(msg.sender, _to, _value);
        return true;
    }

    function transferFrom(address _from, address _to, uint256 _value) public returns (bool success) {
        //same as above. Replace this line with the following if you want to protect against wrapping uints.
        //require(balances[_from] >= _value && allowed[_from][msg.sender] >= _value && balances[_to] + _value > balances[_to]);
        uint256 allowance = allowed[_from][msg.sender];
        require(balances[_from] >= _value && allowance >= _value);
        balances[_to] += _value;
        balances[_from] -= _value;
        if (allowance < MAX_UINT256) {
            allowed[_from][msg.sender] -= _value;
        }
        Transfer(_from, _to, _value);
        return true;
    }

    function balanceOf(address _owner) view public returns (uint256 balance) {
        return balances[_owner];
    }

    function getBalance(address _owner) view public returns (uint256 balance) {
        return balances[_owner];
    }

    function approve(address _spender, uint256 _value) public returns (bool success) {
        allowed[msg.sender][_spender] = _value;
        Approval(msg.sender, _spender, _value);
        return true;
    }

    function allowance(address _owner, address _spender) view public returns (uint256 remaining) {
        return allowed[_owner][_spender];
    }

    mapping (address => uint256) balances;
    mapping (address => mapping (address => uint256)) allowed;
}

contract Standard {
    function balanceOf(address _owner) public constant returns (uint256);
    function transfer(address _to, uint256 _value) public returns (bool);
}

contract PrayerCoin is PrayerCoinToken {
  using SafeMath for uint256;
  address public god;

  string public name = "PrayerCoin";
  uint8 public decimals = 18;
  string public symbol = "PRAY";
  string public version = 'H1.0';  

  uint256 public publicSupply = 666666666 ether;
 
  uint private PRAY_ETH_RATIO = 6666;
  uint private PRAY_ETH_RATIO_BONUS1 = 7106;
  uint private PRAY_ETH_RATIO_BONUS2 = 11066;

  uint256 public totalDonations = 0;
  uint256 public totalPrayers = 0;

  bool private acceptingDonations = true;
  
  modifier divine {
    require(msg.sender == god);
    _;
  }

  function PrayerCoin() public { // initialize contract
    god = msg.sender;
    balances[god] = publicSupply; // god holds all of the PRAY
  } 

  function approveAndCall(address _spender, uint256 _value, bytes _extraData) public returns (bool success) {
    allowed[msg.sender][_spender] = _value;
    Approval(msg.sender, _spender, _value);

    //call the receiveApproval function on the contract you want to be notified. This crafts the function signature manually so one doesn't have to include a contract in here just for this.
    //receiveApproval(address _from, uint256 _value, address _tokenContract, bytes _extraData)
    //it is assumed that when does this that the call *should* succeed, otherwise one would use vanilla approve instead.
    require(false == _spender.call(bytes4(bytes32(keccak256("receiveApproval(address,uint256,address,bytes)"))), msg.sender, _value, this, _extraData));
    return true;
  } 

  function startDonations() public divine {
    acceptingDonations = true;
  }

  function endDonations() public divine {
    acceptingDonations = false;
  }

  function fiatSend(address _to, uint256 amt, uint256 prayRatio) public divine {
    totalDonations += amt;
    uint256 prayersIssued = amt.mul(prayRatio);
    totalPrayers += prayersIssued;
    balances[_to] += prayersIssued;
    balances[god] -= prayersIssued;

    Transfer(address(this), _to, prayersIssued);
  }
  
  function() public payable {
    require(acceptingDonations == true);
    if (msg.value == 0) { return; }

    god.transfer(msg.value);

    totalDonations += msg.value;
    
    uint256 prayersIssued = 0;

    if (totalPrayers <= (6666666 * 1 ether)) {
        if (totalPrayers <= (666666 * 1 ether)) {
            prayersIssued = msg.value.mul(PRAY_ETH_RATIO_BONUS2);
        } else {
            prayersIssued = msg.value.mul(PRAY_ETH_RATIO_BONUS1);
        }
    } else {
        prayersIssued = msg.value.mul(PRAY_ETH_RATIO);
    }

    totalPrayers += prayersIssued;
    balances[msg.sender] += prayersIssued;
    balances[god] -= prayersIssued;

    Transfer(address(this), msg.sender, prayersIssued);
  }
 
}