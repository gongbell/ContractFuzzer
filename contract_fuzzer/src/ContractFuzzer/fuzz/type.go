package fuzz

import (
	"fmt"
	"math/big"
	"strconv"
)
type Type uint32
const(
		Undefined Type = iota
		Int8
		Int16
		Int24
		Int32
		Int40
		Int48
		Int56
		Int64
		Int72
		Int80
		Int88
		Int96
		Int104
		Int112
		Int120
		Int128
		Int136
		Int144
		Int152
		Int160
		Int168
		Int176
		Int184
		Int192
		Int200
		Int208
		Int216
		Int224
		Int232
		Int240
		Int248
		Int256
		Uint8
		Uint16
		Uint24
		Uint32
		Uint40
		Uint48
		Uint56
		Uint64
		Uint72
		Uint80
		Uint88
		Uint96
		Uint104
		Uint112
		Uint120
		Uint128
		Uint136
		Uint144
		Uint152
		Uint160
		Uint168
		Uint176
		Uint184
		Uint192
		Uint200
		Uint208
		Uint216
		Uint224
		Uint232
		Uint240
		Uint248
		Uint256
		Bytes1
		Bytes2
		Bytes3
		Bytes4
		Bytes5
		Bytes6
		Bytes7
		Bytes8
		Bytes9
		Bytes10
		Bytes11
		Bytes12
		Bytes13
		Bytes14
		Bytes15
		Bytes16
		Bytes17
		Bytes18
		Bytes19
		Bytes20
		Bytes21
		Bytes22
		Bytes23
		Bytes24
		Bytes25
		Bytes26
		Bytes27
		Bytes28
		Bytes29
		Bytes30
		Bytes31
		Bytes32
		Address
		String
		Bytes
		Bool
)
var typeToString = map[Type]string{
	Undefined:"Undefined",
	Int8 : "int8",
	Int16: "int16",
	Int24: "int24",
	Int32: "int32",
	Int40: "int40",
	Int48: "int48",
	Int56: "int56",
	Int64: "int64",
	Int72: "int72",
	Int80: "int80",
	Int88: "int88",
	Int96: "int96",
	Int104: "int104",
	Int112: "int112",
	Int120: "int120",
	Int128: "int128",
	Int136: "int136",
	Int144: "int144",
	Int152: "int152",
	Int160: "int160",
	Int168: "int168",
	Int176: "int176",
	Int184: "int184",
	Int192: "int192",
	Int200: "int200",
	Int208: "int208",
	Int216: "int216",
	Int224: "int224",
	Int232: "int232",
	Int240: "int240",
	Int248: "int248",
	Int256: "int256",
	Uint8: "uint8",
	Uint16: "uint16",
	Uint24: "uint24",
	Uint32: "uint32",
	Uint40: "uint40",
	Uint48: "uint48",
	Uint56: "uint56",
	Uint64: "uint64",
	Uint72: "uint72",
	Uint80: "uint80",
	Uint88: "uint88",
	Uint96: "uint96",
	Uint104: "uint104",
	Uint112: "uint112",
	Uint120: "uint120",
	Uint128: "uint128",
	Uint136: "uint136",
	Uint144: "uint144",
	Uint152: "uint152",
	Uint160: "uint160",
	Uint168: "uint168",
	Uint176: "uint176",
	Uint184: "uint184",
	Uint192: "uint192",
	Uint200: "uint200",
	Uint208: "uint208",
	Uint216: "uint216",
	Uint224: "uint224",
	Uint232: "uint232",
	Uint240: "uint240",
	Uint248: "uint248",
	Uint256: "uint256",
	Bytes1: "bytes1",
	Bytes2: "bytes2",
	Bytes3: "bytes3",
	Bytes4: "bytes4",
	Bytes5: "bytes5",
	Bytes6: "bytes6",
	Bytes7: "bytes7",
	Bytes8: "bytes8",
	Bytes9: "bytes9",
	Bytes10: "bytes10",
	Bytes11: "bytes11",
	Bytes12: "bytes12",
	Bytes13: "bytes13",
	Bytes14: "bytes14",
	Bytes15: "bytes15",
	Bytes16: "bytes16",
	Bytes17: "bytes17",
	Bytes18: "bytes18",
	Bytes19: "bytes19",
	Bytes20: "bytes20",
	Bytes21: "bytes21",
	Bytes22: "bytes22",
	Bytes23: "bytes23",
	Bytes24: "bytes24",
	Bytes25: "bytes25",
	Bytes26: "bytes26",
	Bytes27: "bytes27",
	Bytes28: "bytes28",
	Bytes29: "bytes29",
	Bytes30: "bytes30",
	Bytes31: "bytes31",
	Bytes32: "bytes32",
	Address: "address",
	String: "string",
	Bytes: "bytes",
	Bool: "bool",
}
var stringToType = map[string]Type{
	"Undefined":Undefined,
	"int8": Int8 ,
	"int16": Int16,
	"int24": Int24,
	"int32": Int32,
	"int40": Int40,
	"int48": Int48,
	"int56": Int56,
	"int64": Int64,
	"int72": Int72,
	"int80": Int80,
	"int88": Int88,
	"int96": Int96,
	"int104": Int104,
	"int112": Int112,
	"int120": Int120,
	"int128": Int128,
	"int136": Int136,
	"int144": Int144,
	"int152": Int152,
	"int160": Int160,
	"int168": Int168,
	"int176": Int176,
	"int184": Int184,
	"int192": Int192,
	"int200": Int200,
	"int208": Int208,
	"int216": Int216,
	"int224": Int224,
	"int232": Int232,
	"int240": Int240,
	"int248": Int248,
	"int256": Int256,
	"uint8": Uint8,
	"uint16": Uint16,
	"uint24": Uint24,
	"uint32": Uint32,
	"uint40": Uint40,
	"uint48": Uint48,
	"uint56": Uint56,
	"uint64": Uint64,
	"uint72": Uint72,
	"uint80": Uint80,
	"uint88": Uint88,
	"uint96": Uint96,
	"uint104": Uint104,
	"uint112": Uint112,
	"uint120": Uint120,
	"uint128": Uint128,
	"uint136": Uint136,
	"uint144": Uint144,
	"uint152": Uint152,
	"uint160": Uint160,
	"uint168": Uint168,
	"uint176": Uint176,
	"uint184": Uint184,
	"uint192": Uint192,
	"uint200": Uint200,
	"uint208": Uint208,
	"uint216": Uint216,
	"uint224": Uint224,
	"uint232": Uint232,
	"uint240": Uint240,
	"uint248": Uint248,
	"uint256": Uint256,
	"bytes1": Bytes1,
	"bytes2": Bytes2,
	"bytes3": Bytes3,
	"bytes4": Bytes4,
	"bytes5": Bytes5,
	"bytes6": Bytes6,
	"bytes7": Bytes7,
	"bytes8": Bytes8,
	"bytes9": Bytes9,
	"bytes10": Bytes10,
	"bytes11": Bytes11,
	"bytes12": Bytes12,
	"bytes13": Bytes13,
	"bytes14": Bytes14,
	"bytes15": Bytes15,
	"bytes16": Bytes16,
	"bytes17": Bytes17,
	"bytes18": Bytes18,
	"bytes19": Bytes19,
	"bytes20": Bytes20,
	"bytes21": Bytes21,
	"bytes22": Bytes22,
	"bytes23": Bytes23,
	"bytes24": Bytes24,
	"bytes25": Bytes25,
	"bytes26": Bytes26,
	"bytes27": Bytes27,
	"bytes28": Bytes28,
	"bytes29": Bytes29,
	"bytes30": Bytes30,
	"bytes31": Bytes31,
	"bytes32": Bytes32,
	"address": Address,
	"string": String,
	"bytes": Bytes,
	"bool": Bool,
}
func (t Type) String() string{
	if v,ok := typeToString[t];ok{
		return v
	}
	return  ""
}
func strToType(s string) (Type,error){
	if v, ok := stringToType[s];ok==true{
		return  v,nil
	}
	return  0,fmt.Errorf("Missing type %s",s)
}
func (t Type) isBool() bool{
	return uint32(t)==uint32(Bool)
}
func (t Type) fuzz(isElem bool) ([]interface{},error){
	switch {
	case uint32(t)<=uint32(Int256) && uint32(t)>=uint32(Int8):
		{
			var myInt = Int(t)
			out,err := myInt.fuzz(isElem)
			if err!=nil{
				return nil,err
			}
			v := out
			strs := make([]interface{},0)
			for _,elem := range v{
				switch elem.(type){
				case big.Int:
					strs = append(strs,BigInt(elem.(big.Int)).String())	
				case string:
					strs = append(strs,elem)
				case int:
					strs = append(strs,strconv.Itoa(elem.(int)))
				}				
			}
			return strs,nil
		}
	case uint32(t)<=uint32(Uint256) && uint32(t)>=uint32(Uint8):
		{
			var myInt = Uint(t)
			out,err := myInt.fuzz(isElem)
			if err!=nil{
				return nil,err
			}
			v := out
			strs := make([]interface{},0)
			for _,elem := range v{
				switch elem.(type){
				case big.Int:
					strs = append(strs,BigInt(elem.(big.Int)).String())	
				case string:
					strs = append(strs,elem)
				case int:
					strs = append(strs,strconv.Itoa(elem.(int)))
				}				
			}
			return strs,nil
		}
	case uint32(t)<=uint32(Bytes32) && uint32(t)>=uint32(Bytes1):
		{
			var myByte = Byte(t)
			out,err := myByte.fuzz(isElem)
			if err!=nil{
				return nil,err
			}
			v := out

			return v,nil
		}
	case uint32(t) == uint32(Bytes):
		{
			var myByte = BYtes(t)
			out,err := myByte.fuzz(isElem)
			if err!=nil{
				return nil,err
			}
			v := out
			return v,nil
		}
	case uint32(t) == uint32(String):
		{
			var myString = STring(t)
			out,err := myString.fuzz(isElem)
			if err!=nil{
				return nil,err
			}
			v := out
			return v,nil
		}
	case uint32(t)==uint32(Address):
		{
			var myAddr = ADdress(t)
			out,err := myAddr.fuzz(isElem)
			if err!=nil{
				return nil,err
			}
			v := out
			return v,nil
		}
	case uint32(t) ==uint32(Bool):
		{
			var myBool = BOol(t)
			out,err := myBool.fuzz(isElem)
			if err!=nil{
				return nil,err
			}
			v := out
			return v,nil
		}
	default:
			return nil,SWICTH_DEFAULT_ERROR(nil)
	}
	return nil,nil
}