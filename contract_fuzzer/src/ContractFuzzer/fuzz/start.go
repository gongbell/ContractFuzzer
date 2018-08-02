package fuzz

import (
	"os"
	"bufio"
	"strings"
	"log"
	"io"
	"fmt"
	"net/http"
	"net/url"
	abi_gen "ContractFuzzer/abi"
)
var (
	error_log = "/list/error-line.log"
)
var transport = http.Transport{
	DisableKeepAlives: false,
	}
var Client = http.Client{Transport:&transport}
var (
	Global_contractList []string=[]string{""}
	Global_addrSeed  string = ""
	Global_intSeed string = ""
	Global_uintSeed string = ""
	Global_stringSeed string = ""
	Global_byteSeed string = ""
	Global_bytesSeed string = ""
	Global_scale  int = 2
	Global_fun_scale int = 8
	Global_fstart int = 0
	Global_fend int  = 0
	Global_addr_map = ""
	Global_abi_sigs_dir = ""
	Global_bin_sigs_dir = ""
	Global_listen_port = ""
	Global_tester_port = ""
	GlobalADDR_MAP = make(map[string]string)
	GlobalFUNSIG_CONTRACT_MAP = make(map[string][]string)
)

func  createADDR_MAP() error {
	f, err := os.Open(Global_addr_map)
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
			f, _ := os.OpenFile(error_log,os.O_CREATE|os.O_APPEND,0666)
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
		GlobalADDR_MAP[name] = addr
	}
	return nil
}
func  getFUNSIG_CONTRACT_by_file(file string) error{
	f, err := os.Open(Global_abi_sigs_dir+"/"+file)
	defer func() {
		f.Close()
	}()
	if err != nil {
		return err
	}
	strs := strings.Split(file,"/")
	name := strings.Split(strs[len(strs)-1],".")[0]
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
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
		str2 := strings.Split(line,":")
		fun_sig := str2[0]
		_, found := GlobalFUNSIG_CONTRACT_MAP[fun_sig]
		if !found{
			GlobalFUNSIG_CONTRACT_MAP[fun_sig] = []string{name}
		}else{
			GlobalFUNSIG_CONTRACT_MAP[fun_sig] = append(GlobalFUNSIG_CONTRACT_MAP[fun_sig],name)
		}
	}
	return  nil
}
func  createFUNSIG_CONTRACT_MAP() error{
	files, err := readDir(Global_abi_sigs_dir)
	if err!=nil{
		return err
	}
	for _,file := range  files{
		if err := getFUNSIG_CONTRACT_by_file(file); err!=nil{
			return fmt.Errorf("createFUNSIG_CONTRACT_MAP.getFUNSIG_CONTRACT_by_file %s: %s",file,err)
		}
	}
	return  nil
}
func setG_current_bin_fun_sigs() error{
	f, err := os.Open(Global_bin_sigs_dir+"/"+G_current_contract.(string)+".bin.sig")
	defer func() {
		f.Close()
	}()
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
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
		str2 := strings.Split(line,":")
		fun_sig := str2[0]
		innercall_funsigs := strings.Split(str2[1]," ")
		for _,sig := range innercall_funsigs{
			_, found := G_current_bin_fun_sigs[fun_sig]
			if !found{
				G_current_bin_fun_sigs[fun_sig] = []string{sig}
			}else{
				G_current_bin_fun_sigs[fun_sig] = append(G_current_bin_fun_sigs[fun_sig],sig)
			}
		}
	}
	return  nil
}
func setG_current_abi_sigs() error{
	f, err := os.Open(Global_abi_sigs_dir+"/"+G_current_contract.(string)+".abi")
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
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
		str2 := strings.Split(line,":")
		fun_sig := str2[0]
		fun_sig_str := strings.TrimSpace(str2[1])
		G_current_abi_sigs[fun_sig_str] = fun_sig
	}
	return  nil
}
func  Init(contractListPath ,addrSeed,intSeed,uintSeed,stringSeed,byteSeed,
	 bytesSeed string,scale ,fun_scale, fstart,fend int,addr_map, abi_sigs_dir,bin_sigs_dir string,
	 listen_port,tester_port string)(error){
	Global_contractList = make([]string,0,0)
	if contractListPath != "null"{
		file,err:=os.Open(contractListPath)
		if err!=nil{
			return err
		}
		reader := bufio.NewReader(file)
		for contract,e:=reader.ReadString('\n');e==nil;contract,e=reader.ReadString('\n'){
			Global_contractList = append(Global_contractList,strings.Trim(contract,"\n")+".abi")
		}
	}
	Global_addrSeed = addrSeed
	Global_bytesSeed = bytesSeed
	Global_intSeed = intSeed
	Global_stringSeed = stringSeed
	Global_uintSeed = uintSeed
	Global_scale = scale
	Global_fstart = fstart
	Global_fend  = fend
	Global_fun_scale = fun_scale
	Global_addr_map = addr_map
	Global_abi_sigs_dir = abi_sigs_dir
	Global_bin_sigs_dir = bin_sigs_dir
	Global_listen_port = listen_port
	Global_tester_port = tester_port
	createADDR_MAP()
	createFUNSIG_CONTRACT_MAP()
	return nil
}
func sendMsg2RunnerMonitor(address string, msgs []string)(bool){
	values := url.Values{"address":[]string{address},"msg":msgs}
	go func(){
			// res,_ := Client.Get("http://localhost:6666/runnerMonitor?"+values.Encode())
			res,_ := Client.Get(Global_tester_port+"/runnerMonitor?"+values.Encode())
			// log.Println(res)
			defer func(){
				if err:= recover();err!=nil{
					log.Println(err)
				}else{
					res.Body.Close()
				}
			}()
	}()
	return true
}
var(
	rand_case_ranges = []interface{}{20,25,30,35,40}
    rand_case_scales = []interface{}{6,7,8,9,10}
)
var(
	RAND_CASE_RANGE = 10
	RAND_CASE_SCALE = 10
)
// no buffered channel.
// synchronized.
var(
	G_stop = make(chan bool,0)
	G_start = make(chan bool,0)
	G_finish = make(chan bool,0)
	G_sig_continue = make(chan bool,0)
) 

