package chat

import (
	"context"

	"github.com/nqxcode/chat_microservice/internal/model"
	"github.com/nqxcode/chat_microservice/internal/service/log/constants"
)

func (s *service) SendMessage(ctx context.Context, message *model.Message) (int64, error) {
	var id int64
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error

		id, errTx = s.messageRepository.Create(ctx, message)
		if errTx != nil {
			return errTx
		}

		err := s.logService.Create(ctx, &model.Log{
			Message: constants.SendMessage,
			Payload: message,
		})

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}
