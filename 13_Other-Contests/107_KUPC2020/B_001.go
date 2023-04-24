package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const mod = 1000000007

	var N, K int
	fmt.Fscan(in, &N, &K)
	A := make([][]int, N)
	B := make([][]int, N)
	for i := 0; i < N; i++ {
		A[i] = make([]int, K)
		B[i] = make([]int, K+1)
	}
	for i := 0; i < N; i++ {
		for j := 0; j < K; j++ {
			fmt.Fscan(in, &A[i][j])
		}
	}
	for i := 0; i < K; i++ {
		B[0][i] = 1
	}
	for i := 1; i < N; i++ {
		for j := 0; j < K; j++ {
			a := lowerBound(A[i], A[i-1][j])
			B[i][a] = (B[i][a] + B[i-1][j]) % mod
		}
		for j := 0; j < K; j++ {
			B[i][j+1] = (B[i][j+1] + B[i][j]) % mod
		}
	}

	ans := 0
	for i := 0; i < K; i++ {
		ans = (ans + B[N-1][i]) % mod
	}
	fmt.Println(ans)
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}
