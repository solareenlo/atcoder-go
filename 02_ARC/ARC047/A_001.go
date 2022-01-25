package main

import "fmt"

func main() {
	var n, l int
	var s string
	fmt.Scan(&n, &l, &s)

	cnt := 0
	tab := 1
	for i := 0; i < n; i++ {
		if s[i] == '+' {
			tab++
		}
		if s[i] == '-' && tab >= 0 {
			tab--
		}
		if tab > l {
			cnt++
			tab = 1
		}
	}

	fmt.Println(cnt)
}
