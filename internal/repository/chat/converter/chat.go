package converter

import (
	"github.com/nqxcode/chat_microservice/internal/model"
	modelRepo "github.com/nqxcode/chat_microservice/internal/repository/chat/model"
)

func ToChatFromRepo(chat *modelRepo.Chat) *model.Chat {
	return &model.Chat{
		ID: chat.ID,
		Info: model.ChatInfo{
			Name: chat.Info.Name,
		},
		CreatedAt: chat.CreatedAt,
		UpdatedAt: chat.UpdatedAt,
	}
}

func ToChatInfoFromRepo(info modelRepo.ChatInfo) model.ChatInfo {
	return model.ChatInfo{
		Name:    info.Name,
		UserIDs: info.UserIDs,
	}
}
