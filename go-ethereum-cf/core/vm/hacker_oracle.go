/**
*@hacker_oracle.go
* define
*       1 vulunerable oracle(such as Reentrancy,ExceptionDisorder,GaslessSend,CallEtherFailed.etc) formally.
*       2 other Operation (such as sendOp,CallOp.etc) Oracle.
*       3 simple count(such as RepeatedCall) Oracle.
* And give strict check condition for each oracle upper.
*/
package vm

import (
	"github.com/ethereum/go-ethereum/common"
	"strings"
	"io"
	"fmt"
	"encoding/hex"
)

//This is a simple trick.
//help us differentiate account address and memory address.
const prefixOfNoAccount  ="0x00000000000000000000"
func IsAccountAddress(addr common.Address) bool{
    addr_str := strings.ToLower(addr.Hex())
	return !strings.Contains(addr_str,prefixOfNoAccount)
}
type Oracle interface {
	InitOracle(hacker_call_hashs []common.Hash,hacker_calls []*HackerContractCall)
	TestOracle() bool
	Write(writer io.Writer)
	String() string
}


type HackerRootCallFailed struct{
	hacker_call_hashs []common.Hash
	hacker_calls []*HackerContractCall
}
func NewHackerRootCallFailed() *HackerRootCallFailed{
	return &HackerRootCallFailed{}
}
func  (oracle *HackerRootCallFailed) InitOracle(hacker_call_hashs []common.Hash,hacker_calls []*HackerContractCall){
	oracle.hacker_calls = hacker_calls
	oracle.hacker_call_hashs = hacker_call_hashs

}
func (oracle *HackerRootCallFailed) TestOracle() bool{
	var rootCall = oracle.hacker_calls[0]
	return rootCall.throwException
}
func (oracle *HackerRootCallFailed) Write(writer io.Writer){
	var str string
	str = "\nHackerRootCallFailed\n"
	writer.Write([]byte(str))
}
func (Oracle *HackerRootCallFailed) String() string{
	return "HackerRootCallFailed"
}
/**
*Oracle: HackerReentrancy
*to help us detect reentrancy bugs.
*The oracle use a trick to check wheter a contract provide a anti-reentrancy policy
*via comparing opcode len between the two 'reentrancy' calls.
*/
type HackerReentrancy struct{
    hacker_call_hashs []common.Hash
	hacker_calls []*HackerContractCall
	repeatedPairs [][2]*HackerContractCall
	feauture  string
}
func NewHackerReentrancy() *HackerReentrancy{
	return &HackerReentrancy{}
}
func  (oracle *HackerReentrancy) InitOracle(hacker_call_hashs []common.Hash,hacker_calls []*HackerContractCall){
	oracle.hacker_calls = hacker_calls
	oracle.hacker_call_hashs = hacker_call_hashs
	oracle.repeatedPairs = make([][2]*HackerContractCall,0,0)
	oracle.feauture = ""
}
func (oracle *HackerReentrancy) TestOracle() bool{
	var hasReen bool
	hasReen = false
    i :=0
    hash1 := oracle.hacker_call_hashs[i]
	for j := i+1;j<len(oracle.hacker_call_hashs);j++{
		    hash2 := oracle.hacker_call_hashs[j]
		    //compare two call hash equal?
		    //compare two call operationLen equal? detect anti-reentrancy protection
		    // if strings.Compare(hash1.String(),hash2.String()) == 0&&oracle.hacker_calls[i].OperationStack.len()<=oracle.hacker_calls[j].OperationStack.len()&&
		    // 	(oracle.hacker_calls[i].isAncestor(oracle.hacker_calls[j]) || oracle.hacker_calls[j].isAncestor(oracle.hacker_calls[i])){
			if strings.Compare(hash1.String(),hash2.String()) == 0{
			    if oracle.hacker_calls[i].OperationStack.len()<=oracle.hacker_calls[j].OperationStack.len(){
					oracle.feauture = "le"
				}else{
					oracle.feauture = "anti-reentrancy"
				}
				repeatedPair := [2]*HackerContractCall{oracle.hacker_calls[i],oracle.hacker_calls[j]}
			    oracle.repeatedPairs = append(oracle.repeatedPairs,repeatedPair)
			    hasReen = true
		    }
	}
    return hasReen
}
func (oracle *HackerReentrancy) Write(writer io.Writer){
	var str string
	str = "\nReentrancyException Call Pairs:\n"
	if len(oracle.repeatedPairs)<=0{
		return
	}
	for _,pair := range oracle.repeatedPairs{
		str += fmt.Sprintf("<%p,%p>\n",pair[0],pair[1])
		str +=fmt.Sprintf("profile:\n")
		for _,call := range  pair {
			str += fmt.Sprintf("%s=>%s  (value:%s,gas:%s)  (input:%s)\n",call.caller.Hex(),call.callee.Hex(),call.value.Text(10),call.gas.Text(10),hex.EncodeToString(call.input))
		}
	}
	writer.Write([]byte(str))
}
func (Oracle *HackerReentrancy) String() string{
	return "HackerReentrancy"+Oracle.feauture
}
/**
* Oracle: HackerRepeatedCall
* to help us detect reductant inner call to others in the contract.
* and this could help developer to write better code to save gas according gas mechanism on Ethereum.
*/
type HackerRepeatedCall struct{
	hacker_call_hashs []common.Hash
	hacker_calls []*HackerContractCall
	repeatedPairs [][2]*HackerContractCall
}
func NewHackerRepeatedCall() *HackerRepeatedCall{
	return &HackerRepeatedCall{}
}
func  (oracle *HackerRepeatedCall) InitOracle(hacker_call_hashs []common.Hash,hacker_calls []*HackerContractCall){
	oracle.hacker_calls = hacker_calls
	oracle.hacker_call_hashs = hacker_call_hashs
	oracle.repeatedPairs = make([][2]*HackerContractCall,0,0)
}
func (oracle *HackerRepeatedCall) TestOracle() bool{
	hasRepeated := false
	//nextcalls_len := len(oracle.hacker_calls[0].nextcalls)
	for i,_ := range oracle.hacker_call_hashs{
		hash1 := oracle.hacker_call_hashs[i]
		for j := i+1;j<len(oracle.hacker_call_hashs);j++{
			hash2 := oracle.hacker_call_hashs[j]
			if strings.Compare(hash1.String(),hash2.String()) == 0&&
				oracle.hacker_calls[i].isBrother(i,oracle.hacker_calls[j]){
				repeatedPair := [2]*HackerContractCall{oracle.hacker_calls[i],oracle.hacker_calls[j]}
				oracle.repeatedPairs = append(oracle.repeatedPairs,repeatedPair)
				hasRepeated = true
			}
		}
	}
	return hasRepeated
}
func (oracle *HackerRepeatedCall) Write(writer io.Writer){
	var str string
	str = "\nHackerRepeated Call Pairs:\n"
	if len(oracle.repeatedPairs)<=0{
		return
	}
	for _,pair := range oracle.repeatedPairs{
		str += fmt.Sprintf("<%p,%p>\n",pair[0],pair[1])
		str +=fmt.Sprintf("profile:\n")
		for _,call := range  pair {
			str += fmt.Sprintf("%s=>%s  (value:%s,gas:%s)  (input:%s)\n",call.caller.Hex(),call.callee.Hex(),call.value.Text(10),call.gas.Text(10),hex.EncodeToString(call.input))
		}
	}
	writer.Write([]byte(str))
}
func (oracle *HackerRepeatedCall) String() string{
	return "HackerRepeatedCall"
}

