package main

import (
	"fmt"
	"strings"
)

var n int
var s []string

func main() {
	fmt.Scan(&n)
	s = make([]string, n)
	for i := range s {
		s[i] = "0"
	}
	m := 1
	for i := 0; i < n; i++ {
		m *= 10
	}
	fmt.Println(m - 1)
	f(0)
}

func f(id int) {
	if id == n {
		fmt.Println(strings.Join(s, ""))
	} else {
		if s[id] == "0" {
			for s[id] <= "9" {
				f(id + 1)
				if s[id] == "9" {
					break
				}
				s[id] = string(s[id][0] + 1)
			}
		} else {
			for s[id] >= "0" {
				f(id + 1)
				if s[id] == "0" {
					break
				}
				s[id] = string(s[id][0] - 1)
			}
		}
	}
}
