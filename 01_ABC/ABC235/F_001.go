package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N string
	var m int
	fmt.Fscan(in, &N, &m)
	nn := len(N)
	c := make([]int, m)
	for i := range c {
		fmt.Fscan(in, &c[i])
	}
	n := make([]int, nn)
	for i := 0; i < nn; i++ {
		n[i] = int(N[i] - '0')
	}

	type pair struct{ x, y int }
	dp := [10005][1024]pair{}
	k := 0
	v := 0
	const mod = 998244353
	for i := 0; i < nn; i++ {
		for s := 0; s < 1024; s++ {
			for d := 0; d < 10; d++ {
				dp[i+1][s|1<<d].x += dp[i][s].x
				dp[i+1][s|1<<d].y += dp[i][s].y*10 + dp[i][s].x*d
			}
		}
		if i != 0 {
			for d := 0; d < 10; d++ {
				if d != 0 {
					dp[i+1][1<<d].x += 1
					dp[i+1][1<<d].y += d
				}
			}
		}
		for d := 0; d < n[i]; d++ {
			if i != 0 || d != 0 {
				dp[i+1][k|1<<d].x += 1
				dp[i+1][k|1<<d].y += v*10 + d
			}
		}

		for s := 0; s < 1024; s++ {
			dp[i+1][s].x %= mod
			dp[i+1][s].y %= mod
		}
		k |= 1 << n[i]
		v = (v*10 + n[i]) % mod
	}

	p := 0
	for i := 0; i < m; i++ {
		p |= 1 << c[i]
	}

	ans := 0
	for s := 0; s < 1024; s++ {
		if (s & p) == p {
			ans += dp[nn][s].y
			ans %= mod
		}
	}
	if (k & p) == p {
		ans += v
		ans %= mod
	}
	fmt.Println(ans)
}
