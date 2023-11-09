package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var dp [200005]int

	var n int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			var x int
			fmt.Fscan(in, &x)
			t := (1 << i) | (1 << j)
			for k := 0; k < (1 << n); k++ {
				if (k & t) == 0 {
					dp[k|t] = max(dp[k|t], dp[k]+x)
				}
			}
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
