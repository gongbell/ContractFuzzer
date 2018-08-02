/**
*  @hub.go   define data structure for recording infos
*  @Note: most parts of this file has been useless other than last function 'isRelOracle()'.
 */
package vm

import (
	//"io"
	//"math/big"
	//"os"
	//
	//"encoding/hex"
	//
	//simplejson "github.com/bitly/go-simplejson"
	"github.com/ethereum/go-ethereum/common"
)

//const (
//	WARNNING = iota
//	ERROR
//)
//const (
//	MSG_START = iota
//	MSG_CREATE
//	MSG_CALL
//	MSG_CALLER
//	MSG_DELEGATECALL
//	MSG_SEND
//	MSG_TRANSFER
//	MSG_THROW
//	MSG_ISZERO
//	MSG_JUMP
//	MSG_CRASH
//	MSG_TIMESTAMP
//	MSG_RETURN
//)
//
//type Change interface {
//	OnCall()
//	OnDelegateCall()
//	OnSend()
//	OnTransfer()
//	OnThrow()
//	OnIsZero()
//	OnJump()
//	OnCrash()
//	OnExit()
//}
//type Signature []byte
//
//func (Sig Signature) equal(other Signature) bool {
//	if len(Sig) != len(other) {
//		return false
//	} else {
//		for index, _ := range Sig {
//			if Sig[index] != other[index] {
//				return false
//			}
//		}
//		return true
//	}
//}
//
//type MSG struct {
//	Type byte
//	//feature msg signature
//	Sig  Signature
//	Pre  *MSG
//	Next *MSG
//}
//
//func (msg *MSG) Marshal() *simplejson.Json {
//	json := simplejson.New()
//	json.Set("type", string(msg.Type))
//	json.Set("sig", string(msg.Sig))
//	return json
//}
//func (msg *MSG) UnMarshal(r io.Reader) *simplejson.Json {
//	json, err := simplejson.NewFromReader(r)
//	if err != nil {
//		fmt.Errorf("%s", err)
//		return nil
//	}
//	if bytes, err := json.Get("type").Bytes(); err != nil {
//		msg.Type = bytes[0]
//	}
//	msg.Sig, _ = json.Get("sig").Bytes()
//	return json
//}
//func NewMSG(_type byte, _sig Signature) *MSG {
//	return &MSG{Type: _type, Sig: _sig}
//}
//func (msg *MSG) equal(other *MSG) bool {
//	return msg.Type == other.Type && msg.Sig.equal(other.Sig)
//}
//func (msg *MSG) pre() *MSG {
//	return msg.Pre
//}
//func (msg *MSG) next() *MSG {
//	return msg.Next
//}
//func (msg *MSG) String() string {
//	return fmt.Sprintf("<msg:%x,content:%s,%s>", msg, msg.Type, string(msg.Sig))
//}
//
//type MsgLinkList struct {
//	root *MSG
//	tail *MSG
//}
//
//func (msglist *MsgLinkList) Marshal() *simplejson.Json {
//	//json := simplejson.New()
//	rootJson := msglist.root.Marshal()
//	curJson := rootJson
//	msg := msglist.root
//	tail := msglist.tail
//	for msg != tail {
//		nextJson := msg.next().Marshal()
//		curJson.Set("next", nextJson)
//		curJson = nextJson
//	}
//	return rootJson
//}
//func NewMsgLinkList() *MsgLinkList {
//	_root := NewMSG(MSG_START, []byte("MSG_STRAT"))
//	_root.Next = _root
//	_root.Pre = _root
//	_tail := _root
//	return &MsgLinkList{root: _root, tail: _tail}
//}
//func (msglist *MsgLinkList) String() string {
//	msg := msglist.root
//	tail := msglist.tail
//	var str string
//	str += "msgList:{"
//	for msg != tail {
//		str = msg.String() + ","
//	}
//	str += "}"
//	return str
//}
//func (msglist *MsgLinkList) Add(_type byte, _sig Signature) {
//	fmt.Printf("%s,%s", string(_type), string(_sig))
//	msg := NewMSG(_type, _sig)
//	msglist.AddMSG(msg)
//}
//func (msglist *MsgLinkList) AddMSG(_msg *MSG) {
//	_msg.Pre = msglist.tail.Pre
//	_msg.Next = msglist.tail.Next
//	msglist.tail.Next = _msg
//	msglist.tail = _msg
//}
//func (msglist *MsgLinkList) HasRepeatedMSG() bool {
//	msg := msglist.root
//	for msg.next() != msglist.root {
//		if msglist.RepeatedMSG(msg) != nil {
//			return true
//		}
//	}
//	return false
//}
//func (msglist *MsgLinkList) RepeatedMSG(_msg *MSG) *MSG {
//	if _msg.next() == msglist.root {
//		return nil
//	} else {
//		msg := _msg
//		mnext := msg.next()
//		for mnext != msglist.root {
//			if mnext.equal(_msg) {
//				return mnext
//			}
//			msg = mnext
//			mnext = msg.next()
//		}
//		return nil
//	}
//}
//
//type ContractState struct {
//	ContractAddr common.Address
//	Storage      map[common.Hash]common.Hash
//	Balance      *big.Int
//}
//
//func NewContractState(addr common.Address) *ContractState {
//	storage := make(map[common.Hash]common.Hash)
//	balance := new(big.Int).SetInt64(0)
//	return &ContractState{ContractAddr: addr, Storage: storage, Balance: balance}
//}
//func (cstate *ContractState) CaptureState(env *EVM) {
//	cstate.Storage = make(map[common.Hash]common.Hash)
//	env.StateDB.ForEachStorage(cstate.ContractAddr, func(key, value common.Hash) bool {
//		cstate.Storage[key] = value
//		return true
//	})
//	cstate.Balance = env.StateDB.GetBalance(cstate.ContractAddr)
//}
//func (cstate *ContractState) WriteDiff(writer io.Writer, other *ContractState) {
//	if other == nil {
//		fmt.Errorf("other contractState is null")
//		return
//	}
//	if cstate.Balance.Cmp(other.Balance) != 0 {
//		delta := new(big.Int).Sub(cstate.Balance, other.Balance)
//		diff := "\n\t[Balance:(" + cstate.Balance.String() + "," + other.Balance.String() + ")|delta:" + delta.String() + "]"
//		fmt.Println(diff)
//		data := []byte(diff)
//		writer.Write(data)
//	}
//	for key, _ := range cstate.Storage {
//		if cstate.Storage[key].Big().Cmp(other.Storage[key].Big()) != 0 {
//			diff := "\n\t[key:" + key.String() + "," + ":(" + cstate.Storage[key].String() + "," + other.Storage[key].String() + ")]"
//			fmt.Println(diff)
//			data := []byte(diff)
//			writer.Write(data)
//		}
//	}
//}
//
//type ContractCall struct {
//	Callee               common.Address
//	Caller               ContractRef
//	Value                big.Int
//	Gas                  big.Int
//	Input                []byte
//	OldState             *ContractState
//	NewState             *ContractState
//	Path                 *MsgLinkList
//	Env                  *EVM
//	EmbeddedCall         *ContractCall
//	EmbeddedDelegateCall *ContractCall
//}
//
//func NewContractCall(_env *EVM, _caller ContractRef, _callee common.Address, _value big.Int, _gas big.Int,
//	_input []byte) *ContractCall {
//	size := len(_input)
//	input := make([]byte, size)
//	copy(input, _input)
//	return &ContractCall{Env: _env, Callee: _callee, Caller: _caller,
//		Value: _value, Gas: _gas, Input: input,
//		OldState: NewContractState(_caller.Address()),
//		NewState: NewContractState(_caller.Address()),
//		Path:     NewMsgLinkList()}
//}
//func (ccall *ContractCall) _captureOldState() {
//	ccall.OldState.CaptureState(ccall.Env)
//}
//func (ccall *ContractCall) _captureNewState() {
//	ccall.NewState.CaptureState(ccall.Env)
//}
//func (ccall *ContractCall) Check(writer io.Writer) {
//	ccall.OldState.WriteDiff(writer, ccall.NewState)
//}
//
//func (ccall *ContractCall) Start() {
//	ccall._captureOldState()
//}
//func (ccall *ContractCall) End() {
//	ccall._captureNewState()
//}
//func (ccall *ContractCall) String() string {
//	return fmt.Sprintf("\nContractCall:\n{\n\tcaller:%s, \n\tcallee:%s, \n\tvalue:%s,\tgas:%s, \n\tinput:%s, \n}", ccall.Caller.Address().Hex(), ccall.Callee.Hex(), ccall.Value.Text(10), ccall.Gas.Text(10), hex.EncodeToString(ccall.Input))
//}
//
//type ChangeSubscriber struct {
//	ccall   *ContractCall
//	checker func(writer io.Writer)
//}
//
//func NewChangeSubscriber(_ccall *ContractCall) *ChangeSubscriber {
//	return &ChangeSubscriber{ccall: _ccall, checker: _ccall.Check}
//}
//func (subscriber *ChangeSubscriber) Open() {
//	fmt.Printf("open...")
//	subscriber.ccall.Start()
//}
//func (subscriber *ChangeSubscriber) Close() {
//	fmt.Printf("close...")
//	subscriber.ccall.Start()
//}
//func (subscriber *ChangeSubscriber) OnCreate() {
//	subscriber.ccall.Path.Add(MSG_CREATE, []byte("create"))
//}
//func (subscriber *ChangeSubscriber) OnCall() {
//	subscriber.ccall.Path.Add(MSG_CALL, []byte("call"))
//}
//func (subscriber *ChangeSubscriber) OnCallCode() {
//	subscriber.ccall.Path.Add(MSG_CALL, []byte("callcode"))
//}
//func (subscriber *ChangeSubscriber) OnDelegateCall() {
//	subscriber.ccall.Path.Add(MSG_DELEGATECALL, []byte("delegateCall"))
//}
//func (subscriber *ChangeSubscriber) OnSend() {
//	subscriber.ccall.Path.Add(MSG_SEND, []byte("Send"))
//}
//func (subscriber *ChangeSubscriber) OnTransfer() {
//	subscriber.ccall.Path.Add(MSG_TRANSFER, []byte("Transfer"))
//}
//func (subscriber *ChangeSubscriber) OnThrow() {
//	subscriber.ccall.Path.Add(MSG_THROW, []byte("Throw"))
//}
//func (subscriber *ChangeSubscriber) OnIsZero() {
//	subscriber.ccall.Path.Add(MSG_ISZERO, []byte("IsZero"))
//}
//func (subscriber *ChangeSubscriber) OnJump() {
//	subscriber.ccall.Path.Add(MSG_JUMP, []byte("Jump"))
//}
//func (subscriber *ChangeSubscriber) OnJumpI() {
//	subscriber.ccall.Path.Add(MSG_JUMP, []byte("Jump"))
//}
//func (subscriber *ChangeSubscriber) OnReturn() {
//	subscriber.ccall.Path.Add(MSG_RETURN, []byte("timestamp"))
//}
//func (subscriber *ChangeSubscriber) OnCaller() {
//	subscriber.ccall.Path.Add(MSG_CALLER, []byte("caller"))
//}
//func (subscriber *ChangeSubscriber) OnTimstamp() {
//	subscriber.ccall.Path.Add(MSG_TIMESTAMP, []byte("timestamp"))
//}
//func (subscriber *ChangeSubscriber) OnCrash() {
//	subscriber.ccall.Path.Add(MSG_CRASH, []byte("Crash"))
//}
//
//func (subscriber *ChangeSubscriber) Write() {
//	file, err := os.OpenFile("hacker.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
//	if err != nil {
//		fmt.Errorf("%s", err)
//		return
//	}
//	var writer io.Writer
//	writer = file
//	str := subscriber.ccall.String()
//	writer.Write([]byte(str))
//	subscriber.checker(writer)
//	file.Close()
//}
//
//var (
//	rootCall   *ContractCall     = nil
//	ccall      *ContractCall     = nil
//	subscriber *ChangeSubscriber = nil
//)
//
//func ExistCCall() bool {
//	return ccall != nil
//}
//func GetCCall() (*ContractCall, error) {
//	if ccall == nil {
//		return nil, fmt.Errorf("ccall is null")
//	} else {
//		return ccall, nil
//	}
//}
//func CreateAndGetCCall(_env *EVM, _caller ContractRef, _callee common.Address, _input []byte, _gas uint64, _value *big.Int, isDelegateCall bool) *ContractCall {
//	if ccall == nil {
//		ccall = CreateCCall(_env, _caller, _callee, *_value, *new(big.Int).SetUint64(_gas), _input)
//		rootCall = ccall
//	} else if ccall.Callee != _callee || ccall.Caller == _caller {
//		if isDelegateCall == true {
//			newcall := CreateCCall(_env, _caller, _callee, *_value, *new(big.Int).SetUint64(_gas), _input)
//			ccall.EmbeddedDelegateCall = newcall
//			ccall = newcall
//		} else {
//			newcall := CreateCCall(_env, _caller, _callee, *_value, *new(big.Int).SetUint64(_gas), _input)
//			ccall.EmbeddedCall = newcall
//			ccall = newcall
//		}
//	}
//	return ccall
//}
//func CreateCCall(_env *EVM, _caller ContractRef, _callee common.Address, _value big.Int, _gas big.Int,
//	_input []byte) *ContractCall {
//	ccall = NewContractCall(_env, _caller, _callee, _value, _gas, _input)
//	subscriber = NewChangeSubscriber(ccall)
//	return ccall
//}
//func MarkCall() {
//	subscriber.OnCall()
//}
//func MarkDelegateCall() {
//	subscriber.OnDelegateCall()
//}
//func MarkJump() {
//	subscriber.OnJump()
//}
//func MarkSend() {
//	subscriber.OnSend()
//}
//func MarkTransfer() {
//	subscriber.OnTransfer()
//}
//func MarkThrow() {
//	subscriber.OnThrow()
//}
//func MarkIsZero() {
//	subscriber.OnIsZero()
//}
//func MarkCrash() {
//	subscriber.OnCrash()
//}
//func MarkExit() {
//	subscriber.OnReturn()
//}

var relOracle = common.HexToAddress("0xfa7b9770ca4cb04296cac84f37736d4041251cdf")

func isRelOracle(addr common.Address) bool {
	//fmt.Printf("addr:%s", relOracle.Hex())
	if addr.Big().Cmp(relOracle.Big()) == 0 {
		//fmt.Printf("\nrel:%s, consistent", addr.Hex())
		return true
	} else {
		//fmt.Printf("\nnon-consistent")
		//fmt.Printf("addr:%s,%s nonconsistent", addr.Hex(), relOracle.Hex())
		return false
	}
}
