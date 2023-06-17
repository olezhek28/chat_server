package app

import (
	"context"
	"log"

	accessV1 "github.com/olezhek28/auth/pkg/access_v1"
	chatV1 "github.com/olezhek28/chat_server/internal/api/chat_v1"
	authClient "github.com/olezhek28/chat_server/internal/clients/grpc/auth"
	"github.com/olezhek28/chat_server/internal/closer"
	chatService "github.com/olezhek28/chat_server/internal/service/chat"
	"github.com/satanaroom/auth/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type serviceProvider struct {
	authClient  authClient.Client
	chatService chatService.Service

	tlsCredentials credentials.TransportCredentials

	chatImpl *chatV1.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) ChatService(ctx context.Context) chatService.Service {
	if s.chatService == nil {
		s.chatService = chatService.NewService(s.AuthClient(ctx))
	}

	return s.chatService
}

func (s *serviceProvider) AuthClient(ctx context.Context) authClient.Client {
	if s.authClient == nil {
		creds, err := credentials.NewClientTLSFromFile("service.pem", "")
		if err != nil {
			log.Fatalf("could not process the credentials: %v", err)
		}

		conn, err := grpc.DialContext(ctx, "localhost:50051", grpc.WithTransportCredentials(creds))
		if err != nil {
			logger.Fatalf("failed to connect %s: %s", "localhost:50051", err.Error())
		}
		closer.Add(conn.Close)

		client := accessV1.NewAccessV1Client(conn)
		s.authClient = authClient.NewClient(client)
	}

	return s.authClient
}

func (s *serviceProvider) ChatImpl(ctx context.Context) *chatV1.Implementation {
	if s.chatImpl == nil {
		s.chatImpl = chatV1.NewImplementation(s.ChatService(ctx))
	}

	return s.chatImpl
}
