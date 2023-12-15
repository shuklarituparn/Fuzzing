package main

import "fmt"

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev := reverse(input)
	doublerev := reverse(rev)
	fmt.Printf("Original: %q\n", input)
	fmt.Printf("Reverse: %q\n", rev)
	fmt.Printf("reverse again: %q\n", doublerev)
}

func reverse(s string) string {
	r := []rune(s) //converting the string to runes
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
