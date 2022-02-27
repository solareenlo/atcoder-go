package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var K, N, M int
	fmt.Fscan(in, &K, &N, &M)

	A := make([]int, K)
	for i := range A {
		fmt.Fscan(in, &A[i])
	}

	type pair struct{ x, y int }
	P := make([]pair, K)
	B := make([]int, K)
	sum := 0
	for i := 0; i < K; i++ {
		B[i] = M * A[i] / N
		P[i].x = N*B[i] - M*A[i]
		P[i].y = i
		sum += B[i]
	}

	sort.Slice(P, func(i, j int) bool {
		return P[i].x < P[j].x
	})

	for i := 0; i < M-sum; i++ {
		B[P[i].y]++
	}
	for i := 0; i < K; i++ {
		fmt.Fprint(out, B[i], " ")
	}
	fmt.Fprintln(out)
}
