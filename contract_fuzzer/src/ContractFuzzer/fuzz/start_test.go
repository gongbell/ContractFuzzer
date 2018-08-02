package fuzz

import (
	"testing"
	//"os"
)

func  Test_createADDR_MAP(t *testing.T)  {
	Global_addr_map = "../resource/addrmap.csv"
	Global_abi_sigs_dir = "/home/liuye/tested_contracts/abi_sig"
	//f, err := os.Open(Global_addr_map)
	//if err==nil{
	//	buf := bufio.NewReader(f)
	//	if line,err := buf.ReadString('\n');err==nil {
	//
	//		t.Log(line)
	//	}else{
	//		t.Fatalf("no byte read")
	//	}
	//	return
	//}
	//t.Log("hello")
	//t.Log(err)
	//if err := createADDR_MAP();err!=nil{
	//	t.Fatalf("%s",err)
	//}
	//err := createADDR_MAP()
	//if err!=nil{
	//	f, _ := os.OpenFile("error.log",os.O_CREATE|os.O_APPEND,0666)
	//	f.Write([]byte("error"))
	//	f.Close()
	//}else{
	//	f, _ := os.OpenFile("addr_map.log",os.O_CREATE|os.O_APPEND|os.O_RDWR,0666)
	//	f.Write([]byte("no error"))
	//	f.Close()
	//}
	//err = createFUNSIG_CONTRACT_MAP()
	//if err!=nil{
	//	t.Fatal(err)
	//}else{
	//	t.Log(GlobalFUNSIG_CONTRACT_MAP)
	//}
	//t.Log("ok")
	//t.Log(len(GlobalADDR_MAP))
	//
	G_current_contract = "ABCToken"
	if err := setG_current_abi_sigs();err==nil{
		t.Log(G_current_abi_sigs)
	}else{
		t.Fatal(err)
	}

}
