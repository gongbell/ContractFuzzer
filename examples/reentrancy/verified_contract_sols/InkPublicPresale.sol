pragma solidity ^0.4.18;

// File: zeppelin-solidity/contracts/math/SafeMath.sol

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

// File: zeppelin-solidity/contracts/ownership/Ownable.sol

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

// File: zeppelin-solidity/contracts/token/ERC20Basic.sol

/**
 * @title ERC20Basic
 * @dev Simpler version of ERC20 interface
 * @dev see https://github.com/ethereum/EIPs/issues/179
 */
contract ERC20Basic {
  uint256 public totalSupply;
  function balanceOf(address who) public view returns (uint256);
  function transfer(address to, uint256 value) public returns (bool);
  event Transfer(address indexed from, address indexed to, uint256 value);
}

// File: zeppelin-solidity/contracts/token/BasicToken.sol

/**
 * @title Basic token
 * @dev Basic version of StandardToken, with no allowances.
 */
contract BasicToken is ERC20Basic {
  using SafeMath for uint256;

  mapping(address => uint256) balances;

  /**
  * @dev transfer token for a specified address
  * @param _to The address to transfer to.
  * @param _value The amount to be transferred.
  */
  function transfer(address _to, uint256 _value) public returns (bool) {
    require(_to != address(0));
    require(_value <= balances[msg.sender]);

    // SafeMath.sub will throw if there is not enough balance.
    balances[msg.sender] = balances[msg.sender].sub(_value);
    balances[_to] = balances[_to].add(_value);
    Transfer(msg.sender, _to, _value);
    return true;
  }

  /**
  * @dev Gets the balance of the specified address.
  * @param _owner The address to query the the balance of.
  * @return An uint256 representing the amount owned by the passed address.
  */
  function balanceOf(address _owner) public view returns (uint256 balance) {
    return balances[_owner];
  }

}

// File: contracts/InkPublicPresale.sol

