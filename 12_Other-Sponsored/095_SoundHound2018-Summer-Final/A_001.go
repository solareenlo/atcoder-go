package main

import "fmt"

func main() {
	var C, D int
	fmt.Scan(&C, &D)
	ans := 0
	mi, ma := 140, 170
	for mi < D {
		ans += max(0, min(ma, D)-max(mi, C))
		mi *= 2
		ma *= 2
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
