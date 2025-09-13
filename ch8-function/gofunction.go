package main

import (
	"fmt"
)

func Func1() {
	fmt.Println("Good Morning")
}

func Func2(a int, b int) {
	div := float32(a) / float32(b)
	fmt.Println(div)
}

func Func3(a int, b int) (int, int) {
	sum := a + b
	sub := a - b
	return sum, sub
}