contract InkPublicPresale is Ownable {
  using SafeMath for uint256;

  // Flag to indicate whether or not the presale is currently active or is paused.
  // This flag is used both before the presale is finalized as well as after.
  // Pausing the presale before finalize means that no further contributions can
  // be made. Pausing the presale after finalize means that no one can claim
  // XNK tokens.
  bool public active;

  // Flag to indicate whether or not contributions can be refunded.
  bool private refundable;

  // The global minimum contribution (in Wei) imposed on all contributors.
  uint256 public globalMin;
  // The global maximum contribution (in Wei) imposed on all contributors.
  // Contributor also have a personal max. When evaluating whether or not they
  // can make a contribution, the lower of the global max and personal max is
  // used.
  uint256 public globalMax;
  // The max amount of Ether (in Wei) that is available for contribution.
  uint256 public etherCap;
  // The running count of Ether (in Wei) that is already contributed.
  uint256 private etherContributed;
  // The running count of XNK that is purchased by contributors.
  uint256 private xnkPurchased;
  // The address of the XNK token contract. When this address is set, the
  // presale is considered finalized and no further contributions can be made.
  address public tokenAddress;
  // Max gas price for contributing transactions.
  uint256 public maxGasPrice;

  // Contributors storage mapping.
  mapping(address => Contributor) private contributors;

  struct Contributor {
    bool whitelisted;
    // The individual rate (in XNK).
    uint256 rate;
    // The individual max contribution (in Wei).
    uint256 max;
    // The amount (in Wei) the contributor has contributed.
    uint256 balance;
  }

  // The presale is considered finalized when the token address is set.
  modifier finalized {
    require(tokenAddress != address(0));
    _;
  }

  // The presale is considered not finalized when the token address is not set.
  modifier notFinalized {
    require(tokenAddress == address(0));
    _;
  }

  function InkPublicPresale() public {
    globalMax = 1000000000000000000; // 1.0 Ether
    globalMin = 100000000000000000;  // 0.1 Ether
    maxGasPrice = 40000000000;       // 40 Gwei
  }

  function updateMaxGasPrice(uint256 _maxGasPrice) public onlyOwner {
    require(_maxGasPrice > 0);

    maxGasPrice = _maxGasPrice;
  }

  // Returns the amount of Ether contributed by all contributors.
  function getEtherContributed() public view onlyOwner returns (uint256) {
    return etherContributed;
  }

  // Returns the amount of XNK purchased by all contributes.
  function getXNKPurchased() public view onlyOwner returns (uint256) {
    return xnkPurchased;
  }

  // Update the global ether cap. If the new cap is set to something less than
  // or equal to the current contributed ether (etherContributed), then no
  // new contributions can be made.
  function updateEtherCap(uint256 _newEtherCap) public notFinalized onlyOwner {
    etherCap = _newEtherCap;
  }

  // Update the global max contribution.
  function updateGlobalMax(uint256 _globalMax) public notFinalized onlyOwner {
    require(_globalMax > globalMin);

    globalMax = _globalMax;
  }

  // Update the global minimum contribution.
  function updateGlobalMin(uint256 _globalMin) public notFinalized onlyOwner {
    require(_globalMin > 0);
    require(_globalMin < globalMax);

    globalMin = _globalMin;
  }

  function updateTokenAddress(address _tokenAddress) public finalized onlyOwner {
    require(_tokenAddress != address(0));

    tokenAddress = _tokenAddress;
  }

  // Pause the presale (disables contributions and token claiming).
  function pause() public onlyOwner {
    require(active);
    active = false;
  }

  // Resume the presale (enables contributions and token claiming).
  function resume() public onlyOwner {
    require(!active);
    active = true;
  }

  // Allow contributors to call the refund function to get their contributions
  // returned to their whitelisted address.
  function enableRefund() public onlyOwner {
    require(!refundable);
    refundable = true;
  }

  // Disallow refunds (this is the case by default).
  function disableRefund() public onlyOwner {
    require(refundable);
    refundable = false;
  }

  // Add a contributor to the whitelist.
  function addContributor(address _account, uint256 _rate, uint256 _max) public onlyOwner notFinalized {
    require(_account != address(0));
    require(_rate > 0);
    require(_max >= globalMin);
    require(!contributors[_account].whitelisted);

    contributors[_account].whitelisted = true;
    contributors[_account].max = _max;
    contributors[_account].rate = _rate;
  }

  // Updates a contributor's rate and/or max.
  function updateContributor(address _account, uint256 _newRate, uint256 _newMax) public onlyOwner notFinalized {
    require(_account != address(0));
    require(_newRate > 0);
    require(_newMax >= globalMin);
    require(contributors[_account].whitelisted);

    // Account for any changes in rate since we are keeping track of total XNK
    // purchased.
    if (contributors[_account].balance > 0 && contributors[_account].rate != _newRate) {
      // Put back the purchased XNK for the old rate.
      xnkPurchased = xnkPurchased.sub(contributors[_account].balance.mul(contributors[_account].rate));

      // Purchase XNK at the new rate.
      xnkPurchased = xnkPurchased.add(contributors[_account].balance.mul(_newRate));
    }

    contributors[_account].rate = _newRate;
    contributors[_account].max = _newMax;
  }

  // Remove the contributor from the whitelist. This also refunds their
  // contribution if they have made any.
  function removeContributor(address _account) public onlyOwner {
    require(_account != address(0));
    require(contributors[_account].whitelisted);

    // Remove from whitelist.
    contributors[_account].whitelisted = false;

    // If contributions were made, refund it.
    if (contributors[_account].balance > 0) {
      uint256 balance = contributors[_account].balance;

      contributors[_account].balance = 0;
      xnkPurchased = xnkPurchased.sub(balance.mul(contributors[_account].rate));
      etherContributed = etherContributed.sub(balance);

      // XXX: The exclamation point does nothing. We just want to get rid of the
      // compiler warning that we're not using the returned value of the Ether
      // transfer. The transfer *can* fail but we don't want it to stop the
      // removal of the contributor. We will deal if the transfer failure
      // manually outside this contract.
      !_account.call.value(balance)();
    }

    delete contributors[_account];
  }

  function withdrawXNK(address _to) public onlyOwner {
    require(_to != address(0));

    BasicToken token = BasicToken(tokenAddress);
    assert(token.transfer(_to, token.balanceOf(this)));
  }

  function withdrawEther(address _to) public finalized onlyOwner {
    require(_to != address(0));

    assert(_to.call.value(this.balance)());
  }

  // Returns a contributor's balance.
  function balanceOf(address _account) public view returns (uint256) {
    require(_account != address(0));

    return contributors[_account].balance;
  }

  // When refunds are enabled, contributors can call this function get their
  // contributed Ether back. The contributor must still be whitelisted.
  function refund() public {
    require(active);
    require(refundable);
    require(contributors[msg.sender].whitelisted);

    uint256 balance = contributors[msg.sender].balance;

    require(balance > 0);

    contributors[msg.sender].balance = 0;
    etherContributed = etherContributed.sub(balance);
    xnkPurchased = xnkPurchased.sub(balance.mul(contributors[msg.sender].rate));

    assert(msg.sender.call.value(balance)());
  }

  function airdrop(address _account) public finalized onlyOwner {
    _processPayout(_account);
  }

  // Finalize the presale by specifying the XNK token's contract address.
  // No further contributions can be made. The presale will be in the
  // "token claiming" phase.
  function finalize(address _tokenAddress) public notFinalized onlyOwner {
    require(_tokenAddress != address(0));

    tokenAddress = _tokenAddress;
  }

  // Fallback/payable method for contributions and token claiming.
  function () public payable {
    // Allow the owner to send Ether to the contract arbitrarily.
    if (msg.sender == owner && msg.value > 0) {
      return;
    }

    require(active);
    require(contributors[msg.sender].whitelisted);

    if (tokenAddress == address(0)) {
      // Presale is still accepting contributions.
      _processContribution();
    } else {
      // Presale has been finalized and the user is attempting to claim
      // XNK tokens.
      _processPayout(msg.sender);
    }
  }

  // Process the contribution.
  function _processContribution() private {
    // Must be contributing a positive amount.
    require(msg.value > 0);
    // Limit the transaction's gas price.
    require(tx.gasprice <= maxGasPrice);
    // The sum of the contributor's total contributions must be higher than the
    // global minimum.
    require(contributors[msg.sender].balance.add(msg.value) >= globalMin);
    // The global contribution cap must be higher than what has been contributed
    // by everyone. Otherwise, there's zero room for any contribution.
    require(etherCap > etherContributed);
    // Make sure that this specific contribution does not take the total
    // contribution by everyone over the global contribution cap.
    require(msg.value <= etherCap.sub(etherContributed));

    uint256 newBalance = contributors[msg.sender].balance.add(msg.value);

    // We limit the individual's contribution based on whichever is lower
    // between their individual max or the global max.
    if (globalMax <= contributors[msg.sender].max) {
      require(newBalance <= globalMax);
    } else {
      require(newBalance <= contributors[msg.sender].max);
    }

    // Increment the contributor's balance.
    contributors[msg.sender].balance = newBalance;
    // Increment the total amount of Ether contributed by everyone.
    etherContributed = etherContributed.add(msg.value);
    // Increment the total amount of XNK purchased by everyone.
    xnkPurchased = xnkPurchased.add(msg.value.mul(contributors[msg.sender].rate));
  }

  // Process the token claim.
  function _processPayout(address _recipient) private {
    // The transaction must be 0 Ether.
    require(msg.value == 0);

    uint256 balance = contributors[_recipient].balance;

    // The contributor must have contributed something.
    require(balance > 0);

    // Figure out the amount of XNK the contributor will receive.
    uint256 amount = balance.mul(contributors[_recipient].rate);

    // Zero out the contributor's balance to denote that they have received
    // their tokens.
    contributors[_recipient].balance = 0;

    // Transfer XNK to the contributor.
    assert(BasicToken(tokenAddress).transfer(_recipient, amount));
  }
}