package service

import (
	"fmt"
)

type LamportClock map[string]int

func (LcClock *LamportClock) AddClock(username string) {
	if *LcClock == nil {
		*LcClock = make(map[string]int)
	}
	(*LcClock)[username] = 0
}

func (LcClock *LamportClock) Tick(username string) {
	if _, exists := (*LcClock)[username]; !exists {
		(*LcClock)[username] = 0
	}
	(*LcClock)[username]++
}

func (LcClock LamportClock) GetClock(username string) int {
	return LcClock[username]
}

func (LcClock *LamportClock) DetermineNewClock(sender, receiver string) {
	if (*LcClock)[sender] > (*LcClock)[receiver] {
		(*LcClock)[receiver] = (*LcClock)[sender]
	}

	LcClock.Tick(receiver)
}

func (LcClock LamportClock) PrintUserNameClock(username string) {
	fmt.Println(username, ":", LcClock[username])
}

func (LcClock LamportClock) PrintAllClocks() {
	for username, clock := range LcClock {
		fmt.Println(username, ":", clock)
	}
}
