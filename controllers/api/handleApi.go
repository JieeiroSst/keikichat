package api

import (
	"database/sql"
	"github.com/gorilla/mux"
	"keiki/controllers/User"
	_ "keiki/database"
	"log"
	"net/http"
)
var (
	user=User.User{}
	db *sql.DB
)

//GetEnventAccounts
func GetEnventAccounts(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","text/json")
	result,err:=db.Query("SELECT * from user")
	if err!=nil{
		log.Fatal(err)
	}
	defer result.Close()
}

//GetEnventAcoount
func GetEnventAcoount(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","text/json")
	params := mux.Vars(r)
	result,err:=db.Query("SELECT id,Name,Email,UserName,Password,VerifiedPassword FROM user where id=?", params["id"])
	if err!=nil{
		log.Fatal(err)
	}
	defer result.Close()
	for result.Next(){

	}
}

func GetAccountLogin(w http.ResponseWriter,r*http.Request){
	w.Header().Set("Content-Type","text/json")
	params := mux.Vars(r)
	result,err:= db.Query("SELECT id,Email,Password FROM user where id=?",params["id"])
	if err!=nil{
		log.Fatal(err)
	}
	defer result.Close()
}

func PostAccount(w http.ResponseWriter,r*http.Request){
	w.Header().Set("Content-Type","text/json")
	parms:= mux.Vars(r)
	result,err:= db.Query("INSERT INTO user (id,Name,Email,UserName,Password,VerifiedPassword) " +
		"VALUE (id=?,Name=?,Email=?,UserName=?,Password=?,VerifiedPassword=?)",
		parms["id"],parms["Name"],parms["UserName"],parms["Password"],parms["VerifiedPassword"])
	if err!=nil{
		log.Fatal(err)
	}
	result.Close()
}

func PutAccount(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	parms:= mux.Vars(r)
	result,err:=db.Query("update user set Name=?,Email=?,UserName=?,Password=?,VerifiedPassword=? where id=?",
		parms["Name"],parms["Email"],parms["UserName"],parms["Password"],parms["VerifiedPassword"])
	if err!=nil{
		log.Fatal(err)
	}
	result.Close()
}

func DeleteAccount(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Tye","application/json")
	parms:= mux.Vars(r)
	result,err:=db.Query("delete from user where id=?",parms["id"])
	if err!=nil{
		log.Fatal(err)
	}
	result.Close()
}

func CheckLogin(w http.ResponseWriter,r*http.Request){
	w.Header().Set("Content-Tye","application/json")
}