package server

import "handin-3/service"

type server struct {
	clock service.LamportClock

	name string
}

func (s *server) init() {
	// Init clock on server
	s.clock.AddClock(s.name)
}

func (s *server) incrementClock() {
	s.clock.Tick(s.name)
}

func (s *server) determineNewClock(sender string) {
	s.clock.DetermineNewClock(sender, s.name)
}

func (s *server) getName() string {
	return s.name
}
