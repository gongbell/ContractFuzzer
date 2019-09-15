pragma solidity ^0.4.18;

// File: contracts\configs\FabricTokenConfig.sol

/**
 * @title FabricTokenConfig
 *
 * @dev The static configuration for the Fabric Token.
 */
contract FabricTokenConfig {
    // The name of the token.
    string constant NAME = "Fabric Token";

    // The symbol of the token.
    string constant SYMBOL = "FT";

    // The number of decimals for the token.
    uint8 constant DECIMALS = 18;  // Same as ethers.

    // Decimal factor for multiplication purposes.
    uint constant DECIMALS_FACTOR = 10 ** uint(DECIMALS);
}

// File: contracts\interfaces\ERC20TokenInterface.sol

/**
 * @dev The standard ERC20 Token interface.
 */
contract ERC20TokenInterface {
    uint public totalSupply;  /* shorthand for public function and a property */
    event Transfer(address indexed _from, address indexed _to, uint _value);
    event Approval(address indexed _owner, address indexed _spender, uint _value);
    function balanceOf(address _owner) public constant returns (uint balance);
    function transfer(address _to, uint _value) public returns (bool success);
    function transferFrom(address _from, address _to, uint _value) public returns (bool success);
    function approve(address _spender, uint _value) public returns (bool success);
    function allowance(address _owner, address _spender) public constant returns (uint remaining);
}

// File: contracts\libraries\SafeMath.sol

/**
 * @dev Library that helps prevent integer overflows and underflows,
 * inspired by https://github.com/OpenZeppelin/zeppelin-solidity
 */
library SafeMath {
    function plus(uint a, uint b) internal pure returns (uint) {
        uint c = a + b;
        assert(c >= a);

        return c;
    }

    function minus(uint a, uint b) internal pure returns (uint) {
        assert(b <= a);

        return a - b;
    }

    function mul(uint a, uint b) internal pure returns (uint) {
        uint c = a * b;
        assert(a == 0 || c / a == b);
        
        return c;
    }

    function div(uint a, uint b) internal pure returns (uint) {
        uint c = a / b;

        return c;
    }
}

// File: contracts\traits\ERC20Token.sol

/**
 * @title ERC20Token
 *
 * @dev Implements the operations declared in the `ERC20TokenInterface`.
 */
contract ERC20Token is ERC20TokenInterface {
    using SafeMath for uint;

    // Token account balances.
    mapping (address => uint) balances;

    // Delegated number of tokens to transfer.
    mapping (address => mapping (address => uint)) allowed;

    /**
     * @dev Checks the balance of a certain address.
     *
     * @param _account The address which's balance will be checked.
     *
     * @return Returns the balance of the `_account` address.
     */
    function balanceOf(address _account) public constant returns (uint balance) {
        return balances[_account];
    }

    /**
     * @dev Transfers tokens from one address to another.
     *
     * @param _to The target address to which the `_value` number of tokens will be sent.
     * @param _value The number of tokens to send.
     *
     * @return Whether the transfer was successful or not.
     */
    function transfer(address _to, uint _value) public returns (bool success) {
        if (balances[msg.sender] < _value || _value == 0) {

            return false;
        }

        balances[msg.sender] -= _value;
        balances[_to] = balances[_to].plus(_value);

        Transfer(msg.sender, _to, _value);

        return true;
    }

    /**
     * @dev Send `_value` tokens to `_to` from `_from` if `_from` has approved the process.
     *
     * @param _from The address of the sender.
     * @param _to The address of the recipient.
     * @param _value The number of tokens to be transferred.
     *
     * @return Whether the transfer was successful or not.
     */
    function transferFrom(address _from, address _to, uint _value) public returns (bool success) {
        if (balances[_from] < _value || allowed[_from][msg.sender] < _value || _value == 0) {
            return false;
        }

        balances[_to] = balances[_to].plus(_value);
        balances[_from] -= _value;
        allowed[_from][msg.sender] -= _value;

        Transfer(_from, _to, _value);

        return true;
    }

    /**
     * @dev Allows another contract to spend some tokens on your behalf.
     *
     * @param _spender The address of the account which will be approved for transfer of tokens.
     * @param _value The number of tokens to be approved for transfer.
     *
     * @return Whether the approval was successful or not.
     */
    function approve(address _spender, uint _value) public returns (bool success) {
        allowed[msg.sender][_spender] = _value;

        Approval(msg.sender, _spender, _value);

        return true;
    }

    /**
     * @dev Shows the number of tokens approved by `_owner` that are allowed to be transferred by `_spender`.
     *
     * @param _owner The account which allowed the transfer.
     * @param _spender The account which will spend the tokens.
     *
     * @return The number of tokens to be transferred.
     */
    function allowance(address _owner, address _spender) public constant returns (uint remaining) {
        return allowed[_owner][_spender];
    }    
}

