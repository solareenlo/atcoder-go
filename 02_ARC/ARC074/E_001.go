package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	type pair struct{ l, x int }
	p := make([][]pair, 303)
	for i := 0; i < m; i++ {
		var l, r, x int
		fmt.Fscan(in, &l, &r, &x)
		p[r-1] = append(p[r-1], pair{l, x})
	}

	mod := int(1e9 + 7)
	dp := [303][303]int{}
	dp[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			for k := 0; k < i; k++ {
				dp[i][j] += dp[max(j, k)][min(j, k)]
				dp[i][j] %= mod
			}
		}
		for _, t := range p[i] {
			l := t.l
			x := t.x
			for j := 0; j < i+1; j++ {
				for k := 0; k < j+1; k++ {
					tmp1, tmp2 := 0, 0
					if j < l {
						tmp1 = 1
					}
					if k < l {
						tmp2 = 1
					}
					if x+tmp1+tmp2-3 != 0 {
						dp[j][k] = 0
					}
				}
			}
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			ans += 6 * dp[i][j]
			ans %= mod
		}
	}
	fmt.Println((ans + dp[0][0]*3) % mod)
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
