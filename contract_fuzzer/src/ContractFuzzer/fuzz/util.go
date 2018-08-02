package fuzz

import (
	"encoding/json"
	"os"
	"io/ioutil"
	"time"
	"math/rand"
	"math/big"
	"fmt"
)
var (
	count = 0
)
var(
	fuzz_rand = rand.Intn(time.Now().Nanosecond())
	s = rand.NewSource(time.Now().UnixNano()+int64(fuzz_rand))
    r = rand.New(s)
)
func randBint(max,min big.Int)(*big.Int){
	count += 1
	fuzz_rand = rand.Intn(count)
	s = rand.NewSource(time.Now().UnixNano()+int64(fuzz_rand))
	r = rand.New(s)
	return new(big.Int).Add(new(big.Int).Rand(r,new(big.Int).Sub(&max,&min)),&min)
}
func toJsonStr(v interface{})([]byte){
	buf,_:=json.Marshal(v)
	return buf
}
func readFile(path string)([]byte,error){
	file,err:=os.Open(path)
	if err!=nil{
		return nil,err
	}
	fileinfo,_:=file.Stat()
	Size :=fileinfo.Size()
	out := make([]byte,Size)
	if _,err :=file.Read(out);err==nil{
		return out,nil 
	}else{
		return nil,err
	}
}
func readDir(path string)([]string,error){
  fileinfos,err := ioutil.ReadDir(path)
  if err!=nil{
  	return nil,err
  }
  files := make([]string,0,len(fileinfos))
  for _,fileinfo := range fileinfos{
  	file := fileinfo.Name()
  	files = append(files,file)
  }
  return files,nil
}
func  Max(a,b int)int{
	if a>b{
		return a
	}
	return b;
}
func randintOne(max,min int) int{
	count += 1
	fuzz_rand = rand.Intn(count)
	s = rand.NewSource(time.Now().UnixNano()+int64(fuzz_rand))
	r = rand.New(s)
	if max-min<=0{
		return Max(max,min)
	}
	tmp := r.Intn(max-min)
	count += tmp
	return tmp+min
}
func randIntN(max,min big.Int,size int,out []big.Int){
	var bInt big.Int
	for i:=0;i<size;i++{
		n := new(big.Int)
		n.Set(bInt.Add(bInt.Rand(r,bInt.Sub(&max,&min)),&min))
		out[i] = *n
	}
}

func randintN(max,min int,size int,out []int){
	for i:=0;i<size;i++{
		n := r.Intn(max-min)+min
		out[i] = n
	}
}
type BigInt big.Int
func (bInt BigInt) String()string{
	var myInt big.Int
	var bint = big.Int(bInt)
	sign := bint.Sign()
	if sign == 0{
		return "-0x80"
	}else if sign == 1{
		return fmt.Sprintf("0x%s",bint.Text(16))
	}else{
		abs := myInt.Abs(&bint)
		return fmt.Sprintf("-0x%s",abs.Text(16))
	}
}
// const(
// 	ERR_ZERO_SIZED_SLICE = iota+101
// 	ERR_OVER_RANDOM_LIMIT
// 	ERR_UNKNOWN_COMPLEX_TYPE
// 	ERR_TYPE_NOT_FOUND
// 	ERR_FUZZ_TYPE_FAILED
// )
var errorCodeMap = map[int]string{
	101:"Zero-sized slice to handle",
	102:"Try over random time limit",
	103:"Cannot understand the complex type",
	104:"Cannot found type",
	105:"Fuzz type failed",
	106:"Abi fuzz failed out of unknown error",
	107:"Random select failed out of unknown error",
}
type Error struct{
	errorCode int 
}
func NewError(code int) Error{
	return Error{errorCode:code}
}
func (err Error) Error() string{
	return fmt.Sprintf("%d:%s",err.errorCode,errorCodeMap[err.errorCode])
}
var(
	ERR_ZERO_SIZED_SLICE = NewError(101)
	ERR_OVER_RANDOM_LIMIT = NewError(102)
	ERR_UNKNOWN_COMPLEX_TYPE = NewError(103)
	ERR_TYPE_NOT_FOUND = NewError(104)
	ERR_FUZZ_TYPE_FAILED = NewError(105)
	ERR_ABI_FUZZ_FAILED = NewError(106)
	ERR_RANDOM_SELECT_FAILED = NewError(107)
)