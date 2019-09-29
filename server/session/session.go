package session

import (
	"github.com/prometheus/common/log"
	"net/http"
)
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
		http.Redirect(w,r,"/forbiden",http.StatusFound)
		return
	}
	username:=r.FormValue("username")
	password:=r.FormValue("password")
	user:=&user{
		UserName:      username,
		Password:      password,
		Authenticated: true,
	}
	session.Values["user"]=user
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
	session.Values["user"]=user{}
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
	user:=getUser(session)
	tpl.ExecuteTemplate(w,"index.html",user)
}
// secret displays the secret message for authorized users
// serves the secret page if authentication is successful
func secret(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := getUser(session)

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

	tpl.ExecuteTemplate(w, "secret.html", user.UserName)
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
	tpl.ExecuteTemplate(w, "forbidden.html", flashMessages)
}