package usecase

const (
	violatesForeignKeyCode      = "23503"
	saveMessageForNonExistsUser = "unable to save user(%d) message(%s), because user(%d) does not exists"
	saveMessagePGError          = "unable to save message in db"
	saveMessageRedisError       = "unable to save message in redis"
)
