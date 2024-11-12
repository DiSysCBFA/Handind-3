package server

import (
	"context"

	chat "github.com/DiSysCBFA/Handind-3/api"
	"google.golang.org/grpc"
)

type Server struct {
	chat.UnimplementedChittyChatServer
}

func Broadcast(context.Context, *chat.Message) (*chat.Empty, error) {
	return nil, nil
}

func Join(*chat.Empty, grpc.ServerStreamingServer[chat.Message]) error {
	return nil
}
