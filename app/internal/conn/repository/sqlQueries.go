package repository

const (
	querySaveMessageText = `
  INSERT INTO messages.message_text(text) VALUES ($1) RETURNING id
  `
	querySaveMessageInfo = `
  INSERT INTO messages.message_info(sender_id, recipient_id, message_text_id) VALUES($1, $2, $3) RETURNING id
  `
)
