pragma solidity ^0.4.11;

contract BLOCKCHAIN_DEPOSIT_BETA {
	
	/* CONTRACT SETUP */

	uint constant PAYOUT_INTERVAL = 1 days;

	/* NB: Solidity doesn't support fixed or floats yet, so we use promille instead of percent */	
	uint constant DEPOSITOR_INTEREST= 10;
	uint constant INTEREST_DENOMINATOR = 1000;

	/* DATA TYPES */

	/* the payout happend */
	event Payout(uint paidPeriods, uint depositors);
	
	/* Depositor struct: describes a single Depositor */
	struct Depositor
	{	
		address etherAddress;
		uint deposit;
		uint depositTime;
	}


	/* the time of last payout */
	uint private contract_latestPayoutTime;

	/* Array of depositors */
	Depositor[] private contract_depositors;

	
	/* PUBLIC FUNCTIONS */

	/* contract constructor */
	function BLOCKCHAIN_DEPOSIT_BETA() 
	{
		contract_latestPayoutTime = now;		
	}

	/* fallback function: called when the contract received plain ether */
	function() payable
	{
		addDepositor();
	}

	function Make_Deposit() payable
	{
		addDepositor();	
	}

	function status() constant returns (uint deposit_fund_sum, uint depositorsCount, uint unpaidTime, uint unpaidIntervals)
	{
		deposit_fund_sum = this.balance;
		depositorsCount = contract_depositors.length;
		unpaidTime = now - contract_latestPayoutTime;
		unpaidIntervals = unpaidTime / PAYOUT_INTERVAL;
	}


	/* checks if it's time to make payouts. if so, send the ether */
	function performPayouts()
	{
		uint paidPeriods = 0;
		uint depositorsDepositPayout;

		while(contract_latestPayoutTime + PAYOUT_INTERVAL < now)
		{						
			uint idx;

			/* pay the depositors  */
			/* we use reverse iteration here */
			for (idx = contract_depositors.length; idx-- > 0; )
			{
				if(contract_depositors[idx].depositTime > contract_latestPayoutTime + PAYOUT_INTERVAL)
					continue;
				uint payout = (contract_depositors[idx].deposit * DEPOSITOR_INTEREST) / INTEREST_DENOMINATOR;
				if(!contract_depositors[idx].etherAddress.send(payout))
					throw;
				depositorsDepositPayout += payout;	
			}
			
			/* save the latest paid time */
			contract_latestPayoutTime += PAYOUT_INTERVAL;
			paidPeriods++;
		}
			
		/* emit the Payout event */
		Payout(paidPeriods, depositorsDepositPayout);
	}

	/* PRIVATE FUNCTIONS */
	function addDepositor() private 
	{
		contract_depositors.push(Depositor(msg.sender, msg.value, now));
	}


}