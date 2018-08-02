package fuzz

import (
	"encoding/json"
)

type StringSeeds struct{
	Name string `json:"name"`
	Seeds []string `json:"seeds"`
}
func (self StringSeeds) String() string{
	buf,_ :=json.Marshal(self)
	return string(buf)
}
type STring Type
func (st STring) seeds(jsondata []byte)(*StringSeeds,error){
	var seeds = new(StringSeeds)
	if err:=json.Unmarshal(jsondata,seeds);err!=nil{
		return nil,JSON_UNMARSHAL_ERROR(err)
	}else{
		return seeds,nil
	}
}
func(self STring) fuzz(isElem bool)([]interface{},error){
	var (
		seedFile = "/config/stringSeed.json"
		seed *StringSeeds = nil
		Seeds []interface{} = make([]interface{},0)
	)
	if Global_stringSeed!=""{
		seedFile = Global_stringSeed
	}
	if jsondata,err := readFile(seedFile);err==nil{
		if seed,err = self.seeds(jsondata);err!=nil{
			return nil,err
		}	
	}else{
		return nil,err
	}
	Seeds = ConvertStringSlice2InterfaceSlice(seed.Seeds)
	bid,err := g_paramval_String_Robin.Random_select(Seeds)
	return []interface{}{bid},err
}
