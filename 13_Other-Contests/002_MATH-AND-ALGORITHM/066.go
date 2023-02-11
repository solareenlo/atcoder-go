package main

import "fmt"

func main() {
	var N, K int
	fmt.Scan(&N, &K)

	ans := N * N * N
	for i := 0; i < N; i++ {
		for j := max(i-K+1, 0); j < min(i+K, N); j++ {
			for k := max(max(i-K, j-K)+1, 0); k < min(min(i+K, j+K), N); k++ {
				ans--
			}
		}
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
