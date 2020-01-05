pragma solidity ^0.4.4;


contract EDProxy {

    function EDProxy() public {
    }

    function dtrade(address _callee, uint8 v1, uint8 v2, uint256[] uints,address[] addresses,bytes32[] b) public {
        
        if (_callee.delegatecall(bytes4(keccak256("trade(address,uint256,address,uint256,uint256,uint256,address,uint8,bytes32,bytes32,uint256)")),
          addresses[0],
          uints[0],
          addresses[2],
          uints[2],
          uints[4],
          uints[6],
          addresses[4],
          v1,
          b[0],
          b[2],
          uints[8]
          )) {
        (_callee.delegatecall(bytes4(keccak256("trade(address,uint256,address,uint256,uint256,uint256,address,uint8,bytes32,bytes32,uint256)")),
           addresses[1],
           uints[1],
           addresses[3],
           uints[3],
           uints[5],
           uints[7],
           addresses[5],
           v2,
           b[1],
           b[3],
           uints[9]
           ));
          }
    }
    
     function testcall(address _callee)  public {
        bytes32[] memory b = new bytes32[](4);
        address[] memory addrs = new address[](6);
        uint256[] memory ints = new uint256[](12);
        uint8 v1;
        uint8 v2;

        bytes32 somebytes;
        ints[0]=1;
        ints[1]=2;
        ints[2]=3;
        ints[3]=4;
        ints[4]=5;
        ints[5]=6;
        ints[6]=7;
        ints[7]=8;
        ints[8]=9;
        ints[9]=10;
        v1=11;
        v2=12;
        b[0]=somebytes;
        b[1]=somebytes;
        b[2]=somebytes;
        b[3]=somebytes;
        addrs[0]=0xdc04977a2078c8ffdf086d618d1f961b6c54111;
        addrs[1]=0xdc04977a2078c8ffdf086d618d1f961b6c54222;
        addrs[2]=0xdc04977a2078c8ffdf086d618d1f961b6c54333;
        addrs[3]=0xdc04977a2078c8ffdf086d618d1f961b6c54444;
        addrs[4]=0xdc04977a2078c8ffdf086d618d1f961b6c54555;
        addrs[5]=0xdc04977a2078c8ffdf086d618d1f961b6c54666;
        dtrade(_callee, v1, v2, ints, addrs,b);
    }
    
}