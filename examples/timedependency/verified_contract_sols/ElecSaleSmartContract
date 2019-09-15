pragma solidity ^0.4.11;


/**
 * @title ERC20Basic
 * @dev Simpler version of ERC20 interface
 * @dev see https://github.com/ethereum/EIPs/issues/179
 */
contract ERC20Basic {
  uint256 public totalSupply;
  function balanceOf(address who) public constant returns (uint256);
  function transfer(address to, uint256 value) public returns (bool);
  event Transfer(address indexed _from, address indexed _to, uint _value);
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
    if (newOwner != address(0)) {
      owner = newOwner;
    }
  }

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
  function balanceOf(address _owner) public constant returns (uint256 balance) {
    return balances[_owner];
  }

}







/**
 * @title ERC20 interface
 * @dev see https://github.com/ethereum/EIPs/issues/20
 */
contract ERC20 is ERC20Basic {
  function allowance(address owner, address spender) public constant returns (uint256);
  function transferFrom(address from, address to, uint256 value) public returns (bool);
  function approve(address spender, uint256 value) public returns (bool);
  event Approval(address indexed _owner, address indexed _spender, uint _value);
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
  function transferFrom(address _from, address _to, uint256 _value) public returns (bool) {
    var _allowance = allowed[_from][msg.sender];

    // Check is not needed because sub(_allowance, _value) will already throw if this condition is not met
    // require (_value <= _allowance);

    balances[_from] = balances[_from].sub(_value);
    balances[_to] = balances[_to].add(_value);
    allowed[_from][msg.sender] = _allowance.sub(_value);
    Transfer(_from, _to, _value);
    return true;
  }

  /**
   * @dev Aprove the passed address to spend the specified amount of tokens on behalf of msg.sender.
   * @param _spender The address which will spend the funds.
   * @param _value The amount of tokens to be spent.
   */
  function approve(address _spender, uint256 _value) public returns (bool) {

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
  function allowance(address _owner, address _spender) public constant returns (uint256 remaining) {
    return allowed[_owner][_spender];
  }

}



contract ElecTokenSmartContract is StandardToken, Ownable {
    string  public  constant name = "ELECTRIFY TOKEN";
    string  public  constant symbol = "ELEC";
    uint8    public  constant decimals = 18;

    uint    public  saleStartTime;
    uint    public  saleEndTime;
    uint    public lockedDays = 0;

    address public  tokenSaleContract;
    address public adminAddress;

    modifier onlyWhenTransferEnabled() {
        if( now <= (saleEndTime + lockedDays * 1 days) && now >= saleStartTime ) {
            require( msg.sender == tokenSaleContract || msg.sender == adminAddress );
        }
        _;
    }

    modifier validDestination( address to ) {
        require(to != address(0x0));
        require(to != address(this) );
        _;
    }

    function ElecTokenSmartContract( uint tokenTotalAmount, uint startTime, uint endTime, uint lockedTime, address admin ) public {
        // Mint all tokens. Then disable minting forever.
        balances[msg.sender] = tokenTotalAmount;
        totalSupply = tokenTotalAmount;
        Transfer(address(0x0), msg.sender, tokenTotalAmount);

        saleStartTime = startTime;
        saleEndTime = endTime;
        lockedDays = lockedTime;

        tokenSaleContract = msg.sender;
        adminAddress = admin;
        transferOwnership(admin); // admin could drain tokens that were sent here by mistake
    }

    function transfer(address _to, uint _value)
    public
    onlyWhenTransferEnabled
    validDestination(_to)
    returns (bool) {
        return super.transfer(_to, _value);
    }

    function transferFrom(address _from, address _to, uint _value)
    public
    onlyWhenTransferEnabled
    validDestination(_to)
    returns (bool) {
        return super.transferFrom(_from, _to, _value);
    }

    event Burn(address indexed _burner, uint _value);

    function burn(uint _value) public onlyWhenTransferEnabled
    returns (bool){
        balances[msg.sender] = balances[msg.sender].sub(_value);
        totalSupply = totalSupply.sub(_value);
        Burn(msg.sender, _value);
        Transfer(msg.sender, address(0x0), _value);
        return true;
    }


    function emergencyERC20Drain( ERC20 token, uint amount ) public onlyOwner {
        token.transfer( owner, amount );
    }
}








/**
 * @title SafeMath
 * @dev Math operations with safety checks that throw on error
 */
library SafeMath {
  function mul(uint256 a, uint256 b) internal pure returns (uint256) {
    uint256 c = a * b;
    assert(a == 0 || c / a == b);
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


contract ElecApprover {
    ElecWhitelist public list;
    mapping(address=>uint)    public participated;

    uint                      public saleStartTime;
    uint                      public firstRoundTime;
    uint                      public saleEndTime;
    uint                      public xtime = 5;/// multiply the cap

    using SafeMath for uint;


    function ElecApprover( ElecWhitelist _whitelistContract,
    uint                      _saleStartTime,
    uint                      _firstRoundTime,
    uint                      _saleEndTime ) public {
        list = _whitelistContract;
        saleStartTime = _saleStartTime;
        firstRoundTime = _firstRoundTime;
        saleEndTime = _saleEndTime;

        require( list != ElecWhitelist(0x0) );
        require( saleStartTime < firstRoundTime );
        require(  firstRoundTime < saleEndTime );
    }

    // this is a seperate function so user could query it before crowdsale starts
    function contributorCap( address contributor ) public constant returns(uint) {
        uint  cap= list.getCap( contributor );
        uint higherCap = cap;

        if ( now > firstRoundTime ) {
            higherCap = cap.mul(xtime);
        }
        return higherCap;
    }


    function eligible( address contributor, uint amountInWei ) public constant returns(uint) {
        if( now < saleStartTime ) return 0;
        if( now >= saleEndTime ) return 0;

        uint cap = list.getCap( contributor );

        if( cap == 0 ) return 0;

        uint higherCap = cap;
        if ( now > firstRoundTime ) {
            higherCap = cap.mul(xtime);
        }

        uint remainedCap = higherCap.sub(participated[ contributor ]);
        if( remainedCap > amountInWei ) return amountInWei;
              else return remainedCap;

    }

    function eligibleTestAndIncrement( address contributor, uint amountInWei ) internal returns(uint) {
        uint result = eligible( contributor, amountInWei );
        if ( result > 0) {
            participated[contributor] = participated[contributor].add( result );
        }
        return result;
    }


    function contributedCap(address _contributor) public constant returns(uint) {
        if (participated[_contributor] == 0 ) return 0;

        return participated[_contributor];
    }

     function contributedInternalCap(address _contributor) view internal returns(uint) {
         if (participated[_contributor] == 0 ) return 0;

        return participated[_contributor];
    }


    function saleEnded() public constant returns(bool) {
        return now > saleEndTime;
    }

    function saleStarted() public constant returns(bool) {
        return now >= saleStartTime;
    }
}





contract ElecWhitelist is Ownable {
    // cap is in wei. The value of 1 is just a stub.
    // after kyc registration ends, we change it to the actual value with setUsersCap
    /// Currenty we set the cap to 1 ETH and the owner is able to change it in the future by call function: setUsersCap
    uint public communityusersCap = (10**18);
    mapping(address=>uint) public addressCap;

    function ElecWhitelist() public {}

    event ListAddress( address _user, uint _cap, uint _time );

    // Owner can remove by setting cap = 0.
    // Onwer can also change it at any time
    function listAddress( address _user, uint _cap ) public onlyOwner {
        addressCap[_user] = _cap;
        ListAddress( _user, _cap, now );
    }

    // an optimization in case of network congestion
    function listAddresses( address[] _users, uint[] _cap ) public onlyOwner {
        require(_users.length == _cap.length );
        for( uint i = 0 ; i < _users.length ; i++ ) {
            listAddress( _users[i], _cap[i] );
        }
    }

    function setUsersCap( uint _cap ) public  onlyOwner {
        communityusersCap = _cap;
    }

    function getCap( address _user ) public constant returns(uint) {
        uint cap = addressCap[_user];
        if( cap == 1 ) return communityusersCap;
        else return cap;
    }

    function destroy() public onlyOwner {
        selfdestruct(owner);
    }
}


contract ElecSaleSmartContract is ElecApprover{
    address             public admin;
    address             public multiSigWallet; // can be a single wallet
    ElecTokenSmartContract public token;
    uint                public raisedWei;
    bool                public haltSale;
    uint                constant toWei = (10**18);
    uint                public minCap = toWei.div(2);

    mapping(bytes32=>uint) public proxyPurchases;

    function ElecSaleSmartContract( address _admin,
    address _multiSigWallet,
    ElecWhitelist _whiteListContract,
    uint _totalTokenSupply,
    uint _companyTokenSupply,
    uint _saleStartTime,
    uint _firstRoundTime,
    uint _saleEndTime,
    uint _lockedDays)

    public

    ElecApprover( _whiteListContract,
    _saleStartTime,
    _firstRoundTime,
    _saleEndTime )
    {
        admin = _admin;
        multiSigWallet = _multiSigWallet;

        token = new ElecTokenSmartContract( _totalTokenSupply,
        _saleStartTime,
        _saleEndTime,
        _lockedDays, ///change depending on each project
        _admin );

        // transfer preminted tokens to company wallet
        token.transfer( multiSigWallet, _companyTokenSupply );
    }

    function setHaltSale( bool halt ) public {
        require( msg.sender == admin );
        haltSale = halt;
    }

    function() public payable {
        buy( msg.sender );
    }

    event ProxyBuy( bytes32 indexed _proxy, address _recipient, uint _amountInWei );
    function proxyBuy( bytes32 proxy, address recipient ) public payable returns(uint){
        uint amount = buy( recipient );
        proxyPurchases[proxy] = proxyPurchases[proxy].add(amount);
        ProxyBuy( proxy, recipient, amount );


        return amount;
    }

    event Buy( address _buyer, uint _tokens, uint _payedWei );
    function buy( address recipient ) public payable returns(uint){
        require( tx.gasprice <= 50000000000 wei );

        require( ! haltSale );
        require( saleStarted() );
        require( ! saleEnded() );

        // check min buy at least 0.5 ETH;
        uint weiContributedCap = contributedInternalCap(recipient);

        if (weiContributedCap == 0 ) require( msg.value >= minCap);



        uint weiPayment = eligibleTestAndIncrement( recipient, msg.value );

        require( weiPayment > 0 );


        // send to msg.sender, not to recipient
        if( msg.value > weiPayment ) {
            msg.sender.transfer( msg.value.sub( weiPayment ) );
        }

        // send payment to wallet
        sendETHToMultiSig( weiPayment );
        raisedWei = raisedWei.add( weiPayment );
        uint recievedTokens = weiPayment.mul( 5000 );

        assert( token.transfer( recipient, recievedTokens ) );


        Buy( recipient, recievedTokens, weiPayment );

        return weiPayment;
    }

    function sendETHToMultiSig( uint value ) internal {
        multiSigWallet.transfer( value );
    }

    event FinalizeSale();
    // function is callable by everyone
    function finalizeSale() public {
        require( saleEnded() );
        require( msg.sender == admin );

        // burn remaining tokens
        token.burn(token.balanceOf(this));

        FinalizeSale();
    }

    // ETH balance is always expected to be 0.
    // but in case something went wrong, we use this function to extract the eth.
    function emergencyDrain(ERC20 anyToken) public returns(bool){
        require( msg.sender == admin );
        require( saleEnded() );

        if( this.balance > 0 ) {
            sendETHToMultiSig( this.balance );
        }

        if( anyToken != address(0x0) ) {
            assert( anyToken.transfer(multiSigWallet, anyToken.balanceOf(this)) );
        }

        return true;
    }

    // just to check that funds goes to the right place
    // tokens are not given in return
    /*function debugBuy() public payable {
        require( msg.value == 123 );
        sendETHToMultiSig( msg.value );
    }*/
}