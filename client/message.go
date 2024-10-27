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

func (m *Message) setUsername(username string) {
	m.username = username
}

func (m *Message) setContent(content string) {
	m.content = content

}

func (m *Message) setTimestamp() {
	m.timestamp = time.Now()
}
