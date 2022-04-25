package main

import (
	"fmt"
	"strings"
)

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	if s == t {
		fmt.Println("same")
		return
	}

	s = strings.ToLower(s)
	t = strings.ToLower(t)
	if s == t {
		fmt.Println("case-insensitive")
		return
	}

	fmt.Println("different")
}
