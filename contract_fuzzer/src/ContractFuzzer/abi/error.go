// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package abi

import (
	"errors"
	"fmt"
	"reflect"
	// "log"
)

var (
	errBadBool = errors.New("abi: improperly encoded boolean value")
)

// formatSliceString formats the reflection kind with the given slice size
// and returns a formatted string representation.
func formatSliceString(kind reflect.Kind, sliceSize int) string {
	if sliceSize == -1 {
		return fmt.Sprintf("[]%v", kind)
	}
	return fmt.Sprintf("[%d]%v", sliceSize, kind)
}

// sliceTypeCheck checks that the given slice can by assigned to the reflection
// type in t.
func sliceTypeCheck(t Type, val reflect.Value) error {
	//log.Printf("error.go :%j, %s",t,val)
	// var tKind string
	// if t.IsSlice{
	// 	tKind = "slice"
	// }else{
	// 	tKind = "array"
	// }
	// log.Printf("error.go :%s, %s",tKind,val.Kind())
	if val.Kind() != reflect.Slice && val.Kind() != reflect.Array {
		return typeErr(formatSliceString(t.Kind, t.SliceSize), val.Type())
	}
	// log.Printf("val.Len():%d t.SliceSize:%d",val.Len(),t.SliceSize)
	if t.IsArray && val.Len() != t.SliceSize {
		return typeErr(formatSliceString(t.Elem.Kind, t.SliceSize), formatSliceString(val.Type().Elem().Kind(), val.Len()))
	}

	if t.Elem.IsSlice {
		if val.Len() > 0 {
			return sliceTypeCheck(*t.Elem, val.Index(0))
		}
	} else if t.Elem.IsArray {
		return sliceTypeCheck(*t.Elem, val.Index(0))
	}

	if elemKind := val.Type().Elem().Kind(); elemKind != t.Elem.Kind {
		return typeErr(formatSliceString(t.Elem.Kind, t.SliceSize), val.Type())
	}
	return nil
}

// typeCheck checks that the given reflection value can be assigned to the reflection
// type in t.
func typeCheck(t Type, value reflect.Value) error {
	// log.Printf("%s",t.Kind)
	// log.Printf("%s",value.Kind())
	// log.Printf("%v",value)

	if t.IsSlice || t.IsArray {
		return sliceTypeCheck(t, value)
	}
	if t.T == AddressTy&&value.Kind()==reflect.Slice&&t.Size == value.Len(){
		return nil
	}
	// log.Printf("%s",t.Kind)
	// log.Printf("%s",value.Kind())
	// log.Printf("%v",value)

	// Check base type validity. Element types will be checked later on.
	if t.Kind != value.Kind() {
		// log.Printf("t.Kind!=value.Kind()")
		return typeErr(t.Kind, value.Kind())
	}
	return nil
}

// varErr returns a formatted error.
func varErr(expected, got reflect.Kind) error {
	return typeErr(expected, got)
}

// typeErr returns a formatted type casting error.
func typeErr(expected, got interface{}) error {
	return fmt.Errorf("abi: cannot use %v as type %v as argument", got, expected)
}
