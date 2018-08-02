package fuzz

import "testing"

func TestBOol_fuzz(t *testing.T){
	var b BOol = BOol(Bool)
	out,_:=b.fuzz(false)
	v := out
	t.Logf("%v",v)
}