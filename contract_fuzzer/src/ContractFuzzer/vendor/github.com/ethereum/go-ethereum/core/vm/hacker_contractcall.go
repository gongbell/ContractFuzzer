/**
* @hacker_contractcall.go
* 1 record the contract call info at call start&over
* 2 while all contract calls triggered by one transaction finish,check oracle status.
* 3 Write corresponding  info to 0x***-UTime.log in detail
*    and append this info profile to Oracle.log
*/

package vm

import (
	"encoding/hex"
	"fmt"
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"runtime"
	"log"
	//"strings"
	"encoding/json"
	"net/http"
	"net/url"
)

var hacker_env *EVM
var hacker_call_stack *HackerContractCallStack
var hacker_call_hashs []common.Hash
var hacker_calls []*HackerContractCall

type HackerContractCall struct {
	isInitCall     bool
	caller         common.Address
	callee         common.Address
	value          big.Int
	gas            big.Int
	finalgas       big.Int
	input          []byte
	nextcalls      []*HackerContractCall
	OperationStack *HackerOperationStack
	StateStack     *HackerStateStack
	throwException  bool
	snapshotId      int
	nextRevisionId  int
}
func CallsPointerToString(calls []*HackerContractCall) string{
	if len(calls)== 0{
		return ""
	}
	var Data string
	Data = ""
	for _,call := range  calls{
		Data = Data +fmt.Sprintf("%p",call)+" "
	}
	return Data
}
func  CallsToString(calls []*HackerContractCall) string{
	if len(calls)== 0{
		return ""
	}
	var Data string
	Data = ""
	for _,call := range calls{
		tmp := call.ToString()
		Data = Data + tmp +"\n"
	}
	return Data
}
func (call *HackerContractCall) Sig() string{
	return fmt.Sprintf("{caller:'%s',callee:'%s',value:'%s',gas:'%s',input:'%s'}",call.caller.Hex(),call.callee.Hex(),call.value.Text(10),
		call.gas.Text(10),hex.EncodeToString(call.input))
}
func (call *HackerContractCall) Hash() []byte{
	var hash = make([]byte,0,0)
	for _,nextcall := range call.nextcalls{
		hash = append(hash,([]byte(nextcall.Sig()))...)
	}
	hash = append(hash,([]byte(string(call.OperationStack.len())))...)
	hash = append(hash,([]byte(call.StateStack.String()))...)
	return  hash
}
func (call *HackerContractCall) ToString() string{
	var Data string

	Data = fmt.Sprintf(""+
			"call@%p:<caller:%s,callee:%s," +
		    "value:%s," +
			"gas:%s,"+
			"finalgas:%s"+
			"\n\tlen(input):%d,input:%s> "+
			"\n\t\t},",
			call,
		    call.caller.Hex(), call.callee.Hex(),
		    call.value.Text(10),
			call.gas.Text(10),
		    call.finalgas.Text(10),
		    len(call.input), hex.EncodeToString(call.input),
			)
	return Data
}
func (call *HackerContractCall) Write(writer io.Writer) {
	var Data string
	//Data = fmt.Sprintf("%s",call)
	Data = fmt.Sprintf(""+
			"\ncall@%p:" +
		    "\n<caller:%s," +
		    "callee:%s," +
		    "value:%s,gas:%s,finalgas:%s,"+
			"\n\tlen(input):%d,input:%s> "+
			"\n\tnextcalls:{"+
			"\n\t\tlen:%d," +
			"\n\t\tcalls:[%s]" +
			"\n\t\tcalls:{%s}" +
			"\n\t\t},"+
			"\n\tOperationStack:{\n\t%s}"+
			"\n\tStateStack:{\n\t%s}>",
			call,
		    call.caller.Hex(), call.callee.Hex(),
		    call.value.Text(10),
			call.gas.Text(10),
			call.finalgas.Text(10),
		    len(call.input),
		    hex.EncodeToString(call.input),
			len(call.nextcalls),
			CallsPointerToString(call.nextcalls),
		    CallsToString(call.nextcalls),
			call.OperationStack,
			call.StateStack)
	writer.Write([]byte(Data))
}
func newHackerContractCall(operation string, caller, callee common.Address,
	value, gas big.Int, _input []byte) *HackerContractCall {
	_operationStack := newHackerOperationStack()
	_operationStack.push(operation)

	_stateStack := newHackerStateStack()
	initState := newHackerState(caller, callee)
	_stateStack.push(initState)
	nextcalls := make([]*HackerContractCall, 0)
	input := make([]byte, len(_input))
	copy(input, _input)

	return &HackerContractCall{isInitCall:false,caller: caller, callee: callee, value: value, gas: gas, input: input,
		OperationStack: _operationStack, StateStack: _stateStack, nextcalls: nextcalls,throwException:false}
}

