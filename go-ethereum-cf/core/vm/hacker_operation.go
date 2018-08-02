/**
*@ hacker_operation.go
* 1 store some vital operations which recorded in file @instruction.go
* 2 key unit:operation_stack
*/
package vm

import (
	"fmt"
)

//HackerOperationStack record
//opCode's String
//
type HackerOperationStack struct {
	data []string
}

func (stack *HackerOperationStack) String() string {
	len := stack.len()
	var Data string
	Data = fmt.Sprintf("\n### stack ###\n")
	if len > 0 {
		for i, val := range stack.data {
			Data = Data + fmt.Sprintf("%-3d  %v\n", i, val)
		}
	} else {
		Data = Data + fmt.Sprintf("-- empty --\n")
	}
	Data = Data + fmt.Sprintf("#############\n")
	return Data
}
func newHackerOperationStack() *HackerOperationStack {
	_data := make([]string, 0)
	return &HackerOperationStack{data: _data}
}

func (st *HackerOperationStack) Data() []string {
	return st.data
}

func (st *HackerOperationStack) push(d string) {
	// NOTE push limit (1024) is checked in baseCheck
	//stackItem := new(big.Int).Set(d)
	//st.data = append(st.data, stackItem)
	st.data = append(st.data, d)
}
func (st *HackerOperationStack) pushN(ds ...string) {
	st.data = append(st.data, ds...)
}

func (st *HackerOperationStack) pop() (ret string) {
	ret = st.data[len(st.data)-1]
	st.data = st.data[:len(st.data)-1]
	return
}

func (st *HackerOperationStack) len() int {
	return len(st.data)
}

func (st *HackerOperationStack) swap(n int) {
	st.data[st.len()-n], st.data[st.len()-1] = st.data[st.len()-1], st.data[st.len()-n]
}

func (st *HackerOperationStack) peek() string {
	return st.data[st.len()-1]
}

// Back returns the n'th item in stack
func (st *HackerOperationStack) Back(n int) string {
	return st.data[st.len()-n-1]
}

func (st *HackerOperationStack) require(n int) error {
	if st.len() < n {
		return fmt.Errorf("stack underflow (%d <=> %d)", len(st.data), n)
	}
	return nil
}

func (st *HackerOperationStack) Print() {
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
