package main

import "fmt"

func cntDisits(n int) int {
	digits := 0
	for n > 0 {
		n /= 10
		digits++
	}
	return digits
}

func main() {
	var n int
	fmt.Scan(&n)

	res := cntDisits(n)
	for a := 1; a*a <= n; a++ {
		if n%a != 0 {
			continue
		}
		now := max(cntDisits(a), cntDisits(n/a))
		res = min(res, now)
	}
	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