func (call *HackerContractCall) isAncestor(callA *HackerContractCall) (bool){
	for _,childcall := range call.nextcalls{
		if childcall == callA {
			return true
		}
	}
	for _,childcall := range  call.nextcalls{
		if childcall.isAncestor(callA)==true{
			return true
		}
	}
	return false
}
func (call *HackerContractCall) findFather(index int) *HackerContractCall{
	for i:= index-1;i>=0;i--{
		if hacker_calls[i].isAncestor(call) {
			return hacker_calls[i]
		}
	}
	return nil
}
func (call *HackerContractCall) isBrother(callindex int ,callA *HackerContractCall)(bool){
	father := call.findFather(callindex)
	if father == nil{
		return false
	}
	return  father.isAncestor(callA)
	//return  !call.isAncestor(callA)&&!callA.isAncestor(call)
}
func (call *HackerContractCall) OnCall(_caller ContractRef, _callee common.Address, _value, _gas big.Int,
	_input []byte) *HackerContractCall {
	call.OperationStack.push(opCodeToString[CALL])
	call.StateStack.push(newHackerState(_caller.Address(), _callee))
	nextcall := newHackerContractCall(opCodeToString[CALL], _caller.Address(), _callee, _value, _gas, _input)
	call.nextcalls = append(call.nextcalls, nextcall)
	
	var util HackerUtils
	hash := util.Hash(nextcall)
	hacker_call_hashs= append(hacker_call_hashs,hash)
	hacker_calls = append(hacker_calls,nextcall)
	
	return nextcall
}
func (call *HackerContractCall) OnDelegateCall(_caller ContractRef, _callee common.Address, _gas big.Int,
	_input []byte) *HackerContractCall {
	call.OperationStack.push(opCodeToString[DELEGATECALL])
	call.StateStack.push(newHackerState(_caller.Address(), _callee))
	nextcall := newHackerContractCall(opCodeToString[DELEGATECALL], _caller.Address(), _callee, *new(big.Int).SetUint64(0), _gas, _input)
	call.nextcalls = append(call.nextcalls, nextcall)
	
	var util HackerUtils
	hash := util.Hash(nextcall)
	hacker_call_hashs= append(hacker_call_hashs,hash)
	hacker_calls = append(hacker_calls,nextcall)
	
	return nextcall
}
func (call *HackerContractCall) OnCallCode(_caller ContractRef, _callee common.Address,  _value,_gas big.Int,
	_input []byte) *HackerContractCall {
	call.OperationStack.push(opCodeToString[CALLCODE])
	call.StateStack.push(newHackerState(_caller.Address(), _callee))
	nextcall := newHackerContractCall(opCodeToString[CALLCODE], _caller.Address(), _callee, _value, _gas, _input)
	call.nextcalls = append(call.nextcalls, nextcall)
	
	var util HackerUtils
	hash := util.Hash(nextcall)
	hacker_call_hashs= append(hacker_call_hashs,hash)
	hacker_calls = append(hacker_calls,nextcall)
	
	return nextcall
}
func (call *HackerContractCall) OnCloseCall(finalgas big.Int) {
	call.finalgas = finalgas
	//fmt.Println("CloseCall..")
	call.OperationStack.push(opCodeToString[RETURN])
	call.StateStack.push(newHackerState(call.caller, call.callee))
	fmt.Printf("\ncall@%pClosed",call)
	
	//call.Write(hacker_writer)
}
func (call *HackerContractCall) OnBlockHash() {
	call.OperationStack.push(opCodeToString[BLOCKHASH])
	call.StateStack.push(newHackerState(call.caller, call.callee))
}
func (call *HackerContractCall) OnGas() {
	call.OperationStack.push(opCodeToString[BLOCKHASH])
	call.StateStack.push(newHackerState(call.caller, call.callee))
}
func (call *HackerContractCall) OnTimestamp() {
	call.OperationStack.push(opCodeToString[TIMESTAMP])
	call.StateStack.push(newHackerState(call.caller, call.callee))
}
func (call *HackerContractCall) OnRelationOp(relation OpCode) {
	call.OperationStack.push(opCodeToString[relation])
	call.StateStack.push(newHackerState(call.caller, call.callee))
}
func (call *HackerContractCall) OnSha3() {
	call.OperationStack.push(opCodeToString[SHA3])
	call.StateStack.push(newHackerState(call.caller, call.callee))
}
func (call *HackerContractCall) OnCreate() {
	call.OperationStack.push(opCodeToString[CREATE])
	call.StateStack.push(newHackerState(call.caller, call.callee))
}
func (call *HackerContractCall) OnAddress() {
	call.OperationStack.push(opCodeToString[ADDRESS])
	call.StateStack.push(newHackerState(call.caller, call.callee))
}
func (call *HackerContractCall) OnOrigin() {
	call.OperationStack.push(opCodeToString[ADDRESS])
	call.StateStack.push(newHackerState(call.caller, call.callee))
}
func (call *HackerContractCall) OnCaller() {
	call.OperationStack.push(opCodeToString[CALLER])
	call.StateStack.push(newHackerState(call.caller, call.callee))
}
func (call *HackerContractCall) OnDiv() {
	call.OperationStack.push(opCodeToString[DIV])
	call.StateStack.push(newHackerState(call.caller, call.callee))
}
func (call *HackerContractCall) OnBalance() {
	call.OperationStack.push(opCodeToString[BALANCE])
	call.StateStack.push(newHackerState(call.caller, call.callee))
}
func (call *HackerContractCall) OnCallValue() {
	call.OperationStack.push(opCodeToString[CALLVALUE])
	call.StateStack.push(newHackerState(call.caller, call.callee))
}
func (call *HackerContractCall) OnCalldataLoad() {
	call.OperationStack.push(opCodeToString[CALLDATALOAD])
	call.StateStack.push(newHackerState(call.caller, call.callee))
}
//Memory,Storage operation
func (call *HackerContractCall) OnMload() {
	call.OperationStack.push(opCodeToString[MLOAD])
	call.StateStack.push(newHackerState(call.caller, call.callee))
}
func (call *HackerContractCall) OnMstore() {
	call.OperationStack.push(opCodeToString[MSTORE])
	call.StateStack.push(newHackerState(call.caller, call.callee))
}
func (call *HackerContractCall) OnSload() {
	call.OperationStack.push(opCodeToString[SLOAD])
	call.StateStack.push(newHackerState(call.caller, call.callee))
}
func (call *HackerContractCall) OnSstore() {
	call.OperationStack.push(opCodeToString[SSTORE])
	call.StateStack.push(newHackerState(call.caller, call.callee))
}
//函数跳转
func (call *HackerContractCall) OnJumpi() {
	call.OperationStack.push(opCodeToString[JUMPI])
	call.StateStack.push(newHackerState(call.caller, call.callee))
}
func (call *HackerContractCall) OnJump() {
	call.OperationStack.push(opCodeToString[JUMP])
	call.StateStack.push(newHackerState(call.caller, call.callee))
}
func (call *HackerContractCall) OnSuicide() {
	call.OperationStack.push(opCodeToString[SELFDESTRUCT])
	call.StateStack.push(newHackerState(call.caller, call.callee))
}

