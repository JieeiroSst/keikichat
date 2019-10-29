package main

import (
	"fmt"
	"keiki/controllers/api"
	_ "keiki/database"
	"keiki/models/cache"
	"keiki/models/cookie"
	"keiki/models/httplibs"
	"keiki/models/session"
	"keiki/controllers/User"
	"keiki/routers"
	"log"
	"net/http"
)

var port =":8080"
var u *User.User
var w http.ResponseWriter
func main(){
	fmt.Println("Loading...")
	r:=routers.Router()
	fHome:=http.FileServer(http.Dir("./views/home/assets"))
	fLogin:=http.FileServer(http.Dir("./views/Login/"))
	fSignUp:=http.FileServer(http.Dir("./views/sign-up/assets"))
	fWall:=http.FileServer(http.Dir("./views/wallPage/assets"))

	//session
	r.HandleFunc("/",session.Home)
	r.HandleFunc("/",session.Login)
	r.HandleFunc("/",session.Logout)

	//cookie
	r.HandleFunc("/create",cookie.CreateCookie)
	r.HandleFunc("/read",cookie.ReadCookie)
	r.HandleFunc("/delete",cookie.DeleteCookie)
	r.HandleFunc("/",cookie.GetUserAccount)
	if cookie.SetSession(u,w)==false{
		cookie.ClearSession(w)
	}


	//export css and js to client
	r.Handle("./views/home/assets",http.StripPrefix("views/home/assets", fHome))
	r.Handle("./views/Login/",http.StripPrefix("views/Login/", fLogin))
	r.Handle("./views/sign-up/",http.StripPrefix("views/sign-up/", fSignUp))
	r.Handle("./views/wallPage/assets/",http.StripPrefix("views/wallPage/assets", fWall))

	//cache
	r.HandleFunc("/",cache.GetFromCache)

	//api connect from server to client
	r.HandleFunc("/",api.GetAccountLogin)
	r.HandleFunc("/",api.GetEnventAcoount)
	r.HandleFunc("/",api.GetEventAccounts)
	r.HandleFunc("/",api.PostAccount)
	r.HandleFunc("/",api.PutAccount)
	r.HandleFunc("/",api.DeleteAccount)

	//render views to client
	r.HandleFunc("/keiki", httplibs.Renderhome)
	r.HandleFunc("/keiki/login", httplibs.Renderlogin)
	r.HandleFunc("/keiki/signup", httplibs.Rendersignup)
	r.HandleFunc("/keiki/{id}", httplibs.Renderwall)

	if err:=http.ListenAndServe(port,r);err!=nil{
		log.Fatal(err)
	}
}