pragma solidity ^0.4.18;

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

contract ERC20Basic {
    uint256 public totalSupply;

    function balanceOf(address who) constant public returns (uint256);

    function transfer(address to, uint256 value) public returns (bool);

    event Transfer(address indexed from, address indexed to, uint256 value);
}

contract ERC20 is ERC20Basic {
    function allowance(address owner, address spender) constant public returns (uint256);

    function transferFrom(address from, address to, uint256 value) public returns (bool);

    function approve(address spender, uint256 value) public returns (bool);

    event Approval(address indexed owner, address indexed spender, uint256 value);
}

contract Owned {

    address public owner;

    address public newOwner;

    function Owned() public payable {
        owner = msg.sender;
    }

    modifier onlyOwner {
        require(owner == msg.sender);
        _;
    }

    function changeOwner(address _owner) onlyOwner public {
        require(_owner != 0);
        newOwner = _owner;
    }

    function confirmOwner() public {
        require(newOwner == msg.sender);
        owner = newOwner;
        delete newOwner;
    }
}

contract Blocked {

    uint public blockedUntil;

    modifier unblocked {
        require(now > blockedUntil);
        _;
    }
}

contract BalancingToken is ERC20 {
    mapping (address => uint256) public balances;      //!< array of all balances

    function balanceOf(address _owner) public constant returns (uint256 balance) {
        return balances[_owner];
    }
}

contract DividendToken is BalancingToken, Blocked, Owned {

    using SafeMath for uint256;

    event DividendReceived(address indexed dividendReceiver, uint256 dividendValue);

    mapping (address => mapping (address => uint256)) public allowed;

    uint public totalReward;
    uint public lastDivideRewardTime;

    // Fix for the ERC20 short address attack
    modifier onlyPayloadSize(uint size) {
        require(msg.data.length >= size + 4);
        _;
    }

    // Fix for the ERC20 short address attack
    modifier rewardTimePast() {
        require(now > lastDivideRewardTime + rewardDays);
        _;
    }

    struct TokenHolder {
        uint256 balance;
        uint    balanceUpdateTime;
        uint    rewardWithdrawTime;
    }

    mapping(address => TokenHolder) holders;

    uint public rewardDays = 0;

    function transfer(address _to, uint256 _value) onlyPayloadSize(2 * 32) unblocked public returns (bool) {
        return transferSimple(_to, _value);
    }

    function transferSimple(address _to, uint256 _value) internal returns (bool) {
        beforeBalanceChanges(msg.sender);
        beforeBalanceChanges(_to);
        balances[msg.sender] = balances[msg.sender].sub(_value);
        balances[_to] = balances[_to].add(_value);
        Transfer(msg.sender, _to, _value);
        return true;
    }

    function transferFrom(address _from, address _to, uint256 _value) onlyPayloadSize(3 * 32) unblocked public returns (bool) {
        beforeBalanceChanges(_from);
        beforeBalanceChanges(_to);
        var _allowance = allowed[_from][msg.sender];
        balances[_to] = balances[_to].add(_value);
        balances[_from] = balances[_from].sub(_value);
        allowed[_from][msg.sender] = _allowance.sub(_value);
        Transfer(_from, _to, _value);
        return true;
    }

    function approve(address _spender, uint256 _value) onlyPayloadSize(2 * 32) unblocked public returns (bool) {
        require((_value == 0) || (allowed[msg.sender][_spender] == 0));
        allowed[msg.sender][_spender] = _value;
        Approval(msg.sender, _spender, _value);
        return true;
    }

    function allowance(address _owner, address _spender) onlyPayloadSize(2 * 32) unblocked constant public returns (uint256 remaining) {
        return allowed[_owner][_spender];
    }

    function reward() constant public returns (uint256) {
        if (holders[msg.sender].rewardWithdrawTime >= lastDivideRewardTime) {
            return 0;
        }
        uint256 balance;
        if (holders[msg.sender].balanceUpdateTime <= lastDivideRewardTime) {
            balance = balances[msg.sender];
        } else {
            balance = holders[msg.sender].balance;
        }
        return totalReward.mul(balance).div(totalSupply);
    }

    function withdrawReward() public returns (uint256) {
        uint256 rewardValue = reward();
        if (rewardValue == 0) {
            return 0;
        }
        if (balances[msg.sender] == 0) {
            // garbage collector
            delete holders[msg.sender];
        } else {
            holders[msg.sender].rewardWithdrawTime = now;
        }
        require(msg.sender.call.gas(3000000).value(rewardValue)());
        DividendReceived(msg.sender, rewardValue);
        return rewardValue;
    }

    // Divide up reward and make it accesible for withdraw
    function divideUpReward(uint inDays) rewardTimePast onlyOwner external payable {
        require(inDays >= 15 && inDays <= 45);
        lastDivideRewardTime = now;
        rewardDays = inDays;
        totalReward = this.balance;
    }

    function withdrawLeft() rewardTimePast onlyOwner external {
        require(msg.sender.call.gas(3000000).value(this.balance)());
    }

    function beforeBalanceChanges(address _who) public {
        if (holders[_who].balanceUpdateTime <= lastDivideRewardTime) {
            holders[_who].balanceUpdateTime = now;
            holders[_who].balance = balances[_who];
        }
    }
}

