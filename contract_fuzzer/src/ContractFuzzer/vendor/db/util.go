package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

func getDb()(*sql.DB,error){
	if db,err :=sql.Open("mysql",dblinkinfo);err!=nil{
		return nil,err
	}else{
		return  db,nil
	}
}
