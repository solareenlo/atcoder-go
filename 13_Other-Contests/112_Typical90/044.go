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

	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
	}

	r := 0
	for Q > 0 {
		Q--
		var T, x, y int
		fmt.Fscan(in, &T, &x, &y)
		if T == 1 {
			A[(x-1+r)%N], A[(y-1+r)%N] = A[(y-1+r)%N], A[(x-1+r)%N]
		} else if T == 2 {
			r = (r - 1 + N) % N
		} else {
			fmt.Fprintln(out, A[(x-1+r)%N])
		}
	}
}
