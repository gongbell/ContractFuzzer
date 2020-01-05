pragma solidity ^0.4.24;
/**
* Audited by VZ Chains (vzchains.com)
* HashRushICO.sol creates the client's token for crowdsale and allows for subsequent token sales and minting of tokens
*   Crowdsale contracts edited from original contract code at https://www.ethereum.org/crowdsale#crowdfund-your-idea
*   Additional crowdsale contracts, functions, libraries from OpenZeppelin
*       at https://github.com/OpenZeppelin/zeppelin-solidity/tree/master/contracts/token
*   Token contract edited from original contract code at https://www.ethereum.org/token
*   ERC20 interface and certain token functions adapted from https://github.com/ConsenSys/Tokens
**/
/**
 * @title ERC20 interface
 * @dev see https://github.com/ethereum/EIPs/issues/20
 */
contract ERC20 {
    //Sets events and functions for ERC20 token
    event Approval(address indexed _owner, address indexed _spender, uint256 _value);
    event Transfer(address indexed _from, address indexed _to, uint256 _value);
    function totalSupply() public view returns (uint256);
    function balanceOf(address _owner) public view returns (uint256);
    function transfer(address _to, uint256 _value) public returns (bool);
    function allowance(address _owner, address _spender) public view returns (uint256);
    function approve(address _spender, uint256 _value) public returns (bool);
    function transferFrom(address _from, address _to, uint256 _value) public returns (bool);
}
/**
 * @title Owned
 * @dev The Owned contract has an owner address, and provides basic authorization control
 * functions, this simplifies the implementation of "user permissions".
 */
