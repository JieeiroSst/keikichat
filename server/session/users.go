package session

import (
	"encoding/gob"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"html/template"
)

type user struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Authenticated bool `json:"Authenticated"`
}

var store *sessions.CookieStore
var tpl *template.Template

func init(){
	authKeyOne:=securecookie.GenerateRandomKey(64)
	encryptionKeyOne:=securecookie.GenerateRandomKey(32)

	store=sessions.NewCookieStore(
		authKeyOne,
		encryptionKeyOne,
		)
	store.Options=&sessions.Options{
		maxAge:60*15,
		httpOnly:true,
	}
	gob.Register(user{})
	tpl=template.Must(template.ParseGlob("templates/*.html"))
}