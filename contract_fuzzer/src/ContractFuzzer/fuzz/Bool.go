package fuzz

import (
	"encoding/json"
)

type BOol Type
func (self BOol) name() string{
	return typeToString[Type(self)]
}
func (self BOol) fuzz(isElem bool)([]interface{},error){
	v := []interface{}{
		true,
		false,
	}
	bid,err := g_paramval_Bool_Robin.Random_select(v)
	return []interface{}{bid},err
}
type MyBOOL bool
func (self MyBOOL) String() string{
	buf,_ :=json.Marshal(self)
	return string(buf)
}