package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	var a [55]int
	sum := 0
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		sum += a[i]
	}
	if sum < m {
		fmt.Println(-1)
		return
	}
	sum -= m
	var dp [550000]int
	for i := 1; i <= n; i++ {
		for j := sum; j >= a[i]; j-- {
			dp[j] = max(dp[j], dp[j-a[i]]+a[i])
		}
	}
	fmt.Println(sum + m - dp[sum])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
