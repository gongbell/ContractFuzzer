/**
 *Submitted for verification at Etherscan.io on 2017-11-27
*/

pragma solidity ^0.4.11;

/**
 * @title SafeMath
 * @dev Math operations with safety checks that throw on error
 */
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

/**
 * @title Ownable
 * @dev The Ownable contract has an owner address, and provides basic authorization control
 * functions, this simplifies the implementation of "user permissions".
 */
contract Ownable {
    address public owner;


    /**
     * @dev The Ownable constructor sets the original `owner` of the contract to the sender
     * account.
     */
    function Ownable() {
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
    function transferOwnership(address newOwner) onlyOwner {
        if (newOwner != address(0)) {
            owner = newOwner;
        }
    }

}

/**
 * @title Pausable
 * @dev Base contract which allows children to implement an emergency stop mechanism.
 */
contract Pausable is Ownable {
    event Pause();
    event Unpause();

    bool public paused = false;


    /**
     * @dev modifier to allow actions only when the contract IS paused
     */
    modifier whenNotPaused() {
        require(!paused);
        _;
    }

    /**
     * @dev modifier to allow actions only when the contract IS NOT paused
     */
    modifier whenPaused {
        require(paused);
        _;
    }

    /**
     * @dev called by the owner to pause, triggers stopped state
     */
    function pause() onlyOwner whenNotPaused returns (bool) {
        paused = true;
        Pause();
        return true;
    }

    /**
     * @dev called by the owner to unpause, returns to normal state
     */
    function unpause() onlyOwner whenPaused returns (bool) {
        paused = false;
        Unpause();
        return true;
    }
}

/**
 * @title ERC20Basic
 * @dev Simpler version of ERC20 interface
 * @dev see https://github.com/ethereum/EIPs/issues/179
 */
contract ERC20Basic {
    uint256 public totalSupply;
    function balanceOf(address who) constant returns (uint256);
    function transfer(address to, uint256 value) returns (bool);
    event Transfer(address indexed from, address indexed to, uint256 value);
}

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
    function transfer(address _to, uint256 _value) returns (bool) {
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
    function balanceOf(address _owner) constant returns (uint256 balance) {
        return balances[_owner];
    }

}

/**
 * @title ERC20 interface
 * @dev see https://github.com/ethereum/EIPs/issues/20
 */
contract ERC20 is ERC20Basic {
    function allowance(address owner, address spender) constant returns (uint256);
    function transferFrom(address from, address to, uint256 value) returns (bool);
    function approve(address spender, uint256 value) returns (bool);
    event Approval(address indexed owner, address indexed spender, uint256 value);
}

/**
 * @title Standard ERC20 token
 *
 * @dev Implementation of the basic standard token.
 * @dev https://github.com/ethereum/EIPs/issues/20
 * @dev Based on code by FirstBlood: https://github.com/Firstbloodio/token/blob/master/smart_contract/FirstBloodToken.sol
 */
contract StandardToken is ERC20, BasicToken {

    mapping (address => mapping (address => uint256)) allowed;


    /**
     * @dev Transfer tokens from one address to another
     * @param _from address The address which you want to send tokens from
     * @param _to address The address which you want to transfer to
     * @param _value uint256 the amout of tokens to be transfered
     */
    function transferFrom(address _from, address _to, uint256 _value) returns (bool) {
        var _allowance = allowed[_from][msg.sender];

        // Check is not needed because sub(_allowance, _value) will already throw if this condition is not met
        // require (_value <= _allowance);

        balances[_to] = balances[_to].add(_value);
        balances[_from] = balances[_from].sub(_value);
        allowed[_from][msg.sender] = _allowance.sub(_value);
        Transfer(_from, _to, _value);
        return true;
    }

    /**
     * @dev Aprove the passed address to spend the specified amount of tokens on behalf of msg.sender.
     * @param _spender The address which will spend the funds.
     * @param _value The amount of tokens to be spent.
     */
    function approve(address _spender, uint256 _value) returns (bool) {

        // To change the approve amount you first have to reduce the addresses`
        //  allowance to zero by calling `approve(_spender, 0)` if it is not
        //  already 0 to mitigate the race condition described here:
        //  https://github.com/ethereum/EIPs/issues/20#issuecomment-263524729
        require((_value == 0) || (allowed[msg.sender][_spender] == 0));

        allowed[msg.sender][_spender] = _value;
        Approval(msg.sender, _spender, _value);
        return true;
    }

    /**
     * @dev Function to check the amount of tokens that an owner allowed to a spender.
     * @param _owner address The address which owns the funds.
     * @param _spender address The address which will spend the funds.
     * @return A uint256 specifing the amount of tokens still avaible for the spender.
     */
    function allowance(address _owner, address _spender) constant returns (uint256 remaining) {
        return allowed[_owner][_spender];
    }

}

/**
 * @title HoQuToken
 * @dev HoQu.io token contract.
 */
contract HoQuToken is StandardToken, Pausable {

    string public constant name = "HOQU Token";
    string public constant symbol = "HQX";
    uint32 public constant decimals = 18;

    /**
     * @dev Give all tokens to msg.sender.
     */
    function HoQuToken(uint _totalSupply) {
        require (_totalSupply > 0);
        totalSupply = balances[msg.sender] = _totalSupply;
    }

    function transfer(address _to, uint _value) whenNotPaused returns (bool) {
        return super.transfer(_to, _value);
    }

    function transferFrom(address _from, address _to, uint _value) whenNotPaused returns (bool) {
        return super.transferFrom(_from, _to, _value);
    }
}

/**
 * @title ClaimableCrowdsale
 * @title HoQu.io claimable crowdsale contract.
 */
contract ClaimableCrowdsale is Pausable {
    using SafeMath for uint256;

    // all accepted ethers will be sent to this address
    address beneficiaryAddress;

    // all remain tokens after ICO should go to that address
    address public bankAddress;

    // token instance
    HoQuToken public token;

    uint256 public maxTokensAmount;
    uint256 public issuedTokensAmount = 0;
    uint256 public minBuyableAmount;
    uint256 public tokenRate; // amount of HQX per 1 ETH
    
    uint256 endDate;

    bool public isFinished = false;

    // buffer for claimable tokens
    mapping(address => uint256) public tokens;
    mapping(address => bool) public approved;
    mapping(uint32 => address) internal tokenReceivers;
    uint32 internal receiversCount;

    /**
    * Events for token purchase logging
    */
    event TokenBought(address indexed _buyer, uint256 _tokens, uint256 _amount);
    event TokenAdded(address indexed _receiver, uint256 _tokens, uint256 _equivalentAmount);
    event TokenToppedUp(address indexed _receiver, uint256 _tokens, uint256 _equivalentAmount);
    event TokenSubtracted(address indexed _receiver, uint256 _tokens, uint256 _equivalentAmount);
    event TokenSent(address indexed _receiver, uint256 _tokens);

    modifier inProgress() {
        require (!isFinished);
        require (issuedTokensAmount < maxTokensAmount);
        require (now <= endDate);
        _;
    }
    
    /**
    * @param _tokenAddress address of a HQX token contract
    * @param _bankAddress address for remain HQX tokens accumulation
    * @param _beneficiaryAddress accepted ETH go to this address
    * @param _tokenRate rate HQX per 1 ETH
    * @param _minBuyableAmount min ETH per each buy action (in ETH wei)
    * @param _maxTokensAmount ICO HQX capacity (in HQX wei)
    * @param _endDate the date when ICO will expire
    */
    function ClaimableCrowdsale(
        address _tokenAddress,
        address _bankAddress,
        address _beneficiaryAddress,
        uint256 _tokenRate,
        uint256 _minBuyableAmount,
        uint256 _maxTokensAmount,
        uint256 _endDate
    ) {
        token = HoQuToken(_tokenAddress);

        bankAddress = _bankAddress;
        beneficiaryAddress = _beneficiaryAddress;

        tokenRate = _tokenRate;
        minBuyableAmount = _minBuyableAmount;
        maxTokensAmount = _maxTokensAmount;

        endDate = _endDate;
    }

    /*
     * @dev Set new HoQu token exchange rate.
     */
    function setTokenRate(uint256 _tokenRate) onlyOwner {
        require (_tokenRate > 0);
        tokenRate = _tokenRate;
    }

    /**
     * Buy HQX. Tokens will be stored in contract until claim stage
     */
    function buy() payable inProgress whenNotPaused {
        uint256 payAmount = msg.value;
        uint256 returnAmount = 0;

        // calculate token amount to be transfered to investor
        uint256 tokensAmount = tokenRate.mul(payAmount);
    
        if (issuedTokensAmount + tokensAmount > maxTokensAmount) {
            tokensAmount = maxTokensAmount.sub(issuedTokensAmount);
            payAmount = tokensAmount.div(tokenRate);
            returnAmount = msg.value.sub(payAmount);
        }
    
        issuedTokensAmount = issuedTokensAmount.add(tokensAmount);
        require (issuedTokensAmount <= maxTokensAmount);

        storeTokens(msg.sender, tokensAmount);
        TokenBought(msg.sender, tokensAmount, payAmount);

        beneficiaryAddress.transfer(payAmount);
    
        if (returnAmount > 0) {
            msg.sender.transfer(returnAmount);
        }
    }

    /**
     * Add HQX payed by another crypto (BTC, LTC). Tokens will be stored in contract until claim stage
     */
    function add(address _receiver, uint256 _equivalentEthAmount) onlyOwner inProgress whenNotPaused {
        uint256 tokensAmount = tokenRate.mul(_equivalentEthAmount);
        issuedTokensAmount = issuedTokensAmount.add(tokensAmount);

        storeTokens(_receiver, tokensAmount);
        TokenAdded(_receiver, tokensAmount, _equivalentEthAmount);
    }

    /**
     * Add HQX by referral program. Tokens will be stored in contract until claim stage
     */
    function topUp(address _receiver, uint256 _equivalentEthAmount) onlyOwner whenNotPaused {
        uint256 tokensAmount = tokenRate.mul(_equivalentEthAmount);
        issuedTokensAmount = issuedTokensAmount.add(tokensAmount);

        storeTokens(_receiver, tokensAmount);
        TokenToppedUp(_receiver, tokensAmount, _equivalentEthAmount);
    }

    /**
     * Reduce bought HQX amount. Emergency use only
     */
    function sub(address _receiver, uint256 _equivalentEthAmount) onlyOwner whenNotPaused {
        uint256 tokensAmount = tokenRate.mul(_equivalentEthAmount);

        require (tokens[_receiver] >= tokensAmount);

        tokens[_receiver] = tokens[_receiver].sub(tokensAmount);
        issuedTokensAmount = issuedTokensAmount.sub(tokensAmount);

        TokenSubtracted(_receiver, tokensAmount, _equivalentEthAmount);
    }

    /**
     * Internal method for storing tokens in contract until claim stage
     */
    function storeTokens(address _receiver, uint256 _tokensAmount) internal whenNotPaused {
        if (tokens[_receiver] == 0) {
            tokenReceivers[receiversCount] = _receiver;
            receiversCount++;
            approved[_receiver] = false;
        }
        tokens[_receiver] = tokens[_receiver].add(_tokensAmount);
    }

    /**
     * Claim all bought HQX. Available tokens will be sent to transaction sender address if it is approved
     */
    function claim() whenNotPaused {
        claimFor(msg.sender);
    }

    /**
     * Claim all bought HQX for specific approved address
     */
    function claimOne(address _receiver) onlyOwner whenNotPaused {
        claimFor(_receiver);
    }

    /**
     * Claim all bought HQX for all approved addresses
     */
    function claimAll() onlyOwner whenNotPaused {
        for (uint32 i = 0; i < receiversCount; i++) {
            address receiver = tokenReceivers[i];
            if (approved[receiver] && tokens[receiver] > 0) {
                claimFor(receiver);
            }
        }
    }

    /**
     * Internal method for claiming tokens for specific approved address
     */
    function claimFor(address _receiver) internal whenNotPaused {
        require(approved[_receiver]);
        require(tokens[_receiver] > 0);

        uint256 tokensToSend = tokens[_receiver];
        tokens[_receiver] = 0;

        token.transferFrom(bankAddress, _receiver, tokensToSend);
        TokenSent(_receiver, tokensToSend);
    }

    function approve(address _receiver) onlyOwner whenNotPaused {
        approved[_receiver] = true;
    }
    
    /**
     * Finish Sale.
     */
    function finish() onlyOwner {
        require (issuedTokensAmount >= maxTokensAmount || now > endDate);
        require (!isFinished);
        isFinished = true;
        token.transfer(bankAddress, token.balanceOf(this));
    }

    function getReceiversCount() constant onlyOwner returns (uint32) {
        return receiversCount;
    }

    function getReceiver(uint32 i) constant onlyOwner returns (address) {
        return tokenReceivers[i];
    }
    
    /**
     * Buy HQX. Tokens will be stored in contract until claim stage
     */
    function() external payable {
        buy();
    }
}

/**
 * @title ChangeableRateCrowdsale
 * @dev HoQu.io Main Sale stage
 */
contract ChangeableRateCrowdsale is ClaimableCrowdsale {

    struct RateBoundary {
        uint256 amount;
        uint256 rate;
    }

    mapping (uint => RateBoundary) public rateBoundaries;
    uint public currentBoundary = 0;
    uint public numOfBoundaries = 0;
    uint256 public nextBoundaryAmount;

    /**
    * @param _tokenAddress address of a HQX token contract
    * @param _bankAddress address for remain HQX tokens accumulation
    * @param _beneficiaryAddress accepted ETH go to this address
    * @param _tokenRate rate HQX per 1 ETH
    * @param _minBuyableAmount min ETH per each buy action (in ETH wei)
    * @param _maxTokensAmount ICO HQX capacity (in HQX wei)
    * @param _endDate the date when ICO will expire
    */
    function ChangeableRateCrowdsale(
        address _tokenAddress,
        address _bankAddress,
        address _beneficiaryAddress,
        uint256 _tokenRate,
        uint256 _minBuyableAmount,
        uint256 _maxTokensAmount,
        uint256 _endDate
    ) ClaimableCrowdsale(
        _tokenAddress,
        _bankAddress,
        _beneficiaryAddress,
        _tokenRate,
        _minBuyableAmount,
        _maxTokensAmount,
        _endDate
    ) {
        rateBoundaries[numOfBoundaries++] = RateBoundary({
            amount : 13777764 ether,
            rate : 6000
        });
        rateBoundaries[numOfBoundaries++] = RateBoundary({
            amount : 27555528 ether,
            rate : 5750
        });
        rateBoundaries[numOfBoundaries++] = RateBoundary({
            amount : 41333292 ether,
            rate : 5650
        });
        rateBoundaries[numOfBoundaries++] = RateBoundary({
            amount : 55111056 ether,
            rate : 5550
        });
        rateBoundaries[numOfBoundaries++] = RateBoundary({
            amount : 68888820 ether,
            rate : 5450
        });
        rateBoundaries[numOfBoundaries++] = RateBoundary({
            amount : 82666584 ether,
            rate : 5350
        });
        rateBoundaries[numOfBoundaries++] = RateBoundary({
            amount : 96444348 ether,
            rate : 5250
        });
        rateBoundaries[numOfBoundaries++] = RateBoundary({
            amount : 110222112 ether,
            rate : 5150
        });
        rateBoundaries[numOfBoundaries++] = RateBoundary({
            amount : 137777640 ether,
            rate : 5000
        });
        nextBoundaryAmount = rateBoundaries[currentBoundary].amount;
    }

    /**
     * Internal method to change rate if boundary is hit
     */
    function touchRate() internal {
        if (issuedTokensAmount >= nextBoundaryAmount) {
            currentBoundary++;
            if (currentBoundary >= numOfBoundaries) {
                nextBoundaryAmount = maxTokensAmount;
            }
            else {
                nextBoundaryAmount = rateBoundaries[currentBoundary].amount;
                tokenRate = rateBoundaries[currentBoundary].rate;
            }
        }
    }

    /**
     * Inherited internal method for storing tokens in contract until claim stage
     */
    function storeTokens(address _receiver, uint256 _tokensAmount) internal whenNotPaused {
        ClaimableCrowdsale.storeTokens(_receiver, _tokensAmount);
        touchRate();
    }
}
