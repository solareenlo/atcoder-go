package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	mini := 1 << 32
	var t int
	for i := 0; i < n; i++ {
		fmt.Scan(&t)
		mini = min(mini, t)
		x -= t
	}
	fmt.Println(n + x/mini)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
