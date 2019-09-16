/**
 *Submitted for verification at Etherscan.io on 2018-02-06
*/

pragma solidity ^0.4.18;

// import "./Pausable.sol";
// import "./CratePreSale.sol";

contract NewCratePreSale {
    
    // migration functions migrate the data from the previous contract in stages
    // all addresses are included for transparency and easy verification
    // however addresses with no robots (i.e. failed transaction and never bought properly) have been commented out.
    // to view the full list of state assignments, go to etherscan.io/address/{address} and you can view the verified
    mapping (address => uint[]) public userToRobots; 

    function _migrate(uint _index) external onlyOwner {
        bytes4 selector = bytes4(sha3("setData()"));
        address a = migrators[_index];
        require(a.delegatecall(selector));
    }
    // source code - feel free to verify the migration
    address[6] migrators = [
        0x700febd9360ac0a0a72f371615427bec4e4454e5, //0x97ae01893e42d6d33fd9851a28e5627222af7bbb,
        0x72cc898de0a4eac49c46ccb990379099461342f6,
        0xc3cc48da3b8168154e0f14bf0446c7a93613f0a7,
        0x4cc96f2ddf6844323ae0d8461d418a4d473b9ac3,
        0xa52bfcb5ff599e29ee2b9130f1575babaa27de0a,
        0xe503b42aabda22974e2a8b75fa87e010e1b13584
    ];
    
    function NewCratePreSale() public payable {
        
            owner = msg.sender;
        // one time transfer of state from the previous contract
        // var previous = CratePreSale(0x3c7767011C443EfeF2187cf1F2a4c02062da3998); //MAINNET

        // oldAppreciationRateWei = previous.appreciationRateWei();
        oldAppreciationRateWei = 100000000000000;
        appreciationRateWei = oldAppreciationRateWei;
  
        // oldPrice = previous.currentPrice();
        oldPrice = 232600000000000000;
        currentPrice = oldPrice;

        // oldCratesSold = previous.cratesSold();
        oldCratesSold = 1075;
        cratesSold = oldCratesSold;

        // Migration Rationale
        // due to solidity issues with enumerability (contract calls cannot return dynamic arrays etc)
        // no need for trust -> can still use web3 to call the previous contract and check the state
        // will only change in the future if people send more eth
        // and will be obvious due to change in crate count. Any purchases on the old contract
        // after this contract is deployed will be fully refunded, and those robots bought will be voided. 
        // feel free to validate any address on the old etherscan:
        // https://etherscan.io/address/0x3c7767011C443EfeF2187cf1F2a4c02062da3998
        // can visit the exact contracts at the addresses listed above
    }

    // ------ STATE ------
    uint256 constant public MAX_CRATES_TO_SELL = 3900; // Max no. of robot crates to ever be sold
    uint256 constant public PRESALE_END_TIMESTAMP = 1518699600; // End date for the presale - no purchases can be made after this date - Midnight 16 Feb 2018 UTC

    uint256 public appreciationRateWei;
    uint32 public cratesSold;
    uint256 public currentPrice;

    // preserve these for later verification
    uint32 public oldCratesSold;
    uint256 public oldPrice;
    uint256 public oldAppreciationRateWei;
    // mapping (address => uint32) public userCrateCount; // replaced with more efficient method
    

    // store the unopened crates of this user
    // actually stores the blocknumber of each crate 
    mapping (address => uint[]) public addressToPurchasedBlocks;
    // store the number of expired crates for each user 
    // i.e. crates where the user failed to open the crate within 256 blocks (~1 hour)
    // these crates will be able to be opened post-launch
    mapping (address => uint) public expiredCrates;
    // store the part information of purchased crates



    function openAll() public {
        uint len = addressToPurchasedBlocks[msg.sender].length;
        require(len > 0);
        uint8 count = 0;
        // len > i to stop predicatable wraparound
        for (uint i = len - 1; i >= 0 && len > i; i--) {
            uint crateBlock = addressToPurchasedBlocks[msg.sender][i];
            require(block.number > crateBlock);
            // can't open on the same timestamp
            var hash = block.blockhash(crateBlock);
            if (uint(hash) != 0) {
                // different results for all different crates, even on the same block/same user
                // randomness is already taken care of
                uint rand = uint(keccak256(hash, msg.sender, i)) % (10 ** 20);
                userToRobots[msg.sender].push(rand);
                count++;
            } else {
                // all others will be expired
                expiredCrates[msg.sender] += (i + 1);
                break;
            }
        }
        CratesOpened(msg.sender, count);
        delete addressToPurchasedBlocks[msg.sender];
    }

    // ------ EVENTS ------
    event CratesPurchased(address indexed _from, uint8 _quantity);
    event CratesOpened(address indexed _from, uint8 _quantity);

    // ------ FUNCTIONS ------
    function getPrice() view public returns (uint256) {
        return currentPrice;
    }

    function getRobotCountForUser(address _user) external view returns(uint256) {
        return userToRobots[_user].length;
    }

    function getRobotForUserByIndex(address _user, uint _index) external view returns(uint) {
        return userToRobots[_user][_index];
    }

    function getRobotsForUser(address _user) view public returns (uint[]) {
        return userToRobots[_user];
    }

    function getPendingCratesForUser(address _user) external view returns(uint[]) {
        return addressToPurchasedBlocks[_user];
    }

    function getPendingCrateForUserByIndex(address _user, uint _index) external view returns(uint) {
        return addressToPurchasedBlocks[_user][_index];
    }

    function getExpiredCratesForUser(address _user) external view returns(uint) {
        return expiredCrates[_user];
    }

    function incrementPrice() private {
        // Decrease the rate of increase of the crate price
        // as the crates become more expensive
        // to avoid runaway pricing
        // (halving rate of increase at 0.1 ETH, 0.2 ETH, 0.3 ETH).
        if ( currentPrice == 100000000000000000 ) {
            appreciationRateWei = 200000000000000;
        } else if ( currentPrice == 200000000000000000) {
            appreciationRateWei = 100000000000000;
        } else if (currentPrice == 300000000000000000) {
            appreciationRateWei = 50000000000000;
        }
        currentPrice += appreciationRateWei;
    }

    function purchaseCrates(uint8 _cratesToBuy) public payable whenNotPaused {
        require(now < PRESALE_END_TIMESTAMP); // Check presale is still ongoing.
        require(_cratesToBuy <= 10); // Can only buy max 10 crates at a time. Don't be greedy!
        require(_cratesToBuy >= 1); // Sanity check. Also, you have to buy a crate. 
        require(cratesSold + _cratesToBuy <= MAX_CRATES_TO_SELL); // Check max crates sold is less than hard limit
        uint256 priceToPay = _calculatePayment(_cratesToBuy);
         require(msg.value >= priceToPay); // Check buyer sent sufficient funds to purchase
        if (msg.value > priceToPay) { //overpaid, return excess
            msg.sender.transfer(msg.value-priceToPay);
        }
        //all good, payment received. increment number sold, price, and generate crate receipts!
        cratesSold += _cratesToBuy;
      for (uint8 i = 0; i < _cratesToBuy; i++) {
            incrementPrice();
            addressToPurchasedBlocks[msg.sender].push(block.number);
        }

        CratesPurchased(msg.sender, _cratesToBuy);
    } 

    function _calculatePayment (uint8 _cratesToBuy) private view returns (uint256) {
        
        uint256 tempPrice = currentPrice;

        for (uint8 i = 1; i < _cratesToBuy; i++) {
            tempPrice += (currentPrice + (appreciationRateWei * i));
        } // for every crate over 1 bought, add current Price and a multiple of the appreciation rate
          // very small edge case of buying 10 when you the appreciation rate is about to halve
          // is compensated by the great reduction in gas by buying N at a time.
        
        return tempPrice;
    }


    //owner only withdrawal function for the presale
    function withdraw() onlyOwner public {
        owner.transfer(this.balance);
    }

    function addFunds() onlyOwner external payable {

    }

  event SetPaused(bool paused);

  // starts unpaused
  bool public paused = false;

  modifier whenNotPaused() {
    require(!paused);
    _;
  }

  modifier whenPaused() {
    require(paused);
    _;
  }

  function pause() external onlyOwner whenNotPaused returns (bool) {
    paused = true;
    SetPaused(paused);
    return true;
  }

  function unpause() external onlyOwner whenPaused returns (bool) {
    paused = false;
    SetPaused(paused);
    return true;
  }


  address public owner;

  event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);




  modifier onlyOwner() {
    require(msg.sender == owner);
    _;
  }

  function transferOwnership(address newOwner) public onlyOwner {
    require(newOwner != address(0));
    OwnershipTransferred(owner, newOwner);
    owner = newOwner;
  }
    
}
