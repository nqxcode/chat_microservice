package converter

import (
	"github.com/nqxcode/chat_microservice/internal/model"
	modelRepo "github.com/nqxcode/chat_microservice/internal/repository/message/model"
)

func ToMessageFromRepo(m *modelRepo.Message) *model.Message {
	return &model.Message{
		ID:        m.ID,
		ChatID:    m.ChatID,
		From:      m.From,
		Message:   m.Message,
		SentAt:    m.SentAt,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func ToManyChatToUserFromRepo(models []modelRepo.Message) []model.Message {
	result := make([]model.Message, 0, len(models))
	for _, m := range models {
		result = append(result, *ToMessageFromRepo(&m))
	}

	return result
}
