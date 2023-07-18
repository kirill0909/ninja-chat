package repository

const (
	queryRegistration = `
  INSERT INTO users.user (login, password_hash) VALUES ($1, $2)
  `

	queryLogin = `
  SELECT
  id AS id,
  password_hash AS password_hash
  FROM users.user WHERE login = $1
  `
)
