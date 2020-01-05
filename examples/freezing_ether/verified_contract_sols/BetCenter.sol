pragma solidity ^0.4.18;

// zeppelin-solidity: 1.8.0

contract DataCenterInterface {
  function getResult(bytes32 gameId) view public returns (uint16, uint16, uint8);
}

contract DataCenterAddrResolverInterface {
  function getAddress() public returns (address _addr);
}

contract DataCenterBridge {
  uint8 constant networkID_auto = 0;
  uint8 constant networkID_mainnet = 1;
  uint8 constant networkID_testnet = 3;
  string public networkName;

  address public mainnetAddr = 0x6690E2698Bfa407DB697E69a11eA56810454549b;
  address public testnetAddr = 0x282b192518fc09568de0E66Df8e2533f88C16672;

  DataCenterAddrResolverInterface DAR;

  DataCenterInterface dataCenter;

  modifier dataCenterAPI() {
    if((address(DAR) == 0) || (getCodeSize(address(DAR)) == 0))
      setNetwork(networkID_auto);
    if(address(dataCenter) != DAR.getAddress())
      dataCenter = DataCenterInterface(DAR.getAddress());
    _;
  }

  /**
   * @dev set network will indicate which net will be used
   * @notice comment out `networkID` to avoid &#39;unused parameter&#39; warning
   */
  function setNetwork(uint8 /*networkID*/) internal returns(bool){
    return setNetwork();
  }

  function setNetwork() internal returns(bool){
    if (getCodeSize(mainnetAddr) > 0) {
      DAR = DataCenterAddrResolverInterface(mainnetAddr);
      setNetworkName("eth_mainnet");
      return true;
    }
    if (getCodeSize(testnetAddr) > 0) {
      DAR = DataCenterAddrResolverInterface(testnetAddr);
      setNetworkName("eth_ropsten");
      return true;
    }
    return false;
  }

  function setNetworkName(string _networkName) internal {
    networkName = _networkName;
  }

  function getNetworkName() internal view returns (string) {
    return networkName;
  }

  function dataCenterGetResult(bytes32 _gameId) dataCenterAPI internal returns (uint16, uint16, uint8){
    return dataCenter.getResult(_gameId);
  }

  function getCodeSize(address _addr) view internal returns (uint _size) {
    assembly {
      _size := extcodesize(_addr)
    }
  }
}

/**
 * @title Ownable
 * @dev The Ownable contract has an owner address, and provides basic authorization control
 * functions, this simplifies the implementation of "user permissions".
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
 * @title SafeMath
 * @dev Math operations with safety checks that throw on error
 */
library SafeMath {

  /**
  * @dev Multiplies two numbers, throws on overflow.
  */
  

  /**
  * @dev Integer division of two numbers, truncating the quotient.
  */
  

  /**
  * @dev Subtracts two numbers, throws on overflow (i.e. if subtrahend is greater than minuend).
  */
  

  /**
  * @dev Adds two numbers, throws on overflow.
  */
  
}

