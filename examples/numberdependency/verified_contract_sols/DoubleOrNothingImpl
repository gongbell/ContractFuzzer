pragma solidity ^0.4.11;

// SafeMath handles math with overflow.
contract SafeMath {
    function safeAdd(uint256 x, uint256 y) internal returns(uint256) {
        uint256 z = x + y;
        assert((z >= x) && (z >= y));
        return z;
    }

    function safeSubtract(uint256 x, uint256 y) internal returns(uint256) {
        assert(x >= y);
        uint256 z = x - y;
        return z;
    }

    function safeMult(uint256 x, uint256 y) internal returns(uint256) {
        uint256 z = x * y;
        assert((x == 0)||(z/x == y));
        return z;
    }
}

// Random is a block hash based random number generator.
contract Random {
    // Generates a random number from 0 to max based on the last block hash.
    function getRand(uint blockNumber, uint max) constant returns(uint) {
        return(uint(sha3(block.blockhash(blockNumber))) % max);
    }
}

// Manages contract ownership.
contract Owned {
    address public owner;
    function owned() {
        owner = msg.sender;
    }
    modifier onlyOwner {
        if (msg.sender != owner) throw;
        _;
    }
    function transferOwnership(address newOwner) onlyOwner {
        owner = newOwner;
    }
}

// DoubleOrNothing is the main public interface for gambling.
// To gamble:
//   Step 1: Send up to maxWagerEth ETH to this contract address.
//   Step 2: Wait waitTimeBlocks.
//   Step 3: Call payout() to receive your payment, if any.
contract DoubleOrNothing {
    // maxWagerWei is the maximum wager in Wei.
    uint256 public maxWagerWei;
    
    // waitTime is the number of blocks before payout is available.
    uint public waitTimeBlocks;
    
    // payoutOdds is the value / 10000 that a payee will win a wager.
    // eg. payoutOdds of 4950 implies a 49.5% chance of winning.
    uint public payoutOdds;
    
    // Wager represents one wager.
    struct Wager {
        address sender;
        uint256 wagerWei;
        uint256 creationBlockNumber;
        bool active;
    }
    
    // wagers contains all current outstanding wagers.
    // TODO: Support multiple Wagers per address.
    mapping (address => Wager) wagers;
    
    function makeWager() payable public;
    function payout() public;
}

contract DoubleOrNothingImpl is DoubleOrNothing, Owned, Random, SafeMath {
    
    // Initialize state by assigning the owner to the contract deployer.
    function DoubleOrNothingImpl() {
        owner = msg.sender;
        maxWagerWei = 100000000000000000;
        waitTimeBlocks = 2;
        payoutOdds = 4950;
    }
    
    // Allow the owner to set maxWagerWei.
    function setMaxWagerWei(uint256 maxWager) public onlyOwner {
        maxWagerWei = maxWager;
    }
    
    // Allow the owner to set waitTimeBlocks.
    function setWaitTimeBlocks(uint waitTime) public onlyOwner {
        waitTimeBlocks = waitTime;
    }
    
    // Allow the owner to set payoutOdds.
    function setPayoutOdds(uint odds) public onlyOwner {
        payoutOdds = odds;
    }
    
    // Allow the owner to cash out the holdings of this contract.
    function withdraw(address recipient, uint256 balance) public onlyOwner {
        recipient.transfer(balance);
    }
    
    // View your wager.
    function getWagerOwner(address wager_owner) constant public returns (
        uint256 wagerWei,
        uint creationBlockNumber,
        bool active) {
        return _getWager(wager_owner);
    }
    
    // Allow the owner to payout outstanding wagers on others' behalf.
    function ownerPayout(address wager_owner) public onlyOwner {
        _payout(wager_owner);
    }
    
    // Assume that simple transactions are trying to make a wager, unless it is
    // from the owner.
    function () payable public {
        if (msg.sender != owner) {
            makeWager();
        }
    }
    
    // Make a wager.
    function makeWager() payable public {
        if (msg.value == 0 || msg.value > maxWagerWei) throw;
        if (wagers[msg.sender].active) {
            // A Wager already exists for this user.
            throw;
        }
        wagers[msg.sender] = Wager({
            sender: msg.sender,
            wagerWei: msg.value,
            creationBlockNumber: block.number,
            active: true,
        });
    }
    
    // View your wager.
    function getWager() constant public returns (
        uint256 wagerWei,
        uint creationBlockNumber,
        bool active) {
        return _getWager(msg.sender);
    }
    
    // Payout any wagers associated with the sending address.
    function payout() public {
        _payout(msg.sender);
    }
    
    // Internal implementation of getWager().
    function _getWager(address wager_owner) constant public returns (
        uint256 wagerWei,
        uint creationBlockNumber,
        bool active) {
        Wager thisWager = wagers[wager_owner];
        return (thisWager.wagerWei, thisWager.creationBlockNumber, thisWager.active);
    }
    
    // Internal implementation of payout().
    function _payout(address wager_owner) internal {
        if (!wagers[wager_owner].active) {
            // No outstanding active Wager.
            throw;
        }
        uint256 blockDepth = block.number - wagers[wager_owner].creationBlockNumber;
        if (blockDepth > waitTimeBlocks) {
            // waitTimeBlocks has passed, resolve and payout this wager.
            uint256 payoutBlock = wagers[wager_owner].creationBlockNumber + waitTimeBlocks - 1;
            uint randNum = getRand(payoutBlock, 10000);
            if (randNum < payoutOdds) {
                // Wager wins, payout wager.
                uint256 winnings = safeMult(wagers[wager_owner].wagerWei, 2);
                if (wagers[wager_owner].sender.send(winnings)) {
                    wagers[wager_owner].active = false;
                }
            } else {
                // Wager loses, disable wager.
                wagers[wager_owner].active = false;
            }
        }
    }
}