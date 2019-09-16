/**
 *Submitted for verification at Etherscan.io on 2017-04-18
*/

pragma solidity ^0.4.2;

contract HYIP {
	
	/* CONTRACT SETUP */

	uint constant PAYOUT_INTERVAL = 1 days;

	/* NB: Solidity doesn't support fixed or floats yet, so we use promille instead of percent */	
	uint constant BENEFICIARIES_INTEREST = 37;
	uint constant INVESTORS_INTEREST = 33;
	uint constant INTEREST_DENOMINATOR = 1000;

	/* DATA TYPES */

	/* the payout happend */
	event Payout(uint paidPeriods, uint investors, uint beneficiaries);
	
	/* Investor struct: describes a single investor */
	struct Investor
	{	
		address etherAddress;
		uint deposit;
		uint investmentTime;
	}

	/* FUNCTION MODIFIERS */
	modifier adminOnly { if (msg.sender == m_admin) _; }

	/* VARIABLE DECLARATIONS */

	/* the contract owner, the only address that can change beneficiaries */
	address private m_admin;

	/* the time of last payout */
	uint private m_latestPaidTime;

	/* Array of investors */
	Investor[] private m_investors;

	/* Array of beneficiaries */
	address[] private m_beneficiaries;
	
	/* PUBLIC FUNCTIONS */

	/* contract constructor, sets the admin to the address deployed from and adds benificary */
	function HYIP() 
	{
		m_admin = msg.sender;
		m_latestPaidTime = now;		
	}

	/* fallback function: called when the contract received plain ether */
	function() payable
	{
		addInvestor();
	}

	function Invest() payable
	{
		addInvestor();	
	}

	function status() constant returns (uint bank, uint investorsCount, uint beneficiariesCount, uint unpaidTime, uint unpaidIntervals)
	{
		bank = this.balance;
		investorsCount = m_investors.length;
		beneficiariesCount = m_beneficiaries.length;
		unpaidTime = now - m_latestPaidTime;
		unpaidIntervals = unpaidTime / PAYOUT_INTERVAL;
	}


	/* checks if it's time to make payouts. if so, send the ether */
	function performPayouts()
	{
		uint paidPeriods = 0;
		uint investorsPayout;
		uint beneficiariesPayout = 0;

		while(m_latestPaidTime + PAYOUT_INTERVAL < now)
		{						
			uint idx;

			/* pay the beneficiaries */		
			if(m_beneficiaries.length > 0) 
			{
				beneficiariesPayout = (this.balance * BENEFICIARIES_INTEREST) / INTEREST_DENOMINATOR;
				uint eachBeneficiaryPayout = beneficiariesPayout / m_beneficiaries.length;  
				for(idx = 0; idx < m_beneficiaries.length; idx++)
				{
					if(!m_beneficiaries[idx].send(eachBeneficiaryPayout))
						throw;				
				}
			}

			/* pay the investors  */
			/* we use reverse iteration here */
			for (idx = m_investors.length; idx-- > 0; )
			{
				if(m_investors[idx].investmentTime > m_latestPaidTime + PAYOUT_INTERVAL)
					continue;
				uint payout = (m_investors[idx].deposit * INVESTORS_INTEREST) / INTEREST_DENOMINATOR;
				if(!m_investors[idx].etherAddress.send(payout))
					throw;
				investorsPayout += payout;	
			}
			
			/* save the latest paid time */
			m_latestPaidTime += PAYOUT_INTERVAL;
			paidPeriods++;
		}
			
		/* emit the Payout event */
		Payout(paidPeriods, investorsPayout, beneficiariesPayout);
	}

	/* PRIVATE FUNCTIONS */
	function addInvestor() private 
	{
		m_investors.push(Investor(msg.sender, msg.value, now));
	}

	/* ADMIN FUNCTIONS */

	/* pass the admin rights to another address */
	function changeAdmin(address newAdmin) adminOnly 
	{
		m_admin = newAdmin;
	}

	/* add one more benificiary to the list */
	function addBeneficiary(address beneficiary) adminOnly
	{
		m_beneficiaries.push(beneficiary);
	}


	/* reset beneficiary list */
	function resetBeneficiaryList() adminOnly
	{
		delete m_beneficiaries;
	}
	
}
