package chat

import (
	"context"

	"github.com/nqxcode/chat_microservice/internal/model"
)

func (s *service) Create(ctx context.Context, info *model.ChatInfo) (int64, error) {
	var id int64
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.chatRepository.Create(ctx, info)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}
