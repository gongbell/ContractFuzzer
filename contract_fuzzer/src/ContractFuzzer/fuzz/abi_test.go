package fuzz

import (
	"testing"
	"os"
)

func Test_getAbi(t *testing.T){
	str := `[{"constant":true,"inputs":[],"name":"name","outputs":[{"name":"","type":"string"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"_bidId","type":"uint256"},{"name":"_slotId","type":"uint256"},{"name":"_peer","type":"bytes32"}],"name":"acceptBid","outputs":[],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"_adunitId","type":"uint256"},{"name":"_target","type":"uint256"},{"name":"_rewardAmount","type":"uint256"},{"name":"_timeout","type":"uint256"},{"name":"_peer","type":"bytes32"}],"name":"placeBid","outputs":[],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"_bidId","type":"uint256"}],"name":"refundBid","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"_adunitId","type":"uint256"}],"name":"getAllBidsByAdunit","outputs":[{"name":"","type":"uint256[]"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"_adslotId","type":"uint256"},{"name":"_state","type":"uint256"}],"name":"getBidsByAdslot","outputs":[{"name":"","type":"uint256[]"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"_bidId","type":"uint256"}],"name":"giveupBid","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"_bidId","type":"uint256"}],"name":"getBid","outputs":[{"name":"","type":"uint256"},{"name":"","type":"uint256"},{"name":"","type":"uint256"},{"name":"","type":"uint256"},{"name":"","type":"uint256"},{"name":"","type":"uint256"},{"name":"","type":"bytes32"},{"name":"","type":"bytes32"},{"name":"","type":"uint256"},{"name":"","type":"bytes32"},{"name":"","type":"bytes32"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"bidsCount","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"_adunitId","type":"uint256"},{"name":"_state","type":"uint256"}],"name":"getBidsByAdunit","outputs":[{"name":"","type":"uint256[]"}],"payable":false,"type":"function"},{"constant":false,"inputs":[],"name":"withdrawEther","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"registry","outputs":[{"name":"","type":"address"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"tokenaddr","type":"address"}],"name":"withdrawToken","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"owner","outputs":[{"name":"","type":"address"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"_bidId","type":"uint256"},{"name":"_report","type":"bytes32"}],"name":"verifyBid","outputs":[],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"_bidId","type":"uint256"}],"name":"cancelBid","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"_bidId","type":"uint256"}],"name":"getBidReports","outputs":[{"name":"","type":"bytes32"},{"name":"","type":"bytes32"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"_bidId","type":"uint256"}],"name":"claimBidReward","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"_adslotId","type":"uint256"}],"name":"getAllBidsByAdslot","outputs":[{"name":"","type":"uint256[]"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"newOwner","type":"address"}],"name":"transferOwnership","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"token","outputs":[{"name":"","type":"address"}],"payable":false,"type":"function"},{"inputs":[{"name":"_token","type":"address"},{"name":"_registry","type":"address"}],"payable":false,"type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"name":"bidId","type":"uint256"},{"indexed":false,"name":"advertiser","type":"address"},{"indexed":false,"name":"adunitId","type":"uint256"},{"indexed":false,"name":"adunitIpfs","type":"bytes32"},{"indexed":false,"name":"target","type":"uint256"},{"indexed":false,"name":"rewardAmount","type":"uint256"},{"indexed":false,"name":"timeout","type":"uint256"},{"indexed":false,"name":"advertiserPeer","type":"bytes32"}],"name":"LogBidOpened","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"bidId","type":"uint256"},{"indexed":false,"name":"publisher","type":"address"},{"indexed":false,"name":"adslotId","type":"uint256"},{"indexed":false,"name":"adslotIpfs","type":"bytes32"},{"indexed":false,"name":"acceptedTime","type":"uint256"},{"indexed":false,"name":"publisherPeer","type":"bytes32"}],"name":"LogBidAccepted","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"bidId","type":"uint256"}],"name":"LogBidCanceled","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"bidId","type":"uint256"}],"name":"LogBidExpired","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"bidId","type":"uint256"},{"indexed":false,"name":"advReport","type":"bytes32"},{"indexed":false,"name":"pubReport","type":"bytes32"}],"name":"LogBidCompleted","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"_bidId","type":"uint256"},{"indexed":false,"name":"_wallet","type":"address"},{"indexed":false,"name":"_amount","type":"uint256"}],"name":"LogBidRewardClaimed","type":"event"}]`
    abi,err := newAbi([]byte(str))
	abi.fuzz()
	writer,_ := os.Create("test.txt")
	abi.OutputValue(writer)
	if err!=nil{
		t.Fatalf("%s",err)
	}else {
		t.Logf("%s",abi)
	}
}
func TestAbi_fuzz(t *testing.T){
	//str := `[{"constant":true,"inputs":[],"name":"mintingFinished","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"name","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_value","type":"uint256"}],"name":"approve","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_from","type":"address"},{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transferFrom","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"decimals","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_to","type":"address"},{"name":"_amount","type":"uint256"}],"name":"mint","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"balance","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[],"name":"finishMinting","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"owner","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"symbol","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transfer","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"},{"name":"_spender","type":"address"}],"name":"allowance","outputs":[{"name":"remaining","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"newOwner","type":"address"}],"name":"transferOwnership","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"anonymous":false,"inputs":[{"indexed":true,"name":"to","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Mint","type":"event"},{"anonymous":false,"inputs":[],"name":"MintFinished","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"owner","type":"address"},{"indexed":true,"name":"spender","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Approval","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":true,"name":"to","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Transfer","type":"event"}]`
	str,_ := readFile("/home/liuye/Workplace/PycharmProjects/DownloadContracts/verified_contract_abis/AiraRegistrarService.abi")
	abi,err := newAbi([]byte(str))

	if err!=nil{
		t.Fatalf("%s",err)
	}else {
		t.Logf("%s",abi)
	}
	if _,err:=abi.fuzz();err!=nil{
		t.Fatalf("%s",err)
	}
	for _,fun := range *abi{
		t.Logf("%s:%s",fun.Name,fun.Inputs)
	}
	//t.Logf("%s",abi)
}
type Person struct{
	Name string
}
func Test_Obj(t *testing.T){
	var p1 = Person{Name:"jack"}
	var p2 = &p1
	p2.Name = "lily"
	t.Logf("p1.name:%s",p1.Name)
	t.Logf("p2.name:%s",p2.Name)
}
func TestInput_fuzz(t *testing.T){
	str := `[{"name":"_addr","type":"uint256"},{"name":"_value","type":"address[5][]"}]`
	input,_:= newIOput([]byte(str))
	input.fuzz()
	t.Logf("%s",input)
}
