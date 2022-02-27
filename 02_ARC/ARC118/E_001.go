package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = 998244353

func Add(a *int, b int) {
	*a += b
	if *a >= mod {
		*a -= mod
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	const maxn = 205

	var n int
	fmt.Fscan(in, &n)

	m := 0
	a := make([]int, n+1)
	bad := [maxn]bool{}
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		if a[i] >= 0 {
			bad[a[i]] = true
		} else {
			m++
		}
	}

	fac := make([]int, m+1)
	fac[0] = 1
	for i := 1; i <= m; i++ {
		fac[i] = fac[i-1] * i % mod
	}

	dp := [maxn][maxn][maxn][4]int{}
	dp[0][0][0][3] = 1
	for i := 0; i <= n+1; i++ {
		for j := 0; j <= n+1; j++ {
			for k := 0; k <= m; k++ {
				for f := 1; f < 4; f++ {
					cur := dp[i][j][k][f]
					if cur == 0 {
						continue
					}
					Add(&dp[i+1][j][k][f|1], cur)
					Add(&dp[i][j+1][k][f|2], cur)
					if 1 <= i && i <= n && 1 <= j && j <= n {
						if a[i] < 0 && !bad[j] && f == 3 {
							Add(&dp[i+1][j][k+1][1], mod-cur)
							Add(&dp[i][j+1][k+1][2], mod-cur)
						}
						if a[i] >= 0 && a[i] == j {
							Add(&dp[i+1][j][k][1], mod-cur)
							Add(&dp[i][j+1][k][2], mod-cur)
						}
					}
				}
			}
		}
	}

	ans := 0
	for i := 0; i <= m; i++ {
		ans = (ans + dp[n+1][n+1][i][3]*fac[m-i]) % mod
	}
	fmt.Println(ans)
}
