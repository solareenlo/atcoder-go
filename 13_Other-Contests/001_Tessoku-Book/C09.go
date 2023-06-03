package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	A := make([]int, n)
	B := make([]int, n+1)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &A[i])
	}
	B[1] = A[0]
	for i := 2; i <= n; i++ {
		B[i] = max(B[i-1], B[i-2]+A[i-1])
	}
	fmt.Println(max(B[n], B[n-1]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
