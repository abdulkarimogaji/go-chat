package db

import "time"

type User struct {
	Id        int64
	Username  string
	Fullname  string
	Email     string
	Password  string
	CreatedAt time.Time
}
