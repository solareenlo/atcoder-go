package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	s = strings.Replace(s, s[3:4], "8", 1)
	fmt.Println(s)
}
