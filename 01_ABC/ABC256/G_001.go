package main

import "fmt"

func main() {
	var n, d int
	fmt.Scan(&n, &d)
	K = 2

	binom = make([]int, d)
	binom[0] = 1
	for i := 1; i < d; i++ {
		binom[i] = divMod(binom[i-1]*(d-1-(i-1))%mod, i)
	}

	ans := 0
	coef := make([][]int, d+1)
	for i, _ := range coef {
		coef[i] = make([]int, d+1)
	}
	for c := 0; c <= d+1; c++ {
		coef[0][0] = b(c)
		coef[0][1] = b(c - 1)
		coef[1][0] = b(c - 1)
		coef[1][1] = b(c - 2)
		ret := matPowMod(coef, n)
		ans += (ret[0][0] + ret[1][1]) % mod
		ans %= mod
	}

	fmt.Println(ans)
}

var binom []int

func b(i int) int {
	if i < 0 || i >= len(binom) {
		return 0
	}
	return binom[i]
}

const mod = 998244353

var K int

func matMulMod(A, B [][]int) [][]int {
	C := make([][]int, K)
	for i := range C {
		C[i] = make([]int, K)
	}
	for i := 0; i < K; i++ {
		for j := 0; j < K; j++ {
			for k := 0; k < K; k++ {
				C[i][j] += A[i][k] * B[k][j] % mod
				C[i][j] %= mod
			}
		}
	}
	return C
}

func matPowMod(A [][]int, n int) [][]int {
	T := make([][]int, K)
	for i := range T {
		T[i] = make([]int, K)
	}
	if n == 0 {
		for i := 0; i < K; i++ {
			for j := 0; j < K; j++ {
				if i == j {
					T[i][j] = 1
				} else {
					T[i][j] = 0
				}
			}
		}
		return T
	}
	T = matPowMod(A, n>>1)
	T = matMulMod(T, T)
	if n&1 != 0 {
		T = matMulMod(T, A)
	}
	return T
}

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
