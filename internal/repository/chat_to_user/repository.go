package chat

import (
	"context"

	"github.com/nqxcode/chat_microservice/internal/client/db"
	"github.com/nqxcode/chat_microservice/internal/model"
	"github.com/nqxcode/chat_microservice/internal/repository"
	"github.com/nqxcode/chat_microservice/internal/repository/chat_to_user/converter"
	modelRepo "github.com/nqxcode/chat_microservice/internal/repository/chat_to_user/model"

	sq "github.com/Masterminds/squirrel"
)

const (
	tableName = "chat_to_user"

	idColumn        = "chat_to_user_id"
	chatIDColumn    = "chat_id"
	userIDColumn    = "user_id"
	createdAtColumn = "created_at"
)

type repo struct {
	db db.Client
}

// NewRepository Create chat to user repository
func NewRepository(db db.Client) repository.ChatToUserRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, model *model.ChatToUser) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(chatIDColumn, userIDColumn).
		Values(model.ChatID, model.UserID).
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

func (r *repo) Get(ctx context.Context, chatID int64, limit repository.Limit) ([]model.ChatToUser, error) {
	builder := sq.Select(idColumn, chatIDColumn, userIDColumn, createdAtColumn).
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

	var chatToUser []modelRepo.ChatToUser
	err = r.db.DB().ScanAllContext(ctx, &chatToUser, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToManyChatToUserFromRepo(chatToUser), nil
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
