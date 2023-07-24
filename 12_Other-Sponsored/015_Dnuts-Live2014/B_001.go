package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var s string
	fmt.Fscan(in, &s)

	n := len(s)

	x := int(s[0] - '0')

	var dp [20][7][2]int
	for i := 1; i < x; i++ {
		dp[1][i%7][0] += 1
	}
	dp[1][x%7][1] += 1

	for i := 1; i < n; i++ {
		x := int(s[i] - '0')
		for j := 0; j < 7; j++ {
			for k := 1; k <= 7; k++ {
				dp[i+1][(10*j+k)%7][0] += dp[i][j][0]
			}
			for k := 1; k < x; k++ {
				dp[i+1][(10*j+k)%7][0] += dp[i][j][1]
			}
			dp[i+1][(10*j+x)%7][1] += dp[i][j][1]
		}
		for j := 1; j <= 7; j++ {
			dp[i+1][j%7][0] += 1
		}
	}

	fmt.Println(dp[n][0][0] + dp[n][0][1])
}
