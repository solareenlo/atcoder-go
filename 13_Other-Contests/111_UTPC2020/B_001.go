package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 998244353
const N = 10100

var n, k int
var dp, dp2 [N]int
var dp3 [N][2]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var s string
	fmt.Fscan(in, &s)
	fmt.Fscan(in, &k)
	n = len(s)
	dp[0] = 1
	dp2[0] = 1
	dp3[0][0] = 1
	for i := 1; i < n+k+1; i++ {
		dp[i] = dp[i-1]
		dp2[i] = dp2[i-1]
		if i <= k {
			dp2[i] = (dp2[i] + dp2[i-1]) % MOD
		}
		for j := 0; j < 2; j++ {
			dp3[i][j] = dp3[i-1][j]
		}
		if i <= k {
			dp3[i][0] = (dp3[i][0] + dp3[i-1][0]) % MOD
		}
		if i >= n {
			dp3[i][1] = (dp3[i][1] + dp3[i-1][0] + dp3[i-1][1]) % MOD
		}
		for j := 1; j <= k && j < i; j++ {
			dp[i] = (dp[i] + dp[j-1]*dp[i-j-1]) % MOD
			dp2[i] = (dp2[i] + dp[j-1]*dp2[i-j-1]) % MOD
			for z := 0; z < 2; z++ {
				dp3[i][z] = (dp3[i][z] + dp[j-1]*dp3[i-j-1][z]) % MOD
			}
		}
	}
	var g [N]int
	for i := 0; i < n; i++ {
		if s[i] == 'R' {
			g[i] = 0
		} else {
			if s[i] == 'S' {
				g[i] = 1
			} else {
				g[i] = 2
			}
		}
	}
	var a [N]int
	for i := 1; i < n; i++ {
		a[i] = (g[i] - g[i-1]) % 3
		if a[i] == -2 {
			a[i] = 1
		}
		if a[i] == 2 {
			a[i] = -1
		}
	}
	pre := 0
	cur := -1
	ans := 1
	for i := 1; i < n; i++ {
		if a[i] == 1 {
			cur = i
			ans = ans * get(pre, cur) % MOD
			pre = cur
		} else if a[i] == -1 {
			cur = i + k
			ans = ans * get(pre, cur) % MOD
			pre = cur
		}
	}
	ans = ans * get(pre, n+k) % MOD
	fmt.Println(ans)
}

func get(l, r int) int {
	if l == 0 && r == n+k {
		if r-l-1 < 0 {
			return 0
		}
		return (dp3[r-l-1][0] + dp3[r-l-1][1]) % MOD
	}
	if l == 0 || r == n+k {
		if r-l-1 < 0 {
			return 0
		}
		return dp2[r-l-1]
	} else {
		if r-l-1 < 0 {
			return 0
		}
		return dp[r-l-1]
	}
}
