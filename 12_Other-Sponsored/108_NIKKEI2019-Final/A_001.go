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

	var A [3003]int

	var N int
	fmt.Fscan(in, &N)
	for i := 1; i <= N; i++ {
		fmt.Fscan(in, &A[i])
		A[i] += A[i-1]
	}
	for k := 1; k <= N; k++ {
		s := 0
		for i := 0; i <= N-k; i++ {
			s = max(s, A[i+k]-A[i])
		}
		fmt.Fprintln(out, s)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
