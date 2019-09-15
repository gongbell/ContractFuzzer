pragma solidity ^0.4.18;

/**
 * @title SafeMath for performing valid mathematics.
 */
library SafeMath {
 
  function Mul(uint a, uint b) internal pure returns (uint) {
    uint256 c = a * b;
    assert(a == 0 || c / a == b);
    return c;
  }

  function Div(uint a, uint b) internal pure returns (uint) {
    //assert(b > 0); // Solidity automatically throws when Dividing by 0
    uint256 c = a / b;
    // assert(a == b * c + a % b); // There is no case in which this doesn't hold
    return c;
  }

  function Sub(uint a, uint b) internal pure returns (uint) {
    assert(b <= a);
    return a - b;
  } 

  function Add(uint a, uint b) internal pure returns (uint) {
    uint256 c = a + b;
    assert(c >= a);
    return c;
  } 
}

/**
* @title Contract that will work with ERC223 tokens.
*/
contract ERC223ReceivingContract { 
    /**
     * @dev Standard ERC223 function that will handle incoming token transfers.
     *
     * @param _from  Token sender address.
     * @param _value Amount of tokens.
     * @param _data  Transaction metadata.
     */
    function tokenFallback(address _from, uint _value, bytes _data) public;
}

/**
 * Contract "Ownable"
 * Purpose: Defines Owner for contract and provide functionality to transfer ownership to another account
 */
contract Ownable {

  //owner variable to store contract owner account
  address public owner;
  //add another owner to transfer ownership
  address oldOwner;

  //Constructor for the contract to store owner's account on deployement
  function Ownable() public {
    owner = msg.sender;
    oldOwner = msg.sender;
  }

  //modifier to check transaction initiator is only owner
  modifier onlyOwner() {
    require (msg.sender == owner || msg.sender == oldOwner);
      _;
  }

  //ownership can be transferred to provided newOwner. Function can only be initiated by contract owner's account
  function transferOwnership(address newOwner) public onlyOwner {
    require (newOwner != address(0));
    owner = newOwner;
  }

}

/**
 * @title ERC20 interface
 */
contract ERC20 is Ownable {
    uint256 public totalSupply;
    function balanceOf(address _owner) public view returns (uint256 value);
    function transfer(address _to, uint256 _value) public returns (bool _success);
    function allowance(address owner, address spender) public view returns (uint256 _value);
    function transferFrom(address from, address to, uint256 value) public returns (bool _success);
    function approve(address spender, uint256 value) public returns (bool _success);
    event Approval(address indexed owner, address indexed spender, uint256 value);
    event Transfer(address indexed _from, address indexed _to, uint _value, bytes comment);
}

