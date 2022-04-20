package main

import "fmt"

var (
	N int
	B = [15]int{}
	A = [15][15]int{}
)

func a(n int) int {
	if n == N {
		cnt := 0
		for i := 0; i < N; i++ {
			for j := i + 1; j < N; j++ {
				if B[i] == B[j] {
					cnt += A[i][j]
				}
			}
		}
		return cnt
	}
	mx := -1000000000
	for B[n] = 0; B[n] < 3; B[n]++ {
		mx = max(mx, a(n+1))
	}
	return mx
}

func main() {
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			fmt.Scan(&A[i][j])
		}
	}
	fmt.Println(a(0))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
