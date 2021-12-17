package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	ss := strings.Split(s, "")

	m := map[string]bool{}
	for i := range ss {
		m[ss[i]] = true
	}
	n := len(m)
	switch n {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(3)
	default:
		fmt.Println(6)
	}
}
