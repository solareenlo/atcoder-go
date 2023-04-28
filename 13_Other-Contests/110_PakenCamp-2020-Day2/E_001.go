package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 1000000007
	const L = 50000
	const K = 224

	var N, M int
	fmt.Fscan(in, &N, &M)
	ni := make([]int, N+1)
	for i := range ni {
		ni[i] = 1
	}
	for i := 0; i < N; i++ {
		ni[i+1] = (ni[i] * 2) % MOD
	}

	var ten, sen [2*K + 1][L + 1]int
	for i := 0; i < N; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		if a*a < L {
			sen[a+K][b]++
		} else {
			A := (L - b) / a
			B := (L + b) / a
			B *= -1
			for j := max(0, min(A, B)-1); j <= max(A, B)+1; j++ {
				if a*j+b > L || a*j+b < 0 {
					continue
				} else {
					ten[j][a*j+b]++
				}
			}
		}
	}

	Z := 0
	for i := 0; i < M; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		A := 0
		if x <= 2*K {
			A += ten[x][y]
		}
		for j := -223; j < 224; j++ {
			if y-j*x >= 0 && y-j*x <= L {
				A += sen[j+K][y-j*x]
			}
		}
		Z += ni[N] - ni[N-A]
		Z %= MOD
	}
	fmt.Println((Z + MOD) % MOD)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