// File: contracts\traits\HasOwner.sol

/**
 * @title HasOwner
 *
 * @dev Allows for exclusive access to certain functionality.
 */
contract HasOwner {
    // Current owner.
    address public owner;

    // Conditionally the new owner.
    address public newOwner;

    /**
     * @dev The constructor.
     *
     * @param _owner The address of the owner.
     */
    function HasOwner(address _owner) internal {
        owner = _owner;
    }

    /** 
     * @dev Access control modifier that allows only the current owner to call the function.
     */
    modifier onlyOwner {
        require(msg.sender == owner);
        _;
    }

    /**
     * @dev The event is fired when the current owner is changed.
     *
     * @param _oldOwner The address of the previous owner.
     * @param _newOwner The address of the new owner.
     */
    event OwnershipTransfer(address indexed _oldOwner, address indexed _newOwner);

    /**
     * @dev Transfering the ownership is a two-step process, as we prepare
     * for the transfer by setting `newOwner` and requiring `newOwner` to accept
     * the transfer. This prevents accidental lock-out if something goes wrong
     * when passing the `newOwner` address.
     *
     * @param _newOwner The address of the proposed new owner.
     */
    function transferOwnership(address _newOwner) public onlyOwner {
        newOwner = _newOwner;
    }
 
    /**
     * @dev The `newOwner` finishes the ownership transfer process by accepting the
     * ownership.
     */
    function acceptOwnership() public {
        require(msg.sender == newOwner);

        OwnershipTransfer(owner, newOwner);

        owner = newOwner;
    }
}

// File: contracts\traits\Freezable.sol

/**
 * @title Freezable
 * @dev This trait allows to freeze the transactions in a Token
 */
contract Freezable is HasOwner {
  bool public frozen = false;

  /**
   * @dev Modifier makes methods callable only when the contract is not frozen.
   */
  modifier requireNotFrozen() {
    require(!frozen);
    _;
  }

  /**
   * @dev Allows the owner to "freeze" the contract.
   */
  function freeze() onlyOwner public {
    frozen = true;
  }

  /**
   * @dev Allows the owner to "unfreeze" the contract.
   */
  function unfreeze() onlyOwner public {
    frozen = false;
  }
}

// File: contracts\traits\FreezableERC20Token.sol

/**
 * @title FreezableERC20Token
 *
 * @dev Extends ERC20Token and adds ability to freeze all transfers of tokens.
 */
contract FreezableERC20Token is ERC20Token, Freezable {
    /**
     * @dev Overrides the original ERC20Token implementation by adding whenNotFrozen modifier.
     *
     * @param _to The target address to which the `_value` number of tokens will be sent.
     * @param _value The number of tokens to send.
     *
     * @return Whether the transfer was successful or not.
     */
    function transfer(address _to, uint _value) public requireNotFrozen returns (bool success) {
        return super.transfer(_to, _value);
    }

    /**
     * @dev Send `_value` tokens to `_to` from `_from` if `_from` has approved the process.
     *
     * @param _from The address of the sender.
     * @param _to The address of the recipient.
     * @param _value The number of tokens to be transferred.
     *
     * @return Whether the transfer was successful or not.
     */
    function transferFrom(address _from, address _to, uint _value) public requireNotFrozen returns (bool success) {
        return super.transferFrom(_from, _to, _value);
    }

    /**
     * @dev Allows another contract to spend some tokens on your behalf.
     *
     * @param _spender The address of the account which will be approved for transfer of tokens.
     * @param _value The number of tokens to be approved for transfer.
     *
     * @return Whether the approval was successful or not.
     */
    function approve(address _spender, uint _value) public requireNotFrozen returns (bool success) {
        return super.approve(_spender, _value);
    }

}

// File: contracts\FabricToken.sol

