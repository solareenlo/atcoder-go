package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	var S string
	fmt.Scan(&n, &S)
	s := strings.Split(S, "")

	for i := 0; i < n-3+1; i++ {
		t := s[i] + s[i+1] + s[i+2]
		if t == "axa" || t == "ixi" || t == "uxu" || t == "exe" || t == "oxo" {
			s[i] = "."
			s[i+1] = "."
			s[i+2] = "."
		}
	}
	fmt.Println(strings.Join(s, ""))
}
