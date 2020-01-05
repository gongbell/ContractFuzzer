pragma solidity ^0.4.10;

// Copyright 2017 Bittrex







contract UserWallet {






    AbstractSweeperList sweeperList;
    function UserWallet(address _sweeperlist) {


        
sweeperList = AbstractSweeperList(_sweeperlist);
    }

    function () public payable {
 }

    function tokenFallback(address _from, uint _value, bytes _data) {


        (_from);

        (_value);

        (_data);
     }

    function sweep(address _token, uint _amount)
    returns (bool) {


        (_amount);

        
return sweeperList.sweeperOf(_token).delegatecall(msg.data);
    }
}

contract AbstractSweeperList {






    function sweeperOf(address _token) returns (address);
}

