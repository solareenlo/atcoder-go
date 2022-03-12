package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	dp := [2][1 << 20]int{}
	for i := 0; i < 1<<20; i++ {
		dp[0][i] = 1
	}

	const mod = 1_000_000_007
	b := 0
	z := make([]int, 500005)
	pr := make([]int, 1<<20)
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		b ^= a
		if b != 0 {
			dp[0][b] += (z[i] - z[pr[b]]) * dp[1][b]
			dp[0][b] %= mod
			dp[1][b] += dp[0][b]
			dp[1][b] %= mod
			pr[b] = i
		} else {
			z[i]++
			if dp[1][0] != 0 {
				dp[1][0] *= 2
				dp[1][0] %= mod
			} else {
				dp[1][0]++
			}
		}
		z[i+1] += z[i]
	}

	ans := 0
	if b != 0 {
		ans = dp[0][b]
	} else {
		for i := 0; i < 1<<20; i++ {
			ans += dp[1][i]
			ans %= mod
		}
	}

	fmt.Println(ans)
}
