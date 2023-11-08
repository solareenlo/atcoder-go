package main

import "fmt"

func main() {
	var n, a int
	fmt.Scan(&n, &a)
	b := 0
	for i := 1; i <= n; i++ {
		var t int
		fmt.Scan(&t)
		b = max(b, t)
	}
	fmt.Println(max(0, b-a+1))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
