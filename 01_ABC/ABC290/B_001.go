package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	var S string
	fmt.Scan(&S)
	s := strings.Split(S, "")
	a := 0
	for i := 0; i < n; i++ {
		if s[i] == "o" && k == a {
			s[i] = "x"
		}
		if s[i] == "o" && a < k {
			a++
		}
	}
	fmt.Println(strings.Join(s, ""))
}
