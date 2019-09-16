/**
 *Submitted for verification at Etherscan.io on 2018-02-19
*/

pragma solidity 0.4.19;

/**
 * @title SafeMath
 * @dev Math operations with safety checks that throw on error
 */
library SafeMath {
    /**
    * @dev Multiplies two numbers, throws on overflow.
    */
    function mul(uint a, uint b) internal pure returns (uint) {
        if (a == 0) {
        return 0;
        }
        uint c = a * b;
        assert(c / a == b);
        return c;
    }

    /**
    * @dev Integer division of two numbers, truncating the quotient.
    */
    function div(uint a, uint b) internal pure returns (uint) {
        // assert(b > 0); // Solidity automatically throws when dividing by 0
        uint c = a / b;
        // assert(a == b * c + a % b); // There is no case in which this doesn't hold
        return c;
    }

    /**
    * @dev Substracts two numbers, throws on overflow (i.e. if subtrahend is greater than minuend).
    */
    function sub(uint a, uint b) internal pure returns (uint) {
        assert(b <= a);
        return a - b;
    }

    /**
    * @dev Adds two numbers, throws on overflow.
    */
    function add(uint a, uint b) internal pure returns (uint) {
        uint c = a + b;
        assert(c >= a);
        return c;
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
 * @title Heritable
 * @dev The Heritable contract provides ownership transfer capabilities, in the
 * case that the current owner stops "heartbeating". Only the heir can pronounce the
 * owner's death.
 */
contract Heritable is Ownable {
    address public heir;

    // Time window the owner has to notify they are alive.
    uint public heartbeatTimeout;

    // Timestamp of the owner's death, as pronounced by the heir.
    uint public timeOfDeath;

    event HeirChanged(address indexed owner, address indexed newHeir);
    event OwnerHeartbeated(address indexed owner);
    event OwnerProclaimedDead(address indexed owner, address indexed heir, uint timeOfDeath);
    event HeirOwnershipClaimed(address indexed previousOwner, address indexed newOwner);


    /**
    * @dev Throw an exception if called by any account other than the heir's.
    */
    modifier onlyHeir() {
        require(msg.sender == heir);
        _;
    }


    /**
    * @notice Create a new Heritable Contract with heir address 0x0.
    * @param _heartbeatTimeout time available for the owner to notify they are alive,
    * before the heir can take ownership.
    */
    function Heritable(uint _heartbeatTimeout) public {
        setHeartbeatTimeout(_heartbeatTimeout);
    }

    function setHeir(address newHeir) public onlyOwner {
        require(newHeir != owner);
        heartbeat();
        HeirChanged(owner, newHeir);
        heir = newHeir;
    }

    /**
    * @dev set heir = 0x0
    */
    function removeHeir() public onlyOwner {
        heartbeat();
        heir = 0;
    }

    /**
    * @dev Heir can pronounce the owners death. To claim the ownership, they will
    * have to wait for `heartbeatTimeout` seconds.
    */
    function proclaimDeath() public onlyHeir {
        require(owner != heir); // added
        require(ownerLives());
        OwnerProclaimedDead(owner, heir, timeOfDeath);
        timeOfDeath = now;
    }

    /**
    * @dev Owner can send a heartbeat if they were mistakenly pronounced dead.
    */
    function heartbeat() public onlyOwner {
        OwnerHeartbeated(owner);
        timeOfDeath = 0;
    }

    /**
    * @dev Allows heir to transfer ownership only if heartbeat has timed out.
    */
    function claimHeirOwnership() public onlyHeir {
        require(!ownerLives());
        require(now >= timeOfDeath + heartbeatTimeout);
        OwnershipTransferred(owner, heir);
        HeirOwnershipClaimed(owner, heir);
        owner = heir;
        timeOfDeath = 0;
    }

    function setHeartbeatTimeout(uint newHeartbeatTimeout) internal onlyOwner {
        require(ownerLives());
        heartbeatTimeout = newHeartbeatTimeout;
    }

    function ownerLives() internal view returns (bool) {
        return timeOfDeath == 0;
    }
}

/// @title Interface for contracts conforming to ERC-721: Non-Fungible Tokens
/// @author Dieter Shirley <dete@axiomzen.co> (https://github.com/dete)
contract ERC721 {
    // Required methods
    function approve(address _to, uint _tokenId) public;
    function balanceOf(address _owner) public view returns (uint balance);
    function implementsERC721() public pure returns (bool);
    function ownerOf(uint _tokenId) public view returns (address addr);
    function takeOwnership(uint _tokenId) public;
    function totalSupply() public view returns (uint total);
    function transferFrom(address _from, address _to, uint _tokenId) public;
    function transfer(address _to, uint _tokenId) public;

    event Transfer(address indexed from, address indexed to, uint tokenId);
    event Approval(address indexed owner, address indexed approved, uint tokenId);

    // Optional
    // function name() public view returns (string name);
    // function symbol() public view returns (string symbol);
    // function tokenOfOwnerByIndex(address _owner, uint _index) external view returns (uint tokenId);
    // function tokenMetadata(uint _tokenId) public view returns (string infoUrl);
}

contract BitArtToken is Heritable, ERC721 {
    string public constant NAME = "BitGallery";
    string public constant SYMBOL = "BitArt";

    struct Art {
        bytes32 data;
    }

    Art[] internal arts;

    mapping (uint => address) public tokenOwner;
    mapping (address => uint) public ownedTokenCount;
    mapping (uint => address) public tokenApprovals;

    event Transfer(address from, address to, uint tokenId);
    event Approval(address owner, address approved, uint tokenId);

    // 30 days to change owner
    function BitArtToken() Heritable(2592000) public {}

    function tokensOf(address _owner) external view returns(uint[]) {
        uint tokenCount = balanceOf(_owner);

        if (tokenCount == 0) {
            return new uint[](0);
        } else {
            uint[] memory result = new uint[](tokenCount);
            uint totaltokens = totalSupply();
            uint index = 0;
            
            for (uint tokenId = 0; tokenId < totaltokens; tokenId++) {
                if (tokenOwner[tokenId] == _owner) {
                    result[index] = tokenId;
                    index++;
                }
            }
            
            return result;
        }
    }

    function approve(address _to, uint _tokenId) public {
        require(_owns(msg.sender, _tokenId));
        tokenApprovals[_tokenId] = _to;
        Approval(msg.sender, _to, _tokenId);
    }

    function balanceOf(address _owner) public view returns (uint balance) {
        return ownedTokenCount[_owner];
    }

    function getArts() public view returns (bytes32[]) {
        uint count = totalSupply();
        bytes32[] memory result = new bytes32[](count);

        for (uint i = 0; i < count; i++) {
            result[i] = arts[i].data;
        }

        return result;
    }

    function implementsERC721() public pure returns (bool) {
        return true;
    }

    function name() public pure returns (string) {
        return NAME;
    }

    function ownerOf(uint _tokenId) public view returns (address owner) {
        owner = tokenOwner[_tokenId];
        require(_addressNotNull(owner));
    }

    function symbol() public pure returns (string) {
        return SYMBOL;
    }

    function takeOwnership(uint _tokenId) public {
        address newOwner = msg.sender;
        require(_addressNotNull(newOwner));
        require(_approved(newOwner, _tokenId));
        address oldOwner = tokenOwner[_tokenId];

        _transfer(oldOwner, newOwner, _tokenId);
    }

    function totalSupply() public view returns (uint total) {
        return arts.length;
    }

    function transfer(address _to, uint _tokenId) public {
        require(_owns(msg.sender, _tokenId));
        require(_addressNotNull(_to));

        _transfer(msg.sender, _to, _tokenId);
    }

    function transferFrom(address _from, address _to, uint _tokenId) public {
        require(_owns(_from, _tokenId));
        require(_approved(_to, _tokenId));
        require(_addressNotNull(_to));

        _transfer(_from, _to, _tokenId);
    }

    function _mint(address _to, uint256 _tokenId) internal {
        require(_to != address(0));
        require(tokenOwner[_tokenId] == address(0));

        _transfer(0x0, _to, _tokenId);
    }

    function _transfer(address _from, address _to, uint _tokenId) internal {
        require(_from != _to);
        ownedTokenCount[_to]++;
        tokenOwner[_tokenId] = _to;

        if (_addressNotNull(_from)) {
            ownedTokenCount[_from]--;
            delete tokenApprovals[_tokenId];
        }

        Transfer(_from, _to, _tokenId);
    }

    function _addressNotNull(address _address) private pure returns (bool) {
        return _address != address(0);
    }

    function _approved(address _to, uint _tokenId) private view returns (bool) {
        return tokenApprovals[_tokenId] == _to;
    }

    function _owns(address _claimant, uint _tokenId) private view returns (bool) {
        return _claimant == tokenOwner[_tokenId];
    }
}

contract BitAuction is BitArtToken {
    using SafeMath for uint;

    struct Auction {
        uint basePrice;
        uint64 time1;
        uint64 time2;
        uint8 pct1;
        uint8 pct2;
        uint8 discount;
    }

    uint internal _auctionStartsAfter;
    uint internal _auctionDuration;
    uint internal _auctionFee;

    mapping (uint => Auction) public tokenAuction;

    event AuctionRulesChanged(uint startsAfter, uint duration, uint fee);
    event NewAuction(uint tokenId, uint discount);
    event NewSaleDiscount(uint tokenId, uint discount);

    function BitAuction() public { }

    function setSaleDiscount(uint _tokenId, uint _discount) external {      
        require(ownerOf(_tokenId) == msg.sender);
        require(_discount <= 90);
        require(_discount >= 10);

        Auction storage auction = tokenAuction[_tokenId];
        require(auction.basePrice > 0);        
        require(auction.time2 <= now);
        auction.discount = uint8(_discount);

        NewSaleDiscount(_tokenId, _discount);
    }

    function canPurchase(uint _tokenId) public view returns (bool) {
        Auction storage auction = tokenAuction[_tokenId];
        require(auction.time1 > 0);
        return (now >= auction.time1 && priceOf(_tokenId) > 0);
    }

    function getPrices(uint[] _ids) public view returns (uint[]) {
        uint count = _ids.length;
        bool isEmpty = count == 0;

        if (isEmpty) {
            count = totalSupply();
        }

        uint[] memory result = new uint[](count);
        
        for (uint i = 0; i < count; i++) {
            uint tokenId = isEmpty ? i : _ids[i];
            result[i] = priceOf(tokenId);
        }        
        
        return result;
    }

    function priceOf(uint _tokenId) public view returns (uint) {
        Auction storage auction = tokenAuction[_tokenId];
        return _currentPrice(auction);
    }

    function setAuctionDurationRules(uint _timeAfter, uint _duration, uint _fee) public onlyOwner {  
        require(_timeAfter >= 0 seconds);
        require(_timeAfter <= 7 days);
        require(_duration >= 24 hours);
        require(_duration <= 30 days);
        require(_fee >= 1);
        require(_fee <= 5);
        
        _auctionStartsAfter = _timeAfter;
        _auctionDuration = _duration;
        _auctionFee = _fee;

        AuctionRulesChanged(_timeAfter, _duration, _fee);
    }

    function _createCustomAuction(uint _tokenId, uint _basePrice, uint _time1, uint _time2, uint _pct1, uint _pct2) private {
        require(_time1 >= now);
        require(_time2 >= _time1);
        require(_pct1 > 0);
        require(_pct2 > 0);
        
        Auction memory auction = Auction({
            basePrice: _basePrice, 
            time1: uint64(_time1), 
            time2: uint64(_time2), 
            pct1: uint8(_pct1), 
            pct2: uint8(_pct2), 
            discount: 0           
        });

        tokenAuction[_tokenId] = auction;
    }

    function _createNewTokenAuction(uint _tokenId, uint _basePrice) internal {
        _createCustomAuction(_tokenId, _basePrice, now, now + _auctionStartsAfter + _auctionDuration, 100, 10);
    }

    function _createStandartAuction(uint _tokenId, uint _basePrice) internal {
        uint start = now + _auctionStartsAfter;
        _createCustomAuction(_tokenId, _basePrice, start, start + _auctionDuration, 200, 110);
    }

    function _currentPrice(Auction _auction) internal view returns (uint) {
        if (_auction.discount > 0) {
            return uint((_auction.basePrice * (100 - _auction.discount)) / 100);
        }

        uint _startingPrice = uint((_auction.basePrice * _auction.pct1) / 100);

        if (_auction.time1 > now) {
            return _startingPrice;
        }

        uint _secondsPassed = uint(now - _auction.time1);
        uint _duration = uint(_auction.time2 - _auction.time1);
        uint _endingPrice = uint((_auction.basePrice * _auction.pct2) / 100);

        if (_secondsPassed >= _duration) {
            return _endingPrice;
        } else {
            int totalPriceChange = int(_endingPrice) - int(_startingPrice);
            int currentPriceChange = totalPriceChange * int(_secondsPassed) / int(_duration);
            int currentPrice = int(_startingPrice) + currentPriceChange;

            return uint(currentPrice);
        }
    }

    function _computePrice(uint _secondsPassed, uint _duration, uint _startingPrice, uint _endingPrice) private pure returns (uint) {
        if (_secondsPassed >= _duration) {
            return _endingPrice;
        } else {
            int totalPriceChange = int(_endingPrice) - int(_startingPrice);
            int currentPriceChange = totalPriceChange * int(_secondsPassed) / int(_duration);
            int currentPrice = int(_startingPrice) + currentPriceChange;

            return uint(currentPrice);
        }
    }
}

contract BitGallery is BitAuction {
    using SafeMath for uint;

    string public infoMessage;

    event TokenSold(uint tokenId, uint price, address from, address to);
    event NewToken(uint tokenId, string metadata);

    function BitGallery() public {
        setAuctionDurationRules(24 hours, 6 days, 3);

        setMessage("Our web site is www.bitgallery.co");                          
    }

    function() public payable {}

    function addArt(string _keyData, uint _basePrice) public onlyOwner {
        return addArtTo(address(this), _keyData, _basePrice);
    }

    function addArtTo(address _owner, string _keyData, uint _basePrice) public onlyOwner {
        require(_basePrice >= 1 finney);
        
        Art memory _art = Art({
            data: keccak256(_keyData)
        });

        uint tokenId = arts.push(_art) - 1;
        NewToken(tokenId, _keyData);
        _mint(_owner, tokenId);
        _createNewTokenAuction(tokenId, _basePrice);
    }

    function artExists(string _keydata) public view returns (bool) {
        for (uint i = 0; i < totalSupply(); i++) {
            if (arts[i].data == keccak256(_keydata)) {
                return true;
            }
        }

        return false;
    }

    function fullDataOf(uint _tokenId) public view returns (
        uint basePrice,
        uint64 time1,
        uint64 time2,
        uint8 pct1,
        uint8 pct2,
        uint8 discount,
        uint currentPrice,
        bool _canPurchase,
        address owner
    ) {
        Auction storage auction = tokenAuction[_tokenId];
        basePrice = auction.basePrice;
        time1 = auction.time1;
        time2 = auction.time2;
        pct1 = auction.pct1;
        pct2 = auction.pct2;
        discount = auction.discount;
        currentPrice = priceOf(_tokenId);
        _canPurchase = canPurchase(_tokenId);
        owner = ownerOf(_tokenId);
    }

    function payout(address _to) public onlyOwner {
        require(_to != address(this));
        
        if (_to == address(0)) { 
            _to = msg.sender;
        }

        _to.transfer(this.balance);
    }

    function purchase(uint _tokenId) public payable {
        Auction storage auction = tokenAuction[_tokenId];
        require(now >= auction.time1);
        uint price = _currentPrice(auction);
        require(msg.value >= price);

        uint payment = uint((price * (100 - _auctionFee)) / 100);
        uint purchaseExcess = msg.value - price;
        _createStandartAuction(_tokenId, price);

        address from = ownerOf(_tokenId);
        address to = msg.sender;
        _transfer(from, to, _tokenId);

        if (from != address(this)) {
            from.transfer(payment);
        }

        TokenSold(_tokenId, price, from, to);
        msg.sender.transfer(purchaseExcess);
    }

    function setMessage(string _message) public onlyOwner {        
        infoMessage = _message;
    }
}
