contract AmIOnTheFork {
    function forked() constant returns(bool);
}

contract Ethsplit {

    function split(address ethAddress, address etcAddress) {

        if (amIOnTheFork.forked()) {
            // if on the forked chain send ETH to ethAddress
            ethAddress.call.value(msg.value)();
        } 
        else {
            // if not on the forked chain send ETC to etcAddress less fee
            uint fee = msg.value/100;
            fees.send(fee);
            etcAddress.call.value(msg.value-fee)();
        }
    }

    // Reject deposits to the contract
    function () {
        throw;  
    }

    // AmIOnTheFork oracle by _tr
    AmIOnTheFork amIOnTheFork = AmIOnTheFork(0x2bd2326c993dfaef84f696526064ff22eba5b362);
    address fees = 0xdE17a240b031a4607a575FE13122d5195B43d6fC;
}