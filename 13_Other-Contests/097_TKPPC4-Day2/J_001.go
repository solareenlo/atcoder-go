package main

import (
	"bufio"
	"fmt"
	"os"
)

var N, M, P, S, T, K int
var A, B, C, D, H [3777]int
var X [77][]int
var dp1 [77][177][177]int
var dp2 [77][177]int

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &N, &M, &P, &S, &T, &K)
	S--
	T--
	for i := 1; i <= M; i++ {
		fmt.Fscan(in, &A[i], &B[i])
		A[i]--
		B[i]--
		X[A[i]] = append(X[A[i]], B[i])
	}
	for i := 0; i < N; i++ {
		H[i] = -1
	}
	for i := 1; i <= P; i++ {
		fmt.Fscan(in, &C[i], &D[i])
		C[i]--
		H[C[i]] = D[i]
	}
	init_1()
	init_2()
	fmt.Println(solve())
}

func init_1() {
	for i := 0; i < N*2; i++ {
		for j := 0; j < N*2; j++ {
			dp1[0][i][j] = -(1 << 60)
		}
	}
	for i := 0; i < N; i++ {
		dp1[0][i*2+0][i*2+0] = 0
		dp1[0][i*2+1][i*2+1] = 0
		for _, j := range X[i] {
			if H[j] >= 0 {
				dp1[0][i*2+0][j*2+0] = 0
				dp1[0][i*2+0][j*2+1] = -2 * H[j]
				dp1[0][i*2+1][j*2+0] = 2 * H[j]
				dp1[0][i*2+1][j*2+1] = 0
			} else {
				dp1[0][i*2+0][j*2+0] = 0
				dp1[0][i*2+1][j*2+1] = 0
			}
		}
	}
}

func init_2() {
	for t := 0; t < 63; t++ {
		for i := 0; i < N*2; i++ {
			for j := 0; j < N*2; j++ {
				dp1[t+1][i][j] = -(1 << 60)
			}
		}
		for k := 0; k < N*2; k++ {
			for i := 0; i < N*2; i++ {
				for j := 0; j < N*2; j++ {
					if dp1[t][i][k] == -(1<<60) || dp1[t][k][j] == -(1<<60) {
						continue
					}
					dp1[t+1][i][j] = max(dp1[t+1][i][j], min(1<<60, dp1[t][i][k]+dp1[t][k][j]))
					if t == 0 && i == 2 && j == 6 && dp1[t+1][i][j] == 8 {
						t += 0
					}
				}
			}
		}
	}
}

func solve() int {
	for i := 0; i < 64; i++ {
		for j := 0; j < N*2; j++ {
			dp2[i][j] = -(1 << 59)
		}
	}
	dp2[61][S*2+0] = H[S]
	dp2[61][S*2+1] = -H[S]

	r := 0
	for i := 60; i >= 0; i-- {
		for j := 0; j < N*2; j++ {
			for k := 0; k < N*2; k++ {
				if dp2[i+1][j] == -(1<<59) || dp1[i][j][k] == -(1<<60) {
					continue
				}
				dp2[i][k] = max(dp2[i][k], dp2[i+1][j]+dp1[i][j][k])
			}
		}
		v1 := dp2[i][T*2+0] - H[T]
		v2 := dp2[i][T*2+1] + H[T]
		if max(v1, v2) < K {
			r += (1 << i)
		} else {
			for j := 0; j < N*2; j++ {
				dp2[i][j] = dp2[i+1][j]
			}
		}
	}
	if r == (1<<61)-1 {
		return -1
	}
	return r + 1
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
