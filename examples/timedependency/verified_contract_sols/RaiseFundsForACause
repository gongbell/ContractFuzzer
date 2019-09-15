pragma solidity ^0.4.17;

contract RaiseFundsForACause {

    // Base definitions
    address public owner;
    address public receiver;
    string public cause;
    uint256 public expirationInSeconds;
    bool public hasBeenClaimed;
    uint256 public timeStarted;

    // Dynamic data
    uint256 public minimumAmountRequired;
    uint256 public numPayments;
    uint256 public totalAmountRaised;
    mapping(address => uint256) donationData;

    function RaiseFundsForACause(address beneficiary, string message, uint256 secondsUntilExpiration)
        public
    {
        require(beneficiary != 0x0);
        require(secondsUntilExpiration > 0);

        owner = msg.sender;
        receiver = beneficiary;
        cause = message;
        expirationInSeconds = secondsUntilExpiration;
        hasBeenClaimed = false;

        minimumAmountRequired = 0;
        numPayments = 0;
        totalAmountRaised = 0;
        timeStarted = block.timestamp;
    }

    function ()
        public
    {
        throw;
    }

    function donate()
        public
        payable
    {
        require(msg.sender != receiver);
        require(block.timestamp < (timeStarted + expirationInSeconds));
        require(msg.value > 0);
        require(minimumAmountRequired != 0);
        require(hasBeenClaimed == false);

        assert(donationData[msg.sender] + msg.value >= donationData[msg.sender]);
        assert(totalAmountRaised + msg.value >= totalAmountRaised);
        assert(numPayments + 1 >= numPayments);

        donationData[msg.sender] += msg.value;
        totalAmountRaised += msg.value;
        numPayments += 1;
    }

    // Note: can only be set once
    function receiverSetAmountRequired(uint256 minimum)
        public
    {
        require(msg.sender == receiver);
        require(minimumAmountRequired == 0);
        require(minimum > 0);

        minimumAmountRequired = minimum;
    }

    function receiverWithdraw()
        public
    {
        require(msg.sender == receiver);
        require(totalAmountRaised >= minimumAmountRequired);
        require(this.balance > 0);
        require(block.timestamp < (timeStarted + expirationInSeconds));
        require(hasBeenClaimed == false);

        hasBeenClaimed = true;
        receiver.transfer(this.balance);
        // Expecting transfer to throw on error
        // assert(this.balance == 0);
    }

    function withdraw()
        public
    {
        require(donationData[msg.sender] > 0);
        require(block.timestamp > (timeStarted + expirationInSeconds));
        require(hasBeenClaimed == false);

        var value = donationData[msg.sender];
        donationData[msg.sender] = 0;
        msg.sender.transfer(value);
        // Expecting transfer to throw on error
        // assert(donationData[donor] == 0);
    }

    function currentTotalExcess()
        public
        constant returns (uint256)
    {
        if (totalAmountRaised > minimumAmountRequired) {
            return totalAmountRaised - minimumAmountRequired;
        }
        else {
            return 0;
        }
    }

    function expirationTimestamp()
        public
        constant returns (uint256)
    {
        assert((timeStarted + expirationInSeconds) >= timeStarted);
        return (timeStarted + expirationInSeconds);
    }
}