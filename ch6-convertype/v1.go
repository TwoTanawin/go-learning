package main

import (
	"fmt"
	"strconv"
)

func V1Conv() {
	i := 42
	var f float64 = float64(i)
	fmt.Printf("i : %d f : %f\n", i, f)
}

func V3Conv() {
	var i int = 42
	var s string = strconv.Itoa(i)
	fmt.Printf("s : %s\n", s)
}

func V4Conv() {
	var s string = "xyz"
	var i int
	var err error

	i, err = strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error :", err)
	}
	fmt.Println("i : ", i)
}
