package tests

import (
	"context"
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

func TestGet(t *testing.T) {
	t.Parallel()

	type chatRepositoryMock func(mc *minimock.Controller) repository.ChatRepository
	type chatToUserRepositoryMock func(mc *minimock.Controller) repository.ChatToUserRepository
	type messageRepositoryMock func(mc *minimock.Controller) repository.MessageRepository
	type logServiceMock func(mc *minimock.Controller) service.LogService

	type input struct {
		ctx    context.Context
		chatID int64
	}

	type expected struct {
		resp *model.Chat
		err  error
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		chatID  = gofakeit.Int64()
		name    = gofakeit.Name()
		userIDs = []int64{gofakeit.Int64()}

		repoErr = fmt.Errorf("repo error")
	)

	cht := &model.Chat{
		ID: chatID,
		Info: model.ChatInfo{
			Name:    name,
			UserIDs: userIDs,
		},
	}

	chatToUser := []model.ChatToUser{
		{
			ID:     gofakeit.Int64(),
			ChatID: chatID,
			UserID: userIDs[0],
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
				resp: cht,
			},
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repoMocks.NewChatRepositoryMock(mc)
				mock.GetMock.Expect(ctx, chatID).Return(cht, nil)
				return mock
			},
			chatToUserRepositoryMock: func(mc *minimock.Controller) repository.ChatToUserRepository {
				mock := repoMocks.NewChatToUserRepositoryMock(mc)

				mock.GetMock.When(ctx, chatID, pagination.Limit{
					Limit:  pagination.DefaultLimit,
					Offset: 0,
				}).Then(chatToUser, nil)

				mock.GetMock.When(ctx, chatID, pagination.Limit{
					Limit:  pagination.DefaultLimit,
					Offset: 1,
				}).Then([]model.ChatToUser{}, nil)
				return mock
			},
			logServiceMock: func(mc *minimock.Controller) service.LogService {
				mock := serviceMocks.NewLogServiceMock(mc)
				mock.CreateMock.Expect(ctx, &model.Log{
					Message: constants.ChatFound,
					Payload: cht,
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
				ctx:    ctx,
				chatID: chatID,
			},
			expected: expected{
				err:  repoErr,
				resp: nil,
			},
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repoMocks.NewChatRepositoryMock(mc)
				mock.GetMock.Expect(ctx, chatID).Return(nil, repoErr)
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

			ar, err := srv.Get(tt.input.ctx, tt.input.chatID)
			require.Equal(t, tt.expected.err, err)
			require.Equal(t, tt.expected.resp, ar)
		})
	}
}
