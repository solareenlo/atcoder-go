package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const mod = 100000
	const N = 3000010

	var K, R int
	fmt.Fscan(in, &K, &R)

	latte := make([]bool, N)
	malta := make([]bool, N)
	var dp [N][2][2]int
	for i := 0; i < R; i++ {
		var t, a int
		fmt.Fscan(in, &t, &a)
		if t == 1 {
			latte[a] = true
		} else {
			malta[a] = true
		}
	}

	dp[0][0][0] = 1
	for i := 1; i <= K; i++ {
		if latte[i] == false {
			dp[i][0][1] = (dp[i-1][0][0] + dp[i-1][1][0]) % mod
		}
		if malta[i] == false {
			dp[i][1][0] = (dp[i-1][0][0] + dp[i-1][0][1]) % mod
		}
		if latte[i] == false && malta[i] == false {
			dp[i][0][0] = (dp[i-1][0][0] + dp[i-1][0][1] + dp[i-1][1][0]) % mod
		}
	}
	fmt.Println(dp[K][0][0])
}
