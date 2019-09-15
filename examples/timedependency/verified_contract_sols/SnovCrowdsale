pragma solidity ^0.4.11;

contract SafeMath {

    function safeMul(uint256 a, uint256 b) internal constant returns (uint256 ) {
        uint256 c = a * b;
        assert(a == 0 || c / a == b);
        return c;
    }

    function safeDiv(uint256 a, uint256 b) internal constant returns (uint256 ) {
        assert(b > 0);
        uint256 c = a / b;
        assert(a == b * c + a % b);
        return c;
    }

    function safeSub(uint256 a, uint256 b) internal constant returns (uint256 ) {
        assert(b <= a);
        return a - b;
    }

    function safeAdd(uint256 a, uint256 b) internal constant returns (uint256 ) {
        uint256 c = a + b;
        assert(c >= a);
        return c;
    }
}

contract ERC20 {

    /* This is a slight change to the ERC20 base standard.
    function totalSupply() constant returns (uint256 supply);
    is replaced with:
    uint256 public totalSupply;
    This automatically creates a getter function for the totalSupply.
    This is moved to the base contract since public getter functions are not
    currently recognised as an implementation of the matching abstract
    function by the compiler.
    */
    /// total amount of tokens
    uint256 public totalSupply;

    /// @param _owner The address from which the balance will be retrieved
    /// @return The balance
    function balanceOf(address _owner) constant returns (uint256 balance);

    /// @notice send `_value` token to `_to` from `msg.sender`
    /// @param _to The address of the recipient
    /// @param _value The amount of token to be transferred
    /// @return Whether the transfer was successful or not
    function transfer(address _to, uint256 _value) returns (bool success);

    /// @notice send `_value` token to `_to` from `_from` on the condition it is approved by `_from`
    /// @param _from The address of the sender
    /// @param _to The address of the recipient
    /// @param _value The amount of token to be transferred
    /// @return Whether the transfer was successful or not
    function transferFrom(address _from, address _to, uint256 _value) returns (bool success);

    /// @notice `msg.sender` approves `_spender` to spend `_value` tokens
    /// @param _spender The address of the account able to transfer the tokens
    /// @param _value The amount of tokens to be approved for transfer
    /// @return Whether the approval was successful or not
    function approve(address _spender, uint256 _value) returns (bool success);

    /// @param _owner The address of the account owning tokens
    /// @param _spender The address of the account able to transfer the tokens
    /// @return Amount of remaining tokens allowed to spent
    function allowance(address _owner, address _spender) constant returns (uint256 remaining);

    event Transfer(address indexed _from, address indexed _to, uint256 _value);
    event Approval(address indexed _owner, address indexed _spender, uint256 _value);
}

contract StandardToken is ERC20, SafeMath {

    mapping(address => uint256) balances;
    mapping(address => mapping(address => uint256)) allowed;

    /// @dev Returns number of tokens owned by given address.
    /// @param _owner Address of token owner.
    function balanceOf(address _owner) constant returns (uint256 balance) {
        return balances[_owner];
    }

    /// @dev Transfers sender's tokens to a given address. Returns success.
    /// @param _to Address of token receiver.
    /// @param _value Number of tokens to transfer.
    function transfer(address _to, uint256 _value) returns (bool) {
        if (balances[msg.sender] >= _value && _value > 0) {
            balances[msg.sender] = safeSub(balances[msg.sender], _value);
            balances[_to] = safeAdd(balances[_to], _value);
            Transfer(msg.sender, _to, _value);
            return true;
        } else return false;
    }

    /// @dev Allows allowed third party to transfer tokens from one address to another. Returns success.
    /// @param _from Address from where tokens are withdrawn.
    /// @param _to Address to where tokens are sent.
    /// @param _value Number of tokens to transfer.
    function transferFrom(address _from, address _to, uint256 _value) returns (bool) {
        if (balances[_from] >= _value && allowed[_from][msg.sender] >= _value && _value > 0) {
            balances[_to] = safeAdd(balances[_to], _value);
            balances[_from] = safeSub(balances[_from], _value);
            allowed[_from][msg.sender] = safeSub(allowed[_from][msg.sender], _value);
            Transfer(_from, _to, _value);
            return true;
        } else return false;
    }

    /// @dev Sets approved amount of tokens for spender. Returns success.
    /// @param _spender Address of allowed account.
    /// @param _value Number of approved tokens.
    function approve(address _spender, uint256 _value) returns (bool) {
        allowed[msg.sender][_spender] = _value;
        Approval(msg.sender, _spender, _value);
        return true;
    }

    /// @dev Returns number of allowed tokens for given address.
    /// @param _owner Address of token owner.
    /// @param _spender Address of token spender.
    function allowance(address _owner, address _spender) constant returns (uint256 remaining) {
        return allowed[_owner][_spender];
    }
}

