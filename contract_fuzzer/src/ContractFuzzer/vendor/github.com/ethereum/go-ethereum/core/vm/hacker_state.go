/**@hacker_state.go
* 1  Record snapshot of contract state after reach any keypoint
*       (such as operation Call,DelegateCall to be initiated);
*/
package vm

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)
func StorageToString(storage map[common.Hash]common.Hash) string {
	var Data string
	Data = "{"
	for key, value := range storage {
		Data = Data + fmt.Sprintf("\n\t<key:%s,value:%s>,", key.TerminalString(), value.Hex())
	}
	Data = Data + "\n}\n"
	return Data
}
func getStorage(_addr common.Address) map[common.Hash]common.Hash {
	storage := make(map[common.Hash]common.Hash)
	hacker_env.StateDB.ForEachStorage(_addr, func(key, value common.Hash) bool {
		storage[key] = value
		return true
	})
	return storage
}
func getBalance(_addr common.Address) big.Int {
	balance := hacker_env.StateDB.GetBalance(_addr)
	return *balance
}
type Hacker_ContractState struct {
	addr    common.Address
	storage map[common.Hash]common.Hash
	balance big.Int
}

func newHacker_ContractState(_addr common.Address) *Hacker_ContractState {
	_storage := getStorage(_addr)
	_balance := getBalance(_addr)
	return &Hacker_ContractState{addr: _addr, storage: _storage, balance: _balance}
}
func (state *Hacker_ContractState) String() string {
	return fmt.Sprintf("<addr:%s,"+
		"\n\tbalance:%s,"+
		"\n\tstorage:%s",
		state.addr.Hex(),
		state.balance.Text(10),
		StorageToString(state.storage))
}
func (state *Hacker_ContractState) Cmp(other *Hacker_ContractState) (int,string){
   s1 := state.storage
   s2 := other.storage
   if i :=state.balance.Cmp(&other.balance);i !=0{
   	    return i,fmt.Sprintf("Balance Delta: %s",new(big.Int).Sub(&state.balance,&other.balance).Text(16))
   }
   for key,_ := range  s1{
   	 if s1[key].Big().Cmp(s2[key].Big()) != 0{
   	 	i := s1[key].Big().Cmp(s2[key].Big())
  	 	str := fmt.Sprintf("Value Delta: %s",new(big.Int).Sub(s1[key].Big(),s2[key].Big()).Text(16))
   	 	return i,str
     }
   }
   return 0,fmt.Sprintf("Same")
}

type HackerState struct {
	contracts []*Hacker_ContractState
}
func (state *HackerState) Cmp(other *HackerState)(int,string){
	for index,_:= range state.contracts{
		if i,str := state.contracts[index].Cmp(other.contracts[index]);i!=0{
			return i,str
		}
  }
  return 0,fmt.Sprintf("Same")
}
func (state *HackerState) String() string {
	len := len(state.contracts)
	var Data string
	Data = fmt.Sprintf("### @contracts ###\n")
	if len > 0 {
		for i, val := range state.contracts {
			Data = Data + fmt.Sprintf("%-3d  %s\n", i, val)
		}
	} else {
		Data = Data + fmt.Sprintf("-- empty --\n")
	}
	Data = Data + fmt.Sprintf("#############\n")
	return Data
}
func newHackerState(addrs ...common.Address) *HackerState {
	size := len(addrs)
	_contracts := make([]*Hacker_ContractState, 0,size)
	for _, addr := range addrs {
		_contracts = append(_contracts, newHacker_ContractState(addr))
	}
	return &HackerState{contracts: _contracts}
}

type HackerStateStack struct {
	data []*HackerState
}
func (stack *HackerStateStack) Cmp(other *HackerStateStack)(int,string){
	fmt.Printf("storage record stack size between initState and lastState <%d,%d>",stack.len(),other.len())
	start := 0
	last := other.len()-1
	if i,str := stack.data[start].Cmp(other.data[last]);i!=0{
			return i,str
	}
	return 0,fmt.Sprintf("Same")
}
func (stack *HackerStateStack) String() string {
	len := stack.len()
	var Data string
	Data = fmt.Sprintf("\n### stack ###\n")
	if len > 0 {
		for i, val := range stack.data {
			Data = Data + fmt.Sprintf("%-3d  %s\n", i, val)
		}
	} else {
		Data = Data + fmt.Sprintf("-- empty --\n")
	}
	Data = Data + fmt.Sprintf("#############\n")
	return Data
}
func newHackerStateStack() *HackerStateStack {
	_data := make([]*HackerState, 0, 1024)
	return &HackerStateStack{data: _data}
}

func (st *HackerStateStack) Data() []*HackerState {
	return st.data
}

func (st *HackerStateStack) push(d *HackerState) {
	// NOTE push limit (1024) is checked in baseCheck
	//stackItem := new(big.Int).Set(d)
	//st.data = append(st.data, stackItem)
	st.data = append(st.data, d)
}
func (st *HackerStateStack) pushN(ds ...*HackerState) {
	st.data = append(st.data, ds...)
}

func (st *HackerStateStack) pop() (ret *HackerState) {
	ret = st.data[len(st.data)-1]
	st.data = st.data[:len(st.data)-1]
	return
}

func (st *HackerStateStack) len() int {
	return len(st.data)
}

func (st *HackerStateStack) swap(n int) {
	st.data[st.len()-n], st.data[st.len()-1] = st.data[st.len()-1], st.data[st.len()-n]
}

func (st *HackerStateStack) peek() *HackerState {
	return st.data[st.len()-1]
}

// Back returns the n'th item in stack
func (st *HackerStateStack) Back(n int) *HackerState {
	return st.data[st.len()-n-1]
}

func (st *HackerStateStack) require(n int) error {
	if st.len() < n {
		return fmt.Errorf("stack underflow (%d <=> %d)", len(st.data), n)
	}
	return nil
}

func (st *HackerStateStack) Print() {
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
