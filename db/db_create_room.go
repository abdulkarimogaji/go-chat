package db

func (s *Store) CreateRoom(room *Room) (Room, error) {
	prep, err := s.DB.Prepare(createRoomQuery)
	if err != nil {
		return Room{}, err
	}

	res, err := prep.Exec(room.UserID, room.OtherUserId, room.RoomName)
	if err != nil {
		return Room{}, err
	}
	newId, err := res.LastInsertId()

	if err == nil {
		room.Id = int(newId)
	}
	return *room, nil
}
