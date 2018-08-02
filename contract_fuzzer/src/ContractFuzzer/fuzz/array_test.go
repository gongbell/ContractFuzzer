package fuzz

import (
	"testing"
)
func Test_getInfo(t *testing.T){
	str := "address[5][]"
	v,err := getInfo(str)
	if err !=nil{
		t.Fatalf("%s",err)
	}else{
		switch v {
		case Cfundemental:
				{
					t.Logf("Fundemental")
					f,_ := strToType(str)
					t.Logf("%s",f)
					isElem := false
					out,_:= f.fuzz(isElem)
					t.Logf("%v",out)
				}
		case CfixedArray:
				{
					t.Logf("FixedArray")
					f := newFixedArray(str)
					out,_ := f.fuzz()
					t.Logf("%v",out)
				}
		case CdynamicArray:
				{
					t.Logf("DynamicArray")
					d := newDynamicArray(str)
					out,_ := d.fuzz()
					t.Logf("%s", out)
				}
		default:
			t.Fatalf("type:%s couldn't be contrusted",str)
		}
		}
	//f,_ := strToType(str)
	//var myString = STring(f)
	//out,err := myString.fuzz(false)
	//if err!=nil{
	//	t.Fatalf("%s",err)
	//}
	//val,ok := out.([]string)
	//e := DYNAMIC_CAST_ERROR(ok)
	//if e!=nil{
	//	t.Fatalf("%s",err)
	//}
	//t.Logf("%s",val)
}
