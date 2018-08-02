package fuzz

import (
	"encoding/json"
	"math/big"
	"strconv"
)
var(
	IntMax =  map[int]string{
		1:"0x7f",
		2:"0x7fff",
		3:"0x7fffff",
		4:"0x7fffffff",
		5:"0x7fffffffff",
		6:"0x7fffffffffff",
		7:"0x7fffffffffffff",
		8:"0x7fffffffffffffff",
		9:"0x7fffffffffffffffff",
		10:"0x7fffffffffffffffffff",
		11:"0x7fffffffffffffffffffff",
		12:"0x7fffffffffffffffffffffff",
		13:"0x7fffffffffffffffffffffffff",
		14:"0x7fffffffffffffffffffffffffff",
		15:"0x7fffffffffffffffffffffffffffff",
		16:"0x7fffffffffffffffffffffffffffffff",
		17:"0x7fffffffffffffffffffffffffffffffff",
		18:"0x7fffffffffffffffffffffffffffffffffff",
		19:"0x7fffffffffffffffffffffffffffffffffffff",
		20:"0x7fffffffffffffffffffffffffffffffffffffff",
		21:"0x7fffffffffffffffffffffffffffffffffffffffff",
		22:"0x7fffffffffffffffffffffffffffffffffffffffffff",
		23:"0x7fffffffffffffffffffffffffffffffffffffffffffff",
		24:"0x7fffffffffffffffffffffffffffffffffffffffffffffff",
		25:"0x7fffffffffffffffffffffffffffffffffffffffffffffffff",
		26:"0x7fffffffffffffffffffffffffffffffffffffffffffffffffff",
		27:"0x7fffffffffffffffffffffffffffffffffffffffffffffffffffff",
		28:"0x7fffffffffffffffffffffffffffffffffffffffffffffffffffffff",
		29:"0x7fffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
		30:"0x7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
		31:"0x7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
		32:"0x7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
	}
	IntMin = map[int]string{
		1:"-0x80",
		2:"-0x8000",
		3:"-0x800000",
		4:"-0x80000000",
		5:"-0x8000000000",
		6:"-0x800000000000",
		7:"-0x80000000000000",
		8:"-0x8000000000000000",
		9:"-0x800000000000000000",
		10:"-0x80000000000000000000",
		11:"-0x8000000000000000000000",
		12:"-0x800000000000000000000000",
		13:"-0x80000000000000000000000000",
		14:"-0x8000000000000000000000000000",
		15:"-0x800000000000000000000000000000",
		16:"-0x80000000000000000000000000000000",
		17:"-0x8000000000000000000000000000000000",
		18:"-0x800000000000000000000000000000000000",
		19:"-0x80000000000000000000000000000000000000",
		20:"-0x8000000000000000000000000000000000000000",
		21:"-0x800000000000000000000000000000000000000000",
		22:"-0x80000000000000000000000000000000000000000000",
		23:"-0x8000000000000000000000000000000000000000000000",
		24:"-0x800000000000000000000000000000000000000000000000",
		25:"-0x80000000000000000000000000000000000000000000000000",
		26:"-0x8000000000000000000000000000000000000000000000000000",
		27:"-0x800000000000000000000000000000000000000000000000000000",
		28:"-0x80000000000000000000000000000000000000000000000000000000",
		29:"-0x8000000000000000000000000000000000000000000000000000000000",
		30:"-0x800000000000000000000000000000000000000000000000000000000000",
		31:"-0x80000000000000000000000000000000000000000000000000000000000000",
		32:"-0x8000000000000000000000000000000000000000000000000000000000000000",
	}
)
type IntSeed struct {
	Name  string	`json:"name"`
	Seed  []int 	`json:"seed"`
}
type IntSeeds struct{
	Name  string	  `json:"name"`
	Seeds []IntSeed   `json:"seeds"`
}
func (seeds IntSeeds) String() string{
	buf,_ :=json.Marshal(seeds)
	return string(buf)
}
func (seeds IntSeeds) getSeeds(name string) ([]int,bool){
	for _,intseed := range seeds.Seeds{
		if intseed.Name == name{
			return intseed.Seed,true
		}
	}
	return nil,false
}
type Int Type
func (self Int) String() string{
	return typeToString[Type(self)]
}
func (self Int) size() uint32{
	return uint32(self)-uint32(Int8)+1
}

func (self Int) seeds(jsondata []byte) (*IntSeeds,error){
	var seeds = new(IntSeeds)
	if err := json.Unmarshal(jsondata,seeds);err!=nil{
		return seeds,JSON_UNMARSHAL_ERROR(err)
	}else {
		return  seeds,nil
	}
}
func (self Int) random_select() ([]interface{},error){
	var Max,Min big.Int
	box := randintOne(int(self.size()),1)
	// Max.SetString(UintMax[int(self.size())],0)
	// Min.SetString(UintMin[int(self.size())],0)
	Max.SetString(UintMax[box],0)
	Min.SetString(UintMin[box],0)
	bid := randBint(Max,Min)
	return []interface{}{BigInt(*bid).String()},nil
}

func (self Int) fuzz(isElem bool) ([]interface{},error){
	var (
		seedFile =  "/config/intSeed.json"
		seed *IntSeeds = nil
		Seeds = make([]interface{},0)
	)
	const (
		Prob_Seeds = 50
		Prob_random = 50
	)
	var Probs = []int{Prob_Seeds,Prob_random}
	if Global_intSeed!=""{
		seedFile = Global_intSeed
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
			Seeds = ConvertIntSlice2InterfaceSlice(seeds)
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
				if out,err := g_paramval_Int_Robin.Random_select(Seeds);err==nil{
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

func Int2BigIntString(arr []int) (ret []string){
	ret = make([]string,0)
	for _,item:= range arr{
		ret = append(ret,strconv.Itoa(item))
	}
	return
}