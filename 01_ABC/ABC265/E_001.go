package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct {
	x, y int
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	var a, b, c, d, e, f int
	fmt.Fscan(in, &a, &b, &c, &d, &e, &f)
	poi := make(map[pair]bool)
	for i := 1; i <= m; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		poi[pair{x, y}] = true
	}

	const mod = 998244353
	ans := 0
	var dp [305][305][305]int
	dp[0][0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j <= i; j++ {
			for k, p := 0, i-j; p != -1; k, p = k+1, p-1 {
				nowx := j*a + k*c + p*e
				nowy := j*b + k*d + p*f
				_, ok := poi[pair{nowx, nowy}]
				if ok {
					continue
				}
				tmp1, tmp2 := 0, 0
				if j != 0 {
					tmp1 = dp[i-1][j-1][k]
				}
				if k != 0 {
					tmp2 = dp[i-1][j][k-1]
				}
				dp[i][j][k] = ((dp[i-1][j][k]+tmp1)%mod + tmp2) % mod
				if i == n {
					ans = (ans + dp[n][j][k]) % mod
				}
			}
		}
	}
	fmt.Println(ans)
}
