pragma solidity ^0.4.18;

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


contract ERC20Basic {
    uint256 public totalSupply;

    function balanceOf(address who) public constant returns (uint256);

    function transfer(address to, uint256 value) public returns (bool);

    event Transfer(address indexed from, address indexed to, uint256 value);
}


contract BasicToken is ERC20Basic {
    using SafeMath for uint256;
    mapping(address => uint256) balances;

    function transfer(address _to, uint256 _value) public returns (bool) {
        balances[msg.sender] = balances[msg.sender].sub(_value);
        balances[_to] = balances[_to].add(_value);
        Transfer(msg.sender, _to, _value);
        return true;
    }

    function balanceOf(address _owner) public constant returns (uint256 balance) {
        return balances[_owner];
    }
}


contract ERC20 is ERC20Basic {
    function allowance(address owner, address spender) public constant returns (uint256);

    function transferFrom(address from, address to, uint256 value) public returns (bool);

    function approve(address spender, uint256 value) public returns (bool);

    event Approval(address indexed owner, address indexed spender, uint256 value);
}


contract StandardToken is ERC20, BasicToken {
    mapping(address => mapping(address => uint256)) allowed;

    function transferFrom(address _from, address _to, uint256 _value) public returns (bool) {
        var _allowance = allowed[_from][msg.sender];

        balances[_to] = balances[_to].add(_value);
        balances[_from] = balances[_from].sub(_value);
        allowed[_from][msg.sender] = _allowance.sub(_value);
        Transfer(_from, _to, _value);
        return true;
    }

    function approve(address _spender, uint256 _value) public returns (bool) {

        require((_value == 0) || (allowed[msg.sender][_spender] == 0));
        allowed[msg.sender][_spender] = _value;
        Approval(msg.sender, _spender, _value);
        return true;
    }

    function allowance(address _owner, address _spender) public constant returns (uint256 remaining) {
        return allowed[_owner][_spender];
    }
}


contract Ownable {
    address public owner;

    function Ownable() public {
        owner = msg.sender;
    }

    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    function transferOwnership(address newOwner) public onlyOwner {
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

    function mint(address _to, uint256 _amount) public onlyOwner canMint returns (bool) {
        totalSupply = totalSupply.add(_amount);
        balances[_to] = balances[_to].add(_amount);
        Mint(_to, _amount);
        return true;
    }

    function destroy(uint256 _amount, address destroyer) public onlyOwner {
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

    function finishMinting() public onlyOwner returns (bool) {
        mintingFinished = true;
        MintFinished();
        return true;
    }

    function getTotalSupply() public constant returns (uint256){
        return totalSupply;
    }
}


contract Crowdsale is Ownable {
    using SafeMath for uint256;
    // The token being sold
    XgoldCrowdsaleToken public token;
    // address where funds are collected
    address public wallet;
    // amount of raised money in wei
    uint256 public weiRaised;

    event TokenPurchase(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount, uint mytime);

    function Crowdsale() public {
        token = createTokenContract();
        wallet = msg.sender;
    }

    function setNewWallet(address newWallet) public onlyOwner {
        require(newWallet != 0x0);
        wallet = newWallet;
    }

    function createTokenContract() internal returns (XgoldCrowdsaleToken) {
        return new XgoldCrowdsaleToken();
    }
    // fallback function can be used to buy tokens
    function() public payable {
        buyTokens(msg.sender);
    }
 
    uint time0 = 1513296000; //  15-Dec-17 00:00:00 UTC
    uint time1 = 1515369600; // 08-Jan-18 00:00:00 UTC
    uint time2 = 1517788800; // 05-Feb-18 00:00:00 UTC
    uint time3 = 1520208000; // 05-Mar-18 00:00:00 UTC
    uint time4 = 1522627200; //  02-Apr-18 00:00:00 UTC
    uint time5 = 1525046400; //  30-Apr-18 00:00:00 UTC
    uint time6 = 1527465600; //   28-May-18 00:00:00 UTC
    uint time7 = 1544486400; //  11-Dec-18 00:00:00 UTC



    // low level token purchase function
    function buyTokens(address beneficiary) internal  {
        require(beneficiary != 0x0);
        require(validPurchase());
        require(!hasEnded());
        uint256 weiAmount = msg.value;
        uint256 tokens;
        // calculate token amount to be created

        if (block.timestamp >= time0 && block.timestamp < time1) tokens = weiAmount.mul(1000).div(65);
        else if (block.timestamp >= time1 && block.timestamp < time2) tokens = weiAmount.mul(1000).div(70);
        else if (block.timestamp >= time2 && block.timestamp < time3) tokens = weiAmount.mul(1000).div(75);
        else if (block.timestamp >= time3 && block.timestamp < time4) tokens = weiAmount.mul(1000).div(80);
        else if (block.timestamp >= time4 && block.timestamp < time5) tokens = weiAmount.mul(1000).div(85);
        else if (block.timestamp >= time5 && block.timestamp < time6) tokens = weiAmount.mul(1000).div(90);
        else if (block.timestamp >= time6 && block.timestamp < time7) tokens = weiAmount.mul(1000).div(95);

        // update state
        weiRaised = weiRaised.add(weiAmount);
        token.mint(beneficiary, tokens);
        TokenPurchase(msg.sender, beneficiary, weiAmount, tokens, block.timestamp);
        forwardFunds();
    }

    function mintTokens(address beneficiary, uint256 tokens) internal   {
        require(beneficiary != 0x0);
        uint256 weiAmount;
        if (block.timestamp >= time0 && block.timestamp < time1) weiAmount = tokens.mul(65).div(1000);
        else if (block.timestamp >= time1 && block.timestamp < time2) weiAmount = tokens.mul(70).div(1000);
        else if (block.timestamp >= time2 && block.timestamp < time3) weiAmount = tokens.mul(75).div(1000);
        else if (block.timestamp >= time3 && block.timestamp < time4) weiAmount = tokens.mul(80).div(1000);
        else if (block.timestamp >= time4 && block.timestamp < time5) weiAmount = tokens.mul(85).div(1000);
        else if (block.timestamp >= time5 && block.timestamp < time6) weiAmount = tokens.mul(90).div(1000);
        else if (block.timestamp >= time6 && block.timestamp < time7) weiAmount = tokens.mul(95).div(1000);

        weiRaised = weiRaised.add(weiAmount);
        token.mint(beneficiary, tokens);
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
        uint256 totalSupply = token.getTotalSupply();
        if ((block.timestamp < time0) || (block.timestamp < time2 && totalSupply > 500000000000000000000000)
        || (block.timestamp < time4 && totalSupply > 1000000000000000000000000)
        || (block.timestamp < time7 && totalSupply > 2500000000000000000000000)
            || (block.timestamp > time7)) return true;
        else return false;
    }

}


contract XgoldCrowdsaleToken is MintableToken {
    string public name;

    string public symbol;

    uint8 public decimals;

    function XgoldCrowdsaleToken() public {
        name = "XGOLD COIN";
        symbol = "XGC";
        decimals = 18;
    }
}


contract XgoldCrowdsale is Crowdsale {

    uint256 public investors;


    function XgoldCrowdsale() public
    Crowdsale()
    {
        investors = 0;
    }


    function buyXgoldTokens(address _sender) public payable {
        investors++;
        buyTokens(_sender);
    }


    function() public payable {
        buyXgoldTokens(msg.sender);
    }

    function sendTokens(address _beneficiary, uint256 _amount) public payable onlyOwner {
        investors++;
        mintTokens(_beneficiary, _amount);
    }

}