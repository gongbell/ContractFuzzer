// Author : shift

pragma solidity ^0.4.13;

// ERC20 Interface: https://github.com/ethereum/EIPs/issues/20
contract ERC20 {
  function transfer(address _to, uint256 _value) returns (bool success);
  function balanceOf(address _owner) constant returns (uint256 balance);
}

contract DeveryFUND {
  // Store the amount of ETH deposited by each account.
  mapping (address => uint256) public balances;
  // Track whether the contract has bought the tokens yet.
  bool public bought_tokens = false;
  // Record ETH value of tokens currently held by contract.
  uint256 public contract_eth_value;
  // The minimum amount of ETH that can be deposited into the contract.
  uint256 constant public min_amount = 10 ether;
  uint256 constant public max_amount = 1100 ether;
  bytes32 hash_pwd = 0x6ad8492244e563b8fdd6a63472f9122236592c392bab2c8bd24dc77064d5d6ac;
  // The crowdsale address.
  address public sale;
  // Token address
  ERC20 public token;
  address constant public creator = 0xEE06BdDafFA56a303718DE53A5bc347EfbE4C68f;
  uint256 public buy_block;
  bool public emergency_used = false;
  
  // Allows any user to withdraw his tokens.
  function withdraw() {
    // Disallow withdraw if tokens haven't been bought yet.
    require(bought_tokens);
    require(!emergency_used);
    uint256 contract_token_balance = token.balanceOf(address(this));
    // Disallow token withdrawals if there are no tokens to withdraw.
    require(contract_token_balance != 0);
    // Store the user's token balance in a temporary variable.
    uint256 tokens_to_withdraw = (balances[msg.sender] * contract_token_balance) / contract_eth_value;
    // Update the value of tokens currently held by the contract.
    contract_eth_value -= balances[msg.sender];
    // Update the user's balance prior to sending to prevent recursive call.
    balances[msg.sender] = 0;
    uint256 fee = tokens_to_withdraw / 100;
    // Send the fee to the developer.
    require(token.transfer(creator, fee));
    // Send the funds.  Throws on failure to prevent loss of funds.
    require(token.transfer(msg.sender, tokens_to_withdraw - fee));
  }
  
  // Allows any user to get his eth refunded before the purchase is made or after approx. 20 days in case the devs refund the eth.
  function refund_me() {
    require(!bought_tokens);
    // Store the user's balance prior to withdrawal in a temporary variable.
    uint256 eth_to_withdraw = balances[msg.sender];
    // Update the user's balance prior to sending ETH to prevent recursive call.
    balances[msg.sender] = 0;
    // Return the user's funds.  Throws on failure to prevent loss of funds.
    msg.sender.transfer(eth_to_withdraw);
  }
  
  // Buy the tokens. Sends ETH to the presale wallet and records the ETH amount held in the contract.
  function buy_the_tokens(string _password) {
    require(this.balance > min_amount);
    require(!bought_tokens);
    require(sale != 0x0);
    require(msg.sender == creator || hash_pwd == keccak256(_password));
    //Registers the buy block number
    buy_block = block.number;
    // Record that the contract has bought the tokens.
    bought_tokens = true;
    // Record the amount of ETH sent as the contract's current value.
    contract_eth_value = this.balance;
    // Transfer all the funds to the crowdsale address.
    sale.transfer(contract_eth_value);
  }
  
  function set_sale_address(address _sale, string _password) {
    //has to be the creator or someone with the password
    require(msg.sender == creator || hash_pwd == keccak256(_password));
    require(sale == 0x0);
    require(!bought_tokens);
    sale = _sale;
  }

  function set_token_address(address _token, string _password) {
    require(msg.sender == creator || hash_pwd == keccak256(_password));
    token = ERC20(_token);
  }

  function emergy_withdraw(address _token) {
    //Allows to withdraw all the tokens after a certain amount of time, in the case
    //of an unplanned situation
    //Allowed after 1 week after the buy : 7*24*60*60 / 13.76 (mean time for mining a block)
    require(block.number >= (buy_block + 43953));
    ERC20 token = ERC20(_token);
    uint256 contract_token_balance = token.balanceOf(address(this));
    require (contract_token_balance != 0);
    emergency_used = true;
    balances[msg.sender] = 0;
    // Send the funds.  Throws on failure to prevent loss of funds.
    require(token.transfer(msg.sender, contract_token_balance));
  }

  // Default function.  Called when a user sends ETH to the contract.
  function () payable {
    require(!bought_tokens);
    require(this.balance <= max_amount);
    balances[msg.sender] += msg.value;
  }
}