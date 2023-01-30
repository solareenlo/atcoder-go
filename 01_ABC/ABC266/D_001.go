package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 500010

	var n int
	fmt.Fscan(in, &n)

	var b [N][10]int
	for i := 1; i <= n; i++ {
		var t, x, a int
		fmt.Fscan(in, &t, &x, &a)
		b[t][x] += a
	}

	mx := 0
	var dp [N][10]int
	for i := 1; i <= 100000; i++ {
		for j := 0; j <= 4; j++ {
			tmp := 0
			if j != 0 {
				tmp = dp[i-1][j-1]
			}
			dp[i][j] = max(tmp, max(dp[i-1][j], dp[i-1][j+1]))
			if i >= j {
				dp[i][j] += b[i][j]
			}
			mx = max(mx, dp[i][j])
		}
	}
	fmt.Println(mx)

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
