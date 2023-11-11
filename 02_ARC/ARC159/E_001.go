package main

import (
	"bufio"
	"fmt"
	"os"
)

var MM int
var A, B [105]int

func so(L, R, q, x int) int {
	q %= MM
	M := (A[q]*L + B[q]*R) / (A[q] + B[q])
	if M == x {
		return 2 * (M - L)
	} else if x < M {
		return so(L, M-1, q+1, x) + 1
	}
	return 2*(M-L) + 1 + so(M+1, R, q+1, x)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N int
	fmt.Fscan(in, &N, &MM)
	for i := 0; i < MM; i++ {
		fmt.Fscan(in, &A[i], &B[i])
	}
	var Q int
	fmt.Fscan(in, &Q)
	for Q > 0 {
		Q--
		var s, t int
		fmt.Fscan(in, &s, &t)
		fmt.Fprintln(out, so(1, N, 0, t)-so(1, N, 0, s))
	}
}
