package interceptor

import (
	"context"

	"github.com/olezhek28/chat_server/internal/clients/grpc/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type authInterceptor struct {
	authClient auth.Client
}

func NewAuthInterceptor(authClient auth.Client) *authInterceptor {
	return &authInterceptor{
		authClient: authClient,
	}
}

func (i *authInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			ctx = metadata.NewOutgoingContext(ctx, md)
		}

		ok, err = i.authClient.Check(ctx, info.FullMethod)
		if err != nil || !ok {
			return nil, err
		}

		return handler(ctx, req)
	}
}
