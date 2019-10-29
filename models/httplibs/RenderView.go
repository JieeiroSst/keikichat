package httplibs

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"keiki/controllers/User"
	"keiki/models/cookie"
	"log"
	"net/http"
)


func Renderhome(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/html")
	w.WriteHeader(http.StatusCreated)
	tpl,err:=template.ParseFiles("../../views/home/index.html")
	if err!=nil{
		log.Fatal(err)
	}
	_ = tpl.Execute(w, nil)
}
func Renderlogin(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	u:=User.User{}
	_ = cookie.GetUserAccount(r)
	json.NewEncoder(w).Encode(u)
	w.WriteHeader(http.StatusCreated)
	if r.Method=="GET" {
		tpl, err := template.ParseFiles("../../views/Login/index.html")
		if err != nil {
			log.Fatal(err)
		}
		_ = tpl.Execute(w, u)
	}else {
		_ = r.ParseForm()
		fmt.Println("Email",r.Form["Email"])
		fmt.Println("Password",r.Form["Password"])
	}

}
func Rendersignup(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	u:= User.User{}
	json.NewEncoder(w).Encode(u)
	w.WriteHeader(http.StatusCreated)
	if r.Method=="GET" {
		tpl, err := template.ParseFiles("../../views/sign-up/index.html")
		if err != nil {
			log.Fatal(err)
		}
		_ = tpl.Execute(w, u)
	}else{
		_ = r.ParseForm()
		fmt.Println("Full Name",r.Form["Name"])
		fmt.Println("Email",r.Form["Email"])
		fmt.Println("UserName",r.Form["Username"])
		fmt.Println("Password",r.Form["Password"])
		fmt.Println("Verified Password",r.Form["Verified Password"])
	}
}

func Renderwall(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	u:= User.User{}
	json.NewEncoder(w).Encode(u)
	ID := u.Username
	w.WriteHeader(http.StatusCreated)
	vars:=mux.Vars(r)
	_ = vars[ID]
	tpl,err:=template.ParseFiles("../../views/wallPage/index.html")
	if err!=nil{
		log.Fatal(err)
	}
	_=tpl.Execute(w,u)
}