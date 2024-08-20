package chat

import (
	"context"
	"github.com/nqxcode/chat_microservice/internal/model"
	"github.com/nqxcode/chat_microservice/internal/service/log/constants"
)

func (s *service) Delete(ctx context.Context, id int64) error {

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error

		errTx = s.messageRepository.DeleteByChatID(ctx, id)
		if errTx != nil {
			return errTx
		}

		errTx = s.chatToUserRepository.DeleteByChatID(ctx, id)
		if errTx != nil {
			return errTx
		}

		errTx = s.chatRepository.Delete(ctx, id)
		if errTx != nil {
			return errTx
		}

		err := s.logService.Create(ctx, &model.Log{
			Message: constants.ChatDeleted,
			Payload: id,
		})

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
