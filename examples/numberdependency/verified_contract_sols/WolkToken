pragma solidity 0.4.11;

contract WolkToken {
  mapping (address => uint256) balances;
  mapping (address => uint256) allocations;
  mapping (address => mapping (address => uint256)) allowed;
  mapping (address => mapping (address => bool)) authorized; //trustee

  function transfer(address _to, uint256 _value) isTransferable returns (bool success) {
    if (balances[msg.sender] >= _value && _value > 0) {
      balances[msg.sender] = safeSub(balances[msg.sender], _value);
      balances[_to] = safeAdd(balances[_to], _value);
      Transfer(msg.sender, _to, _value, balances[msg.sender], balances[_to]);
      return true;
    } else {
      return false;
    }
  }
  
  function transferFrom(address _from, address _to, uint256 _value) isTransferable returns (bool success) {
    var _allowance = allowed[_from][msg.sender];
    if (balances[_from] >= _value && allowed[_from][msg.sender] >= _value && _value > 0) {
      balances[_to] = safeAdd(balances[_to], _value);
      balances[_from] = safeSub(balances[_from], _value);
      allowed[_from][msg.sender] = safeSub(_allowance, _value);
      Transfer(_from, _to, _value, balances[_from], balances[_to]);
      return true;
    } else {
      return false;
    }
  }
 
   // Platform Settlement
  function settleFrom(address _from, address _to, uint256 _value) isTransferable returns (bool success) {
    var _allowance = allowed[_from][msg.sender];
    var isPreauthorized = authorized[_from][msg.sender];
    if (balances[_from] >= _value && ( isPreauthorized || _allowance >= _value ) && _value > 0) {
      balances[_to] = safeAdd(balances[_to], _value);
      balances[_from] = safeSub(balances[_from], _value);
      Transfer(_from, _to, _value, balances[_from], balances[_to]);
      if (isPreauthorized && _allowance < _value){
          allowed[_from][msg.sender] = 0;
      }else{
          allowed[_from][msg.sender] = safeSub(_allowance, _value);
      }
      return true;
    } else {
      return false;
    }
  }


  function totalSupply() external constant returns (uint256) {
        return generalTokens + reservedTokens;
  }
 

  function balanceOf(address _owner) constant returns (uint256 balance) {
    return balances[_owner];
  }


  function approve(address _spender, uint256 _value) returns (bool success) {
    allowed[msg.sender][_spender] = _value;
    Approval(msg.sender, _spender, _value);
    return true;
  }


  /// @param _trustee Grant trustee permission to settle media spend
  /// @return Whether the authorization was successful or not
  function authorize(address _trustee) returns (bool success) {
    authorized[msg.sender][_trustee] = true;
    Authorization(msg.sender, _trustee);
    return true;
  }

  /// @param _trustee_to_remove Revoke trustee's permission on settle media spend 
  /// @return Whether the deauthorization was successful or not
  function deauthorize(address _trustee_to_remove) returns (bool success) {
    authorized[msg.sender][_trustee_to_remove] = false;
    Deauthorization(msg.sender, _trustee_to_remove);
    return true;
  }

  // @param _owner
  // @param _trustee
  // @return authorization_status for platform settlement 
  function checkAuthorization(address _owner, address _trustee) constant returns (bool authorization_status) {
    return authorized[_owner][_trustee];
  }

  function allowance(address _owner, address _spender) constant returns (uint256 remaining) {
    return allowed[_owner][_spender];
  }


  //**** ERC20 TOK Events:
  event Transfer(address indexed _from, address indexed _to, uint256 _value, uint from_final_tok, uint to_final_tok);
  event Approval(address indexed _owner, address indexed _spender, uint256 _value);
  event Authorization(address indexed _owner, address indexed _trustee);
  event Deauthorization(address indexed _owner, address indexed _trustee_to_remove);

  event NewOwner(address _newOwner);
  event MintEvent(uint reward_tok, address recipient);
  event LogRefund(address indexed _to, uint256 _value);
  event WolkCreated(address indexed _to, uint256 _value);
  event Vested(address indexed _to, uint256 _value);

  modifier onlyOwner { assert(msg.sender == owner); _; }
  modifier isOperational { assert(saleCompleted); _; }
  modifier isTransferable { assert(generalTokens > crowdSaleMin); _;}
  modifier is_not_dust { if (msg.value < dust) throw; _; } // must be sufficiently large (1 mwei) to not be considered dust.
  
  //**** ERC20 TOK fields:
  string  public constant name = 'Wolk Coin';
  string  public constant symbol = "WOLK";
  string  public constant version = "1.0";
  uint256 public constant decimals = 18;
  uint256 public constant wolkFund  =  10 * 10**1 * 10**decimals;        //  100 Wolk in operation Fund
  uint256 public constant crowdSaleMin =  10 * 10**3 * 10**decimals; //  10000 Wolk Min
  uint256 public constant crowdSaleMax =  10 * 10**5 * 10**decimals; //  1000000 Wolk Max
  uint256 public constant tokenExchangeRate = 10000;   // 10000 Wolk per 1 ETH
  uint256 public constant dust = 1000000 wei; // 1 Mwei

  uint256 public generalTokens = wolkFund;     // tokens in circulation
  uint256 public reservedTokens;              // Unvested developer tokens

  address public owner = 0x5fcf700654B8062B709a41527FAfCda367daE7b1; // MK - main
  address public multisigWallet = 0x6968a9b90245cB9bD2506B9460e3D13ED4B2FD1e; 

  uint256 public constant start_block = 3843600;   // Sale Starting block
  uint256 public end_block = 3847200;              // Sale Ending block when crowdSaleMax not reached
  uint256 public unlockedAt;                       // team vesting  
  uint256 public end_ts;                           // sale End time
  
  bool public saleCompleted = false;               
  bool public fairsaleProtection = true;         

 

  // Migration support
  //address migrationMaster;


  // Constructor:
  function Wolk() onlyOwner {

    // Wolk Inc has 25MM Wolk, 5MM of which is allocated for Wolk Inc Founding staff, who vest at "unlockedAt" time
    reservedTokens = 25 * 10**decimals;
    allocations[0x564a3f7d98Eb5B1791132F8875fef582d528d5Cf] = 20; // unassigned
    allocations[0x7f512CCFEF05F651A70Fa322Ce27F4ad79b74ffe] = 1;  // Sourabh
    allocations[0x9D203A36cd61b21B7C8c7Da1d8eeB13f04bb24D9] = 2;  // Michael - Test
    allocations[0x5fcf700654B8062B709a41527FAfCda367daE7b1] = 1;  // Michael - Main
    allocations[0xC28dA4d42866758d0Fc49a5A3948A1f43de491e9] = 1;  // Urmi
    
    balances[owner] = wolkFund; // 100 wOlk growth fund
    WolkCreated(owner, wolkFund);
  }

  // VESTING SUPPORT
  function unlock() external {
    if (now < unlockedAt) throw;
    uint256 vested = allocations[msg.sender] * 10**decimals;
    if (vested < 0 ) throw; // Will fail if allocation (and therefore toTransfer) is 0.
    allocations[msg.sender] = 0;
    reservedTokens = safeSub(reservedTokens, vested);
    balances[msg.sender] = safeAdd(balances[msg.sender], vested); 
    Vested(msg.sender, vested);
  }

  // CROWDSALE SUPPORT
  function redeemToken() payable is_not_dust external {
    if (saleCompleted) throw;
    if (block.number < start_block) throw;
    if (block.number > end_block) throw;
    if (msg.value <= dust) throw;
    if (tx.gasprice > 0.46 szabo && fairsaleProtection) throw; 
    if (msg.value > 1 ether && fairsaleProtection) throw; 

    uint256 tokens = safeMul(msg.value, tokenExchangeRate); // check that we're not over totals
    uint256 checkedSupply = safeAdd(generalTokens, tokens);
    if ( checkedSupply > crowdSaleMax) throw; // they need to get their money back if something goes wrong
      generalTokens = checkedSupply;
      balances[msg.sender] = safeAdd(balances[msg.sender], tokens);   // safeAdd not needed; bad semantics to use here
      WolkCreated(msg.sender, tokens); // logs token creation
    
  }
  


  // Disabling fairsale protection  
  function fairsaleProtectionOFF() onlyOwner {
    if ( block.number - start_block < 2000) throw; // fairsale window is strictly enforced for the first 2000 blocks
    fairsaleProtection = false;
  }


  // Finalizing the crowdsale
  function finalize() onlyOwner {
    if ( saleCompleted ) throw;
    if ( generalTokens < crowdSaleMin ) throw; 
    if ( block.number < end_block ) throw;  
    saleCompleted = true;
    end_ts = now;
    end_block = block.number; 
    unlockedAt = end_ts + 30 minutes;
    if ( ! multisigWallet.send(this.balance) ) throw;
  }

  function withdraw() onlyOwner{ 		
		if ( this.balance == 0) throw;
		if ( generalTokens < crowdSaleMin) throw;	
        if ( ! multisigWallet.send(this.balance) ) throw;
  }


  function refund() {
    if ( saleCompleted ) throw; 
    if ( block.number < end_block ) throw;   
    if ( generalTokens >= crowdSaleMin ) throw;  
    if ( msg.sender == owner ) throw;
    uint256 Val = balances[msg.sender];
    balances[msg.sender] = 0;
    generalTokens = safeSub(generalTokens, Val);
    uint256 ethVal = safeDiv(Val, tokenExchangeRate);
    LogRefund(msg.sender, ethVal);
    if ( ! msg.sender.send(ethVal) ) throw;
  }
    

  // MINTING SUPPORT - Rewarding growth tokens for value-addeddata suppliers
  
  modifier onlyMinter { assert(msg.sender == minter_address); _; }
 
  address public minter_address = owner;

 // minting support
  //uint public max_creation_rate_per_second; // Maximum token creation rate per second
  //address public minter_address;            // Has permission to mint
  
  function mintTokens(uint reward_tok, address recipient) external payable onlyMinter isOperational
  {
    balances[recipient] = safeAdd(balances[recipient], reward_tok);
    generalTokens = safeAdd(generalTokens, reward_tok);
    MintEvent(reward_tok, recipient);
  }

  function changeMintingAddress(address newAddress) onlyOwner returns (bool success) { 
    minter_address = newAddress; 
    return true;
  }

  
  //**** SafeMath:
  function safeMul(uint a, uint b) internal returns (uint) {
    uint c = a * b;
    assert(a == 0 || c / a == b);
    return c;
  }
  
  function safeDiv(uint a, uint b) internal returns (uint) {
    assert(b > 0);
    uint c = a / b;
    assert(a == b * c + a % b);
    return c;
  }
  
  function safeSub(uint a, uint b) internal returns (uint) {
    assert(b <= a);
    return a - b;
  }
  
  function safeAdd(uint a, uint b) internal returns (uint) {
    uint c = a + b;
    assert(c>=a && c>=b);
    return c;
  }

  function assert(bool assertion) internal {
    if (!assertion) throw;
  }
}