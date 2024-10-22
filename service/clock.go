package main

// Lapart clock type with a map of username (string) to each individual clock (int)
type LamportClock map[string]int

// Initialize a new LamportClock on username
func (LcClock LamportClock) AddClock(username string) {
	LcClock[username] = 0
}

// Increment the clock of the given username
func (LcClock LamportClock) Tick(username string) {

	// TODO: implement check for if the username exists in the map. otherwise create it.
	LcClock[username]++
}

// Get the clock of the given username
func (LcClock LamportClock) GetClock(username string) int {

	return LcClock[username]
}

// Takes 2 usernames and increments the clock of the receiver to the max of the two clocks + 1
func (LcClock LamportClock) DetermineNewClock(sender, reciever string) {
	if LcClock[sender] > LcClock[reciever] {
		LcClock[reciever] = LcClock[sender]
	}

	LcClock.Tick(reciever)
}
