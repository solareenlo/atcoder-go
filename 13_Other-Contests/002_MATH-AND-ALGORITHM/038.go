package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N, Q int
	fmt.Fscan(in, &N, &Q)

	A := make([]int, N+1)
	for i := 1; i <= N; i++ {
		fmt.Fscan(in, &A[i])
	}

	L := make([]int, Q+1)
	R := make([]int, Q+1)
	for i := 1; i <= Q; i++ {
		fmt.Fscan(in, &L[i], &R[i])
	}

	B := make([]int, N+1)
	B[0] = 0
	for i := 1; i <= N; i++ {
		B[i] = B[i-1] + A[i]
	}

	for i := 1; i <= Q; i++ {
		fmt.Fprintln(out, B[R[i]]-B[L[i]-1])
	}
}
