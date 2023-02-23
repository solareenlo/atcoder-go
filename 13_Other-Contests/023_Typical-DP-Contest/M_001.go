package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 1000000007

type P struct {
	x, y int
}

var h, r int
var M [][]int
var g [16][16]bool
var dp [1 << 16][16][16]int

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &h, &r)
	M = make([][]int, r)
	for i := range M {
		M[i] = make([]int, r)
	}
	for i := 0; i < r; i++ {
		for j := 0; j < r; j++ {
			fmt.Fscan(in, &g[i][j])
		}
	}
	for j := 0; j < r; j++ {
		dp[1<<j][j][j] = 1
	}
	for i := 0; i < 1<<r; i++ {
		for j := 0; j < r; j++ {
			for k := 0; k < r; k++ {
				if dp[i][j][k] == 0 {
					continue
				}
				for t := 0; t < r; t++ {
					if (i>>t)&1 != 0 || !g[k][t] {
						continue
					}
					dp[i|1<<t][j][t] += dp[i][j][k]
					dp[i|1<<t][j][t] %= MOD
				}
				M[j][k] += dp[i][j][k]
				M[j][k] %= MOD
			}
		}
	}
	M = ppow(M, h)
	fmt.Println(M[0][0])
}

func ppow(a [][]int, b int) [][]int {
	res := make([][]int, r)
	for i := range res {
		res[i] = make([]int, r)
	}
	for i := 0; i < r; i++ {
		res[i][i] = 1
	}
	for b != 0 {
		if (b & 1) != 0 {
			res = mul(res, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return res
}

func mul(a, b [][]int) [][]int {
	c := make([][]int, r)
	for i := range c {
		c[i] = make([]int, r)
	}
	for i := 0; i < r; i++ {
		for j := 0; j < r; j++ {
			for k := 0; k < r; k++ {
				c[i][j] += a[i][k] * b[k][j]
				c[i][j] %= MOD
			}
		}
	}
	return c
}
