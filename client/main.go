package client

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	tasks "github.com/DiSysCBFA/Handind-3/api"
	"github.com/DiSysCBFA/Handind-3/service"
)

const (
	defaultAddress = "localhost:4002" // Updated to use port 4002
	name           = "chittyChat"
)

var (
	username string
	address  string
	LcClock  *service.LamportClock // Use a pointer to ensure initialization
	client   tasks.ChittyChatClient
)

func StartClient(NameInput string, addressInput string) {
	username = NameInput

	// Initialize Lamport Clock to avoid nil map issues
	LcClock = service.NewLamportClock()
	LcClock.AddClock(username)

	// Connect to server using grpc.Dial
	conn, err := grpc.Dial(addressInput, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client = tasks.NewChittyChatClient(conn)
	JoinChat(client)
	select {}
}

func JoinChat(client tasks.ChittyChatClient) {
	log.Printf("Joining chat as user: %s on time %d...", username, LcClock.GetClock(username))

	// Create a context with a timeout to avoid indefinite waits
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.Join(ctx, &tasks.Joins{
		Participant: username,
		Timestamp:   int32(LcClock.GetClock(username)),
	})

	if err != nil {
		log.Fatalf("Could not join chat: %v", err)
	}

	LcClock.Tick(username)
	log.Printf("Welcome! You just joined the chat with status: %s at time %d", res.Participant, LcClock.GetClock(username))
}

func LeaveChat() {
	log.Printf("Leaving chat as user: %s on time %d...", username, LcClock.GetClock(username))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := client.Leave(ctx, &tasks.Leaves{
		Participant: username,
	})
	if err != nil {
		log.Fatalf("Could not leave chat: %v", err)
	}
}

func sendMessage() {
	LcClock.Tick(username)
	log.Printf("Sending message as user: %s on time %d...", username, LcClock.GetClock(username))

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := client.Toserver(ctx, &tasks.Message{
		Participant: username,
	})
	if err != nil {
		log.Fatalf("Could not send message: %v", err)
	}
}

func listenerforMessages() {
	/*stream, err := client.Broadcast(context.Background(), &tasks.Message{})
	  if err != nil {
	      log.Fatalf("Error listening for messages: %v", err)
	  }

	  for {
	      msg, err := stream.Recv()
	      if err != nil {
	          log.Fatalf("Error receiving message: %v", err)
	      }
	      fmt.Printf("Received message from %s: %s\n", msg.Participant, msg.Content)
	  }*/
}
