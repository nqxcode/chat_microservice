package converter

import (
	desc "github.com/nqxcode/chat_microservice/pkg/chat_v1"

	"github.com/nqxcode/platform_common/pagination"
)

// ToLimitFromDesc convert to pagination limit
func ToLimitFromDesc(message *desc.Limit) *pagination.Limit {
	if message == nil {
		return &pagination.Limit{}
	}
	return &pagination.Limit{
		Offset: message.Offset,
		Limit:  message.Limit,
	}
}
