package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 998244353

	var v [2020]int
	var dp [2020][2020]int
	var used [4116]bool

	ans := 0
	dp[0][0] = 1

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		used[a] = true
	}
	a := 0
	for i := 0; i < m; i++ {
		for used[a] {
			a++
		}
		v[i] = a
		used[a] = true
	}
	for i := 0; i < m; i++ {
		for j := 0; j < i+1; j++ {
			dp[i+1][j] += dp[i][j] * v[j] % MOD
			dp[i+1][j] %= MOD
			dp[i+1][j+1] += dp[i][j]
			dp[i+1][j+1] %= MOD
		}
	}
	for j := 0; j < m+1; j++ {
		ans += dp[m][j]
		ans %= MOD
	}
	fmt.Println(ans)
}
