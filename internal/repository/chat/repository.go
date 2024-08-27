package chat

import (
	"context"

	"github.com/nqxcode/chat_microservice/internal/model"
	"github.com/nqxcode/chat_microservice/internal/repository"
	"github.com/nqxcode/chat_microservice/internal/repository/chat/converter"
	modelRepo "github.com/nqxcode/chat_microservice/internal/repository/chat/model"
	"github.com/nqxcode/platform_common/client/db"

	sq "github.com/Masterminds/squirrel"
)

const (
	tableName = "chat"

	idColumn        = "chat_id"
	nameColumn      = "name"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db db.Client
}

// NewRepository new chat repository
func NewRepository(db db.Client) repository.ChatRepository {
	return &repo{db: db}
}

// Create chat
func (r *repo) Create(ctx context.Context, model *model.ChatInfo) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn).
		Values(model.Name).
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

// Delete by id
func (r *repo) Delete(ctx context.Context, id int64) error {
	builder := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     tableName + "_repository.Delete",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

// Get chat by id
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
		Name:     tableName + "_repository.Find",
		QueryRaw: query,
	}

	var chat modelRepo.Chat
	err = r.db.DB().ScanOneContext(ctx, &chat, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToChatFromRepo(&chat), nil
}
