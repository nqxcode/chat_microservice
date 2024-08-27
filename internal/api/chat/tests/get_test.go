package tests

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/nqxcode/chat_microservice/internal/api/chat"
	"github.com/nqxcode/chat_microservice/internal/converter"
	"github.com/nqxcode/chat_microservice/internal/model"
	"github.com/nqxcode/chat_microservice/internal/service"
	serviceMocks "github.com/nqxcode/chat_microservice/internal/service/mocks"
	desc "github.com/nqxcode/chat_microservice/pkg/chat_v1"
)

func TestGet(t *testing.T) {
	t.Parallel()

	type ChatServiceMockFunc func(mc *minimock.Controller) service.ChatService

	type input struct {
		ctx context.Context
		req *desc.GetRequest
	}

	type expected struct {
		resp *desc.GetResponse
		err  error
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id        = gofakeit.Int64()
		name      = gofakeit.Name()
		userIDs   = []int64{gofakeit.Int64(), gofakeit.Int64(), gofakeit.Int64()}
		createdAt = gofakeit.Date()
		updatedAt = sql.NullTime{Valid: true, Time: gofakeit.Date()}

		serviceErr = fmt.Errorf("service error")

		req = &desc.GetRequest{
			Id: id,
		}

		cht = &model.Chat{
			ID: id,
			Info: model.ChatInfo{
				Name:    name,
				UserIDs: userIDs,
			},
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}

		resp = &desc.GetResponse{
			Chat: converter.ToChatFromService(cht),
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
				mock.GetMock.Expect(ctx, req.GetId()).Return(cht, nil)
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
				mock.GetMock.Expect(ctx, req.GetId()).Return(nil, serviceErr)
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

			ar, err := api.Get(tt.input.ctx, tt.input.req)
			require.Equal(t, tt.expected.err, err)
			require.Equal(t, tt.expected.resp, ar)
		})
	}
}
