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

	var N int
	fmt.Fscan(in, &N)

	type tuple struct{ x, y, z int }
	P := make([]tuple, N+1)
	A := make([]int, N+1)
	B := make([]int, N+1)
	for i := 1; i <= N; i++ {
		fmt.Fscan(in, &A[i])
	}
	for i := 1; i <= N; i++ {
		fmt.Fscan(in, &B[i])
	}

	for i := 1; i <= N; i++ {
		P[i] = tuple{-A[i] - B[i], -A[i], i}
	}
	P = P[1:]
	sort.Slice(P, func(i, j int) bool {
		return P[i].x < P[j].x || (P[i].x == P[j].x && P[i].y < P[j].y) || (P[i].x == P[j].x && P[i].y == P[j].y && P[i].z < P[j].z)
	})

	for i := 0; i < N; i++ {
		fmt.Fprint(out, P[i].z, " ")
	}
}
