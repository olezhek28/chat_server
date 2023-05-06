package auth

import (
	"context"
	"time"

	"github.com/olezhek28/auth/pkg/note_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ Client = (*client)(nil)

type Info struct {
	Title     string
	Content   string
	CreatedAt time.Time
}

type Client interface {
	Create(ctx context.Context, info Info) (int64, error)
}

type client struct {
	noteClient note_v1.NoteV1Client
}

func NewClient(cl note_v1.NoteV1Client) *client {
	return &client{
		noteClient: cl,
	}
}

func (c *client) Create(ctx context.Context, info Info) (int64, error) {
	res, err := c.noteClient.Create(ctx, &note_v1.CreateRequest{
		Info: &note_v1.NoteInfo{
			Title:     info.Title,
			Content:   info.Content,
			CreatedAt: timestamppb.New(info.CreatedAt),
		},
	})
	if err != nil {
		return 0, err
	}

	return res.GetId(), nil
}
