package repository

const (
	queryRegistration = `
  INSERT INTO users (login, hashed_password) VALUES ($1, $2)
  `
)
