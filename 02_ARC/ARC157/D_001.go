package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 2005
	const MOD = 998244353

	var s [N][N]int
	var c [N]int

	var n, m int
	fmt.Fscan(in, &n, &m)
	var a [N]string
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		a[i] = " " + a[i]
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			tmp := 0
			if a[i][j] == 'Y' {
				tmp = 1
			}
			s[i][j] = s[i-1][j] + s[i][j-1] - s[i-1][j-1] + tmp
		}
	}

	var sum func(int, int, int, int) int
	sum = func(x, y, xx, yy int) int {
		return s[xx][yy] - s[x-1][yy] - s[xx][y-1] + s[x-1][y-1]
	}

	k := s[n][m]
	ans := 0
	for i := 1; i <= m; i++ {
		if k%i == 0 {
			p := k / i
			l := 0
			t := 1
			now := 1
			for j := 1; j <= m; j++ {
				for t*p < sum(1, 1, n, j) {
					now = now * l % MOD
					l = 0
					t++
				}
				if t*p == sum(1, 1, n, j) {
					c[t] = j
					l++
				}
			}
			now = now * min(1, l) % MOD
			c[t] = m
			if now == 0 {
				continue
			}
			lst := 1
			l = 0
			for j := 1; j <= n; j++ {
				mx := 0
				mn := 1 << 30
				for o := 1; o <= t; o++ {
					mx = max(mx, sum(lst, c[o-1]+1, j, c[o]))
					mn = min(mn, sum(lst, c[o-1]+1, j, c[o]))
				}
				for mx > 2 {
					now = now * l % MOD
					mx -= 2
					mn -= 2
					lst = j
					l = 0
				}
				if mn == mx && mx == 2 {
					l++
				}
			}
			ans = (ans + now*min(1, l)) % MOD
		}
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
