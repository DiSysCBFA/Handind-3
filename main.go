package main

import (
	"log"
	"os"

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
		//start new server

		server.SetupServer()

		//use grpc to start new server
	}

	if result == "Start new Client" {
		//start new client

		selectionName := promptui.Prompt{
			Label: "Input desired name, or leave blank for defualt",
		}

		resultName, err := selectionName.Run()

		if err != nil {
			log.Fatalf("Failed to run: %v", err)
		}

		if resultName == "" {
			resultName = "Client"
		}

		/*
			selectionAddress := promptui.Select{
				Label: "Input desired address, or leave blank for defualt",
			}

			_, resultAddress, err := selectionAddress.Run()

			if err != nil {
				log.Fatalf("Failed to run: %v", err)
			}

			// use grpc to start new client
		*/
	}

	if result == "Exit" {

		//exit
		log.Println("Exiting...")
		os.Exit(1)
	}

}
