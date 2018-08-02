package fuzz

import (
	"math/big"
	"encoding/json"
)

type ByteSeed struct {
	Name  string	`json:"name"`
	Seed  []string 	`json:"seed"`
}
type ByteSeeds struct{
	Name  string	  `json:"name"`
	Seeds []ByteSeed   `json:"seeds"`
}
func (seeds ByteSeeds) String() string{
	buf,_ :=json.Marshal(seeds)
	return string(buf)
}
func (seeds ByteSeeds) getSeeds(name string) ([]string,bool){
	for _,seed := range seeds.Seeds{
		if seed.Name == name{
			return seed.Seed,true
		}
	}
	return nil,false
}
type Byte Type
func (self Byte) String() string{
	return typeToString[Type(self)]
}
func (self Byte) size() uint32{
	return uint32(self)-uint32(Bytes1)+1
}
func (self Byte) seeds(jsondata []byte) (*ByteSeeds,error){
	var seeds = new(ByteSeeds)
	if err := json.Unmarshal(jsondata,seeds);err!=nil{
		return seeds,JSON_UNMARSHAL_ERROR(err)
	}else {
		return  seeds,nil
	}
}
var(
	ByteMax =  map[int]string{
		1:"0xff",
		2:"0xffff",
		3:"0xffffff",
		4:"0xffffffff",
		5:"0xffffffffff",
		6:"0xffffffffffff",
		7:"0xffffffffffffff",
		8:"0xffffffffffffffff",
		9:"0xffffffffffffffffff",
		10:"0xffffffffffffffffffff",
		11:"0xffffffffffffffffffffff",
		12:"0xffffffffffffffffffffffff",
		13:"0xffffffffffffffffffffffffff",
		14:"0xffffffffffffffffffffffffffff",
		15:"0xffffffffffffffffffffffffffffff",
		16:"0xffffffffffffffffffffffffffffffff",
		17:"0xffffffffffffffffffffffffffffffffff",
		18:"0xffffffffffffffffffffffffffffffffffff",
		19:"0xffffffffffffffffffffffffffffffffffffff",
		20:"0xffffffffffffffffffffffffffffffffffffffff",
		21:"0xffffffffffffffffffffffffffffffffffffffffff",
		22:"0xffffffffffffffffffffffffffffffffffffffffffff",
		23:"0xffffffffffffffffffffffffffffffffffffffffffffff",
		24:"0xffffffffffffffffffffffffffffffffffffffffffffffff",
		25:"0xffffffffffffffffffffffffffffffffffffffffffffffffff",
		26:"0xffffffffffffffffffffffffffffffffffffffffffffffffffff",
		27:"0xffffffffffffffffffffffffffffffffffffffffffffffffffffff",
		28:"0xffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
		29:"0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
		30:"0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
		31:"0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
		32:"0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
	}
	ByteMin = map[int]string{
		1:"0x0",
		2:"0x0",
		3:"0x0",
		4:"0x0",
		5:"0x0",
		6:"0x0",
		7:"0x0",
		8:"0x0",
		9:"0x0",
		10:"0x0",
		11:"0x0",
		12:"0x0",
		13:"0x0",
		14:"0x0",
		15:"0x0",
		16:"0x0",
		17:"0x0",
		18:"0x0",
		19:"0x0",
		20:"0x0",
		21:"0x0",
		22:"0x0",
		23:"0x0",
		24:"0x0",
		25:"0x0",
		26:"0x0",
		27:"0x0",
		28:"0x0",
		29:"0x0",
		30:"0x0",
		31:"0x0",
		32:"0x0",
    }
)
func (self Byte) random_select()([]interface{},error){
	var Max,Min big.Int
	Max.SetString(ByteMax[int(self.size())],0)
	Min.SetString(ByteMin[int(self.size())],0)
	bid := randBint(Max,Min)
	return []interface{}{BigInt(*bid).String()},nil
}
func (self Byte) fuzz(isElem bool) ([]interface{},error){
	var (
		seedFile = "/config/byteSeed.json"
		seed *ByteSeeds = nil
		Seeds []interface{} = make([]interface{},0)
	)
	const(
		Prob_Seeds = 50
		Prob_Random = 50
	)
	var Probs = []int{Prob_Seeds,Prob_Random}
	if Global_byteSeed!=""{
		seedFile = Global_byteSeed
	}
	if jsondata,err := readFile(seedFile);err==nil{
		if seed,err = self.seeds(jsondata);err!=nil{
			return nil,err
		}	
	}else{
		return nil,err
	}
    if name,found:= typeToString[Type(self)];found{
		if seeds,found:= seed.getSeeds(name);found{
			Seeds = ConvertStringSlice2InterfaceSlice(seeds)
		}
	}else{
		return nil,ERR_TYPE_NOT_FOUND
	}
	box := randintOne(100,0)
	sum := 0
	for index,_ := range Probs{
		sum += Probs[index]
		if box<=sum{
			switch index{
			case 0:
				if out,err := g_paramval_Byte_Robin.Random_select(Seeds);err==nil{
					return []interface{}{out},nil
				}else{
					continue
				}
			case 1:
				if out,err := self.random_select();err==nil{
					return out,nil
				}else{
					continue
				}
			}
		}
	}
	return nil,ERR_FUZZ_TYPE_FAILED
}
