package fuzz

import (
	"testing"
	"math/big"
)

func Test_randIntN(t *testing.T){
	tp := Int(Int8)
	out,err := tp.fuzz(false)
	if err!=nil{
		t.Fatalf("%s fuzz error %s",tp,err)
	}
	v := out
	for _,elem := range v{
			t.Logf("%s",BigInt(elem.(big.Int)))
	}
}
func Test_intseed(t *testing.T){
	//seedsFile :="/home/liuye/GoglandProjects/autoGoFuzz/resource/intseed.json"
	//jsondata := make([]byte,1024)
	//if file,err := os.Open(seedsFile);err==nil{
	//	n,e :=file.Read(jsondata)
	//	if e!=nil{
	//		t.Fatalf("fuzz failed.seed file reading error:%s",e)
	//	}else{
	//		jsondata = jsondata[:n]
	//	}
	//}else{
	//	t.Fatalf("fuzz failed.seed file open error:%s",err)
	//}
	//var myInt Int
	//myInt = Int(Int8)
	//myIntSeeds,err :=myInt.seeds(jsondata)
	//if err!=nil{
	//
	//	t.Fatalf("fuzz failed.seed file json umarshal error:%s",err)
	//}else{
	//	t.Logf("unmarsharl %v",*myIntSeeds)
	//}
	//if out,err := myInt.fuzz();err==nil{
	//	v,ok := out.([]big.Int)
	//	if ok!=true{
	//		t.Fatalf("type translate from interface{} to big.Int failed")
	//	}else{
	//		t.Logf("%v",v)
	//		for _,elem := range v{
	//			t.Logf("elem:%v",elem.Text(10))
	//		}
	//	}
	//}else{
	//	t.Fatalf("%s",err)
	//}
	//var s = rand.NewSource(time.Now().UnixNano())
	//var r = rand.New(s)
	//var bInt big.Int
	//size := 10
	//max := new(big.Int)
	//max.SetString("0x7f",0)
	//min := new(big.Int)
	//min.SetString("-0x80",0)
	//t.Logf("%s",max.Text(16))
	//t.Logf("%s",min.Text(16))
	//for i:=0;i<size;i++{
	//	n := bInt.Add(bInt.Rand(r,bInt.Sub(max,min)),min)
	//	t.Logf("%s",n.Text(10))
	//}
	//out := make([]big.Int,size)
	//randIntN(*max,*min,size,out)
	//t.Logf("%v",out)

}
