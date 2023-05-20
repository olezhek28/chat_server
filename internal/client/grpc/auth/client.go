package auth

import (
	"context"

	"github.com/olezhek28/auth/pkg/access_v1"
)

var _ Client = (*client)(nil)

type Client interface {
	Check(ctx context.Context, endpoint string) (bool, error)
}

type client struct {
	accessClient access_v1.AccessV1Client
}

func NewClient(cl access_v1.AccessV1Client) *client {
	return &client{
		accessClient: cl,
	}
}

func (c *client) Check(ctx context.Context, endpoint string) (bool, error) {
	ok, err := c.accessClient.Check(ctx, &access_v1.CheckRequest{
		EndpointAddress: endpoint,
	})
	if err != nil {
		return false, err
	}

	return ok.GetIsAllowed(), nil
}
