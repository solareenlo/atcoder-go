package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	for c := 'a'; c <= 'z'; c++ {
		if !strings.Contains(s, string(c)) {
			fmt.Println(string(c))
			return
		}
	}
	fmt.Println("None")
}
