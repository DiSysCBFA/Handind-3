package main

import (
	"log"
	"os"

	"github.com/DiSysCBFA/Handind-3/server"

	"github.com/DiSysCBFA/Handind-3/client"

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
		//start new server

		server.SetupServer()

		//use grpc to start new server
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

		// Prompt for server address without a default
		selectionAddress := promptui.Prompt{
			Label: "Enter server address",
		}
		address, err := selectionAddress.Run()
		if err != nil || address == "" {
			log.Fatalf("Address must be provided")
		}
		address = "localhost:" + address;
		client.StartClient(username, address)
		select {
		// Listen for messages
		//listenerforMessages()
		}
	}

	if result == "Exit" {

		//exit
		log.Println("Exiting...")
		os.Exit(1)
	}
}