/**
*Oracle: HackerEtherTransfer
*to help us to check whether ether to be sent
*/
type HackerEtherTransfer struct{
	hacker_call_hashs     []common.Hash
	hacker_calls          []*HackerContractCall
	hacker_exception_calls []*HackerContractCall
}
func NewHackerEtherTransfer() *HackerEtherTransfer{
	return new(HackerEtherTransfer)
}
func (oracle *HackerEtherTransfer) InitOracle(hacker_call_hashs []common.Hash,hacker_calls []*HackerContractCall){
	oracle.hacker_calls = hacker_calls
	oracle.hacker_call_hashs = hacker_call_hashs
	oracle.hacker_exception_calls = make([]*HackerContractCall,0)
}
func (oracle *HackerEtherTransfer) TestOracle() bool{
	ret := false
	// if oracle.triggerOracle(oracle.hacker_calls[0]){
	// 	oracle.hacker_exception_calls = append(oracle.hacker_exception_calls, oracle.hacker_calls[0])
	// 	ret = true
	// }
    calls := oracle.hacker_calls[0].nextcalls
    for _,call := range  calls {
		if oracle.triggerOracle(call) {
			oracle.hacker_exception_calls = append(oracle.hacker_exception_calls, call)
			ret = true
		}
	}
	return ret
}
func (oracle *HackerEtherTransfer) triggerOracle(call *HackerContractCall) bool{
	nextcalls := call.nextcalls
	for _, n_call := range  nextcalls{
		if n_call.value.Uint64()>0{
			return true
		}
	}
	return call.value.Uint64()>0
}
func (oracle *HackerEtherTransfer) Write(writer io.Writer){
	str := ""
	if len(oracle.hacker_exception_calls) >0{
		str = "\nSendOp found:\n"
	}else{
		return
	}
	for _,call := range oracle.hacker_exception_calls{
		str += fmt.Sprintf("%p\n",call)
		str +=fmt.Sprintf("profile:\n")
		str += fmt.Sprintf("%s=>%s  (value:%s,gas:%s)  (input:%s)\n",call.caller.Hex(),call.callee.Hex(),call.value.Text(10),call.gas.Text(10),hex.EncodeToString(call.input))
	}
	writer.Write([]byte(str))
}
func (oracle *HackerEtherTransfer) String() string{
	return "HackerEtherTransfer"
}

