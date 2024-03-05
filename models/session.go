package models

import (
	"time"
)

// Session used to get the session information
type Session struct {
	UserID     uint      `json:"user_id"`
	Token      string    `json:"Token"`
	ExpireAt   time.Time `json:"expire_at"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedOn time.Time `json:"modified_on"`
} //name @Session

// Claims Retrive claims from the token string
type Claims struct {
	UserName    string    `json:"username"`
	Password    string    `json:"password"`
	ExpiaryTime time.Time `json:"expiryTime"`
} //name @Claims
