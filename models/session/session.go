package session

import (
	"fmt"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
	"keiki/controllers/User"
	"keiki/models/cookie"
)
var users=User.User{
	Name:             "",
	Email:            "",
	Username:         "",
	Password:         "",
	VerifiedPassword: "",
}
var store *sessions.CookieStore
func init(){
	store=sessions.NewCookieStore([]byte("secret-key"))
}

func Home(w http.ResponseWriter, r *http.Request){
	session,_:=store.Get(r,"session-name")
	var authenticated interface{} = session.Values["authenticated"]
	if authenticated != nil {
		isAuthenticated := session.Values["authenticated"].(bool)
		if !isAuthenticated {
			http.Error(w, "You are unauthorized to view the page",
				http.StatusForbidden)
			return
		}
		_, _ = fmt.Fprintln(w, "Home Page")
	}else
	{
		http.Error(w, "You are unauthorized to view the page",
			http.StatusForbidden)
		return
	}
}
// gets form values and checks user credentials
func Login(w http.ResponseWriter,r *http.Request){
	session,err:=store.Get(r,"cookie-name")
	if err!=nil{
		log.Fatal(err)
	}

	// Where authentication could be done
	if r.FormValue("code")!="code"{
		if r.FormValue("code")==""{
			session.AddFlash("must enter a code")
		}
		session.AddFlash("the code was incorrect")
		err:=session.Save(r,w)
		if err!=nil{
			log.Fatal(err)
		}
		http.Redirect(w,r,"/forbidden",http.StatusFound)
		return
	}
	users.Username=r.FormValue("Username")
	users.Password=r.FormValue("Password")
	session.Values["user"]=users
	err=session.Save(r,w)
	if err!=nil{
		log.Fatal(err)
	}
	http.Redirect(w,r,"/secret",http.StatusFound)
}
//removes user credentials and clears the session cookie
func Logout(w http.ResponseWriter,r *http.Request){
	session,err:=store.Get(r,"cookie-name")
	if err!=nil{
		log.Fatal(err)
	}
	session.Values["user"]=users
	session.Options.MaxAge=-1

	err=session.Save(r,w)
	if err!=nil{
		log.Fatal(err)
	}
	cookie.ClearSession(w)
	http.Redirect(w,r,"/",http.StatusFound)
}

