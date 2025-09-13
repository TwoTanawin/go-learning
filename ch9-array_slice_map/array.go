package main

import (
	"fmt"
)

func SampleArray() {
	var f [3]string
	f[0] = "a"
	f[1] = "b"
	f[2] = "c"

	fmt.Println("f : ", f)

	i := 0

	for i < len(f) {
		defer fmt.Printf("index %d => %s : \n", i, f[i])
		i++
	}
}

func SampleArray2() {
	f := []string{"a", "b", "c"}

	f = append(f, "d")

	i := 0

	for i < len(f) {
		defer fmt.Printf("index %d => %s : \n", i, f[i])
		i++
	}

	fmt.Println("Size : ", len(f))
	fmt.Println("Cap : ", cap(f))
}
