package main

import (
	"bufio"
	"fmt"
	"os"
)

var N, M, K int
var A [100100]int

func check(k int) bool {
	sum := 0
	for i := 0; i < N; i++ {
		sum += (max(0, k-A[i]) + K - 1) / K
		if sum > M {
			return true
		}
	}
	return false
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &N, &M, &K)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
	}
	l := 0
	r := int(1e18)
	for r-l > 1 {
		m := l + (r-l)/2
		if check(m) {
			r = m
		} else {
			l = m
		}
	}
	for i := 0; i < N; i++ {
		t := (max(0, l-A[i]) + K - 1) / K
		M -= t
		A[i] += t * K
	}
	for i := 0; i < N; i++ {
		if M > 0 && A[i] == l {
			M--
			A[i] += K
		}
	}
	for i := 0; i < N; i++ {
		if i+1 == N {
			fmt.Fprintln(out, A[i])
		} else {
			fmt.Fprintf(out, "%d ", A[i])
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
