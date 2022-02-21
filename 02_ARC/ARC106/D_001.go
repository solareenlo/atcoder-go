package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	C := [310][310]int{}
	for i := 0; i <= m; i++ {
		C[i][0] = 1
	}

	const mod = 998244353
	const inv2 = 499122177
	for i := 1; i <= m; i++ {
		for j := 1; j <= i; j++ {
			C[i][j] = (C[i-1][j] + C[i-1][j-1]) % mod
		}
	}

	s := make([]int, m+1)
	for i := 1; i <= n; i++ {
		k := 1
		for j := 0; j <= m; j++ {
			s[j] += k
			s[j] %= mod
			k = k * a[i] % mod
		}
	}

	dp := make([]int, m+1)
	for i := 1; i <= m; i++ {
		for j := 0; j <= i; j++ {
			dp[i] += s[j] * s[i-j] % mod * C[i][j] % mod
			dp[i] %= mod
		}
	}

	bin := 2
	for i := 1; i <= m; i++ {
		dp[i] += mod - bin*s[i]%mod
		dp[i] %= mod
		dp[i] = dp[i] * inv2 % mod
		bin <<= 1
		bin %= mod
	}

	for i := 1; i <= m; i++ {
		fmt.Fprintln(out, dp[i])
	}
}
