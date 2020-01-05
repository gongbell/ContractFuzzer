pragma solidity ^0.4.11;

contract AnyContract{
    mapping(address => uint256) public numbers;
    mapping(address => string) public texts;
    
    function add(uint256 _a,uint256 _b) public{
        numbers[msg.sender] =_a+_b;
    }
    
    function write(string _text) public{
        texts[msg.sender] = _text;
    }
    
    function batchWrite(uint256 _a,uint256 _b,string _text) public payable{
        numbers[msg.sender] =_a+_b;
        texts[msg.sender] = _text;
    }
    function getBalance() public view returns(uint256){
        return address(this).balance;
    }
}

