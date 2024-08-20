package converter

import (
	"github.com/nqxcode/chat_microservice/internal/model"
	modelRepo "github.com/nqxcode/chat_microservice/internal/repository/message/model"
)

// ToMessageFromRepo convert to message model
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

// ToManyChatToUserFromRepo convert to many chat to user relation models
func ToManyChatToUserFromRepo(messages []modelRepo.Message) []model.Message {
	result := make([]model.Message, 0, len(messages))
	for i := range messages {
		m := ToMessageFromRepo(&messages[i])
		if m != nil {
			result = append(result, *m)
		}
	}

	return result
}
