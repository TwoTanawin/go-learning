package main

import (
	"fmt"
)

func Forloop() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func Forloop2() {
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}
}

func Forloop3() {
	i := 1

	for {
		if i == 1024 {
			break
		}

		fmt.Println(i)
		i++
	}

}
