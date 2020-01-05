/*
 * This file is part of Safether.
 *
 * Safether is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Safether is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with Safether.  If not, see <http://www.gnu.org/licenses/>.
 */

/* 
 * SafetherStorage Contract
 * For storing depositor data.
 */
contract SafetherStorage {
    
    /*
     * Depositor Struct is storage for user.
     * 
     * _token is access key required for find assets.
     * _data is storage for depositor.
     *
     * _data[0] : Register Block Number.
     * _data[1] : Period holding assets.
     * _data[2] : Amount of holding assets.
     *
     */
    struct Depositor {
        bytes8     _token;
        uint256[3]  _data;
    }
    
    mapping (address=>Depositor) internal _depositor;
}

/* 
 * SafetherModifier Contract
 * For declaring modifier function.
 */
contract SafetherModifier is SafetherStorage {
    modifier isRegisterd {
        require(_depositor[msg.sender]._token != 0x0);
        _;
    }
    
    modifier isNotRegisterd {
        require(_depositor[msg.sender]._token == 0x0);
        _;
    }
    
    /*
     * isValidDespositor
     * Modifier function for finding assets.
     */
    modifier isValidDepositor(address depositor, bytes8 token) {
        require(_depositor[depositor]._token == token);
        require(_depositor[depositor]._data[2] > 0);
        require(block.number >= _depositor[depositor]._data[1]);
        _;
    }
}

/* 
 * SafetherInterface Contract
 * Interface contract for declaring Safether Contract.
 */
contract SafetherInterface {
    function authentication(bytes8 token) public constant returns(bool);
    function getDepositor() public constant returns(uint256[3]);
    
    function register(bytes7 password) public;
    function deposit(uint256 period) public payable;
    function withdraw(address depositor, bytes8 token) public payable;
    function cancel() public payable;
}

/* 
 * Safether Contract
 * Implements the inherited functions of interface contract.
 */
contract Safether is SafetherModifier, SafetherInterface {

    /* 
     * authentication Function
     *
     * It only returns information about msg.sender to 
     * prevent hacking through DoS attack.
     * 
     */
    function authentication(bytes8 token) public constant returns(bool) {
        return _depositor[msg.sender]._token == token;
    }
    
    /* 
     * getDepositor Function
     *
     * It only returns information about msg.sender to 
     * prevent hacking through DoS attack.
     * 
     */
    function getDepositor() public constant returns (uint256[3]) {
        return (_depositor[msg.sender]._data);
    }
    
    function register(bytes7 password) public isNotRegisterd {
        _depositor[msg.sender]._token = bytes8(keccak256(block.number, msg.sender, password));
        _depositor[msg.sender]._data[0] = block.number;
    }
    
    function deposit(uint256 period) public payable isRegisterd {
        _depositor[msg.sender]._data[1] = block.number + period;
        _depositor[msg.sender]._data[2] += msg.value;
    }
    
    /* 
     * withdraw Function
     * 
     * Recipients can not know how much money is stored, 
     * nor can they know the duration. 
     
     * This is a preventive measure to prevent DoS attacks from hackers. 
     * This is because the metamask informs the modifier of the error in the context execution stage. 
     * 
     * Hackers can try this information without paying for it, 
     * but they do not know if the password is wrong, the storage period has not expired, 
     * or the money does not exist. 
     *
     * Even if you hacked the correct token, This makes hacking difficult.
     *
     */
    function withdraw(address depositor, bytes8 token) public payable isValidDepositor(depositor, token) {
        uint256 tempDeposit = _depositor[depositor]._data[2];
         _depositor[depositor]._data[2] = 0;
         msg.sender.transfer(tempDeposit + msg.value);
    }
    
    function cancel() public payable isRegisterd {
        uint256 tempDeposit = _depositor[msg.sender]._data[2];
        delete _depositor[msg.sender];
        msg.sender.transfer(tempDeposit + msg.value);
    }
}