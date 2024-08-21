package chat

import (
	"context"
	"github.com/nqxcode/chat_microservice/internal/model"
	"github.com/nqxcode/chat_microservice/internal/pagination"
	"github.com/nqxcode/chat_microservice/internal/service/log/constants"
)

func (s *service) Find(ctx context.Context, id int64) (*model.Chat, error) {
	var chat *model.Chat
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		chat, errTx = s.chatRepository.Find(ctx, id)
		if errTx != nil {
			return errTx
		}

		userIDs := make([]int64, 0)
		offset := uint64(0)
		limit := pagination.DefaultLimit
		for {
			chunk := make([]model.ChatToUser, 0)
			chunk, errTx = s.chatToUserRepository.Get(ctx, id, pagination.Limit{Offset: offset, Limit: limit})
			if errTx != nil {
				return errTx
			}
			if len(chunk) == 0 {
				break
			}

			for _, ctu := range chunk {
				userIDs = append(userIDs, ctu.UserID)
			}

			offset += uint64(len(chunk))
		}
		chat.Info.UserIDs = userIDs

		err := s.logService.Create(ctx, &model.Log{
			Message: constants.ChatFound,
			Payload: chat,
		})

		if err != nil {
			return err
		}

		return nil
	})

	return chat, err
}
