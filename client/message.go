package main

import (
	"errors"
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

func (m *Message) setContent(content string) error {

	if len(content) > 128 {
		return errors.New("Message cannot exceed 128 characters")
	}

	m.content = content

}

func (m *Message) setTimestamp() {
	m.timestamp = time.Now()
}