contract Ownable {

    address public owner;
    address public pendingOwner;

    function Ownable() {
        owner = msg.sender;
    }

    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    // Safe transfer of ownership in 2 steps. Once called, a newOwner needs to call claimOwnership() to prove ownership.
    function transferOwnership(address newOwner) onlyOwner {
        pendingOwner = newOwner;
    }

    function claimOwnership() {
        if (msg.sender == pendingOwner) {
            owner = pendingOwner;
            pendingOwner = 0;
        }
    }
}

contract MultiOwnable {

    mapping (address => bool) ownerMap;
    address[] public owners;

    event OwnerAdded(address indexed _newOwner);
    event OwnerRemoved(address indexed _oldOwner);

    modifier onlyOwner() {
        require(isOwner(msg.sender));
        _;
    }

    function MultiOwnable() {
        // Add default owner
        address owner = msg.sender;
        ownerMap[owner] = true;
        owners.push(owner);
    }

    function ownerCount() public constant returns (uint256) {
        return owners.length;
    }

    function isOwner(address owner) public constant returns (bool) {
        return ownerMap[owner];
    }

    function addOwner(address owner) onlyOwner returns (bool) {
        if (!isOwner(owner) && owner != 0) {
            ownerMap[owner] = true;
            owners.push(owner);

            OwnerAdded(owner);
            return true;
        } else return false;
    }

    function removeOwner(address owner) onlyOwner returns (bool) {
        if (isOwner(owner)) {
            ownerMap[owner] = false;
            for (uint i = 0; i < owners.length - 1; i++) {
                if (owners[i] == owner) {
                    owners[i] = owners[owners.length - 1];
                    break;
                }
            }
            owners.length -= 1;

            OwnerRemoved(owner);
            return true;
        } else return false;
    }
}

contract Pausable is Ownable {

    bool public paused;

    modifier ifNotPaused {
        require(!paused);
        _;
    }

    modifier ifPaused {
        require(paused);
        _;
    }

    // Called by the owner on emergency, triggers paused state
    function pause() external onlyOwner {
        paused = true;
    }

    // Called by the owner on end of emergency, returns to normal state
    function resume() external onlyOwner ifPaused {
        paused = false;
    }
}

contract TokenSpender {
    function receiveApproval(address _from, uint256 _value);
}

contract BsToken is StandardToken, MultiOwnable {

    bool public locked;

    string public name;
    string public symbol;
    uint256 public totalSupply;
    uint8 public decimals = 18;
    string public version = 'v0.1';

    address public creator;
    address public seller;
    uint256 public tokensSold;
    uint256 public totalSales;

    event Sell(address indexed _seller, address indexed _buyer, uint256 _value);
    event SellerChanged(address indexed _oldSeller, address indexed _newSeller);

    modifier onlyUnlocked() {
        if (!isOwner(msg.sender) && locked) throw;
        _;
    }

    function BsToken(string _name, string _symbol, uint256 _totalSupplyNoDecimals, address _seller) MultiOwnable() {

        // Lock the transfer function during the presale/crowdsale to prevent speculations.
        locked = true;

        creator = msg.sender;
        seller = _seller;

        name = _name;
        symbol = _symbol;
        totalSupply = _totalSupplyNoDecimals * 1e18;

        balances[seller] = totalSupply;
        Transfer(0x0, seller, totalSupply);
    }

    function changeSeller(address newSeller) onlyOwner returns (bool) {
        require(newSeller != 0x0 && seller != newSeller);

        address oldSeller = seller;

        uint256 unsoldTokens = balances[oldSeller];
        balances[oldSeller] = 0;
        balances[newSeller] = safeAdd(balances[newSeller], unsoldTokens);
        Transfer(oldSeller, newSeller, unsoldTokens);

        seller = newSeller;
        SellerChanged(oldSeller, newSeller);
        return true;
    }

    function sellNoDecimals(address _to, uint256 _value) returns (bool) {
        return sell(_to, _value * 1e18);
    }

    function sell(address _to, uint256 _value) onlyOwner returns (bool) {
        if (balances[seller] >= _value && _value > 0) {
            balances[seller] = safeSub(balances[seller], _value);
            balances[_to] = safeAdd(balances[_to], _value);
            Transfer(seller, _to, _value);

            tokensSold = safeAdd(tokensSold, _value);
            totalSales = safeAdd(totalSales, 1);
            Sell(seller, _to, _value);
            return true;
        } else return false;
    }

    function transfer(address _to, uint256 _value) onlyUnlocked returns (bool) {
        return super.transfer(_to, _value);
    }

    function transferFrom(address _from, address _to, uint256 _value) onlyUnlocked returns (bool) {
        return super.transferFrom(_from, _to, _value);
    }

    function lock() onlyOwner {
        locked = true;
    }

    function unlock() onlyOwner {
        locked = false;
    }

    function burn(uint256 _value) returns (bool) {
        if (balances[msg.sender] >= _value && _value > 0) {
            balances[msg.sender] = safeSub(balances[msg.sender], _value) ;
            totalSupply = safeSub(totalSupply, _value);
            Transfer(msg.sender, 0x0, _value);
            return true;
        } else return false;
    }

    /* Approve and then communicate the approved contract in a single tx */
    function approveAndCall(address _spender, uint256 _value) {
        TokenSpender spender = TokenSpender(_spender);
        if (approve(_spender, _value)) {
            spender.receiveApproval(msg.sender, _value);
        }
    }
}

