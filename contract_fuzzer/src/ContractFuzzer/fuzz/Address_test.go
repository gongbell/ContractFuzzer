package fuzz

import "testing"

func TestADdress_fuzz(t *testing.T){
	var addr ADdress=ADdress(Address)
	out,err := addr.fuzz(false)
	if err!=nil{
		t.Fatalf("%s",err)
	}
	v := out
	//if err:= DYNAMIC_CAST_ERROR(ok);err!=nil{
	//	t.Fatalf("%s",err)
	//}
	t.Logf("%v",v)
}