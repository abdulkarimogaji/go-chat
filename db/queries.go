package db

const createTableUser = `CREATE TABLE IF NOT EXISTS user (
  id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  username VARCHAR(50) UNIQUE NOT NULL,
  fullname VARCHAR(200) NOT NULL,
  email VARCHAR(50) NULL,
  password TEXT NOT NULL,
  role int NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
)`

const createTableRoom = `CREATE TABLE IF NOT EXISTS room (
  id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  user_id INT NOT NULL,
  other_user_id INT NOT NULL,
  last_chat_id INT,
  room_name VARCHAR(50),
  status INT NOT NULL,
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

const createUserQuery = `INSERT INTO user(username, fullname, email, password, role, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?)`

const getUserByUsernameQuery = `SELECT * FROM user WHERE username = ?`
