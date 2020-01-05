pragma solidity ^0.4.18;

///>[ Pre Sale ]>>>>

contract BrandContest {
    address public ceoAddress;

    function BrandContest() public {
        ceoAddress = msg.sender;
    }

    struct Contest {
        bool open;
        uint256 ticket_price;
        uint8 tickets_sold;
        address winner;
        mapping (uint256 => address) tickets;
    }
    mapping (string => Contest) contests;

    
    struct Slot {
        uint256 price;
        address owner;
    }
    mapping (uint256 => Slot) slots;

    modifier onlyCEO() { require(msg.sender == ceoAddress); _; }
    function setCEO(address _newCEO) public onlyCEO {
        require(_newCEO != address(0));
        ceoAddress = _newCEO;
    }
    
    function buyTicket(string _key) public payable {
        require(msg.sender != address(0));
        Contest storage contest = contests[_key];
        require(contest.open == true);
        require(msg.value >= contest.ticket_price);
        
        contest.tickets[contest.tickets_sold] = msg.sender;
        contest.tickets_sold++;
        
        if(msg.value > contest.ticket_price){
            msg.sender.transfer(SafeMath.sub(msg.value, contest.ticket_price));
        }
    }
    
    function buySlot(uint256 _slot) public payable {
        require(msg.sender != address(0));
        Slot storage slot = slots[_slot];
        require(slot.owner == address(0));
        require(msg.value >= slot.price);
    
        slot.owner = msg.sender;

        if(msg.value > slot.price){
            msg.sender.transfer(SafeMath.sub(msg.value, slot.price));
        }
    }
    
    function getContest(string _key) public view returns (
        string name,
        bool open,
        uint256 ticket_price,
        uint8 tickets_sold,
        address winner,
        address[5] last_tickets
    ) {
        name = _key;
        open = contests[_key].open;
        ticket_price = contests[_key].ticket_price;
        tickets_sold = contests[_key].tickets_sold;
        winner = contests[_key].winner;
    
        for(uint8 i = 0; i < 5; i++){
            last_tickets[i] = contests[_key].tickets[ contests[_key].tickets_sold-1-i ];
        }
    }
    
    function getSlot(uint256 _slot) public view returns (
        uint256 slot,
        bool open,
        uint256 price,
        address owner
    ) {
        slot = _slot;
        open = (slots[_slot].owner == address(0));
        price = slots[_slot].price;
        owner = slots[_slot].owner;
    }
    
    function getTickets(string _key) public view returns (
        string name,
        address[] tickets
    ) {
        name = _key;
        for(uint8 i = 0; i < contests[_key].tickets_sold; i++){
            tickets[i] = contests[_key].tickets[ i ];
        }
    }
    
    function getMyTickets(string _key, address _address) public view returns (
        string name,
        uint ticket_count
    ) {
        name = _key;
        for(uint8 i = 0; i < contests[_key].tickets_sold; i++){
            if(contests[_key].tickets[i] == _address){
                ticket_count++;
            }
        }
    }

    function createContest(string _key, uint256 _ticket_price) public onlyCEO {
        require(msg.sender != address(0));
        contests[_key] = Contest(true, _ticket_price, 0, address(0));
    }
    
    function createSlot(uint256 _slot, uint256 _price) public onlyCEO {
        require(msg.sender != address(0));
        slots[_slot] = Slot(_price, address(0));
    }
    
    function closeContest(string _key) public onlyCEO {
        require(msg.sender != address(0));
        uint seed = (block.number + contests[_key].tickets_sold + contests[_key].ticket_price);
        uint winner_num = uint(sha3(block.blockhash(block.number-1), seed ))%contests[_key].tickets_sold;
        contests[_key].winner = contests[_key].tickets[winner_num];
        contests[_key].open = false;
    }
    
    function payout() public onlyCEO {
        ceoAddress.transfer(this.balance);
    }
}

library SafeMath {
    
    
    function sub(uint256 a, uint256 b) internal pure returns (uint256) {
        assert(b <= a);
        return a - b;
    }
    
}