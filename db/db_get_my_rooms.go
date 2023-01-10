package db

func (s *Store) GetMyRooms(userId int) ([]Room, error) {
	rows, err := s.DB.Query(getMyRoomsQuery, userId)
	if err != nil {
		return []Room{}, err
	}
	var rooms []Room

	for rows.Next() {
		var room Room
		err = rows.Scan(&room.Id, &room.UserID, &room.OtherUserId, &room.LastChatID, &room.RoomName, &room.Status, &room.CreatedAt, &room.UpdatedAt)
		if err != nil {
			return rooms, err
		}
		rooms = append(rooms, room)
	}
	return rooms, nil
}
