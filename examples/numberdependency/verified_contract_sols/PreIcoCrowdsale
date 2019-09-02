pragma solidity 0.4.18;


/*
 * https://github.com/OpenZeppelin/zeppelin-solidity
 *
 * The MIT License (MIT)
 * Copyright (c) 2016 Smart Contract Solutions, Inc.
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


/*
 * https://github.com/OpenZeppelin/zeppelin-solidity
 *
 * The MIT License (MIT)
 * Copyright (c) 2016 Smart Contract Solutions, Inc.
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


/**
 * @title One-time schedulable contract
 * @author Jakub Stefanski (https://github.com/jstefanski)
 *
 * https://github.com/OnLivePlatform/onlive-contracts
 *
 * The BSD 3-Clause Clear License
 * Copyright (c) 2018 OnLive LTD
 */
contract Schedulable is Ownable {

    /**
     * @dev First block when contract is active (inclusive). Zero if not scheduled.
     */
    uint256 public startBlock;

    /**
     * @dev Last block when contract is active (inclusive). Zero if not scheduled.
     */
    uint256 public endBlock;

    /**
     * @dev Contract scheduled within given blocks
     * @param startBlock uint256 The first block when contract is active (inclusive)
     * @param endBlock uint256 The last block when contract is active (inclusive)
     */
    event Scheduled(uint256 startBlock, uint256 endBlock);

    modifier onlyNotZero(uint256 value) {
        require(value != 0);
        _;
    }

    modifier onlyScheduled() {
        require(isScheduled());
        _;
    }

    modifier onlyNotScheduled() {
        require(!isScheduled());
        _;
    }

    modifier onlyActive() {
        require(isActive());
        _;
    }

    modifier onlyNotActive() {
        require(!isActive());
        _;
    }

    /**
     * @dev Schedule contract activation for given block range
     * @param _startBlock uint256 The first block when contract is active (inclusive)
     * @param _endBlock uint256 The last block when contract is active (inclusive)
     */
    function schedule(uint256 _startBlock, uint256 _endBlock)
        public
        onlyOwner
        onlyNotScheduled
        onlyNotZero(_startBlock)
        onlyNotZero(_endBlock)
    {
        require(_startBlock < _endBlock);

        startBlock = _startBlock;
        endBlock = _endBlock;

        Scheduled(_startBlock, _endBlock);
    }

    /**
     * @dev Check whether activation is scheduled
     */
    function isScheduled() public view returns (bool) {
        return startBlock > 0 && endBlock > 0;
    }

    /**
     * @dev Check whether contract is currently active
     */
    function isActive() public view returns (bool) {
        return block.number >= startBlock && block.number <= endBlock;
    }
}


/**
 * @title Pre-ICO Crowdsale with constant price and limited supply
 * @author Jakub Stefanski (https://github.com/jstefanski)
 *
 * https://github.com/OnLivePlatform/onlive-contracts
 *
 * The BSD 3-Clause Clear License
 * Copyright (c) 2018 OnLive LTD
 */
contract Mintable {
    uint256 public decimals;

    function mint(address to, uint256 amount) public;
}


/**
 * @title Crowdsale for off-chain payment methods
 * @author Jakub Stefanski (https://github.com/jstefanski)
 *
 * https://github.com/OnLivePlatform/onlive-contracts
 *
 * The BSD 3-Clause Clear License
 * Copyright (c) 2018 OnLive LTD
 */
contract PreIcoCrowdsale is Schedulable {

    using SafeMath for uint256;

    /**
     * @dev Address of contribution wallet
     */
    address public wallet;

    /**
     * @dev Address of mintable token instance
     */
    Mintable public token;

    /**
     * @dev Current amount of tokens available for sale
     */
    uint256 public availableAmount;

    /**
     * @dev Price of token in Wei
     */
    uint256 public price;

    /**
     * @dev Minimum ETH value sent as contribution
     */
    uint256 public minValue;

    /**
     * @dev Indicates whether contribution identified by bytes32 id is already registered
     */
    mapping (bytes32 => bool) public isContributionRegistered;

    function PreIcoCrowdsale(
        address _wallet,
        Mintable _token,
        uint256 _availableAmount,
        uint256 _price,
        uint256 _minValue
    )
        public
        onlyValid(_wallet)
        onlyValid(_token)
        onlyNotZero(_availableAmount)
        onlyNotZero(_price)
    {
        wallet = _wallet;
        token = _token;
        availableAmount = _availableAmount;
        price = _price;
        minValue = _minValue;
    }

    /**
     * @dev Contribution is accepted
     * @param contributor address The recipient of the tokens
     * @param value uint256 The amount of contributed ETH
     * @param amount uint256 The amount of tokens
     */
    event ContributionAccepted(address indexed contributor, uint256 value, uint256 amount);

    /**
     * @dev Off-chain contribution registered
     * @param id bytes32 A unique contribution id
     * @param contributor address The recipient of the tokens
     * @param amount uint256 The amount of tokens
     */
    event ContributionRegistered(bytes32 indexed id, address indexed contributor, uint256 amount);

    modifier onlyValid(address addr) {
        require(addr != address(0));
        _;
    }

    modifier onlySufficientValue(uint256 value) {
        require(value >= minValue);
        _;
    }

    modifier onlySufficientAvailableTokens(uint256 amount) {
        require(availableAmount >= amount);
        _;
    }

    modifier onlyUniqueContribution(bytes32 id) {
        require(!isContributionRegistered[id]);
        _;
    }

    /**
     * @dev Accept ETH transfers as contributions
     */
    function () public payable {
        acceptContribution(msg.sender, msg.value);
    }

    /**
     * @dev Contribute ETH in exchange for tokens
     * @param contributor address The address that receives tokens
     */
    function contribute(address contributor) public payable returns (uint256) {
        return acceptContribution(contributor, msg.value);
    }

    /**
     * @dev Register contribution with given id
     * @param id bytes32 A unique contribution id
     * @param contributor address The recipient of the tokens
     * @param amount uint256 The amount of tokens
     */
    function registerContribution(bytes32 id, address contributor, uint256 amount)
        public
        onlyOwner
        onlyActive
        onlyValid(contributor)
        onlyNotZero(amount)
        onlyUniqueContribution(id)
    {
        isContributionRegistered[id] = true;
        mintTokens(contributor, amount);

        ContributionRegistered(id, contributor, amount);
    }

    /**
     * @dev Calculate amount of ONL tokens received for given ETH value
     * @param value uint256 Contribution value in ETH
     * @return uint256 Amount of received ONL tokens
     */
    function calculateContribution(uint256 value) public view returns (uint256) {
        return value.mul(10 ** token.decimals()).div(price);
    }

    function acceptContribution(address contributor, uint256 value)
        private
        onlyActive
        onlyValid(contributor)
        onlySufficientValue(value)
        returns (uint256)
    {
        uint256 amount = calculateContribution(value);
        mintTokens(contributor, amount);

        wallet.transfer(value);

        ContributionAccepted(contributor, value, amount);

        return amount;
    }

    function mintTokens(address to, uint256 amount)
        private
        onlySufficientAvailableTokens(amount)
    {
        availableAmount = availableAmount.sub(amount);
        token.mint(to, amount);
    }
}