package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = 1_000_000_007

type P struct{ x, y int }

var (
	x   = [88]int{}
	y   = [88]int{}
	ok  = [2][88][88][88]bool{}
	dp  = [2][88][88][88]int{}
	det = [88][88][88]int{}
	cnt = [88][88][88]int{}
	p2  = [88]int{}
	n   int
)

func cmp1(a, b P) bool {
	if a.x == b.x {
		return a.y < b.y
	}
	return a.x < b.x
}

func cmp2(a, b P) bool {
	return a.x <= b.x
}

func solve(s, i, j, p int) int {
	if j == s {
		if p == 0 && cmp1(P{x[s], y[s]}, P{x[i], y[i]}) {
			return 1
		} else {
			return 0
		}
	}
	if ok[p][s][i][j] {
		return dp[p][s][i][j]
	}
	res := 0
	for k := 0; k < n; k++ {
		q := det[s][j][k]
		if cmp2(P{x[s], y[s]}, P{x[k], y[k]}) && q >= 0 && det[j][k][i] > 0 {
			res += solve(s, j, k, p^(q&1)) * p2[cnt[j][k][s]]
			res %= mod
		}
	}
	ok[p][s][i][j] = true
	dp[p][s][i][j] = res
	return res
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)
	p2[0] = 1
	for i := 1; i < n; i++ {
		p2[i] = p2[i-1] * 2 % mod
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x[i], &y[i])
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				det[i][j][k] = (x[j]-x[i])*(y[k]-y[i]) - (y[j]-y[i])*(x[k]-x[i])
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				for l := 0; l < n; l++ {
					if det[l][i][j] > 0 && det[l][j][k] > 0 && det[l][k][i] > 0 {
						cnt[i][j][k]++
					}
				}
			}
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if cmp1(P{x[i], y[i]}, P{x[j], y[j]}) {
				ans += solve(i, i, j, 0)
				ans %= mod
			}
		}
	}
	fmt.Println(ans)
}
