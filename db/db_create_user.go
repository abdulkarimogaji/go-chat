package db

func (s *Store) CreateUser(user *User) (User, error) {
	prep, err := s.DB.Prepare(createUserQuery)
	if err != nil {
		return User{}, err
	}

	res, err := prep.Exec(user.Username, user.Fullname, user.Email, user.Password)
	if err != nil {
		return User{}, err
	}
	newId, err := res.LastInsertId()

	if err == nil {
		user.Id = int(newId)
	}
	return *user, nil
}
