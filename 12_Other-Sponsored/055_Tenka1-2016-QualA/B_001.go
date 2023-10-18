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
	var p [1001]int
	for i := 1; i < n; i++ {
		fmt.Fscan(in, &p[i])
	}
	var a [1001]int
	for i := 0; i < m; i++ {
		var P, Q int
		fmt.Fscan(in, &P, &Q)
		a[P] = Q
	}
	H := make([][]int, 1001)
	for i := 1; i < n; i++ {
		G := i
		if a[i] == 0 {
			continue
		}
		for G != 0 {
			H[G] = append(H[G], i)
			G = p[G]
		}
	}
	S := 0
	for i := 1; i < n; i++ {
		M := int(1e8)
		for j := 0; j < len(H[i]); j++ {
			M = min(M, a[H[i][j]])
		}
		for j := 0; j < len(H[i]); j++ {
			a[H[i][j]] -= M
		}
		S += M
	}
	fmt.Println(S)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
