package main

import (
	"fmt"
	"strings"
)

func main() {
	var a, b, k int
	fmt.Scan(&a, &b, &k)

	var dp [31][31]int
	dp[0][0] = 1
	for i := 0; i < a+1; i++ {
		for j := 0; j < b+1; j++ {
			if i != 0 {
				dp[i][j] += dp[i-1][j]
			}
			if j != 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}
	for a > 0 && b > 0 {
		if k <= dp[a-1][b] {
			fmt.Print("a")
			a--
		} else {
			fmt.Print("b")
			k -= dp[a-1][b]
			b--
		}
	}
	fmt.Print(strings.Repeat("a", a), strings.Repeat("b", b))
	fmt.Println()
}
