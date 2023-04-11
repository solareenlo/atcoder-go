package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, w int
	fmt.Fscan(in, &n, &w)
	var dp [100010]int
	for i := 1; i <= n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		for j := w; j >= x; j-- {
			dp[j] = max(dp[j], dp[j-x]+y)
		}
	}
	fmt.Println(dp[w])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
