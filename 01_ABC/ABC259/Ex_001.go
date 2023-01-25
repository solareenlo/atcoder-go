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

	const mod = 998244353
	const N = 410

	var c [N << 1][N << 1]int
	var a [N][N]int
	for i := 0; i < 2*N; i++ {
		c[i][0] = 1
		c[i][i] = 1
		for j := 1; j < i; j++ {
			c[i][j] = (c[i-1][j] + c[i-1][j-1]) % mod
		}
	}

	var n int
	fmt.Fscan(in, &n)

	v := make([][]pair, N*N)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Fscan(in, &a[i][j])
			v[a[i][j]] = append(v[a[i][j]], pair{i, j})
		}
	}

	ans := 0
	for i := 1; i <= n*n; i++ {
		l := len(v[i])
		if l <= n {
			ans = (ans + l) % mod
			for j := 0; j < l; j++ {
				for k := j + 1; k < l; k++ {
					x := v[i][k].x - v[i][j].x
					y := v[i][k].y - v[i][j].y
					if x >= 0 && y >= 0 {
						ans = (ans + c[x+y][x]) % mod
					}
				}
			}
		} else {
			var dp [N][N]int
			for j := 1; j <= n; j++ {
				for k := 1; k <= n; k++ {
					dp[j][k] = dp[j-1][k]
					dp[j][k] = (dp[j][k] + dp[j][k-1]) % mod
					if a[j][k] == i {
						dp[j][k] = (dp[j][k] + 1) % mod
						ans = (ans + dp[j][k]) % mod
					}
				}
			}
		}
	}
	fmt.Println(ans)
}
