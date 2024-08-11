package chat

import (
	"context"
	"github.com/nqxcode/chat_microservice/internal/converter"
	"log"

	desc "github.com/nqxcode/chat_microservice/pkg/chat_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.chatService.Create(ctx, converter.ToChatInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	log.Printf("inserted chat with id: %d", id)

	//log.Printf("Chat name: %v, user ids: %v", req.GetInfo().GetName(), req.GetInfo().GetUserIds())
	//
	//conn, err := s.pool.Acquire(ctx)
	//if err != nil {
	//	return nil, status.Errorf(codes.Internal, "failed to acquire connection: %v", err)
	//}
	//defer conn.Release()
	//
	//tx, err := conn.Begin(ctx)
	//if err != nil {
	//	return nil, status.Errorf(codes.Internal, "failed to begin transaction: %v", err)
	//}
	//defer tx.Rollback(ctx) //nolint:errcheck
	//
	//builderInsert := sq.Insert("\"chat\"").
	//	PlaceholderFormat(sq.Dollar).
	//	Columns("name").
	//	Values(req.GetInfo().GetName()).
	//	Suffix("RETURNING chat_id")
	//
	//query, args, err := builderInsert.ToSql()
	//if err != nil {
	//	log.Fatalf("failed to build query: %v", err)
	//}
	//
	//var chatID int64
	//err = s.pool.QueryRow(ctx, query, args...).Scan(&chatID)
	//if err != nil {
	//	log.Fatalf("failed to insert chat: %v", err)
	//}
	//
	//log.Printf("inserted chat with id: %d", chatID)
	//
	//for _, userID := range req.GetInfo().GetUserIds() {
	//	builderInsert = sq.Insert("\"chat_to_user\"").
	//		PlaceholderFormat(sq.Dollar).
	//		Columns("chat_id", "user_id").
	//		Values(chatID, userID).
	//		Suffix("RETURNING chat_to_user_id")
	//
	//	query, args, err = builderInsert.ToSql()
	//	if err != nil {
	//		log.Fatalf("failed to build query: %v", err)
	//	}
	//
	//	var chatToUserID int64
	//	err = s.pool.QueryRow(ctx, query, args...).Scan(&chatToUserID)
	//	if err != nil {
	//		log.Fatalf("failed to insert chat to user: %v", err)
	//	}
	//
	//	log.Printf("inserted chat to user with id: %d", chatToUserID)
	//}
	//
	//err = tx.Commit(ctx)
	//if err != nil {
	//	return nil, status.Errorf(codes.Internal, "failed to commit transaction: %v", err)
	//}
	//
	//return &desc.CreateResponse{
	//	Id: chatID,
	//}, nil

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
