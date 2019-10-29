package cookie

import (
	"bytes"
	"encoding/json"
	"keikibook/account"
	"log"
	"net/http"
	"keiki/controllers/User"
)
var u *User.User
func GetUserAccount(w http.ResponseWriter ,r* http.Request){
	cookie,err:=r.Cookie("first-cookie")
	if err!=nil{
		log.Fatal(err)
	}else{
		cookieValue := make(map[string]string)
		err = cookieHandler.Decode("first-cookie", cookie.Value,
			&cookieValue)
		if err != nil{
			u.Id=cookieValue["ID"]
			u.Email=cookieValue["Email"]
			u.Password=cookieValue["Password"]
		}
	}
}
func SetSession(u *User.User, w http.ResponseWriter) bool{
	value := map[string]string{
		"ID": u.Id,
		"Email":u.Email,
		"Password":u.Password,
	}
	encoded, err := cookieHandler.Encode("first-cookie", value)
	if err == nil {
		cookie := &http.Cookie{
		Name: "first-cookie",
		Value: encoded,
		Path: "/",
	}
		http.SetCookie(w, cookie)
	}
	return false
}

func LoginOnNonAccountSignUp(w http.ResponseWriter,r *http.Request){
	var jsonStr=[]byte(`{
		"ID":"`+	u.Id+`",
		"Email"`+ 	u.Email+`",
		"Username"`+ u.Username+`",
		"Password"`+ u.Password+`",
		"VerifiedPassword"`+u.VerifiedPassword+`"
	}`)
	var USER_ROOT_URL string
	request,_:=http.NewRequest("POST",USER_ROOT_URL+"login",bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, _:= client.Do(request)

	decoder := json.NewDecoder(response.Body)
	if err:= decoder.Decode(&account.Login{}); err != nil {
		log.Fatal(err)
	}

}

func ClearSession(w http.ResponseWriter){
	cookie:=&http.Cookie{
		Name:       "first-cookie",
		Value:      "",
		Path:       "/",
		MaxAge:     -1,
	}
	http.SetCookie(w,cookie)
}
