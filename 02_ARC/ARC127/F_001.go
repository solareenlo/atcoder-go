package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(A, B, V, M, X, Y int) int {
	if V >= A {
		return solve(A, B, V%A, M, X, Y) + V/A*Y
	}
	if V+B > M {
		return 0
	}
	return solve(B%A, A, M-B+B%A-V, M-B+B%A, Y, X+B/A*Y)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var T int
	fmt.Fscan(in, &T)

	for i := 0; i < T; i++ {
		var A, B, V, M int
		fmt.Fscan(in, &A, &B, &V, &M)
		if A+B-1 <= M {
			fmt.Fprintln(out, M+1)
		} else {
			fmt.Fprintln(out, solve(A, B, V, M, 1, 1)+solve(A, B, M-V, M, 1, 1)+1)
		}
	}
}
