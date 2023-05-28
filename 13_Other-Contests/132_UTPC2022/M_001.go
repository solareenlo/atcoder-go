package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const mod = 998244353

	var n int
	fmt.Fscan(in, &n)
	ans := 0
	v := make([]int, n)
	for i := range v {
		fmt.Fscan(in, &v[i])
		ans += v[i]
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			ans += F(v[i] + v[j])
		}
	}
	ans %= mod
	var dp [5][2]int
	dp[0][0] = 1
	for _, i := range v {
		b := i & 1
		for j := 4; j > 0; j-- {
			dp[j][b] += dp[j-1][0]
			dp[j][b] %= mod
			tmp := 0
			if b == 0 {
				tmp = 1
			}
			dp[j][tmp] += dp[j-1][1]
			dp[j][tmp] %= mod
		}
		dp[3][0] += dp[4][0]
		dp[3][1] += dp[4][1]
		dp[3][0] %= mod
		dp[3][1] %= mod
		dp[4][0] = 0
		dp[4][1] = 0
	}
	ans += dp[3][1]
	ans %= mod
	fmt.Println(ans)
}

func F(n int) int {
	ret := 0
	for (n & 1) != 0 {
		n >>= 1
		ret = ret*2 + 1
	}
	return ret
}