/**
 * @title Fabric Token
 *
 * @dev A standard token implementation of the ERC20 token standard with added
 *      HasOwner trait and initialized using the configuration constants.
 */
contract FabricToken is FabricTokenConfig, HasOwner, FreezableERC20Token {
    // The name of the token.
    string public name;

    // The symbol for the token.
    string public symbol;

    // The decimals of the token.
    uint8 public decimals;

    /**
     * @dev The constructor. Initially sets `totalSupply` and the balance of the
     *      `owner` address according to the initialization parameter.
     */
    function FabricToken(uint _totalSupply) public
        HasOwner(msg.sender)
    {
        name = NAME;
        symbol = SYMBOL;
        decimals = DECIMALS;
        totalSupply = _totalSupply;
        balances[owner] = _totalSupply;
    }
}

// File: contracts\configs\FabricTokenFundraiserConfig.sol

/**
 * @title FabricTokenFundraiserConfig
 *
 * @dev The static configuration for the Fabric Token fundraiser.
 */
contract FabricTokenFundraiserConfig is FabricTokenConfig {
    // The number of FT per 1 ETH.
    uint constant CONVERSION_RATE = 9000;

    // The public sale hard cap of the fundraiser.
    uint constant TOKENS_HARD_CAP = 71250 * (10**3) * DECIMALS_FACTOR;

    // The start date of the fundraiser: Thursday, 2018-02-15 10:00:00 UTC.
    uint constant START_DATE = 1518688800;

    // The end date of the fundraiser: Sunday, 2018-04-01 10:00:00 UTC (45 days after `START_DATE`).
    uint constant END_DATE = 1522576800;
    
    // Total number of tokens locked for the FT core team.
    uint constant TOKENS_LOCKED_CORE_TEAM = 12 * (10**6) * DECIMALS_FACTOR;

    // Total number of tokens locked for FT advisors.
    uint constant TOKENS_LOCKED_ADVISORS = 7 * (10**6) * DECIMALS_FACTOR;

    // The release date for tokens locked for the FT core team.
    uint constant TOKENS_LOCKED_CORE_TEAM_RELEASE_DATE = START_DATE + 1 years;

    // The release date for tokens locked for FT advisors.
    uint constant TOKENS_LOCKED_ADVISORS_RELEASE_DATE = START_DATE + 180 days;

    // Total number of tokens locked for bounty program.
    uint constant TOKENS_BOUNTY_PROGRAM = 1 * (10**6) * DECIMALS_FACTOR;

    // Maximum gas price limit
    uint constant MAX_GAS_PRICE = 50000000000 wei; // 50 gwei/shanon

    // Minimum individual contribution
    uint constant MIN_CONTRIBUTION =  0.1 ether;

    // Individual limit in ether
    uint constant INDIVIDUAL_ETHER_LIMIT =  9 ether;
}

// File: contracts\traits\TokenSafe.sol

/**
 * @title TokenSafe
 *
 * @dev A multi-bundle token safe contract that contains locked tokens released after a date for the specific bundle type.
 */
contract TokenSafe {
    using SafeMath for uint;

    struct AccountsBundle {
        // The total number of tokens locked.
        uint lockedTokens;
        // The release date for the locked tokens
        // Note: Unix timestamp fits uint32, however block.timestamp is uint
        uint releaseDate;
        // The balances for the FT locked token accounts.
        mapping (address => uint) balances;
    }

    // The account bundles of locked tokens grouped by release date
    mapping (uint8 => AccountsBundle) public bundles;

    // The `ERC20TokenInterface` contract.
    ERC20TokenInterface token;

    /**
     * @dev The constructor.
     *
     * @param _token The address of the Fabric Token (fundraiser) contract.
     */
    function TokenSafe(address _token) public {
        token = ERC20TokenInterface(_token);
    }

    /**
     * @dev The function initializes the bundle of accounts with a release date.
     *
     * @param _type Bundle type.
     * @param _releaseDate Unix timestamp of the time after which the tokens can be released
     */
    function initBundle(uint8 _type, uint _releaseDate) internal {
        bundles[_type].releaseDate = _releaseDate;
    }

    /**
     * @dev Add new account with locked token balance to the specified bundle type.
     *
     * @param _type Bundle type.
     * @param _account The address of the account to be added.
     * @param _balance The number of tokens to be locked.
     */
    function addLockedAccount(uint8 _type, address _account, uint _balance) internal {
        var bundle = bundles[_type];
        bundle.balances[_account] = bundle.balances[_account].plus(_balance);
        bundle.lockedTokens = bundle.lockedTokens.plus(_balance);
    }

    /**
     * @dev Allows an account to be released if it meets the time constraints.
     *
     * @param _type Bundle type.
     * @param _account The address of the account to be released.
     */
    function releaseAccount(uint8 _type, address _account) internal {
        var bundle = bundles[_type];
        require(now >= bundle.releaseDate);
        uint tokens = bundle.balances[_account];
        require(tokens > 0);
        bundle.balances[_account] = 0;
        bundle.lockedTokens = bundle.lockedTokens.minus(tokens);
        if (!token.transfer(_account, tokens)) {
            revert();
        }
    }
}

