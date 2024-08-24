package tests

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/nqxcode/chat_microservice/internal/api/chat"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/nqxcode/chat_microservice/internal/service"
	serviceMocks "github.com/nqxcode/chat_microservice/internal/service/mocks"
	desc "github.com/nqxcode/chat_microservice/pkg/chat_v1"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestDelete(t *testing.T) {
	t.Parallel()

	type ChatServiceMockFunc func(mc *minimock.Controller) service.ChatService

	type input struct {
		ctx context.Context
		req *desc.DeleteRequest
	}

	type expected struct {
		resp *empty.Empty
		err  error
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id = gofakeit.Int64()

		serviceErr = fmt.Errorf("service error")

		req = &desc.DeleteRequest{
			Id: id,
		}

		resp = (*empty.Empty)(nil)
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
				mock.DeleteMock.Expect(ctx, req.GetId()).Return(nil)
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
				mock.DeleteMock.Expect(ctx, req.GetId()).Return(serviceErr)
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

			ar, err := api.Delete(tt.input.ctx, tt.input.req)
			require.Equal(t, tt.expected.err, err)
			require.Equal(t, tt.expected.resp, ar)
		})
	}
}
