package main

import "fmt"

func main() {
	ans := 0
	for i := 0; i < 8; i++ {
		var t int
		fmt.Scan(&t)
		ans = max(ans, t)
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
