package chat

import (
	"context"
	"github.com/nqxcode/chat_microservice/internal/model"
	"github.com/nqxcode/chat_microservice/internal/service/log/constants"
	"github.com/nqxcode/platform_common/pagination"
)

func (s *service) GetMessages(ctx context.Context, id int64, limit *pagination.Limit) ([]model.Message, error) {
	var messages []model.Message
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error

		messages, errTx = s.messageRepository.Get(ctx, id, limit)
		if errTx != nil {
			return errTx
		}

		err := s.logService.Create(ctx, &model.Log{
			Message: constants.GetMessages,
			Payload: messages,
		})

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return messages, nil
}
