package client

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/DiSysCBFA/Handind-3/api"
	chat "github.com/DiSysCBFA/Handind-3/api"
	"google.golang.org/grpc"
)

// Client represents a chat client with a gRPC connection
type Client struct {
	api.ChittyChatClient
	conn      *grpc.ClientConn
	port      string
	name      string
	timestamp int
}

// NewClient creates a new client instance and initializes the gRPC connection
func NewClient(name, port string) *Client {
	address := "localhost:" + port
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}

	client := chat.NewChittyChatClient(conn)
	return &Client{
		ChittyChatClient: client,
		conn:             conn,
		port:             port,
		name:             name,
		timestamp:        0,
	}
}

// Join connects to the server and starts listening for messages on the Join stream
func (c *Client) Join() {
	// Start the Join stream to listen for incoming messages
	stream, err := c.ChittyChatClient.Join(context.Background(), &chat.Empty{})
	if err != nil {
		log.Fatalf("Failed to join chat: %v", err)
	}

	// Listen for messages in a separate goroutine
	go func() {
		for {
			in, err := stream.Recv()
			if err != nil {
				log.Fatalf("Failed to receive message: %v", err)
			}
			log.Printf("[%s]: %s", in.Participant, in.Content)
		}
	}()

	// Start broadcasting messages to the server
	c.BroadcastMessages()
}

// BroadcastMessages prompts the user to send messages
func (c *Client) BroadcastMessages() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter message: ")
		content, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Failed to read message: %v", err)
		}

		// Trim newline characters from the message
		content = strings.TrimSpace(content)

		// Send the message to the server using the Broadcast method
		_, err = c.ChittyChatClient.Broadcast(context.Background(), &chat.Message{
			Participant: c.name,
			Content:     content,
			Timestamp:   time.Now().Unix(),
		})
		if err != nil {
			log.Fatalf("Failed to send message: %v", err)
		}
	}
}

// Close closes the gRPC connection
func (c *Client) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}
