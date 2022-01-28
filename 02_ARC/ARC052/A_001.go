package main

import (
	"fmt"
	"unicode"
)

func main() {
	var s string
	fmt.Scan(&s)

	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			fmt.Print(string(s[i]))
		}
	}
	fmt.Println()
}
