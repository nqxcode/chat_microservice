package main

import (
	"context"
	"fmt"
	"log"
	"net"

	desc "github.com/nqxcode/chat_microservice/pkg/chat_v1"

	"github.com/brianvoe/gofakeit"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = 50052

type server struct {
	desc.UnimplementedChatV1Server
}

func (s *server) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("User names: %v", req.GetUsernames())

	return &desc.CreateResponse{
		Id: int64(gofakeit.Number(1, 1000)),
	}, nil
}

func (s *server) Delete(_ context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {
	log.Printf("Delete user: %v", req.GetId())

	return nil, nil
}

func (s *server) SendMessage(_ context.Context, req *desc.SendMessageRequest) (*empty.Empty, error) {
	log.Printf("Send message: %v from %v at %v", req.GetMessage(), req.GetFrom(), req.GetTimestamp())

	return nil, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
