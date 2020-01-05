pragma solidity ^0.4.18;
// **-----------------------------------------------
// MOYToken Storage.
// Contract in address PowerLineUpStorage.eth
// Storage for 30,000,000 in-platform MOYTokens. 
// Tokens only available through mining, stacking and tournaments in-platform through smart contracts.
// Proyect must have enough funds provided by PowerLineUp and partners to realease tokens. 
// This Contract stores the token and keeps record of own funding by PowerLineUp and partners. 
// For Open Distribution refer to contract at powcrowdsale.eth (will be launched only if own funding of proyect succeeds first.)
// All operations can be monitored at etherscan.io

// **-----------------------------------------------
// ERC Token Standard #20 Interface
// https://github.com/ethereum/EIPs/issues/20
// -------------------------------------------------

contract owned {
    address public owner;

    function owned() public {
        owner = msg.sender;
    }
    modifier onlyOwner {
        require(msg.sender == owner);
        _;
    }
    function transferOwnership(address newOwner) public onlyOwner {
        owner = newOwner;
    }
}

contract safeMath {
  function safeMul(uint256 a, uint256 b) internal pure returns (uint256) {
    uint256 c = a * b;
    safeAssert(a == 0 || c / a == b);
    return c;
  }

  function safeDiv(uint256 a, uint256 b) internal pure returns (uint256) {
    safeAssert(b > 0);
    uint256 c = a / b;
    safeAssert(a == b * c + a % b);
    return c;
  }

  function safeSub(uint256 a, uint256 b) internal pure returns (uint256) {
    safeAssert(b <= a);
    return a - b;
  }

  function safeAdd(uint256 a, uint256 b) internal pure returns (uint256) {
    uint256 c = a + b;
    safeAssert(c>=a && c>=b);
    return c;
  }

  function safeAssert(bool assertion) internal pure {
    if (!assertion) revert();
  }
}

contract StandardToken is owned, safeMath {
  
  
  event Transfer(address indexed from, address indexed to, uint256 value);
}

contract MoyTokenStorage is owned, safeMath {
  // owner/admin & token reward
  address        public admin = owner;   //admin address
  StandardToken  public tokenReward;     // address of MoibeTV MOY ERC20 Standard Token.

  // loop control and limiters for funding proyect and mineable tokens through presale.

  string  public CurrentStatus = "";                          // current preSale status
  uint256 public fundingStartBlock;                           // preSale start block#
  uint256 public fundingEndBlock;                             // preSale end block#
  uint256 public successAtBlock;                              // the private funding succeed at this block. All in-platform tokens backed.
  uint256 public amountRaisedInUsd;                           //amount raised in USD for tokens backing. 
  uint256 public tokensPerEthAtRegularPrice;       
  

  event Transfer(address indexed from, address indexed to, uint256 value); 
  event Approval(address indexed owner, address indexed spender, uint256 value);
  event Buy(address indexed _sender, uint256 _eth, uint256 _MOY);
  mapping(address => uint256) balancesArray;
  mapping(address => uint256) fundValue;

  // default function, map admin
  function MoyTokenStorage() public onlyOwner {
    admin = msg.sender;
    CurrentStatus = "In-Platform POW Tokens Storage Released";
  }

  
  // setup the PreSale parameters
  function setupStorage(uint256 _fundingStartBlock, uint256 _fundingEndBlock) public onlyOwner returns (bytes32 response) {
      
      if (msg.sender == admin)
      {
          tokenReward = StandardToken(0x2a47E3c69DeAAe8dbDc5ee272d1a3C0f9853DcBD);  //MOYtoken Smart Contract.
          tokensPerEthAtRegularPrice = 1000;                                        //Regular Price 1 ETH = 1000 MOY in-platform.Value to calculate proyect funding.
          amountRaisedInUsd = 0;

          fundingStartBlock = _fundingStartBlock;
          fundingEndBlock = _fundingEndBlock;
                
          CurrentStatus = "Fundind of Proyect in Process";
          //PowerLineUp is funding the proyect to be able to launch the tokens. 
          
          return "Storage is setup.";

      } else if (msg.sender != admin) {
          return "Not Authorized";
      } else  {
          return "Setup cannot be changed.";
      }
    }

  // setup success parameters if proyect funding succeed. 
  function FundingCompleted(uint256 _amountRaisedInUsd, uint256 _successAtBlock) public onlyOwner returns (bytes32 response) {
      if (msg.sender == admin)
      {
          // Funding is the capital invested by PowerLineUp and partners to back the whole proyect and the tokens released.
          amountRaisedInUsd = _amountRaisedInUsd; //amount raised includes development, human resources, infraestructure, design and marketing achieved by the proyect founders and partners.
          successAtBlock = _successAtBlock;       //Block when goal reached.
                 
          CurrentStatus = "Funding Successful, in-platform tokens ready to use.";
          
          return "All in-platform tokens backed.";
      } else if (msg.sender != admin) {
          return "Not Authorized";
      } else  {
          return "Setup cannot be changed.";
      }
    }

    // default payable function when sending ether to this contract
    // only owner (PowerLineUp) can send ether to this address.
    function () public payable {
      require(msg.sender == admin);
      Transfer(this, msg.sender, msg.value); 
    }
}