package service

type LamportClock struct {
	clock map[string]int
}

// NewLamportClock initializes the clock map to avoid nil map errors
func NewLamportClock() *LamportClock {
	return &LamportClock{
		clock: make(map[string]int),
	}
}

// AddClock initializes a clock entry for a new user
func (lc *LamportClock) AddClock(username string) {
	if lc.clock == nil {
		lc.clock = make(map[string]int)
	}
	lc.clock[username] = 0
}

// Tick increments the clock for the specified user
func (lc *LamportClock) Tick(username string) {
	if lc.clock == nil {
		lc.clock = make(map[string]int)
	}
	lc.clock[username]++
}

// GetClock retrieves the current clock value for a specified user
func (lc *LamportClock) GetClock(username string) int {
	if lc.clock == nil {
		lc.clock = make(map[string]int)
	}
	return lc.clock[username]
}

func (lc *LamportClock) DetermineNewClock(sender string, receiver string) {
	if lc.clock == nil {
		lc.clock = make(map[string]int)
	}
	if lc.clock[sender] >= lc.clock[receiver] {
		lc.clock[receiver] = lc.clock[sender] + 1
	}
}
