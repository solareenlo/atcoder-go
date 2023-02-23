package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 1010

	var n int
	fmt.Fscan(in, &n)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	pref := make([][]int, n+1)
	for i := range pref {
		pref[i] = make([]int, n+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Fscan(in, &f[i][j])
			pref[i][j] = pref[i][j-1] + f[i][j]
		}
	}

	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		mx := -int(2e9)
		for j := 1; j <= i; j++ {
			mx = max(mx, dp[j])
			dp[j] = pref[i][i] - pref[i][j-1] + mx
		}
	}

	res := 0
	for i := 0; i < n+1; i++ {
		res = max(res, dp[i])
	}
	fmt.Println(2 * res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
