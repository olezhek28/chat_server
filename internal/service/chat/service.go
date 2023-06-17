package chat

import (
	"context"

	authClient "github.com/olezhek28/chat_server/internal/clients/grpc/auth"
	"github.com/olezhek28/chat_server/internal/model"
)

var _ Service = (*service)(nil)

type Service interface {
	CreateChat(ctx context.Context, usernames *model.CreateChat) (int64, error)
}

type service struct {
	authClient authClient.Client
}

func NewService(authClient authClient.Client) *service {
	return &service{
		authClient: authClient,
	}
}
