package main

import (
	"log"
	"time"
)

func main() {
	for {
		log.Printf("Time is %v", time.Now())
		time.Sleep(time.Duration(30 * time.Second))
	}
}
