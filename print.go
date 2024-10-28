package main

import (
	"fmt"
)

func main() {
	// Take input from user
	fmt.Println("Please enter your number.")
	var number int
	fmt.Scanln(&number)

	if number == 42 {
		fmt.Printf("Hello Universe")
	} else {
		fmt.Printf("%d", number)
	}
}
