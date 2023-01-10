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
	user.UpdatedAt = time.Now()
	user.Role = 1

	res, err := prep.Exec(user.Username, user.Fullname, user.Email, user.Password, user.Role, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return User{}, err
	}
	newId, err := res.LastInsertId()

	if err == nil {
		user.Id = newId
	}
	return *user, nil
}
