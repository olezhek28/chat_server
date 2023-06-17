package chat_v1

import (
	"context"

	"github.com/google/uuid"
	desc "github.com/olezhek28/chat_server/pkg/chat_v1"
)

func (i *Implementation) CreateChat(ctx context.Context, req *desc.CreateChatRequest) (*desc.CreateChatResponse, error) {
	chatID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	i.channels[chatID.String()] = make(chan *desc.Message, 100)

	return &desc.CreateChatResponse{
		ChatId: chatID.String(),
	}, nil
}

// chat (chat_id)(u1, u2, u3...)
// u1 -> chat stream
// u2 -> chat stream
// u3 -> chat stream
