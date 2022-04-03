package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)

	a := make([]int, n)
	dp := [40][200010]int{}

	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	for j := 0; j < n; j++ {
		dp[0][j] = a[j]
	}

	for i := 0; i < 39; i++ {
		for j := 0; j < n; j++ {
			dp[i+1][j] = dp[i][j] + dp[i][(j+dp[i][j])%n]
		}
	}

	ans := 0
	for i := 0; i < 40; i++ {
		if k&1 != 0 {
			ans += dp[i][ans%n]
		}
		k = (k >> 1)
	}

	fmt.Println(ans)
}
