package main

import (
	"fmt"
)

func IfCondition() {
	num := 15
	if num < 15 {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}

func SwitchCase() {
	day := 5

	switch day {
	case 1:
		fmt.Println("Monday")
	case 5:
		fmt.Println("Friday")
	default:
		fmt.Println("Error")
	}
}
