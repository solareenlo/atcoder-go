package main

import (
	"fmt"
	"strings"
)

func main() {
	s := make([]string, 3)
	for i := range s {
		fmt.Scan(&s[i])
	}
	fmt.Print(
		strings.ToUpper(string(s[0][0])),
		strings.ToUpper(string(s[1][0])),
		strings.ToUpper(string(s[2][0])))
}