/**
*Oracle: HackerEtherTransferFailed
*to help us check whether the call (with ether received or sent) fail.
*/
type HackerEtherTransferFailed struct{
	hacker_call_hashs     []common.Hash
	hacker_calls          []*HackerContractCall
	hacker_exception_calls []*HackerContractCall
}
func NewHackerEtherTransferFailed() *HackerEtherTransferFailed{
	return new(HackerEtherTransferFailed)
}
func (oracle *HackerEtherTransferFailed) InitOracle(hacker_call_hashs []common.Hash,hacker_calls []*HackerContractCall){
	oracle.hacker_calls = hacker_calls
	oracle.hacker_call_hashs = hacker_call_hashs
	oracle.hacker_exception_calls = make([]*HackerContractCall,0)
}
func (oracle *HackerEtherTransferFailed) TestOracle() bool{
	ret := false
	if oracle.triggerOracle(oracle.hacker_calls[0]){
		oracle.hacker_exception_calls = append(oracle.hacker_exception_calls,oracle.hacker_calls[0])
		ret = true
	}
	calls := oracle.hacker_calls[0].nextcalls
	for _,call := range  calls {
		if oracle.triggerOracle(call) {
			oracle.hacker_exception_calls = append(oracle.hacker_exception_calls, call)
			ret = true
		}
	}
	return ret
}
func (oracle *HackerEtherTransferFailed) triggerOracle(call *HackerContractCall) bool{

	return (call.value.Uint64()>0||strings.Contains(call.OperationStack.String(),opCodeToString[BALANCE]))&&call.throwException
}
func (oracle *HackerEtherTransferFailed) Write(writer io.Writer){
	str := ""
	if len(oracle.hacker_exception_calls) >0{
		str = "\nHackerEtherTransferFailed:\n"
	}else{
		return
	}
	for _,call := range oracle.hacker_exception_calls{
		str += fmt.Sprintf("%p\n",call)
		str +=fmt.Sprintf("profile:\n")
		str += fmt.Sprintf("%s=>%s  (value:%s,gas:%s)  (input:%s)\n",call.caller.Hex(),call.callee.Hex(),call.value.Text(10),call.gas.Text(10),hex.EncodeToString(call.input))
	}
	writer.Write([]byte(str))
}
func (oracle *HackerEtherTransferFailed) String() string{
	return "HackerEtherTransferFailed"
}
/**
* Oracle: HackerCallEtherTransferFailed.
* part of HackerEtherTransferFailed.
* to help us choose the call who send ether with gas>2300
* (2300 is the default gas limit of 'SEND')
*/
type HackerCallEtherTransferFailed struct{
	hacker_call_hashs []common.Hash
	hacker_calls []*HackerContractCall
	hacker_fallback_calls []*HackerContractCall
}
func NewHackerCallEtherTransferFailed() *HackerCallEtherTransferFailed{
	return &HackerCallEtherTransferFailed{}
}
func (oracle *HackerCallEtherTransferFailed) InitOracle(hacker_call_hashs []common.Hash,hacker_calls []*HackerContractCall){
	oracle.hacker_calls = hacker_calls
	oracle.hacker_call_hashs = hacker_call_hashs
	oracle.hacker_fallback_calls = make([]*HackerContractCall,0,0)
}
func (oracle *HackerCallEtherTransferFailed) TestOracle() bool{
	var hasCallEtherTransferFailed bool=false
	calls := oracle.hacker_calls[0].nextcalls
	for _,call := range calls{
		if true == oracle.TriggerFallbackCall(call){
			oracle.hacker_fallback_calls = append(oracle.hacker_fallback_calls, call)
			hasCallEtherTransferFailed = true
		}
	}
	return hasCallEtherTransferFailed;
}
func (oracle *HackerCallEtherTransferFailed) TriggerFallbackCall(call *HackerContractCall) bool{
	return  IsAccountAddress(call.callee)&&call.gas.Uint64()>2300&&call.throwException&&call.value.Uint64()>0
}
func (oracle *HackerCallEtherTransferFailed) Write(writer io.Writer){
	var str  string
	str = "\n HackerCallEtherTransferFailed :\n"
	if len(oracle.hacker_fallback_calls) <=0{
		return
	}
	for _,call := range oracle.hacker_fallback_calls{
		str += fmt.Sprintf("%p\n",call)
		str +=fmt.Sprintf("profile:\n")
		str += fmt.Sprintf("%s=>%s  (value:%s,gas:%s)  (input:%s)\n",call.caller.Hex(),call.callee.Hex(),call.value.Text(10),call.gas.Text(10),hex.EncodeToString(call.input))
	}
	writer.Write([]byte(str))
}
func  (oracle *HackerCallEtherTransferFailed) String() string{
	return  "HackerCallEtherTransferFailed"
}
/**
* Oracle: HackerGaslessSend
* to help us detect the call whether it sends ether via "SEND" but fails out of gas insufficiency.
*/
type HackerGaslessSend struct{
	hacker_call_hashs     []common.Hash
	hacker_calls          []*HackerContractCall
	hacker_exception_calls []*HackerContractCall
}
func NewHackerGaslessSend() *HackerGaslessSend{
	return &HackerGaslessSend{}
}
func (oracle *HackerGaslessSend) InitOracle(hacker_call_hashs []common.Hash,hacker_calls []*HackerContractCall){
	oracle.hacker_calls = hacker_calls
	oracle.hacker_call_hashs = hacker_call_hashs
	oracle.hacker_exception_calls = make([]*HackerContractCall,0,0)
}
func (oracle *HackerGaslessSend) TestOracle() bool{
	hasException := false
	calls := oracle.hacker_calls[0].nextcalls
	for _,call := range calls{
		if oracle.TriggerExceptionCall(call){
			oracle.hacker_exception_calls = append(oracle.hacker_exception_calls, call)
			hasException = true
		}
	}
	return hasException;
}
func (oracle *HackerGaslessSend) TriggerExceptionCall(call *HackerContractCall) bool{
	return  IsAccountAddress(call.callee)&&call.throwException==true&&call.errOutGas==true&&len(call.input)==0&&call.gas.Uint64()==2300
}
func (oracle *HackerGaslessSend) Write(writer io.Writer){
	str := ""
	if len(oracle.hacker_exception_calls) >0{
		str = "\nGaslessSend Exception calls:\n"
	}
	for _,call := range oracle.hacker_exception_calls{
		str += fmt.Sprintf("%p\n",call)
		str +=fmt.Sprintf("profile:\n")
		str += fmt.Sprintf("%s=>%s  (value:%s,gas:%s)  (input:%s)\n",call.caller.Hex(),call.callee.Hex(),call.value.Text(10),call.gas.Text(10),hex.EncodeToString(call.input))
	}
	writer.Write([]byte(str))
}
func  (oracle *HackerGaslessSend) String() string{
	return  "HackerGaslessSend"
}

