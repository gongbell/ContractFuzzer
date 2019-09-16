/**
 *Submitted for verification at Etherscan.io on 2017-12-22
*/

pragma solidity ^0.4.18;

//*** Owner ***//
contract owned {
	address public owner;
    
    //*** OwnershipTransferred ***//
    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

	function owned() public {
		owner = msg.sender;
	}

    //*** Change Owner ***//
	function changeOwner(address newOwner) onlyOwner public {
		owner = newOwner;
	}
    
    //*** Transfer OwnerShip ***//
    function transferOwnership(address newOwner) onlyOwner public {
        require(newOwner != address(0));
        OwnershipTransferred(owner, newOwner);
        owner = newOwner;
    }
    
    //*** Only Owner ***//
	modifier onlyOwner {
		require(msg.sender == owner);
		_;
	}
}

interface tokenRecipient { function receiveApproval(address _from, uint256 _value, address _token, bytes _extraData) public; }

//*** GraphenePowerToken ***//
contract GraphenePowerToken is owned{
    
    //************** Token ************//
	string public standard = 'Token 1';

	string public name = 'Graphene Power';

	string public symbol = 'GRP';

	uint8 public decimals = 18;

	uint256 public totalSupply =0;
	
	//*** Pre-sale ***//
    uint preSaleStart=1513771200;
    uint preSaleEnd=1515585600;
    uint256 preSaleTotalTokens=30000000;
    uint256 preSaleTokenCost=6000;
    address preSaleAddress;
    bool public enablePreSale=false;
    
    //*** ICO ***//
    uint icoStart;
    uint256 icoSaleTotalTokens=400000000;
    address icoAddress;
    bool public enableIco=false;
    
    //*** Advisers,Consultants ***//
    uint256 advisersConsultantTokens=15000000;
    address advisersConsultantsAddress;
    
    //*** Bounty ***//
    uint256 bountyTokens=15000000;
    address bountyAddress;
    
    //*** Founders ***//
    uint256 founderTokens=40000000;
    address founderAddress;
    
    //*** Walet ***//
    address public wallet;
    
    //*** TranferCoin ***//
    bool public transfersEnabled = false;
    
     //*** Balance ***//
    mapping (address => uint256) public balanceOf;
    
    //*** Alowed ***//
    mapping (address => mapping (address => uint256)) allowed;
    
    //*** Tranfer ***//
    event Transfer(address from, address to, uint256 value);
    
	//*** Approval ***//
	event Approval(address indexed _owner, address indexed _spender, uint256 _value);
	
	//*** Destruction ***//
	event Destruction(uint256 _amount);
	
	//*** Burn ***//
	event Burn(address indexed from, uint256 value);
	
	//*** Issuance ***//
	event Issuance(uint256 _amount);
	
	function GraphenePowerToken() public{
        preSaleAddress=0xC07850969A0EC345A84289f9C5bb5F979f27110f;
        icoAddress=0x1C21Cf57BF4e2dd28883eE68C03a9725056D29F1;
        advisersConsultantsAddress=0xe8B6dA1B801b7F57e3061C1c53a011b31C9315C7;
        bountyAddress=0xD53E82Aea770feED8e57433D3D61674caEC1D1Be;
        founderAddress=0xDA0D3Dad39165EA2d7386f18F96664Ee2e9FD8db;
        totalSupply =500000000;
        balanceOf[msg.sender]=totalSupply;
	}

	//*** Payable ***//
    function() payable public {
        require(msg.value>0);
        require(msg.sender != 0x0);
        
        uint256 weiAmount;
        uint256 tokens;
        wallet=owner;
        
        if(isPreSale()){
            wallet=preSaleAddress;
            weiAmount=6000;
        }
        else if(isIco()){
            wallet=icoAddress;
            
            if((icoStart+(7*24*60*60)) >= now){
               weiAmount=4000;
            }
            else if((icoStart+(14*24*60*60)) >= now){
                 weiAmount=3750;
            }
            else if((icoStart+(21*24*60*60)) >= now){
                 weiAmount=3500;
            }
            else if((icoStart+(28*24*60*60)) >= now){
                 weiAmount=3250;
            }
            else if((icoStart+(35*24*60*60)) >= now){
                 weiAmount=3000;
            }
            else{
                weiAmount=2000;
            }
        }
        else{
            weiAmount=4000;
        }
        
        tokens=msg.value*weiAmount/1000000000000000000;
        Transfer(this, msg.sender, tokens);
        balanceOf[msg.sender]+=tokens;
        totalSupply=(totalSupply-tokens);
        wallet.transfer(msg.value);
        balanceOf[this]+=msg.value;
	}
	
	/* Send coins */
	function transfer(address _to, uint256 _value) public returns (bool success) {
	    if(transfersEnabled){
		    require(balanceOf[_to] >= _value);
		    // Subtract from the sender
		    balanceOf[msg.sender] = (balanceOf[msg.sender] -_value);
	        balanceOf[_to] =(balanceOf[_to] + _value);
		    Transfer(msg.sender, _to, _value);
		    return true;
	    }
	    else{
	        return false;
	    }
	
	}

	//*** Transfer From ***//
	function transferFrom(address _from, address _to, uint256 _value) public returns (bool success) {
	    if(transfersEnabled){
	        // Check if the sender has enough
		    require(balanceOf[_from] >= _value);
		    // Check allowed
		    require(_value <= allowed[_from][msg.sender]);

		    // Subtract from the sender
		    balanceOf[_from] = (balanceOf[_from] - _value);
		    // Add the same to the recipient
		    balanceOf[_to] = (balanceOf[_to] + _value);

		    allowed[_from][msg.sender] = (allowed[_from][msg.sender] - _value);
		    Transfer(_from, _to, _value);
		    return true;
	    }
	    else{
	        return false;
	    }
	}
	
	//*** Transfer OnlyOwner ***//
	function transferOwner(address _to,uint256 _value) public onlyOwner returns(bool success){
	    // Subtract from the sender
	    totalSupply=(totalSupply-_value);
		// Add the same to the recipient
		balanceOf[_to] = (balanceOf[_to] + _value);
		Transfer(this, _to, _value);
	}
	
	//*** Allowance ***//
	function allowance(address _owner, address _spender) constant public returns (uint256 remaining) {
		return allowed[_owner][_spender];
	}
	
	//*** Approve ***//
	function approve(address _spender, uint256 _value) public returns (bool success) {
		allowed[msg.sender][_spender] = _value;
		Approval(msg.sender, _spender, _value);
		return true;
	}
	
	//*** Burn Owner***//
	function burnOwner(uint256 _value) public onlyOwner returns (bool success) {
		destroyOwner(msg.sender, _value);
		Burn(msg.sender, _value);
		return true;
	}
	
	//*** Destroy Owner ***//
	function destroyOwner(address _from, uint256 _amount) public onlyOwner{
	    balanceOf[_from] =(balanceOf[_from] - _amount);
		totalSupply = (totalSupply - _amount);
		Transfer(_from, this, _amount);
		Destruction(_amount);
	}
	
	//*** Kill Balance ***//
	function killBalance(uint256 _value) onlyOwner public {
		if(this.balance > 0) {
		    if(_value==1){
		        preSaleAddress.transfer(this.balance);
		        balanceOf[this]=0;
		    }
		    else if(_value==2){
		        icoAddress.transfer(this.balance);
		         balanceOf[this]=0;
		    }
		    else{
		        owner.transfer(this.balance);
		         balanceOf[this]=0;
		    }
		}
		else{
		    owner.transfer(this.balance);
		     balanceOf[this]=0;
		}
	}
	
	//*** Kill Tokens ***//
	function killTokens() onlyOwner public{
	    Transfer(this, bountyAddress, bountyTokens);
	    Transfer(this, founderAddress, founderTokens);
	    Transfer(this, advisersConsultantsAddress, advisersConsultantTokens);
	    totalSupply=totalSupply-(bountyTokens+founderTokens+advisersConsultantTokens);
	    bountyTokens=0;
	    founderTokens=0;
	    advisersConsultantTokens=0;
	}
	
	//*** Contract Balance ***//
	function contractBalance() constant public returns (uint256 balance) {
		return balanceOf[this];
	}
	
	//*** Set ParamsTransfer ***//
	function setParamsTransfer(bool _value) public onlyOwner{
	    transfersEnabled=_value;
	}
	
	//*** Set ParamsICO ***//
    function setParamsIco(bool _value) public onlyOwner returns(bool result){
        enableIco=_value;
    }
    
	//*** Set ParamsPreSale ***//
    function setParamsPreSale(bool _value) public onlyOwner returns(bool result){
        enablePreSale=_value;
    }
	
	//*** Is ico ***//
    function isIco() constant public returns (bool ico) {
		 bool result=((icoStart+(35*24*60*60)) >= now);
		 if(enableIco){
		     return true;
		 }
		 else{
		     return result;
		 }
	}
    
    //*** Is PreSale ***//
    function isPreSale() constant public returns (bool preSale) {
		bool result=(preSaleEnd >= now);
		if(enablePreSale){
		    return true;
		}
		else{
		    return result;
		}
	}
}
