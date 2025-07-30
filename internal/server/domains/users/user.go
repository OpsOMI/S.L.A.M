package users

import "time"

type User struct {
	ID          string
	Username    string
	Password    string
	Nickname    string
	PrivateCode string
	CreatedAt   time.Time
}
