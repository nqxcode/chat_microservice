package tests

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/nqxcode/chat_microservice/internal/service/chat"
	"github.com/nqxcode/platform_common/pagination"
	"testing"

	"github.com/nqxcode/chat_microservice/internal/model"
	"github.com/nqxcode/chat_microservice/internal/repository"
	repoMocks "github.com/nqxcode/chat_microservice/internal/repository/mocks"
	"github.com/nqxcode/chat_microservice/internal/service"
	serviceSupport "github.com/nqxcode/chat_microservice/internal/service/chat/tests/support"
	"github.com/nqxcode/chat_microservice/internal/service/log/constants"
	serviceMocks "github.com/nqxcode/chat_microservice/internal/service/mocks"
	"github.com/nqxcode/platform_common/client/db"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestGetMessages(t *testing.T) {
	t.Parallel()

	type chatRepositoryMock func(mc *minimock.Controller) repository.ChatRepository
	type chatToUserRepositoryMock func(mc *minimock.Controller) repository.ChatToUserRepository
	type messageRepositoryMock func(mc *minimock.Controller) repository.MessageRepository
	type logServiceMock func(mc *minimock.Controller) service.LogService

	type input struct {
		ctx    context.Context
		chatID int64
		limit  pagination.Limit
	}

	type expected struct {
		resp []model.Message
		err  error
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		chatID    = gofakeit.Int64()
		messageID = gofakeit.Int64()
		from      = gofakeit.Email()
		msg       = gofakeit.Generate("???")
		sentAt    = sql.NullTime{Valid: true, Time: gofakeit.Date()}
		createdAt = gofakeit.Date()
		updatedAt = sql.NullTime{Valid: true, Time: gofakeit.Date()}

		limit = pagination.Limit{Limit: pagination.DefaultLimit}

		repoErr = fmt.Errorf("repo error")
	)

	messages := []model.Message{
		{
			ID:        messageID,
			ChatID:    chatID,
			From:      from,
			Message:   msg,
			SentAt:    sentAt,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		},
	}

	cases := []struct {
		name                     string
		input                    input
		expected                 expected
		chatRepositoryMock       chatRepositoryMock
		chatToUserRepositoryMock chatToUserRepositoryMock
		messageRepositoryMock    messageRepositoryMock
		logServiceMock           logServiceMock
		txManagerFake            db.TxManager
	}{
		{
			name: "success case",
			input: input{
				ctx:    ctx,
				chatID: chatID,
			},
			expected: expected{
				err:  nil,
				resp: messages,
			},
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repoMocks.NewChatRepositoryMock(mc)
				return mock
			},
			chatToUserRepositoryMock: func(mc *minimock.Controller) repository.ChatToUserRepository {
				mock := repoMocks.NewChatToUserRepositoryMock(mc)
				return mock
			},
			logServiceMock: func(mc *minimock.Controller) service.LogService {
				mock := serviceMocks.NewLogServiceMock(mc)
				mock.CreateMock.Expect(ctx, &model.Log{
					Message: constants.GetMessages,
					Payload: messages,
				}).Return(nil)
				return mock
			},
			messageRepositoryMock: func(mc *minimock.Controller) repository.MessageRepository {
				mock := repoMocks.NewMessageRepositoryMock(mc)

				mock.GetMock.Expect(ctx, chatID, &limit).Return(messages, nil)
				return mock
			},
			txManagerFake: serviceSupport.NewTxManagerFake(),
		},
		{
			name: "service error case",
			input: input{
				ctx:    ctx,
				chatID: chatID,
			},
			expected: expected{
				err:  repoErr,
				resp: nil,
			},
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repoMocks.NewChatRepositoryMock(mc)
				return mock
			},
			chatToUserRepositoryMock: func(mc *minimock.Controller) repository.ChatToUserRepository {
				mock := repoMocks.NewChatToUserRepositoryMock(mc)
				return mock
			},
			logServiceMock: func(mc *minimock.Controller) service.LogService {
				mock := serviceMocks.NewLogServiceMock(mc)
				return mock
			},
			messageRepositoryMock: func(mc *minimock.Controller) repository.MessageRepository {
				mock := repoMocks.NewMessageRepositoryMock(mc)

				mock.GetMock.Expect(ctx, chatID, &limit).Return(nil, repoErr)
				return mock
			},
			txManagerFake: serviceSupport.NewTxManagerFake(),
		},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatRepoMock := tt.chatRepositoryMock(mc)
			chatToUserRepoMock := tt.chatToUserRepositoryMock(mc)
			messageRepoMock := tt.messageRepositoryMock(mc)
			logSrvMock := tt.logServiceMock(mc)
			txMngFake := tt.txManagerFake

			srv := chat.NewService(chatRepoMock, chatToUserRepoMock, messageRepoMock, logSrvMock, txMngFake)

			ar, err := srv.GetMessages(tt.input.ctx, tt.input.chatID, &limit)
			require.Equal(t, tt.expected.err, err)
			require.Equal(t, tt.expected.resp, ar)
		})
	}
}
