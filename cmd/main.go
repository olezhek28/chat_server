package main

import (
	"context"

	"github.com/olezhek28/auth/pkg/access_v1"
	"github.com/olezhek28/chat_server/internal/client/grpc/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	ctx := context.Background()

	opt := grpc.WithInsecure()

	conn, err := grpc.DialContext(ctx, "localhost:50051", opt)
	if err != nil {
		panic(err)
	}

	cl := access_v1.NewAccessV1Client(conn)

	authClient := auth.NewClient(cl)

	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODQ1ODYyMjksInVzZXJuYW1lIjoib2xlZyIsInJvbGUiOiJ1c2VyIn0.ZJnSeuQMWwXYWtDZQDZ-SCIYvdigZANr5mICyP1eN1Y"

	md := metadata.New(map[string]string{"Authorization": "Bearer " + accessToken})
	ctx = metadata.NewOutgoingContext(ctx, md)

	ok, err := authClient.Check(ctx, "bla")
	if err != nil {
		panic(err)
	}

	println(ok)
}
