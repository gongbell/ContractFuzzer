package abi

import (
	"os"
	// "log"
	"github.com/ethereum/go-ethereum/common"
	"bytes"
	"encoding/json"
	"strings"
	"errors"
)
var errLogger,_= os.Create("./error.log")

func Parse_GenMsg(fun_sig string)(ret string,ret_err error){
		defer func(){
			if err := recover(); err != nil {
				ret = ""
				ret_err = errors.New("Parse_GenMsg Failed")
			}
		}()

	//str := `_startNextCompetition(string,uint32,uint88,uint8,uint8,uint16,uint64,uint32,bytes32,uint32[]):["world","0x5d39ad1f","0x4300fac8fcc88c0ff33739","0xa8","0x17","0x1fec","0xdd87e95ce32da82e","0xe35d1e45","0xfbd57879ec4e6a57fd57b683ae0e5d16cb134bebb8bb478583bbdca9795d5a",["0xa58fbc22","0x6161dc4f","0xc673347e","0x9a9d897c"]]`
	//str :=`approveAndCall(address,uint256,bytes):["0x437998dacb1b3684567e84971da8c6257132aa99","0x737d631b4164c9a3c2bff55da1785064c3f319b574a11d8f712e2ed6a04cdf2d","0x123543523f95723e5432"]`l
	// log.Printf("%s",fun_sig)
	name,data := parse(fun_sig)
	// log.Printf("%s",name)
	// log.Printf("data:%s",common.Bytes2Hex(data))
	abi,err := JSON(bytes.NewReader(data))
	if err!=nil{
		errLogger.Write([]byte(fun_sig))
		errLogger.Write([]byte("\n"))
		errLogger.Write([]byte(err.Error()))
		errLogger.Write([]byte("\n"))
		return "",err
	}
	var args interface{}
	if strings.Contains(fun_sig,":")==true{
		// log.Printf(":exist")
		strs := strings.Split(fun_sig,":")
		json.Unmarshal([]byte(strs[1]), &args)
		packed, errr := abi.Pack(name, args.([]interface{})...)
		if errr!=nil{
			errLogger.Write([]byte(fun_sig))
			errLogger.Write([]byte("\n"))
			errLogger.Write([]byte(errr.Error()))
			errLogger.Write([]byte("\n"))
			return "",errr
		}
		// log.Printf("0x%s",common.Bytes2Hex(packed))
		return "0x"+common.Bytes2Hex(packed),nil
	}else{
		// log.Printf(":not exist")
		packed, errr := abi.Pack(name)
		if errr!=nil{
			errLogger.Write([]byte(fun_sig))
			errLogger.Write([]byte("\n"))
			errLogger.Write([]byte(errr.Error()))
			errLogger.Write([]byte("\n"))
			return "",errr
		}
		// log.Printf("0x%s",common.Bytes2Hex(packed))
		return "0x"+common.Bytes2Hex(packed),nil
	}

}
func parse(fun_sig string)(string, []byte){
	name := strings.Split(fun_sig,"(")[0]
	//log.Printf("%s",name)
	paramlistStart := strings.LastIndex(fun_sig,"(")
	paramlistEnd := strings.LastIndex(fun_sig,")")
	paramlist := strings.Split(string(fun_sig[paramlistStart+1:paramlistEnd]),",")
	//log.Printf("%s",paramlist)
	abi := make([]map[string]interface{},0,0)
	method := make(map[string]interface{})
	method["name"] = name
	method["type"] = "function"
	method["inputs"] = make([]map[string]string, 0, 0)
	if len(paramlist[0])>0 {
		for _, param := range paramlist {
			input := make(map[string]string)
			input["type"] = param
			method["inputs"] = append(method["inputs"].([]map[string]string), input)
		}
	}
	abi = append(abi, method)

	data,_ :=json.Marshal(abi)
	// log.Printf("%s",string(data))
	return name,data
}