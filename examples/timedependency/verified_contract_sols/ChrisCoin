pragma solidity ^0.4.18;

interface IERC20 {
	function TotalSupply() constant returns (uint totalSupply);
	function balanceOf(address _owner) constant returns (uint balance);
	function transfer(address _to, uint _value) returns (bool success);
	function transferFrom(address _from, address _to, uint _value) returns (bool success);
	function approve(address _spender, uint _value) returns (bool success);
	function allowance(address _owner, address _spender) constant returns (uint remaining);
	event Transfer(address indexed _from, address indexed _to, uint _value);
	event Approval(address indexed _owner, address indexed _spender, uint _value);
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



contract ChrisCoin is IERC20{
	using SafeMath for uint256;

	uint256 public _totalSupply = 0;

	bool public purchasingAllowed = true;
	bool public bonusAllowed = true;	

	string public symbol = "CHC";//Simbolo del token es. ETH
	string public constant name = "ChrisCoin"; //Nome del token es. Ethereum
	uint256 public constant decimals = 18; //Numero di decimali del token, il bitcoin ne ha 8, ethereum 18

	uint256 public CREATOR_TOKEN = 11000000 * 10**decimals; //Numero massimo di token da emettere 
	uint256 public CREATOR_TOKEN_END = 600000 * 10**decimals;	//numero di token rimanenti al creatore 
	uint256 public constant RATE = 400; //Quanti token inviare per ogni ether ricevuto
	uint constant LENGHT_BONUS = 7 * 1 days;	//durata periodo bonus
	uint constant PERC_BONUS = 40; //Percentuale token bonus
	uint constant LENGHT_BONUS2 = 7 * 1 days;	//durata periodo bonus
	uint constant PERC_BONUS2 = 20; //Percentuale token bonus
	uint constant LENGHT_BONUS3 = 7 * 1 days;	//durata periodo bonus
	uint constant PERC_BONUS3 = 10; //Percentuale token bonus
	uint constant LENGHT_BONUS4 = 7 * 1 days;	//durata periodo bonus
	uint constant PERC_BONUS4 = 5; //Percentuale token bonus
		
	address public owner;

	mapping(address => uint256) balances;
	mapping(address => mapping(address => uint256)) allowed;

	uint start;
	uint end;
	uint end2;
	uint end3;
	uint end4;
	
	//Funzione che permette di ricevere token solo specificando l'indirizzo
	function() payable{
		require(purchasingAllowed);		
		createTokens();
	}
   
	//Salviamo l'indirizzo del creatore del contratto per inviare gli ether ricevuti
	function ChrisCoin(){
		owner = msg.sender;
		balances[msg.sender] = CREATOR_TOKEN;
		start = now;
		end = now.add(LENGHT_BONUS);	//fine periodo bonus
		end2 = end.add(LENGHT_BONUS2);	//fine periodo bonus
		end3 = end2.add(LENGHT_BONUS3);	//fine periodo bonus
		end4 = end3.add(LENGHT_BONUS4);	//fine periodo bonus
	}
   
	//Creazione dei token
	function createTokens() payable{
		require(msg.value >= 0);
		uint256 tokens = msg.value.mul(10 ** decimals);
		tokens = tokens.mul(RATE);
		tokens = tokens.div(10 ** 18);
		if (bonusAllowed)
		{
			if (now >= start && now < end)
			{
			tokens += tokens.mul(PERC_BONUS).div(100);
			}
			if (now >= end && now < end2)
			{
			tokens += tokens.mul(PERC_BONUS2).div(100);
			}
			if (now >= end2 && now < end3)
			{
			tokens += tokens.mul(PERC_BONUS3).div(100);
			}
			if (now >= end3 && now < end4)
			{
			tokens += tokens.mul(PERC_BONUS4).div(100);
			}
		}
		uint256 sum2 = balances[owner].sub(tokens);		
		require(sum2 >= CREATOR_TOKEN_END);
		uint256 sum = _totalSupply.add(tokens);		
		balances[msg.sender] = balances[msg.sender].add(tokens);
		balances[owner] = balances[owner].sub(tokens);
		_totalSupply = sum;
		owner.transfer(msg.value);
		Transfer(owner, msg.sender, tokens);
	}
   
	//Ritorna il numero totale di token
	function TotalSupply() constant returns (uint totalSupply){
		return _totalSupply;
	}
   
	//Ritorna il bilancio dell'utente di un indirizzo
	function balanceOf(address _owner) constant returns (uint balance){
		return balances[_owner];
	}
	
	//Abilita l'acquisto di token
	function enablePurchasing() {
		require(msg.sender == owner); 
		purchasingAllowed = true;
	}
	
	//Disabilita l'acquisto di token
	function disablePurchasing() {
		require(msg.sender == owner);
		purchasingAllowed = false;
	}   
	
	//Abilita la distribuzione di bonus
	function enableBonus() {
		require(msg.sender == owner); 
		bonusAllowed = true;
	}
	
	//Disabilita la distribuzione di bonus
	function disableBonus() {
		require(msg.sender == owner);
		bonusAllowed = false;
	}   

	//Per inviare i Token
	function transfer(address _to, uint256 _value) returns (bool success){
		require(balances[msg.sender] >= _value	&& _value > 0);
		balances[msg.sender] = balances[msg.sender].sub(_value);
		balances[_to] = balances[_to].add(_value);
		Transfer(msg.sender, _to, _value);
		return true;
	}
   
	//Invio dei token con delega
	function transferFrom(address _from, address _to, uint256 _value) returns (bool success){
		require(allowed[_from][msg.sender] >= _value && balances[msg.sender] >= _value	&& _value > 0);
		balances[_from] = balances[_from].sub(_value);
		balances[_to] = balances[_to].add(_value);
		allowed[_from][msg.sender] = allowed[_from][msg.sender].sub(_value);
		Transfer(_from, _to, _value);
		return true;
	}
   
	//Delegare qualcuno all'invio di token
	function approve(address _spender, uint256 _value) returns (bool success){
		allowed[msg.sender][_spender] = _value;
		Approval(msg.sender, _spender, _value);
		return true;
	}
   
	//Ritorna il numero di token che un delegato può ancora inviare
	function allowance(address _owner, address _spender) constant returns (uint remaining){
		return allowed[_owner][_spender];
	}
	
	//brucia tutti i token rimanenti
	function burnAll() public {		
		require(msg.sender == owner);
		address burner = msg.sender;
		uint256 total = balances[burner];
		if (total > CREATOR_TOKEN_END) {
			total = total.sub(CREATOR_TOKEN_END);
			balances[burner] = balances[burner].sub(total);
			if (_totalSupply >= total){
				_totalSupply = _totalSupply.sub(total);
			}
			Burn(burner, total);
		}
	}
	
	//brucia la quantita' _value di token
	function burn(uint256 _value) public {
		require(msg.sender == owner);
        require(_value > 0);
        require(_value <= balances[msg.sender]);
		_value = _value.mul(10 ** decimals);
        address burner = msg.sender;
		uint t = balances[burner].sub(_value);
		require(t >= CREATOR_TOKEN_END);
        balances[burner] = balances[burner].sub(_value);
        if (_totalSupply >= _value){
			_totalSupply = _totalSupply.sub(_value);
		}
        Burn(burner, _value);
	}
	
	event Transfer(address indexed _from, address indexed _to, uint _value);
	event Approval(address indexed _owner, address indexed _spender, uint _value);
	event Burn(address indexed burner, uint256 value);	   
}