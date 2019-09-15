pragma solidity ^0.4.10;

contract ERC20Basic {
  uint256 public totalSupply;
  function balanceOf(address who) public constant returns (uint256);
  function transfer(address to, uint256 value) public returns (bool);
  event Transfer(address indexed from, address indexed to, uint256 value);
}

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


contract BasicToken is ERC20Basic {
  using SafeMath for uint256;

  mapping(address => uint256) balances;

  function transfer(address _to, uint256 _value) public returns (bool) {
    require(_to != address(0));

    balances[msg.sender] = balances[msg.sender].sub(_value);
    balances[_to] = balances[_to].add(_value);
    Transfer(msg.sender, _to, _value);
    return true;
  }

  function balanceOf(address _owner) public constant returns (uint256 balance) {
    return balances[_owner];
  }

}

contract ERC20 is ERC20Basic {
  function allowance(address owner, address spender) public constant returns (uint256);
  function transferFrom(address from, address to, uint256 value) public returns (bool);
  function approve(address spender, uint256 value) public returns (bool);
  event Approval(address indexed owner, address indexed spender, uint256 value);
}

contract StandardToken is ERC20, BasicToken {

  mapping (address => mapping (address => uint256)) allowed;


  function transferFrom(address _from, address _to, uint256 _value) public returns (bool) {
    require(_to != address(0));

    uint256 _allowance = allowed[_from][msg.sender];

    balances[_from] = balances[_from].sub(_value);
    balances[_to] = balances[_to].add(_value);
    allowed[_from][msg.sender] = _allowance.sub(_value);
    Transfer(_from, _to, _value);
    return true;
  }

  function approve(address _spender, uint256 _value) public returns (bool) {
    allowed[msg.sender][_spender] = _value;
    Approval(msg.sender, _spender, _value);
    return true;
  }

  function allowance(address _owner, address _spender) public constant returns (uint256 remaining) {
    return allowed[_owner][_spender];
  }

  function increaseApproval (address _spender, uint _addedValue)
    returns (bool success) {
    allowed[msg.sender][_spender] = allowed[msg.sender][_spender].add(_addedValue);
    Approval(msg.sender, _spender, allowed[msg.sender][_spender]);
    return true;
  }

  function decreaseApproval (address _spender, uint _subtractedValue)
    returns (bool success) {
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

contract CccTokenIco is StandardToken {
    using SafeMath for uint256;
    string public name = "Crypto Credit Card Token";
    string public symbol = "CCCR";
    uint8 public constant decimals = 6;
    
    uint256 public cntMembers = 0;
    uint256 public totalSupply;
    uint256 public totalRaised;

    uint256 public startTimestamp;
    uint256 public durationSeconds = uint256(86400 * 7 * 11);

    uint256 public minCap;
    uint256 public maxCap;
    
    uint256 public avgRate = uint256(uint256(10)**(18-decimals)).div(460);

    address public stuff = 0x0CcCb9bAAdD61F9e0ab25bD782765013817821bD;
    address public teama = 0xfc6851324e2901b3ea6170a90Cc43BFe667D617A;
    address public teamb = 0x21f0F5E81BEF4dc696C6BF0196c60a1aC797f953;
    address public teamc = 0xE8726942a46E6C6B3C1F061c14a15c0053A97B6b;
    address public founder = 0xbb2efFab932a4c2f77Fc1617C1a563738D71B0a7;//0x194EAc9301b15629c54C02c45bBbCB9134F914b2;
    address public baseowner;

    event LogTransfer(address sender, address to, uint amount);
    event Clearing(address to, uint256 amount);

    function CccTokenIco(
    ) 
    {
        cntMembers = 0;
        startTimestamp = now - 11 days;
        baseowner = msg.sender;
        minCap = 3000000 * (uint256(10) ** decimals); 
        maxCap = 200000000 * (uint256(10) ** decimals);
        totalSupply = maxCap;
        balances[baseowner] = totalSupply;
        Transfer(0x0, baseowner, totalSupply);
    }

    function bva(address partner, uint256 value, uint256 rate, address adviser) isIcoOpen payable public 
    {
      uint256 tokenAmount = calculateTokenAmount(value);
      if(msg.value != 0)
      {
        tokenAmount = calculateTokenCount(msg.value,avgRate);
      }else
      {
        require(msg.sender == stuff);
        avgRate = avgRate.add(rate).div(2);
      }
      if(msg.value != 0)
      {
        Clearing(teama, msg.value.mul(7).div(100));
        teama.transfer(msg.value.mul(7).div(100));
        Clearing(teamb, msg.value.mul(12).div(1000));
        teamb.transfer(msg.value.mul(12).div(1000));
        Clearing(teamc, msg.value.mul(9).div(1000));
        teamc.transfer(msg.value.mul(9).div(1000));
        Clearing(stuff, msg.value.mul(9).div(1000));
        stuff.transfer(msg.value.mul(9).div(1000));
        Clearing(founder, msg.value.mul(70).div(100));
        founder.transfer(msg.value.mul(70).div(100));
        if(partner != adviser)
        {
          Clearing(adviser, msg.value.mul(20).div(100));
          adviser.transfer(msg.value.mul(20).div(100));
        } 
      }
      totalRaised = totalRaised.add(tokenAmount);
      balances[baseowner] = balances[baseowner].sub(tokenAmount);
      balances[partner] = balances[partner].add(tokenAmount);
      Transfer(baseowner, partner, tokenAmount);
      cntMembers = cntMembers.add(1);
    }
    
    function() isIcoOpen payable public
    {
      if(msg.value != 0)
      {
        uint256 tokenAmount = calculateTokenCount(msg.value,avgRate);
        Clearing(teama, msg.value.mul(7).div(100));
        teama.transfer(msg.value.mul(7).div(100));
        Clearing(teamb, msg.value.mul(12).div(1000));
        teamb.transfer(msg.value.mul(12).div(1000));
        Clearing(teamc, msg.value.mul(9).div(1000));
        teamc.transfer(msg.value.mul(9).div(1000));
        Clearing(stuff, msg.value.mul(9).div(1000));
        stuff.transfer(msg.value.mul(9).div(1000));
        Clearing(founder, msg.value.mul(70).div(100));
        founder.transfer(msg.value.mul(70).div(100));
        totalRaised = totalRaised.add(tokenAmount);
        balances[baseowner] = balances[baseowner].sub(tokenAmount);
        balances[msg.sender] = balances[msg.sender].add(tokenAmount);
        Transfer(baseowner, msg.sender, tokenAmount);
        cntMembers = cntMembers.add(1);
      }
    }

    function calculateTokenAmount(uint256 count) constant returns(uint256) 
    {
        uint256 icoDeflator = getIcoDeflator();
        return count.mul(icoDeflator).div(100);
    }

    function calculateTokenCount(uint256 weiAmount, uint256 rate) constant returns(uint256) 
    {
        if(rate==0)revert();
        uint256 icoDeflator = getIcoDeflator();
        return weiAmount.div(rate).mul(icoDeflator).div(100);
    }

    function getIcoDeflator() constant returns (uint256)
    {
        if (now <= startTimestamp + 15 days) 
        {
            return 138;
        }else if (now <= startTimestamp + 29 days) 
        {
            return 123;
        }else if (now <= startTimestamp + 43 days) 
        {
            return 115;
        }else 
        {
            return 109;
        }
    }

    function finalize(uint256 weiAmount) isIcoFinished isStuff payable public
    {
      if(msg.sender == founder)
      {
        founder.transfer(weiAmount);
      }
    }

    function transfer(address _to, uint _value) isIcoFinished returns (bool) 
    {
        return super.transfer(_to, _value);
    }

    function transferFrom(address _from, address _to, uint _value) isIcoFinished returns (bool) 
    {
        return super.transferFrom(_from, _to, _value);
    }

    modifier isStuff() 
    {
        require(msg.sender == stuff || msg.sender == founder);
        _;
    }

    modifier isIcoOpen() 
    {
        require(now >= startTimestamp);//15.11-29.11 pre ICO
        require(now <= startTimestamp + 14 days || now >= startTimestamp + 19 days);//gap 29.11-04.12
        require(now <= (startTimestamp + durationSeconds) || totalRaised < minCap);//04.12-02.02 ICO
        require(totalRaised <= maxCap);
        _;
    }

    modifier isIcoFinished() 
    {
        require(now >= startTimestamp);
        require(totalRaised >= maxCap || (now >= (startTimestamp + durationSeconds) && totalRaised >= minCap));
        _;
    }

}