/**
*Oracle: HackerDelegateCallInfo
* to help us to choose the call featured with 'DELEAGATECALL' initiating.
*/
type HackerDelegateCallInfo struct {
	hacker_call_hashs     []common.Hash
	hacker_calls          []*HackerContractCall
	hacker_delegate_calls []*HackerContractCall
	feautures			  []string
}
func NewHackerDelegateCallInfo() *HackerDelegateCallInfo{
	return &HackerDelegateCallInfo{}
}
func (oracle *HackerDelegateCallInfo) InitOracle(hacker_call_hashs []common.Hash,hacker_calls []*HackerContractCall){
	oracle.hacker_calls = hacker_calls
	oracle.hacker_call_hashs = hacker_call_hashs
	oracle.hacker_delegate_calls = make([]*HackerContractCall,0,0)
	oracle.feautures = make([]string,0,0)
}
func (oracle *HackerDelegateCallInfo) TestOracle() bool{
	var hasDelegate bool
	hasDelegate = false
	nextcalls := oracle.hacker_calls[0].nextcalls
	for _,call := range nextcalls {
		if oracle.TriggerDelegateCall(call) {
			oracle.hacker_delegate_calls = append(oracle.hacker_delegate_calls, call)
			hasDelegate = true
			oracle.GetFeatures(oracle.hacker_calls[0],call)
		}
	}
	return hasDelegate;
}
func (oracle *HackerDelegateCallInfo) GetFeatures(rootcall, call *HackerContractCall)  {
	if strings.EqualFold(strings.ToLower(rootcall.caller.Hex()),strings.ToLower(call.callee.Hex()))||strings.Contains(strings.ToLower(hex.EncodeToString(rootcall.input)),strings.ToLower(call.callee.Hex()[2:])) {
		oracle.feautures = append(oracle.feautures, "DANGEROUS_CALLER")
	}
	if strings.Contains(strings.ToLower(hex.EncodeToString(rootcall.input)),strings.ToLower(hex.EncodeToString(call.input))) {
		oracle.feautures = append(oracle.feautures, "DANGEROUS_INPUT")
	}
}
func (oracle *HackerDelegateCallInfo) TriggerDelegateCall(call *HackerContractCall) bool{
	return  strings.Contains(call.OperationStack.String(),opCodeToString[DELEGATECALL])
}
func (oracle *HackerDelegateCallInfo) Write(writer io.Writer){
	var str string
	str = ""
	if len(oracle.hacker_delegate_calls) >0{
		str = "\nHackerDelegateCallInfo:\n"
	}
	for _,call := range oracle.hacker_delegate_calls{
		str += fmt.Sprintf("%p\n",call)
		str += fmt.Sprintf("profile:\n")
		str += fmt.Sprintf("%s=>%s  (value:%s,gas:%s)  (input:%s)\n",call.caller.Hex(),call.callee.Hex(),call.value.Text(10),call.gas.Text(10),hex.EncodeToString(call.input))
	}
   writer.Write([]byte(str))
}
func (oracle *HackerDelegateCallInfo) String() string{
	if len(oracle.feautures)>0{
		return "HackerDelegateCallInfo:"+strings.Join(oracle.feautures," ")
	}else{
		return "delegatecallop"
	}
}

