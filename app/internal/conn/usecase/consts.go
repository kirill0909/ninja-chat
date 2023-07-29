package usecase

const (
	violatesForeignKeyCode   = "23503"
	sendMessageNonExistsUser = "attampt to send a message to a user(%d) who does not exists"
	sendMessagePGError       = "unable to save message in db"
	sendMessageRedisError    = "unable to send message in redis"
)
