package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const mod = 1000000009

	var N, T, M int
	fmt.Fscan(in, &N, &T, &M)

	var dp [105][10005]int
	dp[0][0] = 1
	for i := 0; i < N; i++ {
		var x int
		fmt.Fscan(in, &x)
		for j := i + 1; j >= 1; j-- {
			for k := x; k <= T; k++ {
				dp[j][k] += dp[j-1][k-x]
				if dp[j][k] >= mod {
					dp[j][k] -= mod
				}
			}
		}
	}

	ans := 0
	for j := M; j <= N; j++ {
		for k := 0; k <= T; k++ {
			ans += dp[j][k]
			if ans >= mod {
				ans -= mod
			}
		}
	}
	fmt.Println(ans)
}
