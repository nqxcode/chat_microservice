package chat

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	desc "github.com/nqxcode/chat_microservice/pkg/chat_v1"
	"log"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {
	log.Printf("Delete chat: %v", req.GetId())
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
	//builderDelete := sq.Delete("\"chat_to_user\"").
	//	PlaceholderFormat(sq.Dollar).
	//	Where(sq.Eq{"chat_id": req.GetId()})
	//
	//query, args, err := builderDelete.ToSql()
	//if err != nil {
	//	log.Fatalf("failed to build query: %v", err)
	//}
	//
	//res, err := s.pool.Exec(ctx, query, args...)
	//if err != nil {
	//	log.Fatalf("failed to delete chat to user: %v", err)
	//}
	//
	//log.Printf("delete %d rows from chat_to_user", res.RowsAffected())
	//
	//builderDelete = sq.Delete("\"message\"").
	//	PlaceholderFormat(sq.Dollar).
	//	Where(sq.Eq{"chat_id": req.GetId()})
	//
	//query, args, err = builderDelete.ToSql()
	//if err != nil {
	//	log.Fatalf("failed to build query: %v", err)
	//}
	//
	//res, err = s.pool.Exec(ctx, query, args...)
	//if err != nil {
	//	log.Fatalf("failed to delete message: %v", err)
	//}
	//
	//log.Printf("delete %d rows from message", res.RowsAffected())
	//
	//builderDelete = sq.Delete("\"chat\"").
	//	PlaceholderFormat(sq.Dollar).
	//	Where(sq.Eq{"chat_id": req.GetId()})
	//
	//query, args, err = builderDelete.ToSql()
	//if err != nil {
	//	log.Fatalf("failed to build query: %v", err)
	//}
	//
	//res, err = s.pool.Exec(ctx, query, args...)
	//if err != nil {
	//	log.Fatalf("failed to delete chat: %v", err)
	//}
	//
	//err = tx.Commit(ctx)
	//if err != nil {
	//	return nil, status.Errorf(codes.Internal, "failed to commit transaction: %v", err)
	//}
	//
	//log.Printf("delete %d rows from chat", res.RowsAffected())
	//
	//if res.RowsAffected() == 0 {
	//	return nil, status.Errorf(codes.NotFound, "no chat deleted")
	//}

	return nil, nil
}
