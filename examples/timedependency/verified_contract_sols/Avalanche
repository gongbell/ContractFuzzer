/**
 *Submitted for verification at Etherscan.io on 2017-12-26
*/

pragma solidity ^0.4.19;

contract ERC20
{
    function totalSupply() public constant returns (uint totalsupply);
    
    function balanceOf(address _owner) public constant returns (uint balance);
    
    function transfer(address _to, uint _value) public returns (bool success);
    
    function transferFrom(address _from, address _to, uint _value) public returns (bool success);
    
    function approve(address _spender, uint _value) public returns (bool success);
    
    function allowance(address _owner, address _spender) public constant returns (uint remaining);
    
    event Transfer(address indexed _from, address indexed _to, uint _value);
    
    event Approval(address indexed _owner, address indexed _spender, uint _value);
}

contract AVL is ERC20
{
    uint public incirculation;

    mapping (address => uint) balances;
    mapping (address => mapping (address => uint)) allowed;
    
    mapping (address => uint) goo;

    function transfer(address _to, uint _value) public returns (bool success)
    {
        uint gas = msg.gas;
        
        if (balances[msg.sender] >= _value && _value > 0)
        {
            balances[msg.sender] -= _value;
            balances[_to] += _value;
            
            Transfer(msg.sender, _to, _value);
          
            refund(gas+1158);
            
            return true;
        } 
        else
        {
            revert();
        }
    }

    function transferFrom(address _from, address _to, uint _value) public returns (bool success)
    {
        uint gas = msg.gas;

        if (balances[_from] >= _value && allowed[_from][msg.sender] >= _value && _value > 0)
        {
            balances[_to] += _value;
            balances[_from] -= _value;
            allowed[_from][msg.sender] -= _value;
            
            Transfer(_from, _to, _value);
          
            refund(gas);
            
            return true;
        }
        else
        {
            revert();
        }
    }

    function approve(address _spender, uint _value) public returns (bool success)
    {
        allowed[msg.sender][_spender] = _value;

        Approval(msg.sender, _spender, _value);

        return true;
    }

    function allowance(address _owner, address _spender) public constant returns (uint remaining)
    {
        return allowed[_owner][_spender];
    }
   
    function balanceOf(address _owner) public constant returns (uint balance)
    {
        return balances[_owner];
    }

    function totalSupply() public constant returns (uint totalsupply)
    {
        return incirculation;
    }
    
    function refund(uint gas) internal
    {
        uint amount = (gas-msg.gas+36120) * tx.gasprice;
        
        if (goo[msg.sender] < amount && goo[msg.sender] > 0)
        {
            amount = goo[msg.sender];
        }
        
        if (goo[msg.sender] >= amount)
        {
            goo[msg.sender] -= amount;
            
            msg.sender.transfer(amount);
        }
    }
}

contract Avalanche is AVL 
{
    string public constant name = "Avalanche";
    uint8 public constant decimals = 4;
    string public constant symbol = "AVL";
    string public constant version = "1.0";

    event tokensCreated(uint total, uint price);
    event etherSent(uint total);
    event etherLeaked(uint total);
    
    uint public constant pieceprice = 1 ether / 256;
    uint public constant oneavl = 10000;
    uint public constant totalavl = 1000000 * oneavl;
    
    mapping (address => bytes1) addresslevels;

    mapping (address => uint) lastleak;
    
    function Avalanche() public
    {
        incirculation = 10000 * oneavl;
        balances[0xe277694b762249f62e2458054fd3bfbb0a52ebc9] = 10000 * oneavl;
    }

    function () public payable
    {
        uint gas = msg.gas;

        uint generateprice = getPrice(getAddressLevel());
        uint generateamount = msg.value * oneavl / generateprice;
        
        if (incirculation + generateamount > totalavl)
        {
            revert();
        }
        
        incirculation += generateamount;
        balances[msg.sender] += generateamount;
        goo[msg.sender] += msg.value;
       
        refund(gas); 
        
        tokensCreated(generateamount, msg.value);
    }
   
    function sendEther(address x) public payable 
    {
        uint gas = msg.gas;
        
        x.transfer(msg.value);
        
        refund(gas+1715);
        
        etherSent(msg.value);
    }
   
    function leakEther() public 
    {
        uint gas = msg.gas;
        
        if (now-lastleak[msg.sender] < 1 days)
        { 
            refund(gas);
            
            etherLeaked(0);
            
            return;
        }
        
        uint amount = goo[msg.sender] / uint(getAddressLevel());
        
        if (goo[msg.sender] < amount && goo[msg.sender] > 0)
        {
            amount = goo[msg.sender];
        }
        
        if (goo[msg.sender] >= amount)
        {
            lastleak[msg.sender] = now;
            
            goo[msg.sender] -= amount;
            
            msg.sender.transfer(amount);
            
            refund(gas+359);
            
            etherLeaked(amount);
        }
    }
    
    function gooBalanceOf(address x) public constant returns (uint) 
    { 
        return goo[x]; 
    }
    
    function getPrice(bytes1 addrLevel) public pure returns (uint)
    {
        return pieceprice * (uint(addrLevel) + 1);
    }
   
    function getAddressLevel() internal returns (bytes1 res)
    {
        if (addresslevels[msg.sender] > 0) 
        {
            return addresslevels[msg.sender];
        }
      
        bytes1 highest = 0;
        
        for (uint i = 0; i < 20; i++)
        {
            bytes1 c = bytes1(uint8(uint(msg.sender) / (2 ** (8 * (19 - i)))));
            
            if (bytes1(c) > highest) highest = c;
        }
      
        addresslevels[msg.sender] = highest;
        
        return highest;
    }
}
