package repository

import (
	"context"
	"database/sql"
	"log"
	"ninja-chat-core-api/internal/conn"
	models "ninja-chat-core-api/internal/models/conn"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type ConnPGRepo struct {
	db *sqlx.DB
}

func NewConnPGRepo(db *sqlx.DB) conn.PGRepo {
	return &ConnPGRepo{db: db}
}

func (r *ConnPGRepo) SendMessage(ctx context.Context, request models.SendMessageRequest) (
	result models.SendMessageResponse, err error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return
	}
	defer func() {
		err := tx.Rollback()
		if err != nil && !errors.Is(err, sql.ErrTxDone) {
			log.Println(err)
		}
	}()

	var messageID int
	if err = tx.GetContext(ctx, &messageID, querySendMessageText, request.Message); err != nil {
		return
	}

	var messageInfoID int
	if err = tx.GetContext(ctx, &messageInfoID, querySaveMessageInfo, request.UserID, request.RecipientID, messageID); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		return
	}

	return models.SendMessageResponse{Success: true, Code: 200}, nil
}
