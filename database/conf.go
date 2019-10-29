package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"keiki/models/config"
	"log"
)

var(
	nameDB=config.NameDB
	userDB=config.Mysqluser
	userpass=config.Mysqlpass
	mysqlDB=config.Mysqldb
	myPort=config.MyPort
	myUrl=config.MyUrl
)


func init(){
	db,err:=sql.Open(nameDB,userDB+":"+userpass+"@("+myUrl+":"+myPort+")/"+mysqlDB)
	if err!=nil{
		log.Fatal(err)
	}
	defer db.Close()

	if err:=db.Ping(); err!=nil{
		log.Fatal(err)
	}

	log.Println("connected to mysql database")
}