/**
* Oracle: HackerExceptionDisorder
* to help us detect whether exception happens and exception propagates bottom-up.
*/
type HackerExceptionDisorder struct{
	hacker_call_hashs     []common.Hash
	hacker_calls          []*HackerContractCall
	hacker_exception_calls []*HackerContractCall
}
func NewHackerExceptionDisorder() *HackerExceptionDisorder{
	return &HackerExceptionDisorder{}
}
func (oracle *HackerExceptionDisorder) InitOracle(hacker_call_hashs []common.Hash,hacker_calls []*HackerContractCall){
	oracle.hacker_calls = hacker_calls
	oracle.hacker_call_hashs = hacker_call_hashs
	oracle.hacker_exception_calls = make([]*HackerContractCall,0,5)
}
func (oracle *HackerExceptionDisorder) TestOracle() bool{
	exception := false
	nextcalls := oracle.hacker_calls[0].nextcalls
	for _,call := range nextcalls {
		if oracle.TriggerExceptionCall(hacker_calls[0],call) {
			oracle.hacker_exception_calls = append(oracle.hacker_exception_calls, call)
			exception = true
		}
	}
	return exception;
}
func (oracle *HackerExceptionDisorder) TriggerExceptionCall(root,call *HackerContractCall) bool{
	return root.throwException==false && IsAccountAddress(call.callee)&& call.throwException==true
}
func (oracle *HackerExceptionDisorder) Write(writer io.Writer){
	var str  string
	str = ""
	if len(oracle.hacker_exception_calls) >0{
		str = "\nExceptionDisorder calls:\n"
	}
	for _,call := range oracle.hacker_exception_calls{
		str += fmt.Sprintf("%p\n",call)
		str +=fmt.Sprintf("profile:\n")
		str += fmt.Sprintf("%s=>%s  (value:%s,gas:%s)  (input:%s)\n",call.caller.Hex(),call.callee.Hex(),call.value.Text(10),call.gas.Text(10),hex.EncodeToString(call.input))
	}
	writer.Write([]byte(str))
}
func (oracle *HackerExceptionDisorder) String() string{
	return "HackerExceptionDisorder"
}

/**
*Oracle: HackerSendOpInfo
*to help us detect the call whether cases such as "send()" works in the process.
*/
type HackerSendOpInfo struct{
	hacker_call_hashs     []common.Hash
	hacker_calls          []*HackerContractCall
	hacker_exception_calls []*HackerContractCall
}
func NewHackerSendOpInfo() *HackerSendOpInfo{
	return new(HackerSendOpInfo)
}
func (oracle *HackerSendOpInfo) InitOracle(hacker_call_hashs []common.Hash,hacker_calls []*HackerContractCall){
	oracle.hacker_calls = hacker_calls
	oracle.hacker_call_hashs = hacker_call_hashs
	oracle.hacker_exception_calls = make([]*HackerContractCall,0)
}
func (oracle *HackerSendOpInfo) TestOracle() bool{
	ret := false
	nextcalls := oracle.hacker_calls[0].nextcalls
	for _,call := range  nextcalls{
		if oracle.triggerOracle(call){
			oracle.hacker_exception_calls = append(oracle.hacker_exception_calls,call)
			ret = true
		}
	}
	return ret
}
func (oracle *HackerSendOpInfo) triggerOracle(call *HackerContractCall) bool{
	return IsAccountAddress(call.callee)&&len(call.input)==0&&call.gas.Uint64()==2300
}
func (oracle *HackerSendOpInfo) Write(writer io.Writer){
	str := ""
	if len(oracle.hacker_exception_calls) >0{
		str = "\nSendOp found:\n"
	}else{
		return
	}
	for _,call := range oracle.hacker_exception_calls{
		str += fmt.Sprintf("%p\n",call)
		str +=fmt.Sprintf("profile:\n")
		str += fmt.Sprintf("%s=>%s  (value:%s,gas:%s)  (input:%s)\n",call.caller.Hex(),call.callee.Hex(),call.value.Text(10),call.gas.Text(10),hex.EncodeToString(call.input))
	}
	writer.Write([]byte(str))
}
func (oracle *HackerSendOpInfo) String() string{
	return "HackerSendOpInfo"
}

/**
*Oracle: HackerCallOpInfo
*to help us detect the call whether cases such as "call.value(_value).gas(_gas)(_sig)" work in the process.
*/
type HackerCallOpInfo struct{
	hacker_call_hashs     []common.Hash
	hacker_calls          []*HackerContractCall
	hacker_exception_calls []*HackerContractCall
}
func NewHackerCallOpInfo() *HackerCallOpInfo{
	return new(HackerCallOpInfo)
}
func (oracle *HackerCallOpInfo) InitOracle(hacker_call_hashs []common.Hash,hacker_calls []*HackerContractCall){
	oracle.hacker_calls = hacker_calls
	oracle.hacker_call_hashs = hacker_call_hashs
	oracle.hacker_exception_calls = make([]*HackerContractCall,0)
}
func (oracle *HackerCallOpInfo) TestOracle() bool{
	ret := false
	nextcalls := oracle.hacker_calls[0].nextcalls
	for _,call := range  nextcalls{
		if oracle.triggerOracle(call){
			oracle.hacker_exception_calls = append(oracle.hacker_exception_calls,call)
			ret = true
		}
	}
	return ret
}
func (oracle *HackerCallOpInfo) triggerOracle(call *HackerContractCall) bool{
	return IsAccountAddress(call.callee)&&call.gas.Uint64()>2300
}
func (oracle *HackerCallOpInfo) Write(writer io.Writer){
	str := ""
	if len(oracle.hacker_exception_calls) >0{
		str = "\nCallOpInfo found:\n"
	}else{
		return
	}
	for _,call := range oracle.hacker_exception_calls{
		str += fmt.Sprintf("%p\n",call)
		str +=fmt.Sprintf("profile:\n")
		str += fmt.Sprintf("%s=>%s  (value:%s,gas:%s)  (input:%s)\n",call.caller.Hex(),call.callee.Hex(),call.value.Text(10),call.gas.Text(10),hex.EncodeToString(call.input))
	}
	writer.Write([]byte(str))
}
func (oracle *HackerCallOpInfo) String() string{
	return "HackerCallOpInfo"
}

