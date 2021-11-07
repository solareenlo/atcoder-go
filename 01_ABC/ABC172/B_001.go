package main

import "fmt"

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	n := len(s)

	cnt := 0
	for i := 0; i < n; i++ {
		if s[i] != t[i] {
			cnt++
		}
	}

	fmt.Println(cnt)
}
