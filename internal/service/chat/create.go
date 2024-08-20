package chat

import (
	"context"

	"github.com/nqxcode/chat_microservice/internal/model"
)

func (s *service) Create(ctx context.Context, info *model.ChatInfo) (int64, error) {
	var chatID int64
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		chatID, errTx = s.chatRepository.Create(ctx, info)
		if errTx != nil {
			return errTx
		}

		for _, userID := range info.UserIDs {
			_, errTx = s.chatToUserRepository.Create(ctx, &model.ChatToUser{
				ChatID: chatID,
				UserID: userID,
			})
			if errTx != nil {
				return errTx
			}
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return chatID, nil
}
