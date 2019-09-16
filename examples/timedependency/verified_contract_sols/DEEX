/**
 *Submitted for verification at Etherscan.io on 2017-10-26
*/

pragma solidity ^0.4.15;


/*
*  deex.exchange pre-ICO tokens smart contract
*  implements [ERC-20 Token Standard](https://github.com/ethereum/EIPs/blob/master/EIPS/eip-20-token-standard.md)
*
*  Style
*  1) before start coding, run Python and type 'import this' in Python console.
*  2) we avoid using inheritance (contract B is A) as it makes code less clear for observer
*  ("Flat is better than nested", "Readability counts")
*  3) we avoid using -= ; =- ; +=; =+
*  see: https://github.com/ether-camp/virtual-accelerator/issues/8
*  https://www.ethnews.com/ethercamps-hkg-token-has-a-bug-and-needs-to-be-reissued
*  4) always explicitly mark variables and functions visibility ("Explicit is better than implicit")
*  5) every function except constructor should trigger at leas one event.
*  6) smart contracts have to be audited and reviewed, comment your code.
*
*  Code is published on https://github.com/thedeex/thedeex.github.io
*/


/* "Interfaces" */

//  this is expected from another contracts
//  if it wants to spend tokens of behalf of the token owner in our contract
//  this can be used in many situations, for example to convert pre-ICO tokens to ICO tokens
//  see 'approveAndCall' function
contract allowanceRecipient {
    function receiveApproval(address _from, uint256 _value, address _inContract, bytes _extraData) returns (bool success);
}


// see:
// https://github.com/ethereum/EIPs/issues/677
contract tokenRecipient {
    function tokenFallback(address _from, uint256 _value, bytes _extraData) returns (bool success);
}


