package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev, revErr := reverse(input)
	doubleRev, doubleRevErr := reverse(rev)
	fmt.Printf("Original: %q\n", input)
	fmt.Printf("Reverse: %q, Errror: %q\n", rev, revErr)
	fmt.Printf("reverse again: %q, Error: %q\n", doubleRev, doubleRevErr)
}

func reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("The string is not a valid")
	}
	r := []rune(s) //converting the string to runes
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
}

//If the function name didn't start with Cap they wont be exported
