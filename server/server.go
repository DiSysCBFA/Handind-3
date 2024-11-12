package server

import (
	"context"
	"log"
	"sync"

	chat "github.com/DiSysCBFA/Handind-3/api"
	"google.golang.org/grpc"
)

// Server represents the gRPC server with connected clients
type Server struct {
	chat.UnimplementedChittyChatServer                                       // Embed for forward compatibility
	mu                                 sync.Mutex                            // Mutex to protect access to clients
	clients                            map[string]chat.ChittyChat_JoinServer // Store clients for broadcasting
}

// CreateGrpcServer initializes a new gRPC server
func CreateGrpcServer(name string) (*grpc.Server, error) {
	grpcServer := grpc.NewServer()
	server := &Server{
		clients: make(map[string]chat.ChittyChat_JoinServer),
	}
	chat.RegisterChittyChatServer(grpcServer, server)
	log.Printf("gRPC server '%s' created", name)
	return grpcServer, nil
}

// Broadcast sends a message to all connected clients
func (s *Server) Broadcast(ctx context.Context, msg *chat.Message) (*chat.Empty, error) {
	log.Printf("Broadcasting message from [%s]: %s", msg.Participant, msg.Content)

	s.mu.Lock()
	defer s.mu.Unlock()

	for _, client := range s.clients {
		if err := client.Send(msg); err != nil {
			log.Printf("Failed to send message to client: %v", err)
		}
	}

	return &chat.Empty{}, nil
}

// Join registers a client for receiving messages
func (s *Server) Join(_ *chat.Empty, stream chat.ChittyChat_JoinServer) error {
	// Generate a unique client ID for demonstration
	clientID := stream.Context().Value("clientID").(string)
	log.Printf("Client %s joined the chat", clientID)

	s.mu.Lock()
	s.clients[clientID] = stream
	s.mu.Unlock()

	// Remove client on disconnect
	defer func() {
		s.mu.Lock()
		delete(s.clients, clientID)
		s.mu.Unlock()
		log.Printf("Client %s left the chat", clientID)
	}()

	// Keep the stream open so the client can receive messages
	for {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		}
	}
}
