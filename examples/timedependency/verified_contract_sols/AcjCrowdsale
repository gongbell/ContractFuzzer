/**
 *Submitted for verification at Etherscan.io on 2018-01-16
*/

pragma solidity ^0.4.18;


/**
 * Math operations with safety checks
 */

library SafeMath {


    function mul(uint a, uint b) internal pure returns (uint) {
        uint c = a * b;
        assert(a == 0 || c / a == b);
        return c;
    }

    function sub(uint a, uint b) internal pure returns (uint) {
        assert(b <= a);
        return a - b;
    }

    function add(uint a, uint b) internal pure returns (uint) {
        uint c = a + b;
        assert(c>=a && c>=b);
        return c;
    }
}


/* 
 * Token related contracts 
 */


/*
 * ERC20Basic
 * Simpler version of ERC20 interface
 * see https://github.com/ethereum/EIPs/issues/20
 */

contract ERC20Basic {
    uint public totalSupply;
    function balanceOf(address who) public view returns (uint);
    function transfer(address to, uint256 value) public returns (bool);
    event Transfer(address indexed from, address indexed to, uint value);
}


/**
 * @title ERC20 interface
 * @dev see https://github.com/ethereum/EIPs/issues/20
 */

contract ERC20 is ERC20Basic {
    function allowance(address owner, address spender) public view returns (uint256);
    function transferFrom(address from, address to, uint256 value) public returns (bool);
    function approve(address spender, uint256 value) public returns (bool);
    event Approval(address indexed owner, address indexed spender, uint256 value);
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
     * @return An uint representing the amount owned by the passed address.
     */
    function balanceOf(address _owner) public view returns (uint) {
        return balances[_owner];
    }

}


/**
 * @title Standard ERC20 token
 *
 * @dev Implementation of the basic standard token.
 * @dev https://github.com/ethereum/EIPs/issues/20
 * @dev Based on code by FirstBlood: https://github.com/Firstbloodio/token/blob/master/smart_contract/FirstBloodToken.sol
 */

