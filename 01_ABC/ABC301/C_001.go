package main

import (
	"fmt"
	"strings"
)

func main() {
	var s, t string
	fmt.Scan(&s, &t)
	cs := make(map[rune]int)
	ct := make(map[rune]int)
	for _, c := range s {
		cs[c]++
	}
	for _, c := range t {
		ct[c]++
	}
	a := "atcoder"
	extra := 0
	tmp := "abcdefghijklmnopqrstuvwxyz"
	for _, c := range tmp {
		if cs[c] != ct[c] {
			if idx := strings.IndexRune(a, c); idx == -1 {
				fmt.Println("No")
				return
			}
			if cs[c] < ct[c] {
				extra += ct[c] - cs[c]
			}
		}
	}
	if extra <= cs['@'] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
