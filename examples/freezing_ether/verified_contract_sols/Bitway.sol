pragma solidity ^0.4.18;


    contract ERC20 {
    function totalSupply() public constant returns (uint256);
    function balanceOf(address who) public view returns (uint256);
    function transfer(address to, uint256 value) public returns (bool);
    function transferFrom(address from, address to, uint256 value) public returns (bool);
    function allowance(address owner, address spender) public view returns (uint256);
    function approve(address spender, uint256 value) public returns (bool);
    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);
    }

    library SafeMath {
    function mul(uint256 a, uint256 b) internal pure returns (uint256) {
    if (a == 0) {
      return 0;
    }
    uint256 c = a * b;
    assert(c / a == b);
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


    contract Bitway is ERC20 {
    
    using SafeMath for uint256;
    
    
    uint256 public totalSupply = 0;
    uint256 public maxSupply = 22000000 * 10 ** uint256(decimals);
    
    string public constant symbol = "BTW";
    string public constant name = "Bitway";
    uint256 public constant decimals = 18;
    
    
    
    uint256 public constant RATE = 10000;
    address public owner;
    
   
    mapping(address => uint256) balances;
    mapping(address => mapping(address => uint256)) allowed;
    
    
    
    function () public payable {
        createTokens();
        
    }
    
    function Bitway() public {
        owner = msg.sender;
        
    }
    
   
    function createTokens() public payable {
        require(msg.value > 0);
        require(totalSupply < maxSupply);
        uint256 tokens = msg.value.mul(RATE);
        balances[msg.sender] = balances[msg.sender].add(tokens);
        totalSupply = totalSupply.add(tokens);
        
    }
    
    
    
    function totalSupply() public constant returns (uint256){
        return totalSupply;
    }

  
    function balanceOf(address _owner) public view returns (uint256 balance) {
        return balances[_owner];
    }
    
    function transfer(address _to, uint256 _value) public returns (bool) {
        require(_to != address(0));
        require(_value <= balances[msg.sender]);

        // SafeMath.sub will throw if there is not enough balance.
        balances[msg.sender] = balances[msg.sender].sub(_value);
        balances[_to] = balances[_to].add(_value);
        Transfer(msg.sender, _to, _value);
        return true;
    }
    
    function transferFrom(address _from, address _to, uint256 _value) public returns (bool) {
        require(
        allowed[_from][msg.sender] >= _value
        && balances[_from] >= _value
        && _value > 0
        );

        balances[_from] = balances[_from].sub(_value);
        balances[_to] = balances[_to].add(_value);
        allowed[_from][msg.sender] = allowed[_from][msg.sender].sub(_value);
        Transfer(_from, _to, _value);
        return true;
    }
    
    function approve(address _spender, uint256 _value) public returns (bool) {
        allowed[msg.sender][_spender] = _value;
        Approval(msg.sender, _spender, _value);
        return true;
    }
   
    function allowance(address _owner, address _spender) public view returns (uint256) {
        return allowed[_owner][_spender];
    }
    
    event Transfer(address indexed _from, address indexed _to, uint256 _value);

    event Approval(address indexed _owner, address indexed _spender, uint256 _value);
    
}