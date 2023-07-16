package repository

const (
	queryRegistration = `
  INSERT INTO users (login, password_hash) VALUES ($1, $2)
  `

	queryLogin = `
  SELECT
  id AS id,
  password_hash AS password_hash
  FROM users WHERE login = $1
  `
)
