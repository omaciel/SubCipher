package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func userPrompt() string {
	fmt.Printf("Message to encrypt:")
	input, _ := reader.ReadString('\n')

	return strings.TrimSpace(input)
}

func main() {
	text := userPrompt()

	for _, letter := range text {
		switch newLetter := string(letter); newLetter {
		case "P":
			fmt.Print("Z")
		default:
			fmt.Print(newLetter)
		}

	}
}
