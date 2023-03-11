package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 2003
	const mod = 1_000_000_007

	var n int
	fmt.Fscan(in, &n)
	var p [N]int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &p[i])
	}
	var q [N]int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &q[i])
	}
	var catalan [N]int
	catalan[0] = 1
	for i := 1; i <= n; i++ {
		sum := 0
		for j := 0; j < i; j++ {
			sum += catalan[j] * catalan[i-j-1] % mod
		}
		catalan[i] = sum % mod
	}

	var dt [N][N]int
	dt[0][0] = 1
	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			s := 0
			sum := 0
			for i-s > 0 && j-s > 0 && p[i-s-1] == q[j-s-1] {
				s++
				sum -= dt[i-s][j-s] * catalan[s-1] % mod
			}
			if i > 0 {
				sum += dt[i-1][j]
			}
			if j > 0 {
				sum += dt[i][j-1]
			}
			sum = (sum%mod + mod) % mod
			dt[i][j] = sum % mod
		}
	}
	fmt.Println(dt[n][n] % mod)
}
