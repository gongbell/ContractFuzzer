package fuzz

import "fmt"

func FILE_OPEN_ERROR(err error)(error){
	return fmt.Errorf("file open error. %s",err)
}
func FILE_READ_ERROR(err error)(error){
	return fmt.Errorf("file read error. %s",err)
}
func FILE_WRITE_ERROR(err error)(error){
	return fmt.Errorf("file write error. %s",err)
}
func JSON_MARSHAL_ERROR(err error)(error){
	return fmt.Errorf("json marshal error. %s",err)
}
func JSON_UNMARSHAL_ERROR(err error)(error){
	return fmt.Errorf("json unmarshal error. %s",err)
}
func DYNAMIC_CAST_ERROR(ok bool)(error){
	if ok!=true{
		return fmt.Errorf("dynamic cast error. ")
	}else{
		return nil
	}
}
func SWICTH_DEFAULT_ERROR(err error) (error){
	if err!=nil{
		return fmt.Errorf("swictch default error. %s",err)
	}else{
		return fmt.Errorf("swictch default error. ")
	}
}
