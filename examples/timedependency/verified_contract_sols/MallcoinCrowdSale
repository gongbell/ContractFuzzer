pragma solidity ^0.4.18;

library SafeMath {
    function mul(uint256 a, uint256 b) internal pure returns (uint256) {
      if (a == 0) {
          return 0;
      }
      uint256 c = a * b;
      assert(c / a == b);
      return c;
    }

    function div(uint256 a, uint256 b) internal pure returns (uint256) {
      uint256 c = a / b;
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

contract Ownable {
    address public owner;

    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

    function Ownable() public {
      owner = msg.sender;
    }

    modifier onlyOwner() {
      require(msg.sender == owner);
      _;
    }

    function transferOwnership(address newOwner) public onlyOwner {
      require(newOwner != address(0));

      OwnershipTransferred(owner, newOwner);
      owner = newOwner;
    }
}

contract Authorizable {
    mapping(address => bool) authorizers;

    modifier onlyAuthorized {
      require(isAuthorized(msg.sender));
      _;
    }

    function Authorizable() public {
      authorizers[msg.sender] = true;
    }


    function isAuthorized(address _addr) public constant returns(bool) {
      require(_addr != address(0));

      bool result = bool(authorizers[_addr]);
      return result;
    }

    function addAuthorized(address _addr) external onlyAuthorized {
      require(_addr != address(0));

      authorizers[_addr] = true;
    }

    function delAuthorized(address _addr) external onlyAuthorized {
      require(_addr != address(0));
      require(_addr != msg.sender);

      //authorizers[_addr] = false;
      delete authorizers[_addr];
    }
}

interface tokenRecipient { function receiveApproval(address _from, uint256 _value, address _token, bytes _extraData) external; }

contract ERC20Basic {
    function totalSupply() public view returns (uint256);
    function balanceOf(address who) public view returns (uint256);
    function transfer(address to, uint256 value) public returns (bool);
    event Transfer(address indexed from, address indexed to, uint256 value);
}

contract ERC20 is ERC20Basic {
    function allowance(address owner, address spender) public view returns (uint256);
    function transferFrom(address from, address to, uint256 value) public returns (bool);
    function approve(address spender, uint256 value) public returns (bool);
    event Approval(address indexed owner, address indexed spender, uint256 value);
}

contract BasicToken is ERC20Basic {
    using SafeMath for uint256;

    mapping(address => uint256) balances;

    uint256 totalSupply_;

    //modifier onlyPayloadSize(uint size) {
    //  require(msg.data.length < size + 4);
    //  _;
    //}

    function totalSupply() public view returns (uint256) {
      return totalSupply_;
    }

    function transfer(address _to, uint256 _value) public returns (bool) {
      //requeres in FrozenToken
      //require(_to != address(0));
      //require(_value <= balances[msg.sender]);

      balances[msg.sender] = balances[msg.sender].sub(_value);
      balances[_to] = balances[_to].add(_value);
      Transfer(msg.sender, _to, _value);
      return true;
    }

    function balanceOf(address _owner) public view returns (uint256 balance) {
      return balances[_owner];
    }
}

contract StandardToken is ERC20, BasicToken {

    mapping (address => mapping (address => uint256)) internal allowed;

    function transferFrom(address _from, address _to, uint256 _value) public returns (bool) {
      //requires in FrozenToken
      //require(_to != address(0));
      //require(_value <= balances[_from]);
      //require(_value <= allowed[_from][msg.sender]);

      balances[_from] = balances[_from].sub(_value);
      balances[_to] = balances[_to].add(_value);
      allowed[_from][msg.sender] = allowed[_from][msg.sender].sub(_value);
      Transfer(_from, _to, _value);
      return true;
    }

    function approve(address _spender, uint256 _value) public returns (bool) {
      require((_value == 0) || (allowed[msg.sender][_spender] == 0));
      allowed[msg.sender][_spender] = _value;
      Approval(msg.sender, _spender, _value);
      return true;
    }

    function allowance(address _owner, address _spender) public view returns (uint256) {
      return allowed[_owner][_spender];
    }

    function increaseApproval(address _spender, uint _addedValue) public returns (bool) {
      allowed[msg.sender][_spender] = allowed[msg.sender][_spender].add(_addedValue);
      Approval(msg.sender, _spender, allowed[msg.sender][_spender]);
      return true;
    }

    function decreaseApproval(address _spender, uint _subtractedValue) public returns (bool) {
      uint oldValue = allowed[msg.sender][_spender];
      if (_subtractedValue > oldValue) {
        allowed[msg.sender][_spender] = 0;
      } else {
        allowed[msg.sender][_spender] = oldValue.sub(_subtractedValue);
      }
      Approval(msg.sender, _spender, allowed[msg.sender][_spender]);
      return true;
    }

    function approveAndCall(address _spender, uint256 _value, bytes _extraData) public returns (bool success) {
        tokenRecipient spender = tokenRecipient(_spender);
        if (approve(_spender, _value)) {
            spender.receiveApproval(msg.sender, _value, this, _extraData);
            return true;
        }
    }
}

contract FrozenToken is StandardToken, Ownable {
    mapping(address => bool) frozens;
    mapping(address => uint256) frozenTokens;

    event FrozenAddress(address addr);
    event UnFrozenAddress(address addr);
    event FrozenTokenEvent(address addr, uint256 amount);
    event UnFrozenTokenEvent(address addr, uint256 amount);

    modifier isNotFrozen() {
      require(frozens[msg.sender] == false);
      _;
    }

    function frozenAddress(address _addr) onlyOwner public returns (bool) {
      require(_addr != address(0));

      frozens[_addr] = true;
      FrozenAddress(_addr);
      return frozens[_addr];
    }

    function unFrozenAddress(address _addr) onlyOwner public returns (bool) {
      require(_addr != address(0));

      delete frozens[_addr];
      //frozens[_addr] = false;
      UnFrozenAddress(_addr);
      return frozens[_addr];
    }

    function isFrozenByAddress(address _addr) public constant returns(bool) {
      require(_addr != address(0));

      bool result = bool(frozens[_addr]);
      return result;
    }

    function balanceFrozenTokens(address _addr) public constant returns(uint256) {
      require(_addr != address(0));

      uint256 result = uint256(frozenTokens[_addr]);
      return result;
    }

    function balanceAvailableTokens(address _addr) public constant returns(uint256) {
      require(_addr != address(0));

      uint256 frozen = uint256(frozenTokens[_addr]);
      uint256 balance = uint256(balances[_addr]);
      require(balance >= frozen);

      uint256 result = balance.sub(frozen);

      return result;
    }

    function frozenToken(address _addr, uint256 _amount) onlyOwner public returns(bool) {
      require(_addr != address(0));
      require(_amount > 0);

      uint256 balance = uint256(balances[_addr]);
      require(balance >= _amount);

      frozenTokens[_addr] = frozenTokens[_addr].add(_amount);
      FrozenTokenEvent(_addr, _amount);
      return true;
    }
    

    function unFrozenToken(address _addr, uint256 _amount) onlyOwner public returns(bool) {
      require(_addr != address(0));
      require(_amount > 0);
      require(frozenTokens[_addr] >= _amount);

      frozenTokens[_addr] = frozenTokens[_addr].sub(_amount);
      UnFrozenTokenEvent(_addr, _amount);
      return true;
    }

    function transfer(address _to, uint256 _value) isNotFrozen() public returns (bool) {
      require(_to != address(0));
      require(_value <= balances[msg.sender]);

      uint256 balance = balances[msg.sender];
      uint256 frozen = frozenTokens[msg.sender];
      uint256 availableBalance = balance.sub(frozen);
      require(availableBalance >= _value);

      return super.transfer(_to, _value);
    }

    function transferFrom(address _from, address _to, uint256 _value) isNotFrozen() public returns (bool) {
      require(_to != address(0));
      require(_value <= balances[_from]);
      require(_value <= allowed[_from][msg.sender]);

      uint256 balance = balances[_from];
      uint256 frozen = frozenTokens[_from];
      uint256 availableBalance = balance.sub(frozen);
      require(availableBalance >= _value);

      return super.transferFrom(_from ,_to, _value);
    }
}

contract MallcoinToken is FrozenToken, Authorizable {
      string public constant name = "Mallcoin Token";
      string public constant symbol = "MLC";
      uint8 public constant decimals = 18;
      uint256 public MAX_TOKEN_SUPPLY = 250000000 * 1 ether;

      event CreateToken(address indexed to, uint256 amount);
      event CreateTokenByAtes(address indexed to, uint256 amount, string data);

      modifier onlyOwnerOrAuthorized {
        require(msg.sender == owner || isAuthorized(msg.sender));
        _;
      }

      function createToken(address _to, uint256 _amount) onlyOwnerOrAuthorized public returns (bool) {
        require(_to != address(0));
        require(_amount > 0);
        require(MAX_TOKEN_SUPPLY >= totalSupply_ + _amount);

        totalSupply_ = totalSupply_.add(_amount);
        balances[_to] = balances[_to].add(_amount);

        // KYC
        frozens[_to] = true;
        FrozenAddress(_to);

        CreateToken(_to, _amount);
        Transfer(address(0), _to, _amount);
        return true;
      }

      function createTokenByAtes(address _to, uint256 _amount, string _data) onlyOwnerOrAuthorized public returns (bool) {
        require(_to != address(0));
        require(_amount > 0);
        require(bytes(_data).length > 0);
        require(MAX_TOKEN_SUPPLY >= totalSupply_ + _amount);

        totalSupply_ = totalSupply_.add(_amount);
        balances[_to] = balances[_to].add(_amount);

        // KYC
        frozens[_to] = true;
        FrozenAddress(_to);

        CreateTokenByAtes(_to, _amount, _data);
        Transfer(address(0), _to, _amount);
        return true;
      }
} 

contract MallcoinCrowdSale is Ownable, Authorizable {
      using SafeMath for uint256;

      MallcoinToken public token;
      address public wallet; 

      uint256 public PRE_ICO_START_TIME = 1519297200; // Thursday, 22 February 2018, 11:00:00 GMT
      uint256 public PRE_ICO_END_TIME = 1520550000; // Thursday, 8 March 2018, 23:00:00 GMT
      uint256 public PRE_ICO_BONUS_TIME_1 =  1519556400; // Sunday, 25 February 2018, 11:00:00 GMT
      uint256 public PRE_ICO_BONUS_TIME_2 =  1519988400; // Friday, 2 March 2018, 11:00:00 GMT
      uint256 public PRE_ICO_BONUS_TIME_3 =  1520334000; // Tuesday, 6 March 2018, 11:00:00 GMT
      uint256 public PRE_ICO_RATE = 3000 * 1 ether; // 1 Ether = 3000 MLC
      uint256 public PRE_ICO_BONUS_RATE = 75 * 1 ether; // 75 MLC = 2.5%
      uint256 public preIcoTokenSales;

      uint256 public ICO_START_TIME = 1521716400; // Thursday, 22 March 2018, 11:00:00 GMT
      uint256 public ICO_END_TIME = 1523574000; // Thursday, 12 April 2018, 23:00:00 GMT
      uint256 public ICO_BONUS_TIME_1 = 1521975600; // Sunday, 25 March 2018, 11:00:00 GMT
      uint256 public ICO_BONUS_TIME_2 = 1522839600; // Wednesday, 4 April 2018, 11:00:00 GMT
      uint256 public ICO_BONUS_TIME_3 = 1523358000; // Tuesday, 10 April 2018, 11:00:00 GMT
      uint256 public ICO_RATE = 2000 * 1 ether; // 1 Ether = 2000 MLC
      uint256 public ICO_BONUS_RATE = 50 * 1 ether; // 50 MLC = 2.5%
      uint256 public icoTokenSales;

      uint256 public SECRET_BONUS_FACTOR = 0;

      bool public crowdSaleStop = false;

      uint256 public MAX_TOKEN_SUPPLY = 250000000 * 1 ether;
      uint256 public MAX_CROWD_SALE_TOKENS = 185000000 * 1 ether;
      uint256 public weiRaised;
      uint256 public tokenSales;
      uint256 public bountyTokenFund;
      uint256 public reserveTokenFund;
      uint256 public teamTokenFund;


      event TokenPurchase(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount);
      event ChangeCrowdSaleDate(uint8 period, uint256 unixtime);

      modifier onlyOwnerOrAuthorized {
        require(msg.sender == owner || isAuthorized(msg.sender));
        _;
      }

      function MallcoinCrowdSale() public {
        wallet = owner;
        preIcoTokenSales = 0;
        icoTokenSales = 0;
        weiRaised = 0;
        tokenSales = 0;

        bountyTokenFund = 0;
        reserveTokenFund = 0;
        teamTokenFund = 0;
      
      }

   function () external payable {
     buyTokens(msg.sender);
   }

    function buyTokens(address beneficiary) public payable {
      require(beneficiary != address(0));
      require(validPurchase());

      uint256 weiAmount = msg.value;
      uint256 _buyTokens = 0;
      uint256 rate = 0;
      if (now >= PRE_ICO_START_TIME && now <= PRE_ICO_END_TIME) {
        rate = PRE_ICO_RATE.add(getPreIcoBonusRate());
        _buyTokens = rate.mul(weiAmount).div(1 ether);
        preIcoTokenSales = preIcoTokenSales.add(_buyTokens);
      } else if (now >= ICO_START_TIME && now <= ICO_END_TIME) {
        rate = ICO_RATE.add(getIcoBonusRate());
        _buyTokens = rate.mul(weiAmount).div(1 ether);
        icoTokenSales = icoTokenSales.add(_buyTokens);
      }

      require(MAX_CROWD_SALE_TOKENS >= tokenSales.add(_buyTokens));

      tokenSales = tokenSales.add(_buyTokens);
      weiRaised = weiRaised.add(weiAmount);
      wallet.transfer(msg.value);
      token.createToken(beneficiary, _buyTokens);
      TokenPurchase(msg.sender, beneficiary, weiAmount, _buyTokens);
    }

    function buyTokensByAtes(address addr_, uint256 amount_, string data_) onlyOwnerOrAuthorized  public returns (bool) {
      require(addr_ != address(0));
      require(amount_ > 0);
      require(bytes(data_).length > 0);
      require(validPurchase());

      uint256 _buyTokens = 0;
      uint256 rate = 0;
      if (now >= PRE_ICO_START_TIME && now <= PRE_ICO_END_TIME) {
        rate = PRE_ICO_RATE.add(getPreIcoBonusRate());
        _buyTokens = rate.mul(amount_).div(1

ether);
        preIcoTokenSales = preIcoTokenSales.add(_buyTokens);
      } else if (now >= ICO_START_TIME && now <= ICO_END_TIME) {
        rate = ICO_RATE.add(getIcoBonusRate());
        _buyTokens = rate.mul(amount_).div(1 ether);
        icoTokenSales = icoTokenSales.add(_buyTokens);
      }

      require(MAX_CROWD_SALE_TOKENS >= tokenSales.add(_buyTokens));

      tokenSales = tokenSales.add(_buyTokens);
      weiRaised = weiRaised.add(amount_);
      token.createTokenByAtes(addr_, _buyTokens, data_);
      TokenPurchase(msg.sender, addr_, amount_, _buyTokens);

      return true;
    }

    function getPreIcoBonusRate() private view returns (uint256 bonus) {
      bonus = 0;
      uint256 factorBonus = getFactorBonus();

      if (factorBonus > 0) {
        if (now >= PRE_ICO_START_TIME && now < PRE_ICO_BONUS_TIME_1) { // Sunday, 25 February 2018, 11:00:00 GMT
          factorBonus = factorBonus.add(7);
          bonus = PRE_ICO_BONUS_RATE.mul(factorBonus); // add 600-750 MLC
        } else if (now >= PRE_ICO_BONUS_TIME_1 && now < PRE_ICO_BONUS_TIME_2) { // Friday, 2 March 2018, 11:00:00 GMT
          factorBonus = factorBonus.add(5);
          bonus = PRE_ICO_BONUS_RATE.mul(factorBonus); // add 450-600 MLC
        } else if (now >= PRE_ICO_BONUS_TIME_2 && now < PRE_ICO_BONUS_TIME_3) { // Tuesday, 6 March 2018, 11:00:00 GMT
          factorBonus = factorBonus.add(1);
          bonus = PRE_ICO_BONUS_RATE.mul(factorBonus); // add 150-300 MLC
        } 
      }

      return bonus;
    }

    function getIcoBonusRate() private view returns (uint256 bonus) {
      bonus = 0;
      uint256 factorBonus = getFactorBonus();

      if (factorBonus > 0) {
        if (now >= ICO_START_TIME && now < ICO_BONUS_TIME_1) { // Sunday, 25 March 2018, 11:00:00 GMT
          factorBonus = factorBonus.add(7);
          bonus = ICO_BONUS_RATE.mul(factorBonus); // add 400-500 MLC
        } else if (now >= ICO_BONUS_TIME_1 && now < ICO_BONUS_TIME_2) { // Wednesday, 4 April 2018, 11:00:00 GMT
          factorBonus = factorBonus.add(5);
          bonus = ICO_BONUS_RATE.mul(factorBonus); // add 300-400 MLC
        } else if (now >= ICO_BONUS_TIME_2 && now < ICO_BONUS_TIME_3) { // Tuesday, 10 April 2018, 11:00:00 GMT
          factorBonus = factorBonus.add(1);
          bonus = ICO_BONUS_RATE.mul(factorBonus); // add 100-200 MLC
        } else if (now >= ICO_BONUS_TIME_3 && now < ICO_END_TIME) { // Secret bonus dates
          factorBonus = factorBonus.add(SECRET_BONUS_FACTOR);
          bonus = ICO_BONUS_RATE.mul(factorBonus); // add 150-300 MLC
        } 
      }

      return bonus;
    }

    function getFactorBonus() private view returns (uint256 factor) {
      factor = 0;
      if (msg.value >= 5 ether && msg.value < 10 ether) {
        factor = 1;
      } else if (msg.value >= 10 ether && msg.value < 100 ether) {
        factor = 2;
      } else if (msg.value >= 100 ether) {
        factor = 3;
      }
      return factor;
    }

   function validPurchase() internal view returns (bool) {
      bool withinPeriod = false;
     if (now >= PRE_ICO_START_TIME && now <= PRE_ICO_END_TIME && !crowdSaleStop) {
        withinPeriod = true;
      } else if (now >= ICO_START_TIME && now <= ICO_END_TIME && !crowdSaleStop) {
        withinPeriod = true;
      }
     bool nonZeroPurchase = msg.value > 0;
      
     return withinPeriod && nonZeroPurchase;
   }

    function stopCrowdSale() onlyOwner public {
      crowdSaleStop = true;
    }

    function startCrowdSale() onlyOwner public {
      crowdSaleStop = false;
    }

    function changeCrowdSaleDates(uint8 _period, uint256 _unixTime) onlyOwner public {
      require(_period > 0 && _unixTime > 0);

      if (_period == 1) {
        PRE_ICO_START_TIME = _unixTime;
        ChangeCrowdSaleDate(_period, _unixTime);
      } else if (_period == 2) {
        PRE_ICO_END_TIME = _unixTime;
        ChangeCrowdSaleDate(_period, _unixTime);
      } else if (_period == 3) {
        PRE_ICO_BONUS_TIME_1 = _unixTime;
        ChangeCrowdSaleDate(_period, _unixTime);
      } else if (_period == 4) {

PRE_ICO_BONUS_TIME_2 = _unixTime;
        ChangeCrowdSaleDate(_period, _unixTime);
      } else if (_period == 5) {
        PRE_ICO_BONUS_TIME_3 = _unixTime;
        ChangeCrowdSaleDate(_period, _unixTime);
      } else if (_period == 6) {
        ICO_START_TIME = _unixTime;
        ChangeCrowdSaleDate(_period, _unixTime);
      } else if (_period == 7) {
        ICO_END_TIME = _unixTime;
        ChangeCrowdSaleDate(_period, _unixTime);
      } else if (_period == 8) {
        ICO_BONUS_TIME_1 = _unixTime;
        ChangeCrowdSaleDate(_period, _unixTime);
      } else if (_period == 9) {
        ICO_BONUS_TIME_2 = _unixTime;
        ChangeCrowdSaleDate(_period, _unixTime);
      } else if (_period == 10) {
        ICO_BONUS_TIME_3 = _unixTime;
        ChangeCrowdSaleDate(_period, _unixTime);
      } 
    }

    function setSecretBonusFactor(uint256 _factor) onlyOwner public {
      require(_factor >= 0);

      SECRET_BONUS_FACTOR = _factor;
    }
    
    function changeMallcoinTokenAddress(address _token) onlyOwner public {
      require(_token != address(0));

      token = MallcoinToken(_token);
    }

    function finishCrowdSale() onlyOwner public returns (bool) {
      crowdSaleStop = true;
      teamTokenFund = tokenSales.div(100).mul(10); // Team fund 10%
      bountyTokenFund = tokenSales.div(100).mul(7); // Bounty fund 7%;
      reserveTokenFund = tokenSales.div(100).mul(9); // Reserve fund 9%;

      uint256 tokensFund = teamTokenFund.add(bountyTokenFund).add(reserveTokenFund);
      wallet.transfer(this.balance);
      token.createToken(wallet, tokensFund);

      return true;
    }
}