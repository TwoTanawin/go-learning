package greeting

import (
	"fmt"

	personal "github.com/TwoTanawin/samplego/greeting/internal"
)

func SayGreeting() {
	personal.PrintPersonal()
	fmt.Println("Hello, from Greeting")
	sayHi()
}
