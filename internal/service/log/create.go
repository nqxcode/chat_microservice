package log

import (
	"context"
	"encoding/json"

	"github.com/nqxcode/chat_microservice/internal/model"
	"github.com/nqxcode/platform_common/helper/grpc"
)

func (s *service) Create(ctx context.Context, log *model.Log) error {
	ip, _ := grpc.ClientIP(ctx)
	jsonPayload, _ := json.Marshal(log.Payload)

	return s.logRepository.Create(ctx, &model.Log{
		Message: log.Message,
		Payload: jsonPayload,
		IP:      ip,
	})
}
