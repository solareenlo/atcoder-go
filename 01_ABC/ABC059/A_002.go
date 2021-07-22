package main

import (
	"fmt"
	"strings"
)

func main() {
	s := make([]string, 3)
	fmt.Scan(&s[0], &s[1], &s[2])
	fmt.Println(strings.ToUpper(s[0][:1] + s[1][:1] + s[2][:1]))
}
