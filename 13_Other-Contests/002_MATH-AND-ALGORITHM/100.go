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

	K = 3

	var q int
	fmt.Fscan(in, &q)

	for i := 0; i < q; i++ {
		var x, y, z float64
		var t int
		fmt.Fscan(in, &x, &y, &z, &t)

		a := make([][]float64, 3)
		for i := range a {
			a[i] = make([]float64, 3)
		}
		a[0][0] = 1 - x
		a[0][1] = y
		a[1][1] = 1 - y
		a[1][2] = z
		a[2][0] = x
		a[2][2] = 1 - z

		b := matPowMod(a, t)

		A, B, C := 0.0, 0.0, 0.0
		for i := 0; i < 3; i++ {
			A += b[0][i]
			B += b[1][i]
			C += b[2][i]
		}

		fmt.Fprintln(out, A, B, C)
	}
}

var K int

func matMulMod(A, B [][]float64) [][]float64 {
	C := make([][]float64, K)
	for i := range C {
		C[i] = make([]float64, K)
	}
	for i := 0; i < K; i++ {
		for j := 0; j < K; j++ {
			for k := 0; k < K; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}
	return C
}

func matPowMod(A [][]float64, n int) [][]float64 {
	T := make([][]float64, K)
	for i := range T {
		T[i] = make([]float64, K)
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
