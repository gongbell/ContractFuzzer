package fuzz

import (
	"testing"
	//"os"
	//"encoding/json"
)

func TestSTring_fuzz(t *testing.T){
	var st = STring(String)
	out,err := st.fuzz(false)
	if err!=nil{
		t.Fatalf("%s",err)
	}
	v := out
	for _,seed := range v{
		t.Logf("seed:%s",seed)
	}
}