package main

import (
	"fmt"
	"github.com/gorilla/mux"
	_ "keiki/database"
	"keiki/models/config"
	"keiki/models/httplibs"
	_ "keiki/routers"
	"keiki/controllers/api"
	"log"
	"net/http"
)

var port =config.Mysqlport

func main(){
	fmt.Println("Loading...")
	r:=mux.NewRouter()
	fHome:=http.FileServer(http.Dir("views/home/assets"))
	fLogin:=http.FileServer(http.Dir("views/Login/"))
	fSignUp:=http.FileServer(http.Dir("views/sign-up/assets"))
	fWall:=http.FileServer(http.Dir("views/wallPage/assets"))

	//export css and js to client
	r.Handle("./views/home/assets",http.StripPrefix("views/home/assets", http.StripPrefix("/static/", fHome)))
	r.Handle("./views/Login/",http.StripPrefix("views/Login/", http.StripPrefix("/static/", fLogin)))
	r.Handle("./views/sign-up/",http.StripPrefix("views/sign-up/", http.StripPrefix("/static/", fSignUp)))
	r.Handle("./views/wallPage/assets",http.StripPrefix("views/wallPage/assets", http.StripPrefix("/static/", fWall)))

	//api connect from server to client
	r.HandleFunc("/",api.GetAccountLogin)
	r.HandleFunc("/",api.GetEnventAcoount)
	r.HandleFunc("/",api.GetEnventAccounts)
	r.HandleFunc("/",api.PostAccount)
	r.HandleFunc("",api.PutAccount)
	r.HandleFunc("/",api.CheckLogin)
	r.HandleFunc("/",api.DeleteAccount)

	//render views to client
	r.HandleFunc("/keiki/", httplibs.Renderhome).Methods("POST")
	r.HandleFunc("/keiki/login", httplibs.Renderlogin).Methods("POST")
	r.HandleFunc("/keiki/signup", httplibs.Rendersignup).Methods("POST")
	r.HandleFunc("/keiki/{id}", httplibs.Renderwall).Methods("POST")

	if err:=http.ListenAndServe(":"+port,r);err!=nil{
		log.Fatal(err)
	}
}