// File: contracts\FabricTokenSafe.sol

/**
 * @title FabricTokenSafe
 *
 * @dev The Fabric Token safe containing all details about locked tokens.
 */
contract FabricTokenSafe is TokenSafe, FabricTokenFundraiserConfig {
    // Bundle type constants
    uint8 constant CORE_TEAM = 0;
    uint8 constant ADVISORS = 1;

    /**
     * @dev The constructor.
     *
     * @param _token The address of the Fabric Token (fundraiser) contract.
     */
    function FabricTokenSafe(address _token) public
        TokenSafe(_token)
    {
        token = ERC20TokenInterface(_token);

        /// Core team.
        initBundle(CORE_TEAM,
            TOKENS_LOCKED_CORE_TEAM_RELEASE_DATE
        );

        // Accounts with tokens locked for the FT core team.
        addLockedAccount(CORE_TEAM, 0xB494096548aA049C066289A083204E923cBf4413, 4 * (10**6) * DECIMALS_FACTOR);
        addLockedAccount(CORE_TEAM, 0xE3506B01Bee377829ee3CffD8bae650e990c5d68, 4 * (10**6) * DECIMALS_FACTOR);
        addLockedAccount(CORE_TEAM, 0x3d13219dc1B8913E019BeCf0772C2a54318e5718, 4 * (10**6) * DECIMALS_FACTOR);

        // Verify that the tokens add up to the constant in the configuration.
        assert(bundles[CORE_TEAM].lockedTokens == TOKENS_LOCKED_CORE_TEAM);

        /// Advisors.
        initBundle(ADVISORS,
            TOKENS_LOCKED_ADVISORS_RELEASE_DATE
        );

        // Accounts with FT tokens locked for advisors.
        addLockedAccount(ADVISORS, 0x4647Da07dAAb17464278B988CDE59A4b911EBe44, 2 * (10**6) * DECIMALS_FACTOR);
        addLockedAccount(ADVISORS, 0x3eA2caac5A0A4a55f9e304AcD09b3CEe6cD4Bc39, 1 * (10**6) * DECIMALS_FACTOR);
        addLockedAccount(ADVISORS, 0xd5f791EC3ED79f79a401b12f7625E1a972382437, 1 * (10**6) * DECIMALS_FACTOR);
        addLockedAccount(ADVISORS, 0xcaeae3CD1a5d3E6E950424C994e14348ac3Ec5dA, 1 * (10**6) * DECIMALS_FACTOR);
        addLockedAccount(ADVISORS, 0xb6EA6193058F3c8A4A413d176891d173D62E00bE, 1 * (10**6) * DECIMALS_FACTOR);
        addLockedAccount(ADVISORS, 0x8b3E184Cf5C3bFDaB1C4D0F30713D30314FcfF7c, 1 * (10**6) * DECIMALS_FACTOR);

        // Verify that the tokens add up to the constant in the configuration.
        assert(bundles[ADVISORS].lockedTokens == TOKENS_LOCKED_ADVISORS);
    }

    /**
     * @dev Returns the total locked tokens. This function is called by the fundraiser to determine number of tokens to create upon finalization.
     *
     * @return The current total number of locked Fabric Tokens.
     */
    function totalTokensLocked() public constant returns (uint) {
        return bundles[CORE_TEAM].lockedTokens.plus(bundles[ADVISORS].lockedTokens);
    }

    /**
     * @dev Allows core team account FT tokens to be released.
     */
    function releaseCoreTeamAccount() public {
        releaseAccount(CORE_TEAM, msg.sender);
    }

    /**
     * @dev Allows advisors account FT tokens to be released.
     */
    function releaseAdvisorsAccount() public {
        releaseAccount(ADVISORS, msg.sender);
    }
}

