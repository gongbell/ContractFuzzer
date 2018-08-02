package fuzz

import (
	"testing"
	"os"
)

func Test_readFile(t *testing.T){
	path := "/home/liuye/GoglandProjects/autoGoFuzz/resource/intSeed.json"
	data,err :=readFile(path)
	if err!=nil{
		t.Fatalf("%s",err)
	}
	t.Logf("size:%d",len(data))
	t.Logf("data:%s",data)
	buf := make([]byte,10)
	if file,err:=os.Open(path);err==nil{
		n,err:=file.Read(buf)
		if err!=nil ||n==0{
			t.Fatalf("%s or n ==0",err)
		}
		t.Logf("%s",buf)
	}else{
		t.Fatalf("%s",err)
	}
}
func Test_readDir(t *testing.T){
	dir := "/home/liuye/PycharmProjects/DownloadContracts/verified_contract_abis"
	files,err:=readDir(dir)
	if err!=nil{
		t.Fatalf("%s",err)
	}
	t.Logf("%s",files)
	for _,file :=range files{
		path := dir+"/"+file
		t.Logf("path:%s",path)
		data,err :=readFile(path)
		if err!=nil{
			t.Fatalf("%s",err)
		}
		t.Logf("data:%s",data)
		abi,err := newAbi(data)
		if err!=nil{
			t.Fatalf("%s",err)
		}else {
			t.Logf("%s",abi)
		}
		if _,err:=abi.fuzz();err!=nil{
			t.Fatalf("%s",err)
		}
		t.Logf("%s",abi)
		break
	}
}