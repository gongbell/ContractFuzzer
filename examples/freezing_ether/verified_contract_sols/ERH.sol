pragma solidity ^0.4.18;

////////////////////
// STANDARD TOKEN //
////////////////////

contract Token {







    uint256 public totalSupply;
    function balanceOf(address _owner) public  returns (uint256 balance);
    function transfer(address _to, uint256 _value) public returns (bool success);
    function transferFrom(address _from, address _to, uint256 _value) public returns (bool success);
    function approve(address _spender, uint256 _value) public returns (bool success);
    function allowance(address _owner, address _spender) public  returns (uint256 remaining);
    event Transfer(address indexed _from, address indexed _to, uint256 _value);
    event Approval(address indexed _owner, address indexed _spender, uint256 _value);

}

contract StandardToken is Token {







     function balanceOf(address _owner) public  returns (uint256 balance) {


        
return balances[_owner];
    }

    function transfer(address _to, uint256 _value) public returns (bool success) {


      
if (balances[msg.sender] >= _value && _value > 0) {

        
balances[msg.sender] -= _value;

        
balances[_to] += _value;

        
Transfer(msg.sender, _to, _value);

        
return true;
      } else {

        
return false;
      }
    }

    function transferFrom(address _from, address _to, uint256 _value) public returns (bool success) {


      
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

    function approve(address _spender, uint256 _value) public returns (bool success) {


        
allowed[msg.sender][_spender] = _value;

        
Approval(msg.sender, _spender, _value);

        
return true;
    }

    function allowance(address _owner, address _spender) public  returns (uint256 remaining) {


      
return allowed[_owner][_spender];
    }

    mapping (address => uint256) balances;
    mapping (address => mapping (address => uint256)) allowed;

}

////////////////////////////////////////////
//   Emergency Response Ethereum-Based   //
///////////////////////////////////////////


contract ERH is StandardToken {








    string public constant name = "Emergency Response Ethereum-Based";
    string public constant symbol = "ERH";
    uint256 public constant decimals = 18;
    uint256 public totalSupply = 10000000000 * 10**decimals;


    function ERH (address _addressFounder)  {



      
balances[_addressFounder] = totalSupply;

      
Transfer(0x0, _addressFounder, totalSupply);

    }

    function () payable public {


      

require(msg.value == 0);

    }

}