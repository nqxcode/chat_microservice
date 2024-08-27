package chat

import (
	"context"

	"github.com/nqxcode/platform_common/pagination"

	"github.com/nqxcode/chat_microservice/internal/model"
	"github.com/nqxcode/chat_microservice/internal/service/audit_log/constants"
)

func (s *service) GetMessages(ctx context.Context, chatID int64, limit *pagination.Limit) ([]model.Message, error) {
	var messages []model.Message
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error

		messages, errTx = s.messageRepository.Get(ctx, chatID, limit)
		if errTx != nil {
			return errTx
		}

		err := s.auditLogService.Create(ctx, &model.Log{
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
