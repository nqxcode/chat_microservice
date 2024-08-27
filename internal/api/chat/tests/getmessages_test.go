package tests

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/nqxcode/chat_microservice/internal/api/chat"
	"github.com/nqxcode/chat_microservice/internal/converter"
	"github.com/nqxcode/chat_microservice/internal/model"
	"github.com/nqxcode/chat_microservice/internal/service"
	serviceMocks "github.com/nqxcode/chat_microservice/internal/service/mocks"
	desc "github.com/nqxcode/chat_microservice/pkg/chat_v1"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGetMessage(t *testing.T) {
	t.Parallel()

	type ChatServiceMockFunc func(mc *minimock.Controller) service.ChatService

	type input struct {
		ctx context.Context
		req *desc.GetMessagesRequest
	}

	type expected struct {
		resp *desc.GetMessagesResponse
		err  error
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		messageID = gofakeit.Int64()
		chatID    = gofakeit.Int64()
		from      = gofakeit.Email()
		message   = gofakeit.Generate("???")
		sentAt    = gofakeit.Date()
		createdAt = gofakeit.Date()
		updatedAt = sql.NullTime{Valid: true, Time: gofakeit.Date()}

		offset = uint64(1)
		limit  = uint64(10)

		serviceErr = fmt.Errorf("service error")

		req = &desc.GetMessagesRequest{
			ChatId: chatID,
			Limit: &desc.Limit{
				Offset: offset,
				Limit:  limit,
			},
		}

		messages = []model.Message{
			{
				ID:        messageID,
				ChatID:    chatID,
				From:      from,
				Message:   message,
				SentAt:    sql.NullTime{Valid: true, Time: sentAt},
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
			},
		}

		resp = &desc.GetMessagesResponse{
			Message: converter.ToMessagesFromService(messages),
		}
	)

	cases := []struct {
		name                string
		input               input
		expected            expected
		ChatServiceMockFunc ChatServiceMockFunc
	}{
		{
			name: "success case",
			input: input{
				ctx: ctx,
				req: req,
			},
			expected: expected{
				resp: resp,
			},
			ChatServiceMockFunc: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				mock.GetMessagesMock.Expect(ctx, req.GetChatId(), converter.ToLimitFromDesc(req.GetLimit())).Return(messages, nil)
				return mock
			},
		},
		{
			name: "service error case",
			input: input{
				ctx: ctx,
				req: req,
			},
			expected: expected{
				err: status.Error(codes.Internal, serviceErr.Error()),
			},
			ChatServiceMockFunc: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				mock.GetMessagesMock.Expect(ctx, req.GetChatId(), converter.ToLimitFromDesc(req.GetLimit())).Return(nil, serviceErr)
				return mock
			},
		},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ChatServiceMock := tt.ChatServiceMockFunc(mc)
			api := chat.NewImplementation(ChatServiceMock)

			ar, err := api.GetMessages(tt.input.ctx, tt.input.req)
			require.Equal(t, tt.expected.err, err)
			require.Equal(t, tt.expected.resp, ar)
		})
	}
}
