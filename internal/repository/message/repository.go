package message

import (
	"context"
	"fmt"

	"github.com/nqxcode/chat_microservice/internal/model"
	"github.com/nqxcode/chat_microservice/internal/repository"
	"github.com/nqxcode/chat_microservice/internal/repository/message/converter"
	modelRepo "github.com/nqxcode/chat_microservice/internal/repository/message/model"
	"github.com/nqxcode/platform_common/client/db"
	"github.com/nqxcode/platform_common/pagination"

	sq "github.com/Masterminds/squirrel"
)

const (
	tableName = "message"

	idColumn        = "message_id"
	chatIDColumn    = "chat_id"
	fromColumn      = "from"
	messageColumn   = "message"
	sentAt          = "sent_at"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db db.Client
}

// NewRepository new message repository
func NewRepository(db db.Client) repository.MessageRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, model *model.Message) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(chatIDColumn, escape(fromColumn), messageColumn, sentAt).
		Values(model.ChatID, model.From, model.Message, model.SentAt).
		Suffix("RETURNING " + idColumn)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     tableName + "_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) Get(ctx context.Context, chatID int64, limit *pagination.Limit) ([]model.Message, error) {
	builder := sq.Select(idColumn, chatIDColumn, escape(fromColumn), messageColumn, sentAt, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{chatIDColumn: chatID}).
		Offset(limit.Offset).
		Limit(limit.Limit)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     tableName + "_repository.Get",
		QueryRaw: query,
	}

	var messages []modelRepo.Message
	err = r.db.DB().ScanAllContext(ctx, &messages, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToManyChatToUserFromRepo(messages), nil
}

func (r *repo) DeleteByChatID(ctx context.Context, chatID int64) error {
	builder := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{chatIDColumn: chatID})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     tableName + "_repository.DeleteByChatID",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func escape(value string) string {
	return fmt.Sprintf("\"%s\"", value)
}
