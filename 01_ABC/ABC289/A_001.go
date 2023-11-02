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
		if s[i] == "0" {
			s[i] = "1"
		} else {
			s[i] = "0"
		}
	}
	fmt.Println(strings.Join(s, ""))
}
