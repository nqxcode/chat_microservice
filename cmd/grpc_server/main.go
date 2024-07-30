package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/nqxcode/chat_microservice/internal/config"
	desc "github.com/nqxcode/chat_microservice/pkg/chat_v1"

	sq "github.com/Masterminds/squirrel"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

type server struct {
	desc.UnimplementedChatV1Server
	pool *pgxpool.Pool
}

func main() {
	flag.Parse()
	ctx := context.Background()

	err := config.Load(configPath)
	if err != nil {
		log.Printf("No %s file found, using environment variables: %v", configPath, err)
	}

	grpcConfig, err := config.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	pgConfig, err := config.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	fmt.Println("PG DSN: ", pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{pool: pool})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Chat name: %v, user ids: %v", req.GetName(), req.GetUserIdList())

	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to acquire connection: %v", err)
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to begin transaction: %v", err)
	}
	//nolint:errcheck
	defer tx.Rollback(ctx)

	builderInsert := sq.Insert("\"chat\"").
		PlaceholderFormat(sq.Dollar).
		Columns("name").
		Values(req.GetName()).
		Suffix("RETURNING chat_id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	var chatID int64
	err = s.pool.QueryRow(ctx, query, args...).Scan(&chatID)
	if err != nil {
		log.Fatalf("failed to insert chat: %v", err)
	}

	log.Printf("inserted chat with id: %d", chatID)

	for _, userID := range req.GetUserIdList() {
		builderInsert = sq.Insert("\"chat_to_user\"").
			PlaceholderFormat(sq.Dollar).
			Columns("chat_id", "user_id").
			Values(chatID, userID).
			Suffix("RETURNING chat_to_user_id")

		query, args, err = builderInsert.ToSql()
		if err != nil {
			log.Fatalf("failed to build query: %v", err)
		}

		var chatToUserID int64
		err = s.pool.QueryRow(ctx, query, args...).Scan(&chatToUserID)
		if err != nil {
			log.Fatalf("failed to insert chat to user: %v", err)
		}

		log.Printf("inserted chat to user with id: %d", chatToUserID)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to commit transaction: %v", err)
	}

	return &desc.CreateResponse{
		Id: chatID,
	}, nil
}

func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {
	log.Printf("Delete chat: %v", req.GetId())

	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to acquire connection: %v", err)
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to begin transaction: %v", err)
	}
	//nolint:errcheck
	defer tx.Rollback(ctx)

	builderDelete := sq.Delete("\"chat_to_user\"").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"chat_id": req.GetId()})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	res, err := s.pool.Exec(ctx, query, args...)
	if err != nil {
		log.Fatalf("failed to delete chat to user: %v", err)
	}

	log.Printf("delete %d rows from chat_to_user", res.RowsAffected())

	builderDelete = sq.Delete("\"message\"").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"chat_id": req.GetId()})

	query, args, err = builderDelete.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	res, err = s.pool.Exec(ctx, query, args...)
	if err != nil {
		log.Fatalf("failed to delete message: %v", err)
	}

	log.Printf("delete %d rows from message", res.RowsAffected())

	builderDelete = sq.Delete("\"chat\"").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"chat_id": req.GetId()})

	query, args, err = builderDelete.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	res, err = s.pool.Exec(ctx, query, args...)
	if err != nil {
		log.Fatalf("failed to delete chat: %v", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to commit transaction: %v", err)
	}

	log.Printf("delete %d rows from chat", res.RowsAffected())

	if res.RowsAffected() == 0 {
		return nil, status.Errorf(codes.NotFound, "no chat deleted")
	}

	return nil, nil
}

func (s *server) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*empty.Empty, error) {
	log.Printf("Chat ID: %d, Send message: %v from %v at %v", req.GetChatId(), req.GetMessage(), req.GetFrom(), req.GetTimestamp())

	// Делаем запрос на выборку записей из таблицы note
	builderSelect := sq.Select("chat_id").
		From("chat").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"chat_id": req.GetChatId()})

	query, args, err := builderSelect.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	var chatID int64
	err = s.pool.QueryRow(ctx, query, args...).Scan(&chatID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "chat not found")
		}
		log.Fatalf("failed to select chat: %v", err)
	}

	var sentAt *time.Time
	if req.GetTimestamp() != nil {
		t := req.GetTimestamp().AsTime()
		sentAt = &t
	}

	builderInsert := sq.Insert("\"message\"").
		PlaceholderFormat(sq.Dollar).
		Columns("chat_id", "message", "\"from\"", "sent_at").
		Values(req.GetChatId(), req.GetMessage(), req.GetFrom(), sentAt).
		Suffix("RETURNING message_id")

	query, args, err = builderInsert.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	var messageID int64
	err = s.pool.QueryRow(ctx, query, args...).Scan(&messageID)
	if err != nil {
		log.Fatalf("failed to insert message: %v", err)
	}

	log.Printf("inserted message with id: %d", messageID)

	return nil, nil
}
