package main

import "handin-3/service"

type server struct {
	clock service.LamportClock

	name string
}
