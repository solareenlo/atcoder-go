package main

import (
	"bufio"
	"fmt"
	"os"
)

const K = 3
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

var (
	A  = make([][][]int, 33)
	R  [][]int
	b  = make([]int, K)
	b2 = make([]int, K)
)

func mul(m [][]int) {
	for i := 0; i < 3; i++ {
		b2[i] = b[i]
		b[i] = 0
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			b[i] = (b[i] + m[i][j]*b2[j]) % mod
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(in, &N, &M)

	A[0] = [][]int{{1, 0, 1}, {1, 1, 1}, {1, 2, 2}}
	R := [][]int{{0, mod - 2, mod - 1}, {1, 1, 1}, {1, 2, 2}}
	for i := 0; i < 30; i++ {
		A[i+1] = matMulMod(A[i], A[i])
	}

	b[0] = 1
	x := 0
	for j := 0; j < M; j++ {
		var y int
		fmt.Fscan(in, &y)
		for i := 30; i >= 0; i-- {
			if (y-x-1)>>i&1 != 0 {
				mul(A[i])
			}
		}
		mul(R)
		x = y
	}

	for i := 30; i >= 0; i-- {
		if (N-x)>>i&1 != 0 {
			mul(A[i])
		}
	}
	fmt.Println(b[2])
}
