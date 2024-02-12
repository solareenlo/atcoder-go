package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	sz := len(s)
	s += "\n"

	for n > 0 {
		n--
		var t string
		cnt := 1
		for i := 0; i < sz; i++ {
			if s[i] == s[i+1] {
				cnt++
			} else {
				t += strconv.Itoa(cnt) + string(s[i])
				cnt = 1
			}
		}
		sz = len(t)
		s = t + "\n"
	}

	fmt.Print(s)
}
