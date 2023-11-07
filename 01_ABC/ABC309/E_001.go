package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 300300

	var p, dp [N]int

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 2; i <= n; i++ {
		fmt.Fscan(in, &p[i])
	}
	for i := 1; i <= m; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		dp[x] = max(dp[x], y+1)
	}
	ans := 0
	for i := 1; i <= n; i++ {
		dp[i] = max(dp[p[i]]-1, dp[i])
		if dp[i] != 0 {
			ans++
		}
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
