package server

import (
	"context"
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

func (s *server) memberJoin(member string) error {
	s.users[member] = nil
	log.Printf("User %s joined the chat", member)
	return nil
}

func (s *server) GetClock() int32 {
	return int32(s.clock.GetClock(s.getName()))
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

func (s server) Join(context context.Context, req *tasks.JoinRequest) error {
	log.Printf("a client wants to join the chat")

	// add the client to the broadcast
	err := s.memberJoin(req.GetUsername())
	if err != nil {
		return err
	}
	// if the client is not in the clients map, add it to the clients map
	// and return a JoinResponse with a status of OK
	log.Printf("[%s: %d] Received a JOIN req from node %s", s.getName(), s.GetClock())
	s.incrementClock()
	return nil
}
