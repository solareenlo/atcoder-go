package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, q int
	fmt.Fscan(in, &n, &q)

	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	dp := [4001][4001]int{}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if a[i] > a[j] {
				dp[i][j] = 1
			}
		}
	}

	const mod = 1_000_000_007
	inv2 := (mod + 1) / 2
	for k := 1; k <= q; k++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		dp[y][x] = inv2 * (dp[x][y] + dp[y][x]) % mod
		dp[x][y] = dp[y][x]
		for i := 1; i <= n; i++ {
			if i != x && i != y {
				dp[i][y] = inv2 * (dp[i][x] + dp[i][y]) % mod
				dp[i][x] = dp[i][y]
				dp[y][i] = inv2 * (dp[x][i] + dp[y][i]) % mod
				dp[x][i] = dp[y][i]
			}
		}
	}

	ans := 0
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			ans += dp[i][j]
		}
	}
	for i := 0; i < q; i++ {
		ans = ans * 2 % mod
	}
	fmt.Println(ans)
}
