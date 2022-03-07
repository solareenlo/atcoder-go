package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, d int
	fmt.Fscan(in, &n, &d)

	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	const mod = 998244353
	dp := [505][2050]int{}
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j < (1 << (d*2 + 1)); j++ {
			if dp[i-1][j] != 0 {
				nj := j >> 1
				for k := i - d; k <= i+d; k++ {
					if k >= 1 && k <= n && (a[i] == -1 || a[i] == k) {
						p := k - i + d
						if (nj>>p)&1 != 0 {
							continue
						}
						dp[i][nj|(1<<p)] += dp[i-1][j]
						dp[i][nj|(1<<p)] %= mod
					}
				}
			}
		}
	}

	for i := 0; i < (1 << (d*2 + 1)); i++ {
		if dp[n][i] != 0 {
			fmt.Println(dp[n][i])
			return
		}
	}
	fmt.Println(0)
}
