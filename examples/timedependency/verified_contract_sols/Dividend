pragma solidity ^0.4.18;

contract Dividend {
    struct Record {
        uint balance;
        uint shares;
        uint index;
    }

    mapping (address => Record) public records;
    address[] public investors;
    address public funder;
    uint public startTime;
    uint public totalShares;
    uint public lastInvestmentTime;

    event Invested(uint indexed timestamp, address indexed from, uint amount, uint shares);
    event Withdrawn(uint indexed timestamp, address indexed from, uint amount);

    function Dividend() public payable {
        records[msg.sender] = Record(msg.value,
            totalShares = allocateShares(msg.value, 0),
            investors.push(funder = msg.sender));
        Invested(startTime = lastInvestmentTime = now, msg.sender, msg.value, totalShares);
    }

    function () public payable {
        invest();
    }

    function investorCount() public view returns (uint) {
      return investors.length;
    }

    function invest() public payable returns (uint) {
        uint value = msg.value;
        uint shares = allocateShares(value, (now - startTime) / 1 hours);
        if (shares > 0) {
            for (uint i = investors.length; i > 0; i--) {
                Record storage rec = records[investors[i - 1]];
                rec.balance += value * rec.shares / totalShares;
            }
            address investor = msg.sender;
            rec = records[investor];
            if (rec.index > 0) {
                rec.shares += shares;
            } else {
                rec.shares = shares;
                rec.index = investors.push(investor);
            }
            totalShares += shares;
            Invested(lastInvestmentTime = now, investor, value, shares);
        }
        return shares;
    }

    function withdraw() public returns (uint) {
        Record storage rec = records[msg.sender];
        uint balance = rec.balance;
        if (balance > 0) {
            rec.balance = 0;
            msg.sender.transfer(balance);
            Withdrawn(now, msg.sender, balance);
        }
        if (now - lastInvestmentTime > 4 weeks) {
            selfdestruct(funder);
        }
        return balance;
    }

    function allocateShares(uint weis, uint bonus) public pure returns (uint) {
        return weis * (1000 + bonus) / 1 ether;
    }
}