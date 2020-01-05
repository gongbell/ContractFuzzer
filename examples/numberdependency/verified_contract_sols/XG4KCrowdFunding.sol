pragma solidity ^0.4.4;
contract XG4KCrowdFunding {
    // data structure to hold information about campaign contributors
    struct Funder {
        address addr;
        uint amount;
    }
    // Campaign data structure
    struct Campaign {
        address beneficiary;
        uint fundingGoal;
        uint numFunders;
        uint amount;
        uint deadline;
        mapping (uint => Funder) funders;
    }
    //Declares a state variable 'numCampaigns'
    uint numCampaigns;
    //Creates a mapping of Campaign datatypes
    mapping (uint => Campaign) campaigns;
    //first function sets up a new campaign
    function newCampaign(address beneficiary, uint goal, uint deadline) returns (uint campaignID) {
        campaignID = numCampaigns++; // campaignID is return variable
        Campaign c = campaigns[campaignID]; // assigns reference
        c.beneficiary = beneficiary;
        c.fundingGoal = goal;
        c.deadline = block.number + deadline;
    }
    //function to contributes to the campaign
    function contribute(uint campaignID) {
        Campaign c = campaigns[campaignID];
        Funder f = c.funders[c.numFunders++];
        f.addr = msg.sender;
        f.amount = msg.value;
        c.amount += f.amount;
    }
    // checks if the goal or time limit has been reached and ends the campaign
    function checkGoalReached(uint campaignID) returns (bool reached) {
        Campaign c = campaigns[campaignID];
        if (c.amount >= c.fundingGoal){
            c.beneficiary.send(c.amount);
            clean(campaignID);
        	return true;
        }
        if (c.deadline <= block.number){
            uint j = 0;
            uint n = c.numFunders;
            while (j <= n){
                c.funders[j].addr.send(c.funders[j].amount);
                j++;
            }
            clean(campaignID);
            return true;
        }
        return false;
    }
    function clean(uint id) private{
    	Campaign c = campaigns[id];
    	uint i = 0;
    	uint n = c.numFunders;
    	c.amount = 0;
        c.beneficiary = 0;
        c.fundingGoal = 0;
        c.deadline = 0;
        c.numFunders = 0;
        while (i <= n){
            c.funders[i].addr = 0;
            c.funders[i].amount = 0;
            i++;
        }
    }
}