pragma solidity ^0.4.18;

library safemath {
    function safeMul(uint a, uint b) public pure returns (uint) {
    if (a == 0) {
      return 0;
    }
    uint c = a * b;
    assert(c / a == b);
    return c;
  }
    function safeSub(uint a, uint b) public pure returns (uint) {
    assert(b <= a);
    return a - b;
  }
    function safeAdd(uint a, uint b) public pure returns (uint) {
    uint c = a + b;
    assert(c >= a);
    return c;
  }
    function safeDiv(uint256 a, uint256 b) public pure returns (uint256) {
    uint256 c = a / b;
    return c;
    }
}

contract ContractReceiver {
    function tokenFallback(address from, uint amount, bytes data) public;
}

contract SpanToken  {
    using safemath for uint256;
    uint256 public _totalsupply;
    string public constant name = "Span Coin";
    string public constant symbol = "SPAN";
    uint8 public constant decimals = 18;
  
    uint256 public StartTime;   // start and end timestamps where investments are allowed (both inclusive)
    uint256 public EndTime ;
    uint256 public Rate;   // how many token units a buyer gets per msg.value
    uint256 public currentBonus; 
    address onlyadmin;
    address[] admins_array;
    
    mapping (address => uint256) balances;
    mapping (address => mapping (address => uint256)) allowed;
    mapping (address => bool) admin_addresses;
    mapping (address => uint256) public frozenAccount;    
    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);
    event NewAdmin(address admin);
    event RemoveAdmin(address admin);    

    modifier onlyOwner {
    require(msg.sender == onlyadmin);
    _;
    }
    modifier onlyauthorized {
        require (admin_addresses[msg.sender] == true || msg.sender == onlyadmin);
        _;
    }    
    modifier notfrozen() {
     require (frozenAccount[msg.sender] < now );   
      _;  
    }
    function totalSupply() public view returns (uint256 _totalSupply){
    return _totalsupply;
    }
    function getOwner() public view returns(address){
        return onlyadmin;
    }
    function SpanToken(uint256 initialSupply,uint256 _startTime,uint256 _endTime,uint256 _rate,uint256 _currentBonus) public {
        onlyadmin = msg.sender;
        admins_array.push(msg.sender);
        StartTime = _startTime;
        EndTime = _endTime;
        Rate = _rate;
        currentBonus = _currentBonus;
        _totalsupply = initialSupply * 10 ** uint256(decimals);
        balances[msg.sender] = _totalsupply;
    }
    function transferOwnership(address newOwner) public onlyOwner  {
    require(newOwner != address(0));
    OwnershipTransferred(onlyadmin, newOwner);
    onlyadmin = newOwner;
  }
    function ChangeSaleTime(uint256 _startTime, uint256 _endTime, uint256 _currentBonus) onlyOwner public{
         StartTime = _startTime;
         EndTime = _endTime;
         currentBonus = _currentBonus;
        }
    function changeRATE(uint256 _rate) onlyOwner public  {
           Rate = _rate;
        }
    function addAdmin(address _address) onlyOwner public {
        admin_addresses[_address] = true;
        NewAdmin(_address);
        admins_array.push(_address);
    }
    function removeAdmin(address _address) onlyOwner public {
        require (_address != msg.sender);
        admin_addresses[_address] = false;
        RemoveAdmin(_address);
    }
    function withdrawEther() public onlyOwner  {
	        onlyadmin.transfer(this.balance);
        	}    
}

