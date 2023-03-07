package main

import "fmt"

func main() {
	var n, t int
	fmt.Scan(&n, &t)
	ti := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&ti[i])
	}
	bm := 0
	for i := 1; i <= t; i++ {
		bn := 0
		for j := 1; j <= n; j++ {
			if i%ti[j] == 0 {
				bn++
			}
		}
		bm = max(bn, bm)
	}
	fmt.Println(bm)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
