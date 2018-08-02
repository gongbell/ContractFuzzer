package main

import (
	"database/sql"
	"fmt"
	"os"
)

func getAddrMap(query string)(*sql.Rows, error){
	db,err := getDb()
	if err!=nil{
		return  nil,err
	}
	rows,err:=db.Query(query)
	return  rows,err
}
func  main()  {
	var query = "select address,name from register"
	var out,_ = os.Create("../resource/addrmap.csv")
	rows,err := getAddrMap(query)
	if err!=nil{
		fmt.Printf("%s",err)
	}
	for rows.Next(){
		var address string
		var name  string
		rows.Columns()
		err = rows.Scan(&address,&name)
		str := fmt.Sprintf("%s,\t%s\n",address,name)
		out.Write([]byte(str))
	}
}