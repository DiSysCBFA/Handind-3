package server

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	chat "github.com/DiSysCBFA/Handind-3/api"
	"google.golang.org/grpc"
)

// Server represents the gRPC server with connected clients
type Server struct {
	chat.UnimplementedChittyChatServer                                       // Embed for forward compatibility
	mu                                 sync.Mutex                            // Mutex to protect access to clients
	clients                            map[string]chat.ChittyChat_JoinServer // Store clients for broadcasting
	clock                              int64
}

// CreateGrpcServer initializes a new gRPC server
func CreateGrpcServer(name string) (*grpc.Server, error) {
	grpcServer := grpc.NewServer()
	server := &Server{
		clients: make(map[string]chat.ChittyChat_JoinServer),
		clock:   0,
	}
	chat.RegisterChittyChatServer(grpcServer, server)
	log.Printf("gRPC server '%s' created", name)
	return grpcServer, nil
}

// Broadcast sends a message to all connected clients
func (s *Server) Broadcast(ctx context.Context, msg *chat.Message) (*chat.Empty, error) {
	s.clock++
	log.Printf("Broadcasting message from [%s]: %s ... Timestamp: %d", msg.Participant, msg.Content, s.clock)

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
// Join registers a client for receiving messages
func (s *Server) Join(_ *chat.Empty, stream chat.ChittyChat_JoinServer) error {
	s.clock++
	// Use a unique client ID (you can generate this as needed)
	clientID := fmt.Sprintf("client-%d", time.Now().UnixNano()) // Unique ID based on time

	log.Printf("Client %s joined the chat", clientID)

	// Register the client
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
