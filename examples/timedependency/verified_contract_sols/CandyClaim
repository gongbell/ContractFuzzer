/**
 *Submitted for verification at Etherscan.io on 2018-02-21
*/

pragma solidity ^0.4.18;

/// candy.claims

contract CandyClaim {
  /*** CONSTANTS ***/
  uint256 private fiveHoursInSeconds = 3600; // 18000;
  string public constant NAME = "CandyClaims";
  string public constant SYMBOL = "CandyClaim";

  /*** STORAGE ***/
  mapping (address => uint256) private ownerCount;

  address public ceoAddress;
  address public cooAddress;

  struct Candy {
    address owner;
    uint256 price;
    uint256 last_transaction;
    address approve_transfer_to;
  }
  uint candy_count;
  mapping (string => Candy) candies;

  /*** ACCESS MODIFIERS ***/
  modifier onlyCEO() { require(msg.sender == ceoAddress); _; }
  modifier onlyCOO() { require(msg.sender == cooAddress); _; }
  modifier onlyCXX() { require(msg.sender == ceoAddress || msg.sender == cooAddress); _; }

  /*** ACCESS MODIFIES ***/
  function setCEO(address _newCEO) public onlyCEO {
    require(_newCEO != address(0));
    ceoAddress = _newCEO;
  }
  function setCOO(address _newCOO) public onlyCEO {
    require(_newCOO != address(0));
    cooAddress = _newCOO;
  }

  /*** DEFAULT METHODS ***/
  function symbol() public pure returns (string) { return SYMBOL; }
  function name() public pure returns (string) { return NAME; }
  function implementsERC721() public pure returns (bool) { return true; }

  /*** CONSTRUCTOR ***/
  function CandyClaim() public {
    ceoAddress = msg.sender;
    cooAddress = msg.sender;
  }

  /*** INTERFACE METHODS ***/
  function createCandy(string _candy_id, uint256 _price) public onlyCXX {
    require(msg.sender != address(0));
    _create_candy(_candy_id, address(this), _price);
  }

  function totalSupply() public view returns (uint256 total) {
    return candy_count;
  }

  function balanceOf(address _owner) public view returns (uint256 balance) {
    return ownerCount[_owner];
  }
  function priceOf(string _candy_id) public view returns (uint256 price) {
    return candies[_candy_id].price;
  }

  function getCandy(string _candy_id) public view returns (
    string id,
    address owner,
    uint256 price,
    uint256 last_transaction
  ) {
    id = _candy_id;
    owner = candies[_candy_id].owner;
    price = candies[_candy_id].price;
    last_transaction = candies[_candy_id].last_transaction;
  }

  function purchase(string _candy_id) public payable {
    Candy storage candy = candies[_candy_id];

    require(candy.owner != msg.sender);
    require(msg.sender != address(0));

    uint256 time_diff = (block.timestamp - candy.last_transaction);
    while(time_diff >= fiveHoursInSeconds){
        time_diff = (time_diff - fiveHoursInSeconds);
        candy.price = SafeMath.mul(SafeMath.div(candy.price, 100), 90);
    }
    if(candy.price < 1000000000000000){ candy.price = 1000000000000000; }
    require(msg.value >= candy.price);

    uint256 excess = SafeMath.sub(msg.value, candy.price);

    if(candy.owner == address(this)){
      ceoAddress.transfer(candy.price);
    } else {
      ceoAddress.transfer(uint256(SafeMath.mul(SafeMath.div(candy.price, 100), 10)));
      candy.owner.transfer(uint256(SafeMath.mul(SafeMath.div(candy.price, 100), 90)));
    }

    candy.price = SafeMath.mul(SafeMath.div(candy.price, 100), 160);
    candy.owner = msg.sender;
    candy.last_transaction = block.timestamp;

    msg.sender.transfer(excess);
  }

  function payout() public onlyCEO {
    ceoAddress.transfer(this.balance);
  }

  /*** PRIVATE METHODS ***/

  function _create_candy(string _candy_id, address _owner, uint256 _price) private {
    candy_count++;
    candies[_candy_id] = Candy({
      owner: _owner,
      price: _price,
      last_transaction: block.timestamp,
      approve_transfer_to: address(0)
    });
  }

  function _transfer(address _from, address _to, string _candy_id) private {
    candies[_candy_id].owner = _to;
    candies[_candy_id].approve_transfer_to = address(0);
    ownerCount[_from] -= 1;
    ownerCount[_to] += 1;
  }
}

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
    uint256 c = a / b;
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
