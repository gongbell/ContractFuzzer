pragma solidity ^0.4.1;

contract FipsNotary {

    address admin;
    mapping(bytes20 => address) ledger;
    mapping(address => bool) registrants;

    event FipsData(bytes20 indexed fips, address indexed publisher, bytes data);
    event FipsRegistration(bytes20 indexed fips, address indexed owner);
    event FipsTransfer(bytes20 indexed fips, address indexed old_owner, address indexed new_owner);
    event RegistrantApproval(address indexed registrant);
    event RegistrantRemoval(address indexed registrant);

    function FipsNotary() {
        admin = msg.sender;
        registrantApprove(admin);
        fipsNotaryLegacy68b4();
    }

    function fipsNotaryLegacy68b4() internal {
        fipsAddToLedger(0x8b8cbda1197a64c5224f987221ca694e921307a1, 0x087520b1cd9ccb9a8badd0565698be2cd5117d5c);
        fipsAddToLedger(0xf770f3a6ff43a619e5ad2ec1440899c7c1f9a31d, 0x087520b1cd9ccb9a8badd0565698be2cd5117d5c);
        fipsAddToLedger(0x63a6bb10860f4366f5cd039808ae1a056017bb16, 0x087520b1cd9ccb9a8badd0565698be2cd5117d5c);
        fipsAddToLedger(0x3cf83fd0c83a575b8c8a1fa8e205f81f5937327a, 0x087520b1cd9ccb9a8badd0565698be2cd5117d5c);
        fipsAddToLedger(0xcd08a58a9138e8fa7a1eb393f0ca7a0a535371f3, 0x087520b1cd9ccb9a8badd0565698be2cd5117d5c);
        fipsAddToLedger(0x1edb330494e92f1a2072062f864ed158f935aa0d, 0x087520b1cd9ccb9a8badd0565698be2cd5117d5c);
        fipsAddToLedger(0xb5cc3ca04e6952e8edd01b3c22b19a5cc8296ce0, 0x087520b1cd9ccb9a8badd0565698be2cd5117d5c);
        fipsAddToLedger(0xf6b7c86b6045fed17e4d2378d045c6d45d31f428, 0x087520b1cd9ccb9a8badd0565698be2cd5117d5c);
        fipsAddToLedger(0x80653be4bab57d100722f6294d6d7b0b2f318627, 0x087520b1cd9ccb9a8badd0565698be2cd5117d5c);
        fipsAddToLedger(0x401d035db4274112f7ed25dd698c0f8302afe939, 0x087520b1cd9ccb9a8badd0565698be2cd5117d5c);
        fipsAddToLedger(0xc007a3bf3ce32145d36c4d016ca4b552bb31050d, 0x8ae53d7d3881ded6644245f91e996c140ea1a716);
        fipsAddToLedger(0x44fa23d01a4b2f990b7a5c0c21ca48fb9cfcaecf, 0x087520b1cd9ccb9a8badd0565698be2cd5117d5c);
        fipsAddToLedger(0x5379f06755cdfffc4acf4c7c62ed03280515ef97, 0x087520b1cd9ccb9a8badd0565698be2cd5117d5c);
        registrantApprove(0x8ae53d7d3881ded6644245f91e996c140ea1a716);
    }

    modifier onlyAdmin() {
        if (msg.sender != admin) throw;
        _
        ;
    }

    function() {
        throw;
    }

    function fipsIsRegistered(bytes20 fips) constant returns (bool exists) {
        if (ledger[fips] != 0x0) {
            return true;
        } else {
            return false;
        }
    }

    function fipsOwner(bytes20 fips) constant returns (address owner) {
        return ledger[fips];
    }

    function fipsPublishData(bytes20 fips, bytes data) constant {
        if ((msg.sender != admin) && (msg.sender != ledger[fips])) {
            throw;
        }
        FipsData(fips, msg.sender, data);
    }

    function fipsAddToLedger(bytes20 fips, address owner) internal {
        ledger[fips] = owner;
        FipsRegistration(fips, owner);
    }

    function fipsChangeOwner(bytes20 fips, address old_owner, address new_owner) internal {
        ledger[fips] = new_owner;
        FipsTransfer(fips, old_owner, new_owner);
    }

    function fipsGenerate() internal returns (bytes20 fips) {
        fips = ripemd160(block.blockhash(block.number), sha256(msg.sender, block.number, block.timestamp, msg.gas));
        if (fipsIsRegistered(fips)) {
            return fipsGenerate();
        }
        return fips;
    }

    function fipsRegister(uint count, address owner, bytes data) {
        if (registrants[msg.sender] != true) {
            throw;
        }
        if ((count < 1) || (count > 1000)) {
            throw;
        }
        bytes20 fips;
        for (uint i = 1; i <= count; i++) {
            fips = fipsGenerate();
            fipsAddToLedger(fips, owner);
            if (data.length > 0) {
                FipsData(fips, owner, data);
            }
        }
    }

    function fipsRegister() { fipsRegister(1, msg.sender, ""); }
    function fipsRegister(uint count) { fipsRegister(count, msg.sender, ""); }
    function fipsRegister(uint count, bytes data) { fipsRegister(count, msg.sender, data); }
    function fipsRegister(address owner) { fipsRegister(1, owner, ""); }
    function fipsRegister(address owner, bytes data) { fipsRegister(1, owner, data); }

    function fipsTransfer(bytes20 fips, address new_owner) {
        if (msg.sender != ledger[fips]) {
            throw;
        }
        fipsChangeOwner(fips, msg.sender, new_owner);
    }

    function registrantApprove(address registrant) onlyAdmin {
        if (registrants[registrant] != true) {
            registrants[registrant] = true;
            RegistrantApproval(registrant);
        }
    }

    function registrantRemove(address registrant) onlyAdmin {
        if (registrants[registrant] == true) {
            delete(registrants[registrant]);
            RegistrantRemoval(registrant);
        }
    }

    function withdrawFunds() onlyAdmin {
        if (!admin.send(this.balance)) {
            throw;
        }
    }

}