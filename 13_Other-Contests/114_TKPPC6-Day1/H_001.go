package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, X, Y int
	fmt.Fscan(in, &N, &X, &Y)
	var B, C int
	fmt.Fscan(in, &B, &C)

	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
	}
	sort.Ints(A)

	if B < C {
		X, Y = Y, X
		B, C = C, B
	}

	for i := 0; i < X; i++ {
		A[i] += B
	}
	sort.Ints(A)

	for i := 0; i < Y; i++ {
		A[i] += C
	}

	ma, mi := A[0], A[0]
	for _, a := range A {
		ma = max(ma, a)
		mi = min(mi, a)
	}
	fmt.Println(ma - mi)
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
