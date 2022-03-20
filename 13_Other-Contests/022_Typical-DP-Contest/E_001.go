package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var d int
	var n string
	fmt.Fscan(in, &d, &n)

	const mod = 1_000_000_007
	dp := [10005][105]int{}
	sum := 0
	for i := 0; i < len(n); i++ {
		for j := 0; j < int(n[i]-'0'); j++ {
			dp[i+1][(sum+j)%d] += 1
		}
		sum += int(n[i] - '0')
		for j := 0; j < d; j++ {
			for k, s := 0, j; k < 10; k++ {
				dp[i+1][s] += dp[i][j]
				dp[i+1][s] %= mod
				s = (s + 1) % d
			}
		}
	}

	fmt.Println((dp[len(n)][0] + mod - 1) % mod)
}
