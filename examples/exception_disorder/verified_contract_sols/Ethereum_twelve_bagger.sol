// <ORACLIZE_API>
/*
Copyright (c) 2015-2016 Oraclize srl, Thomas Bertani



Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:



The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.



THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

contract OraclizeI {
    address public cbAddress;
    function query(uint _timestamp, string _datasource, string _arg) returns (bytes32 _id);
    function query_withGasLimit(uint _timestamp, string _datasource, string _arg, uint _gaslimit) returns (bytes32 _id);
    function query2(uint _timestamp, string _datasource, string _arg1, string _arg2) returns (bytes32 _id);
    function query2_withGasLimit(uint _timestamp, string _datasource, string _arg1, string _arg2, uint _gaslimit) returns (bytes32 _id);
    function getPrice(string _datasource) returns (uint _dsprice);
    function getPrice(string _datasource, uint gaslimit) returns (uint _dsprice);
    function useCoupon(string _coupon);
    function setProofType(byte _proofType);
}
contract OraclizeAddrResolverI {
    function getAddress() returns (address _addr);
}
contract usingOraclize {
    uint constant day = 60*60*24;
    uint constant week = 60*60*24*7;
    uint constant month = 60*60*24*30;
    byte constant proofType_NONE = 0x00;
    byte constant proofType_TLSNotary = 0x10;
    byte constant proofStorage_IPFS = 0x01;
    uint8 constant networkID_auto = 0;
    uint8 constant networkID_mainnet = 1;
    uint8 constant networkID_testnet = 2;
    uint8 constant networkID_morden = 2;
    uint8 constant networkID_consensys = 161;

    OraclizeAddrResolverI OAR;
    
    OraclizeI oraclize;
    modifier oraclizeAPI {
        address oraclizeAddr = OAR.getAddress();
        if (oraclizeAddr == 0){
            oraclize_setNetwork(networkID_auto);
            oraclizeAddr = OAR.getAddress();
        }
        oraclize = OraclizeI(oraclizeAddr);
        _
    }
    modifier coupon(string code){
        oraclize = OraclizeI(OAR.getAddress());
        oraclize.useCoupon(code);
        _
    }

    function oraclize_setNetwork(uint8 networkID) internal returns(bool){
        if (getCodeSize(0x1d3b2638a7cc9f2cb3d298a3da7a90b67e5506ed)>0){
            OAR = OraclizeAddrResolverI(0x1d3b2638a7cc9f2cb3d298a3da7a90b67e5506ed);
            return true;
        }
        if (getCodeSize(0x9efbea6358bed926b293d2ce63a730d6d98d43dd)>0){
            OAR = OraclizeAddrResolverI(0x9efbea6358bed926b293d2ce63a730d6d98d43dd);
            return true;
        }
        if (getCodeSize(0x20e12a1f859b3feae5fb2a0a32c18f5a65555bbf)>0){
            OAR = OraclizeAddrResolverI(0x20e12a1f859b3feae5fb2a0a32c18f5a65555bbf);
            return true;
        }
        return false;
    }
    
    function oraclize_query(string datasource, string arg) oraclizeAPI internal returns (bytes32 id){
        uint price = oraclize.getPrice(datasource);
        if (price > 1 ether + tx.gasprice*200000) return 0; // unexpectedly high price
        return oraclize.query.value(price)(0, datasource, arg);
    }
    function oraclize_query(uint timestamp, string datasource, string arg) oraclizeAPI internal returns (bytes32 id){
        uint price = oraclize.getPrice(datasource);
        if (price > 1 ether + tx.gasprice*200000) return 0; // unexpectedly high price
        return oraclize.query.value(price)(timestamp, datasource, arg);
    }
    function oraclize_query(uint timestamp, string datasource, string arg, uint gaslimit) oraclizeAPI internal returns (bytes32 id){
        uint price = oraclize.getPrice(datasource, gaslimit);
        if (price > 1 ether + tx.gasprice*gaslimit) return 0; // unexpectedly high price
        return oraclize.query_withGasLimit.value(price)(timestamp, datasource, arg, gaslimit);
    }
    function oraclize_query(string datasource, string arg, uint gaslimit) oraclizeAPI internal returns (bytes32 id){
        uint price = oraclize.getPrice(datasource, gaslimit);
        if (price > 1 ether + tx.gasprice*gaslimit) return 0; // unexpectedly high price
        return oraclize.query_withGasLimit.value(price)(0, datasource, arg, gaslimit);
    }
    function oraclize_query(string datasource, string arg1, string arg2) oraclizeAPI internal returns (bytes32 id){
        uint price = oraclize.getPrice(datasource);
        if (price > 1 ether + tx.gasprice*200000) return 0; // unexpectedly high price
        return oraclize.query2.value(price)(0, datasource, arg1, arg2);
    }
    function oraclize_query(uint timestamp, string datasource, string arg1, string arg2) oraclizeAPI internal returns (bytes32 id){
        uint price = oraclize.getPrice(datasource);
        if (price > 1 ether + tx.gasprice*200000) return 0; // unexpectedly high price
        return oraclize.query2.value(price)(timestamp, datasource, arg1, arg2);
    }
    function oraclize_query(uint timestamp, string datasource, string arg1, string arg2, uint gaslimit) oraclizeAPI internal returns (bytes32 id){
        uint price = oraclize.getPrice(datasource, gaslimit);
        if (price > 1 ether + tx.gasprice*gaslimit) return 0; // unexpectedly high price
        return oraclize.query2_withGasLimit.value(price)(timestamp, datasource, arg1, arg2, gaslimit);
    }
    function oraclize_query(string datasource, string arg1, string arg2, uint gaslimit) oraclizeAPI internal returns (bytes32 id){
        uint price = oraclize.getPrice(datasource, gaslimit);
        if (price > 1 ether + tx.gasprice*gaslimit) return 0; // unexpectedly high price
        return oraclize.query2_withGasLimit.value(price)(0, datasource, arg1, arg2, gaslimit);
    }
    function oraclize_cbAddress() oraclizeAPI internal returns (address){
        return oraclize.cbAddress();
    }
    function oraclize_setProof(byte proofP) oraclizeAPI internal {
        return oraclize.setProofType(proofP);
    }

    function getCodeSize(address _addr) constant internal returns(uint _size) {
        assembly {
            _size := extcodesize(_addr)
        }
    }


    function parseAddr(string _a) internal returns (address){
        bytes memory tmp = bytes(_a);
        uint160 iaddr = 0;
        uint160 b1;
        uint160 b2;
        for (uint i=2; i<2+2*20; i+=2){
            iaddr *= 256;
            b1 = uint160(tmp[i]);
            b2 = uint160(tmp[i+1]);
            if ((b1 >= 97)&&(b1 <= 102)) b1 -= 87;
            else if ((b1 >= 48)&&(b1 <= 57)) b1 -= 48;
            if ((b2 >= 97)&&(b2 <= 102)) b2 -= 87;
            else if ((b2 >= 48)&&(b2 <= 57)) b2 -= 48;
            iaddr += (b1*16+b2);
        }
        return address(iaddr);
    }


    function strCompare(string _a, string _b) internal returns (int) {
        bytes memory a = bytes(_a);
        bytes memory b = bytes(_b);
        uint minLength = a.length;
        if (b.length < minLength) minLength = b.length;
        for (uint i = 0; i < minLength; i ++)
            if (a[i] < b[i])
                return -1;
            else if (a[i] > b[i])
                return 1;
        if (a.length < b.length)
            return -1;
        else if (a.length > b.length)
            return 1;
        else
            return 0;
   } 

    function indexOf(string _haystack, string _needle) internal returns (int)
    {
        bytes memory h = bytes(_haystack);
        bytes memory n = bytes(_needle);
        if(h.length < 1 || n.length < 1 || (n.length > h.length)) 
            return -1;
        else if(h.length > (2**128 -1))
            return -1;                                  
        else
        {
            uint subindex = 0;
            for (uint i = 0; i < h.length; i ++)
            {
                if (h[i] == n[0])
                {
                    subindex = 1;
                    while(subindex < n.length && (i + subindex) < h.length && h[i + subindex] == n[subindex])
                    {
                        subindex++;
                    }   
                    if(subindex == n.length)
                        return int(i);
                }
            }
            return -1;
        }   
    }

    function strConcat(string _a, string _b, string _c, string _d, string _e) internal returns (string){
        bytes memory _ba = bytes(_a);
        bytes memory _bb = bytes(_b);
        bytes memory _bc = bytes(_c);
        bytes memory _bd = bytes(_d);
        bytes memory _be = bytes(_e);
        string memory abcde = new string(_ba.length + _bb.length + _bc.length + _bd.length + _be.length);
        bytes memory babcde = bytes(abcde);
        uint k = 0;
        for (uint i = 0; i < _ba.length; i++) babcde[k++] = _ba[i];
        for (i = 0; i < _bb.length; i++) babcde[k++] = _bb[i];
        for (i = 0; i < _bc.length; i++) babcde[k++] = _bc[i];
        for (i = 0; i < _bd.length; i++) babcde[k++] = _bd[i];
        for (i = 0; i < _be.length; i++) babcde[k++] = _be[i];
        return string(babcde);
    }
    
    function strConcat(string _a, string _b, string _c, string _d) internal returns (string) {
        return strConcat(_a, _b, _c, _d, "");
    }

    function strConcat(string _a, string _b, string _c) internal returns (string) {
        return strConcat(_a, _b, _c, "", "");
    }

    function strConcat(string _a, string _b) internal returns (string) {
        return strConcat(_a, _b, "", "", "");
    }

    // parseInt
    function parseInt(string _a) internal returns (uint) {
        return parseInt(_a, 0);
    }

    // parseInt(parseFloat*10^_b)
    function parseInt(string _a, uint _b) internal returns (uint) {
        bytes memory bresult = bytes(_a);
        uint mint = 0;
        bool decimals = false;
        for (uint i=0; i<bresult.length; i++){
            if ((bresult[i] >= 48)&&(bresult[i] <= 57)){
                if (decimals){
                   if (_b == 0) break;
                    else _b--;
                }
                mint *= 10;
                mint += uint(bresult[i]) - 48;
            } else if (bresult[i] == 46) decimals = true;
        }
        return mint;
    }
    

}
// </ORACLIZE_API>



    






contract Ethereum_twelve_bagger is usingOraclize
{

 							//declares global variables
string hexcomparisonchr;
string A;
string B;

uint8 lotteryticket;
address creator;
int lastgainloss;
string lastresult;
string K;
string information;
  
 
address player;
uint8 gameResult;
uint128 wager; 
 mapping (bytes32=>uint) bets;
mapping (bytes32 => address) gamesPlayer;
 

   function  Ethereum_twelve_bagger() private 
    { 
        creator = msg.sender; 								
    }

    function Set_your_game_number_between_1_15(string Set_your_game_number_between_1_15)			//sets game number
 {
	player=msg.sender;
    	A=Set_your_game_number_between_1_15;
	wager =uint128(msg.value);
	
	lastresult = "Waiting for a lottery number from Wolfram Alpha";
	lastgainloss = 0;
	B="The new right lottery number is not ready yet";
	information = "The new right lottery number is not ready yet";
	testWager();
	
	WolframAlpha();
}

     	 
	 
	


    

    function WolframAlpha() private {
	if (wager == 0) return;		//if wager is 0, abort 
        oraclize_setNetwork(networkID_testnet);
        oraclize_setProof(proofType_TLSNotary | proofStorage_IPFS);
     	bytes32 myid =  oraclize_query(0,"WolframAlpha", "random number between 1 and 15");
	bets[myid] = wager;
	gamesPlayer[myid] = player;
    }

 	    function __callback(bytes32 myid, string result, bytes proof) {
        if (msg.sender != oraclize_cbAddress()) throw;
	
        B = result;
	
	wager=uint128(bets[myid]);
	player=gamesPlayer[myid];
	test(A,B);
	returnmoneycreator(gameResult,wager);
	return;
        
}
 
function test(string A,string B) private
{ 
information ="The right lottery number is now ready. One Eth is 10**18 Wei.";
K="K";
bytes memory test = bytes(A);
bytes memory kill = bytes(K);
	 if (test[0]==kill[0] && player == creator)			//Creator can kill contract. Contract does not hold players money.
	{
		suicide(creator);} 
 
    	
    


if (equal(A,B))
{
lastgainloss =(12*wager);
	    	lastresult = "Win!";
	    	player.send(wager * 12);  

gameResult=0;
return;}
else 
{
lastgainloss = int(wager) * -1;
	    	lastresult = "Loss";
	    	gameResult=1;
	    									// Player lost. Return nothing.
	    	return;


 
	}
}


 
function testWager() private
{if((wager * 12) > this.balance) 					// contract has to have 12*wager funds to be able to pay out. (current balance includes the wager sent)
    	{
    		lastresult = "Bet is larger than games's ability to pay";
    		lastgainloss = 0;
    		player.send(wager); // return wager
		gameResult=0;
		wager=0;
		B="Bet is larger than games's ability to pay";
		information ="Bet is larger than games's ability to pay";
    		return;
}

else if (wager < 100000000000000000)					// Minimum bet is 0.1 eth 
    	{
    		lastresult = "Minimum bet is 0.1 eth";
    		lastgainloss = 0;
    		player.send(wager); // return wager
		gameResult=0;
		wager=0;
		B="Minimum bet is 0.1 eth";
		information ="Minimum bet is 0.1 eth";
    		return;
}




	else if (wager == 0)
    	{
    		lastresult = "Wager was zero";
    		lastgainloss = 0;
		gameResult=0;
    		// nothing wagered, nothing returned
    		return;
    	}
}



    /// @dev Does a byte-by-byte lexicographical comparison of two strings.
    /// @return a negative number if `_a` is smaller, zero if they are equal
    /// and a positive numbe if `_b` is smaller.
    function compare(string A, string B) private returns (int) {
        bytes memory a = bytes(A);
        bytes memory b = bytes(B);
        uint minLength = a.length;
        if (b.length < minLength) minLength = b.length;
        //@todo unroll the loop into increments of 32 and do full 32 byte comparisons
        for (uint i = 0; i < minLength; i ++)
            if (a[i] < b[i])
                return -1;
            else if (a[i] > b[i])
                return 1;
        if (a.length < b.length)
            return -1;
        else if (a.length > b.length)
            return 1;
        else
            return 0;
    }
    /// @dev Compares two strings and returns true iff they are equal.
    function equal(string A, string B) private returns (bool) 
       {
        return compare(A, B) == 0;
}

function returnmoneycreator(uint8 gameResult,uint wager) private		//If game has over 50 eth, contract will send all additional eth to owner
	{
	if (gameResult==1&&this.balance>50000000000000000000)
	{creator.send(wager);
	return; 
	}
 
	else if
	(
	gameResult==1&&this.balance>20000000000000000000)				//If game has over 20 eth, contract will send œ of any additional eth to owner
	{creator.send(wager/2);
	return; }
	}
 
/**********
functions below give information about the game in Ethereum Wallet
 **********/
 
 	function Results_of_the_last_round() constant returns (uint players_bet_in_Wei, string last_result,string Last_player_s_lottery_ticket,address last_player,string The_right_lottery_number,int Player_s_gain_or_Loss_in_Wei,string info)
    { 
   	last_player=player;	
	Last_player_s_lottery_ticket=A;
	The_right_lottery_number=B;
	last_result=lastresult;
	players_bet_in_Wei=wager;
	Player_s_gain_or_Loss_in_Wei=lastgainloss;
	info = information;
	
 
    }

 
    
   
	function Game_balance_in_Ethers() constant returns (uint balance, string info)
    { 
        info = "Choose number between 1 and 15. Win pays wager*12. Minimum bet is 0.1 eth. Maximum bet is game balance/12. Game balance is shown in full Ethers.";
    	balance=(this.balance/10**18);

    }
    
   
}