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

	var N int
	fmt.Fscan(in, &N)
	S := make([]string, N)
	A := make([]int, N)
	x := 1000000001
	var y int
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &S[i], &A[i])
		if A[i] < x {
			x = A[i]
			y = i
		}
	}
	for i := y; i < N; i++ {
		fmt.Fprintln(out, S[i])
	}
	for i := 0; i < y; i++ {
		fmt.Fprintln(out, S[i])
	}
}
