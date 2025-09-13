package greeting

import (
	"fmt"
)

// Pub
func SayMeeting() {
	fmt.Println("Hello, from Meeting Method")
}

// Private
func sayHi() {
	fmt.Println("Hello, from sayHi Method")
}
