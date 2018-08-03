// package  main
package server

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"os"
	"strings"
	"runtime"
	"io"
	"bufio"
	"sync"
	"strconv"
	"time"
	fuzz "ContractFuzzer/fuzz"
)
var(
	logWriter *os.File
	countWriter *os.File
	count_output *os.File
	receive_count_writer *os.File
	reen_writer  *os.File
	except_disorder_writer *os.File
	delegate_writer *os.File
	gasless_writer *os.File
	timedependency_writer *os.File
	numberdependency_writer *os.File
	freezingether_writer *os.File
	addr_map_file string 
)
func init_file_rw(addr_map string, reporter string){
	addr_map_file = addr_map
	logfile := reporter+"/log.txt"
	countfile := reporter+"/count.txt"
	contract_output_file := reporter+"/contract_fun_vulnerabilities.txt"
	receive_count_file := reporter+"/receive_count.txt"
	logWriter,_ = os.OpenFile(logfile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
    countWriter,_=os.OpenFile(countfile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
    count_output,_  = os.OpenFile(contract_output_file, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
    receive_count_writer,_ = os.OpenFile(receive_count_file, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	reen_writer,_ = os.OpenFile(reporter+"/bug/reentrancy_danger.list",os.O_CREATE|os.O_APPEND|os.O_RDWR,0666)
	except_disorder_writer,_ = os.OpenFile(reporter+"/bug/exception_disorder.list",os.O_CREATE|os.O_APPEND|os.O_RDWR,0666)
	delegate_writer,_ = os.OpenFile(reporter+"/bug/delegate_danger.list",os.O_CREATE|os.O_APPEND|os.O_RDWR,0666)
	gasless_writer,_ = os.OpenFile(reporter+"/bug/gasless_send.list",os.O_CREATE|os.O_APPEND|os.O_RDWR,0666)
	timedependency_writer,_ = os.OpenFile(reporter+"/bug/time_dependency.list",os.O_CREATE|os.O_APPEND|os.O_RDWR,0666)
	numberdependency_writer,_ = os.OpenFile(reporter+"/bug/number_dependency.list",os.O_CREATE|os.O_APPEND|os.O_RDWR,0666)
	freezingether_writer,_ = os.OpenFile(reporter+"/bug/freezing_ether.list",os.O_CREATE|os.O_APPEND|os.O_RDWR,0666)
}
// const (
// 	addr_map_file = "/list/config/addrmap.csv"
// 	logfile = "/list/reporter/log.txt"
// 	countfile = "/list/reporter/count.txt"
// 	contract_output_file = "/list/reporter/contract_fun_vulnerabilities.txt"
// 	receive_count_file = "/list/reporter/receive_count.txt"
// )
// var logWriter,_ = os.OpenFile(logfile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
// var countWriter,_=os.OpenFile(countfile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
// var count_output,_  = os.OpenFile(contract_output_file, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
// var receive_count_writer,_ = os.OpenFile(receive_count_file, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
// var (
// 	reen_writer,_ = os.OpenFile("/list/reporter/bug/reentrancy_danger.list",os.O_CREATE|os.O_APPEND|os.O_RDWR,0666)
// 	except_disorder_writer,_ = os.OpenFile("/list/reporter/bug/exception_disorder.list",os.O_CREATE|os.O_APPEND|os.O_RDWR,0666)
// 	delegate_writer,_ = os.OpenFile("/list/reporter/bug/delegate_danger.list",os.O_CREATE|os.O_APPEND|os.O_RDWR,0666)
// 	gasless_writer,_ = os.OpenFile("/list/reporter/bug/gasless_send.list",os.O_CREATE|os.O_APPEND|os.O_RDWR,0666)
// 	timedependency_writer,_ = os.OpenFile("/list/reporter/bug/time_dependency.list",os.O_CREATE|os.O_APPEND|os.O_RDWR,0666)
// 	numberdependency_writer,_ = os.OpenFile("/list/reporter/bug/number_dependency.list",os.O_CREATE|os.O_APPEND|os.O_RDWR,0666)
// 	freezingether_writer,_ = os.OpenFile("/list/reporter/bug/freezing_ether.list",os.O_CREATE|os.O_APPEND|os.O_RDWR,0666)
// )
type Result struct{
	Ret int
	Reason string
	Data interface{}
}

var(
	GlobalADDR_MAP = make(map[string]string,0)
)
var m *sync.Mutex
var timeLock *sync.Mutex
var (
	cMutex *sync.Mutex
	receive_count = 0
)
var (
	decMutex *sync.Mutex
	group_count int = 0
)
const MAX_TIME uint64= 10*60 // 2.5*60
var ( 
	start_time uint64= uint64(time.Now().Unix())
	time_bound uint64= start_time+ MAX_TIME
)
const MAX_DURATION uint64 = 5
var (
	//if unchanged for two 
	unchanged_duration uint64 = MAX_DURATION  
)

const (
	CALLFAILED = "HackerRootCallFailed"
	REENTRANCY  = "HackerReentrancy"
	REPEATED   = "HackerRepeatedCall"
	ETHERTRANSFER = "HackerEtherTransfer"
	ETHERTRANSFERFAILED = "HackerEtherTransferFailed"
	CALLETHERETRANSFERFAILED = "HackerCallEtherTransferFailed"
	GASLESSSEND = "HackerGaslessSend"
	DELEGATE = "HackerDelegateCallInfo"
	EXCEPTIONDISORDER = "HackerExceptionDisorder"
	SENDOP = "HackerSendOpInfo"
	CALLOP = "HackerCallOpInfo"
	CALLEXCEPTION = "HackerCallException"
	UNKNOWCALL = "HackerUnknownCall"
	STORAGECHANGE = "HackerStorageChanged"
	TIMESTAMP = "HackerTimestampOp"
	BLOCKHAHSH = "HackerBlockHashOp"
	BLOCKNUMBER = "HackerNumberOp"
	FREEZINGETHER = "HackerFreezingEther"
)
var (
	CallFailedCount int = 0
	StorageChangedCount int= 0
	CallOpCount int = 0
	CallExceptionCount int = 0
	ExceptionDisorderCount int = 0
	EtherTransferCount int =0
	EtherTransferFailedCount int = 0
	DelegateCount int =0
	GaslessSendCount int =0
	ReentrancyCount int = 0
	CallEtherFailedCount int = 0
	RepeatedCallCount int = 0
	TimestampCount int = 0
	BlockHashCount int = 0
	BlockNumberCount int = 0
	SendOpCount int = 0
	UnknowCallCount int = 0
	FreezingEtherCount int = 0
)
func profile_unchanged(oldprofile, profile string ) bool{
	return strings.EqualFold(oldprofile,profile)==true
}
func OutputJson(w http.ResponseWriter, ret int, reason string, i interface{}) {
	out := &Result{ret, reason, i}
	b, err := json.Marshal(out)
	if err != nil {
		return
	}
	w.Write(b)
}

var(
	contract string = " "
	oldcontract string = ""
	result =make(map[string]string)
	profile = ""
	oldprofile = ""
	isCallFailed bool = false
	isStorageChanged bool = false
	isCallOp bool = false
	isCallException bool= false
	isExceptionDisorder bool = false
	isEtherTransfer bool = false
	isEtherTransferFailed bool = false
	isDelegate bool = false
	isGaslessSend bool = false
	isReentrancy bool = false
	isCallEtherFailed bool = false
	isRepeatedCall  bool = false
	isTimestamp  bool = false
	isBlockHash  bool = false
	isBlockNumber bool = false
	isSendOp bool = false
	isUnknowCall bool = false
	// isFreezingEther bool = false
)
const (

	Danger_reentrancy = "DR"
	Danger_exception_disorder="DE"
	Danger_delegate="DD"
	Danger_gasless_send ="DGS"
	Danger_timestampdependency = "DT"
	Danger_numberdependency = "DN"
	// Danger_freezingether = "DF"
)
const (
	ownerAccount = "0xA31A0f4653f62aCa35B6e986743C8F4Fc6c8F38D"
	agentAccount = "0xA31A0f4653f62aCa35B6e986743C8F4Fc6c8F38a"
	normalAccount = "0xA31A0f4653f62aCa35B6e986743C8F4Fc6c8F38F"
)

func importAddr_ContractMap(mapfile string) error {
	f, err := os.Open(mapfile)
	defer func() {
		f.Close()
	}()
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		if err!=nil{
			f, _ := os.OpenFile("error.log",os.O_CREATE|os.O_APPEND,0666)
			f.Write([]byte(err.Error()))
			f.Close()
		}

		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		if line == "" {
			return nil
		}
		line = strings.TrimSpace(line)
		str2 := strings.Split(line,",")
		addr := strings.TrimSpace(str2[0])
		name := strings.TrimSpace(str2[1])
		GlobalADDR_MAP[strings.ToLower(addr)] = name
	}
	return nil

}
func Init()  {

	isCallFailed = false
	isStorageChanged  = false
	isCallOp  = false
	isCallException = false
	isExceptionDisorder  = false
	isEtherTransfer  = false
	isEtherTransferFailed  = false
	isDelegate  = false
	isGaslessSend  = false
	isReentrancy  = false
	isCallEtherFailed  = false
	isRepeatedCall  = false
	isTimestamp   = false
	isBlockHash   = false
	isBlockNumber  = false
	isSendOp  = false
	isUnknowCall  = false
	// isFreezingEther = false
}

func Output()  {
	if profile==""{
		return
	}
	packet_split := "\n***********************"
    count_output.Write([]byte(packet_split))
	// packet_info := "\n"+contract
	packet_info := "\n"+GlobalADDR_MAP[strings.ToLower(contract[1:len(contract)-1])]
    packet_info += "\n"+profile
    for fun := range result{
    	packet_info += "\n"+ fun+": "+result[fun]
	}
	if len(result)>0{
		count_output.Write([]byte(packet_info))
		log.Println(packet_info)
	}
	if strings.Contains(profile,Danger_reentrancy){
		reen_writer.Write([]byte(GlobalADDR_MAP[strings.ToLower(contract[1:len(contract)-1])]))
		reen_writer.Write([]byte("\n"))
	}
	if strings.Contains(profile,Danger_exception_disorder){
		except_disorder_writer.Write([]byte(GlobalADDR_MAP[strings.ToLower(contract[1:len(contract)-1])]))
		except_disorder_writer.Write([]byte("\n"))
	}
	if strings.Contains(profile,Danger_delegate){
		delegate_writer.Write([]byte(GlobalADDR_MAP[strings.ToLower(contract[1:len(contract)-1])]))
		delegate_writer.Write([]byte("\n"))
	}
	if strings.Contains(profile,Danger_gasless_send){
		gasless_writer.Write([]byte(GlobalADDR_MAP[strings.ToLower(contract[1:len(contract)-1])]))
		gasless_writer.Write([]byte("\n"))
	}
	if strings.Contains(profile,Danger_timestampdependency){
		timedependency_writer.Write([]byte(GlobalADDR_MAP[strings.ToLower(contract[1:len(contract)-1])]))
		timedependency_writer.Write([]byte("\n"))
	}
	if strings.Contains(profile,Danger_numberdependency){
		numberdependency_writer.Write([]byte(GlobalADDR_MAP[strings.ToLower(contract[1:len(contract)-1])]))
		numberdependency_writer.Write([]byte("\n"))
	}
	// if strings.Contains(profile,Danger_freezingether){
	// 	freezingether_writer.Write([]byte(GlobalADDR_MAP[strings.ToLower(contract[1:len(contract)-1])]))
	// 	freezingether_writer.Write([]byte("\n"))
	// }
}

func isOwner(caller string) bool{
	if strings.ToLower(caller) == strings.ToLower(ownerAccount){
		return  true
	}
	return  false
}
func isAgent(caller string) bool{
	if strings.ToLower(caller) == strings.ToLower(agentAccount){
		return true
	}
	return  false
}
func isNormal(caller string) bool{
	if strings.ToLower(caller) == strings.ToLower(normalAccount){
		return true
	}
	return false
}
func hackCountFunc(caller,callee,value,input string){
	fun  := input
	// fmt.Println(input)
	if len(input)>=8{
		fun = input[:8]
	}
	if isReentrancy && (isStorageChanged||isEtherTransfer||isSendOp){
		if dangers,found:=result[fun];found{
			if !strings.Contains(dangers,Danger_reentrancy){
				result[fun] = result[fun]+" "+Danger_reentrancy
			}
		}else{
			result[fun] = Danger_reentrancy
		}
		if !strings.Contains(profile,Danger_reentrancy){
			profile += " "+ Danger_reentrancy
		}
	}
	if isExceptionDisorder{
		if dangers,found:=result[fun];found{
			if !strings.Contains(dangers,Danger_exception_disorder){
				result[fun] = result[fun]+" "+Danger_exception_disorder
			}
		}else{
			result[fun] = Danger_exception_disorder
		}
		if !strings.Contains(profile,Danger_exception_disorder){
			profile += " "+ Danger_exception_disorder
		}
	}
	if isDelegate{
		if dangers,found:=result[fun];found{
			if !strings.Contains(dangers,Danger_delegate){
				result[fun] = result[fun]+" "+Danger_delegate
			}
		}else{
			result[fun] = Danger_delegate
		}
		if !strings.Contains(profile,Danger_delegate){
			profile += " "+ Danger_delegate
		}
	}
	if isGaslessSend{
		if dangers,found:=result[fun];found{
			if !strings.Contains(dangers,Danger_gasless_send){
				result[fun] = result[fun]+" "+Danger_gasless_send
			}
		}else{
			result[fun] = Danger_gasless_send
		}
		if !strings.Contains(profile,Danger_gasless_send){
			profile += " "+ Danger_gasless_send
		}
	}
	if isTimestamp &&(isStorageChanged||isEtherTransfer||isSendOp){
		if dangers,found:=result[fun];found{
			if !strings.Contains(dangers,Danger_timestampdependency){
				result[fun] = result[fun]+" "+Danger_timestampdependency
			}
		}else{
			result[fun] = Danger_timestampdependency
		}
		if !strings.Contains(profile,Danger_timestampdependency){
			profile += " "+ Danger_timestampdependency
		}
	}
	if isBlockNumber && (isStorageChanged||isEtherTransfer||isSendOp){
		if dangers,found:=result[fun];found{
			if !strings.Contains(dangers,Danger_numberdependency){
				result[fun] = result[fun]+" "+Danger_numberdependency
			}
		}else{
			result[fun] = Danger_numberdependency
		}
		if !strings.Contains(profile,Danger_numberdependency){
			profile += " "+ Danger_numberdependency
		}
	}
	// if isFreezingEther{
	// 	if dangers,found:=result[fun];found{
	// 		if !strings.Contains(dangers,Danger_freezingether){
	// 			result[fun] = result[fun]+" "+Danger_freezingether
	// 		}
	// 	}else{
	// 		result[fun] = Danger_freezingether
	// 	}
	// 	if !strings.Contains(profile,Danger_freezingether){
	// 		profile += " "+ Danger_freezingether
	// 	}
	// }
}
func hackCount(oracles,profiles string){
	caller := strings.Split(strings.Split(profiles,",")[0],"caller:")[1]
	new_contract := strings.Split(strings.Split(profiles,",")[1],"callee:")[1]
	value := strings.Split(strings.Split(profiles,",")[2],"value:")[1]
	input := strings.Split(strings.Split(strings.Split(profiles,",")[4],":")[1],"}")[0]
	//fmt.Println(caller,new_contract,value,input)
	if new_contract!=contract{
		Output()
		contract = new_contract
		result =make(map[string]string)
		profile = ""
	}
	Init()
	if strings.Contains(oracles,CALLFAILED){
		isCallFailed = true
	}
	if strings.Contains(oracles,REENTRANCY){
		isReentrancy = true
	}
	if strings.Contains(oracles,REPEATED){
		isRepeatedCall = true
	}
	if strings.Contains(oracles,ETHERTRANSFER){
		isEtherTransfer = true
	}
	if strings.Contains(oracles,ETHERTRANSFERFAILED){
		isEtherTransferFailed = true
	}
	if strings.Contains(oracles,CALLETHERETRANSFERFAILED){
        isCallEtherFailed = true
	}
	if strings.Contains(oracles,GASLESSSEND){
		isGaslessSend = true
	}
	if strings.Contains(oracles,DELEGATE){
		//isGaslessSend = true; ERROR：校正gaslessSend数据
		isDelegate = true
	}

	if strings.Contains(oracles,EXCEPTIONDISORDER){
		isExceptionDisorder = true
	}
	if strings.Contains(oracles,SENDOP){
		isSendOp = true
	}
	if strings.Contains(oracles,CALLOP){
		isCallOp = true
	}
	if strings.Contains(oracles,CALLEXCEPTION){
		isCallException = true
	}
	if strings.Contains(oracles,UNKNOWCALL){
		isUnknowCall = true
	}
	if strings.Contains(oracles,STORAGECHANGE){
		isStorageChanged = true
	}
	if strings.Contains(oracles,TIMESTAMP){
		isTimestamp = true
	}
	if strings.Contains(oracles,BLOCKHAHSH){
		isBlockHash = true
	}
	if strings.Contains(oracles,BLOCKNUMBER){
		isBlockNumber = true
	}
	// if strings.Contains(oracles,FREEZINGETHER){
	// 	isFreezingEther = true
	// }
	hackCountFunc(caller,contract,value,input)
}
func hackHandler(w http.ResponseWriter, r *http.Request) {
	defer func(){
		if err := recover();err!=nil{
			log.Println(err)
		}else{
			r.Body.Close()
		}
	}()
	cMutex.Lock()
	receive_count ++
	cMutex.Unlock()
	// fmt.Println(receive_count)
    receive_count_writer.Write([]byte(strconv.Itoa(receive_count)))
    //receive_count_writer.Write([]byte(str))
	//fmt.Println("GET params were:", r.URL.Query());
	oracles := r.URL.Query().Get("oracles")
	profile := r.URL.Query().Get("profile")
	if strings.Contains(oracles,CALLFAILED){
		CallFailedCount++
	}
	if strings.Contains(oracles,REENTRANCY){
		ReentrancyCount++
	}
	if strings.Contains(oracles,REPEATED){
		RepeatedCallCount++
	}
	if strings.Contains(oracles,ETHERTRANSFER){
		EtherTransferCount++
	}
	if strings.Contains(oracles,ETHERTRANSFERFAILED){
		EtherTransferFailedCount++
	}
	if strings.Contains(oracles,CALLETHERETRANSFERFAILED){
		CallEtherFailedCount++
	}
	if strings.Contains(oracles,GASLESSSEND){
		GaslessSendCount++
	}
	if strings.Contains(oracles,DELEGATE){
		DelegateCount++
	}
	if strings.Contains(oracles,EXCEPTIONDISORDER){
		ExceptionDisorderCount++
	}
	if strings.Contains(oracles,SENDOP){
		SendOpCount++
	}
	if strings.Contains(oracles,CALLOP){
		CallOpCount++
	}
	if strings.Contains(oracles,CALLEXCEPTION){
		CallExceptionCount++
	}
	if strings.Contains(oracles,UNKNOWCALL){
		UnknowCallCount++
	}
	if strings.Contains(oracles,STORAGECHANGE){
    	StorageChangedCount ++
	}
	if strings.Contains(oracles,TIMESTAMP){
		TimestampCount ++
	}
    if strings.Contains(oracles,BLOCKHAHSH){
    	BlockHashCount ++
	}
	if strings.Contains(oracles,BLOCKNUMBER){
		BlockNumberCount ++
	}
	if strings.Contains(oracles,FREEZINGETHER){
		FreezingEtherCount ++
	}
	countList := fmt.Sprintf("%4d %4d %4d %4d %4d %4d %4d %4d %4d %4d %4d %4d %4d %4d %4d %4d %4d %4d\n ",CallFailedCount,StorageChangedCount,CallOpCount,CallExceptionCount,ExceptionDisorderCount,EtherTransferCount,EtherTransferFailedCount, DelegateCount,GaslessSendCount,ReentrancyCount,CallEtherFailedCount,RepeatedCallCount,TimestampCount,BlockHashCount,BlockNumberCount,SendOpCount,UnknowCallCount,FreezingEtherCount)
	countWriter.Write([]byte(countList))
	msg :=fmt.Sprintf("\n%s  %s\n",r.URL.Query().Get("oracles"),r.URL.Query().Get("profile"))
	logWriter.Write([]byte(msg))
	w.Write([]byte(fmt.Sprintf("GET params were:%s", r.URL.Query())))
	m.Lock()
	hackCount(oracles,profile)
	m.Unlock()
	decMutex.Lock()
	group_count = group_count - 1
	log.Println(group_count)
	if group_count <= 0{
		// fuzz.G_stop <- false
		if profile_unchanged(oldprofile,profile){
			unchanged_duration -= 1
			if unchanged_duration == 0{
				fuzz.G_stop<-true
			}else{
				fuzz.G_stop<-false
			}
		}else{
			fuzz.G_stop<-false
			unchanged_duration = MAX_DURATION
		}
		oldprofile = profile
	}
	decMutex.Unlock()
}

func Start(addr_map string, reporter string) {
	init_file_rw(addr_map,reporter);
	importAddr_ContractMap(addr_map_file)
	cMutex = new(sync.Mutex)
	m = new(sync.Mutex)
	decMutex = new(sync.Mutex)
	timeLock = new(sync.Mutex)
	defer func(){
		if err := recover(); err != nil {
			log.Println(err) // 这里的err其实就是panic传入的内容，55
			for i := 0; i < 10; i++ {
				funcName, file, line, ok := runtime.Caller(i)
				if ok {
					log.Printf("frame %v:[func:%v,file:%v,line:%v]\n", i, runtime.FuncForPC(funcName).Name(), file, line)
				}
			}
		}
	}()
	go func(){
		for true{
			log.Println(time.Now().Unix(),time_bound)
			// log.Println(time_bound)
			if uint64(time.Now().Unix())>time_bound{
				log.Println("TIMEOUT, stop.")
				fuzz.G_stop<-true
			}
			time.Sleep(50*time.Second)
		}
	}()
	go func(){
		for true{
			<-fuzz.G_start
			timeLock.Lock()
			if contract!=oldcontract{
				start_time = uint64(time.Now().Unix())
				time_bound = start_time + MAX_TIME
				
				oldcontract = contract
				unchanged_duration = MAX_DURATION
			}
			timeLock.Unlock()
			log.Println(fuzz.RAND_CASE_SCALE)
			group_count = 3*fuzz.RAND_CASE_SCALE
			log.Println(group_count)
			log.Printf("group_count:%d\n",group_count)
			fuzz.G_sig_continue<-true
		}
	}()
	http.HandleFunc("/hack",hackHandler)
	// log.Fatal(http.ListenAndServe(":8888", nil))
	log.Fatal(http.ListenAndServe(fuzz.Global_listen_port, nil))
}