/**
*Oracle: HackerCallException
*to help us detect the call whether cases such as "call.value(_value).gas(_gas)(_sig)" fails in the process.
*/

type HackerCallExecption struct{
	hacker_call_hashs     []common.Hash
	hacker_calls          []*HackerContractCall
	hacker_exception_calls []*HackerContractCall
}
func NewHackerCallExecption() *HackerCallExecption{
	return new(HackerCallExecption)
}
func (oracle *HackerCallExecption) InitOracle(hacker_call_hashs []common.Hash,hacker_calls []*HackerContractCall){
	oracle.hacker_calls = hacker_calls
	oracle.hacker_call_hashs = hacker_call_hashs
	oracle.hacker_exception_calls = make([]*HackerContractCall,0)
}
func (oracle *HackerCallExecption) TestOracle() bool{
	ret := false
	nextcalls := oracle.hacker_calls[0].nextcalls
	for _,call := range  nextcalls{
		if oracle.triggerOracle(call){
			oracle.hacker_exception_calls = append(oracle.hacker_exception_calls,call)
			ret = true
		}
	}
	return ret
}
func (oracle *HackerCallExecption) triggerOracle(call *HackerContractCall) bool{
	return IsAccountAddress(call.callee) && call.throwException == true&&call.gas.Uint64()>2300
}
func (oracle *HackerCallExecption) Write(writer io.Writer){
	str := ""
	if len(oracle.hacker_exception_calls) >0{
		str = "\nHackerCallExecption:\n"
	}else{
		return
	}
	for _,call := range oracle.hacker_exception_calls{
		str += fmt.Sprintf("%p\n",call)
		str +=fmt.Sprintf("profile:\n")
		str += fmt.Sprintf("%s=>%s  (value:%s,gas:%s)  (input:%s)\n",call.caller.Hex(),call.callee.Hex(),call.value.Text(10),call.gas.Text(10),hex.EncodeToString(call.input))
	}
	writer.Write([]byte(str))
}
func (oracle *HackerCallExecption) String() string{
	return "HackerCallException"
}

/***
*Oracle: HackerUnknownCall
*to help us detech call that are specified by message sender or message input.
*/
type HackerUnknownCall struct{
	hacker_call_hashs     []common.Hash
	hacker_calls          []*HackerContractCall
	hacker_exception_calls []*HackerContractCall
}
func NewHackerUnknowCall() *HackerUnknownCall{
	return  new(HackerUnknownCall)
}
func (oracle *HackerUnknownCall) InitOracle(hacker_call_hashs []common.Hash,hacker_calls []*HackerContractCall){
	oracle.hacker_calls = hacker_calls
	oracle.hacker_call_hashs = hacker_call_hashs
	oracle.hacker_exception_calls = make([]*HackerContractCall,0,10)
}
func (oracle * HackerUnknownCall) TestOracle() bool{
	ret := false
	nextcalls := oracle.hacker_calls[0].nextcalls
	for _,call := range  nextcalls{
		if oracle.TriggerOracle(oracle.hacker_calls[0],call){
			oracle.hacker_exception_calls = append(oracle.hacker_exception_calls,call)
			ret = true
		}
	}
	return ret
}
func (oracle *HackerUnknownCall) TriggerOracle(rootCall,call *HackerContractCall) bool{
	var input_str = string(rootCall.input)
	var callee_str = strings.ToLower(call.callee.Hex()[2:])
	return call.gas.Uint64()>2300&&(strings.EqualFold(strings.ToLower(rootCall.caller.Hex()),strings.ToLower(call.callee.Hex()))||strings.Contains(input_str,string(call.input))||strings.Contains(input_str,callee_str))
}
func (oracle *HackerUnknownCall) Write(writer io.Writer){
	str := ""
	if len(oracle.hacker_exception_calls) >0{
		str = "\nHackerUnknownAddressCall calls:\n"
	}else{
		return
	}
	for _,call := range oracle.hacker_exception_calls{
		str += fmt.Sprintf("%p\n",call)
		str +=fmt.Sprintf("profile:\n")
		str += fmt.Sprintf("%s=>%s  (value:%s,gas:%s)  (input:%s)\n",call.caller.Hex(),call.callee.Hex(),call.value.Text(10),call.gas.Text(10),hex.EncodeToString(call.input))
	}
	writer.Write([]byte(str))
}
func (oracle *HackerUnknownCall) String() string{
	return "HackerUnknownCall"
}

