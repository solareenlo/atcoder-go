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

	D := make([]int, N)
	for i := range D {
		D[i] = -1
	}
	T := make([]int, N)
	for i := 0; i < N; i++ {
		T[i] = i
	}

	for qi := 0; qi < Q; qi++ {
		var f, t, x int
		fmt.Fscan(in, &f, &t, &x)
		f--
		t--
		x--
		T[t], T[f] = T[f], T[t]
		T[f], D[x] = D[x], T[f]
	}

	A := make([]int, N)
	for i := 0; i < N; i++ {
		c := T[i]
		for c != -1 {
			A[c] = i + 1
			c = D[c]
		}
	}
	for i := 0; i < N; i++ {
		fmt.Fprintln(out, A[i])
	}
}
