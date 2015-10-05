package main

import (
	"time"
)

type Message struct {
	Author    string    `json:"author"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

func (self *Message) String() string {
	return self.Author + " says " + self.Body
}