/**
* Oracle: HackerStorageChanged   
* to help us detect the call who change storage of the called contract.
*/
type HackerStorageChanged struct{
	hacker_call_hashs     []common.Hash
	hacker_calls          []*HackerContractCall
	hacker_exception_calls []*HackerContractCall
}
func NewHackerStorageChanged() *HackerStorageChanged{
	return  new(HackerStorageChanged)
}
func (oracle *HackerStorageChanged) InitOracle(hacker_call_hashs []common.Hash,hacker_calls []*HackerContractCall){
	oracle.hacker_calls = hacker_calls
	oracle.hacker_call_hashs = hacker_call_hashs
	oracle.hacker_exception_calls = make([]*HackerContractCall,0,10)
}
func (oracle * HackerStorageChanged) TestOracle() bool{
	ret := false
	if oracle.TriggerOracle(oracle.hacker_calls[0]){
			oracle.hacker_exception_calls = append(oracle.hacker_exception_calls,oracle.hacker_calls[0])
			ret = true
		}
	return ret
}
func (oracle *HackerStorageChanged) TriggerOracle(rootCall *HackerContractCall) bool{
	rootStorage := rootCall.StateStack
	ret,_ := rootStorage.Data()[0].Cmp(rootStorage.Data()[rootStorage.len()-1])
	return ret!=0
}
func (oracle *HackerStorageChanged) Write(writer io.Writer){
	str := ""
	if len(oracle.hacker_exception_calls) >0{
		str = "\nHackerStorageChanged calls:\n"
	}else{
		return
	}
	for _,call := range oracle.hacker_exception_calls{
		str += fmt.Sprintf("%p\n",call)
		str +=fmt.Sprintf("profile:\n")
		str += fmt.Sprintf("%s=>%s  (value:%s,gas:%s)  (input:%s)\n",call.caller.Hex(),call.callee.Hex(),call.value.Text(10),call.gas.Text(10),hex.EncodeToString(call.input))
	}
	writer.Write([]byte(str))
}
func (oracle *HackerStorageChanged) String() string{
	return "HackerStorageChanged"
}

/**
* Oracle: HackerTimestampOp
* to help us detect the call where TIMESTAMP operation works.
*/
type HackerTimestampOp struct{
	hacker_call_hashs     []common.Hash
	hacker_calls          []*HackerContractCall
	hacker_exception_calls []*HackerContractCall
}
func NewHackerTimestampOp() *HackerTimestampOp{
	return  new(HackerTimestampOp)
}
func (oracle *HackerTimestampOp) InitOracle(hacker_call_hashs []common.Hash,hacker_calls []*HackerContractCall){
	oracle.hacker_calls = hacker_calls
	oracle.hacker_call_hashs = hacker_call_hashs
	oracle.hacker_exception_calls = make([]*HackerContractCall,0,10)
}
func (oracle * HackerTimestampOp) TestOracle() bool{
	var rootCall = hacker_calls[0]
	return strings.Contains(rootCall.OperationStack.String(),opCodeToString[TIMESTAMP])
}
func (oracle *HackerTimestampOp) Write(writer io.Writer){
	str := ""
	if len(oracle.hacker_exception_calls) >0{
		str = "\nHackerTimestampOp calls:\n"
	}else{
		return
	}
	for _,call := range oracle.hacker_exception_calls{
		str += fmt.Sprintf("%p\n",call)
		str +=fmt.Sprintf("profile:\n")
		str += fmt.Sprintf("%s=>%s  (value:%s,gas:%s)  (input:%s)\n",call.caller.Hex(),call.callee.Hex(),call.value.Text(10),call.gas.Text(10),hex.EncodeToString(call.input))
	}
	writer.Write([]byte(str))
}
func (oracle *HackerTimestampOp) String() string{
	return "HackerTimestampOp"
}

/**
* Oracle: HackerNumberOp
* to help us detect the call where BLOCKNUMBER operation works.
*/

