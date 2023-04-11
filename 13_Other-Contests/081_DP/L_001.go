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
	var a [30030]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	var dp [30030]int
	for i := 1; i <= n; i++ {
		for j := 1; j+i-1 <= n; j++ {
			dp[j] = max(a[j]-dp[j+1], a[j+i-1]-dp[j])
		}
	}
	fmt.Println(dp[1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
