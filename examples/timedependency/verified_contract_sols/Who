pragma solidity 0.4.20;

/**
 * @title ERC20Basic
 * @dev Simpler version of ERC20 interface
 * @dev see https://github.com/ethereum/EIPs/issues/179
 */
contract ERC20Interface {
    function circulatingSupply() public view returns (uint);
    function balanceOf(address who) public view returns (uint);
    function transfer(address to, uint value) public returns (bool);
    event TransferEvent(address indexed from, address indexed to, uint value);
}



/**
 * @title SafeMath
 * @dev Math operations with safety checks that throw on error
 */
library SafeMath {

  /**
  * @dev Multiplies two numbers, throws on overflow.
  */
    function mul(uint256 a, uint256 b) internal pure returns (uint256) {
        if (a == 0) {
            return 0;
        }
        uint256 c = a * b;
        assert(c / a == b);
        return c;
    }

    /**
    * @dev Integer division of two numbers, truncating the quotient.
    */
    function div(uint256 a, uint256 b) internal pure returns (uint256) {
      // assert(b > 0); // Solidity automatically throws when dividing by 0
        uint256 c = a / b;
      // assert(a == b * c + a % b); // There is no case in which this doesn't hold
        return c;
    }

    /**
    * @dev Substracts two numbers, throws on overflow (i.e. if subtrahend is greater than minuend).
    */
    function sub(uint256 a, uint256 b) internal pure returns (uint256) {
        assert(b <= a);
        return a - b;
    }

    /**
    * @dev Adds two numbers, throws on overflow.
    */
    function add(uint256 a, uint256 b) internal pure returns (uint256) {
        uint256 c = a + b;
        assert(c >= a);
        return c;
    }
}


contract WhoVote {

    mapping (address => bytes32) public voteHash;
    address public parentContract;
    uint public deadline;

    modifier isActive {
        require(now < deadline);
        _;
    }

    modifier isParent {
        require(msg.sender == parentContract);
        _;
    }

    function WhoVote(address _parentContract, uint timespan) public {
        parentContract = _parentContract;
        deadline = now + timespan;
    }

    /**
    * @dev Recieve Vote from Who-Token-Contract
    * @param _sender Contest-participant
    * @param _hash Hash of the JSON-Parameter
    */
    function recieveVote(address _sender, bytes32 _hash) public isActive isParent returns (bool) {
        require(voteHash[_sender] == 0);
        voteHash[_sender] = _hash;
        return true;
    }


}



/**
 * @title Basic token
 * @dev Basic version of StandardToken, with no allowances.
 */
contract StandardToken is ERC20Interface {
    using SafeMath for uint;

    uint public maxSupply_;
    uint public circulatingSupply_;
    uint public timestampMint;
    uint public timestampRelease;
    uint8 public decimals;

    string public symbol;
    string public  name;

    address public owner;

    mapping(address => uint) public balances;
    mapping (address => uint) public permissonedAccounts;

    /**
    * @dev Checks if last mint is 3 weeks in past
    */
    modifier onlyAfter() {
        require(now >= timestampMint + 3 weeks);
        _;
    }

    /**
    * @dev Checks if account has staff-level
    */
    modifier hasPermission(uint _level) {
        require(permissonedAccounts[msg.sender] > 0);
        require(permissonedAccounts[msg.sender] <= _level);
        _;
    }

    /**
    * @dev total number of tokens in existence
    */
    function circulatingSupply() public view returns (uint) {
        return circulatingSupply_;
    }

    /**
    * @dev Gets balance of address
    * @param _owner The address to query the the balance of.
    * @return An uint representing the amount owned by the passed address.
    */
    function balanceOf(address _owner) public view returns (uint balance) {
        return balances[_owner];
    }

    /**
    * @dev Token-transfer from msg.sender to address
    * @param _to target-address
    * @param _value amount of WHO transfered
    */
    function transfer(address _to, uint _value) public returns (bool) {
        require(_to != address(0));
        require(_value <= balances[msg.sender]);
        balances[msg.sender] = balances[msg.sender].sub(_value);
        balances[_to] = balances[_to].add(_value);
        TransferEvent(msg.sender, _to, _value);
        return true;
    }
}


/**
 * @title The Who-Token by WhoHas
 * @author Felix Leber, Christian Siegert
 * @dev Special version of the ERC20 Token
 */
