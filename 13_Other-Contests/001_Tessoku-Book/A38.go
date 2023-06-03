package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var D, N int
	fmt.Fscan(in, &D, &N)
	A := make([]int, D+1)
	for i := range A {
		A[i] = 24
	}
	for i := 0; i < N; i++ {
		var L, R, H int
		fmt.Fscan(in, &L, &R, &H)
		for j := L; j <= R; j++ {
			A[j] = min(A[j], H)
		}
	}
	ans := -24
	for i := range A {
		ans += A[i]
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
