package main

import "fmt"

func add(a int, b int) int {
	sum := a + b
	return sum
}

func main() {
	var c int = add(1, 2)
	fmt.Println(c)

	Func1()
	Func2(1, 2)
	fmt.Println(Func3(1, 2))
}
