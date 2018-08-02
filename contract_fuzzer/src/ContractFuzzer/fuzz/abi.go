package fuzz

import (
	"encoding/json"
	"io"
	"log"
	"strings"
	"fmt"
)

func decarts_product_one(outs [][]interface{}) string{
	var vals = make([]interface{},0) 
	if len(outs)==0{
		return ""
	}
	for _,v := range outs[0]{
		vals = append(vals,stringify(v))
	}
	for i:=1; i<len(outs); i++{
		vals = product(vals,outs[i])
	}
	// println(vals)
	return vals[0].(string)
}

func stringify(val interface{}) string{
	_,ok := val.(string)
	if ok{
		if strings.Contains(val.(string),"\""){
			return val.(string)
		}else{
			return "\"" + val.(string) + "\""
		}

	}else{
		data, _ := json.Marshal(val)
		return string(data)
	}
}
func product(A, B []interface{})[]interface{}{
	var rets = make([]interface{},0,0)
	for _, a := range  A{
		for _,b := range  B{
			rets = append(rets,stringify(a)+","+stringify(b))
		}
	}
	return  rets
}

type Elem struct{
	Name string `json:"name,omitempty"`
	Type string `json:"type"`
	Out  []interface{} `json:"out,omitempty"`
}
type IOput []Elem
func newIOput(jsondata []byte) (*IOput,error){
	var ioput  = new(IOput)
	if err := json.Unmarshal(jsondata,ioput);err!=nil{
		return nil,JSON_UNMARSHAL_ERROR(err)
	}
	return ioput,nil
}
func (input *IOput)String() string{
	buf,_:= json.Marshal(input)
	return string(buf)
}
func (input *IOput) fuzz() (interface{},error){
	for i,_:= range *input{
		elem := &(*input)[i]
		out,err := fuzz(elem.Type)
		if err!=nil{
			return nil,err
		}
		elem.Out = out

	}
	var outs = make([][]interface{},0,0)
	for _,elem := range *input{
		outs = append(outs,elem.Out)
	}
	val := decarts_product_one(outs)
	return val,nil
}
type Function struct{
	Name string `json:"name,omitempty"`
	Type string `json:"type"`
	Inputs IOput `json:"inputs,omitempty"`
	Outputs IOput `json:"outputs,omitempty"`
	Payable bool  `json:"payable"`
	Constant bool `json:"constant,omitempty"`
}
func (fun *Function) Sig() string{
	var elems = ([]Elem)(fun.Inputs)
	sig := fun.Name+"("
	for i,elem := range  elems{
		if i==0{
			sig += elem.Type
		}else{
			sig +=","+elem.Type
		}
	}
	sig = sig+")"
	return  sig
}
func (fun *Function) Values() []interface{}  {
	var elems = ([]Elem)(fun.Inputs)
	var outs = make([][]interface{},0,0)

	for _,elem := range elems{
		outs = append(outs,elem.Out)
	}

	var vals = make([]interface{},0,0)
	if len(outs)==0 {
		return nil
	}else{
		for _,v := range outs[0]{
			vals = append(vals,stringify(v))
		}
		for i:=1; i<len(outs); i++{
			if i > 3 && len(outs[i])>2 {
				c := randintOne(len(outs[i]),0)
				outs[i][0] =  outs[i][c]
				c = randintOne(len(outs[i]),0)
				outs[i][1] =  outs[i][c]
				outs[i] = outs[i][:2]
			}
			vals = product(vals,outs[i])
		}
	}
	return vals
}
type Abi []*Function
func newAbi(jsondata []byte) (*Abi,error){
	var abi = new(Abi)
	if err := json.Unmarshal(jsondata,abi);err!=nil{
		return nil,JSON_UNMARSHAL_ERROR(err)
	}
	return abi,nil
}
func (abi *Abi) OutputValue(writer io.Writer){
	funs := ([]*Function)(*abi)
	for _, fun := range funs{
		if fun!=nil {
			sig := fun.Sig()
			values := fun.Values()
			log.Println(sig + ":" + string(len(values)))
			if len(values) != 0 {

				if len(values) >Global_fun_scale{
					out_values := make([]interface{},0,Global_fun_scale)
					for i:=0;i<Global_fun_scale;i++{
						c := randintOne(len(values),0)
						out_values = append(out_values,values[c])
					}
					//values = values[:Global_fun_scale]
					values = out_values
				}
				for _, value := range values {
					writer.Write([]byte(sig + ":"))
					writer.Write([]byte("[" + value.(string) + "]"))
					writer.Write([]byte("\n"))
				}
			} else if len([]Elem(fun.Inputs))==0 {
				writer.Write([]byte(sig + "\n"))
			}
		}
	}

}
func (abi *Abi) String() string{
	buf,_ := json.Marshal(abi)
	return string(buf)
}
func (abi *Abi) fuzz() (ret interface{},ret_err error){
	defer func(){
		if err := recover(); err != nil {
			log.Println(err)
			// printCallStackIfError()
			ret = nil
			ret_err = ERR_ABI_FUZZ_FAILED
		}
	}()
	funs := ([]*Function)(*abi)
	funcs := make([]interface{},len(funs))
	for i:=0;i<len(funs);i++{
		funcs[i] = funs[i]
	}
	var func_chose interface{}
	var err error
	var f *Function
	func_chose = nil
	for func_chose,err=g_func_Robin.Random_select(funcs);err==nil&&func_chose.(*Function).Type!="function";	func_chose,err=g_func_Robin.Random_select(funcs){
	}
	if err!=nil{
		return "0x0",nil
	}
	f = func_chose.(*Function)
	if len(f.Inputs)>0{
		G_current_fun = f
		// log.Println(f.Sig())
		if  ret,err := f.Inputs.fuzz();err==nil{
			return fmt.Sprintf("%s:[%s]",f.Sig(),ret.(string)),nil
		}else{
			return "0x0",err
		}
  	}else{
		return fmt.Sprintf("%s",f.Sig()),nil
	}
}
