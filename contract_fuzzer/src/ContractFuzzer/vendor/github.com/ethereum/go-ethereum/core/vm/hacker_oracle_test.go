package vm

import (
	"testing"
	"fmt"
	"strings"
)
func EqualStr(a,b string) bool{
	if strings.Compare(a,b) ==0{
		return true
	}else{
		return false
	}
}
func Test_EqualStr(t  *testing.T ){
	var (
		a = "hello"
		b = "world"
	)
	fmt.Printf("%t",EqualStr(a,b))
	if false != EqualStr(a,b){
		t.Error("a is not equal to b, but return yes")
	}
	a = "world"
	if true != EqualStr(a,b){
		t.Error("a is  equal to b, but return false")
	}
}