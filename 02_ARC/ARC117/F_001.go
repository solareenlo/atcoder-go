package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 300300

var (
	n int
	m int
	F = [N][2]int{}
	A = [N]int{}
)

func ok(w int) bool {
	F[0][1] = -(1 << 60)
	A[m+1] = -(1 << 60)
	for i := 1; i < n+1; i++ {
		F[i][0] = max(F[i-1][0], F[i-1][1]+A[i+1+n]-w)
		F[i][1] = max(F[i-1][1], F[i-1][0]+A[i+1])
		if F[i][0] > w {
			return false
		}
	}
	F[0][1] = max(A[1], F[n][0])
	x := F[0][1]
	if F[0][1]+A[n+1] > w {
		return false
	}
	for i := 1; i < n+1; i++ {
		F[i][0] = max(F[i-1][0], F[i-1][1]+A[i+1+n]-w)
		F[i][1] = max(F[i-1][1], F[i-1][0]+A[i+1])
		if F[i][1] > w {
			return false
		}
	}
	return x == max(A[1], F[n][0])
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)
	m = 2 * n
	for i := 1; i < m+1; i++ {
		fmt.Fscan(in, &A[i])
	}

	l := 0
	for i := 1; i < n+1; i++ {
		l = max(l, int(A[i]+A[i+n]))
	}

	r := 1000000000
	for l <= r {
		m := (l + r) >> 1
		if ok(m) {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	fmt.Println(l)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
