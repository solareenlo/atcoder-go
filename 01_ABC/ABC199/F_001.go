package main

import "fmt"

var (
	N int
	G = make([][]int, 100)
)

func mul(A [][]int) {
	T := [100][100]int{}
	for k := 0; k < N; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				T[i][j] += A[i][k] * G[k][j] % mod
				T[i][j] %= mod
			}
		}
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			A[i][j] = T[i][j]
		}
	}
}

func main() {
	var M, K int
	fmt.Scan(&N, &M, &K)

	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}

	for i := range G {
		G[i] = make([]int, 100)
	}

	E := make([][]int, 100)
	for i := range E {
		E[i] = make([]int, 100)
	}
	for i := 0; i < N; i++ {
		G[i][i] = 1
		E[i][i] = 1
	}

	inv2 := divMod(divMod(1, 2), M)
	for i := 0; i < M; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		a--
		b--
		G[a][a] = (G[a][a] - inv2 + mod) % mod
		G[b][b] = (G[b][b] - inv2 + mod) % mod
		G[a][b] += inv2
		G[a][b] %= mod
		G[b][a] += inv2
		G[b][a] %= mod
	}

	for K > 0 {
		if K&1 != 0 {
			mul(E)
		}
		K >>= 1
		mul(G)
	}

	for i := 0; i < N; i++ {
		now := 0
		for j := 0; j < N; j++ {
			now += E[i][j] * A[j] % mod
			now %= mod
		}
		fmt.Println(now)
	}
}

const mod = 1000000007

func divMod(a, b int) int {
	ret := a * modInv(b)
	ret %= mod
	return ret
}

func modInv(a int) int {
	b, u, v := mod, 1, 0
	for b != 0 {
		t := a / b
		a -= t * b
		a, b = b, a
		u -= t * v
		u, v = v, u
	}
	u %= mod
	if u < 0 {
		u += mod
	}
	return u
}
