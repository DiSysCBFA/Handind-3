package main

import (
	"fmt"

	"../service"
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

}