/**
 * In this crowdsale We assume that ETH rate is 320 USD/ETH
 */
contract BsCrowdsale is SafeMath, Ownable, Pausable {

    enum Currency { BTC, LTC, DASH, ZEC, WAVES, USD, EUR }

    struct Backer {
        uint256 weiReceived; // Amount of wei given by backer
        uint256 tokensSent;  // Amount of tokens received in return to the given amount of ETH.
    }

    // TODO rename to buyers?
    // (buyer_eth_address -> struct)
    mapping(address => Backer) public backers;

    // currency_code => (tx_hash => tokens)
    mapping(uint => mapping(bytes32 => uint256)) public externalTxs;

    BsToken public token;           // Token contract reference.
    address public beneficiary;     // Address that will receive ETH raised during this crowdsale.
    address public notifier;        // Address that can this crowdsale about changed external conditions.

    uint256 public usdPerEth;
    uint256 public usdPerEthMin = 200; // Lowest boundary of USD/ETH rate
    uint256 public usdPerEthMax = 500; // Highest boundary of USD/ETH rate

    struct UsdPerEthLog {
        uint256 rate;
        uint256 time;
        address changedBy;
    }

    UsdPerEthLog[] public usdPerEthLog; // History of used rates of USD/ETH

    uint256 public minInvestCents = 1;              // Because 1 token = 1 cent.
    uint256 public maxCapInCents  = 10 * 1e6 * 100; // 10m USD in cents.

    uint256 public tokensPerCents = 1 * 1e18;       // Ordinary price is 1 token = 1 USD cent.
    uint256 public tokensPerCentsDayOne = 1.1 * 1e18;
    uint256 public tokensPerCentsWeekOne = 1.05 * 1e18;

    // TODO set to last value from SnovWhitelist contract.
    uint256 public totalInCents = 71481400; // Total amount of USD raised during this crowdsale including (wei -> USD) + (external txs in USD).

    // Stats for current crowdsale
    uint256 public totalWeiReceived = 0;   // Total amount of wei received during this crowdsale smart contract.
    uint256 public totalTokensSold = 0;    // Total amount of tokens sold during this crowdsale.
    uint256 public totalEthSales = 0;      // Total amount of ETH contributions during this crowdsale.
    uint256 public totalExternalSales = 0; // Total amount of external contributions (USD, BTC, etc.) during this crowdsale.

    uint256 public startTime = 1509451200; // 2017-10-31 12:00:00Z
    uint256 public endTime   = 1512043200; // 2017-11-30 12:00:00Z

    // Use in bonuses:
    uint256 oneDayTime       = 1509537600; // 2017-11-01 12:00:00Z
    uint256 oneWeekTime      = 1510056000; // 2017-11-07 12:00:00Z

    uint256 public finalizedTime = 0;      // Unix timestamp when finalize() was called.

    bool public saleEnabled = true;       // if true, then contract will not sell tokens on payment received

    event BeneficiaryChanged(address indexed _oldAddress, address indexed _newAddress);
    event NotifierChanged(address indexed _oldAddress, address indexed _newAddress);
    event UsdPerEthChanged(uint256 _oldRate, uint256 _newRate);

    event EthReceived(address indexed _buyer, uint256 _amountInWei);
    event ExternalSale(Currency _currency, string _txHash, address indexed _buyer, uint256 _amountInCents, uint256 _tokensE18);

    modifier respectTimeFrame() {
        require(isSaleOn());
        _;
    }

    modifier canNotify() {
        require(msg.sender == owner || msg.sender == notifier);
        _;
    }

    function BsCrowdsale(address _token, address _beneficiary, uint256 _usdPerEth) {
        token = BsToken(_token);

        owner = msg.sender;
        notifier = 0x73E5B12017A141d41c1a14FdaB43a54A4f9BD490;
        beneficiary = _beneficiary;

        setUsdPerEth(_usdPerEth);
    }

    // Override this method to mock current time.
    function getNow() public constant returns (uint256) {
        return now;
    }

    function setSaleEnabled(bool _enabled) public onlyOwner {
        saleEnabled = _enabled;
    }

    function setBeneficiary(address _beneficiary) public onlyOwner {
        BeneficiaryChanged(beneficiary, _beneficiary);
        beneficiary = _beneficiary;
    }

    function setNotifier(address _notifier) public onlyOwner {
        NotifierChanged(notifier, _notifier);
        notifier = _notifier;
    }

    function setUsdPerEth(uint256 _usdPerEth) public canNotify {
        if (_usdPerEth < usdPerEthMin || _usdPerEth > usdPerEthMax) throw;

        UsdPerEthChanged(usdPerEth, _usdPerEth);
        usdPerEth = _usdPerEth;
        usdPerEthLog.push(UsdPerEthLog({ rate: usdPerEth, time: getNow(), changedBy: msg.sender }));
    }

    function usdPerEthLogSize() public constant returns (uint256) {
        return usdPerEthLog.length;
    }

    /*
     * The fallback function corresponds to a donation in ETH
     */
    function() public payable {
        if (saleEnabled) sellTokensForEth(msg.sender, msg.value);
    }

    function sellTokensForEth(address _buyer, uint256 _amountInWei) internal ifNotPaused respectTimeFrame {

        uint256 amountInCents = weiToCents(_amountInWei);
        require(amountInCents >= minInvestCents);

        totalInCents = safeAdd(totalInCents, amountInCents);
        require(totalInCents <= maxCapInCents); // If max cap reached.

        uint256 tokensSold = centsToTokens(amountInCents);
        require(token.sell(_buyer, tokensSold)); // Transfer tokens to buyer.

        totalWeiReceived = safeAdd(totalWeiReceived, _amountInWei);
        totalTokensSold = safeAdd(totalTokensSold, tokensSold);
        totalEthSales++;

        Backer backer = backers[_buyer];
        backer.tokensSent = safeAdd(backer.tokensSent, tokensSold);
        backer.weiReceived = safeAdd(backer.weiReceived, _amountInWei);  // Update the total wei collected during the crowdfunding for this backer

        EthReceived(_buyer, _amountInWei);
    }

    function weiToCents(uint256 _amountInWei) public constant returns (uint256) {
        return safeDiv(safeMul(_amountInWei, usdPerEth * 100), 1 ether);
    }

    function centsToTokens(uint256 _amountInCents) public constant returns (uint256) {
        uint256 rate = tokensPerCents;
        uint _now = getNow();

        if (startTime <= _now && _now < oneDayTime) rate = tokensPerCentsDayOne;
        else if (oneDayTime <= _now && _now < oneWeekTime) rate = tokensPerCentsWeekOne;

        return safeMul(_amountInCents, rate);
    }

    function externalSale(
        Currency _currency,
        string _txHash,
        address _buyer,
        uint256 _amountInCents,
        uint256 _tokensE18
    ) internal ifNotPaused canNotify {

        require(_buyer > 0 && _amountInCents > 0 && _tokensE18 > 0);

        var txsByCur = externalTxs[uint(_currency)];
        bytes32 txKey = keccak256(_txHash);

        // If this foreign transaction has been already processed in this contract.
        require(txsByCur[txKey] == 0);

        totalInCents = safeAdd(totalInCents, _amountInCents);
        require(totalInCents < maxCapInCents); // Max cap should not be reached yet.

        require(token.sell(_buyer, _tokensE18)); // Transfer tokens to buyer.

        totalTokensSold = safeAdd(totalTokensSold, _tokensE18);
        totalExternalSales++;

        txsByCur[txKey] = _tokensE18;
        ExternalSale(_currency, _txHash, _buyer, _amountInCents, _tokensE18);
    }

    function sellTokensForBtc(string _txHash, address _buyer, uint256 _amountInCents, uint256 _tokensE18) public {
        externalSale(Currency.BTC, _txHash, _buyer, _amountInCents, _tokensE18);
    }

    function sellTokensForLtc(string _txHash, address _buyer, uint256 _amountInCents, uint256 _tokensE18) public {
        externalSale(Currency.LTC, _txHash, _buyer, _amountInCents, _tokensE18);
    }

    function sellTokensForDash(string _txHash, address _buyer, uint256 _amountInCents, uint256 _tokensE18) public {
        externalSale(Currency.DASH, _txHash, _buyer, _amountInCents, _tokensE18);
    }

    function sellTokensForZec(string _txHash, address _buyer, uint256 _amountInCents, uint256 _tokensE18) public {
        externalSale(Currency.ZEC, _txHash, _buyer, _amountInCents, _tokensE18);
    }

    function sellTokensForWaves(string _txHash, address _buyer, uint256 _amountInCents, uint256 _tokensE18) public {
        externalSale(Currency.WAVES, _txHash, _buyer, _amountInCents, _tokensE18);
    }

    function sellTokensForUsd(string _txHash, address _buyer, uint256 _amountInCents, uint256 _tokensE18) public {
        externalSale(Currency.USD, _txHash, _buyer, _amountInCents, _tokensE18);
    }

    function sellTokensForEur(string _txHash, address _buyer, uint256 _amountInCents, uint256 _tokensE18) public {
        externalSale(Currency.EUR, _txHash, _buyer, _amountInCents, _tokensE18);
    }

    function tokensByExternalTx(Currency _currency, string _txHash) internal constant returns (uint256) {
        return externalTxs[uint(_currency)][keccak256(_txHash)];
    }

    function tokensByBtcTx(string _txHash) public constant returns (uint256) {
        return tokensByExternalTx(Currency.BTC, _txHash);
    }

    function tokensByLtcTx(string _txHash) public constant returns (uint256) {
        return tokensByExternalTx(Currency.LTC, _txHash);
    }

    function tokensByDashTx(string _txHash) public constant returns (uint256) {
        return tokensByExternalTx(Currency.DASH, _txHash);
    }

    function tokensByZecTx(string _txHash) public constant returns (uint256) {
        return tokensByExternalTx(Currency.ZEC, _txHash);
    }

    function tokensByWavesTx(string _txHash) public constant returns (uint256) {
        return tokensByExternalTx(Currency.WAVES, _txHash);
    }

    function tokensByUsdTx(string _txHash) public constant returns (uint256) {
        return tokensByExternalTx(Currency.USD, _txHash);
    }

    function tokensByEurTx(string _txHash) public constant returns (uint256) {
        return tokensByExternalTx(Currency.EUR, _txHash);
    }

    function totalSales() public constant returns (uint256) {
        return safeAdd(totalEthSales, totalExternalSales);
    }

    function isMaxCapReached() public constant returns (bool) {
        return totalInCents >= maxCapInCents;
    }

    function isSaleOn() public constant returns (bool) {
        uint256 _now = getNow();
        return startTime <= _now && _now <= endTime;
    }

    function isSaleOver() public constant returns (bool) {
        return getNow() > endTime;
    }

    function isFinalized() public constant returns (bool) {
        return finalizedTime > 0;
    }

    /*
     * Finalize the crowdsale. Raised money can be sent to beneficiary only if crowdsale hit end time or max cap (15m USD).
     */
    function finalize() public onlyOwner {

        // Cannot finalise before end day of crowdsale until max cap is reached.
        require(isMaxCapReached() || isSaleOver());

        beneficiary.transfer(this.balance);

        finalizedTime = getNow();
    }
}

contract SnovCrowdsale is BsCrowdsale {

    function SnovCrowdsale() BsCrowdsale(
        0xBDC5bAC39Dbe132B1E030e898aE3830017D7d969,
        0x983F64a550CD9D733f2829275f94B1A3728Fe888,
        310
    ) {}
}