package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var reader = bufio.NewReader(os.Stdin)

func userPrompt() string {
	fmt.Printf("Message to encrypt: ")
	input, _ := reader.ReadString('\n')

	return strings.TrimSpace(input)
}

func cipher(text string, substitutionMap map[string]string) string {
	message := ""

	for _, letter := range text {
		if unicode.IsLetter(letter) {
			if unicode.IsLower(letter) {
				replacement, ok := substitutionMap[string(letter)]
				if ok {
					message = message + replacement
				} else {
					message = message + string(letter)
				}
			} else {
				replacement, ok := substitutionMap[strings.ToLower(string(letter))]
				if ok {
					message = message + strings.ToUpper(replacement)
				} else {
					message = message + string(letter)
				}
			}
		} else {
			message = message + string(letter)
		}
	}
	return message
}
func main() {
	text := userPrompt()
	encryptedMessage := cipher(text, zenit_polar_substitutions())
	fmt.Println("Encrypted message:", encryptedMessage)
	//for _, letter := range text {
	//	switch newLetter := string(letter); newLetter {
	//	case "P":
	//		fmt.Print("Z")
	//	default:
	//		fmt.Print(newLetter)
	//	}
	//
	//}
}
