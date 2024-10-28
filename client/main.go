package main

import (
	"fmt"

	"github.com/DiSysCBFA/Handind-3/service"
)

var (
	username string
	LcClock  service.LamportClock
)

func main() {

	// Init username
	fmt.Println("Enter username:")
	fmt.Scan(&username)
	fmt.Println("Hello", username)

	// Init clock on username
	LcClock.AddClock(username)

	LcClock.PrintUserNameClock(username)

}

func StartClient(NameInput string, adressInput string) {

	
}

func JoinChat(){

}
func LeaveChat(){

}

func sendMessage(){

}

func ListnerforMessages(){}