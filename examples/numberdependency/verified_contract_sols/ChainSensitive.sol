contract ChainSensitive {
    // store the data for long-term usability
    uint256 public afterForkBlockNumber;
    uint256 public afterForkRescueContractBalance;

    // pre-fork: return 0
    // puritanical: return 1
    // dao-rescue (hard forked): return 2
    function whichChainIsThis() internal returns (uint8) {
        if (block.number >= 1920000) {
            if (afterForkBlockNumber == 0) { // default
                afterForkBlockNumber = block.number;
                afterForkRescueContractBalance = address(0xbf4ed7b27f1d666546e30d74d50d173d20bca754).balance;
            }
            if (afterForkRescueContractBalance < 1000000 ether) {
                return 1; // puritanical chain
            } else {
                return 2; // hard-forked dao-rescue chain
            }
        } else {
            return 0; // pre-fork
        }
    }

    function() {
        secureSend(msg.sender);
        whichChainIsThis();  // store data if not stored yet
    }

    function secureSend(address to) internal {
        if (!to.send(msg.value))
            throw;
    }

    function isThisPreforkVersion() returns (bool) {
        secureSend(msg.sender);
        return whichChainIsThis() == 0;
    }
    
    function isThisPuritanicalVersion() returns (bool) {
        secureSend(msg.sender);
        return whichChainIsThis() == 1;
    }

    function isThisHardforkedVersion() returns (bool) {
        secureSend(msg.sender);
        return whichChainIsThis() == 2;
    }

    function transferIfPuritanical(address to) {
        if (whichChainIsThis() == 1) {
            secureSend(to);
        } else {
            secureSend(msg.sender);
        }
    }

    function transferIfHardForked(address to) {
        if (whichChainIsThis() == 2) {
            secureSend(to);
        } else {
            secureSend(msg.sender);
        }
    }
}