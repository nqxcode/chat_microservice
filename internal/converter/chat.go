package converter

import (
	"github.com/nqxcode/chat_microservice/internal/model"
	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/nqxcode/chat_microservice/pkg/chat_v1"
)

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

func ToChatInfoFromService(info model.ChatInfo) *desc.ChatInfo {
	return &desc.ChatInfo{
		Name:    info.Name,
		UserIds: info.UserIDs,
	}
}

func ToChatInfoFromDesc(info *desc.ChatInfo) *model.ChatInfo {
	return &model.ChatInfo{
		Name:    info.Name,
		UserIDs: info.UserIds,
	}
}
