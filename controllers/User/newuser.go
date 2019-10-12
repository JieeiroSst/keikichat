package User

type User struct {
	Name             string `json:"Name" form:"Name"`
	Email            string `json:"Email" form:"Email"`
	Username         string `json:"Username" form:"Username"`
	Password         string `json:"Password" form:"Password"`
	VerifiedPassword string `json:"VerifiedPassword" form:"VerifiedPassword"`
}

