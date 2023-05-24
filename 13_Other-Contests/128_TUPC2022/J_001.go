package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var A, P, C, inv, nP [2 << 17]int

	var N, M int
	fmt.Fscan(in, &N, &M)
	for i := 1; i <= N; i++ {
		P[i] = i
	}
	for i := 0; i < M; i++ {
		fmt.Fscan(in, &A[i])
		P[A[i]], P[A[i]+1] = P[A[i]+1], P[A[i]]
	}
	for i := 1; i < N+1; i++ {
		inv[P[i]] = i
	}
	sz := 0
	for i := 1; i < N+1; i++ {
		if C[i] == 0 {
			sz++
			for C[i] == 0 {
				C[i] = sz
				i = P[i]
			}
		}
	}
	fmt.Println(N - sz)
	for i := 1; i <= N; i++ {
		nP[i] = i
	}
	var U, V []int
	for i := 0; i < M; i++ {
		u := nP[A[i]]
		v := nP[A[i]+1]
		if C[u] != C[v] {
			nP[A[i]], nP[A[i]+1] = nP[A[i]+1], nP[A[i]]
		} else {
			fmt.Println(2*i+1, A[i])
			inv[u], inv[v] = inv[v], inv[u]
			U = make([]int, 0)
			U = append(U, u)
			V = make([]int, 0)
			V = append(V, v)
			for inv[U[len(U)-1]] != U[0] && inv[V[len(V)-1]] != V[0] {
				U = append(U, inv[U[len(U)-1]])
				V = append(V, inv[V[len(V)-1]])
			}
			sz++
			if inv[U[len(U)-1]] != U[0] {
				U, V = V, U
			}
			for _, w := range U {
				C[w] = sz
			}
		}
	}
}
