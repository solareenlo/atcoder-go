package main

import "fmt"

var K, N int

const mod = 1_000_000_007

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

func main() {
	fmt.Scan(&K, &N)
	K = 2*K - 2
	N -= 2
	var a int
	var A [][]int
	if K < 3 {
		A = [][]int{{1, 1}, {1, 0}}
		A = matPowMod(A, N)
		a = A[0][0]*2%mod + A[0][1]
		a %= mod
	}
	if K == 4 {
		A = [][]int{{0, 1, 2, 0}, {1, 0, 0, 0}, {1, 0, 0, 1}, {0, 0, 1, 0}}
		A = matPowMod(A, N)
		a = A[0][0]*3%mod + A[0][3]
		a %= mod
	}
	if K > 5 {
		A = [][]int{{1, 1, 2, 0, 1, 0}, {1, 0, 0, 0, 0, 0}, {1, 0, 1, 0, 0, 0},
			{0, 0, 1, 0, 0, 0}, {1, 0, 0, 0, 0, 1}, {0, 0, 0, 0, 1, 0}}
		A = matPowMod(A, N)
		a = A[0][0]*5%mod + A[0][1] + A[0][2]*2%mod + A[0][3] + A[0][4] + A[0][5]
		a %= mod
	}
	fmt.Println(a)
}
