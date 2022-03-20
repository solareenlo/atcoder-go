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

	L := make([]int, 100009)
	R := make([]int, 100009)
	X := make([]int, 100009)
	B := make([]int, 100009)
	for i := 1; i <= Q; i++ {
		fmt.Fscan(in, &L[i], &R[i], &X[i])
		B[L[i]] += X[i]
		B[R[i]+1] -= X[i]
	}

	for i := 2; i <= N; i++ {
		if B[i] > 0 {
			fmt.Fprint(out, "<")
		}
		if B[i] == 0 {
			fmt.Fprint(out, "=")
		}
		if B[i] < 0 {
			fmt.Fprint(out, ">")
		}
	}
}