contract Who is StandardToken {

    mapping (address => uint) public votings_;
    mapping (address => uint8) public icoAccounts;
    address public prizePool;
    uint public icoPool;
    uint public raisedIcoValue;
    uint public maxMint;


    event WinningEvent(address[] winner, address contest, uint payoutValue);
    event VotingStarted(address _voting, uint _duration, uint _costPerVote);
    event ParticipatedInVoting(address _sender, address _votingContract, bytes32 _hash, uint _voteAmount);

    modifier icoPhase() {
        require(now >= timestampRelease);
        require(now <= 32 days + timestampRelease);
        require(msg.value >= 2*(10**16));
        _;

    }

    function Who() public {
        owner = 0x4c556b28A7D62D3b7A84481521308fbb9687f38F;

        name = "WhoHas";
        symbol = "WHO";
        decimals = 18;

        permissonedAccounts[owner] = 1;
        timestampRelease = now + 4 hours + 2 days;

        balances[0x4c556b28A7D62D3b7A84481521308fbb9687f38F] = 150000000*(10**18); //150 Millionen
        icoPool = 100000000*(10**18); //100 Millionen
        maxSupply_ = 1500000000*(10**18); //1,5 Billion
        maxMint = 150000*(10**18); //150 k
        circulatingSupply_ = circulatingSupply_.add(balances[msg.sender]).add(icoPool); //250 Million
    }

    /**
    * @dev Buy option during ICO, payable
    * @notice Please make sure that ICO Pool is at least equal to your bid
    */
    function icoBuy() public icoPhase() payable {
        prizePool.transfer(msg.value);
        raisedIcoValue = raisedIcoValue.add(msg.value);
        uint256 tokenAmount = calculateTokenAmountICO(msg.value);

        require(icoPool >= tokenAmount);

        icoPool = icoPool.sub(tokenAmount);
        balances[msg.sender] += tokenAmount;
        TransferEvent(prizePool, msg.sender, tokenAmount);
    }

    /**
    * @dev Calculation of Token Ratio in ICO
    * @param _etherAmount Amount in Ether in order to be spent on WHO Token
    */
    function calculateTokenAmountICO(uint256 _etherAmount) public icoPhase constant returns(uint256) {
          // ICO standard rate: 1 ETH : 3315 WHO - 0,20 Euro
          // ICO Phase 1:   1 ETH : 4420 WHO - 0,15 Euro
        if (now <= 10 days + timestampRelease) {
            require(icoAccounts[msg.sender] == 1);
            return _etherAmount.mul(4420);
        } else {
            require(icoAccounts[msg.sender] == 2);
            return _etherAmount.mul(3315);
        }
    }

    /**
    * @dev Set/Unset address as permissioned
    * @param _account The address to give/take away the permissiones.
    * @param _level Permission-Level: 7:none, 1: owner, 2: admin, 3: pyFactory
    */
    function updatePermissions(address _account, uint _level) public hasPermission(1) {
        require(_level != 1 && msg.sender != _account);
        permissonedAccounts[_account] = _level;
    }

    /**
    * @dev Update Address recieving & distributing tokens in votings
    * @param _account Address of the new prize Pool
    */
    function updatePrizePool(address _account) public hasPermission(1) {
        prizePool = _account;
    }

    /**
    * @dev Increases circulatingSupply_ by specified amount. Available every three weeks until maxSupply_ is reached.
    * @param _mintAmount Amount of increase, must be smaller than 100000000
    */
    function mint(uint _mintAmount) public onlyAfter hasPermission(2) {
        require(_mintAmount <= maxMint);
        require(circulatingSupply_ + _mintAmount <= maxSupply_);
        balances[owner] = balances[owner].add(_mintAmount);
        circulatingSupply_ = circulatingSupply_.add(_mintAmount);
        timestampMint = now;
    }

    function registerForICO(address[] _icoAddresses, uint8 _level) public hasPermission(3) {
        for (uint i = 0; i < _icoAddresses.length; i++) {
            icoAccounts[_icoAddresses[i]] = _level;
        }
    }

    /**
    * @dev Manually add an existing WhoVote contract
    * @param _timespan Amount of time the contract is valid
    * @param _votePrice Price in Who(x10^18) per Vote
    */
    function gernerateVoting(uint _timespan, uint _votePrice) public hasPermission(3) {
        require(_votePrice > 0 && _timespan > 0);
        address generatedVoting = new WhoVote(this, _timespan);
        votings_[generatedVoting] = _votePrice;
        VotingStarted(generatedVoting, _timespan, _votePrice);
    }

    /**
    * @dev Manually add an existing WhoVote contract
    * @param _votingContract Adress of Voting-Contrac
    * @param _votePrice Price in Who(x10^18) per Vote
    */
    function addVoting(address _votingContract, uint _votePrice) public hasPermission(3) {
        votings_[_votingContract] = _votePrice;
    }

    /**
    * @dev Disable voting
    * @param _votingContract Adress of Voting-Contract
    */
    function finalizeVoting(address _votingContract) public hasPermission(3) {
        votings_[_votingContract] = 0;
    }

    /**
    * @dev PyFactory payout of winner
    * @param _winner Account which paricipated in the voting
    * @param _payoutValue Amount of Who payed to the winning account
    * @param _votingAddress Address of the Voting-Contract
    */
    function payout(address[] _winner, uint _payoutValue, address _votingAddress) public hasPermission(3) {
        for (uint i = 0; i < _winner.length; i++) {
            transfer(_winner[i], _payoutValue);
        }
        WinningEvent(_winner, _votingAddress, _payoutValue);
    }

    /**
    * @dev Participating in a Voting
    * @param _votingContract Adress of Voting-Contract
    * @param _hash Hash of the JSON-Parameter
    * @param _quantity Quantity of Votes
    */
    function payForVote(address _votingContract, bytes32 _hash, uint _quantity) public {
        require(_quantity >= 1 && _quantity <= 5);
        uint votePrice = votings_[_votingContract];
        require(votePrice > 0);
        transfer(prizePool, _quantity.mul(votePrice));
        sendVote(_votingContract, msg.sender, _hash);
        ParticipatedInVoting(msg.sender, _votingContract, _hash, _quantity);
    }

    /**
    * @dev [Internal] Send vote to Voting-Contract
    * @param _contract Address of Voting-Contract
    * @param _sender Sender of Votes
    * @param _hash Hash of the JSON-Parameter
    */
    function sendVote(address _contract, address _sender, bytes32 _hash) private returns (bool) {
        return WhoVote(_contract).recieveVote(_sender, _hash);
    }

}