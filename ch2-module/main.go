package main

import (
	"fmt"

	"github.com/TwoTanawin/samplego/greeting"
	"github.com/google/uuid"
)

func main() {
	fmt.Printf("Welcome Pattaya\n")

	var id string = uuid.New().String()
	fmt.Printf("UUID: %s\n", id)

	greeting.SayGreeting()
	greeting.SayMeeting()
}