func (call *HackerContractCall) OnNumber() {
	call.OperationStack.push(opCodeToString[NUMBER])
	call.StateStack.push(newHackerState(call.caller, call.callee))
}
func (call *HackerContractCall) OnReturn() {
	call.OperationStack.push(opCodeToString[RETURN])
	call.StateStack.push(newHackerState(call.caller, call.callee))
}
type HackerContractCallStack struct {
	data []*HackerContractCall
}

func newHackerContractCallStack() *HackerContractCallStack {
	_data := make([]*HackerContractCall, 0, 1024)
	return &HackerContractCallStack{data: _data}
}

func (st *HackerContractCallStack) Data() []*HackerContractCall {
	return st.data
}

func (st *HackerContractCallStack) push(d *HackerContractCall) {
	// NOTE push limit (1024) is checked in baseCheck
	//stackItem := new(big.Int).Set(d)
	//st.data = append(st.data, stackItem)
	st.data = append(st.data, d)
}
func (st *HackerContractCallStack) pushN(ds ...*HackerContractCall) {
	st.data = append(st.data, ds...)
}

func (st *HackerContractCallStack) pop() (ret *HackerContractCall) {
	ret = st.data[len(st.data)-1]
	st.data = st.data[:len(st.data)-1]
	return
}
func (st *HackerContractCallStack) len() int {
	return len(st.data)
}

