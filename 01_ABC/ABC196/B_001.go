package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	pos := strings.Index(s, ".")
	if pos == -1 {
		fmt.Println(s)
	} else {
		fmt.Println(s[:pos])
	}
}
