package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const mod = 1000000007

	var n, k int
	fmt.Fscan(in, &n, &k)

	var dp [2 << 17]int
	dp[k] = 1
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		for j := k; j >= 0; j-- {
			dp[j] += dp[j+1]
		}
		for j := 0; j < k; j++ {
			dp[j] = (dp[j] - dp[j+a+1]) % mod
		}
	}
	fmt.Println(dp[0])
}
