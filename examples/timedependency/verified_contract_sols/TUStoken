/**
 *Submitted for verification at Etherscan.io on 2018-02-28
*/

pragma solidity ^0.4.18;
/*https://hashnode.com/post/how-to-build-your-own-ethereum-based-erc20-token-and-launch-an-ico-in-next-20-minutes-cjbcpwzec01c93awtbij90uzn*/
contract Token {
    event Transfer(address indexed _from, address indexed _to, uint256 _value);
    event Approval(address indexed _owner, address indexed _spender, uint256 _value);
    function transfer(address _to, uint256 _value) public returns (bool success) {
        if (balances[msg.sender] >= _value && _value > 0) {
            balances[msg.sender] -= _value;
            balances[_to] += _value;
            Transfer(msg.sender, _to, _value);
            return true;
        } else {
            return false; 
        }
    }
    function transferFrom(address _from, address _to, uint256 _value) public returns (bool success) {
        if (balances[_from] >= _value && allowed[_from][msg.sender] >= _value && _value > 0) {
            balances[_to] += _value;
            balances[_from] -= _value;
            allowed[_from][msg.sender] -= _value;
            Transfer(_from, _to, _value);
            return true;
        } else {
            return false;
        }
    }
    function balanceOf(address _owner) constant public returns (uint256 balance) {
        return balances[_owner];
    }
    function approve(address _spender, uint256 _value) public returns (bool success) {
        allowed[msg.sender][_spender] = _value;
        Approval(msg.sender, _spender, _value);
        return true;
    }
    function allowance(address _owner, address _spender) constant public returns (uint256 remaining) {
      return allowed[_owner][_spender];
    }
    mapping (address => uint256) balances;
    mapping (address => mapping (address => uint256)) allowed;
    uint256 public totalSupply;
}

contract TUStoken is Token {
    string public version = "0x02"; 

    string public name;
    uint8 public decimals;
    string public symbol;

    uint256 public totalEthInWei;
    address public hommie; 
    uint public stopsale;

    uint256 public JP_sum;
    address public JP_winner;
    bytes32 public JP_winningHash;

    function TUStoken() public {
        totalSupply = 0; 
        hommie = msg.sender;

        name = "true underground system token";
        decimals = 0;
        symbol = "TU$";

        totalEthInWei = 0;
        stopsale = 1522804800;   // 04.04.2018 |__4:20__| MSK (GMT+3)

        JP_sum = 0;
        JP_winner = hommie;
        JP_winningHash = "";
        
    }

    function finishICO() public {
        require(msg.sender == hommie);
        require(now > stopsale);
        uint256 tempo = JP_sum;
        JP_sum = 0;
        JP_winner.transfer(tempo);
    }

    function makeMoveBro() public payable {
        require(now < stopsale);
        uint256 amount = msg.value / (1 ether);
        uint toDonats = msg.value - (amount * (1 ether));  //сдача
        uint bonus = 1;
        if (amount > 1) {
            bonus = 2;
        } else if (amount >= 8) {
            bonus = 3;
        } else if (amount >= 96) {
            bonus = 4;
        } else if (amount >= 1618) {
            bonus = 5;
        }
        bytes32 pseudoRnd = keccak256(block.blockhash(block.number), now, msg.sender, msg.data);
        if (pseudoRnd > JP_winningHash) {
            JP_winner = msg.sender;
            JP_winningHash = pseudoRnd;
        }
        uint toJP = (amount * (1 ether) / 20) + (toDonats / 2);
        JP_sum += toJP;
        if (amount > 0) {
            uint256 tokens = amount * bonus;
            balances[msg.sender] += tokens;
            totalSupply += tokens;
            Transfer(hommie, msg.sender, tokens); 
        }

        totalEthInWei = totalEthInWei + msg.value;
        hommie.transfer(msg.value - toJP);
    }

    function approveAndCall(address _spender, uint256 _value, bytes _extraData) public returns (bool success) {
        require(_spender.call(bytes4(bytes32(keccak256("receiveApproval(address,uint256,address,bytes)"))), msg.sender, _value, this, _extraData));
        allowed[msg.sender][_spender] = _value;
        Approval(msg.sender, _spender, _value);
        return true;
    }

    function() public payable {
        makeMoveBro();
    }
}