func (st *HackerContractCallStack) swap(n int) {
	st.data[st.len()-n], st.data[st.len()-1] = st.data[st.len()-1], st.data[st.len()-n]
}

func (st *HackerContractCallStack) peek() *HackerContractCall {
	return st.data[st.len()-1]
}

// Back returns the n'th item in stack
func (st *HackerContractCallStack) Back(n int) *HackerContractCall {
	return st.data[st.len()-n-1]
}

func (st *HackerContractCallStack) require(n int) error {
	if st.len() < n {
		return fmt.Errorf("stack underflow (%d <=> %d)", len(st.data), n)
	}
	return nil
}

func (st *HackerContractCallStack) Print() {
	fmt.Println("### stack ###")
	if len(st.data) > 0 {
		for i, val := range st.data {
			fmt.Printf("%-3d  %v\n", i, val)
		}
	} else {
		fmt.Println("-- empty --")
	}
	fmt.Println("#############")
}
//var hacker_writer io.Writer = nil
//const Hacker_LOG_DIR = "/home/liuye/Ethereum/Hacker_LOG"
//var Oracle_Writer,_ =  os.OpenFile(Hacker_LOG_DIR+"/oracle.log",os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
//var logfile string
func hacker_init(evm *EVM, contract *Contract, input []byte) {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println("hacker_init")
			fmt.Println(err) // 这里的err其实就是panic传入的内容，55
		}
	}()
	if hacker_env == nil || hacker_call_stack == nil {
		hacker_env = evm
		hacker_call_stack = newHackerContractCallStack()
		hacker_call_hashs = make([]common.Hash,0,0)
		hacker_calls = make([]*HackerContractCall,0,0)
		initCall := newHackerContractCall("STARTRECORD", contract.Caller(), contract.Address(), *contract.Value(), *new(big.Int).SetUint64(contract.Gas), contract.Input)
		initCall.isInitCall = true
		hacker_call_stack.push(initCall)
		//if _, err := os.Lstat(Hacker_LOG_DIR); err != nil {
		//	os.Mkdir(Hacker_LOG_DIR, os.ModePerm)
		//}
		//if _, err := os.Lstat(Hacker_LOG_DIR + "/" + contract.Address().Hex()); err != nil {
		//	os.Mkdir(Hacker_LOG_DIR+"/"+contract.Address().Hex(), os.ModePerm)
		//}
		//logfile = fmt.Sprintf("%s/%s/%s.log",Hacker_LOG_DIR,contract.Address().Hex(),time.Now())
		//print(logfile)
		//Oracle_Writer,_ =  os.OpenFile(Hacker_LOG_DIR+"/oracle.log",os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	}

}

