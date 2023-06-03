package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	A := make([][]int, m)
	for i := range A {
		A[i] = make([]int, n)
	}
	D := make([]int, 1<<n)
	for i := range D {
		D[i] = int(1e9)
	}
	D[0] = 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &A[i][j])
		}
	}
	for i := 0; i < m; i++ {
		x := 0
		for j := 0; j < n; j++ {
			if A[i][j] != 0 {
				x |= 1 << j
			}
		}
		E := make([]int, len(D))
		copy(E, D)
		for j := 0; j < 1<<n; j++ {
			E[x|j] = min(E[x|j], D[j]+1)
		}
		D = E
	}
	if D[(1<<n)-1] == int(1e9) {
		fmt.Println(-1)
	} else {
		fmt.Println(D[(1<<n)-1])
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
