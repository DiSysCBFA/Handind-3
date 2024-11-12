package server

import (
	"context"
	"log"

	chat "github.com/DiSysCBFA/Handind-3/api"
	"github.com/DiSysCBFA/Handind-3/service"
	"google.golang.org/grpc"
)

type server struct {
	chat.UnimplementedChittyChatServer

	name  string
	clock service.LamportClock
}

// Create and initialize a new server instance
func CreateServer(name string) (*server, error) {
	chittyChatServer := &server{
		clock: service.LamportClock{},
		name:  name,
	}
	chittyChatServer.clock.AddClock(name)
	return chittyChatServer, nil
}

// Create a new gRPC server instance
func CreateGrpcServer(name string) (*grpc.Server, error) {
	grpcServer := grpc.NewServer()
	chatServer, err := CreateServer(name)
	if err != nil {
		return nil, err
	}

	// Register the ChittyChat server
	chat.RegisterChittyChatServer(grpcServer, chatServer)

	log.Printf("Starting gRPC server with name: %s", name)

	chatServer.clock.Tick(chatServer.name)

	return grpcServer, nil
}

func Broadcast(context.Context, *chat.Message) (*chat.Empty, error) {
	return nil, nil
}

func Join(*chat.Empty, grpc.ServerStreamingServer[chat.Message]) error {
	return nil
}