type HackerNumberOp struct{
	hacker_call_hashs     []common.Hash
	hacker_calls          []*HackerContractCall
	hacker_exception_calls []*HackerContractCall
}
func NewHackerNumberOp() *HackerNumberOp{
	return  new(HackerNumberOp)
}
func (oracle *HackerNumberOp) InitOracle(hacker_call_hashs []common.Hash,hacker_calls []*HackerContractCall){
	oracle.hacker_calls = hacker_calls
	oracle.hacker_call_hashs = hacker_call_hashs
	oracle.hacker_exception_calls = make([]*HackerContractCall,0,10)
}
func (oracle * HackerNumberOp) TestOracle() bool{
	var rootCall = hacker_calls[0]
	return strings.Contains(rootCall.OperationStack.String(),opCodeToString[NUMBER])
}
func (oracle *HackerNumberOp) Write(writer io.Writer){
	str := ""
	if len(oracle.hacker_exception_calls) >0{
		str = "\nHackerNumberOp calls:\n"
	}else{
		return
	}
	for _,call := range oracle.hacker_exception_calls{
		str += fmt.Sprintf("%p\n",call)
		str +=fmt.Sprintf("profile:\n")
		str += fmt.Sprintf("%s=>%s  (value:%s,gas:%s)  (input:%s)\n",call.caller.Hex(),call.callee.Hex(),call.value.Text(10),call.gas.Text(10),hex.EncodeToString(call.input))
	}
	writer.Write([]byte(str))
}
func (oracle *HackerNumberOp) String() string{
	return "HackerNumberOp"
}
/**
* Oracle: HackerBlockHashOp
* to help us detect the call where BLOCKHASH operation works.
*/
type HackerBlockHashOp struct{
	hacker_call_hashs     []common.Hash
	hacker_calls          []*HackerContractCall
	hacker_exception_calls []*HackerContractCall
}
func NewHackerBlockHashOp() *HackerBlockHashOp{
	return  new(HackerBlockHashOp)
}
func (oracle *HackerBlockHashOp) InitOracle(hacker_call_hashs []common.Hash,hacker_calls []*HackerContractCall){
	oracle.hacker_calls = hacker_calls
	oracle.hacker_call_hashs = hacker_call_hashs
	oracle.hacker_exception_calls = make([]*HackerContractCall,0,10)
}
func (oracle * HackerBlockHashOp) TestOracle() bool{
	var rootCall = hacker_calls[0]
	return strings.Contains(rootCall.OperationStack.String(),opCodeToString[BLOCKHASH])
}
func (oracle *HackerBlockHashOp) Write(writer io.Writer){
	str := ""
	if len(oracle.hacker_exception_calls) >0{
		str = "\nHackerBlockHashOp calls:\n"
	}else{
		return
	}
	for _,call := range oracle.hacker_exception_calls{
		str += fmt.Sprintf("%p\n",call)
		str +=fmt.Sprintf("profile:\n")
		str += fmt.Sprintf("%s=>%s  (value:%s,gas:%s)  (input:%s)\n",call.caller.Hex(),call.callee.Hex(),call.value.Text(10),call.gas.Text(10),hex.EncodeToString(call.input))
	}
	writer.Write([]byte(str))
}
func (oracle *HackerBlockHashOp) String() string{
	return "HackerBlockHashOp"
}
/**
* Oracle: HackerBalanceGtZero
* to help us detect the contract whose balance>0.
* this could help us find 'high-value' target contracts with other test oracles result. 
*/
type HackerBalanceGtZero struct{
	hacker_call_hashs     []common.Hash
	hacker_calls          []*HackerContractCall
}
func NewHackerBalanceGtZero() *HackerBalanceGtZero{
	return  new(HackerBalanceGtZero)
}
func (oracle *HackerBalanceGtZero) InitOracle(hacker_call_hashs []common.Hash,hacker_calls []*HackerContractCall){
	oracle.hacker_calls = hacker_calls
	oracle.hacker_call_hashs = hacker_call_hashs
}
func (oracle * HackerBalanceGtZero) TestOracle() bool{
	var rootCall = hacker_calls[0]
	return rootCall.StateStack.Data()[0].contracts[1].balance.Uint64()>0
}
func (oracle *HackerBalanceGtZero) Write(writer io.Writer){
	str := "\nHackerBalanceGtZero\n"
	writer.Write([]byte(str))
}
func (oracle *HackerBalanceGtZero) String() string{
	return "HackerBalanceGtZero"
}



/**
* Test oracle result report
*/
var  _reportor *HackerCallInfoReportor = nil
type HackerCallInfoReportor struct{
	hacker_call_hashs     []common.Hash
	hacker_calls          []*HackerContractCall
	callsLen 			  int
	root                  *HackerContractCall
	operationLen          int
	operationStack        *HackerOperationStack
}
func GetReportor() *HackerCallInfoReportor{
	if _reportor!=nil{
		return _reportor
	}else{
		_reportor = new(HackerCallInfoReportor)
		return  _reportor
	}
}
func  (report *HackerCallInfoReportor) Profile(hacker_call_hashs []common.Hash,hacker_calls []*HackerContractCall)  string{

	report.hacker_calls = hacker_calls
	report.hacker_call_hashs = hacker_call_hashs
	report.root = report.hacker_calls[0]
	report.operationLen = report.root.OperationStack.len()
	report.operationStack = report.root.OperationStack
	report.callsLen = len(hacker_calls)
	sig := report.root.Sig()
	//return fmt.Sprintf("Profile:{sig:%s,callsLen:%d,operationLen:%d,operationStack:%s}",sig,report.callsLen,report.operationLen,report.operationStack)
	return fmt.Sprintf("Profile:{sig:%s,callsLen:%d,operationLen:%d,operationStack:%s}",sig,report.callsLen,report.operationLen,strings.Replace(report.operationStack.String(),"\n"," ",-1))
}
