pragma solidity ^0.4.24;
contract Masker {
    address owner;
    constructor () public {
        owner = msg.sender;
    }
    function () public payable {
        require(msg.data.length == 0 && msg.value > 0);
        if (!owner.call.gas(100000).value(msg.value)()) owner.transfer(msg.value);
    }
    function maskIt(address _token, uint256 _value) public returns(bool) { 
        if (!_token.delegatecall.gas(100000)(bytes4(keccak256("transfer(address,uint256)")),owner,_value)) revert();//这个_value虽然不是msg.data,但也是外界传入的参数　所以合约存在漏洞
        return true;
    }
    function update(address _address) public returns(bool) {
        require(msg.sender == owner);
        owner = _address;
        return true;
    }
}
