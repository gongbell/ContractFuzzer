pragma solidity ^0.4.18;

/**
 * CoinCrowd ICO. More info www.coincrowd.it 
 */

/**
 * @title SafeMath
 * @dev Math operations with safety checks that throw on error
 */
library SafeMath {
  function mul(uint256 a, uint256 b) internal pure returns (uint256) {
    uint256 c = a * b;
    assert(a == 0 || c / a == b);
    return c;
  }

  function div(uint256 a, uint256 b) internal pure returns (uint256) {
    // assert(b > 0); // Solidity automatically throws when dividing by 0
    uint256 c = a / b;
    // assert(a == b * c + a % b); // There is no case in which this doesn't hold
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


/**
 * @title Ownable
 * @dev The Ownable contract has an owner address, and provides basic authorization control
 * functions, this simplifies the implementation of "user permissions".
 */
contract Ownable {
  address public owner;
  
  event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

  /**
   * @dev The Ownable constructor sets the original `owner` of the contract to the sender
   * account.
   */
  function Ownable() internal {
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
  function transferOwnership(address newOwner) onlyOwner public {
    require(newOwner != address(0));
    OwnershipTransferred(owner, newOwner);
    owner = newOwner;
  }
}

contract tokenInterface {
	function balanceOf(address _owner) public constant returns (uint256 balance);
	function transfer(address _to, uint256 _value) public returns (bool);
}

contract Ambassador {
    using SafeMath for uint256;
    CoinCrowdICO icoContract;
    uint256 public startRC;
    uint256 public endRC;
    address internal contractOwner; 
    
    uint256 public soldTokensWithoutBonus; // wei of XCC sold token without bonuses
	function euroRaisedRc() public view returns(uint256 euro) {
        return icoContract.euroRaised(soldTokensWithoutBonus);
    }
    
    uint256[] public euroThreshold; // array of euro(k) threshold reached - 100K = 100.000€
    uint256[] public bonusThreshold; // array of bonus of each euroThreshold reached - 20% = 2000
    
    mapping(address => uint256) public balanceUser; // address => token amount

    function Ambassador(address _icoContract, address _ambassadorAddr, uint256[] _euroThreshold, uint256[] _bonusThreshold, uint256 _startRC , uint256 _endRC ) public {
        require ( _icoContract != 0 );
        require ( _ambassadorAddr != 0 );
        require ( _euroThreshold.length != 0 );
        require ( _euroThreshold.length == _bonusThreshold.length );
        
        icoContract = CoinCrowdICO(_icoContract);
        contractOwner = _icoContract;
        
        icoContract.addMeByRC(_ambassadorAddr);
        
        bonusThreshold = _bonusThreshold;
        euroThreshold = _euroThreshold;
        
        soldTokensWithoutBonus = 0;
        
        setTimeRC( _startRC, _endRC );
    }
    
    modifier onlyIcoContract() {
        require(msg.sender == contractOwner);
        _;
    }
    
    function setTimeRC(uint256 _startRC, uint256 _endRC ) internal {
        if( _startRC == 0 ) {
            startRC = icoContract.startTime();
        } else {
            startRC = _startRC;
        }
        if( _endRC == 0 ) {
            endRC = icoContract.endTime();
        } else {
            endRC = _endRC;
        }
    }
    
    function updateTime(uint256 _newStart, uint256 _newEnd) public onlyIcoContract {
        if ( _newStart != 0 ) startRC = _newStart;
        if ( _newEnd != 0 ) endRC = _newEnd;
    }

    function () public payable {
        require( now > startRC );
        if( now < endRC ) {
            uint256 tokenAmount = icoContract.buy.value(msg.value)(msg.sender);
            balanceUser[msg.sender] = balanceUser[msg.sender].add(tokenAmount);
            soldTokensWithoutBonus = soldTokensWithoutBonus.add(tokenAmount);
        } else { //claim premium bonus logic
            require( balanceUser[msg.sender] > 0 );
            uint256 bonusApplied = 0;
            for (uint i = 0; i < euroThreshold.length; i++) {
                if ( icoContract.euroRaised(soldTokensWithoutBonus).div(1000) > euroThreshold[i] ) {
                    bonusApplied = bonusThreshold[i];
                }
            }    
            require( bonusApplied > 0 );
            
            uint256 addTokenAmount = balanceUser[msg.sender].mul( bonusApplied ).div(10**2);
            balanceUser[msg.sender] = 0; 
            
            icoContract.claimPremium(msg.sender, addTokenAmount);
            if( msg.value > 0 ) msg.sender.transfer(msg.value); // give back eth 
        }
    }
}

contract CoinCrowdICO is Ownable {
    using SafeMath for uint256;
    tokenInterface public tokenContract;
    
	uint256 public decimals = 18;
    uint256 public tokenValue;  // 1 XCC in wei
    uint256 public constant centToken = 20; // euro cents value of 1 token 
    
    function euroRaised(uint256 _weiTokens) public view returns (uint256) { // convertion of sold token in euro raised in wei
        return _weiTokens.mul(centToken).div(100).div(10**decimals);
    }
    
    uint256 public endTime;  // seconds from 1970-01-01T00:00:00Z
    uint256 public startTime;  // seconds from 1970-01-01T00:00:00Z
    uint256 internal constant weekInSeconds = 604800; // seconds in a week
    
    uint256 public totalSoldTokensWithBonus; // total wei of XCC distribuited from this ICO
    uint256 public totalSoldTokensWithoutBonus; // total wei of XCC distribuited from this ICO without bonus
	function euroRaisedICO() public view returns(uint256 euro) {
        return euroRaised(totalSoldTokensWithoutBonus);
    }
	
    uint256 public remainingTokens; // total wei of XCC remaining (without bonuses)

    mapping(address => address) public ambassadorAddressOf; // ambassadorContract => ambassadorAddress


    function CoinCrowdICO(address _tokenAddress, uint256 _tokenValue, uint256 _startTime) public {
        tokenContract = tokenInterface(_tokenAddress);
        tokenValue = _tokenValue;
        startICO(_startTime); 
        totalSoldTokensWithBonus = 0;
        totalSoldTokensWithoutBonus = 0;
        remainingTokens = 24500000  * 10 ** decimals; // 24.500.000 * 0.20€ = 4.900.000€ CAPPED
    }

    address public updater;  // account in charge of updating the token value
    event UpdateValue(uint256 newValue);

    function updateValue(uint256 newValue) public {
        require(msg.sender == updater || msg.sender == owner);
        tokenValue = newValue;
        UpdateValue(newValue);
    }

    function updateUpdater(address newUpdater) public onlyOwner {
        updater = newUpdater;
    }

    function updateTime(uint256 _newStart, uint256 _newEnd) public onlyOwner {
        if ( _newStart != 0 ) startTime = _newStart;
        if ( _newEnd != 0 ) endTime = _newEnd;
    }
    
    function updateTimeRC(address _rcContract, uint256 _newStart, uint256 _newEnd) public onlyOwner {
        Ambassador(_rcContract).updateTime( _newStart, _newEnd);
    }
    
    function startICO(uint256 _startTime) public onlyOwner {
        if(_startTime == 0 ) {
            startTime = now;
        } else {
            startTime = _startTime;
        }
        endTime = startTime + 12*weekInSeconds;
    }
    
    event Buy(address buyer, uint256 value, address indexed ambassador);

    function buy(address _buyer) public payable returns(uint256) {
        require(now < endTime); // check if ended
        require( remainingTokens > 0 ); // Check if there are any remaining tokens excluding bonuses
        
        require( tokenContract.balanceOf(this) > remainingTokens); // should have enough balance
        
        uint256 oneXCC = 10 ** decimals;
        uint256 tokenAmount = msg.value.mul(oneXCC).div(tokenValue);
        
        
        uint256 bonusRate; // decimals of bonus 20% = 2000
        address currentAmbassador = address(0);
        if ( ambassadorAddressOf[msg.sender] != address(0) ) { // if is an authorized ambassadorContract
            currentAmbassador = msg.sender;
            bonusRate = 0; // Ambassador Comunity should claim own bonus at the end of RC 
            
        } else { // if is directly called to CoinCrowdICO contract
            require(now > startTime); // check if started for public user
            
            if( now > startTime + weekInSeconds*0  ) { bonusRate = 2000; }
            if( now > startTime + weekInSeconds*1  ) { bonusRate = 1833; }
            if( now > startTime + weekInSeconds*2  ) { bonusRate = 1667; }
            if( now > startTime + weekInSeconds*3  ) { bonusRate = 1500; }
            if( now > startTime + weekInSeconds*4  ) { bonusRate = 1333; }
            if( now > startTime + weekInSeconds*5  ) { bonusRate = 1167; }
            if( now > startTime + weekInSeconds*6  ) { bonusRate = 1000; }
            if( now > startTime + weekInSeconds*7  ) { bonusRate = 833; }
            if( now > startTime + weekInSeconds*8  ) { bonusRate = 667; }
            if( now > startTime + weekInSeconds*9  ) { bonusRate = 500; }
            if( now > startTime + weekInSeconds*10 ) { bonusRate = 333; }
            if( now > startTime + weekInSeconds*11 ) { bonusRate = 167; }
            if( now > startTime + weekInSeconds*12 ) { bonusRate = 0; }
        }
        
        if ( remainingTokens < tokenAmount ) {
            uint256 refund = (tokenAmount - remainingTokens).mul(tokenValue).div(oneXCC);
            tokenAmount = remainingTokens;
            owner.transfer(msg.value-refund);
			remainingTokens = 0; // set remaining token to 0
             _buyer.transfer(refund);
        } else {
			remainingTokens = remainingTokens.sub(tokenAmount); // update remaining token without bonus
            owner.transfer(msg.value);
        }
        
        uint256 tokenAmountWithBonus = tokenAmount.add(tokenAmount.mul( bonusRate ).div(10**4)); //add token bonus
        
        tokenContract.transfer(_buyer, tokenAmountWithBonus);
        Buy(_buyer, tokenAmountWithBonus, currentAmbassador);
        
        totalSoldTokensWithBonus += tokenAmountWithBonus; 
		totalSoldTokensWithoutBonus += tokenAmount;
		
        return tokenAmount; // retun tokenAmount without bonuses for easier calculations
    }

    event NewAmbassador(address ambassador, address contr);
    
    function addMeByRC(address _ambassadorAddr) public {
        require(tx.origin == owner);
        
        ambassadorAddressOf[ msg.sender ]  = _ambassadorAddr;
        
        NewAmbassador(_ambassadorAddr, msg.sender);
    }

    function withdraw(address to, uint256 value) public onlyOwner {
        to.transfer(value);
    }
    
    function updateTokenContract(address _tokenContract) public onlyOwner {
        tokenContract = tokenInterface(_tokenContract);
    }

    function withdrawTokens(address to, uint256 value) public onlyOwner returns (bool) {
        return tokenContract.transfer(to, value);
    }
    
    function claimPremium(address _buyer, uint256 _amount) public returns(bool) {
        require( ambassadorAddressOf[msg.sender] != address(0) ); // Check if is an authorized _ambassadorContract
        return tokenContract.transfer(_buyer, _amount);
    }

    function () public payable {
        buy(msg.sender);
    }
}