package main

import (
	"context"
	"github.com/nqxcode/chat_microservice/internal/app"
	desc "github.com/nqxcode/chat_microservice/pkg/chat_v1"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

type server struct {
	desc.UnimplementedChatV1Server
	pool *pgxpool.Pool
}

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