contract StandardToken is ERC20, BasicToken {

    mapping (address => mapping (address => uint256)) internal allowed;


    /**
     * @dev Transfer tokens from one address to another
     * @param _from address The address which you want to send tokens from
     * @param _to address The address which you want to transfer to
     * @param _value uint256 the amount of tokens to be transferred
     */
    function transferFrom(address _from, address _to, uint256 _value) public returns (bool) {
        require(_to != address(0));
        require(_value <= balances[_from]);
        require(_value <= allowed[_from][msg.sender]);

        balances[_from] = balances[_from].sub(_value);
        balances[_to] = balances[_to].add(_value);
        allowed[_from][msg.sender] = allowed[_from][msg.sender].sub(_value);
        Transfer(_from, _to, _value);
        return true;
    }

    /**
     * @dev Approve the passed address to spend the specified amount of tokens on behalf of msg.sender.
     *
     * Beware that changing an allowance with this method brings the risk that someone may use both the old
     * and the new allowance by unfortunate transaction ordering. One possible solution to mitigate this
     * race condition is to first reduce the spender's allowance to 0 and set the desired value afterwards:
     * https://github.com/ethereum/EIPs/issues/20#issuecomment-263524729
     * @param _spender The address which will spend the funds.
     * @param _value The amount of tokens to be spent.
     */
    function approve(address _spender, uint256 _value) public returns (bool) {
        allowed[msg.sender][_spender] = _value;
        Approval(msg.sender, _spender, _value);
        return true;
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

}


/*
 * Ownable
 *
 * Base contract with an owner.
 * Provides onlyOwner modifier, which prevents function from running if it is called by anyone other than the owner.
 */

contract Ownable {
    address public owner;
    address public newOwner;

    function Ownable() public {
        owner = msg.sender;
    }

    modifier onlyOwner() { 
        require(msg.sender == owner);
        _;
    }

    modifier onlyNewOwner() {
        require(msg.sender == newOwner);
        _;
    }
    /*
    // This code is dangerous because an error in the newOwner 
    // means that this contract will be ownerless 
    function transfer(address newOwner) public onlyOwner {
        require(newOwner != address(0)); 
        owner = newOwner;
    }
   */

    function proposeNewOwner(address _newOwner) external onlyOwner {
        require(_newOwner != address(0));
        newOwner = _newOwner;
    }

    function acceptOwnership() external onlyNewOwner {
        require(newOwner != owner);
        owner = newOwner;
    }
}


/**
 * @title Mintable token
 * @dev Simple ERC20 Token example, with mintable token creation
 * @dev Issue: * https://github.com/OpenZeppelin/zeppelin-solidity/issues/120
 * Based on code by TokenMarketNet: https://github.com/TokenMarketNet/ico/blob/master/contracts/MintableToken.sol
 */

contract MintableToken is StandardToken, Ownable {
    event Mint(address indexed to, uint256 amount);
    event MintFinished();

    bool public mintingFinished = false;


    modifier canMint() {
        require(!mintingFinished);
        _;
    }

    /**
     * @dev Function to mint tokens
     * @param _to The address that will receive the minted tokens.
     * @param _amount The amount of tokens to mint.
     * @return A boolean that indicates if the operation was successful.
     */
    function mint(address _to, uint256 _amount) public onlyOwner canMint returns (bool) {
        totalSupply = totalSupply.add(_amount);
        balances[_to] = balances[_to].add(_amount);
        Mint(_to, _amount);
        Transfer(address(0), _to, _amount);
        return true;
    }

    /**
     * @dev Function to stop minting new tokens.
     * @return True if the operation was successful.
     */
    function finishMinting() public onlyOwner canMint returns (bool) {
        mintingFinished = true;
        MintFinished();
        return true;
    }
}


/**
 * @title Burnable Token
 * @dev Token that can be irreversibly burned (destroyed).
 */
contract BurnableToken is BasicToken {

    event Burn(address indexed burner, uint256 value);

    /**
     * @dev Burns a specific amount of tokens.
     * @param _value The amount of token to be burned.
     */
    function burn(uint256 _value) public  {
        require(_value <= balances[msg.sender]);
        // no need to require value <= totalSupply, since that would imply the
        // sender's balance is greater than the totalSupply, which *should* be an assertion failure

        address burner = msg.sender;
        balances[burner] = balances[burner].sub(_value);
        totalSupply = totalSupply.sub(_value);
        Burn(burner, _value);
    }
}


/**
 * @title Pausable
 * @dev Base contract which allows children to implement an emergency stop mechanism.
 */

contract Pausable is Ownable {


    event Pause();
    event Unpause();

    bool public paused = true;


    /**
     * @dev Modifier to make a function callable only when the contract is not paused.
     */
    modifier whenNotPaused() {
        require(!paused);
        _;
    }

    /**
     * @dev Modifier to make a function callable only when the contract is paused.
     */
    modifier whenPaused() {
        require(paused);
        _;
    }

    /**
     * @dev called by the owner to pause, triggers stopped state
     */
    function pause() onlyOwner whenNotPaused public {
        paused = true;
        Pause();
    }

    /**
     * @dev called by the owner to unpause, returns to normal state
     */
    function unpause() onlyOwner whenPaused public {
        paused = false;
        Unpause();
    }
}



/* @title Pausable token
 *
 * @dev StandardToken modified with pausable transfers.
 **/

contract PausableToken is StandardToken, Pausable {

    function transfer(address _to, uint256 _value) public whenNotPaused returns (bool) {
        return super.transfer(_to, _value);
    }

    function transferFrom(address _from, address _to, uint256 _value) public whenNotPaused returns (bool) {
        return super.transferFrom(_from, _to, _value);
    }

    function approve(address _spender, uint256 _value) public whenNotPaused returns (bool) {
        return super.approve(_spender, _value);
    }

}


/*
 * Actual token contract
 */

contract AcjToken is BurnableToken, MintableToken, PausableToken {
    using SafeMath for uint256;

    string public constant name = "Artist Connect Coin";
    string public constant symbol = "ACJ";
    uint public constant decimals = 18;
    
    function AcjToken() public {
        totalSupply = 150000000 ether; 
        balances[msg.sender] = totalSupply;
        paused = true;
    }

    function activate() external onlyOwner {
        unpause();
        finishMinting();
    }

    // This method will be used by the crowdsale smart contract 
    // that owns the AcjToken and will distribute 
    // the tokens to the contributors
    function initialTransfer(address _to, uint _value) external onlyOwner returns (bool) {
        require(_to != address(0));
        require(_value <= balances[msg.sender]);

        balances[msg.sender] = balances[msg.sender].sub(_value);
        balances[_to] = balances[_to].add(_value);
        Transfer(msg.sender, _to, _value);
        return true;
    }

    function burn(uint256 _amount) public onlyOwner {
        super.burn(_amount);
    }

}

 


contract AcjCrowdsale is Ownable {

    using SafeMath for uint256;
    
    // Presale bonus percentage
    uint public constant BONUS_PRESALE = 10;            
    // Medium bonus percentage
    uint public constant BONUS_MID = 10;                
    // High bonus percentage
    uint public constant BONUS_HI = 20;                 
    // Medium bonus threshold
    uint public constant BONUS_MID_QTY = 150 ether;     
    // High bonus threshold
    uint public constant BONUS_HI_QTY = 335 ether;      
    // Absolute dates as timestamps
    uint public startPresale;            
    uint public endPresale;             
    uint public startIco;              
    uint public endIco;               
    // 30 days refund period on fail
    uint public constant REFUND_PERIOD = 30 days;
    // Indicative token balances during the crowdsale 
    mapping(address => uint256) public tokenBalances;    
    // Token smart contract address
    address public token;
    // Total tokens created
    uint256 public constant TOKENS_TOTAL_SUPPLY = 150000000 ether; 
    // Tokens available for sale
    uint256 public constant TOKENS_FOR_SALE = 75000000 ether;    
    // soft cap in Tokens
    uint256 public constant TOKENS_SOFT_CAP = 500000 ether;       
    // Tokens sold via buyTokens
    uint256 public tokensSold;                             
    // Tokens created during the sale
    uint256 public tokensDistributed;                                         
    // ICO flat rate subject to bonuses
    uint256 public ethTokenRate;                                 
    // Allow multiple administrators
    mapping(address => bool) public admins;                    
    // Total wei received 
    uint256 public weiReceived;                            
    // Minimum contribution in ETH
    uint256 public constant MIN_CONTRIBUTION = 100 finney;           
    // Contributions in wei for each address
    mapping(address => uint256) public contributions;
    // Refund state for each address
    mapping(address => bool) public refunds;
    // Company wallet that will receive the ETH
    address public companyWallet;     

    // Yoohoo someone contributed !
    event Contribute(address indexed _from, uint _amount); 
    // Token <> ETH rate updated
    event TokenRateUpdated(uint _newRate);                  
    // ETH Refund 
    event Refunded(address indexed _from, uint _amount);    
    
    modifier belowTotalSupply {
        require(tokensDistributed < TOKENS_TOTAL_SUPPLY);
        _;
    }

    modifier belowHardCap {
        require(tokensDistributed < TOKENS_FOR_SALE);
        _;
    }

    modifier adminOnly {
        require(msg.sender == owner || admins[msg.sender] == true);
        _;
    }

    modifier crowdsaleFailed {
        require(isFailed());
        _;
    }

    modifier crowdsaleSuccess {
        require(isSuccess());
        _;
    }

    modifier duringSale {
        require(now < endIco);
        require((now > startPresale && now < endPresale) || now > startIco);
        _;
    }

    modifier afterSale {
        require(now > endIco);
        _;
    }

    modifier aboveMinimum {
        require(msg.value >= MIN_CONTRIBUTION);
        _;
    }

    /* 
     * Constructor
     * Creating the new Token smart contract
     * and setting its owner to the current sender
     * 
     */
    function AcjCrowdsale(
        uint _presaleStart,
        uint _presaleEnd,
        uint _icoStart,
        uint _icoEnd,
        uint256 _rate,
        address _token
    ) public 
    {
        require(_presaleEnd > _presaleStart);
        require(_icoStart > _presaleEnd);
        require(_icoEnd > _icoStart);
        require(_rate > 0); 

        startPresale = _presaleStart;
        endPresale = _presaleEnd;
        startIco = _icoStart;
        endIco = _icoEnd;
        ethTokenRate = _rate;
        
        admins[msg.sender] = true;
        companyWallet = msg.sender;

        token = _token;
    }

    /*
     * Fallback payable
     */
    function () external payable {
        buyTokens(msg.sender);
    }

    /* Crowdsale staff only */
    /*
     * Admin management
     */
    function addAdmin(address _adr) external onlyOwner {
        require(_adr != address(0));
        admins[_adr] = true;
    }

    function removeAdmin(address _adr) external onlyOwner {
        require(_adr != address(0));
        admins[_adr] = false;
    }

    /*
     * Change the company wallet
     */
    function updateCompanyWallet(address _wallet) external adminOnly {
        companyWallet = _wallet;
    }

    /*
     *  Change the owner of the token
     */
    function proposeTokenOwner(address _newOwner) external adminOnly {
        AcjToken _token = AcjToken(token);
        _token.proposeNewOwner(_newOwner);
    }

    function acceptTokenOwnership() external onlyOwner {    
        AcjToken _token = AcjToken(token);
        _token.acceptOwnership();
    }

    /*
     * Activate the token
     */
    function activateToken() external adminOnly crowdsaleSuccess afterSale {
        AcjToken _token = AcjToken(token);
        _token.activate();
    }

    /* 
     * Adjust the token value before the ICO
     */
    function adjustTokenExchangeRate(uint _rate) external adminOnly {
        require(now > endPresale && now < startIco);
        ethTokenRate = _rate;
        TokenRateUpdated(_rate);
    }

    /* 
     * Start therefund period
     * Each contributor has to claim own  ETH 
     */     
    function refundContribution() external crowdsaleFailed afterSale {
        require(!refunds[msg.sender]);
        require(contributions[msg.sender] > 0);

        uint256 _amount = contributions[msg.sender];
        tokenBalances[msg.sender] = 0;
        refunds[msg.sender] = true;
        Refunded(msg.sender, contributions[msg.sender]);
        msg.sender.transfer(_amount);
    }

    /*
     * After the refund period, remaining tokens
     * are transfered to the company wallet
     * Allow withdrawal at any time if the ICO is a success.
     */     
    function withdrawUnclaimed() external adminOnly {
        require(now > endIco + REFUND_PERIOD || isSuccess());
        companyWallet.transfer(this.balance);
    }

    /*
     * Pre-ICO and offline Investors, collaborators and team tokens
     */
    function reserveTokens(address _beneficiary, uint256 _tokensQty) external adminOnly belowTotalSupply {
        require(_beneficiary != address(0));
        uint _distributed = tokensDistributed.add(_tokensQty);

        require(_distributed <= TOKENS_TOTAL_SUPPLY);

        tokenBalances[_beneficiary] = _tokensQty.add(tokenBalances[_beneficiary]);
        tokensDistributed = _distributed;

        AcjToken _token = AcjToken(token);
        _token.initialTransfer(_beneficiary, _tokensQty);
    }

    /*
     * Actually buy the tokens
     * requires an active sale time
     * and amount above the minimum contribution
     * and sold tokens inferior to tokens for sale
     */     
    function buyTokens(address _beneficiary) public payable duringSale aboveMinimum belowHardCap {
        require(_beneficiary != address(0));
        uint256 _weiAmount = msg.value;        
        uint256 _tokensQty = msg.value.mul(getBonus(_weiAmount));
        uint256 _distributed = _tokensQty.add(tokensDistributed);
        uint256 _sold = _tokensQty.add(tokensSold);

        require(_distributed <= TOKENS_TOTAL_SUPPLY);
        require(_sold <= TOKENS_FOR_SALE);

        contributions[_beneficiary] = _weiAmount.add(contributions[_beneficiary]);
        tokenBalances[_beneficiary] = _tokensQty.add(tokenBalances[_beneficiary]);
        weiReceived = weiReceived.add(_weiAmount);
        tokensDistributed = _distributed;
        tokensSold = _sold;

        Contribute(_beneficiary, msg.value);

        AcjToken _token = AcjToken(token);
        _token.initialTransfer(_beneficiary, _tokensQty);
    }

    /*
     * Crowdsale Helpers 
     */
    function hasEnded() public view returns(bool) {
        return now > endIco;
    }

    /*
     * Checks if the crowdsale is a success
     */
    function isSuccess() public view returns(bool) {
        if (tokensDistributed >= TOKENS_SOFT_CAP) {
            return true;
        }
        return false;
    }

    /*
     * Checks if the crowdsale failed
     */
    function isFailed() public view returns(bool) {
        if (tokensDistributed < TOKENS_SOFT_CAP && now > endIco) {
            return true;
        }
        return false;
    }

    /* 
     * Bonus calculations
     * Either time or ETH quantity based 
     */
    function getBonus(uint256 _wei) internal constant returns(uint256 ethToAcj) {
        uint256 _bonus = 0;

        // Time based bonus
        if (endPresale > now) {
            _bonus = _bonus.add(BONUS_PRESALE); 
        }

        // ETH Quantity based bonus
        if (_wei >= BONUS_HI_QTY) { 
            _bonus = _bonus.add(BONUS_HI);
        } else if (_wei >= BONUS_MID_QTY) {
            _bonus = _bonus.add(BONUS_MID);
        }

        return ethTokenRate.mul(100 + _bonus) / 100;
    }

}
