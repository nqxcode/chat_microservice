package chat

import (
	"context"
	"github.com/nqxcode/chat_microservice/internal/client/db"
	"github.com/nqxcode/chat_microservice/internal/model"
	"github.com/nqxcode/chat_microservice/internal/repository"

	sq "github.com/Masterminds/squirrel"

	"github.com/nqxcode/chat_microservice/internal/repository/chat/converter"
	modelRepo "github.com/nqxcode/chat_microservice/internal/repository/chat/model"
)

const (
	tableName = "chat"

	idColumn        = "id"
	nameColumn      = "name"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.ChatRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, info *model.ChatInfo) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn).
		Values(info.Name).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "chat_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.Chat, error) {
	builder := sq.Select(idColumn, nameColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "chat_repository.Get",
		QueryRaw: query,
	}

	var chat modelRepo.Chat
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&chat.ID, &chat.Info.Name, &chat.CreatedAt, &chat.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return converter.ToChatFromRepo(&chat), nil
}
