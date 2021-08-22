package main

import (
	"fmt"
	"strings"
)

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	at := "atcoder"
	res := "You can win"
	for i := 0; i < len(s); i++ {
		if s[i] == '@' && strings.Contains(at, string(t[i])) {
			continue
		}
		if t[i] == '@' && strings.Contains(at, string(s[i])) {
			continue
		}
		if s[i] != t[i] {
			res = "You will lose"
			break
		}
	}
	fmt.Println(res)
}
