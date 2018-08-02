package fuzz

import "testing"

func Test_strToType(t *testing.T){
	str := "int8"
	tp,err := strToType(str)
	if err==nil{
		t.Logf("%d",tp)
	}else{
		t.Logf("%v",stringToType["int8"])
		t.Fatalf("%s",err)
	}

}