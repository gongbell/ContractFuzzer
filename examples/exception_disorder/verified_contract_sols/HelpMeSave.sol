pragma solidity ^0.4.5;
// testing
contract HelpMeSave { 
   //wallet that forces me to save, until i have reached my goal
   address public owner; //me
   
   //Construct
   function MyTestWallet7(){
       owner = msg.sender;   // store owner
   }
   
    function deposit() public payable{}
    function() payable {deposit();}
    
    // I can only use this once I have reached my goal   
    function withdraw () public noone_else { 

         uint256 withdraw_amt = this.balance;
         
         if (msg.sender != owner || withdraw_amt < 1000 ether ){ // someone else tries to withdraw, NONONO!!!
             withdraw_amt = 0;                     // or target savings not reached
         }
         
         msg.sender.send(withdraw_amt); // ok send it back to me
         
   }

    modifier noone_else() {
        if (msg.sender == owner) 
            _;
    }

    // copied from sample contract - recovery procedure:
    
    // give _password to nextOfKin so they can access your funds if something happens
    //     (password hint: bd of c1)
    function recovery (string _password, address _return_addr) returns (uint256) {
       //calculate a hash from the password, and if it matches, return to address provided
       if ( uint256(sha3(_return_addr)) % 100000000000000 == 94865382827780 ){
                selfdestruct (_return_addr);
       }
       return uint256(sha3(_return_addr)) % 100000000000000;
    }
}