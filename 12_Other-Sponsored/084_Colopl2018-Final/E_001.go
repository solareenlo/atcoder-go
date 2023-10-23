package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 1000000007
	const N = 45

	var n, m int
	fmt.Fscan(in, &n, &m)
	s := 0
	a := make([]int, m+1)
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &a[i])
		s += a[i]
	}
	sort.Ints(a[1:])
	var js [N]int
	js[0] = 1
	js[1] = 1
	for i := 2; i <= n; i++ {
		js[i] = js[i-1] * i % MOD
	}
	var inv [N]int
	inv[0] = 1
	inv[1] = 1
	for i := 2; i <= n; i++ {
		inv[i] = (MOD - MOD/i) * inv[MOD%i] % MOD
	}
	for i := 2; i <= n; i++ {
		inv[i] = inv[i] * inv[i-1] % MOD
	}
	var v [N][N * N]int
	for i := 0; i < n+1; i++ {
		for j := 0; j < s+1; j++ {
			v[i][j] = j + i*m
			pre := 0
			for x := 1; x <= m; x++ {
				pre += a[x]
				v[i][j] = min(v[i][j], pre+j+i*(m-x))
			}
		}
	}
	var f [N][N][N * N]int
	for i := 0; i < n+1; i++ {
		if v[i][0] >= s {
			f[0][i][s] = inv[n-i]
		}
	}
	for i := 0; i <= m-1; i++ {
		for p := 0; p <= n; p++ {
			for r := 0; r <= s; r++ {
				if f[i][p][r] != 0 {
					for q := 0; q <= min(p, r/(i+1)); q++ {
						if v[p-q][s-r+(i+1)*q] >= s {
							f[i+1][p-q][r-(i+1)*q] = (f[i+1][p-q][r-(i+1)*q] + f[i][p][r]*inv[q]) % MOD
						}
					}
				}
			}
		}
	}
	fmt.Println(f[m][0][0] * js[n] % MOD)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
