package main

import "fmt"

func main() {
	var N, T int
	fmt.Scan(&N, &T)

	pos := 0
	for i := 0; i < N; i++ {
		var a int
		fmt.Scan(&a)
		pos = max(pos, a)
		pos = a + (pos-a+T-1)/T*T
	}
	fmt.Println(pos)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
