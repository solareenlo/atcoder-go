package main

import "fmt"

func main() {
	type pair struct {
		x, y int
	}

	var n int
	fmt.Scan(&n)
	var s string
	fmt.Scan(&s)
	x, y := 0, 0
	vis := make(map[pair]bool)
	vis[pair{x, y}] = true
	ok := false
	for i := 0; i < len(s); i++ {
		if s[i] == 'R' {
			x++
		}
		if s[i] == 'L' {
			x--
		}
		if s[i] == 'U' {
			y++
		}
		if s[i] == 'D' {
			y--
		}
		if (vis[pair{x, y}]) {
			ok = true
		}
		vis[pair{x, y}] = true
	}
	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
