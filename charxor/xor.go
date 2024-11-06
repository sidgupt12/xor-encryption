package main

import (
	"fmt"
)

const KEY rune = 'X'
const CLEARTEXT rune = 'A'

func printchar(c rune, text string) {
	fmt.Printf("%s = '%c' (0x%x)\n", text, c, c)
}

func main() {

	var cleartext, key, ciphertext, newcleartext rune

	cleartext = CLEARTEXT
	printchar(cleartext, "cleartext")

	key = KEY
	printchar(key, "key")

	// ^ character is often called caret character
	ciphertext = cleartext ^ key

	printchar(ciphertext, "ciphertext")

	newcleartext = ciphertext ^ key
	printchar(newcleartext, "newcleartext")

}
