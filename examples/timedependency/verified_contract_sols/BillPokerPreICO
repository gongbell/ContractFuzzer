/**
 *Submitted for verification at Etherscan.io on 2017-11-30
*/

pragma solidity ^0.4.16;


contract Token {
    uint256 public totalSupply;

    function balanceOf(address who) constant returns (uint256);

    function transferFrom(address _from, address _to, uint256 _value) returns (bool);

    event Transfer(address indexed from, address indexed to, uint256 value);
}


/**
 * @title SafeMath
 * @dev Math operations with safety checks that throw on error
 */
library SafeMath {
    function mul(uint256 a, uint256 b) internal constant returns (uint256) {
        uint256 c = a * b;
        assert(a == 0 || c / a == b);
        return c;
    }

    function div(uint256 a, uint256 b) internal constant returns (uint256) {
        // assert(b > 0); // Solidity automatically throws when dividing by 0
        uint256 c = a / b;
        // assert(a == b * c + a % b); // There is no case in which this doesn't hold
        return c;
    }

    function sub(uint256 a, uint256 b) internal constant returns (uint256) {
        assert(b <= a);
        return a - b;
    }

    function add(uint256 a, uint256 b) internal constant returns (uint256) {
        uint256 c = a + b;
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


    /**
     * @dev The Ownable constructor sets the original `owner` of the contract to the sender
     * account.
     */
    function Ownable() {
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
    function transferOwnership(address newOwner) onlyOwner {
        require(newOwner != address(0));
        owner = newOwner;
    }

}


/**
 * @title Pausable
 * @dev Base contract which allows children to implement an emergency stop mechanism.
 */
contract Pausable is Ownable {
    event Pause();

    event Unpause();

    bool public paused = false;


    /**
     * @dev modifier to allow actions only when the contract IS paused
     */
    modifier whenNotPaused() {
        require(!paused);
        _;
    }

    /**
     * @dev modifier to allow actions only when the contract IS NOT paused
     */
    modifier whenPaused() {
        require(paused);
        _;
    }

    /**
     * @dev called by the owner to pause, triggers stopped state
     */
    function pause() onlyOwner whenNotPaused {
        paused = true;
        Pause();
    }

    /**
     * @dev called by the owner to unpause, returns to normal state
     */
    function unpause() onlyOwner whenPaused {
        paused = false;
        Unpause();
    }
}


contract BillPokerPreICO is Ownable, Pausable {
    using SafeMath for uint;

    /* The party who holds the full token pool and has approve()'ed tokens for this crowdsale */
    address public tokenWallet = 0xf91E6d611ec35B985bADAD2F0DA96820930B9BD2;

    uint public tokensSold;

    uint public weiRaised;

    mapping (address => uint256) public holdTokens;

    mapping (address => uint256) public purchaseTokens;

    address[] public holdTokenInvestors;

    Token public token = Token(0xc305fcdc300fa43c527e9327711f360e79528a70);

    uint public constant minInvest = 0.0001 ether;

    uint public constant tokensLimit = 25000000 ether;

    // start and end timestamps where investments are allowed
    uint256 public startTime = 1510339500; // 14 November 2017 00:00 UTC

    uint256 public endTime = 1519689600; // 28 December 2017 00:00 UTC

    uint public price = 0.0001 ether;

    bool public isHoldTokens = false;

    uint public investorCount;

    mapping (bytes32 => Promo) public promoMap;

    struct Promo {
    bool enable;
    uint investorPercentToken;
    address dealer;
    uint dealerPercentToken;
    uint dealerPercentETH;
    uint buyCount;
    uint investorTokenAmount;
    uint dealerTokenAmount;
    uint investorEthAmount;
    uint dealerEthAmount;
    }
    
    function addPromo(bytes32 promoPublicKey, uint userPercentToken, address dealer, uint dealerPercentToken, uint dealerPercentETH) public onlyOwner {
        promoMap[promoPublicKey] = Promo(true, userPercentToken, dealer, dealerPercentToken, dealerPercentETH, 0, 0, 0, 0, 0);
    }

    function removePromo(bytes32 promoPublicKey) public onlyOwner {
        promoMap[promoPublicKey].enable = false;
    }


    /**
     * event for token purchase logging
     * @param purchaser who paid for the tokens
     * @param beneficiary who got the tokens
     * @param value weis paid for purchase
     * @param amount amount of tokens purchased
     */
    event TokenPurchase(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount);

    // fallback function can be used to buy tokens
    function() public payable {
        buyTokens(msg.sender);
    }

    // low level token purchase function
    function buyTokens(address beneficiary) public whenNotPaused payable {
        require(startTime <= now && now <= endTime);

        uint weiAmount = msg.value;

        require(weiAmount >= minInvest);

        uint tokenAmountEnable = tokensLimit.sub(tokensSold);

        require(tokenAmountEnable > 0);

        uint tokenAmount = weiAmount / price * 1 ether;

        if (tokenAmount > tokenAmountEnable) {
            tokenAmount = tokenAmountEnable;
            weiAmount = tokenAmount * price / 1 ether;
            msg.sender.transfer(msg.value.sub(weiAmount));


            if (msg.data.length > 0) {
                Promo storage promo = promoMap[sha3(msg.data)];
                if (promo.enable && promo.dealerPercentETH > 0) {
                    uint dealerEthAmount = weiAmount * promo.dealerPercentETH / 10000;
                    promo.dealer.transfer(dealerEthAmount);
                    weiAmount = weiAmount.sub(dealerEthAmount);

                    promo.dealerEthAmount += dealerEthAmount;
                }
            }
        }
        else {
            uint countBonusAmount = tokenAmount * getCountBonus(weiAmount) / 1000;
            uint timeBonusAmount = tokenAmount * getTimeBonus(now) / 1000;

            if (msg.data.length > 0) {
                bytes32 promoPublicKey = sha3(msg.data);
                promo = promoMap[promoPublicKey];
                if (promo.enable) {
                    
                    promo.buyCount++;
                    promo.investorTokenAmount += tokenAmount;
                    promo.investorEthAmount += weiAmount;
                    
                    if (promo.dealerPercentToken > 0) {
                        uint dealerTokenAmount = tokenAmount * promo.dealerPercentToken / 10000;
                        sendTokens(promo.dealer, dealerTokenAmount);
                        promo.dealerTokenAmount += dealerTokenAmount;
                    }

                    if (promo.dealerPercentETH > 0) {
                        dealerEthAmount = weiAmount * promo.dealerPercentETH / 10000;
                        promo.dealer.transfer(dealerEthAmount);
                        weiAmount = weiAmount.sub(dealerEthAmount);
                        promo.dealerEthAmount += dealerEthAmount;
                    }

                        
                    if (promo.investorPercentToken > 0) {
                        uint promoBonusAmount = tokenAmount * promo.investorPercentToken / 10000;
                        tokenAmount += promoBonusAmount;
                    }

                }
            }

            tokenAmount += countBonusAmount + timeBonusAmount;

            if (tokenAmount > tokenAmountEnable) {
                tokenAmount = tokenAmountEnable;
            }
        }


        if (purchaseTokens[beneficiary] == 0) investorCount++;

        purchaseTokens[beneficiary] = purchaseTokens[beneficiary].add(tokenAmount);

        sendTokens(beneficiary, tokenAmount);

        weiRaised = weiRaised.add(weiAmount);

        TokenPurchase(msg.sender, beneficiary, weiAmount, tokenAmount);
    }

    function sendTokens(address to, uint tokenAmount) private {
        if (isHoldTokens) {
            if (holdTokens[to] == 0) holdTokenInvestors.push(to);
            holdTokens[to] = holdTokens[to].add(tokenAmount);
        }
        else {
            require(token.transferFrom(tokenWallet, to, tokenAmount));
        }

        tokensSold = tokensSold.add(tokenAmount);
    }

    uint[] etherForCountBonus = [2 ether, 3 ether, 5 ether, 7 ether, 9 ether, 12 ether, 15 ether, 20 ether, 25 ether, 30 ether, 35 ether, 40 ether, 45 ether, 50 ether, 60 ether, 70 ether, 80 ether, 90 ether, 100 ether, 120 ether, 150 ether, 200 ether, 250 ether, 300 ether, 350 ether, 400 ether, 450 ether, 500 ether];

    uint[] amountForCountBonus = [0, 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 90, 100, 105, 110, 115, 120, 125, 130, 135, 140, 145, 150];


    function getCountBonus(uint weiAmount) public constant returns (uint) {
        for (uint i = 0; i < etherForCountBonus.length; i++) {
            if (weiAmount < etherForCountBonus[i]) return amountForCountBonus[i];
        }
        return amountForCountBonus[amountForCountBonus.length - 1];
    }

    function getTimeBonus(uint time) public constant returns (uint) {
        if (time < startTime + 604800) return 250;
        if (time < startTime + 604800) return 200;
        if (time < startTime + 259200) return 100;
        return 0;
    }

    function withdrawal(address to) public onlyOwner {
        to.transfer(this.balance);
    }

    function holdTokenInvestorsCount() public constant returns(uint){
        return holdTokenInvestors.length;
    }

    uint public sendInvestorIndex = 0;

    function finalSendTokens() public onlyOwner {
        isHoldTokens = false;
        
        for (uint i = sendInvestorIndex; i < holdTokenInvestors.length; i++) {
            address investor = holdTokenInvestors[i];
            uint tokenAmount = holdTokens[investor];

            if (tokenAmount > 0) {
                holdTokens[investor] = 0;
                require(token.transferFrom(tokenWallet, investor, tokenAmount));
            }

            if (msg.gas < 100000) {
                sendInvestorIndex = i;
                return;
            }
        }

        sendInvestorIndex = holdTokenInvestors.length;
    }

}
