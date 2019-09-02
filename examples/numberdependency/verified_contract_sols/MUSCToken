pragma solidity ^0.4.19; 




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


/* taking ideas from FirstBlood token */
contract SafeMath {

    /* function assert(bool assertion) internal { */
    /*   if (!assertion) { */
    /*     throw; */
    /*   } */
    /* }      // assert no longer needed once solidity is on 0.4.10 */

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

contract MUSCToken is StandardToken, SafeMath {

    // metadata
    string public constant name = "Manchester United SC";
    string public constant symbol = "MUSC";
    uint256 public constant decimals = 18;
    string public version = "1.0";

    // contracts
    address public ethFundDeposit;      // deposit address for the raised ETH 
    address public muscFundDeposit;      // deposit address for the MUSC coins fund

    // crowdsale parameters
    bool public isFinalized;              // switched to true in operational state
    uint256 public fundingStartBlock;
    uint256 public fundingEndBlock;
    uint256 public constant muscFund = 30 * (10**6) * 10**decimals;   // 30MM coins goes to the muscFund 
    uint256 public constant tokenExchangeRate = 2000; // 2000 MUSC tokens per 1 ETH == 0.0005 eth per token
    uint256 public constant tokenCreationCap =  100 * (10**6) * 10**decimals; // 100m cap total hard cap
    uint256 public constant tokenCreationMin =  0 * (10**6) * 10**decimals; // 0 is the minimum cap 


    // events
    event LogRefund(address indexed _to, uint256 _value);
    event CreateMUSC(address indexed _to, uint256 _value);

    // constructor
    function MUSCToken(
        
        )
    {
      isFinalized = false ;                   //controls pre through crowdsale state
      ethFundDeposit = 0xeEad6BE557441c568A3984eC15B0fDCC85C3e008 ;
      muscFundDeposit = 0xEaBd227E940a9e876C604eF4CEb46DDF577c5977 ;
      fundingStartBlock = 5073730 ;
      fundingEndBlock = 5240572 ;
      totalSupply = muscFund;
      balances[muscFundDeposit] = muscFund;    // Deposit the reserverd supply initially to the musc fund
      CreateMUSC(muscFundDeposit, muscFund);  // logs 
    }

    /// @dev Accepts ether and creates new MUSC tokens.
    function createTokens() payable external {
      // if (isFinalized) throw;
      require (!isFinalized); 
      // if (block.number < fundingStartBlock) throw;
      require(block.number > fundingStartBlock) ; 
      // if (block.number > fundingEndBlock) throw;
      require(block.number < fundingEndBlock) ; 
      // if (msg.value == 0) throw;
      require(msg.value != 0) ;


      uint256 tokens = safeMult(msg.value, tokenExchangeRate); // check that we're not over totals
      uint256 checkedSupply = safeAdd(totalSupply, tokens);

      // return money if something goes wrong
      // if (tokenCreationCap < checkedSupply) throw;  // odd fractions won't be found
      require(tokenCreationCap > checkedSupply) ; 

      totalSupply = checkedSupply;
      balances[msg.sender] += tokens;  // safeAdd not needed; bad semantics to use here
      CreateMUSC(msg.sender, tokens);  // logs token creation
    }

    /// @dev Ends the funding period and sends the ETH home
    function finalize() external {
      // if (isFinalized) throw;
      require(!isFinalized) ; 
      
      // if (msg.sender != ethFundDeposit) throw; // locks finalize to the ultimate ETH owner
      require(msg.sender == ethFundDeposit) ; 
      
      // if(totalSupply < tokenCreationMin) throw;      // have to sell minimum to move to operational
      require(totalSupply > tokenCreationMin) ; 
      
      // if(block.number <= fundingEndBlock && totalSupply != tokenCreationCap) throw;
      // require(block.number > fundingEndBlock ) ; 
      // move to operational
      isFinalized = true;
      
      // if(!ethFundDeposit.send(this.balance)) throw;  // send the eth to the founding team address
      require(ethFundDeposit.send(this.balance)) ; 
    }

    /// @dev Allows contributors to recover their ether in the case of a failed funding campaign.
    function refund() external {
      // if(isFinalized) throw;                       // prevents refund if operational
      require(!isFinalized) ; 
      // if (block.number <= fundingEndBlock) throw; // prevents refund until sale period is over
      require(block.number > fundingEndBlock) ; 
      // if(totalSupply >= tokenCreationMin) throw;  // no refunds if we sold enough
      require(totalSupply < tokenCreationMin) ; 
      // if(msg.sender == muscFundDeposit) throw;    // MUSC fund not entitled ever to a refund
      require(msg.sender != muscFundDeposit) ; 
      uint256 muscVal = balances[msg.sender];
      // if (muscVal == 0) throw;
      require (muscVal != 0) ; 

      balances[msg.sender] = 0;
      totalSupply = safeSubtract(totalSupply, muscVal); // extra safe
      uint256 ethVal = muscVal / tokenExchangeRate;     // should be safe; previous throws covers edges
      LogRefund(msg.sender, ethVal);               // log it 
      // if (!msg.sender.send(ethVal)) throw;       // if you're using a contract; make sure it works with .send gas limits
      require(msg.sender.send(ethVal)) ; 
    }

}