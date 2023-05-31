package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)
	A := make([][]int, N)
	for i := range A {
		A[i] = make([]int, N)
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Fscan(in, &A[i][j])
			if 0 < A[i][j] {
				A[i][j] -= 1
			}
		}
	}
	var C [4][4][16]int
	for t := 0; t < 4; t++ {
		for u := 0; u < 4; u++ {
			for i := t; i < N; i += 4 {
				for j := u; j < N; j += 4 {
					if A[i][j] != -1 {
						C[t][u][A[i][j]] = 1
					}
				}
			}
		}
	}
	var P, Q [16]int
	for i := 0; i < N; i++ {
		var T [16]int
		for j := range T {
			T[j] = -1
		}
		for j := 0; j < N; j++ {
			if A[i][j] == -1 {
				continue
			}
			if T[A[i][j]] != -1 && T[A[i][j]] != j%4 {
				fmt.Println("No")
				return
			}
			T[A[i][j]] = j % 4
		}
	}
	for j := 0; j < N; j++ {
		var T [16]int
		for i := range T {
			T[i] = -1
		}
		for i := 0; i < N; i++ {
			if A[i][j] == -1 {
				continue
			}
			if T[A[i][j]] != -1 && T[A[i][j]] != i%4 {
				fmt.Println("No")
				return
			}
			T[A[i][j]] = i % 4
		}
	}
	for t := 0; t < (1 << 16); t++ {
		var F, G [16]int
		for i := 0; i < 16; i++ {
			F[i] = -1
			G[i] = -1
		}
		var x, y [4]int
		flg := true
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				b := ((t >> (4*i + j)) & 1)
				if b == 1 {
					x[i]++
				} else {
					y[j]++
				}
				for k := 0; k < 16; k++ {
					if C[i][j][k] != 0 {
						tmp := j
						if b != 0 {
							tmp = i
						}
						if F[k] != -1 && (F[k] != b || G[k] != tmp) {
							flg = false
						}
						F[k] = b
						if b != 0 {
							G[k] = i
						} else {
							G[k] = j
						}
					}
				}
			}
		}
		for s := 0; s < 16; s++ {
			if F[s] == 1 {
				if P[s] != 0 {
					flg = false
				}
				x[G[s]]--
			}
			if F[s] == 0 {
				if Q[s] != 0 {
					flg = false
				}
				y[G[s]]--
			}
		}
		for i := 0; i < 4; i++ {
			if x[i] < 0 {
				flg = false
			}
			if y[i] < 0 {
				flg = false
			}
		}
		if flg {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
