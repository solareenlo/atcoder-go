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

	var N, M, K int
	fmt.Fscan(in, &N, &M, &K)

	H := make([]int, N)
	for i := range H {
		fmt.Fscan(in, &H[i])
	}

	D := make([]int, N)
	for i := 0; i < N; i++ {
		D[i] = -1
	}
	Q := make([]int, 0)
	for i := 0; i < K; i++ {
		var k int
		fmt.Fscan(in, &k)
		k--
		Q = append(Q, k)
		D[k] = 0
	}

	E := make([][]int, 200000)
	for i := 0; i < M; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		if H[u] < H[v] {
			E[u] = append(E[u], v)
		} else {
			E[v] = append(E[v], u)
		}
	}
	for len(Q) > 0 {
		p := Q[0]
		Q = Q[1:]
		for _, e := range E[p] {
			if D[e] != -1 {
				continue
			}
			D[e] = D[p] + 1
			Q = append(Q, e)
		}
	}

	for i := 0; i < N; i++ {
		fmt.Fprintln(out, D[i])
	}
}