contract CTV is ERC20 {

    using SafeMath for uint256;
    //The name of the  token
    string public constant name = "Coin TV";
    //The token symbol
    string public constant symbol = "CTV";
    //To denote the locking on transfer of tokens among token holders
    bool public locked;
    //The precision used in the calculations in contract
    uint8 public constant decimals = 18;
    //maximum number of tokens
    uint256 constant MAXCAP = 29999990e18;
    // maximum number of tokens that can be supplied by referrals
    uint public constant MAX_REFERRAL_TOKENS = 2999999e18;
    //set the softcap of ether received
    uint256 constant SOFTCAP = 70 ether;
    //Refund eligible or not
    // 0: sale not started yet, refunding invalid
    // 1: refund not required
    // 2: softcap not reached, refund required
    // 3: Refund in progress
    // 4: Everyone refunded
    uint256 public refundStatus = 0;
    //the account which will receive all balance
    address ethCollector;
    //to save total number of ethers received
    uint256 totalWeiReceived;
    //count tokens earned by referrals
    uint256 public tokensSuppliedFromReferral = 0;

    //Mapping to relate owner and spender to the tokens allowed to transfer from owner
    mapping(address => mapping(address => uint256)) allowed;
    //to manage referrals
    mapping(address => address) public referredBy;
    //Mapping to relate number of  token to the account
    mapping(address => uint256) balances;

    //Structure for investors; holds received wei amount and Token sent
    struct Investor {
        //wei received during PreSale
        uint weiReceived;
        //Tokens sent during CrowdSale
        uint tokensPurchased;
        //user has been refunded or not
        bool refunded;
        //Uniquely identify an investor(used for iterating)
        uint investorID;
    }

    //time when the sale starts
    uint256 public startTime;
    //time when the presale ends
    uint256 public endTime;
    //to check the sale status
    bool public saleRunning;
    //investors indexed by their ETH address
    mapping(address => Investor) public investors;
    //investors indexed by their IDs
    mapping (uint256 => address) public investorList;
    //count number of investors
    uint256 public countTotalInvestors;
    //to keep track of how many investors have been refunded
    uint256 countInvestorsRefunded;
    //by default any new account will show false for registered mapping
    mapping(address => bool) registered;

    address[] listOfAddresses;

    //events
    event StateChanged(bool);

    function CTV() public{
        totalSupply = 0;
        startTime = 0;
        endTime = 0;
        saleRunning = false;
        locked = true;
        setEthCollector(0xAf3BBf663769De9eEb6C2b235262Cf704eD4EA4b);
    }
    //To handle ERC20 short address attack
    modifier onlyPayloadSize(uint size) {
        require(msg.data.length >= size + 4);
        _;
    }

    modifier onlyUnlocked() { 
        require (!locked); 
        _; 
    }

    modifier validTimeframe(){
        require(saleRunning && now >=startTime && now < endTime);
        _;
    }
    
    function setEthCollector(address _ethCollector) public onlyOwner{
        require(_ethCollector != address(0));
        ethCollector = _ethCollector;
    }
    
    function startSale() public onlyOwner{
        require(startTime == 0);
        startTime = now;
        endTime = startTime.Add(7 weeks);
        saleRunning = true;
    }

    //To enable transfer of tokens
    function unlockTransfer() external onlyOwner{
        locked = false;
    }

    /**
    * @dev Check if the address being passed belongs to a contract
    *
    * @param _address The address which you want to verify
    * @return A bool specifying if the address is that of contract or not
    */
    function isContract(address _address) private view returns(bool _isContract){
        assert(_address != address(0) );
        uint length;
        //inline assembly code to check the length of address
        assembly{
            length := extcodesize(_address)
        }
        if(length > 0){
            return true;
        }
        else{
            return false;
        }
    }

    /**
    * @dev Check balance of given account address
    *
    * @param _owner The address account whose balance you want to know
    * @return balance of the account
    */
    function balanceOf(address _owner) public view returns (uint256 _value){
        return balances[_owner];
    }

    /**
    * @dev Transfer sender's token to a given address
    *
    * @param _to The address which you want to transfer to
    * @param _value the amount of tokens to be transferred
    * @return A bool if the transfer was a success or not
    */
    function transfer(address _to, uint _value) onlyUnlocked onlyPayloadSize(2 * 32) public returns(bool _success) {
        require( _to != address(0) );
        bytes memory _empty;
        if((balances[msg.sender] > _value) && _value > 0 && _to != address(0)){
            balances[msg.sender] = balances[msg.sender].Sub(_value);
            balances[_to] = balances[_to].Add(_value);
            if(isContract(_to)){
                ERC223ReceivingContract receiver = ERC223ReceivingContract(_to);
                receiver.tokenFallback(msg.sender, _value, _empty);
            }
            Transfer(msg.sender, _to, _value, _empty);
            return true;
        }
        else{
            return false;
        }
    }

    /**
    * @dev Transfer tokens to an address given by sender. To make ERC223 compliant
    *
    * @param _to The address which you want to transfer to
    * @param _value the amount of tokens to be transferred
    * @param _data additional information of account from where to transfer from
    * @return A bool if the transfer was a success or not
    */
    function transfer(address _to, uint _value, bytes _data) onlyUnlocked onlyPayloadSize(3 * 32) public returns(bool _success) {
        if((balances[msg.sender] > _value) && _value > 0 && _to != address(0)){
            balances[msg.sender] = balances[msg.sender].Sub(_value);
            balances[_to] = balances[_to].Add(_value);
            if(isContract(_to)){
                ERC223ReceivingContract receiver = ERC223ReceivingContract(_to);
                receiver.tokenFallback(msg.sender, _value, _data);
            }
            Transfer(msg.sender, _to, _value, _data);
            return true;
        }
        else{
            return false;
        }
    }

    /**
    * @dev Transfer tokens from one address to another, for ERC20.
    *
    * @param _from The address which you want to send tokens from
    * @param _to The address which you want to transfer to
    * @param _value the amount of tokens to be transferred
    * @return A bool if the transfer was a success or not 
    */
    function transferFrom(address _from, address _to, uint256 _value) onlyPayloadSize(3*32) public onlyUnlocked returns (bool){
        bytes memory _empty;
        if((_value > 0)
           && (_to != address(0))
       && (_from != address(0))
       && (allowed[_from][msg.sender] > _value )){
           balances[_from] = balances[_from].Sub(_value);
           balances[_to] = balances[_to].Add(_value);
           allowed[_from][msg.sender] = allowed[_from][msg.sender].Sub(_value);
           if(isContract(_to)){
               ERC223ReceivingContract receiver = ERC223ReceivingContract(_to);
               receiver.tokenFallback(msg.sender, _value, _empty);
           }
           Transfer(_from, _to, _value, _empty);
           return true;
       }
       else{
           return false;
       }
    }

    /**
    * @dev Function to check the amount of tokens that an owner has allowed a spender to recieve from owner.
    *
    * @param _owner address The address which owns the funds.
    * @param _spender address The address which will spend the funds.
    * @return A uint256 specifying the amount of tokens still available for the spender to spend.
    */
    function allowance(address _owner, address _spender) public view returns (uint256){
        return allowed[_owner][_spender];
    }

    /**
    * @dev Approve the passed address to spend the specified amount of tokens on behalf of msg.sender.
    *
    * @param _spender The address which will spend the funds.
    * @param _value The amount of tokens to be spent.
    */
    function approve(address _spender, uint256 _value) public returns (bool){
        if( (_value > 0) && (_spender != address(0)) && (balances[msg.sender] >= _value)){
            allowed[msg.sender][_spender] = _value;
            Approval(msg.sender, _spender, _value);
            return true;
        }
        else{
            return false;
        }
    }

    /**
    * @dev Calculate number of tokens that will be received in one ether
    * 
    */
    function getPrice() public view returns(uint256) {
        uint256 price;
        if(totalSupply <= 1e6*1e18)
            price = 13330;
        else if(totalSupply <= 5e6*1e18)
            price = 12500;
        else if(totalSupply <= 9e6*1e18)
            price = 11760;
        else if(totalSupply <= 13e6*1e18)
            price = 11110;
        else if(totalSupply <= 17e6*1e18)
            price = 10520;
        else if(totalSupply <= 21e6*1e18)
            price = 10000;
        else{
            //zero indicates that no tokens will be allocated when total supply
            //of 21 million tokens is reached
            price = 0;
        }
        return price;
    }
    
    function mintAndTransfer(address beneficiary, uint256 numberOfTokensWithoutDecimal, bytes comment) public onlyOwner {
        uint256 tokensToBeTransferred = numberOfTokensWithoutDecimal*1e18;
        require(totalSupply.Add(tokensToBeTransferred) <= MAXCAP);
        totalSupply = totalSupply.Add(tokensToBeTransferred);
        balances[beneficiary] = balances[beneficiary].Add(tokensToBeTransferred);
        Transfer(owner, beneficiary ,tokensToBeTransferred, comment);
    }

    /**
    * @dev to enable pause sale for break in ICO and Pre-ICO
    *
    */
    function pauseSale() public onlyOwner{
        assert(saleRunning && startTime > 0 && now <= endTime);
        saleRunning = false;
    }

    /**
    * @dev to resume paused sale
    *
    */
    function resumeSale() public onlyOwner{
        assert(!saleRunning && startTime > 0 && now <= endTime);
        saleRunning = true;
    }

    function buyTokens(address beneficiary) internal validTimeframe {
        uint256 tokensBought = msg.value.Mul(getPrice());
        balances[beneficiary] = balances[beneficiary].Add(tokensBought);
        totalSupply = totalSupply.Add(tokensBought);

        //Make entry in Investor indexed with address
        Investor storage investorStruct = investors[beneficiary];
        //If it is a new investor, then create a new id
        if(investorStruct.investorID == 0){
            countTotalInvestors++;
            investorStruct.investorID = countTotalInvestors;
            investorList[countTotalInvestors] = beneficiary;
        }
        else{
            investorStruct.weiReceived = investorStruct.weiReceived.Add(msg.value);
            investorStruct.tokensPurchased = investorStruct.tokensPurchased.Add(tokensBought);
        }
        
        //Award referral tokens
        if(referredBy[msg.sender] != address(0)){
            //give some referral tokens
            balances[referredBy[msg.sender]] = balances[referredBy[msg.sender]].Add(tokensBought/10);
            tokensSuppliedFromReferral = tokensSuppliedFromReferral.Add(tokensBought/10);
            totalSupply = totalSupply.Add(tokensBought/10);
        }
        //if referrer was also referred by someone
        if(referredBy[referredBy[msg.sender]] != address(0)){
            //give 1% tokens to 2nd generation referrer
            balances[referredBy[referredBy[msg.sender]]] = balances[referredBy[referredBy[msg.sender]]].Add(tokensBought/100);
            if(tokensSuppliedFromReferral.Add(tokensBought/100) < MAX_REFERRAL_TOKENS)
                tokensSuppliedFromReferral = tokensSuppliedFromReferral.Add(tokensBought/100);
            totalSupply = totalSupply.Add(tokensBought/100);
        }
        
        assert(totalSupply <= MAXCAP);
        totalWeiReceived = totalWeiReceived.Add(msg.value);
        ethCollector.transfer(msg.value);
    }

    /**
     * @dev This function is used to register a referral.
     * Whoever calls this function, is telling contract,
     * that "I was referred by referredByAddress"
     * Whenever I am going to buy tokens, 10% will be awarded to referredByAddress
     * 
     * @param referredByAddress The address of person who referred the person calling this function
     */
    function registerReferral (address referredByAddress) public {
        require(msg.sender != referredByAddress && referredByAddress != address(0));
        referredBy[msg.sender] = referredByAddress;
    }
    
    /**
     * @dev Owner is allowed to manually register who was referred by whom
     * @param heWasReferred The address of person who was referred
     * @param I_referred_this_person The person who referred the above address
     */
    function referralRegistration(address heWasReferred, address I_referred_this_person) public onlyOwner {
        require(heWasReferred != address(0) && I_referred_this_person != address(0));
        referredBy[heWasReferred] = I_referred_this_person;
    }

    /**
    * Finalize the crowdsale
    */
    function finalize() public onlyOwner {
        //Make sure Sale is running
        assert(saleRunning);
        if(MAXCAP.Sub(totalSupply) <= 1 ether || now > endTime){
            //now sale can be finished
            saleRunning = false;
        }

        //Refund eligible or not
        // 0: sale not started yet, refunding invalid
        // 1: refund not required
        // 2: softcap not reached, refund required
        // 3: Refund in progress
        // 4: Everyone refunded

        //Checks if the fundraising goal is reached in crowdsale or not
        if (totalWeiReceived < SOFTCAP)
            refundStatus = 2;
        else
            refundStatus = 1;

        //crowdsale is ended
        saleRunning = false;
        //enable transferring of tokens among token holders
        locked = false;
        //Emit event when crowdsale state changes
        StateChanged(true);
    }

    /**
    * Refund the investors in case target of crowdsale not achieved
    */
    function refund() public onlyOwner {
        assert(refundStatus == 2 || refundStatus == 3);
        uint batchSize = countInvestorsRefunded.Add(30) < countTotalInvestors ? countInvestorsRefunded.Add(30): countTotalInvestors;
        for(uint i=countInvestorsRefunded.Add(1); i <= batchSize; i++){
            address investorAddress = investorList[i];
            Investor storage investorStruct = investors[investorAddress];
            //If purchase has been made during CrowdSale
            if(investorStruct.tokensPurchased > 0 && investorStruct.tokensPurchased <= balances[investorAddress]){
                //return everything
                investorAddress.transfer(investorStruct.weiReceived);
                //Reduce totalWeiReceived
                totalWeiReceived = totalWeiReceived.Sub(investorStruct.weiReceived);
                //Update totalSupply
                totalSupply = totalSupply.Sub(investorStruct.tokensPurchased);
                // reduce balances
                balances[investorAddress] = balances[investorAddress].Sub(investorStruct.tokensPurchased);
                //set everything to zero after transfer successful
                investorStruct.weiReceived = 0;
                investorStruct.tokensPurchased = 0;
                investorStruct.refunded = true;
            }
        }
        //Update the number of investors that have recieved refund
        countInvestorsRefunded = batchSize;
        if(countInvestorsRefunded == countTotalInvestors){
            refundStatus = 4;
        }
        StateChanged(true);
    }
    
    function extendSale(uint56 numberOfDays) public onlyOwner{
        saleRunning = true;
        endTime = now.Add(numberOfDays*86400);
        StateChanged(true);
    }

    /**
    * @dev This will receive ether from owner so that the contract has balance while refunding
    *
    */
    function prepareForRefund() public payable {}

    function () public payable {
        buyTokens(msg.sender);
    }

    /**
    * Failsafe drain
    */
    function drain() public onlyOwner {
        owner.transfer(this.balance);
    }
}