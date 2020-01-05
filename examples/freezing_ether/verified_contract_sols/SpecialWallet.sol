pragma solidity ^0.4.18;

// File: contracts/ownership/Ownable.sol

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
  function Ownable() public {
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
  function transferOwnership(address newOwner) public onlyOwner {
    require(newOwner != address(0));
    OwnershipTransferred(owner, newOwner);
    owner = newOwner;
  }

}

// File: contracts/InvestedProvider.sol



// File: contracts/AddressesFilterFeature.sol



// File: contracts/math/SafeMath.sol

/**
 * @title SafeMath
 * @dev Math operations with safety checks that throw on error
 */
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

// File: contracts/token/ERC20Basic.sol

/**
 * @title ERC20Basic
 * @dev Simpler version of ERC20 interface
 * @dev see https://github.com/ethereum/EIPs/issues/179
 */


// File: contracts/token/BasicToken.sol

/**
 * @title Basic token
 * @dev Basic version of StandardToken, with no allowances.
 */


// File: contracts/token/ERC20.sol

/**
 * @title ERC20 interface
 * @dev see https://github.com/ethereum/EIPs/issues/20
 */


// File: contracts/token/StandardToken.sol

/**
 * @title Standard ERC20 token
 *
 * @dev Implementation of the basic standard token.
 * @dev https://github.com/ethereum/EIPs/issues/20
 * @dev Based on code by FirstBlood: https://github.com/Firstbloodio/token/blob/master/smart_contract/FirstBloodToken.sol
 */


// File: contracts/MintableToken.sol



// File: contracts/TokenProvider.sol



// File: contracts/MintTokensInterface.sol



// File: contracts/MintTokensFeature.sol



// File: contracts/PercentRateProvider.sol

contract PercentRateProvider {

  uint public percentRate = 100;

}

// File: contracts/PercentRateFeature.sol

contract PercentRateFeature is Ownable, PercentRateProvider {

  function setPercentRate(uint newPercentRate) public onlyOwner {
    percentRate = newPercentRate;
  }

}

// File: contracts/RetrieveTokensFeature.sol



// File: contracts/WalletProvider.sol



// File: contracts/CommonSale.sol



// File: contracts/SpecialWallet.sol

contract SpecialWallet is PercentRateFeature {
  
  using SafeMath for uint;

  uint public endDate;

  uint initialBalance;

  bool public started;

  uint public startDate;

  uint availableAfterStart;

  uint public withdrawed;

  uint public startQuater;

  uint public quater1;

  uint public quater2;

  uint public quater3;

  uint public quater4;

  modifier notStarted() {
    require(!started);
    _;
  }

  function start() public onlyOwner notStarted {
    started = true;
    startDate = now;

    uint year = 1 years;
    uint quater = year.div(4);
    uint prevYear = endDate.sub(1 years);

    quater1 = prevYear;
    quater2 = prevYear.add(quater);
    quater3 = prevYear.add(quater.mul(2));
    quater4 = prevYear.add(quater.mul(3));

    initialBalance = this.balance;

    startQuater = curQuater();
  }

  function curQuater() public view returns (uint) {
    if(now > quater4) 
      return 4;
    if(now > quater3) 
      return 3;
    if(now > quater2) 
      return 2;
    return 1;
  }
 
  function setAvailableAfterStart(uint newAvailableAfterStart) public onlyOwner notStarted {
    availableAfterStart = newAvailableAfterStart;
  }

  function setEndDate(uint newEndDate) public onlyOwner notStarted {
    endDate = newEndDate;
  }

  function withdraw(address to) public onlyOwner {
    require(started);
    if(now >= endDate) {
      to.transfer(this.balance);
    } else {
      uint cQuater = curQuater();
      uint toTransfer = initialBalance.mul(availableAfterStart).div(percentRate);
      if(startQuater < 4 && cQuater > startQuater) {
        uint secondInitialBalance = initialBalance.sub(toTransfer);
        uint quaters = 4;
        uint allQuaters = quaters.sub(startQuater);        
        uint value = secondInitialBalance.mul(cQuater.sub(startQuater)).div(allQuaters);         
        toTransfer = toTransfer.add(value);
      }
      toTransfer = toTransfer.sub(withdrawed); 
      to.transfer(toTransfer);
      withdrawed = withdrawed.add(toTransfer);        
    }
  }

  function () public payable {
  }

}

// File: contracts/AssembledCommonSale.sol



// File: contracts/WalletsPercents.sol



// File: contracts/ExtendedWalletsMintTokensFeature.sol

//import './PercentRateProvider.sol';



// File: contracts/StagedCrowdsale.sol



// File: contracts/ITO.sol



// File: contracts/NextSaleAgentFeature.sol



// File: contracts/SoftcapFeature.sol



// File: contracts/PreITO.sol



// File: contracts/ReceivingContractCallback.sol



// File: contracts/Token.sol



// File: contracts/Configurator.sol

