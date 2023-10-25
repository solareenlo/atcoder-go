package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var dp [333][333]int

	var n int
	fmt.Fscan(in, &n)
	dp[0][0] = 1
	ans := 0
	for i := 0; i < n+30; i++ {
		x := 0
		if i < n {
			fmt.Fscan(in, &x)
		}
		for j := 0; j < 333; j++ {
			for k := 0; k < x+1; k++ {
				if ((j + k) % 2) != 0 {
					continue
				}
				dp[i+1][(j+k)/2] = (dp[i+1][(j+k)/2] + dp[i][j]) % 1000000007
			}
		}
		ans = (ans + dp[i+1][1] + min(x, 1)) % 1000000007
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
