package main

import (
	"fmt"
)

const KEY rune = 'A'

func xorString(text string, key rune) string {
	result := ""
	for _, char := range text {
		result += string(char ^ key)

	}
	return result
}

func main() {
	originalText := "HELLO"

	// Encrypt
	ciphertext := xorString(originalText, KEY)
	fmt.Printf("Ciphertext: %s\n", ciphertext)

	// Decrypt (XORing with the same key again reverses the operation)
	decryptedText := xorString(ciphertext, KEY)
	fmt.Printf("Decrypted Text: %s\n", decryptedText)
}
