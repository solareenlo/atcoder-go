package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 1000000007

	var x, rem, p [2100]int
	var A, C [2100][2100]int
	var N, K, X, Y int
	fmt.Fscan(in, &N, &K, &X, &Y)
	C[0][0] = 1
	for i := 0; i < 2050; i++ {
		for j := 0; j < i+1; j++ {
			C[i+1][j] = (C[i+1][j] + C[i][j]) % MOD
			C[i+1][j+1] = (C[i+1][j+1] + C[i][j]) % MOD
		}
	}
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &x[i])
	}
	for i := 0; i < N; i++ {
		rem[x[i]]++
	}
	for i := 0; i < K; i++ {
		for j := 0; j < K; j++ {
			fmt.Fscan(in, &A[i][j])
			A[i][j] ^= X
		}
	}
	Y ^= X
	ret := 0
	for i := 0; i < K; i++ {
		for j := 0; j < K; j++ {
			if (A[0][i]^A[0][j]) != A[i][j] && (A[0][i]^A[0][j]) != (A[i][j]^Y) {
				fmt.Println(0)
				return
			}
		}
	}
	for i := 0; i < 2048; i++ {
		if rem[i] == 0 {
			continue
		}
		if A[0][0] != 0 && A[0][0] != Y {
			continue
		}
		rem[i]--
		for j := 0; j < 2048; j++ {
			p[j] = 0
		}
		for j := 1; j < K; j++ {
			p[min(A[0][j]^i, A[0][j]^i^Y)]++
		}
		tmp := 1
		for j := 0; j < 2048; j++ {
			P := rem[j]
			Q := rem[j^Y]
			if P+Q < p[j] {
				tmp = 0
				break
			}
			ks := 0
			for k := 0; k <= p[j]; k++ {
				if P < k || Q < p[j]-k {
					continue
				}
				ks = (ks + C[p[j]][k]) % MOD
			}
			tmp = tmp * ks % MOD
		}
		rem[i]++
		ret = (ret + tmp) % MOD
	}
	fmt.Println(ret)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
