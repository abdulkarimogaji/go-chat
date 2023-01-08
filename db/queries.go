package db

const createTableUser = `CREATE TABLE IF NOT EXISTS user (
  id BIGINT(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
  username VARCHAR(50) NOT NULL,
  fullname VARCHAR(200) NOT NULL,
  email VARCHAR(50) UNIQUE NOT NULL,
  password VARCHAR(50) NOT NULL,
  createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
)`

const createUserQuery = `INSERT INTO user(username, fullname, email, password, createdAt) VALUES(?, ?, ?, ?, ?)`