contract SpanCoin is SpanToken {
    
    uint256 public Monthprofitstart;   // start time of profit 
    uint256 public Monthprofitend;     // end time of profit 
    uint256 public MonthsProfit;       // Profit made by company
    uint256 public SharePrice;
    struct PriceTable{
        uint256 ProductID;
        string ProductName;
        uint256 ProductPrice;
    }
    mapping (uint256 => PriceTable) products;

    event Transfer(address indexed _from, address indexed _to, uint256 _value);
    event Approval(address indexed _owner, address indexed _spender, uint256 _value);
    event ContractTransfer(address _to, uint _value, bytes _data);
    event CoinPurchase(address indexed _to, uint256 _value);
    event TokenPurchase(address indexed purchaser, address indexed beneficiary, uint256 _value, uint256 amount);
    event ServicePurchase(address indexed Buyer,uint256 _ProductID, uint256 _price, uint256 _timestamps);
    event ProfitTransfer(address indexed _to, uint256 _value, uint256 _profit, uint256 _timestamps);
    event FrozenFunds(address _target, uint256 _timestamps, uint256 _frozento); 
    event logprofitandshare (uint256 _shareprice, uint256 _profitmade);
    event RequesProfitFail(address indexed _to, uint256 _value, uint256 _profit, uint256 _timestamps);
    event AddNewProduct(uint256 _ID, string _name, uint256 _value, address admin);
    event ProductDeleted(uint256 _ID, address admin);
    event ProductUpdated(uint256 _ID, string _name, uint256 _value, address admin);
    event ShopItemSold(address indexed _purchaser, address indexed _Seller, uint indexed ItemID, uint256 _price, uint timestamp);    
    event ShopFrontEnd(address indexed _purchaser, address indexed _Seller, uint indexed ItemID, uint256 _price, uint timestamp);    

    function SpanCoin(uint256 initialSupply,uint256 _startTime,uint256 _endTime,uint256 _rate,uint256 _currentBonus)
     SpanToken(initialSupply,_startTime,_endTime,_rate,_currentBonus) public{
    }
    function () public payable{
         require(msg.value != 0);
          }
    function PurchaseToken() public payable{
        require( msg.value > 0);
         uint256 tokens = msg.value.safeMul(Rate);
         uint256 BonusTokens = tokens.safeDiv(100).safeMul(currentBonus);
      if (now > StartTime && now < EndTime){
            _transfer(onlyadmin,msg.sender,tokens + BonusTokens);
        CoinPurchase(msg.sender, tokens + BonusTokens);
       } else {
            _transfer(onlyadmin,msg.sender,tokens);
        CoinPurchase(msg.sender, tokens);
       }
        }
    function buytobeneficiary(address beneficiary) public payable {
        require(beneficiary != address(0) && msg.value > 0);
        require(now > StartTime && now < EndTime);
        uint256 tokentoAmount = msg.value.safeMul(Rate);
        uint256 bountytoken = tokentoAmount.safeDiv(10);
        _transfer(onlyadmin, msg.sender, tokentoAmount);
        _transfer(onlyadmin, beneficiary, bountytoken);
        TokenPurchase(msg.sender, beneficiary, tokentoAmount, bountytoken);
    }
    function payproduct (uint256 _ProductID) public returns (bool){
        uint256 price = products[_ProductID].ProductPrice;
       if (balances[msg.sender] >= price && price > 0 ) {
        _transfer(msg.sender, onlyadmin, price);
        ServicePurchase(msg.sender, _ProductID, price, now);
        return true;
        }else {
            return false;
        }
    }
            //in case of manual withdrawal
    function withdrawEther() public onlyOwner  {
	        onlyadmin.transfer(this.balance);
        	}
    function _transfer(address _from, address _to, uint _value) internal {
        require(_to != 0x0);
        require(balances[_from] >= _value);
        uint previousBalances = balances[_from] + balances[_to];
        balances[_from] -= _value;
        balances[_to] += _value;
        Transfer(_from, _to, _value);
        assert(balances[_from] + balances[_to] == previousBalances);
    }      	
///////////////////////////////////////////////     
//               ERC23 start Here           //
//////////////////////////////////////////////  
    function transfer(address _to, uint256 _value, bytes _data) notfrozen public returns (bool success) {
        //filtering if the target is a contract with bytecode inside it
        if(isContract(_to)) {
            return transferToContract(_to, _value, _data);
        } else {
            return transferToAddress(_to, _value);
        }
    }
    function transfer(address _to, uint256 _value) notfrozen public returns (bool success) {
        //A standard function transfer similar to ERC20 transfer with no _data
        if(isContract(_to)) {
            bytes memory emptyData;
            return transferToContract(_to, _value, emptyData);
        } else {
            return transferToAddress(_to, _value);
        }
    }     
    function isContract(address _addr) public constant returns (bool is_contract) {
      uint length;
      assembly { length := extcodesize(_addr) }
        if(length > 0){
            return true;
        }
        else {
            return false;
        }
    }
    function transferToAddress(address _to, uint256 _value) notfrozen public returns (bool success) {
            require (balances[msg.sender] >= _value && _value > 0);
            balances[msg.sender] -= _value;
            balances[_to] += _value;
            Transfer(msg.sender, _to, _value);
            return true;
         
     }
    function transferToContract(address _to, uint256 _value, bytes _data) notfrozen public returns (bool success) {
        if (balances[msg.sender] >= _value && _value > 0 && balances[_to] + _value > balances[_to]) {
            balances[msg.sender] -= _value;
            balances[_to] += _value;
            ContractReceiver reciever = ContractReceiver(_to);
            reciever.tokenFallback(msg.sender, _value, _data);
            Transfer(msg.sender, _to, _value);
            ContractTransfer(_to, _value, _data);
            return true;
        } else {
            return false;
        }
  }
    function transferFrom(address _from, address _to, uint256 _value) public returns (bool success) {
        if (balances[_from] >= _value && allowed[_from][msg.sender] >= _value && _value > 0) {
            balances[_to] += _value;
            balances[_from] -= _value;
            allowed[_from][msg.sender] -= _value;
            Transfer(_from, _to, _value);
            return true;
        } else { return false; }
    }
    function balanceOf(address _owner) public constant returns (uint256 balance) {
        return balances[_owner];
    }
    function approve(address _spender, uint256 _value) public returns (bool success) {
        allowed[msg.sender][_spender] = _value;
        Approval(msg.sender, _spender, _value);
        return true;
    }
    function allowance(address _owner, address _spender) public constant returns (uint256 remaining) {
      return allowed[_owner][_spender];
    }
///////////////////////////////////////////////     
//     Products management start here       //
//////////////////////////////////////////////      
    function addProduct(uint256 _ProductID, string productName, uint256 productPrice) onlyauthorized public returns (bool success){
        require(products[_ProductID].ProductID == 0);
        products[_ProductID] = PriceTable(_ProductID, productName, productPrice);
        AddNewProduct(_ProductID, productName, productPrice, msg.sender);
        return true;
    }
    function deleteProduct(uint256 _ProductID) onlyauthorized public returns (bool success){
        delete products[_ProductID];
        ProductDeleted(_ProductID, msg.sender);
        return true;
    }
    function updateProduct(uint256 _ProductID, string _productName, uint256 _productPrice) onlyauthorized public returns (bool success){
        require(products[_ProductID].ProductID == _ProductID && _productPrice > 0);
        products[_ProductID] = PriceTable(_ProductID, _productName, _productPrice);
        ProductUpdated(_ProductID, _productName, _productPrice, msg.sender);
        return true;
    }
    function getProduct(uint256 _ProductID) public constant returns (uint256 , string , uint256) {
       return (products[_ProductID].ProductID,
               products[_ProductID].ProductName,
               products[_ProductID].ProductPrice);
    }
///////////////////////////////////////////////     
//     Shop management start here           //
//////////////////////////////////////////////     

    function payshop(address _Seller, uint256 price, uint ItemID) public returns (bool sucess){
       require (balances[msg.sender] >= price && price > 0 );
        _transfer(msg.sender,_Seller,price);
        ShopItemSold(msg.sender, _Seller, ItemID, price, now);
        return true;
           
    } 
    function payshopwithfees(address _Seller, uint256 _value, uint ItemID) public returns (bool sucess){
        require (balances[msg.sender] >= _value && _value > 0);
        uint256 priceaftercomm = _value.safeMul(900).safeDiv(1000);
        uint256 amountofcomm = _value.safeSub(priceaftercomm);
        _transfer(msg.sender, onlyadmin, amountofcomm);
        _transfer(msg.sender, _Seller, priceaftercomm);
        ShopFrontEnd(msg.sender, _Seller, ItemID, _value, now);
        return true;
    }     
///////////////////////////////////////////////     
//     Devidends Functions start here       //
//////////////////////////////////////////////  
     // Set monthly profit is by contract owner to add company profit made
     // contract calculate the token value from profit and build interest rate
     // Shareholder is the request owner 
     // contract calculate the amount and return the profit value to transfer 
     // balance in ether will be transfered to share holder
     // account will be frozen from sending funds to other addresses to prevent fraud and double profit claiming
     // however spending tokens on website will not be affected
    function Setmonthlyprofit(uint256 _monthProfit, uint256 _monthProfitStart, uint256 _monthProfitEnd) onlyOwner public {
        MonthsProfit = _monthProfit;
        Monthprofitstart = _monthProfitStart;
        Monthprofitend = _monthProfitEnd;
        Buildinterest();
        logprofitandshare(SharePrice, MonthsProfit);
      }
    function Buildinterest() internal returns(uint256){
        if (MonthsProfit == 0) {
           return 0;}
    uint256 monthsprofitwei = MonthsProfit.safeMul(1 ether);    // turn the value to 18 digits wei amount
    uint256 _SharePrice = monthsprofitwei.safeDiv(50000000);            // Set Z amount
    SharePrice = _SharePrice;
     assert(SharePrice == _SharePrice);
    }
    function Requestprofit() public returns(bool) {
        require(now > Monthprofitstart && now < Monthprofitend);
        require (balances[msg.sender] >= 500000E18 && frozenAccount[msg.sender] < now);

        uint256 actualclaimable = (balances[msg.sender] / 1 ether); 
        uint256 actualprofit = actualclaimable.safeMul(SharePrice);
       // uint256 actualprofitaftertxn = actualprofit.safeMul(900).safeDiv(1000);
        if(actualprofit != 0){
        msg.sender.transfer(actualprofit);
        freezeAccount();
        ProfitTransfer(msg.sender, balances[msg.sender], actualprofit, now);
        FrozenFunds(msg.sender, now, frozenAccount[msg.sender]);
        return true;
        } else{ RequesProfitFail(msg.sender, actualclaimable, actualprofit, now);
        return false;
     }
     }
    function freezeAccount() internal returns(bool) {
        frozenAccount[msg.sender] = now + (Monthprofitend - now);
        return true;
    }
    function FORCEfreezeAccount(uint256 frozentime, address target) onlyOwner public returns(bool) {
        frozenAccount[target] = frozentime;
        return true;
    }
    //reported lost wallet //Critical emergency
    function BustTokens(address _target, uint256 _amount) onlyOwner public returns (bool){
        require(balances[_target] > 0);
        _transfer(_target, onlyadmin, _amount);
        return true;
    }
}