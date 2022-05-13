package main

import (
	"fmt"
	"unicode"
)

func main() {
	var s string
	fmt.Scan(&s)

	u, l, f := 0, 0, 1
	v := [256]int{}
	for _, c := range s {
		if unicode.IsUpper(c) {
			u = 1
		}
		if unicode.IsLower(c) {
			l = 1
		}
		if v[c] != 0 {
			f = 0
		}
		v[c] = 1
	}

	if u&l&f != 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
