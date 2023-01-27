package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var s, t string
	fmt.Fscan(in, &s, &t)

	var n int
	fmt.Fscan(in, &n)

	a := make([]string, 101)
	c := make([]int, 101)
	len := len(t)
	for i := 1; i <= n; i++ {
		var c1 string
		fmt.Fscan(in, &c1, &a[i])
		c[i] = int(c1[0] - 'a')
	}

	n++
	c[n] = 26
	a[n] = s

	var dp [101][101][27]int
	for i, _ := range dp {
		for j, _ := range dp[i] {
			for k, _ := range dp[i][j] {
				dp[i][j][k] = 1061109567
			}
		}
	}
	for i := 1; i <= len; i++ {
		dp[i][i][t[i-1]-'a'] = 0
	}

	g := make([]int, 101)
	for i := len; i >= 1; i-- {
		ok := true
		for ok {
			ok = false
			for j := 1; j <= n; j++ {
				for k, _ := range g {
					g[k] = 1061109567
				}
				g[i-1] = 0
				for k := 0; k < utf8.RuneCountInString(a[j]); k++ {
					for l := len; l >= i-1; l-- {
						g[l] = 1061109567
						for x := i - 1; x <= l-1; x++ {
							g[l] = min(g[l], g[x]+dp[x+1][l][a[j][k]-'a'])
						}
					}
				}
				for r := i; r <= len; r++ {
					if dp[i][r][c[j]] > g[r]+1 {
						dp[i][r][c[j]] = g[r] + 1
						ok = true
					}
				}
			}
		}
	}

	ans := dp[1][len][26]
	if ans == 1061109567 {
		ans = 0
	}
	fmt.Println(ans - 1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
