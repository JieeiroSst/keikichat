package session

import (
	"github.com/gorilla/sessions"
)
import "keiki/controllers/User"

type NoFound struct{
	Response string `json:"response"`
}

var user = User.User{}
func getUser(s *sessions.Session) (User.User,NoFound) {
	val := s.Values[user]

	_, ok := val.(User.User)
	if !ok {
		return _,NoFound{"no find user"}
	}
	return User.User{},_
}
