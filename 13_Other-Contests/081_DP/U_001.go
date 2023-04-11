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
	var a [20][20]int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &a[i][j])
		}
	}
	var dp [1 << 22]int
	for i := 0; i < (1 << n); i++ {
		for j := 0; j < n; j++ {
			if ((i >> j) & 1) != 0 {
				for k := j + 1; k < n; k++ {
					if ((i >> k) & 1) != 0 {
						dp[i] += a[j][k]
					}
				}
			}
		}
		for j := i; j > 0; j = i & (j - 1) {
			dp[i] = max(dp[i], dp[j]+dp[i^j])
		}
	}
	fmt.Println(dp[(1<<n)-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
