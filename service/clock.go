package main

// Lapart clock type with a map of username (string) to each individual clock (int)
type LamportClock map[string]int

// Initialize a new LamportClock on username
func (LcClock LamportClock) AddClock(username string) {
	LcClock[username] = 0
}

// increment the clock of the given username
func (LcClock LamportClock) Tick(username string) {

	// TODO: implement check for if the username exists in the map. otherwise create it.
	LcClock[username]++
}
