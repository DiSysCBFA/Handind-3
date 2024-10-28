package server

import (
	"log"
	"net"
)

const (
	port = ":4002" // default port
	name = "Chitty-Chat-Server"
)

func SetupServer() {
	log.Println("Setting up server on port:", port, "with the name", name)
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer, err := CreateGrpcServer(name)
	if err != nil {
		log.Fatalf("Failed to create gRPC server: %v", err)
	}

	defer func() {
		if err := lis.Close(); err != nil {
			log.Fatalf("Failed to close listener: %v", err)
		}
	}()

	log.Println("Server is now listening on port:", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
