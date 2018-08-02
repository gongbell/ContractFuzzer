package fuzz

import (
	"regexp"
	"strings"
	"fmt"
	"strconv"
	"os"
	"io"
	"encoding/json"
)
var fixReg = regexp.MustCompile("^(.*)\\[([\\d]+)\\]+$")
var DynReg = regexp.MustCompile("^(.*)\\[\\].*$")
type Fuzz interface {
	fuzz(typestr string) string
}

type  FixedArray struct {
	elem Type	         `json:"element_type"`
	size uint32          `json:"size"`
	str  string  		 `json:"description"`
	out  []interface{}     `json:"fuzz_out"`
	ostream io.Writer    `json:"-"`
}
func newFixedArray(str string) *FixedArray{
	f :=new(FixedArray)
	f.str = str
	match := fixReg.FindStringSubmatch(f.str)
	if len(match)!=0{
		elemstr := match[1]
		size,_ := strconv.Atoi(match[2])
		elem,err := strToType(elemstr)
		if err!=nil{
			fmt.Errorf("%s",err)
		}
		f.elem = elem
		f.size =uint32(size)
		f.out = make([]interface{},0)
	 }
	return f
}
func (f *FixedArray) String() string{
	buf ,_ := json.Marshal(f)
	return string(buf)
}
//fuzz:
//	  generate one array item once.
func (f *FixedArray) fuzz() ([]interface{},error){
	var (
		isElem = true
		out = make([]interface{},0)
		size = f.size
	)
	for  i := uint32(0); i<size; i++{
		if m_out,err := f.elem.fuzz(isElem);err==nil{
			out = append(out,m_out[0])
		}else{
			return nil,err
		}
	}
	return []interface{}{out},nil
}
func (f*FixedArray) SetOstream(file string){
	if ostream ,err := os.OpenFile(file,os.O_CREATE|os.O_APPEND|os.O_RDWR,0666);err!=nil{
		fmt.Printf("%s",FILE_OPEN_ERROR(err))
	}else{
		f.ostream = io.Writer(ostream)
	}
}
func (f *FixedArray) Write(data []byte){
	f.ostream.Write(data)
}
type DynamicArray struct{
	elem Type  `json:"element_type"`
	str string `json:"description"`
	out []interface{} `json:"fuzz_out"`
}
func newDynamicArray(str string) *DynamicArray{
	d := new(DynamicArray)
	d.str = str
	match := DynReg.FindStringSubmatch(d.str)
	if len(match)!=0{
		elemstr := match[1]
		elem,err := strToType(elemstr)
		if err!=nil{
			fmt.Errorf("%s",err)
		}
		d.elem = elem
		d.out  = make([]interface{},0)
	}
	return d
}
func (d *DynamicArray) fuzz() ([]interface{},error){
	const ARRAY_SIZE_LIMIT = 10
	size := randintOne(1,ARRAY_SIZE_LIMIT)
	str_fixArray := fmt.Sprintf("%s[%d]",typeToString[d.elem],size)
	fixArray := newFixedArray(str_fixArray)
	out,err := fixArray.fuzz()
    return out,err
}
func (d *DynamicArray) String()string{
	buf,_ := json.Marshal(d)
	return string(buf)
}
const (
	Cfundemental uint32 = iota
	CfixedArray
	CdynamicArray
)
func getInfo(typestr string) (uint32,error){
 typestr = strings.TrimSpace(typestr)

 if match := fixReg.MatchString(typestr);match==true{
 	return CfixedArray,nil
 }else if match := DynReg.MatchString(typestr);match==true{
		return CdynamicArray,nil
 }else if v,err := strToType(typestr);err==nil{
	return Cfundemental,nil
 }else{
 	return uint32(v),err
 }
}
func fuzz(str string)([]interface{},error){
	v,err := getInfo(str)

	if err !=nil{
		return nil,err
	}else{
		switch v {
		case Cfundemental:
			{
				f,_ := strToType(str)
				isElem := false
				out,_:= f.fuzz(isElem)
				return out,nil
			}
		case CfixedArray:
			{
				f := newFixedArray(str)
				out,_ := f.fuzz()
				return out,nil
			}
		case CdynamicArray:
			{
				d := newDynamicArray(str)
				out,_ := d.fuzz()
				return out,nil
			}
		default:
			return nil,ERR_UNKNOWN_COMPLEX_TYPE
		}
	}

}