package main

import "fmt"

func main() {
	var n int
	var a, b string
	fmt.Scan(&n, &a, &b)
	a = " " + a
	b = " " + b

	dp := [200010][10]int{}
	dp[n][0] = 1
	for i := n; i >= 1; i-- {
		for j := 0; j < 7; j++ {
			if b[i] == 'T' {
				dp[i-1][j] = dp[i][(j*10+int(a[i]-'0'))%7] | dp[i][j*10%7]
			} else {
				dp[i-1][j] = dp[i][(j*10+int(a[i]-'0'))%7] & dp[i][j*10%7]
			}
		}
	}

	if dp[0][0] != 0 {
		fmt.Println("Takahashi ")
	} else {
		fmt.Println("Aoki")
	}
}
