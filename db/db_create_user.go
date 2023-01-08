package db

import (
	"time"
)

func (s *Store) CreateUser(user *User) (User, error) {
	prep, err := s.DB.Prepare(createUserQuery)
	if err != nil {
		return User{}, err
	}

	// set timestamps
	user.CreatedAt = time.Now()

	res, err := prep.Exec(user.Username, user.Fullname, user.Email, user.Password, user.CreatedAt)
	if err != nil {
		return User{}, err
	}
	newId, err := res.LastInsertId()

	if err == nil {
		user.Id = newId
	}
	return *user, nil
}