var transport = http.Transport{
DisableKeepAlives: false,
}
var Client = http.Client{Transport:&transport}
func hacker_close() {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
	    hacker_env = nil
		hacker_call_stack = nil
		hacker_call_hashs = nil
		hacker_calls = nil
		fmt.Println("hacker_closed!")
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容，55
			for i := 0; i < 10; i++ {
				funcName, file, line, ok := runtime.Caller(i)
				if ok {
					log.Printf("frame %v:[func:%v,file:%v,line:%v]\n", i, runtime.FuncForPC(funcName).Name(), file, line)
				}
			}
		}
	
	}()
	if hacker_env != nil || hacker_call_stack != nil {
		fmt.Println("hacker_close...")

		for ;hacker_call_stack.len()>0;{
			call:= hacker_call_stack.pop()
			//contract = call.callee
			call.OnCloseCall(*new(big.Int).SetUint64(0))
		}

		oracles :=make([]Oracle,0,0)
		oracles = append(oracles,NewHackerReentrancy())
		oracles = append(oracles,NewHackerCallEtherTransferFailed())
		oracles = append(oracles,NewHackerDelegateCallInfo())
		oracles = append(oracles,NewHackerExceptionDisorder())
        oracles = append(oracles,NewHackerGaslessSend())
        oracles = append(oracles,NewHackerCallOpInfo())
        oracles = append(oracles,NewHackerSendOpInfo())
        oracles = append(oracles,NewHackerCallExecption())
        oracles = append(oracles,NewHackerRepeatedCall())
        oracles = append(oracles,NewHackerEtherTransfer())
        oracles = append(oracles,NewHackerStorageChanged())
        oracles = append(oracles,NewHackerUnknowCall())
        oracles = append(oracles,NewHackerTimestampOp())
        oracles = append(oracles,NewHackerRootCallFailed())
        features := make([]string,0,0)
 		for _,oracle := range  oracles{
			oracle.InitOracle(hacker_call_hashs,hacker_calls)
			if true == oracle.TestOracle(){
				//file, err := os.OpenFile(logfile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
				//if err != nil {
				//	fmt.Printf("%s",err)
				//}
				//hacker_writer = file
				///*
				//* call info written to address+time.log
				//**/
				//for _,call := range  hacker_calls{
				//	call.Write(hacker_writer)
				//}
				//hacker_writer.Write([]byte("\n********hashs*****"))
				//str := fmt.Sprintf("\nid\tcall\thash")
				//hacker_writer.Write([]byte(str))
				//for i,hash := range  hacker_call_hashs{
				//	var str string
				//	str = fmt.Sprintf("\n[%d]\t%p\t%s",i,hacker_calls[i],hash.String())
				//	hacker_writer.Write([]byte(str))
				//}
				//hacker_writer.Write([]byte("\n******************************"))
				///*
				//* oracle written to Oracle.log
				//**/
				//oracle.Write(hacker_writer)
				//Oracle_Writer.Write([]byte("************************\n"))
				//Oracle_Writer.Write([]byte(logfile+":\n"))
				//oracle.Write(Oracle_Writer)
				//
				///*
				//*close file.finished
				//**/
				//file.Close()
				features = append(features,oracle.String())
			}
		}
		features_str,_:= json.Marshal(features)
		values := url.Values{"oracles":{string(features_str)},"profile":{GetReportor().Profile(hacker_call_hashs,hacker_calls)}}
		res,e := Client.Get("http://localhost:8888/hack?"+values.Encode())
		if e ==nil{
			res.Body.Close()
		}
	}
}