// File: contracts\traits\Whitelist.sol

contract Whitelist is HasOwner
{
    // Whitelist mapping
    mapping(address => bool) public whitelist;

    /**
     * @dev The constructor.
     */
    function Whitelist(address _owner) public
        HasOwner(_owner)
    {

    }

    /**
     * @dev Access control modifier that allows only whitelisted address to call the method.
     */
    modifier onlyWhitelisted {
        require(whitelist[msg.sender]);
        _;
    }

    /**
     * @dev Internal function that sets whitelist status in batch.
     *
     * @param _entries An array with the entries to be updated
     * @param _status The new status to apply
     */
    function setWhitelistEntries(address[] _entries, bool _status) internal {
        for (uint32 i = 0; i < _entries.length; ++i) {
            whitelist[_entries[i]] = _status;
        }
    }

    /**
     * @dev Public function that allows the owner to whitelist multiple entries
     *
     * @param _entries An array with the entries to be whitelisted
     */
    function whitelistAddresses(address[] _entries) public onlyOwner {
        setWhitelistEntries(_entries, true);
    }

    /**
     * @dev Public function that allows the owner to blacklist multiple entries
     *
     * @param _entries An array with the entries to be blacklist
     */
    function blacklistAddresses(address[] _entries) public onlyOwner {
        setWhitelistEntries(_entries, false);
    }
}

// File: contracts\FabricTokenFundraiser.sol

/**
 * @title FabricTokenFundraiser
 *
 * @dev The Fabric Token fundraiser contract.
 */
