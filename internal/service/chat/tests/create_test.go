package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/nqxcode/platform_common/client/db"
	"github.com/stretchr/testify/require"

	"github.com/nqxcode/chat_microservice/internal/model"
	"github.com/nqxcode/chat_microservice/internal/repository"
	repoMocks "github.com/nqxcode/chat_microservice/internal/repository/mocks"
	"github.com/nqxcode/chat_microservice/internal/service"
	"github.com/nqxcode/chat_microservice/internal/service/audit_log/constants"
	"github.com/nqxcode/chat_microservice/internal/service/chat"
	serviceSupport "github.com/nqxcode/chat_microservice/internal/service/chat/tests/support"
	serviceMocks "github.com/nqxcode/chat_microservice/internal/service/mocks"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	type chatRepositoryMock func(mc *minimock.Controller) repository.ChatRepository
	type chatToUserRepositoryMock func(mc *minimock.Controller) repository.ChatToUserRepository
	type messageRepositoryMock func(mc *minimock.Controller) repository.MessageRepository
	type logServiceMock func(mc *minimock.Controller) service.AuditLogService

	type input struct {
		ctx  context.Context
		chat *model.Chat
	}

	type expected struct {
		resp any
		err  error
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id      = gofakeit.Int64()
		name    = gofakeit.Name()
		userIDs = []int64{gofakeit.Int64()}

		repoErr = fmt.Errorf("repo error")
	)

	cht := &model.Chat{
		ID: id,
		Info: model.ChatInfo{
			Name:    name,
			UserIDs: userIDs,
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
				ctx:  ctx,
				chat: cht,
			},
			expected: expected{
				err:  nil,
				resp: id,
			},
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repoMocks.NewChatRepositoryMock(mc)
				mock.CreateMock.Expect(ctx, &cht.Info).Return(id, nil)
				return mock
			},
			chatToUserRepositoryMock: func(mc *minimock.Controller) repository.ChatToUserRepository {
				mock := repoMocks.NewChatToUserRepositoryMock(mc)
				mock.CreateMock.Expect(ctx, &model.ChatToUser{
					ChatID: cht.ID,
					UserID: userIDs[0],
				}).Return(gofakeit.Int64(), nil)
				return mock
			},
			logServiceMock: func(mc *minimock.Controller) service.AuditLogService {
				mock := serviceMocks.NewAuditLogServiceMock(mc)
				mock.CreateMock.Expect(ctx, &model.Log{
					Message: constants.ChatCreated,
					Payload: model.Chat{ID: id, Info: cht.Info},
				}).Return(nil)
				return mock
			},
			messageRepositoryMock: func(mc *minimock.Controller) repository.MessageRepository {
				return repoMocks.NewMessageRepositoryMock(mc)
			},
			txManagerFake: serviceSupport.NewTxManagerFake(),
		},
		{
			name: "service error case",
			input: input{
				ctx:  ctx,
				chat: cht,
			},
			expected: expected{
				err:  repoErr,
				resp: int64(0),
			},
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repoMocks.NewChatRepositoryMock(mc)
				mock.CreateMock.Expect(ctx, &cht.Info).Return(0, repoErr)
				return mock
			},
			chatToUserRepositoryMock: func(mc *minimock.Controller) repository.ChatToUserRepository {
				mock := repoMocks.NewChatToUserRepositoryMock(mc)
				return mock
			},
			logServiceMock: func(mc *minimock.Controller) service.AuditLogService {
				mock := serviceMocks.NewAuditLogServiceMock(mc)
				return mock
			},
			messageRepositoryMock: func(mc *minimock.Controller) repository.MessageRepository {
				return repoMocks.NewMessageRepositoryMock(mc)
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

			ar, err := srv.Create(tt.input.ctx, &tt.input.chat.Info)
			require.Equal(t, tt.expected.err, err)
			require.Equal(t, tt.expected.resp, ar)
		})
	}
}
