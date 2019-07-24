package model

import (
	"time"
)

var (
	UserTable = "users"
)

type User struct{
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Password string `json:"password"`
	Gender string `json:"gender"`
	Status string `json:"status"`
	LastLoginTime time.Time `json:"lastLoginTime"`
}