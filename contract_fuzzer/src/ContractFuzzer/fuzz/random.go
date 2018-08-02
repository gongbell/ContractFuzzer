package fuzz
import (
	"math/rand"
	"time"
	"log"
	"runtime"
	"math/big"
)
type Fuzzer_Rand struct{
   seed int
   last_chose int 
   r  *rand.Rand 
}
func NewFuzzer_Rand() *Fuzzer_Rand{
	return &Fuzzer_Rand{seed:time.Now().Nanosecond(),last_chose:-1,r:rand.New(rand.NewSource((int64)(time.Now().Nanosecond())))}
}
func (this *Fuzzer_Rand) Random_select(bids []interface{}) (ret interface{},err error){
	defer func(){
		if err := recover(); err != nil {
			log.Println(err)
			printCallStackIfError()
			ret = nil
			err = ERR_RANDOM_SELECT_FAILED
		}
	}()
	if len(bids)==0{
		// log.Fatalf("Random_select: bids' length cannot be zero")
		// printCallStackIfError()
		// os.Exit(-1)	
		return nil,ERR_ZERO_SIZED_SLICE	
	}
	const robinTimeLimit = 10
	var count = 0
	var chose = this.last_chose
	for chose = this.r.Intn(len(bids));this.last_chose==chose&&count<robinTimeLimit;chose=this.r.Intn(len(bids)){
		count += 1
	}
	if count<robinTimeLimit{
		this.last_chose = chose
		return bids[chose],nil
	}else{
		this.last_chose = -1
		return nil,ERR_OVER_RANDOM_LIMIT
	}
}

var(
	g_func_Robin = NewFuzzer_Rand()
	g_paramval_Robin = NewFuzzer_Rand()
	g_paramval_Address_Robin = NewFuzzer_Rand()
	g_paramval_Int_Robin = NewFuzzer_Rand()
	g_paramval_Uint_Robin = NewFuzzer_Rand()
	g_paramval_Bool_Robin = NewFuzzer_Rand()
	g_paramval_Bytes_Robin = NewFuzzer_Rand()
	g_paramval_Byte_Robin = NewFuzzer_Rand()
	g_paramval_String_Robin = NewFuzzer_Rand()
	g_paramval_FixArray_Robin = NewFuzzer_Rand()
	g_paramval_DynamicArray_Robin = NewFuzzer_Rand()
	g_robin = NewFuzzer_Rand()
)
func Convert2InterfaceSlice(slice interface{}) (ret []interface{}){
	ret = make([]interface{},0)
	switch slice.(type){
	case []string:
		for _,item := range(slice.([]string)){
			ret = append(ret,item)
		}
		return ret
	case []int:
		for _,item := range(slice.([]int)){
			ret = append(ret,item)
		}
		return ret
	case []bool:
		for _,item := range(slice.([]bool)){
			ret = append(ret,item)
		}
		return ret
    default:
		return nil
	}
}

func ConvertStringSlice2InterfaceSlice(strSlice []string) []interface{}{
	var interSlice []interface{} = make([]interface{},0)
	for _,item := range(strSlice){
		interSlice = append(interSlice,item)
	}
	return interSlice
}
func ConvertIntSlice2InterfaceSlice(strSlice []int) []interface{}{
	var interSlice []interface{} = make([]interface{},0)
	for _,item := range(strSlice){
		// BigInt(elem.(big.Int)).String()
		interSlice = append(interSlice,BigInt(*new(big.Int).SetInt64(int64(item))).String())
	}
	return interSlice
}
func  printCallStackIfError(){
	for i := 0; i < 10; i++ {
		funcName, file, line, ok := runtime.Caller(i)
		if ok {
			log.Printf("frame %v:[func:%v,file:%v,line:%v]\n", i, runtime.FuncForPC(funcName).Name(), file, line)
		}
	}
}