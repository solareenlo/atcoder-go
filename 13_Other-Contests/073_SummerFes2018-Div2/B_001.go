package main

import "fmt"

func main() {
	var N int
	var S string
	fmt.Scan(&N, &S)
	A, B := -1, N
	for i := 0; i < N; i++ {
		if S[i] == 'D' {
			A = i
		}
		if S[i] == 'K' {
			B = min(B, i)
		}
	}
	fmt.Println(max(0, B-A))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
