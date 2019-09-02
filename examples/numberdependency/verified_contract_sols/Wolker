pragma solidity 0.4.11;

contract Wolker {
  mapping (address => uint256) balances;
  mapping (address => uint256) allocations;
  mapping (address => mapping (address => uint256)) allowed;
  mapping (address => mapping (address => bool)) authorized; //trustee

  /// @param _to The address of the recipient
  /// @param _value The amount of token to be transferred
  /// @return Whether the transfer was successful or not
  function transfer(address _to, uint256 _value) returns (bool success) {
    if (balances[msg.sender] >= _value && _value > 0) {
      balances[msg.sender] = safeSub(balances[msg.sender], _value);
      balances[_to] = safeAdd(balances[_to], _value);
      Transfer(msg.sender, _to, _value, balances[msg.sender], balances[_to]);
      return true;
    } else {
      throw;
    }
  }
  
  /// @param _from The address of the sender
  /// @param _to The address of the recipient
  /// @param _value The amount of token to be transferred
  /// @return Whether the transfer was successful or not
  function transferFrom(address _from, address _to, uint256 _value) returns (bool success) {
    var _allowance = allowed[_from][msg.sender];
    if (balances[_from] >= _value && allowed[_from][msg.sender] >= _value && _value > 0) {
      balances[_to] = safeAdd(balances[_to], _value);
      balances[_from] = safeSub(balances[_from], _value);
      allowed[_from][msg.sender] = safeSub(_allowance, _value);
      Transfer(_from, _to, _value, balances[_from], balances[_to]);
      return true;
    } else {
      throw;
    }
  }
 
  /// @return total amount of tokens
  function totalSupply() external constant returns (uint256) {
        return generalTokens + reservedTokens;
  }
 
  /// @param _owner The address from which the balance will be retrieved
  /// @return The balance
  function balanceOf(address _owner) constant returns (uint256 balance) {
    return balances[_owner];
  }


  /// @param _spender The address of the account able to transfer the tokens
  /// @param _value The amount of Wolk token to be approved for transfer
  /// @return Whether the approval was successful or not
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
  function check_authorization(address _owner, address _trustee) constant returns (bool authorization_status) {
    return authorized[_owner][_trustee];
  }

  /// @param _owner The address of the account owning tokens
  /// @param _spender The address of the account able to transfer the tokens
  /// @return Amount of remaining tokens allowed to spent
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
  event CreateWolk(address indexed _to, uint256 _value);
  event Vested(address indexed _to, uint256 _value);

  modifier onlyOwner {
    assert(msg.sender == owner);
    _;
  }

  modifier isOperational() {
    assert(isFinalized);
    _;
  }


  //**** ERC20 TOK fields:
  string  public constant name = 'Wolk';
  string  public constant symbol = "WOLK";
  string  public constant version = "0.2";
  uint256 public constant decimals = 18;
  uint256 public constant wolkFund  =  10 * 10**1 * 10**decimals;        //  100 Wolk in operation Fund
  uint256 public constant tokenCreationMin =  20 * 10**1 * 10**decimals; //  200 Wolk Min
  uint256 public constant tokenCreationMax = 100 * 10**1 * 10**decimals; // 1000 Wolk Max
  uint256 public constant tokenExchangeRate = 10000;   // 10000 Wolk per 1 ETH
  uint256 public generalTokens = wolkFund; // tokens in circulation
  uint256 public reservedTokens; 

  //address public owner = msg.sender;
  address public owner = 0xC28dA4d42866758d0Fc49a5A3948A1f43de491e9; // michael - main
  address public multisig_owner = 0x6968a9b90245cB9bD2506B9460e3D13ED4B2FD1e; // new multi-sig

  bool public isFinalized = false;          // after token sale success, this is true
  uint public constant dust = 1000000 wei; 
  bool public fairsale_protection = true;

  
     // Actual crowdsale
  uint256 public start_block;                // Starting block
  uint256 public end_block;                  // Ending block
  uint256 public unlockedAt;                 // Unlocking block 
 
  uint256 public end_ts;                     // Unix End time


  // minting support
  //uint public max_creation_rate_per_second; // Maximum token creation rate per second
  //address public minter_address;            // Has permission to mint

  // migration support
  //address migrationMaster;


  //**** Constructor:
  function Wolk() 
  {

    // Actual crowdsale
    start_block = 3831300;
    end_block = 3831900;

    // wolkFund is 100
    balances[msg.sender] = wolkFund;

    // Wolk Inc has 25MM Wolk, 5MM of which is allocated for Wolk Inc Founding staff, who vest at "unlockedAt" time
    reservedTokens = 25 * 10**decimals;
    allocations[0x564a3f7d98Eb5B1791132F8875fef582d528d5Cf] = 20; // unassigned
    allocations[0x7f512CCFEF05F651A70Fa322Ce27F4ad79b74ffe] = 1;  // Sourabh
    allocations[0x9D203A36cd61b21B7C8c7Da1d8eeB13f04bb24D9] = 2;  // Michael - Test
    allocations[0x5fcf700654B8062B709a41527FAfCda367daE7b1] = 1;  // Michael - Main
    allocations[0xC28dA4d42866758d0Fc49a5A3948A1f43de491e9] = 1;  // Urmi
    
    
    CreateWolk(msg.sender, wolkFund); 
  }

  // ****** VESTING SUPPORT
  /// @notice Allow developer to unlock allocated tokens by transferring them to developer's address on vesting schedule of "vested 100% on 1 year)
  function unlock() external {
    if (now < unlockedAt) throw;
    uint256 vested = allocations[msg.sender] * 10**decimals;
    if (vested < 0 ) throw; // Will fail if allocation (and therefore toTransfer) is 0.
    allocations[msg.sender] = 0;
    reservedTokens = safeSub(reservedTokens, vested);
    balances[msg.sender] = safeAdd(balances[msg.sender], vested); 
    Vested(msg.sender, vested);
  }

  // ******* CROWDSALE SUPPORT
  // Accepts ETH and creates WOLK
  function createTokens() payable external is_not_dust {
    if (isFinalized) throw;
    if (block.number < start_block) throw;
    if (block.number > end_block) throw;
    if (msg.value == 0) throw;
    if (tx.gasprice > 0.021 szabo && fairsale_protection) throw; 
    if (msg.value > 0.04 ether && fairsale_protection) throw; 

    uint256 tokens = safeMul(msg.value, tokenExchangeRate); // check that we're not over totals
    uint256 checkedSupply = safeAdd(generalTokens, tokens);
    if ( checkedSupply > tokenCreationMax) { 
      throw; // they need to get their money back if something goes wrong
    } else {
      generalTokens = checkedSupply;
      balances[msg.sender] = safeAdd(balances[msg.sender], tokens);   // safeAdd not needed; bad semantics to use here
      CreateWolk(msg.sender, tokens); // logs token creation
    }
  }
  
  // The value of the message must be sufficiently large to not be considered dust.
  modifier is_not_dust { if (msg.value < dust) throw; _; }

  // Disabling fairsale protection  
  function fairsale_protectionOFF() external {
    if ( block.number - start_block < 200) throw; // fairsale window is strictly enforced
    if ( msg.sender != owner ) throw;
    fairsale_protection = false;
  }

  // Finalizing the crowdsale
  function finalize() external {
    if ( isFinalized ) throw;
    if ( msg.sender != owner ) throw;  // locks finalize to ETH owner
    if ( generalTokens < tokenCreationMin ) throw; // have to sell tokenCreationMin to finalize
    if ( block.number < end_block ) throw;  
    isFinalized = true;
    end_ts = now;
    unlockedAt = end_ts + 2 minutes;
    if ( ! multisig_owner.send(this.balance) ) throw;
  }

  function refund() external {
    if ( isFinalized ) throw; 
    if ( block.number < end_block ) throw;   
    if ( generalTokens >= tokenCreationMin ) throw;  
    if ( msg.sender == owner ) throw;
    uint256 Val = balances[msg.sender];
    balances[msg.sender] = 0;
    generalTokens = safeSub(generalTokens, Val);
    uint256 ethVal = safeDiv(Val, tokenExchangeRate);
    LogRefund(msg.sender, ethVal);
    if ( ! msg.sender.send(ethVal) ) throw;
  }
    
  // ****** Platform Settlement
  function settleFrom(address _from, address _to, uint256 _value) isOperational() external returns (bool success) {
    if ( msg.sender != owner ) throw;
    var _allowance = allowed[_from][msg.sender];
    if (balances[_from] >= _value && (allowed[_from][msg.sender] >= _value || authorized[_from][msg.sender] == true ) && _value > 0) {
      balances[_to] = safeAdd(balances[_to], _value);
      balances[_from] = safeSub(balances[_from], _value);
      allowed[_from][msg.sender] = safeSub(_allowance, _value);
      if ( allowed[_from][msg.sender] < 0 ){
         allowed[_from][msg.sender] = 0;
      }
      Transfer(_from, _to, _value, balances[_from], balances[_to]);
      return true;
    } else {
      throw;
    }
  }

  // ****** MINTING SUPPORT
  // Mint new tokens
  modifier only_minter {
    assert(msg.sender == minter_address);
    _;
  }
  
  address public minter_address = owner;            // Has permission to mint

  function mintTokens(uint reward_tok, address recipient) external payable only_minter
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