package main

import (
	"fmt"
	"time"
)

type Message struct {
	username  string
	content   string
	timestamp time.Time
}

func (m *Message) display() {
	fmt.Println(m.username, "@ ", m.timestamp, "says: ", m.content)
}
