package main

import (
	"bufio"
	"fmt"
	"os"
)

const M = 101000

var (
	F = make([]int, M)
	N = make([]int, 2*M)
	y = make([]int, 2*M)
	Q = make([]int, 2*M)
	t int
)

func DFS(u, f int) int {
	K := 0
	P := 0
	S := 0
	for j := F[u]; j > 0; j = N[j] {
		if y[j] != f {
			K = DFS(y[j], u)
			S |= K & P
			P |= K
		}
	}
	return (P | ((1 << (Q[S] + 1)) - 1)) + 1
}

func A(a, b int) {
	t++
	y[t] = b
	N[t] = F[a]
	F[a] = t
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		A(u, v)
		A(v, u)
	}

	Q[0] = -1
	for i := 2; (i >> 17) == 0; i++ {
		Q[i] = Q[i>>1] + 1
	}
	fmt.Println(Q[DFS(1, 0)])
}
