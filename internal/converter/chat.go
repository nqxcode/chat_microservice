package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/nqxcode/chat_microservice/internal/model"
	desc "github.com/nqxcode/chat_microservice/pkg/chat_v1"
)

// ToChatFromService convert to chat model
func ToChatFromService(chat *model.Chat) *desc.Chat {
	var updatedAt *timestamppb.Timestamp
	if chat.UpdatedAt.Valid {
		updatedAt = timestamppb.New(chat.UpdatedAt.Time)
	}

	return &desc.Chat{
		Id:        chat.ID,
		Info:      ToChatInfoFromService(chat.Info),
		CreatedAt: timestamppb.New(chat.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

// ToChatInfoFromService convert to chat info model
func ToChatInfoFromService(info model.ChatInfo) *desc.ChatInfo {
	return &desc.ChatInfo{
		Name:    info.Name,
		UserIds: info.UserIDs,
	}
}

// ToChatInfoFromDesc to chat info model
func ToChatInfoFromDesc(info *desc.ChatInfo) *model.ChatInfo {
	return &model.ChatInfo{
		Name:    info.Name,
		UserIDs: info.UserIds,
	}
}
