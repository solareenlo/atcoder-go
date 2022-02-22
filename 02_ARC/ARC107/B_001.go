package main

import "fmt"

var (
	n int
	k int
)

func f(k int) int { return max(0, n-abs(k-n-1)) }

func main() {
	fmt.Scan(&n, &k)

	res := 0
	for i := 2; i <= 2*n; i++ {
		res += f(i) * f(i-k)
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
