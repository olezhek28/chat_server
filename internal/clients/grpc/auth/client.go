package auth

import (
	"context"
	"fmt"

	accessV1 "github.com/olezhek28/auth/pkg/access_v1"
)

var _ Client = (*client)(nil)

type Client interface {
	Check(ctx context.Context, endpoint string) (bool, error)
}

type client struct {
	accessClient accessV1.AccessV1Client
}

func NewClient(cl accessV1.AccessV1Client) *client {
	return &client{
		accessClient: cl,
	}
}

func (c *client) Check(ctx context.Context, endpoint string) (bool, error) {
	if _, err := c.accessClient.Check(ctx, &accessV1.CheckRequest{
		EndpointAddress: endpoint,
	}); err != nil {
		return false, fmt.Errorf("accessClient.Check: %w", err)
	}

	return true, nil
}
