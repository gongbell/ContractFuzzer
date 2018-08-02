
package fuzz
import (
"testing"
"math/big"
)

func Test_randUintN(t *testing.T){
	tp := Uint(Uint8)
	t.Logf("size:%d",tp.size())
	out,err := tp.fuzz(false)
	if err!=nil{
		t.Fatalf("%s fuzz error %s",tp,err)
	}
	v := out
	//v,ok := out.([]big.Int)
	//if ok!=true{
	//	t.Fatalf("%s fuzz error,v not big.Int type",tp)
	//}
	for _,elem := range v{
		t.Logf("%s",BigUint(elem.(big.Int)))
	}
}