package httplibs

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"html/template"
	"keiki/controllers/User"
	"log"
	"net/http"
)

func Renderlogin(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	u:=User.User{}
	json.NewEncoder(w).Encode(u)
	if r.Method=="POST" {
		w.WriteHeader(http.StatusCreated)
		tpl:=template.New("")
		tpl, _ = tpl.ParseFiles("./../../views/Login/index.html")
		_ = tpl.Execute(w, u) 	
	}else{
		w.WriteHeader(http.StatusNotFound)
		log.Fatal(r.Method)
		 w.Write([]byte("No found page index in file views"))
	}

}
func Renderhome(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	if r.Method=="POST" {
		w.WriteHeader(http.StatusCreated)
		tpl:=template.New("")
		tpl, _= tpl.ParseFiles("./../../views/home/index.html")
		_ = tpl.Execute(w, nil)
	}else{
		w.WriteHeader(http.StatusNotFound)
		log.Fatal(r.Method)
		w.Write([]byte("No found page index in file views"))
	}
}

func Rendersignup(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	u:= User.User{}
	json.NewEncoder(w).Encode(u)
	if r.Method=="POST" {
		w.WriteHeader(http.StatusCreated)
		tpl:=template.New("")
		tpl, _ = tpl.ParseFiles("./../../views/sign-up/index.html")
		_ = tpl.Execute(w, u)
	}else{
		w.WriteHeader(http.StatusNotFound)
		log.Fatal(r.Method)
		w.Write([]byte("No found page index in file views"))
	}
}

func Renderwall(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	u:= User.User{}
	json.NewEncoder(w).Encode(u)
	ID := u.Username
	if r.Method=="POST" {
		w.WriteHeader(http.StatusCreated)
		vars:=mux.Vars(r)
		_ = vars[ID]
		tpl:=template.New("")
		tpl,_= tpl.ParseFiles("./../../views/wallPage/index.html")
		_ = tpl.Execute(w, u)
	}else{
		w.WriteHeader(http.StatusNotFound)
		log.Fatal(r.Method)
		w.Write([]byte("No found page index in file views"))
	}
}