package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	l := 0
	r := 2_000_000_000
	for r-l > 1 {
		m := (l + r) / 2
		if m*(m+1) <= (n+1)*2 {
			l = m
		} else {
			r = m
		}
	}
	fmt.Println(n - l + 1)
}
