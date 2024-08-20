package converter

import (
	"database/sql"
	"github.com/nqxcode/chat_microservice/internal/model"
	desc "github.com/nqxcode/chat_microservice/pkg/chat_v1"
)

func ToMessageFromDesc(message *desc.Message) *model.Message {
	var sentAt sql.NullTime
	if message.GetTimestamp() != nil {
		sentAt = sql.NullTime{
			Time:  message.GetTimestamp().AsTime(),
			Valid: true,
		}
	}

	return &model.Message{
		ChatID:  message.ChatId,
		From:    message.From,
		Message: message.Message,
		SentAt:  sentAt,
	}
}
