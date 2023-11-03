package main

import (
	"fmt"
	"strings"
)

func main() {
	var S string
	fmt.Scan(&S)
	s := strings.Split(S, "")
	for i := 0; i < len(s); i++ {
		if i%2 == 0 && i+1 < len(s) {
			tmp := s[i]
			s[i] = s[i+1]
			s[i+1] = tmp
		}
	}
	fmt.Println(strings.Join(s, ""))
}
