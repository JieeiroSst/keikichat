package config

import (
	"github.com/micro/go-micro/config"
)
var (
	NameDB    ="mysql"
	Mysqluser = "root"
	Mysqlpass = " "
	Mysqlurls = "127.0.0.1"
	Mysqldb   = "keikibook"
	Mysqlport ="8080"
)

func init(){
	_ = config.LoadFile("./../../conf/web.conf")

}