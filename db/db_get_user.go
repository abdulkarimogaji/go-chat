package db

func (s *Store) GetUserByUsername(username string) (User, error) {
	row := s.DB.QueryRow(getUserByUsernameQuery, username)
	var user User
	err := row.Scan(&user.Id, &user.Username, &user.Fullname, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
