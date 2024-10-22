package main

import (
	"fmt"
)

var (
	username string
)

func main() {

	fmt.Println("Enter username:")
	fmt.Scan(&username)
	fmt.Println("Hello", username)
}
