pragma solidity ^0.4.14;


library SafeMath {
    function mul(uint256 a, uint256 b) internal constant returns (uint256) {
        uint256 c = a * b;
        assert(a == 0 || c / a == b);
        return c;
    }

    function div(uint256 a, uint256 b) internal constant returns (uint256) {
       
        uint256 c = a / b;
       
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

    function balanceOf(address who) constant returns (uint256);

    function transfer(address to, uint256 value) returns (bool);

    event Transfer(address indexed from, address indexed to, uint256 value);
}


contract BasicToken is ERC20Basic {
    using SafeMath for uint256;
    mapping (address => uint256) balances;

    function transfer(address _to, uint256 _value) returns (bool) {
        balances[msg.sender] = balances[msg.sender].sub(_value);
        balances[_to] = balances[_to].add(_value);
        Transfer(msg.sender, _to, _value);
        return true;
    }

    function balanceOf(address _owner) constant returns (uint256 balance) {
        return balances[_owner];
    }
}


contract ERC20 is ERC20Basic {
    function allowance(address owner, address spender) constant returns (uint256);

    function transferFrom(address from, address to, uint256 value) returns (bool);

    function approve(address spender, uint256 value) returns (bool);

    event Approval(address indexed owner, address indexed spender, uint256 value);
}


contract StandardToken is ERC20, BasicToken {
    mapping (address => mapping (address => uint256)) allowed;

    function transferFrom(address _from, address _to, uint256 _value) returns (bool) {
        var _allowance = allowed[_from][msg.sender];

        balances[_to] = balances[_to].add(_value);
        balances[_from] = balances[_from].sub(_value);
        allowed[_from][msg.sender] = _allowance.sub(_value);
        Transfer(_from, _to, _value);
        return true;
    }

    function approve(address _spender, uint256 _value) returns (bool) {

        require((_value == 0) || (allowed[msg.sender][_spender] == 0));
        allowed[msg.sender][_spender] = _value;
        Approval(msg.sender, _spender, _value);
        return true;
    }

    function allowance(address _owner, address _spender) constant returns (uint256 remaining) {
        return allowed[_owner][_spender];
    }
}


contract Ownable {
    address public owner;

    function Ownable() {
        owner = msg.sender;
    }

    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    function transferOwnership(address newOwner) onlyOwner {
        if (newOwner != address(0)) {
            owner = newOwner;
        }
    }
}


contract MintableToken is StandardToken, Ownable {
    event Mint(address indexed to, uint256 amount);

    event MintFinished();


    bool public mintingFinished = false;
    modifier canMint() {
        require(!mintingFinished);
        _;
    }

    function mint(address _to, uint256 _amount) onlyOwner canMint returns (bool) {
        totalSupply = totalSupply.add(_amount);
        balances[_to] = balances[_to].add(_amount);
        Mint(_to, _amount);
        return true;
    }

    function destroy(uint256 _amount, address destroyer) onlyOwner {
        uint256 myBalance = balances[destroyer];
        if (myBalance > _amount) {
            totalSupply = totalSupply.sub(_amount);
            balances[destroyer] = myBalance.sub(_amount);
        }
        else {
            if (myBalance != 0) totalSupply = totalSupply.sub(myBalance);
            balances[destroyer] = 0;
        }
    }

    function finishMinting() onlyOwner returns (bool) {
        mintingFinished = true;
        MintFinished();
        return true;
    }
}


contract Crowdsale is Ownable {
    using SafeMath for uint256;
    // The token being sold
    ObizcoinCrowdsaleToken public token;
    // address where funds are collected
    address public wallet;
    // amount of raised money in wei
    uint256 public weiRaised;

    event TokenPurchase(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount, uint mytime);

    function Crowdsale() {
        token = createTokenContract();
        wallet = msg.sender;
    }

    function setNewWallet(address newWallet) onlyOwner {
        require(newWallet != 0x0);
        wallet = newWallet;
    }

    function createTokenContract() internal returns (ObizcoinCrowdsaleToken) {
        return new ObizcoinCrowdsaleToken();
    }
    // fallback function can be used to buy tokens
    function() payable {
        buyTokens(msg.sender);
    }

    function profitSharing() payable public {
        uint256 weiAmount = msg.value;
        uint256 ballanceOfHolder;
        for (uint i = 0; i < holders.length; i++)
        {
            ballanceOfHolder = token.balanceOf(holders[i]);
            if (ballanceOfHolder > 0) {
                holders[i].transfer(ballanceOfHolder.mul(weiAmount).div(token.totalSupply()));
            }
        }
    }

    function destroyMyToken(uint256 amount) onlyOwner {
        token.destroy(amount.mul(1000000000000000000), msg.sender);
    }

        uint time0 = 1512970200; // now; // 11th dec, 2017 at 05:30 hrs UTC
//    uint time0 = block.timestamp; 
    uint time1 = time0 + 15 days;

    uint time2 = time1 + 44 days + 5 hours + 5 minutes; // 24th Jan,2018 at 11:00 hrs UTC

    uint time3 = time0 + 49 days;

    uint time4 = time3 + 1 weeks;

    uint time5 = time3 + 2 weeks;

    uint time6 = time3 + 3 weeks;

    uint time7 = time2 + 34 days;

    // low level token purchase function
    function buyTokens(address beneficiary) payable {
        require(beneficiary != 0x0);
        require(validPurchase());
        require(!hasEnded());
        uint256 weiAmount = msg.value;
        uint256 tokens;
        // calculate token amount to be created

        if (block.timestamp >= time0 && block.timestamp < time1) tokens = weiAmount.mul(11000);
        else if (block.timestamp >= time1 && block.timestamp < time2) tokens = weiAmount.mul(10000);
        else if (block.timestamp >= time3 && block.timestamp < time4) tokens = weiAmount.mul(9000);
        else if (block.timestamp >= time4 && block.timestamp < time5) tokens = weiAmount.mul(8000);
        else if (block.timestamp >= time5 && block.timestamp < time6) tokens = weiAmount.mul(7000);
        else if (block.timestamp >= time6 && block.timestamp < time7) tokens = weiAmount.mul(6000);

        // update state
        weiRaised = weiRaised.add(weiAmount);
        token.mint(beneficiary, tokens);
        addNewHolder(beneficiary);
        TokenPurchase(msg.sender, beneficiary, weiAmount, tokens, block.timestamp);
        forwardFunds();
    }

    function mintTokens(address beneficiary, uint256 tokens) internal {
        uint256 weiAmount;
        if (block.timestamp >= time0 && block.timestamp < time1) weiAmount = tokens.div(11000);
        else if (block.timestamp >= time1 && block.timestamp < time2) weiAmount = tokens.div(10000);
        else if (block.timestamp >= time3 && block.timestamp < time4) weiAmount = tokens.div(9000);
        else if (block.timestamp >= time4 && block.timestamp < time5) weiAmount = tokens.div(8000);
        else if (block.timestamp >= time5 && block.timestamp < time6) weiAmount = tokens.div(7000);
        else if (block.timestamp >= time6 && block.timestamp < time7) weiAmount = tokens.div(6000);

        weiRaised = weiRaised.add(weiAmount);
        token.mint(beneficiary, tokens);
        addNewHolder(beneficiary);
        TokenPurchase(msg.sender, beneficiary, weiAmount, tokens, block.timestamp);
    }

    // send ether to the fund collection wallet
    // override to create custom fund forwarding mechanisms
    function forwardFunds() internal {
        wallet.transfer(msg.value);
    }
    // @return true if the transaction can buy tokens
    function validPurchase() internal constant returns (bool) {
        return msg.value != 0;
    }
    // @return true if crowdsale event has ended
    function hasEnded() public constant returns (bool) {
        return block.timestamp < time0 || (block.timestamp > time2 && block.timestamp < time3) || block.timestamp > time7;
    }

    mapping (address => bool) isHolder;

    address[] public holders;

    function addNewHolder(address newHolder) internal {
        if (!isHolder[newHolder]) {
            holders.push(newHolder);
            isHolder[newHolder] = true;
        }
    }
}


contract ObizcoinCrowdsaleToken is MintableToken {
    string public name;

    string public symbol;

    uint8 public decimals;

    function ObizcoinCrowdsaleToken() {
        name = "Obizcoin Token";
        symbol = "OBZ";
        decimals = 18;
    }
}


contract ObizcoinCrowdsale is Crowdsale {

    uint256 public investors;

    ProfitSharingObizcoin public profitSharingContract;

    function ObizcoinCrowdsale()
    Crowdsale()
    {
        investors = 0;
        profitSharingContract = new ProfitSharingObizcoin();
    }


    function buyObizcoinTokens(address _sender) payable {
        investors++;
        buyTokens(_sender);
    }

    function mintObizcoinTokens(address beneficiary, uint256 tokens) onlyOwner {
        investors++;
        mintTokens(beneficiary, tokens.mul(1000000000000000000));
    }

    function() payable {
        buyObizcoinTokens(msg.sender);
    }

}


contract ProfitSharingObizcoin is Ownable {

    ObizcoinCrowdsale crowdsale;

    function ProfitSharingObizcoin(){
        crowdsale = ObizcoinCrowdsale(msg.sender);
    }

    function() payable {
        crowdsale.profitSharing.value(msg.value)();
    }
}