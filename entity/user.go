package entity

import "time"

type User struct {
	Id          string
	Password    string
	DisplayName string
	ProfilePic  string
	Bio         string
	Email       string
	IsVerified  bool
	CreatedAt   time.Time
}
