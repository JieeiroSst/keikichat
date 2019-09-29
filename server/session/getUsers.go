package session

import "github.com/gorilla/sessions"

func getUser(s *sessions.Session) user {
	val := s.Values["user"]
	var User = user{}
	User, ok := val.(user)
	if !ok {
		return user{Authenticated:false}
	}
	return User
}