contract Bet is Ownable, DataCenterBridge {
  using SafeMath for uint;

  event LogDistributeReward(address addr, uint reward, uint index);
  event LogGameResult(bytes32 indexed category, bytes32 indexed gameId, uint leftPts, uint rightPts);
  event LogParticipant(address addr, uint choice, uint betAmount);
  event LogRefund(address addr, uint betAmount);
  event LogBetClosed(bool isRefund, uint timestamp);
  event LogDealerWithdraw(address addr, uint withdrawAmount);

  /** 
   * @desc
   * gameId: is a fixed string just like "0021701030"
   *   the full gameId encode(include football, basketball, esports..) will publish on github
   * leftOdds: need divide 100, if odds is 216 means 2.16
   * middleOdds: need divide 100, if odds is 175 means 1.75
   * rightOdds: need divide 100, if odds is 250 means 2.50
   * spread: need sub 0.5, if spread is 1 means 0.5, 0 means no spread
   * flag: indicate which team get spread, 1 means leftTeam, 3 means rightTeam
   */
  struct BetInfo {
    bytes32 category;
    bytes32 gameId;
    uint8   spread;
    uint8   flag;
    uint16  leftOdds;
    uint16  middleOdds;
    uint16  rightOdds;
    uint    minimumBet;
    uint    startTime;
    uint    deposit;
    address dealer;
  }

  struct Player {
    uint betAmount;
    uint choice;
  }

  /**
   * @desc
   * winChoice: Indicate the winner choice of this betting
   *   1 means leftTeam win, 3 means rightTeam win, 2 means draw(leftTeam is not always equivalent to the home team)
   */
  uint8 public winChoice;
  uint8 public confirmations = 0;
  uint8 public neededConfirmations = 1;
  uint16 public leftPts;
  uint16 public rightPts;
  bool public isBetClosed = false;

  uint public totalBetAmount = 0;
  uint public leftAmount;
  uint public middleAmount;
  uint public rightAmount;
  uint public numberOfBet;

  address [] public players;
  mapping(address => Player) public playerInfo;

  /**
   * @dev Throws if called by any account other than the dealer
   */
  modifier onlyDealer() {
    require(msg.sender == betInfo.dealer);
    _;
  }

  

  BetInfo betInfo;

  function Bet(address _dealer, bytes32 _category, bytes32 _gameId, uint _minimumBet, 
                  uint8 _spread, uint16 _leftOdds, uint16 _middleOdds, uint16 _rightOdds, uint8 _flag,
                  uint _startTime, uint8 _neededConfirmations, address _owner) payable public {
    require(_flag == 1 || _flag == 3);
    require(_startTime > now);
    require(msg.value >= 0.1 ether);
    require(_neededConfirmations >= neededConfirmations);

    betInfo.dealer = _dealer;
    betInfo.deposit = msg.value;
    betInfo.flag = _flag;
    betInfo.category = _category;
    betInfo.gameId = _gameId;
    betInfo.minimumBet = _minimumBet;
    betInfo.spread = _spread;
    betInfo.leftOdds = _leftOdds;
    betInfo.middleOdds = _middleOdds;
    betInfo.rightOdds = _rightOdds;
    betInfo.startTime = _startTime;

    neededConfirmations = _neededConfirmations;
    owner = _owner;
  }

  /**
   * @dev get basic information of this bet
   */
  

  /**
   * @dev get basic information of this bet
   *
   *  uint public numberOfBet;
   *  uint public totalBetAmount = 0;
   *  uint public leftAmount;
   *  uint public middleAmount;
   *  uint public rightAmount;
   *  uint public deposit;
   */
  

  /**
   * @dev get bet result information
   *
   *  uint8 public winChoice;
   *  uint8 public confirmations = 0;
   *  uint8 public neededConfirmations = 1;
   *  uint16 public leftPts;
   *  uint16 public rightPts;
   *  bool public isBetClosed = false;
   */
  

  /**
   * @dev calculate the gas whichdistribute rewards will cost
   * set default gasPrice is 5000000000
   */
  

  /**
   * @dev find a player has participanted or not
   * @param player the address of the participant
   */
  

  /**
   * @dev to check the dealer is solvent or not
   * @param choice indicate which team user choose
   * @param amount indicate how many ether user bet
   */
  

  /**
   * @dev update this bet some state
   * @param choice indicate which team user choose
   * @param amount indicate how many ether user bet
   */
  

  /**
   * @dev place a bet with his/her choice
   * @param choice indicate which team user choose
   */
  

  /**
   * @dev in order to let more people participant, dealer can recharge
   */
  

  /**
   * @dev given game result, _return win choice by specific spread
   */
  

  /**
   * @dev manualCloseBet could only be called by owner,
   *      this method only be used for ropsten,
   *      when ethereum-events-data deployed,
   *      game result should not be upload by owner
   */
  

  /**
   * @dev closeBet could be called by everyone, but owner/dealer should to this.
   */
  

  /**
   * @dev get the players
   */
  

  /**
   * @dev get contract balance
   */
  

  /**
   * @dev if there are some reasons lead game postpone or cancel
   *      the bet will also cancel and refund every bet
   */
  

  /**
   * @dev dealer can withdraw the remain ether after refund or closed
   */
  

  /**
   * @dev distribute ether to every winner as they choosed odds
   */
  
}

contract BetCenter is Ownable {

  mapping(bytes32 => Bet[]) public bets;
  mapping(bytes32 => bytes32[]) public gameIds;

  event LogCreateBet(address indexed dealerAddr, address betAddr, bytes32 indexed category, uint indexed startTime);

  function() payable public {}

  function createBet(bytes32 category, bytes32 gameId, uint minimumBet, 
                  uint8 spread, uint16 leftOdds, uint16 middleOdds, uint16 rightOdds, uint8 flag,
                  uint startTime, uint8 confirmations) payable public {
    Bet bet = (new Bet).value(msg.value)(msg.sender, category, gameId, minimumBet, 
                  spread, leftOdds, middleOdds, rightOdds , flag, startTime, confirmations, owner);
    bets[category].push(bet);
    gameIds[category].push(gameId);
    LogCreateBet(msg.sender, bet, category, startTime);
  }

  /**
   * @dev fetch bets use category
   * @param category Indicate the sports events type
   */
  function getBetsByCategory(bytes32 category) view public returns (Bet[]) {
    return bets[category];
  }

  function getGameIdsByCategory(bytes32 category) view public returns (bytes32 []) {
    return gameIds[category];
  }

}