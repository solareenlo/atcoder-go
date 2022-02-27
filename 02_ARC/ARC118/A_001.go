package main

import "fmt"

var t int

func f(x int) int {
	return (100 + t) * x / 100
}

func main() {
	var n int
	fmt.Scan(&t, &n)

	ok := 0
	ng := 100 * n
	for ok+1 < ng {
		x := (ok + ng) / 2
		if f(x)-x < n {
			ok = x
		} else {
			ng = x
		}
	}
	fmt.Println(f(ok) + 1)
}