contract RENTCoin is DividendToken {

    string public constant name = "RentAway Coin";

    string public constant symbol = "RTW";

    uint32 public constant decimals = 18;

    function RENTCoin(uint256 initialSupply, uint unblockTime) public {
        totalSupply = initialSupply;
        balances[owner] = initialSupply;
        blockedUntil = unblockTime;
    }

    function manualTransfer(address _to, uint256 _value) onlyPayloadSize(2 * 32) onlyOwner public returns (bool) {
        return transferSimple(_to, _value);
    }
}

contract TimingCrowdsale {

    // Date of start pre-ICO and ICO.
    uint public constant preICOstartTime = 1517461200; // start at Thursday, February 1, 2018 5:00:00 AM
    uint public constant ICOstartTime =    1518670800; // start at Thursday, February 15, 2018 5:00:00 AM
    uint public constant ICOendTime =      1521090000; // end at Thursday, March 15, 2018 5:00:00 AM

    function currentTime() internal view returns (uint) {
        return now;
    }

    function isPreICO() public view returns (bool) {
        var curTime = currentTime();
        return curTime < ICOstartTime && curTime >= preICOstartTime;
    }

    function isICO() public view returns (bool) {
        var curTime = currentTime();
        return curTime < ICOendTime && curTime >= ICOstartTime;
    }

    function isPreICOFinished() public view returns (bool) {
        return currentTime() > ICOstartTime;
    }

    function isICOFinished() public view returns (bool) {
        return currentTime() > ICOendTime;
    }
}

contract BonusCrowdsale is TimingCrowdsale {

    function getBonus(uint256 amount) public view returns (uint) {
        uint bonus = getAmountBonus(amount);
        if (isPreICO()) {
            bonus += 25;
        }
        return bonus;
    }

    function getAmountBonus(uint256 amount) public view returns (uint) {
        if (amount >= 25 ether) {
            return 15;
        }
        if (amount >= 10 ether) {
            return 5;
        }
        return 0;
    }
}

contract ManualSendingCrowdsale is BonusCrowdsale, Owned {
    using SafeMath for uint256;

    struct AmountData {
        bool exists;
        uint256 value;
    }

    mapping (uint => AmountData) public amountsByCurrency;

    function addCurrency(uint currency) external onlyOwner {
        addCurrencyInternal(currency);
    }

    function addCurrencyInternal(uint currency) internal {
        AmountData storage amountData = amountsByCurrency[currency];
        amountData.exists = true;
    }

    function manualTransferTokensToInternal(address to, uint256 givenTokens, uint currency, uint256 amount) internal returns (uint256) {
        AmountData memory tempAmountData = amountsByCurrency[currency];
        require(tempAmountData.exists);
        AmountData storage amountData = amountsByCurrency[currency];
        amountData.value = amountData.value.add(amount);
        return transferTokensTo(to, givenTokens);
    }

    function transferTokensTo(address to, uint256 givenTokens) internal returns (uint256);
}

contract WithdrawCrowdsale is ManualSendingCrowdsale {

    function isWithdrawAllowed() public view returns (bool);

    modifier canWithdraw() {
        require(isWithdrawAllowed());
        _;
    }

    function withdraw() external onlyOwner canWithdraw {
        require(msg.sender.call.gas(3000000).value(this.balance)());
    }

    function withdrawAmount(uint256 amount) external onlyOwner canWithdraw {
        uint256 givenAmount = amount;
        if (this.balance < amount) {
            givenAmount = this.balance;
        }
        require(msg.sender.call.gas(3000000).value(givenAmount)());
    }
}

