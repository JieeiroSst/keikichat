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
	urls=config.Mysqlurls
	port=config.Mysqlport
	mysqlDB=config.Mysqldb
)

func init(){
	db,err:=sql.Open("nameDB","userDB:userpass@tcp(urls:port)/mysqlDB")
	if err!=nil{
		log.Fatal(err)
	}
	defer db.Close()

	if err:=db.Ping(); err!=nil{
		log.Fatal(err)
	}
	log.Println("connected to mysql database")
}

