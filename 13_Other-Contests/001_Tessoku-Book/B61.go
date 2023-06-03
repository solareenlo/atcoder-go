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
	A := make([][]int, n+1)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		A[a] = append(A[a], b)
		A[b] = append(A[b], a)
	}
	a := 0
	b := 0
	for i := 1; i <= n; i++ {
		if a < len(A[i]) {
			a = len(A[i])
			b = i
		}
	}
	fmt.Println(b)
}
