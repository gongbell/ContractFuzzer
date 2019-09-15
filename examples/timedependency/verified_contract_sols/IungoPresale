pragma solidity 0.4.15;

contract Owned {
    address public owner;

    function Owned() {
        owner = msg.sender;
    }

    modifier onlyOwner() {
        if (msg.sender != owner) {
            throw;
        }
        _;
    }
}

contract IungoPresale is Owned {
    // Total funding amount
    uint256 public totalFunding;

    // Minimum and maximum amounts per transaction for participants
    uint256 public constant MIN_AMOUNT = 500 finney;
    uint256 public constant MAX_AMOUNT = 50 ether;

    // Minimum and maximum goals of the presale
    uint256 public constant PRESALE_MINIMUM_FUNDING = 120 ether;
    uint256 public constant PRESALE_MAXIMUM_FUNDING = 1100 ether;

    // Public presale period
    // Starts Saturday, October 28, 2017 12:00:00 AM GMT
    // Ends Sunday, November 19, 2017 12:00:00 AM GTM
    uint256 public constant PRESALE_START_DATE = 1509148800;
    uint256 public constant PRESALE_END_DATE = 1511049600;

    // Owner can clawback after a date in the future, so no ethers remain
    // trapped in the contract. This will only be relevant if the
    // minimum funding level is not reached
    // Jan 01 2018 @ 12:00pm (UTC) 2018-01-01T12:00:00+00:00 in ISO 8601
    uint256 public constant OWNER_CLAWBACK_DATE = 1514808000;

    /// @notice Keep track of all participants contributions, including both the
    ///         preallocation and public phases
    /// @dev Name complies with ERC20 token standard, etherscan for example will recognize
    ///      this and show the balances of the address
    mapping (address => uint256) public balanceOf;

    /// @notice Log an event for each funding contributed during the public phase
    /// @notice Events are not logged when the constructor is being executed during
    ///         deployment, so the preallocations will not be logged
    event LogParticipation(address indexed sender, uint256 value, uint256 timestamp);

    /// @notice A participant sends a contribution to the contract's address
    ///         between the PRESALE_STATE_DATE and the PRESALE_END_DATE
    /// @notice Only contributions between the MIN_AMOUNT and
    ///         MAX_AMOUNT are accepted. Otherwise the transaction
    ///         is rejected and contributed amount is returned to the participant's
    ///         account
    /// @notice A participant's contribution will be rejected if the presale
    ///         has been funded to the maximum amount
    function () payable {
        // A participant cannot send funds before the presale start date
        if (now < PRESALE_START_DATE) throw;
        // A participant cannot send funds after the presale end date
        if (now > PRESALE_END_DATE) throw;
        // A participant cannot send less than the minimum amount
        if (msg.value < MIN_AMOUNT) throw;
        // A participant cannot send more than the maximum amount
        if (msg.value > MAX_AMOUNT) throw;
        // A participant cannot send funds if the presale has been reached the maximum
        // funding amount
        if (safeIncrement(totalFunding, msg.value) > PRESALE_MAXIMUM_FUNDING) throw;
        // Register the participant's contribution
        addBalance(msg.sender, msg.value);
    }

    /// @notice The owner can withdraw ethers already during presale,
    ///         only if the minimum funding level has been reached
    function ownerWithdraw(uint256 value) external onlyOwner {
        // The owner cannot withdraw if the presale did not reach the minimum funding amount
        if (totalFunding < PRESALE_MINIMUM_FUNDING) throw;
        // Withdraw the amount requested
        if (!owner.send(value)) throw;
    }

    /// @notice The participant will need to withdraw their funds from this contract if
    ///         the presale has not achieved the minimum funding level
    function participantWithdrawIfMinimumFundingNotReached(uint256 value) external {
        // Participant cannot withdraw before the presale ends
        if (now <= PRESALE_END_DATE) throw;
        // Participant cannot withdraw if the minimum funding amount has been reached
        if (totalFunding >= PRESALE_MINIMUM_FUNDING) throw;
        // Participant can only withdraw an amount up to their contributed balance
        if (balanceOf[msg.sender] < value) throw;
        // Participant's balance is reduced by the claimed amount.
        balanceOf[msg.sender] = safeDecrement(balanceOf[msg.sender], value);
        // Send ethers back to the participant's account
        if (!msg.sender.send(value)) throw;
    }

    /// @notice The owner can clawback any ethers after a date in the future, so no
    ///         ethers remain trapped in this contract. This will only be relevant
    ///         if the minimum funding level is not reached
    function ownerClawback() external onlyOwner {
        // The owner cannot withdraw before the clawback date
        if (now < OWNER_CLAWBACK_DATE) throw;
        // Send remaining funds back to the owner
        if (!owner.send(this.balance)) throw;
    }

    /// @dev Keep track of participants contributions and the total funding amount
    function addBalance(address participant, uint256 value) private {
        // Participant's balance is increased by the sent amount
        balanceOf[participant] = safeIncrement(balanceOf[participant], value);
        // Keep track of the total funding amount
        totalFunding = safeIncrement(totalFunding, value);
        // Log an event of the participant's contribution
        LogParticipation(participant, value, now);
    }

    /// @dev Add a number to a base value. Detect overflows by checking the result is larger
    ///      than the original base value.
    function safeIncrement(uint256 base, uint256 increment) private constant returns (uint256) {
        uint256 result = base + increment;
        if (result < base) throw;
        return result;
    }

    /// @dev Subtract a number from a base value. Detect underflows by checking that the result
    ///      is smaller than the original base value
    function safeDecrement(uint256 base, uint256 increment) private constant returns (uint256) {
        uint256 result = base - increment;
        if (result > base) throw;
        return result;
    }
}