package main

import (
	"log"
	"net"
	"os"

	"github.com/DiSysCBFA/Handind-3/client"
	"github.com/DiSysCBFA/Handind-3/server"
	"github.com/manifoldco/promptui"
)

func main() {
	selection := promptui.Select{
		Label: "Select an option",
		Items: []string{"Start Server", "Start new Client", "Exit"},
	}

	_, result, err := selection.Run()
	if err != nil {
		log.Fatalf("Failed to run: %v", err)
	}

	if result == "Start Server" {
		SetupServer() // Start the server
	}

	if result == "Start new Client" {
		// Prompt for client name
		selectionName := promptui.Prompt{
			Label: "Enter desired name",
		}
		username, err := selectionName.Run()
		if err != nil {
			log.Fatalf("Failed to run: %v", err)
		}

		// Prompt for server address
		selectionAddress := promptui.Prompt{
			Label: "Enter server address",
		}
		address, err := selectionAddress.Run()
		if err != nil || address == "" {
			log.Fatalf("Address must be provided")
		}

		// Initialize and join client
		newClient := client.NewClient(username, address)
		defer newClient.Close() // Ensure the connection closes on exit

		newClient.Join() // Start the client join process and begin messaging

		select {} // Keep the program running
	}

	if result == "Exit" {
		log.Println("Exiting...")
		os.Exit(1)
	}
}

const (
	port = ":4002"
	name = "ChittyChat"
)

func SetupServer() {
	log.Println("Setting up server on port:", port, "with the name", name)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer, err := server.CreateGrpcServer(name)
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
