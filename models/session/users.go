package session


import (
	"encoding/gob"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"html/template"
	"keiki/controllers/User"
)

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
		Path:     "/",
		MaxAge: 900000,
		HttpOnly:true,
	}
	gob.Register(User.User{})
	tpl=template.Must(template.ParseGlob("templates/*.html"))
}