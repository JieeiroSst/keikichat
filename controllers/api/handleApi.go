package api

import (
	"database/sql"
	"github.com/gorilla/mux"
	"keiki/controllers/User"
	_ "keiki/database"
	"log"
	"net/http"
	_"keiki/models/toolbox"
)
var (
	user=User.User{}
	db *sql.DB
)

//GetEventAccounts
func GetEventAccounts(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","text/json")
	result,err:=db.Query("SELECT * from user")
	if err!=nil{
		log.Fatal(err)
	}
	defer result.Close()
	for result.Next(){
		err:=result.Scan(user.Id,user.Name,user.Username,user.Password,user.VerifiedPassword)
		if err!=nil{
			log.Fatal(err)

		}
	}
}

//GetEnventAcoount
func GetEnventAcoount(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","text/json")
	params := mux.Vars(r)
	result,err:=db.Query("SELECT id,Name,Email,UserName,Password,VerifiedPassword FROM user where Id=?", user.Id==params["Id"])
	if err!=nil{
		log.Fatal(err)
	}
	defer result.Close()
	for result.Next(){
		err:=result.Scan(user.Id,user.Name,user.Username,user.Password,user.VerifiedPassword)
		if err!=nil{
			log.Fatal(err)
		}
	}
}

func GetAccountLogin(w http.ResponseWriter,r*http.Request){
	w.Header().Set("Content-Type","text/json")
	params := mux.Vars(r)
	result,err:= db.Query("SELECT id,Email,Password FROM user where id=?",user.Id==params["Id"])
	if err!=nil{
		log.Fatal(err)
	}
	defer result.Close()
	for result.Next(){
		err:=result.Scan(user.Id,user.Name,user.Username,user.Password)
		if err!=nil{
			log.Fatal(err)
		}
	}
}

func PostAccount(w http.ResponseWriter,r*http.Request){
	w.Header().Set("Content-Type","text/json")
	parms:= mux.Vars(r)
	result,err:= db.Query("insert into user (Id,Name,Email,UserName,Password,VerifiedPassword) " +
		"VALUE (id=?,Name=?,Email=?,UserName=?,Password=?,VerifiedPassword=?)",
		user.Id==parms["Id"],user.Name==parms["Name"],user.Username==parms["UserName"],user.Password==parms["Password"],user.VerifiedPassword==parms["VerifiedPassword"])
	if err!=nil{
		log.Fatal(err)
	}
	result.Close()
	for result.Next(){

	}
}

func PutAccount(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","text/json")
	parms:= mux.Vars(r)
	result,err:=db.Query("update user set Name=?,Email=?,UserName=?,Password=?,VerifiedPassword=? where id=?",
		user.Id==parms["id"],user.Name==parms["Name"],user.Email==parms["Email"],user.Username==parms["UserName"],user.Password==parms["Password"],user.VerifiedPassword==parms["VerifiedPassword"])
	if err!=nil{
		log.Fatal(err)
	}
	result.Close()
	for result.Next(){
		err:=result.Scan(&user.Name,&user.Email,&user.Username,&user.Password,&user.VerifiedPassword)
		if err!=nil{
			log.Fatal(err)
		}
	}
}

func DeleteAccount(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Tye","text/json")
	parms:= mux.Vars(r)
	result,err:=db.Query("delete from user where id=?",user.Id==parms["Id"])
	if err!=nil{
		log.Fatal(err)
	}
	result.Close()
}

func HashPassword(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Tye","text/json")
	parms:=mux.Vars(r)
	result,err:=db.Query("update user set Password=MD5(Password)",user.Password==parms["Password"])
	if err!=nil{
		log.Fatal(err)
	}
	defer result.Close()
}

func HashVerifiedPassword(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Tye","text/json")
	parms:=mux.Vars(r)
	result,err:=db.Query("update user set VerifiedPassword=MD5(VerifiedPassword)",user.VerifiedPassword==parms["VerifiedPassword"])
	if err!=nil{
		log.Fatal(err)
	}
	defer result.Close()
}
