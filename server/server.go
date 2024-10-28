package server

import (
	"context"
	"log"

	"github.com/DiSysCBFA/Handind-3/service"
	"google.golang.org/grpc"

	tasks "github.com/DiSysCBFA/Handind-3/api"
)

type server struct {
	tasks.UnimplementedChittyChatServer // Embedding the unimplemented server for forward compatibility
	clock                               service.LamportClock
	users                               map[string]tasks.ChittyChat_BroadcastServer
	name                                string
}

// Initialize the Lamport clock for the server instance
func (s *server) init() {
	s.clock.AddClock(s.name)
}

// Method to increment the clock for the server
func (s *server) incrementClock() {
	s.clock.Tick(s.name)
}

// Method to determine the new clock based on a sender's clock
func (s *server) determineNewClock(sender string) {
	s.clock.DetermineNewClock(sender, s.name)
}

// Retrieve the server's name
func (s *server) getName() string {
	return s.name
}

// Method for handling a new member joining the chat
func (s *server) memberJoin(member string) error {
	s.users[member] = nil
	log.Printf("User %s joined the chat", member)
	return nil
}

// Retrieve the current clock for the server as an int32
func (s *server) GetClock() int32 {
	return int32(s.clock.GetClock(s.getName()))
}

// Create and initialize a new server instance
func CreateServer(name string) (*server, error) {
	chittyChatServer := &server{
		clock: service.LamportClock{},
		name:  name,
		users: make(map[string]tasks.ChittyChat_BroadcastServer),
	}
	chittyChatServer.init()
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
	tasks.RegisterChittyChatServer(grpcServer, chatServer)

	log.Printf("Starting gRPC server with name: %s", name)

	chatServer.incrementClock()

	return grpcServer, nil
}

// Implementation of the Join method as defined in the ChittyChatServer interface
func (s *server) Join(ctx context.Context, req *tasks.Joins) (*tasks.JoinMessage, error) {
	log.Printf("A client wants to join the chat")

	// Add the client to the broadcast
	err := s.memberJoin(req.GetParticipant())
	if err != nil {
		return nil, err
	}

	// Log the join request
	log.Printf("[%s: %d] Received a JOIN request from: %s", s.getName(), s.GetClock(), req.GetParticipant())

	// Increment the clock after processing the request
	s.incrementClock()

	// Construct and return the JoinMessage with the current Lamport clock
	return &tasks.JoinMessage{
		Participant: req.GetParticipant(),
		Timestamp:   s.GetClock(),
	}, nil
}
