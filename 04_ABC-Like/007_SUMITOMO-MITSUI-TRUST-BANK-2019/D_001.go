package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	var s string
	fmt.Fscan(in, &n, &s)

	dp := [30001][4][1009]bool{}
	dp[0][0][0] = true
	for i := 0; i < n; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 1000; k++ {
				if dp[i][j][k] == false {
					continue
				}
				dp[i+1][j][k] = true
				if j < 3 {
					dp[i+1][j+1][k*10+int(s[i]-'0')] = true
				}
			}
		}
	}

	cnt := 0
	for i := 0; i < 1000; i++ {
		if dp[n][3][i] {
			cnt++
		}
	}
	fmt.Println(cnt)
}
