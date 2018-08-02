package fuzz

import "testing"

func TestBYtes_fuzz(t *testing.T){
	var b BYtes = BYtes(Bytes)
	out,_:=b.fuzz(false)
	v := out
	t.Logf("%v",v)
}