package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var p, q, res, x int
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		p += a
		q = max(q, p)
		res = max(res, x+q)
		x += p
	}

	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
