package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 505
	const mod = 1000000007

	g := make([]int, N)
	n := 0
	for i := 1; i <= 26; i++ {
		var x int
		fmt.Fscan(in, &x)
		if x != 0 {
			n++
			g[n] = x
		}
	}
	var f [N][N]int
	f[1][g[1]-1] = 1
	s := make([]int, N)
	for i := 1; i <= n; i++ {
		s[i] = s[i-1] + g[i]
	}
	var C [N][N]int
	for i := 0; i <= 300; i++ {
		C[i][0] = 1
		for j := 1; j <= i; j++ {
			C[i][j] = (C[i-1][j] + C[i-1][j-1]) % mod
		}
	}
	for i := 2; i <= n; i++ {
		for j := 0; j <= s[i]; j++ {
			if f[i-1][j] != 0 {
				for k := 1; k <= g[i]; k++ {
					for l := 0; l <= min(j, k); l++ {
						f[i][j-l+g[i]-k] += f[i-1][j] * C[g[i]-1][k-1] % mod * C[j][l] % mod * C[s[i-1]+1-j][k-l] % mod
						f[i][j-l+g[i]-k] %= mod
					}
				}
			}
		}
	}
	fmt.Println(f[n][0])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
