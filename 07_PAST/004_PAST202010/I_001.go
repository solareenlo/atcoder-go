package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)
	S := 0
	A := make([]int, 400005)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
		A[N+i] = A[i]
		S += A[N+i]
	}

	X := 0
	ans := 1 << 60
	j := 0
	for i := 0; i < N; i++ {
		for 2*(X+A[j]) <= S {
			X += A[j]
			j++
		}
		ans = min(ans, abs(S-2*X))
		X -= A[i]
	}
	fmt.Println(ans)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
