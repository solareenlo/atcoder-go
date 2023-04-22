package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	var s [5][]string
	for i := 0; i < 5; i++ {
		var t string
		fmt.Fscan(in, &t)
		s[i] = strings.Split(t, "")
	}

	var dp [1005][3]int
	for i := 0; i < n+1; i++ {
		for j := 0; j < 3; j++ {
			dp[i][j] = int(1e18)
		}
	}
	dp[0][0] = 0
	dp[0][1] = 0
	dp[0][2] = 0
	for i := 0; i < n; i++ {
		red := 0
		blue := 0
		white := 0
		for j := 0; j < 5; j++ {
			if s[j][i] == "R" {
				red++
			} else if s[j][i] == "B" {
				blue++
			} else if s[j][i] == "W" {
				white++
			}
		}
		dp[i+1][0] = min(dp[i+1][0], dp[i][1]+(5-blue), dp[i][2]+(5-white))
		dp[i+1][1] = min(dp[i+1][1], dp[i][0]+(5-red), dp[i][2]+(5-white))
		dp[i+1][2] = min(dp[i+1][2], dp[i][0]+(5-red), dp[i][1]+(5-blue))
	}
	fmt.Println(min(dp[n][0], dp[n][1], dp[n][2]))
}

func min(a ...int) int {
	res := a[0]
	for i := range a {
		if res > a[i] {
			res = a[i]
		}
	}
	return res
}
