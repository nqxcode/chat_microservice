package log

import (
	"github.com/nqxcode/chat_microservice/internal/repository"
	def "github.com/nqxcode/chat_microservice/internal/service"
)

type service struct {
	logRepository repository.LogRepository
}

// NewService new log service
func NewService(logRepository repository.LogRepository) def.LogService {
	return &service{
		logRepository: logRepository,
	}
}
