package interceptor

import (
	"context"

	"github.com/olezhek28/chat_server/internal/client/grpc/auth"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type AuthInterceptor struct {
	authClient auth.Client
}

func NewAuthInterceptor(authClient auth.Client) *AuthInterceptor {
	return &AuthInterceptor{authClient: authClient}
}

func (i *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			ctx = metadata.NewOutgoingContext(ctx, md)
		}

		ok, err = i.authClient.Check(ctx, info.FullMethod)
		if err != nil || !ok {
			return nil, errors.Wrap(err, "failed to check access")
		}

		return handler(ctx, req)
	}
}
