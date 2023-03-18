package main

import (
	"fmt"
)

func main() {
	var N, M, S int
	fmt.Scan(&N, &M, &S)
	T := make([]int, N)
	K := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&T[i], &K[i])
	}
	A := make([]int, S)
	for i := 0; i < N; i++ {
		for j := T[i]; j < S; j++ {
			A[j] += K[i]
		}
	}
	ans := 0
	for i := 0; i < S; i++ {
		if A[i] >= M {
			ans++
		}
	}
	fmt.Println(ans)
}