func  Start(dir string,outdir string)(error)  {
	defer func(){
		if err := recover(); err != nil {
			log.Println(err)
			printCallStackIfError()
		}
	}()
	if err:=os.Mkdir(outdir,0777);err!=nil{
		if !os.IsExist(err){
			return err
		}
	}
	files,err:=readDir(dir)
	if err!=nil{
		return err
	}
	if len(Global_contractList)==0 {
		for i:= Global_fstart;i<Global_fend && i<len(files)-1;i++{
			
			file := files[i]
			path := dir + "/" + file

			G_current_contract = strings.TrimSpace(strings.Split(file,".")[0])
		    current_contract_address :=  GlobalADDR_MAP[G_current_contract.(string)]
			setG_current_bin_fun_sigs()
			setG_current_abi_sigs()

			data, err := readFile(path)
			if err != nil {
				continue
			}
			abi, err := newAbi(data)
			if err != nil {
				continue
			}
			if bid,err:=  g_robin.Random_select(rand_case_ranges);err==nil{
				RAND_CASE_RANGE = bid.(int)
			}else{
				continue
			}
			for no:=0;no<RAND_CASE_RANGE;no++{
				funs := make([]string,0)
				msgs := make([]string,0)
				// RAND_CASE_SCALE,_ = g_robin.Random_select(rand_case_scales)
				if bid,err:= g_robin.Random_select(rand_case_scales);err==nil{
					RAND_CASE_SCALE = bid.(int)
				}else{
					continue
				}
				invalid_no := 0
				for i:=0;i<RAND_CASE_SCALE;i++{
					if ret, err := abi.fuzz(); err == nil {
						log.Println(ret)
						if strings.Contains(ret.(string),"0x0")==true{
							msgs = append(msgs,"0xcaffee")
						}else{
							if hex_str,err := abi_gen.Parse_GenMsg(ret.(string));err==nil{
								log.Println(hex_str)
								msgs = append(msgs,hex_str)
								if len(hex_str)>10{
									funs = append(funs,hex_str[:10])
								}else{
									funs = append(funs,hex_str)
								}
							}
					    }
						// return nil
					}else{
						invalid_no++
						continue
					}
				}
				RAND_CASE_SCALE -= invalid_no
				if RAND_CASE_RANGE==0{
					continue
				}
				G_start<-true
				<-G_sig_continue
				sendMsg2RunnerMonitor(current_contract_address,msgs)
				msgs = make([]string,0)
				funs = make([]string,0)
				if no==RAND_CASE_RANGE-1{
					break
				}
				if c:= <-G_stop;c==true{
					break
				} 
			}
			
			// fuzz_file := outdir + "/" + file
			// fuzz, _ := os.Create(fuzz_file)
			// abi.OutputValue(fuzz)
			// fuzz.Close()
		}
		G_finish<-true
	}else{
		for _, file := range Global_contractList {
			path := dir + "/" + file
			data, err := readFile(path)
			if err != nil {
				continue
			}

			G_current_contract = strings.TrimSpace(strings.Split(file,".")[0])
		    current_contract_address :=  GlobalADDR_MAP[G_current_contract.(string)]
			setG_current_bin_fun_sigs()
			setG_current_abi_sigs()

			abi, err := newAbi(data)
			if err != nil {
				continue
			}
			if bid,err:=  g_robin.Random_select(rand_case_ranges);err==nil{
				RAND_CASE_RANGE = bid.(int)
			}else{
				continue
			}
			for no:=0;no<RAND_CASE_RANGE;no++{
				funs := make([]string,0)
				msgs := make([]string,0)
				// RAND_CASE_SCALE,_ = g_robin.Random_select(rand_case_scales)
				if bid,err:= g_robin.Random_select(rand_case_scales);err==nil{
					RAND_CASE_SCALE = bid.(int)
				}else{
					continue
				}
				invalid_no := 0
				for i:=0;i<RAND_CASE_SCALE;i++{
					if ret, err := abi.fuzz(); err == nil {
						log.Println(ret)
						if strings.Contains(ret.(string),"0x0")==true{
							msgs = append(msgs,"0xC0FFEE")
						}else{
							if hex_str,err := abi_gen.Parse_GenMsg(ret.(string));err==nil{
								log.Println(hex_str)
								msgs = append(msgs,hex_str)
								if len(hex_str)>10{
									funs = append(funs,hex_str[:10])
								}else{
									funs = append(funs,hex_str)
								}
							}
					    }
						// return nil
					}else{
						invalid_no++
						continue
					}
				}
				RAND_CASE_SCALE -= invalid_no
				if RAND_CASE_RANGE==0{
					continue
				}
				G_start<-true
				<-G_sig_continue
				sendMsg2RunnerMonitor(current_contract_address,msgs)
				msgs = make([]string,0)
				funs = make([]string,0)
				if no==RAND_CASE_RANGE-1{
					break
				}
				if c:= <-G_stop;c==true{
					break
				} 
			}
		}
		G_finish<-true
	}
	return nil
}
