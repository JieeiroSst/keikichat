package config

var (
	NameDB    ="mysql"
	Mysqluser = "root"
	Mysqlpass = " "
	Mysqldb   = "keiki"
	MyPort="3306"
	MyUrl="localhost"
)

type(
	Servers struct{
		Port string `json:"port"`
		Host string `json:"host"`
	}
	Databases struct{
		SystemDB string `json:"system_db"`
		User string `json:"user"`
		Password string `json:"password"`
		NameDB string `json:"name_db"`
	}
)
