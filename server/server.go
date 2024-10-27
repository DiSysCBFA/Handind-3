package main

import "handin-3/service"

const (
	name = "Chitty-Chat-Server"
)

type server struct {
	clock service.LamportClock

	name string
}

func (s *server) incrementClock() {
	s.clock.Tick(s.name)
}

func (s *server) getName() string {
	return s.name
}
