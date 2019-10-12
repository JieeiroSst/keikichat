package session

import (
	"log"
	"net/http"
	"keiki/controllers/User"
)
var users=User.User{
	Name:             "",
	Email:            "",
	Username:         "",
	Password:         "",
	VerifiedPassword: "",
}
// gets form values and checks user credentials
func login(w http.ResponseWriter,r *http.Request){
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
func logout(w http.ResponseWriter,r *http.Request){
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
	http.Redirect(w,r,"/",http.StatusFound)
}
//serves the index page
func index(w http.ResponseWriter,r *http.Request){
	session,err:=store.Get(r,"cookie-name")
	if err!=nil{
		log.Fatal(err)
	}
	user,_:=getUser(session)
	_ = tpl.ExecuteTemplate(w, "index.html", user)
}
// secret displays the secret message for authorized users
// serves the secret page if authentication is successful
func secret(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user ,_:= getUser(session)

	if auth := user.Authenticated; !auth {
		session.AddFlash("You don't have access!")
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/forbidden", http.StatusFound)
		return
	}

	_ = tpl.ExecuteTemplate(w, "secret.html", users.Username)
}
// shows an error when accessing without authentication
func forbidden(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	flashMessages := session.Flashes()
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = tpl.ExecuteTemplate(w, "forbidden.html", flashMessages)
}