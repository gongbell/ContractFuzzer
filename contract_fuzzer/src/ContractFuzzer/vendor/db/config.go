package main

import "fmt"

const dbhost = "localhost"
const dbport = "3306"
const dbcharset = "utf8"
const dbusr  = "root"
const dbpwd  = "123456"
const dbname = "Contract2"

var dblinkinfo = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",dbusr,dbpwd,dbhost,dbport,dbname,dbcharset)