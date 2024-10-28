package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverse(str string) string {
	name := []rune(str)
	for i, j := 0, len(name)-1; i < j; i, j = i+1, j-1 {
		name[i], name[j] = name[j], name[i]
	}
	return string(name)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		// Take input from the user
		fmt.Println("Please enter your name:")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Split the input by spaces
		parts := strings.Split(input, " ")

		// Check if there are exactly three names
		if len(parts) >= 3 {
			// Reverse and print each name on a new line
			fmt.Println("Reversed names:")
			for _, part := range parts {
				fmt.Println(reverse(part))
			}
			break
		} else {
			fmt.Println("Error: Please enter exactly three names.")
		}
	}
}
