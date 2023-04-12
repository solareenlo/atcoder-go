package main

import (
	"fmt"
	"strings"
)

func main() {
	var t string
	fmt.Scan(&t)
	s := strings.Split(t, "")
	for i := len(s) - 1; 0 < i; i-- {
		if s[i] == "A" && s[i-1] == "W" {
			s[i] = "C"
			s[i-1] = "A"
		}
	}
	fmt.Println(strings.Join(s, ""))
}
