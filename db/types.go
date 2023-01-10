package db

import "time"

type User struct {
	Id        int
	Username  string
	Fullname  string
	Email     string
	Password  string
	Role      int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Room struct {
	Id          int
	UserID      int
	OtherUserId int
	RoomName    string
	Status      int
	LastChatID  int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
