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
	P := make([]int, N)
	Q := make([]int, N)
	for i := range Q {
		Q[i] = 1
	}
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &P[i])
		if 0 < P[i] {
			P[i]--
			Q[P[i]] = 0
		}
	}
	R := make([]int, 0)
	for i := 0; i < N; i++ {
		if Q[i] != 0 {
			R = append(R, i)
		}
	}
	for i := 0; i < N; i++ {
		if P[i] == -1 {
			P[i] = R[len(R)-1]
			R = R[:len(R)-1]
		}
	}
	mn := int(1e9)
	mx := 0
	for i := 0; i < N; i++ {
		mn = min(mn, (P[i]+1)*(i+1))
		mx = max(mx, (P[i]+1)*(i+1))
	}
	fmt.Println(mx - mn)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
