package main

import (
	"log"

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
	}

	if result == "Start new Client" {
		//start new client
	}

	if result == "Exit" {
		//exit
	}

}