contract DEEX {

    // ver. 2.0

    /* ---------- Variables */

    /* --- ERC-20 variables */

    // https://github.com/ethereum/EIPs/blob/master/EIPS/eip-20-token-standard.md#name
    // function name() constant returns (string name)
    string public name = "deex";

    // https://github.com/ethereum/EIPs/blob/master/EIPS/eip-20-token-standard.md#symbol
    // function symbol() constant returns (string symbol)
    string public symbol = "deex";

    // https://github.com/ethereum/EIPs/blob/master/EIPS/eip-20-token-standard.md#decimals
    // function decimals() constant returns (uint8 decimals)
    uint8 public decimals = 0;

    // https://github.com/ethereum/EIPs/blob/master/EIPS/eip-20-token-standard.md#totalsupply
    // function totalSupply() constant returns (uint256 totalSupply)
    // we start with zero and will create tokens as SC receives ETH
    uint256 public totalSupply;

    // https://github.com/ethereum/EIPs/blob/master/EIPS/eip-20-token-standard.md#balanceof
    // function balanceOf(address _owner) constant returns (uint256 balance)
    mapping (address => uint256) public balanceOf;

    // https://github.com/ethereum/EIPs/blob/master/EIPS/eip-20-token-standard.md#allowance
    // function allowance(address _owner, address _spender) constant returns (uint256 remaining)
    mapping (address => mapping (address => uint256)) public allowance;

    /* ----- For tokens sale */

    uint256 public salesCounter = 0;

    uint256 public maxSalesAllowed;

    bool private transfersBetweenSalesAllowed;

    // initial value should be changed by the owner
    uint256 public tokenPriceInWei = 0;

    uint256 public saleStartUnixTime = 0; // block.timestamp
    uint256 public saleEndUnixTime = 0;  // block.timestamp

    /* --- administrative */
    address public owner;

    // account that can set prices
    address public priceSetter;

    // 0 - not set
    uint256 private priceMaxWei = 0;
    // 0 - not set
    uint256 private priceMinWei = 0;

    // accounts holding tokens for for the team, for advisers and for the bounty campaign
    mapping (address => bool) public isPreferredTokensAccount;

    bool public contractInitialized = false;


    /* ---------- Constructor */
    // do not forget about:
    // https://medium.com/@codetractio/a-look-into-paritys-multisig-wallet-bug-affecting-100-million-in-ether-and-tokens-356f5ba6e90a
    function DEEX() {
        owner = msg.sender;

        // for testNet can be more than 2
        // --------------------------------2------------------------------------------------------change  in production!
        maxSalesAllowed = 2;
        //
        transfersBetweenSalesAllowed = true;
    }


    function initContract(address team, address advisers, address bounty) public onlyBy(owner) returns (bool){

        require(contractInitialized == false);
        contractInitialized = true;

        priceSetter = msg.sender;

        totalSupply = 100000000;

        // tokens for sale go SC own account
        balanceOf[this] = 75000000;

        // for the team
        balanceOf[team] = balanceOf[team] + 15000000;
        isPreferredTokensAccount[team] = true;

        // for advisers
        balanceOf[advisers] = balanceOf[advisers] + 7000000;
        isPreferredTokensAccount[advisers] = true;

        // for the bounty campaign
        balanceOf[bounty] = balanceOf[bounty] + 3000000;
        isPreferredTokensAccount[bounty] = true;

    }

    /* ---------- Events */

    /* --- ERC-20 events */
    // https://github.com/ethereum/EIPs/blob/master/EIPS/eip-20-token-standard.md#events

    // https://github.com/ethereum/EIPs/blob/master/EIPS/eip-20-token-standard.md#transfer-1
    event Transfer(address indexed from, address indexed to, uint256 value);

    // https://github.com/ethereum/EIPs/blob/master/EIPS/eip-20-token-standard.md#approval
    event Approval(address indexed _owner, address indexed spender, uint256 value);

    /* --- Administrative events:  */
    event OwnerChanged(address indexed oldOwner, address indexed newOwner);

    /* ---- Tokens creation and sale events  */

    event PriceChanged(uint256 indexed newTokenPriceInWei);

    event SaleStarted(uint256 startUnixTime, uint256 endUnixTime, uint256 indexed saleNumber);

    event NewTokensSold(uint256 numberOfTokens, address indexed purchasedBy, uint256 indexed priceInWei);

    event Withdrawal(address indexed to, uint sumInWei);

    /* --- Interaction with other contracts events  */
    event DataSentToAnotherContract(address indexed _from, address indexed _toContract, bytes _extraData);

    /* ---------- Functions */

    /* --- Modifiers  */
    modifier onlyBy(address _account){
        require(msg.sender == _account);

        _;
    }

    /* --- ERC-20 Functions */
    // https://github.com/ethereum/EIPs/blob/master/EIPS/eip-20-token-standard.md#methods

    // https://github.com/ethereum/EIPs/blob/master/EIPS/eip-20-token-standard.md#transfer
    function transfer(address _to, uint256 _value) public returns (bool){
        return transferFrom(msg.sender, _to, _value);
    }

    // https://github.com/ethereum/EIPs/blob/master/EIPS/eip-20-token-standard.md#transferfrom
    function transferFrom(address _from, address _to, uint256 _value) public returns (bool){

        // transfers are possible only after sale is finished
        // except for manager and preferred accounts

        bool saleFinished = saleIsFinished();
        require(saleFinished || msg.sender == owner || isPreferredTokensAccount[msg.sender]);

        // transfers can be forbidden until final ICO is finished
        // except for manager and preferred accounts
        require(transfersBetweenSalesAllowed || salesCounter == maxSalesAllowed || msg.sender == owner || isPreferredTokensAccount[msg.sender]);

        // Transfers of 0 values MUST be treated as normal transfers and fire the Transfer event (ERC-20)
        require(_value >= 0);

        // The function SHOULD throw unless the _from account has deliberately authorized the sender of the message via some mechanism
        require(msg.sender == _from || _value <= allowance[_from][msg.sender]);

        // check if _from account have required amount
        require(_value <= balanceOf[_from]);

        // Subtract from the sender
        balanceOf[_from] = balanceOf[_from] - _value;
        //
        // Add the same to the recipient
        balanceOf[_to] = balanceOf[_to] + _value;

        // If allowance used, change allowances correspondingly
        if (_from != msg.sender) {
            allowance[_from][msg.sender] = allowance[_from][msg.sender] - _value;
        }

        // event
        Transfer(_from, _to, _value);

        return true;
    }

    // https://github.com/ethereum/EIPs/blob/master/EIPS/eip-20-token-standard.md#approve
    // there is and attack, see:
    // https://github.com/CORIONplatform/solidity/issues/6,
    // https://drive.google.com/file/d/0ByMtMw2hul0EN3NCaVFHSFdxRzA/view
    // but this function is required by ERC-20
    function approve(address _spender, uint256 _value) public returns (bool success){

        require(_value >= 0);

        allowance[msg.sender][_spender] = _value;

        // event
        Approval(msg.sender, _spender, _value);

        return true;
    }

    /*  ---------- Interaction with other contracts  */

    /* User can allow another smart contract to spend some shares in his behalf
    *  (this function should be called by user itself)
    *  @param _spender another contract's address
    *  @param _value number of tokens
    *  @param _extraData Data that can be sent from user to another contract to be processed
    *  bytes - dynamically-sized byte array,
    *  see http://solidity.readthedocs.io/en/v0.4.15/types.html#dynamically-sized-byte-array
    *  see possible attack information in comments to function 'approve'
    *  > this may be used to convert pre-ICO tokens to ICO tokens
    */
    function approveAndCall(address _spender, uint256 _value, bytes _extraData) public returns (bool success) {

        approve(_spender, _value);

        // 'spender' is another contract that implements code as prescribed in 'allowanceRecipient' above
        allowanceRecipient spender = allowanceRecipient(_spender);

        // our contract calls 'receiveApproval' function of another contract ('allowanceRecipient') to send information about
        // allowance and data sent by user
        // 'this' is this (our) contract address
        if (spender.receiveApproval(msg.sender, _value, this, _extraData)) {
            DataSentToAnotherContract(msg.sender, _spender, _extraData);
            return true;
        }
        else return false;
    }

    function approveAllAndCall(address _spender, bytes _extraData) public returns (bool success) {
        return approveAndCall(_spender, balanceOf[msg.sender], _extraData);
    }

    /* https://github.com/ethereum/EIPs/issues/677
    * transfer tokens with additional info to another smart contract, and calls its correspondent function
    * @param address _to - another smart contract address
    * @param uint256 _value - number of tokens
    * @param bytes _extraData - data to send to another contract
    * > this may be used to convert pre-ICO tokens to ICO tokens
    */
    function transferAndCall(address _to, uint256 _value, bytes _extraData) public returns (bool success){

        transferFrom(msg.sender, _to, _value);

        tokenRecipient receiver = tokenRecipient(_to);

        if (receiver.tokenFallback(msg.sender, _value, _extraData)) {
            DataSentToAnotherContract(msg.sender, _to, _extraData);
            return true;
        }
        else return false;
    }

    // for example for conveting ALL tokens of user account to another tokens
    function transferAllAndCall(address _to, bytes _extraData) public returns (bool success){
        return transferAndCall(_to, balanceOf[msg.sender], _extraData);
    }

    /* --- Administrative functions */

    function changeOwner(address _newOwner) public onlyBy(owner) returns (bool success){
        //
        require(_newOwner != address(0));

        address oldOwner = owner;
        owner = _newOwner;

        OwnerChanged(oldOwner, _newOwner);

        return true;
    }

    /* ---------- Create and sell tokens  */

    /* set time for start and time for end pre-ICO
    * time is integer representing block timestamp
    * in UNIX Time,
    * see: https://www.epochconverter.com
    * @param uint256 startTime - time to start
    * @param uint256 endTime - time to end
    * should be taken into account that
    * "block.timestamp" can be influenced by miners to a certain degree.
    * That means that a miner can "choose" the block.timestamp, to a certain degree,
    * to change the outcome of a transaction in the mined block.
    * see:
    * http://solidity.readthedocs.io/en/v0.4.15/frequently-asked-questions.html#are-timestamps-now-block-timestamp-reliable
    */

    function startSale(uint256 _startUnixTime, uint256 _endUnixTime) public onlyBy(owner) returns (bool success){

        require(balanceOf[this] > 0);
        require(salesCounter < maxSalesAllowed);

        // time for sale can be set only if:
        // this is first sale (saleStartUnixTime == 0 && saleEndUnixTime == 0) , or:
        // previous sale finished ( saleIsFinished() )
        require(
        (saleStartUnixTime == 0 && saleEndUnixTime == 0) || saleIsFinished()
        );
        // time can be set only for future
        require(_startUnixTime > now && _endUnixTime > now);
        // end time should be later than start time
        require(_endUnixTime - _startUnixTime > 0);

        saleStartUnixTime = _startUnixTime;
        saleEndUnixTime = _endUnixTime;
        salesCounter = salesCounter + 1;

        SaleStarted(_startUnixTime, _endUnixTime, salesCounter);

        return true;
    }

    function saleIsRunning() public constant returns (bool){

        if (balanceOf[this] == 0) {
            return false;
        }

        if (saleStartUnixTime == 0 && saleEndUnixTime == 0) {
            return false;
        }

        if (now > saleStartUnixTime && now < saleEndUnixTime) {
            return true;
        }

        return false;
    }

    function saleIsFinished() public constant returns (bool){

        if (balanceOf[this] == 0) {
            return true;
        }

        else if (
        (saleStartUnixTime > 0 && saleEndUnixTime > 0)
        && now > saleEndUnixTime) {

            return true;
        }

        // <<<
        return false;
    }

    function changePriceSetter(address _priceSetter) public onlyBy(owner) returns (bool success) {
        priceSetter = _priceSetter;
        return true;
    }

    function setMinMaxPriceInWei(uint256 _priceMinWei, uint256 _priceMaxWei) public onlyBy(owner) returns (bool success){
        require(_priceMinWei >= 0 && _priceMaxWei >= 0);
        priceMinWei = _priceMinWei;
        priceMaxWei = _priceMaxWei;
        return true;
    }


    function setTokenPriceInWei(uint256 _priceInWei) public onlyBy(priceSetter) returns (bool success){

        require(_priceInWei >= 0);

        // if 0 - not set
        if (priceMinWei != 0 && _priceInWei < priceMinWei) {
            tokenPriceInWei = priceMinWei;
        }
        else if (priceMaxWei != 0 && _priceInWei > priceMaxWei) {
            tokenPriceInWei = priceMaxWei;
        }
        else {
            tokenPriceInWei = _priceInWei;
        }

        PriceChanged(tokenPriceInWei);

        return true;
    }

    // allows sending ether and receiving tokens just using contract address
    // warning:
    // 'If the fallback function requires more than 2300 gas, the contract cannot receive Ether'
    // see:
    // https://ethereum.stackexchange.com/questions/21643/fallback-function-best-practices-when-registering-information
    function() public payable {
        buyTokens();
    }

    //
    function buyTokens() public payable returns (bool success){

        if (saleIsRunning() && tokenPriceInWei > 0) {

            uint256 numberOfTokens = msg.value / tokenPriceInWei;

            if (numberOfTokens <= balanceOf[this]) {

                balanceOf[msg.sender] = balanceOf[msg.sender] + numberOfTokens;
                balanceOf[this] = balanceOf[this] - numberOfTokens;

                NewTokensSold(numberOfTokens, msg.sender, tokenPriceInWei);

                return true;
            }
            else {
                // (payable)
                revert();
            }
        }
        else {
            // (payable)
            revert();
        }
    }

    /*  After sale contract owner
    *  (can be another contract or account)
    *  can withdraw all collected Ether
    */
    function withdrawAllToOwner() public onlyBy(owner) returns (bool) {

        // only after sale is finished:
        require(saleIsFinished());
        uint256 sumInWei = this.balance;

        if (
        // makes withdrawal and returns true or false
        !msg.sender.send(this.balance)
        ) {
            return false;
        }
        else {
            // event
            Withdrawal(msg.sender, sumInWei);
            return true;
        }
    }

    /* ---------- Referral System */

    // list of registered referrers
    // represented by keccak256(address) (returns bytes32)
    // ! referrers can not be removed !
    mapping (bytes32 => bool) private isReferrer;

    uint256 private referralBonus = 0;

    uint256 private referrerBonus = 0;
    // tokens owned by referrers:
    mapping (bytes32 => uint256) public referrerBalanceOf;

    mapping (bytes32 => uint) public referrerLinkedSales;

    function addReferrer(bytes32 _referrer) public onlyBy(owner) returns (bool success){
        isReferrer[_referrer] = true;
        return true;
    }

    function removeReferrer(bytes32 _referrer) public onlyBy(owner) returns (bool success){
        isReferrer[_referrer] = false;
        return true;
    }

    // bonuses are set in as integers (20%, 30%), initial 0%
    function setReferralBonuses(uint256 _referralBonus, uint256 _referrerBonus) public onlyBy(owner) returns (bool success){
        require(_referralBonus > 0 && _referrerBonus > 0);
        referralBonus = _referralBonus;
        referrerBonus = _referrerBonus;
        return true;
    }

    function buyTokensWithReferrerAddress(address _referrer) public payable returns (bool success){

        bytes32 referrer = keccak256(_referrer);

        if (saleIsRunning() && tokenPriceInWei > 0) {

            if (isReferrer[referrer]) {

                uint256 numberOfTokens = msg.value / tokenPriceInWei;

                if (numberOfTokens <= balanceOf[this]) {

                    referrerLinkedSales[referrer] = referrerLinkedSales[referrer] + numberOfTokens;

                    uint256 referralBonusTokens = (numberOfTokens * (100 + referralBonus) / 100) - numberOfTokens;
                    uint256 referrerBonusTokens = (numberOfTokens * (100 + referrerBonus) / 100) - numberOfTokens;

                    balanceOf[this] = balanceOf[this] - numberOfTokens - referralBonusTokens - referrerBonusTokens;

                    balanceOf[msg.sender] = balanceOf[msg.sender] + (numberOfTokens + referralBonusTokens);

                    referrerBalanceOf[referrer] = referrerBalanceOf[referrer] + referrerBonusTokens;

                    NewTokensSold(numberOfTokens + referralBonusTokens, msg.sender, tokenPriceInWei);

                    return true;
                }
                else {
                    // (payable)
                    revert();
                }
            }
            else {
                // (payable)
                buyTokens();
            }
        }
        else {
            // (payable)
            revert();
        }
    }

    event ReferrerBonusTokensTaken(address referrer, uint256 bonusTokensValue);

    function getReferrerBonusTokens() public returns (bool success){
        require(saleIsFinished());
        uint256 bonusTokens = referrerBalanceOf[keccak256(msg.sender)];
        balanceOf[msg.sender] = balanceOf[msg.sender] + bonusTokens;
        ReferrerBonusTokensTaken(msg.sender, bonusTokens);
        return true;
    }

}
