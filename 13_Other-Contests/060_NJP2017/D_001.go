package main

import "fmt"

func main() {
	var N, M, K int
	fmt.Scan(&N, &M, &K)
	swp := false
	if N > M {
		N, M = M, N
		swp = true
	}
	var C [100][100]int
	if K <= N*M*(M-1)/2 {
		for i := 0; i < N; i++ {
			for j := 0; j < M; j++ {
				C[i][j] = i*M + j + 1
			}
		}
	} else {
		for i := 0; i < N; i++ {
			for j := 0; j < M; j++ {
				C[N-1-i][j] = i*M + j + 1
			}
		}
		K -= M * N * (N - 1) / 2
	}
	for i := 0; i < N; i++ {
		for j := 0; j < M-1; j++ {
			if K < M-j-1 {
				continue
			}
			tmp := C[i][M-1]
			for k := M - 1; k > j; k-- {
				C[i][k] = C[i][k-1]
			}
			C[i][j] = tmp
			K -= M - j - 1
		}
	}
	if swp {
		for i := 0; i < M; i++ {
			for j := i + 1; j < M; j++ {
				C[i][j], C[j][i] = C[j][i], C[i][j]
			}
		}
		N, M = M, N
	}
	for i := 0; i < N; i++ {
		fmt.Print(C[i][0])
		for j := 1; j < M; j++ {
			fmt.Print(" ", C[i][j])
		}
		fmt.Println()
	}
}
