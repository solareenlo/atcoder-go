package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)

	var P, A, V, C, cand, MA [202020]int
	var E [202020][]int

	for i := 1; i < N; i++ {
		fmt.Fscan(in, &P[i])
		P[i]--
		E[P[i]] = append(E[P[i]], i)
	}

	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
	}

	var M int
	fmt.Fscan(in, &M)
	for i := 0; i < M; i++ {
		var x int
		fmt.Fscan(in, &x)
		V[x-1]++
		C[x-1]++
	}

	for i := N - 1; i >= 0; i-- {
		if V[i] != 0 {
			cand[i] = 1
		}
		for _, e := range E[i] {
			if C[e] != 0 && C[i] != 0 {
				cand[i] = 1
			}
			C[i] += C[e]
		}
	}

	ret := 0
	for i := 0; i < N; i++ {
		if i != 0 {
			MA[i] = MA[P[i]]
		}
		if cand[i] != 0 {
			MA[i] = max(MA[i], A[i])
		}
		ret += V[i] * MA[i]
	}
	fmt.Println(ret)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
