package converter

import (
	"github.com/nqxcode/chat_microservice/internal/model"
	modelRepo "github.com/nqxcode/chat_microservice/internal/repository/chat_to_user/model"
)

// ToChatToUserFromRepo convert to chat to user relation model
func ToChatToUserFromRepo(chatToUser *modelRepo.ChatToUser) *model.ChatToUser {
	return &model.ChatToUser{
		ID:        chatToUser.ID,
		ChatID:    chatToUser.ChatID,
		UserID:    chatToUser.UserID,
		CreatedAt: chatToUser.CreatedAt,
	}
}

// ToManyChatToUserFromRepo convert to many chat to user models
func ToManyChatToUserFromRepo(chatToUser []modelRepo.ChatToUser) []model.ChatToUser {
	result := make([]model.ChatToUser, 0, len(chatToUser))
	for i := range chatToUser {
		m := ToChatToUserFromRepo(&chatToUser[i])
		if m != nil {
			result = append(result, *m)
		}
	}

	return result
}
