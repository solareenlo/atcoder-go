package main

import (
	"fmt"
	"strings"
)

func main() {
	var a, b int
	var s string
	fmt.Scan(&a, &b, &s)

	t := strings.Count(s, "-")
	if t == 1 && s[a] == '-' {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
