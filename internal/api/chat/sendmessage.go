package chat

import (
	"context"
	desc "github.com/nqxcode/chat_microservice/pkg/chat_v1"

	"github.com/golang/protobuf/ptypes/empty"
)

func (i *Implementation) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*empty.Empty, error) {
	//log.Printf("Chat ID: %d, Send message: %v from %v at %v", req.GetInfo().GetChatId(), req.GetInfo().GetMessage(), req.GetInfo().GetFrom(), req.GetInfo().GetTimestamp())
	//
	//builderSelect := sq.Select("chat_id").
	//	From("chat").
	//	PlaceholderFormat(sq.Dollar).
	//	Where(sq.Eq{"chat_id": req.GetInfo().GetChatId()})
	//
	//query, args, err := builderSelect.ToSql()
	//if err != nil {
	//	log.Fatalf("failed to build query: %v", err)
	//}
	//
	//var chatID int64
	//err = s.pool.QueryRow(ctx, query, args...).Scan(&chatID)
	//if err != nil {
	//	if errors.Is(err, pgx.ErrNoRows) {
	//		return nil, status.Error(codes.NotFound, "chat not found")
	//	}
	//	log.Fatalf("failed to select chat: %v", err)
	//}
	//
	//var sentAt *time.Time
	//if req.GetInfo().GetTimestamp() != nil {
	//	t := req.GetInfo().GetTimestamp().AsTime()
	//	sentAt = &t
	//}
	//
	//builderInsert := sq.Insert("\"message\"").
	//	PlaceholderFormat(sq.Dollar).
	//	Columns("chat_id", "message", "\"from\"", "sent_at").
	//	Values(req.GetInfo().GetChatId(), req.GetInfo().GetMessage(), req.GetInfo().GetFrom(), sentAt).
	//	Suffix("RETURNING message_id")
	//
	//query, args, err = builderInsert.ToSql()
	//if err != nil {
	//	log.Fatalf("failed to build query: %v", err)
	//}
	//
	//var messageID int64
	//err = s.pool.QueryRow(ctx, query, args...).Scan(&messageID)
	//if err != nil {
	//	log.Fatalf("failed to insert message: %v", err)
	//}
	//
	//log.Printf("inserted message with id: %d", messageID)

	return nil, nil
}
