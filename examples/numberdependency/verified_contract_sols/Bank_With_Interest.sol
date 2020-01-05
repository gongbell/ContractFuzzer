pragma solidity ^0.4.5;

contract Bank_With_Interest {
    //
    ////////////////////////////////////////////////////////////////
    //
    //  A term deposit bank that pays interest on withdrawal  
    //
    //  v0.01 beta, use AT OWN RISK - I am not refunding any lost ether!
    //  And check the code before depositing anything.
    //
    //  How to use: 1) transfer min. 250 ether using the deposit() function (5 ether deposit fee per deposit)
    //                 note: minimum ether amount and deposit fee can change, so check public variables
    //                   - minimum_payment
    //                   - deposit_fee
    //                 before depositing!
    //
    //              2) withdraw after 5 days, receive up to 10% interest (1 additional ether for every 10 ether withdrawal)
    //                 note: to get most of the interest paid out, withdraw in lots of 10 ether...
    //
    ///////////////////////////////////////////////////////////////////
    //
    //  Now you may ask - where do the extra ether come from? :)
    //  The answer is simple: From people who dont follow the instructions properly! 
    //                        And there are usually plenty of those...
    //
    //  Common pitfalls:
    //   e.g. - deposit to the fallback function instead of the proper deposit() function
    //        - withdraw MORE than 10 ether AT A TIME ... (that just means you get less interest paid out)
    //        -  or you wait too long after your term deposit ended, and other people have drained the interest pool.
    //
    //  You can always check the availbale interest amount using get_available_interest_amount () before you withdraw
    //  And be quick - everyone gets the same 30850 block term deposit time (~5 days) until they can withdraw.
    //
    //  Also FYI: The bank cannot remove money from the interest pool until the end of the contract life.
    //        And make sure you withdraw your balances before the end of the contract life also.
    //        Check public contract_alive_until_this_block variable to find out when the contract life ends.
    //           Initial end date is block #3000000 (~Mid Jan 2017), but the bank can extend that life.
    //           Note: the bank can only EXTEND that end date, not shorten it.
    //
    // And here we go - happy reading:
    //
    // store of all account balances
    mapping(address => uint256) balances;
    mapping(address => uint256) term_deposit_end_block; // store per address the minimum time for the term deposit
                                                         //
    address public thebank; // the bank
    
    uint256 public minimum_payment; // minimum deposits
    uint256 public deposit_fee;     // fee for deposits
    
    uint256 public contract_alive_until_this_block;
    
    bool public any_customer_payments_yet = false; // the first cutomer payment will make this true and 
                                        // the contract cannot be deleted any more before end of life
    
    function Bank_With_Interest() { // create the contract
        thebank = msg.sender;  
        minimum_payment = 250 ether;
        deposit_fee = 5 ether;
        contract_alive_until_this_block = 3000000; // around 2 months from now (mid Jan 2017)
                                                   // --> can be extended but not brought forward
        //
        // bank cannot touch remaining interest balance until the contract has reached end of life.
        term_deposit_end_block[thebank] = 0;// contract_alive_until_this_block;
        //
    }
    
   //////////////////////////////////////////////////////////////////////////////////////////
    // deposit ether into term-deposit account
    //////////////////////////////////////////////////////////////////////////////////////////
    function deposit() payable {
        //
        if (msg.value < minimum_payment) throw; // minimum deposit is at least minimum_payment.
        //
        // no fee for first payment (if the customers's balance is 0)
        if (balances[msg.sender] == 0) deposit_fee = 0 ether;  
        //
        if ( msg.sender == thebank ){ // thebank is depositing into bank/interest account, without fee
            balances[thebank] += msg.value;
        }
        else { // customer deposit
            any_customer_payments_yet = true; // cannot remove contract any more until end of life
            balances[msg.sender] += msg.value - deposit_fee;  // credit the sender's account
            balances[thebank] += deposit_fee; // difference (fee) to be credited to thebank
            term_deposit_end_block[msg.sender] = block.number + 30850; //  approx 5 days ( 5 * 86400/14 ); 
        }
        //
    }
    
    //////////////////////////////////////////////////////////////////////////////////////////
    // withdraw from account, with 10 ether interest  (after term deposit end)
    //////////////////////////////////////////////////////////////////////////////////////////
    //
    function withdraw(uint256 withdraw_amount) {
        //
        if (withdraw_amount < 10 ether) throw; // minimum withdraw amount is 10 ether
        if ( withdraw_amount > balances[msg.sender]  ) throw; // cannot withdraw more than in customer balance
        if (block.number < term_deposit_end_block[msg.sender] ) throw; // cannot withdraw until the term deposit has ended
        // Note: thebank/interest account cannot be withdrawed from until contract end-of life.
        //       thebank's term-deposit end block is the same as contract_alive_until_this_block
        //
        uint256 interest = 1 ether;  // 1 ether interest paid at time of withdrawal
        //
        if (msg.sender == thebank){ // but no interest for thebank (who can't withdraw until block contract_alive_until_this_block anyways)
            interest = 0 ether;
        }
        //                          
        if (interest > balances[thebank])   // cant pay more interest than available in the thebank/bank
            interest = balances[thebank];  // so send whatever is left anyways
        //
        //
        balances[thebank] -= interest;  // reduce thebank balance, and send bonus to customer
        balances[msg.sender] -= withdraw_amount;
        //
        if (!msg.sender.send(withdraw_amount)) throw;  // send withdraw amount, but check for error to roll back if needed
        if (!msg.sender.send(interest)) throw;         // send interest amount, but check for error to roll back if needed
        //
    }
    
    ////////////////////////////////////////////////////////////////////////////
    // HELPER FUNCTIONS
    ////////////////////////////////////////////////////////////////////////////
    
    // set minimum deposit limits
    function set_minimum_payment(uint256 new_limit) {
        if ( msg.sender == thebank ){
            minimum_payment = new_limit;
        }
    }
    //
    // change deposit fee
    function set_deposit_fee (uint256 new_fee) {
        if ( msg.sender == thebank ){
            deposit_fee = new_fee;
        }
    }
    
    // find out how much money is available for interest payments
    function get_available_interest_amount () returns (uint256) {
        return balances[thebank];
    }
    // find out what the end date of the customers term deposit is
    function get_term_deposit_end_date () returns (uint256) {
        return term_deposit_end_block[msg.sender];
    }    
    // find out how much money is available for interest payments
    function get_balance () returns (uint256) {
        return balances[msg.sender];
    }
    //
    ////////////////////////////////////////////////////////////////
    // this bank won't live forever, so this will handle the exit (or extend its life)
    ////////////////////////////////////////////////////////////
	//
    function extend_life_of_contract (uint256 newblock){
        if ( msg.sender != thebank || newblock < contract_alive_until_this_block ) throw;
        // can only extend
        contract_alive_until_this_block = newblock; 
        // lock thebank/interest account until new end date
        term_deposit_end_block[thebank] = contract_alive_until_this_block;
    }
    //
    // the self destruct after the final block number has been reached (or immediately if there havent been any customer payments yet)
    function close_bank(){
        if (contract_alive_until_this_block < block.number || !any_customer_payments_yet)
            selfdestruct(thebank); 
            // any funds still remaining within the bank will be sent to the creator
            // --> bank customers have to make sure they withdraw their $$$ before the final block.
    }
    ////////////////////////////////////////////////////////////////
    // fallback function
    ////////////////////////////////////////////////////////////
    function () payable { // any unidentified payments (that didnt call the deposit function) 
                          // go into the standard interest account of the bank
                          // and become available for interest withdrawal by bank users
        balances[thebank] += msg.value;
    }
}