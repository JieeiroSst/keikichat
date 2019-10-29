package cookie

import
(
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/securecookie"

)

var cookieHandler *securecookie.SecureCookie
func init() {
	cookieHandler = securecookie.New(securecookie.
	GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))
}
func CreateCookie(w http.ResponseWriter, r *http.Request) {
	value := map[string]string{
	"username": "first-cookie",
	}
	base64Encoded, err := cookieHandler.Encode("key", value)
	if err != nil {
		cookie := &http.Cookie{
		Name: "first-cookie",
		Value: base64Encoded,
		Path: "/",
		HttpOnly:true,
	}
		http.SetCookie(w, cookie)
	}
	w.Write([]byte("new cookie created"))
}

func ReadCookie(w http.ResponseWriter, r *http.Request) {
	log.Printf("Reading Cookie..")
	cookie, err := r.Cookie("first-cookie")
	if cookie != nil && err == nil {
		value := make(map[string]string)
		if err = cookieHandler.Decode("key", cookie.Value, &value);
		err == nil{
		log.Println(fmt.Sprintf("Hello %v \n", value["Username"]))
		log.Fatal(cookie)
		}
	} else {
	log.Printf("Cookie not found..")
	log.Println(fmt.Sprint("Hello"))
	w.Write([]byte("read cookie website for client "))
	}
}

func DeleteCookie(w http.ResponseWriter,r *http.Request){
	c:=http.Cookie{
		Name:       "first-cookie",
		MaxAge:     -1,
	}
	http.SetCookie(w,&c)
	w.Write([]byte("old delete cookie"))
}