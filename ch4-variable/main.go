package main

import (
	"fmt"
)

var xmlGolbal string = "Outside"

func main() {
	xml := "inside main"
	fmt.Println(xml)
	fmt.Println(xmlGolbal)

	var (
		x float32 = 1.1
		y int     = 10_000
		z string  = "John"
	)

	fmt.Println(x, y, z)

	const b string = "Hello"
	const pi float32 = 3.14
	fmt.Println(b, pi)
}