contract Owned {
    address public owner;
    /**
     * @dev The Ownable constructor sets the original `owner` of the contract to the sender
     * account.
     */
    constructor() public {
        owner = msg.sender;
    }
    /**
     * @dev Throws if called by any account other than the owner.
     */
    modifier onlyOwner {
        require(msg.sender == owner);
        _;
    }
    /**
     * @dev Allows the current owner to transfer control of the contract to a newOwner.
     * @param newOwner The address to transfer ownership to.
     */
    function transferOwnership(address newOwner) onlyOwner public {
        owner = newOwner;
    }
}
library SafeMath {
    function add(uint256 a, uint256 b) internal pure returns (uint256) {
        uint256 c = a + b;
        assert(c >= a);
        return c;
    }
    function div(uint256 a, uint256 b) internal pure returns (uint256) {
        require(b > 0); // Solidity only automatically asserts when dividing by 0
        uint256 c = a / b;
        // assert(a == b * c + a % b); // There is no case in which this doesn't hold
        return c;
    }
    
    
    
    
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
        uint256 c = a - b;
        return c;
    }
}
contract HashRush is ERC20, Owned {
    // Applies SafeMath library to uint256 operations
    using SafeMath for uint256;
    // Public variables
    string public name;
    string public symbol;
    uint256 public decimals;
    // Variables
    uint256 totalSupply_;
    uint256 multiplier;
    // Arrays for balances & allowance
    mapping (address => uint256) balance;
    mapping (address => mapping (address => uint256)) allowed;
    // Modifier to prevent short address attack
    modifier onlyPayloadSize(uint size) {
        if(msg.data.length < size.add(4)) revert();
        _;
    }
    constructor(string tokenName, string tokenSymbol, uint8 decimalUnits, uint256 decimalMultiplier) public {
        name = tokenName;
        symbol = tokenSymbol;
        decimals = decimalUnits;
        multiplier = decimalMultiplier;
    }
    /**
    * @dev Total number of tokens in existence
    */
    function totalSupply() public view returns (uint256) {
        return totalSupply_;
    }
    /**
     * @dev Function to check the amount of tokens that an owner allowed to a spender.
     * @param _owner address The address which owns the funds.
     * @param _spender address The address which will spend the funds.
     * @return A uint256 specifying the amount of tokens still available for the spender.
     */
    function allowance(address _owner, address _spender) public view returns (uint256) {
        return allowed[_owner][_spender];
    }
    /**
     * @dev Approve the passed address to spend the specified amount of tokens on behalf of msg.sender.
     * @param _spender The address which will spend the funds.
     * @param _value The amount of tokens to be spent.
     */
    function approve(address _spender, uint256 _value) public returns (bool) {
        allowed[msg.sender][_spender] = _value;
        emit Approval(msg.sender, _spender, _value);
        return true;
    }
    /**
     * @dev Gets the balance of the specified address.
     * @param _owner The address to query the the balance of.
     * @return An uint256 representing the amount owned by the passed address.
     */
    function balanceOf(address _owner) public view returns (uint256) {
        return balance[_owner];
    }
    /**
     * @dev Transfer token to a specified address
     * @param _to The address to transfer to.
     * @param _value The amount to be transferred.
     */
    function transfer(address _to, uint256 _value) onlyPayloadSize(2 * 32) public returns (bool) {
        require(_to != address(0));
        require(_value <= balance[msg.sender]);
        if ((balance[msg.sender] >= _value)
            && (balance[_to].add(_value) > balance[_to])
        ) {
            balance[msg.sender] = balance[msg.sender].sub(_value);
            balance[_to] = balance[_to].add(_value);
            emit Transfer(msg.sender, _to, _value);
            return true;
        } else {
            return false;
        }
    }
    /**
     * @dev Transfer tokens from one address to another
     * @param _from address The address which you want to send tokens from
     * @param _to address The address which you want to transfer to
     * @param _value uint256 the amount of tokens to be transferred
     */
    function transferFrom(address _from, address _to, uint256 _value) onlyPayloadSize(3 * 32) public returns (bool) {
        require(_to != address(0));
        require(_value <= balance[_from]);
        require(_value <= allowed[_from][msg.sender]);
        if ((balance[_from] >= _value) && (allowed[_from][msg.sender] >= _value) && (balance[_to].add(_value) > balance[_to])) {
            balance[_to] = balance[_to].add(_value);
            balance[_from] = balance[_from].sub(_value);
            allowed[_from][msg.sender] = allowed[_from][msg.sender].sub(_value);
            emit Transfer(_from, _to, _value);
            return true;
        } else {
            return false;
        }
    }
}
contract HashRushICO is Owned, HashRush {
    // Applies SafeMath library to uint256 operations
    using SafeMath for uint256;
    // Public Variables
    address public multiSigWallet;
    uint256 public amountRaised;
    uint256 public startTime;
    uint256 public stopTime;
    uint256 public fixedTotalSupply;
    uint256 public price;
    uint256 public minimumInvestment;
    uint256 public crowdsaleTarget;
    // Variables
    bool crowdsaleClosed = true;
    string tokenName = "HashRush";
    string tokenSymbol = "RUSH";
    uint256 multiplier = 100000000;
    uint8 decimalUnits = 8;
    // Initializes the token
    constructor()
        HashRush(tokenName, tokenSymbol, decimalUnits, multiplier) public {
            multiSigWallet = msg.sender;
            fixedTotalSupply = 70000000;
            fixedTotalSupply = fixedTotalSupply.mul(multiplier);
    }
    /**
     * @dev Fallback function creates tokens and sends to investor when crowdsale is open
     */
    function () public payable {
        require(!crowdsaleClosed
            && (now < stopTime)
            && (msg.value >= minimumInvestment)
            && (totalSupply_.add(msg.value.mul(price).mul(multiplier).div(1 ether)) <= fixedTotalSupply)
            && (amountRaised.add(msg.value.div(1 ether)) <= crowdsaleTarget)
        );
        address recipient = msg.sender;
        amountRaised = amountRaised.add(msg.value.div(1 ether));
        uint256 tokens = msg.value.mul(price).mul(multiplier).div(1 ether);
        totalSupply_ = totalSupply_.add(tokens);
    }
    /**
     * @dev Function to mint tokens
     * @param target The address that will receive the minted tokens.
     * @param amount The amount of tokens to mint.
     * @return A boolean that indicates if the operation was successful.
     */
    function mintToken(address target, uint256 amount) onlyOwner public returns (bool) {
        require(amount > 0);
        require(totalSupply_.add(amount) <= fixedTotalSupply);
        uint256 addTokens = amount;
        balance[target] = balance[target].add(addTokens);
        totalSupply_ = totalSupply_.add(addTokens);
        emit Transfer(0, target, addTokens);
        return true;
    }
    /**
     * @dev Function to set token price
     * @param newPriceperEther New price.
     * @return A boolean that indicates if the operation was successful.
     */
    function setPrice(uint256 newPriceperEther) onlyOwner public returns (uint256) {
        require(newPriceperEther > 0);
        price = newPriceperEther;
        return price;
    }
    /**
     * @dev Function to set the multisig wallet for a crowdsale
     * @param wallet Wallet address.
     * @return A boolean that indicates if the operation was successful.
     */
    function setMultiSigWallet(address wallet) onlyOwner public returns (bool) {
        multiSigWallet = wallet;
        return true;
    }
    /**
     * @dev Function to set the minimum investment to participate in crowdsale
     * @param minimum minimum amount in wei.
     * @return A boolean that indicates if the operation was successful.
     */
    function setMinimumInvestment(uint256 minimum) onlyOwner public returns (bool) {
        minimumInvestment = minimum;
        return true;
    }
    /**
     * @dev Function to set the crowdsale target
     * @param target Target amount in ETH.
     * @return A boolean that indicates if the operation was successful.
     */
    function setCrowdsaleTarget(uint256 target) onlyOwner public returns (bool) {
        crowdsaleTarget = target;
        return true;
    }
    /**
     * @dev Function to start the crowdsale specifying startTime and stopTime
     * @param saleStart Sale start timestamp.
     * @param saleStop Sale stop timestamo.
     * @param salePrice Token price per ether.
     * @param setBeneficiary Beneficiary address.
     * @param minInvestment Minimum investment to participate in crowdsale (wei).
     * @param saleTarget Crowdsale target in ETH
     * @return A boolean that indicates if the operation was successful.
     */
    function startSale(uint256 saleStart, uint256 saleStop, uint256 salePrice, address setBeneficiary, uint256 minInvestment, uint256 saleTarget) onlyOwner public returns (bool) {
        require(saleStop > now);
        startTime = saleStart;
        stopTime = saleStop;
        amountRaised = 0;
        crowdsaleClosed = false;
        setPrice(salePrice);
        setMultiSigWallet(setBeneficiary);
        setMinimumInvestment(minInvestment);
        setCrowdsaleTarget(saleTarget);
        return true;
    }
    /**
     * @dev Function that allows owner to stop the crowdsale immediately
     * @return A boolean that indicates if the operation was successful.
     */
    function stopSale() onlyOwner public returns (bool) {
        stopTime = now;
        crowdsaleClosed = true;
        return true;
    }
}