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

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	const mod = 998244353
	ans := 0
	for i := 1; i <= n; i++ {
		var dp [105][105]int
		dp[0][0] = 1
		for _, x := range a {
			for j := i; j >= 1; j-- {
				for k := 0; k < i; k++ {
					dp[j][(k+x)%i] += dp[j-1][k]
					dp[j][(k+x)%i] %= mod
				}
			}
		}
		ans += dp[i][0]
		ans %= mod
	}
	fmt.Println(ans)
}
