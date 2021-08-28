package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	s = strings.ToLower(s)
	s = strings.Title(s)
	fmt.Println(s)
}
