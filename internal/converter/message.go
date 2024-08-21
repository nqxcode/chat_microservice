package converter

import (
	"database/sql"
	"github.com/nqxcode/chat_microservice/internal/model"
	"github.com/nqxcode/chat_microservice/internal/pagination"
	desc "github.com/nqxcode/chat_microservice/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ToMessagesFromService convert to desc messages
func ToMessagesFromService(messages []model.Message) []*desc.Message {
	result := make([]*desc.Message, 0, len(messages))
	for _, message := range messages {

		var timestamp *timestamppb.Timestamp
		if message.SentAt.Valid {
			timestamp = timestamppb.New(message.SentAt.Time)
		}

		result = append(result, &desc.Message{
			From:      message.From,
			Message:   message.Message,
			Timestamp: timestamp,
			ChatId:    message.ChatID,
		})
	}

	return result
}

// ToMessageFromDesc convert to message model
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

// ToLimitFromDesc convert to pagination limit
func ToLimitFromDesc(message *desc.Limit) *pagination.Limit {
	return &pagination.Limit{
		Offset: message.Offset,
		Limit:  message.Limit,
	}
}
