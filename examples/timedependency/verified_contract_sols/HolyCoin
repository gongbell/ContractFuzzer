pragma solidity ^0.4.18;

contract SafeMath {


    function safeAdd(uint256 x, uint256 y) internal returns(uint256) {
      uint256 z = x + y;
      assert((z >= x) && (z >= y));
      return z;
    }

    function safeSubtract(uint256 x, uint256 y) internal returns(uint256) {
      assert(x >= y);
      uint256 z = x - y;
      return z;
    }

    function safeMult(uint256 x, uint256 y) internal returns(uint256) {
      uint256 z = x * y;
      assert((x == 0)||(z/x == y));
      return z;
    }

}

contract Token {
    uint256 public totalSupply;
    function balanceOf(address _owner) constant returns (uint256 balance);
    function transfer(address _to, uint256 _value) returns (bool success);
    function transferFrom(address _from, address _to, uint256 _value) returns (bool success);
    function approve(address _spender, uint256 _value) returns (bool success);
    function allowance(address _owner, address _spender) constant returns (uint256 remaining);
    event Transfer(address indexed _from, address indexed _to, uint256 _value);
    event Approval(address indexed _owner, address indexed _spender, uint256 _value);
}


/*  ERC 20 token */
contract StandardToken is Token {

    function transfer(address _to, uint256 _value) returns (bool success) {
      if (balances[msg.sender] >= _value && _value > 0) {
        balances[msg.sender] -= _value;
        balances[_to] += _value;
        Transfer(msg.sender, _to, _value);
        return true;
      } else {
        return false;
      }
    }

    function transferFrom(address _from, address _to, uint256 _value) returns (bool success) {
      if (balances[_from] >= _value && allowed[_from][msg.sender] >= _value && _value > 0) {
        balances[_to] += _value;
        balances[_from] -= _value;
        allowed[_from][msg.sender] -= _value;
        Transfer(_from, _to, _value);
        return true;
      } else {
        return false;
      }
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
}

contract HolyCoin is StandardToken, SafeMath {

    string public constant name = "HolyCoin";
    string public constant symbol = "HOLY";
    uint256 public constant decimals = 18;
    string public version = "1.0";

    address public ethFundDeposit;
    address public holyFoundersFundDeposit;
    address public holyBountyFundDeposit;

    bool public isFinalized;
    uint256 public fundingStartUnixTimestamp;
    uint256 public fundingEndUnixTimestamp;
    uint256 public constant foundersFund = 2400 * (10**3) * 10**decimals; // 2.4M HolyCoins
    uint256 public constant bountyFund = 600 * (10**3) * 10**decimals; // 0.6M HolyCoins
    uint256 public constant conversionRate = 900; // 900 HolyCoins = 1 ETH

    function tokenRate() constant returns(uint) {
        return conversionRate;
    }

    uint256 public constant tokenCreationCap =  12 * (10**6) * 10**decimals; // 12M HolyCoins maximum


    // events
    event CreateHOLY(address indexed _to, uint256 _value);

    // constructor
    function HolyCoin(
        address _ethFundDeposit,
        address _holyFoundersFundDeposit,
        address _holyBountyFundDeposit,
        uint256 _fundingStartUnixTimestamp,
        uint256 _fundingEndUnixTimestamp)
    {
      isFinalized = false;
      ethFundDeposit = _ethFundDeposit;
      holyFoundersFundDeposit = _holyFoundersFundDeposit;
      holyBountyFundDeposit = _holyBountyFundDeposit;
      fundingStartUnixTimestamp = _fundingStartUnixTimestamp;
      fundingEndUnixTimestamp = _fundingEndUnixTimestamp;
      totalSupply = foundersFund + bountyFund;
      balances[holyFoundersFundDeposit] = foundersFund;
      balances[holyBountyFundDeposit] = bountyFund;
      CreateHOLY(holyFoundersFundDeposit, foundersFund);
      CreateHOLY(holyBountyFundDeposit, bountyFund);
    }


    function makeTokens() payable  {
      if (isFinalized) throw;
      if (block.timestamp < fundingStartUnixTimestamp) throw;
      if (block.timestamp > fundingEndUnixTimestamp) throw;
      if (msg.value < 100 finney || msg.value > 100 ether) throw; // 100 finney = 0.1 ether

      uint256 tokens = safeMult(msg.value, tokenRate());

      uint256 checkedSupply = safeAdd(totalSupply, tokens);

      if (tokenCreationCap < checkedSupply) throw;

      totalSupply = checkedSupply;
      balances[msg.sender] += tokens;
      CreateHOLY(msg.sender, tokens);
    }

    function() payable {
        makeTokens();
    }

    function finalize() external {
      if (isFinalized) throw;
      if (msg.sender != ethFundDeposit) throw;

      if(block.timestamp <= fundingEndUnixTimestamp && totalSupply != tokenCreationCap) throw;

      isFinalized = true;
      if(!ethFundDeposit.send(this.balance)) throw;
    }



}