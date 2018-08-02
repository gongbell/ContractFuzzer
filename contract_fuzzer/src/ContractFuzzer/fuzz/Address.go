package fuzz

import (
	"encoding/json"
	"strings"
)
type AddressSeeds struct{
	Name  string  `json:"name""`
	Seeds []string `json:"seeds"`
	Seeds1 []string `json:"seeds1"`
	Seeds2 []string `json:"seeds2"`
}
type ADdress Type
func (self ADdress) name() string{
	return typeToString[Type(self)]
}
func (self ADdress) seeds(jsondata []byte)(*AddressSeeds,error){
	var seeds = new(AddressSeeds)
	if err:=json.Unmarshal(jsondata,seeds);err!=nil{
		return nil,JSON_UNMARSHAL_ERROR(err)
	}else{
		return seeds,nil
	}
}
func (self ADdress) fuzz(isElem bool) ([]interface{},error){
	defer func(){
		// log.Println("fuzz finished")
	}()
	//Size set to 1. We just randomly generate one value once 'fuzz' initiated.
	const  size = 1
	const (
          Prob_Normal_Account = 20
	      Prob_AgentCall_Contract = 30
	      Prob_InnerCall_Contract = 30
		  Prob_Other_Contract = 20
	)
	//Probability map:
	//    Normal     Account: 20%
	//    Agent     Contract: 30%
	//    InnerCall Contract: 30% . (InnerCall Contract refers to whom itself  have underlying function sig inside in current function call)
	//    Other     Contract: 20%   
	var  Probs = []int{Prob_Normal_Account,Prob_AgentCall_Contract,Prob_InnerCall_Contract,Prob_Other_Contract}
	var seedFile = "/config/addressSeed.json"
	var seed  *AddressSeeds=nil
	if Global_addrSeed!=""{
		seedFile = Global_addrSeed
	}
    if jsondata,err := readFile(seedFile);err==nil{
		if seed,err = self.seeds(jsondata);err!=nil{
			return nil,err
		}	
	}else{
		return nil,err
	}
	var normal_Account = ConvertStringSlice2InterfaceSlice(seed.Seeds)
	var agent_Contract = ConvertStringSlice2InterfaceSlice(seed.Seeds1)
	var other_Contract = ConvertStringSlice2InterfaceSlice(seed.Seeds2)
	var innerCall_ContractAddress []interface{}=make([]interface{},0)
	if len(G_current_bin_fun_sigs)!=0 && len(G_current_abi_sigs)!=0{
		sig := G_current_abi_sigs[strings.TrimSpace(G_current_fun.(*Function).Sig())]
		if inner_call_fun_sigs,found := G_current_bin_fun_sigs[sig];found {
			for _,item := range inner_call_fun_sigs{
				// log.Println(sig,inner_call_fun_sigs,item)
				for _,contract := range GlobalFUNSIG_CONTRACT_MAP[item]{
					innerCall_ContractAddress =append(innerCall_ContractAddress,GlobalADDR_MAP[contract]) 
				}
			}
		}
	}
	box := randintOne(100,0)
	sum := 0
	for i,_:= range Probs{
		sum += Probs[i]
		if box<sum{
			switch i{
			case 0:
				if bid,err:=g_paramval_Address_Robin.Random_select(normal_Account);err==nil{
					return []interface{}{bid},nil
				}else{
					continue
				}
			case 1:
				if bid,err:=g_paramval_Address_Robin.Random_select(agent_Contract);err==nil{
					return []interface{}{bid},nil
				}else{
					continue
				}
			case 2:
				if bid,err:=g_paramval_Address_Robin.Random_select(innerCall_ContractAddress);err==nil{
					return []interface{}{bid},nil
				}else{
					continue
				}
			case 3:
				if bid,err:=g_paramval_Address_Robin.Random_select(other_Contract);err==nil{
					return []interface{}{bid},nil
				}else{
					continue
				}
			default:
				return nil,nil
			}
		}
	}
    return nil,nil
}