contract FabricTokenFundraiser is FabricToken, FabricTokenFundraiserConfig, Whitelist {
    // Indicates whether the fundraiser has ended or not.
    bool public finalized = false;

    // The address of the account which will receive the funds gathered by the fundraiser.
    address public beneficiary;

    // The number of FT participants will receive per 1 ETH.
    uint public conversionRate;

    // Fundraiser start date.
    uint public startDate;

    // Fundraiser end date.
    uint public endDate;

    // Fundraiser tokens hard cap.
    uint public hardCap;

    // The `FabricTokenSafe` contract.
    FabricTokenSafe public fabricTokenSafe;

    // The minimum amount of ether allowed in the public sale
    uint internal minimumContribution;

    // The maximum amount of ether allowed per address
    uint internal individualLimit;

    // Number of tokens sold during the fundraiser.
    uint private tokensSold;

    // Indicates whether the tokens are claimed by the partners
    bool private partnerTokensClaimed = false;

    /**
     * @dev The event fires every time a new buyer enters the fundraiser.
     *
     * @param _address The address of the buyer.
     * @param _ethers The number of ethers sent.
     * @param _tokens The number of tokens received by the buyer.
     * @param _newTotalSupply The updated total number of tokens currently in circulation.
     * @param _conversionRate The conversion rate at which the tokens were bought.
     */
    event FundsReceived(address indexed _address, uint _ethers, uint _tokens, uint _newTotalSupply, uint _conversionRate);

    /**
     * @dev The event fires when the beneficiary of the fundraiser is changed.
     *
     * @param _beneficiary The address of the new beneficiary.
     */
    event BeneficiaryChange(address _beneficiary);

    /**
     * @dev The event fires when the number of FT per 1 ETH is changed.
     *
     * @param _conversionRate The new number of FT per 1 ETH.
     */
    event ConversionRateChange(uint _conversionRate);

    /**
     * @dev The event fires when the fundraiser is successfully finalized.
     *
     * @param _beneficiary The address of the beneficiary.
     * @param _ethers The number of ethers transfered to the beneficiary.
     * @param _totalSupply The total number of tokens in circulation.
     */
    event Finalized(address _beneficiary, uint _ethers, uint _totalSupply);

    /**
     * @dev The constructor.
     *
     * @param _beneficiary The address which will receive the funds gathered by the fundraiser.
     */
    function FabricTokenFundraiser(address _beneficiary) public
        FabricToken(0)
        Whitelist(msg.sender)
    {
        require(_beneficiary != 0);

        beneficiary = _beneficiary;
        conversionRate = CONVERSION_RATE;
        startDate = START_DATE;
        endDate = END_DATE;
        hardCap = TOKENS_HARD_CAP;
        tokensSold = 0;
        minimumContribution = MIN_CONTRIBUTION;
        individualLimit = INDIVIDUAL_ETHER_LIMIT * CONVERSION_RATE;

        fabricTokenSafe = new FabricTokenSafe(this);

        // Freeze the transfers for the duration of the fundraiser.
        freeze();
    }

    /**
     * @dev Changes the beneficiary of the fundraiser.
     *
     * @param _beneficiary The address of the new beneficiary.
     */
    function setBeneficiary(address _beneficiary) public onlyOwner {
        require(_beneficiary != 0);

        beneficiary = _beneficiary;

        BeneficiaryChange(_beneficiary);
    }

    /**
     * @dev Sets converstion rate of 1 ETH to FT. Can only be changed before the fundraiser starts.
     *
     * @param _conversionRate The new number of Fabric Tokens per 1 ETH.
     */
    function setConversionRate(uint _conversionRate) public onlyOwner {
        require(now < startDate);
        require(_conversionRate > 0);

        conversionRate = _conversionRate;
        individualLimit = INDIVIDUAL_ETHER_LIMIT * _conversionRate;

        ConversionRateChange(_conversionRate);
    }

    /**
     * @dev The default function which will fire every time someone sends ethers to this contract's address.
     */
    function() public payable {
        buyTokens();
    }

    /**
     * @dev Creates new tokens based on the number of ethers sent and the conversion rate.
     */
    function buyTokens() public payable onlyWhitelisted {
        require(!finalized);
        require(now >= startDate);
        require(now <= endDate);
        require(tx.gasprice <= MAX_GAS_PRICE);  // gas price limit
        require(msg.value >= minimumContribution);  // required minimum contribution
        require(tokensSold <= hardCap);

        // Calculate the number of tokens the buyer will receive.
        uint tokens = msg.value.mul(conversionRate);
        balances[msg.sender] = balances[msg.sender].plus(tokens);

        // Ensure that the individual contribution limit has not been reached
        require(balances[msg.sender] <= individualLimit);

        tokensSold = tokensSold.plus(tokens);
        totalSupply = totalSupply.plus(tokens);

        Transfer(0x0, msg.sender, tokens);

        FundsReceived(
            msg.sender,
            msg.value, 
            tokens, 
            totalSupply, 
            conversionRate
        );
    }

    /**
     * @dev Distributes the tokens allocated for the strategic partners.
     */
    function claimPartnerTokens() public {
        require(!partnerTokensClaimed);
        require(now >= startDate);

        partnerTokensClaimed = true;

        address partner1 = 0xA6556B9BD0AAbf0d8824374A3C425d315b09b832;
        balances[partner1] = balances[partner1].plus(125 * (10**4) * DECIMALS_FACTOR);

        address partner2 = 0x783A1cBc37a8ef2F368908490b72BfE801DA1877;
        balances[partner2] = balances[partner2].plus(750 * (10**4) * DECIMALS_FACTOR);

        totalSupply = totalSupply.plus(875 * (10**4) * DECIMALS_FACTOR);
    }

    /**
     * @dev Finalize the fundraiser if `endDate` has passed or if `hardCap` is reached.
     */
    function finalize() public onlyOwner {
        require((totalSupply >= hardCap) || (now >= endDate));
        require(!finalized);

        Finalized(beneficiary, this.balance, totalSupply);

        /// Send the total number of ETH gathered to the beneficiary.
        beneficiary.transfer(this.balance);

        /// Allocate locked tokens to the `FabricTokenSafe` contract.
        uint totalTokensLocked = fabricTokenSafe.totalTokensLocked();
        balances[address(fabricTokenSafe)] = balances[address(fabricTokenSafe)].plus(totalTokensLocked);
        totalSupply = totalSupply.plus(totalTokensLocked);

        // Transfer the funds for the bounty program.
        balances[owner] = balances[owner].plus(TOKENS_BOUNTY_PROGRAM);
        totalSupply = totalSupply.plus(TOKENS_BOUNTY_PROGRAM);

        /// Finalize the fundraiser. Keep in mind that this cannot be undone.
        finalized = true;

        // Unfreeze transfers
        unfreeze();
    }
}