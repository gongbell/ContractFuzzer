pragma solidity 0.4.18;


// D.H.A.R.M.A. Initiative Swan Protocol
// The protocol must be executed at least once every 108 minutes
// Failure to do so releases the reward to the last executor
contract Protocol108 {
	// smart contract version
	uint public version = 1;

	// countdown timer reset value
	uint length = 6480;

	// last time protocol was executed
	uint offset;

	// last executor of the protocol
	address public executor;

	// number of times protocol was executed
	// zero value means protocol is in initialization state
	uint public cycle;

	// total value volume passed through
	uint public volume;

	// creates the protocol
	function Protocol108() public {
	}

	// initializes the protocol
	function initialize() public payable {
		// validate protocol state
		assert(cycle == 0);

		// update the protocol
		update();
	}

	// executes the protocol
	function execute() public payable {
		// validate protocol state
		assert(cycle > 0);
		assert(offset + length > now);

		// update the protocol
		update();
	}

	// withdraws the reward to the last executor
	function withdraw() public {
		// validate protocol state
		assert(cycle > 0);
		assert(offset + length <= now);

		// validate input(s)
		require(msg.sender == executor);

		// reset cycle count
		cycle = 0;

		// transfer the reward
		executor.transfer(this.balance);
	}

	// updates the protocol state by
	// updating offset, last executor and cycle count
	function update() private {
		// validate input(s)
		validate(msg.value);

		// update offset (last execution time)
		offset = now;

		// update last executor
		executor = msg.sender;

		// update cycle
		cycle++;

		// update total volume
		volume += msg.value;
	}

	// validates the input sequence of numbers
	// simplest impl (current): positive value
	// proper impl (consideration for future versions): 00..0481516234200..0-like values
	// where any number of leading/trailing zeroes allowed
	// calling this function as part of transaction returns true or throws an exception
	// calling this function as constant returns true or false
	function validate(uint sequence) public constant returns (bool) {
		// validate the sequence
		require(sequence > 0);

		// we won't get here if validation fails
		return true;
	}

	// number of seconds left until protocol terminates
	function countdown() public constant returns (uint) {
		// check if protocol is initialized
		if(cycle == 0) {
			// for uninitialized protocol its equal to length
			return length;
		}

		// for active/terminated protocol calculate the value
		uint n = now;

		// check for negative overflow
		if(offset + length > n) {
			// positive countdown
			return offset + length - n;
		}

		// zero or negative countdown
		return 0;
	}

	// the default payable function, performs one of
	// initialize(), execute() or withdraw() depending on protocol state
	function() public payable {
		if(cycle == 0) {
			// protocol not yet initialized, try to initialize
			initialize();
		}
		else if(offset + length > now) {
			// protocol is eligible for execution, execute
			execute();
		}
		else if(this.balance > 0) {
			// protocol has terminated, withdraw the reward
			withdraw();
		}
		else {
			// invalid protocol state
			revert();
		}
	}

}