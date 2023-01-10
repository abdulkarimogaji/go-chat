package db

const createTableUser = `CREATE TABLE IF NOT EXISTS user (
  id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  username VARCHAR(50) UNIQUE NOT NULL,
  fullname VARCHAR(200) NOT NULL,
  email VARCHAR(50) NULL,
  password TEXT NOT NULL,
  role int NOT NULL DEFAULT 1,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
)`

const createTableRoom = `CREATE TABLE IF NOT EXISTS room (
  id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  user_id INT NOT NULL,
  other_user_id INT NOT NULL,
  last_chat_id INT,
  room_name VARCHAR(50),
  status INT NOT NULL DEFAULT 1,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE (user_id, other_user_id)
)`

const createTableChat = `CREATE TABLE IF NOT EXISTS chat (
  id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  sender_id INT NOT NULL,
  room_id INT NOT NULL,
  message VARCHAR(200) NOT NULL,
  unread BOOLEAN NOT NULL,
  timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
)`

const createUserQuery = `INSERT INTO user(username, fullname, email, password) VALUES(?, ?, ?, ?)`

const getUserByUsernameQuery = `SELECT * FROM user WHERE username = ?`

const deleteUserQuery = `DELETE FROM user WHERE id = ?`

const createRoomQuery = `INSERT INTO room(user_id, other_user_id, room_name) VALUES(?, ?, ?)`

const getMyRoomsQuery = `SELECT * FROM room WHERE user_id = ? OR other_user_id = ?`

const deleteRoomQuery = `DELETE FROM room  WHERE id = ?`

const createChatQuery = `INSERT INTO chat(sender_id, room_id, message, unread) VALUES(?, ?, ?, ?)`

const markChatAsReadQuery = `UPDATE chat SET unread = ? WHERE id = ?`

const setLastChatIdQuery = `UPDATE room SET last_chat_id = ? update_at = now() WHERE id = ?`

const disableRoomQuery = `UPDATE room SET status = 1 update_at = CURRENT_TIMESTAMP() WHERE id = ?`

const makeUserAnAdmin = `UPDATE user SET role = 0 update_at = SELECT now() WHERE id = ?`
