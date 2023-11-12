package main

import "fmt"

func sqrt(x int) int {
	ok, ng := 0, int(2e9)
	for ng-ok > 1 {
		mid := (ng + ok) / 2
		if mid*mid < x {
			ok = mid
		} else {
			ng = mid
		}
	}
	return ok
}

func main() {
	var n int
	fmt.Scan(&n)
	s := sqrt(n)
	n -= s * s
	fmt.Println(abs(s-n+1) + 1)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
