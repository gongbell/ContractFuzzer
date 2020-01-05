pragma solidity ^0.4.13;



contract ERC20Token {

    /// @return total amount of tokens
    function totalSupply() constant returns (uint256) {}

    /// @return The balance
    function balanceOf(address) constant returns (uint256) {}

    /// @notice send `_value` token to `_to` from `msg.sender`
    function transfer(address, uint256) returns (bool) {}

    /// @notice send `_value` token to `_to` from `_from` on the condition it is approved by `_from`
    function transferFrom(address, address, uint256) returns (bool) {}

    /// @notice `msg.sender` approves `_addr` to spend `_value` tokens
    function approve(address, uint256) returns (bool) {}

    function allowance(address, address) constant returns (uint256) {}

    event Transfer(address indexed _from, address indexed _to, uint256 _value);
    event Approval(address indexed _owner, address indexed _spender, uint256 _value);
}


contract StandardToken is ERC20Token {

    function transfer(address _to, uint256 _value) returns (bool success) {
        //Default assumes totalSupply can't be over max (2^256 - 1).
        require(_to != 0x0);
        if (balances[msg.sender] >= _value && _value > 0) {
            balances[msg.sender] -= _value;
            balances[_to] += _value;
            Transfer(msg.sender, _to, _value);
            return true;
        } else { return false; 
                revert();}
    }

    function transferFrom(address _from, address _to, uint256 _value) returns (bool success) {
        //same as above. Replace this line with the following if you want to protect against wrapping uints.
        //if (balances[_from] >= _value && allowed[_from][msg.sender] >= _value && balances[_to] + _value > balances[_to]) {
        if (balances[_from] >= _value && allowed[_from][msg.sender] >= _value && _value > 0) {
            balances[_to] += _value;
            balances[_from] -= _value;
            allowed[_from][msg.sender] -= _value;
            Transfer(_from, _to, _value);
            return true;
        } else { return false; }
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


contract VeganCoin is StandardToken {

    string public name;                   
    uint8 public decimals;                
    string public symbol;                 

    function VeganCoin(){

        balances[msg.sender] = 50000000000000000000000;               // Give the creator all initial tokens
        totalSupply = 50000000000000000000000;                        // Update total supply
        name = "Vegan Coin";                                   // Set the name for display purposes
        decimals = 18;                            // Amount of decimals for display purposes
        symbol = "VGN";                               // Set the symbol for display purposes

    }

    function fundContract() payable returns(bool success) {
        return true;
    }
}