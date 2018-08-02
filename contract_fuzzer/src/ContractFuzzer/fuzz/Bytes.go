package fuzz

import (
	"encoding/json"
	"math/big"
)

type BytesSeeds struct{
	Name string `json:"name"`
	Seeds []string `json:seeds`
}

type BYtes Type
func (self BYtes) name() string{
	return typeToString[Type(self)]
}
func (self BYtes) seeds(jsondata []byte)(*BytesSeeds,error){
	var seeds = new(BytesSeeds)
	if err:=json.Unmarshal(jsondata,seeds);err!=nil{
		return nil,JSON_UNMARSHAL_ERROR(err)
	}else{
		return seeds,nil
	}
}
func (self BYtes) random_select()([]interface{},error){
	var Max,Min big.Int
	max := "0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
	min := "0x0"
	Max.SetString(max,0)
	Min.SetString(min,0)
	bid := randBint(Max,Min)
	return []interface{}{BigInt(*bid).String()},nil
}
func (self BYtes) fuzz(isElem bool) ([]interface{},error){
	var (
		seedFile = "/config/bytesSeed.json"
		seed *BytesSeeds = nil
		Seeds = make([]interface{},0)
	)
	const(
		Prob_Seeds = 50
		Prob_Random = 50
	)
	var Probs = []int{Prob_Seeds,Prob_Random}

	if Global_bytesSeed!=""{
		seedFile = Global_bytesSeed
	}
	if jsondata,err := readFile(seedFile);err==nil{
		if seed,err = self.seeds(jsondata);err==nil{
			Seeds = ConvertStringSlice2InterfaceSlice(seed.Seeds)
		}else{
			return nil,err
		}	
	}else{
		return nil,err
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
	return  nil,ERR_FUZZ_TYPE_FAILED
}
