package models

//"time"

// UserToken is used to get the claims to generatethe token
type UserToken struct {
	UserName string `json:"username"`
	Password string `json:"password"`
} // @name UsetToken

type Token struct {
	Token string `json:"token"`
} // @name Token
