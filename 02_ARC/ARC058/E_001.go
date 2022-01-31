package main

import "fmt"

func main() {
	var n, x, y, z int
	fmt.Scan(&n, &x, &y, &z)

	sum := x + y + z
	g := (((1 << y) + 1) << z) + 1
	ans := 0
	mod := int(1e9 + 7)
	dp := [2][1 << 20]int{}
	dp[0][0] = 1
	for i := 0; i <= n; i++ {
		ans *= 10
		ans %= mod
		for bit := 0; bit < (1 << sum); bit++ {
			if (bit & g) == g {
				ans += dp[i&1][bit]
				ans %= mod
			} else {
				for j := 1; j <= 10; j++ {
					dp[(i&1)^1][(bit+(1<<sum))>>j] += dp[i&1][bit]
					dp[(i&1)^1][(bit+(1<<sum))>>j] %= mod
				}
			}
			dp[i&1][bit] = 0
		}
	}
	fmt.Println(ans)
}
