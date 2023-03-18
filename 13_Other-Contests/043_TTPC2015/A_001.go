package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	if s[2] == 'M' {
		fmt.Printf("Master %c%c\n", s[0], s[1])
	}
	if s[2] == 'D' {
		fmt.Printf("Doctor %c%c\n", s[0], s[1])
	}
	if s[2] == 'B' {
		fmt.Printf("Bachelor %c%c\n", s[0], s[1])
	}
}
