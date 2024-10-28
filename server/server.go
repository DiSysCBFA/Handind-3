package server

import (
	"log"

	"github.com/DiSysCBFA/Handind-3/service"

	"google.golang.org/grpc"

	tasks "github.com/DiSysCBFA/Handind-3/api"
)

type server struct {
	tasks.UnimplementedChittyChatServer
	clock service.LamportClock

	users map[string]tasks.ChittyChat_BroadcastServer

	name string
}

func (s *server) init() {
	// Init clock on server
	s.clock.AddClock(s.name)
}

func (s *server) incrementClock() {
	s.clock.Tick(s.name)
}

func (s *server) determineNewClock(sender string) {
	s.clock.DetermineNewClock(sender, s.name)
}

func (s *server) getName() string {
	return s.name
}

func (s *server) memberJoin(member string) {
	s.users[member] = nil
	log.Printf("User %s joined the chat", member)
}

func CreateServer(name string) (*server, error) {

	chittyChatServer := server{
		clock: service.LamportClock{},
		name:  name,
	}

	return &chittyChatServer, nil
}

func CreateGrpcServer(name string) (*grpc.Server, error) {
	grpcServer := grpc.NewServer()
	chatServer, err := CreateServer(name)
	chatServer.init()

	if err != nil {
		return nil, err
	}

	tasks.RegisterChittyChatServer(grpcServer, *chatServer)

	log.Printf("Starting gRPC server with name: %s", name)

	chatServer.incrementClock()

	return grpcServer, nil

}
