package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, Q int
	fmt.Fscan(in, &N, &Q)

	var A [1 << 17]int
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
	}

	ans := 0
	for i := N - 1; i > 0; i-- {
		A[i] -= A[i-1]
		ans += abs(A[i])
	}
	for i := 0; i < Q; i++ {
		var L, R, V int
		fmt.Fscan(in, &L, &R, &V)
		if L > 1 {
			ans -= abs(A[L-1]) - abs(A[L-1]+V)
			A[L-1] += V
		}
		if R < N {
			ans -= abs(A[R]) - abs(A[R]-V)
			A[R] -= V
		}
		fmt.Println(ans)
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
