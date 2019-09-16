/**
 *Submitted for verification at Etherscan.io on 2018-01-08
*/

pragma solidity ^0.4.18;


contract BlockFoodPreSale {

    enum ApplicationState {Unset, Pending, Rejected, Accepted, Refunded}

    struct Application {
        uint contribution;
        string id;
        ApplicationState state;
    }

    struct Applicant {
        address applicant;
        string id;
    }

    /*
        Set by constructor
    */
    address public owner;
    address public target;
    uint public endDate;
    uint public minContribution;
    uint public minCap;
    uint public maxCap;

    /*
        Set by functions
    */
    mapping(address => Application) public applications;
    Applicant[] public applicants;
    uint public contributionPending;
    uint public contributionRejected;
    uint public contributionAccepted;
    uint public withdrawn;

    /*
        Events
    */
    event PendingApplication(address applicant, uint contribution, string id);
    event RejectedApplication(address applicant, uint contribution, string id);
    event AcceptedApplication(address applicant, uint contribution, string id);
    event Withdrawn(address target, uint amount);
    event Refund(address target, uint amount);
    event ContractUpdate(address owner, address target, uint minContribution, uint minCap, uint maxCap);

    /*
        Modifiers
    */
    modifier onlyBeforeEnd() {
        require(now <= endDate);
        _;
    }

    modifier onlyMoreThanMinContribution() {
        require(msg.value >= minContribution);
        _;
    }

    modifier onlyMaxCapNotReached() {
        require((contributionAccepted + msg.value) <= maxCap);
        _;
    }

    modifier onlyOwner () {
        require(msg.sender == owner);
        _;
    }

    modifier onlyNewApplicant () {
        require(applications[msg.sender].state == ApplicationState.Unset);
        _;
    }

    modifier onlyPendingApplication(address applicant) {
        require(applications[applicant].contribution > 0);
        require(applications[applicant].state == ApplicationState.Pending);
        _;
    }

    modifier onlyMinCapReached() {
        require(contributionAccepted >= minCap);
        _;
    }

    modifier onlyNotWithdrawn(uint amount) {
        require(withdrawn + amount <= contributionAccepted);
        _;
    }

    modifier onlyFailedPreSale() {
        require(now >= endDate);
        require(contributionAccepted + contributionPending < minCap);
        _;
    }

    modifier onlyAcceptedApplication(address applicant) {
        require(applications[applicant].state == ApplicationState.Accepted);
        _;
    }

    modifier onlyAfterTwoMonthsAfterTheEnd() {
        require(now > (endDate + 60 days));
        _;
    }

    modifier sendContractUpdateEvent() {
        _;
        ContractUpdate(owner, target, minContribution, minCap, maxCap);
    }

    /*
        Constructor
    */
    function BlockFoodPreSale(
        address target_,
        uint endDate_,
        uint minContribution_,
        uint minCap_,
        uint maxCap_
    )
    public
    {
        owner = msg.sender;

        target = target_;
        endDate = endDate_;
        minContribution = minContribution_;
        minCap = minCap_;
        maxCap = maxCap_;
    }

    /*
       Public functions
    */

    function apply(string id)
    payable
    public
    onlyBeforeEnd
    onlyMoreThanMinContribution
    onlyMaxCapNotReached
    onlyNewApplicant
    {
        applications[msg.sender] = Application(msg.value, id, ApplicationState.Pending);
        applicants.push(Applicant(msg.sender, id));
        contributionPending += msg.value;
        PendingApplication(msg.sender, msg.value, id);
    }

    function refund()
    public
    onlyFailedPreSale
    onlyAcceptedApplication(msg.sender)
    {
        applications[msg.sender].state = ApplicationState.Refunded;
        msg.sender.transfer(applications[msg.sender].contribution);
        Refund(msg.sender, applications[msg.sender].contribution);
    }

    /*
        Restricted functions (owner only)
    */

    function reject(address applicant)
    public
    onlyOwner
    onlyPendingApplication(applicant)
    {
        applications[applicant].state = ApplicationState.Rejected;

        // protection against function reentry on an overriden transfer() function
        uint contribution = applications[applicant].contribution;
        applications[applicant].contribution = 0;
        applicant.transfer(contribution);

        contributionPending -= contribution;
        contributionRejected += contribution;

        RejectedApplication(applicant, contribution, applications[applicant].id);
    }

    function accept(address applicant)
    public
    onlyOwner
    onlyPendingApplication(applicant)
    {
        applications[applicant].state = ApplicationState.Accepted;

        contributionPending -= applications[applicant].contribution;
        contributionAccepted += applications[applicant].contribution;

        AcceptedApplication(applicant, applications[applicant].contribution, applications[applicant].id);
    }

    function withdraw(uint amount)
    public
    onlyOwner
    onlyMinCapReached
    onlyNotWithdrawn(amount)
    {
        withdrawn += amount;
        target.transfer(amount);
        Withdrawn(target, amount);
    }

    /*
        Views
    */

    function getApplicantsLength()
    view
    public
    returns (uint)
    {
        return applicants.length;
    }

    function getMaximumContributionPossible()
    view
    public
    returns (uint)
    {
        return maxCap - contributionAccepted;
    }

    /*
        Maintenance functions
    */

    function failsafe()
    public
    onlyOwner
    onlyAfterTwoMonthsAfterTheEnd
    {
        target.transfer(this.balance);
    }

    function changeOwner(address owner_)
    public
    onlyOwner
    sendContractUpdateEvent
    {
        owner = owner_;
    }

    function changeTarget(address target_)
    public
    onlyOwner
    sendContractUpdateEvent
    {
        target = target_;
    }

    function changeMinCap(uint minCap_)
    public
    onlyOwner
    sendContractUpdateEvent
    {
        minCap = minCap_;
    }

    function changeMaxCap(uint maxCap_)
    public
    onlyOwner
    sendContractUpdateEvent
    {
        maxCap = maxCap_;
    }

    function changeMinContribution(uint minContribution_)
    public
    onlyOwner
    sendContractUpdateEvent
    {
        minContribution = minContribution_;
    }

}