contract RefundableCrowdsale is WithdrawCrowdsale {

    event Refunded(address indexed beneficiary, uint256 weiAmount);

    mapping (address => uint256) public deposited;

    function refund(address investor) external {
        require(isRefundAllowed());
        uint256 depositedValue = deposited[investor];
        deposited[investor] = 0;
        require(investor.call.gas(3000000).value(depositedValue)());
        Refunded(investor, depositedValue);
    }

    function isRefundAllowed() internal view returns (bool);
}

contract Crowdsale is RefundableCrowdsale {

    using SafeMath for uint256;

    enum State { ICO, REFUND, DONE }
    State public state = State.ICO;

    uint256 public constant maxTokenAmount = 75000000 * 10**18; // max minting
    uint256 public constant bountyTokens =   15000000 * 10**18; // bounty amount
    uint256 public constant softCapTokens =  6000000 * 10**18; // soft cap

    uint public constant unblockTokenTime = 1519880400; // end at Thursday, March 1, 2018 5:00:00 AM

    RENTCoin public token;

    uint256 public leftTokens = 0;

    uint256 public totalAmount = 0;
    uint public transactionCounter = 0;

    bool public bonusesPayed = false;

    uint256 public constant rateToEther = 1000; // rate to ether, how much tokens gives to 1 ether

    uint256 public constant minAmountForDeal = 10**16; // 0.01 ETH

    uint256 public soldTokens = 0;

    modifier canBuy() {
        require(!isFinished());
        require(isPreICO() || isICO());
        _;
    }

    modifier minPayment() {
        require(msg.value >= minAmountForDeal);
        _;
    }

    function Crowdsale() public {
        token = new RENTCoin(maxTokenAmount, unblockTokenTime);
        leftTokens = maxTokenAmount - bountyTokens;
        addCurrencyInternal(0); // add BTC
    }

    function isFinished() public view returns (bool) {
        return isICOFinished() || (leftTokens == 0 && (state == State.ICO || state == State.DONE));
    }

    function isWithdrawAllowed() public view returns (bool) {
        return soldTokens >= softCapTokens;
    }

    function isRefundAllowed() internal view returns (bool) {
        return state == State.REFUND;
    }

    function() external canBuy minPayment payable {
        address investor = msg.sender;
        uint256 amount = msg.value;
        uint bonus = getBonus(amount);
        uint256 givenTokens = amount.mul(rateToEther).div(100).mul(100 + bonus);
        uint256 providedTokens = transferTokensTo(investor, givenTokens);

        if (givenTokens > providedTokens) {
            uint256 needAmount = providedTokens.mul(100).div(100 + bonus).div(rateToEther);
            require(amount > needAmount);
            require(investor.call.gas(3000000).value(amount - needAmount)());
            amount = needAmount;
        }
        totalAmount = totalAmount.add(amount);
        if (!isWithdrawAllowed()) {
            deposited[investor] = deposited[investor].add(msg.value);
        }
    }

    function manualTransferTokensTo(address to, uint256 givenTokens, uint currency, uint256 amount) external onlyOwner canBuy returns (uint256) {
        return manualTransferTokensToInternal(to, givenTokens, currency, amount);
    }

    function finishCrowdsale() external {
        require(isFinished());
        require(state == State.ICO);
        if (!isWithdrawAllowed())  {
            state = State.REFUND;
            bonusesPayed = true;
        } else {
            state = State.DONE;
        }
    }

    function takeBounty() external onlyOwner {
        require(state == State.DONE);
        require(now > ICOendTime);
        require(!bonusesPayed);
        token.changeOwner(msg.sender);
        bonusesPayed = true;
        require(token.transfer(msg.sender, token.balanceOf(this)));
    }

    function addToSoldTokens(uint256 providedTokens) internal {
        soldTokens = soldTokens.add(providedTokens);
    }

    function transferTokensTo(address to, uint256 givenTokens) internal returns (uint256) {
        var providedTokens = givenTokens;
        if (givenTokens > leftTokens) {
            providedTokens = leftTokens;
        }
        leftTokens = leftTokens.sub(providedTokens);
        addToSoldTokens(providedTokens);
        require(token.manualTransfer(to, providedTokens));
        transactionCounter = transactionCounter + 1;
        return providedTokens;
    }
}