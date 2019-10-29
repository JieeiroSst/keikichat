package tests

import (
	"keiki/controllers/User"
	"errors"
	"strconv"
)

var user1 *User.User
func Testpassword(password string ) bool{
	if len(password)<15 {
		_ = errors.New("password fail ")
		return true
	}
	return false
}

func TestUser(user *User.User) (string,bool) {
	if len(user.Username)<0 && len(user.Username)>11 {
		return "test user ok",true
	}
	if user.Password==user.VerifiedPassword{
		return "Password ok",true
	}
	return "fail test case account",false
}

func TestEmail(user *User.User) bool{
	if user.Email!=user1.Email{
		return false
	}
	return false
}