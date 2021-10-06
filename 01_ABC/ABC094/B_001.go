package main

import "fmt"

func main() {
	var n, m, x, t int
	fmt.Scan(&n, &m, &x)

	cnt := 0
	for i := 0; i < m; i++ {
		fmt.Scan(&t)
		if t < x {
			cnt++
		}
	}
	fmt.Println(min(cnt, m-cnt))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
