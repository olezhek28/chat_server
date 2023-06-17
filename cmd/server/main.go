package main

import (
	"context"
	"log"

	"github.com/olezhek28/chat_server/internal/app"
)

func main() {
	ctx := context.Background()

	chatApp, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to initialize app: %s", err.Error())
	}

	log.Println("service starting up")

	if err = chatApp.Run(); err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}

	log.Println("service shutting down")
}
