package main

import (
	"fmt"
	"unicode"
)

func main() {
	var s string
	fmt.Scan(&s)

	i := 0
	for _, r := range s {
		if i%2 != 0 {
			if unicode.IsLower(r) {
				fmt.Println("No")
				return
			}
		} else if unicode.IsUpper(r) {
			fmt.Println("No")
			return
		}
		i++
	}
	fmt.Println("Yes")
}
