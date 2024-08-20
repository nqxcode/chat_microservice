package converter

import (
	"github.com/nqxcode/chat_microservice/internal/model"
	modelRepo "github.com/nqxcode/chat_microservice/internal/repository/chat/model"
)

func ToChatFromRepo(chat *modelRepo.Chat) *model.Chat {
	return &model.Chat{
		ID: chat.ID,
		Info: model.ChatInfo{
			Name: chat.Name,
		},
		CreatedAt: chat.CreatedAt,
		UpdatedAt: chat.UpdatedAt,
	}
}
