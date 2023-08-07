package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var dp [5050]int
	var a [5050][5050]bool
	for i := 0; i < 5050; i++ {
		dp[i] = int(1e9)
		for j := 0; j < 5050; j++ {
			a[i][j] = false
		}
	}
	dp[0] = 0
	var n int
	var s string
	fmt.Fscan(in, &n, &s)
	var c [5050]int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &c[i+1])
	}
	for i := 0; i < n; i++ {
		l := i
		r := i
		for 0 <= l && r < n && s[l] == s[r] {
			a[l][r] = true
			l--
			r++
		}
		l = i
		r = i + 1
		for 0 <= l && r < n && s[l] == s[r] {
			a[l][r] = true
			l--
			r++
		}
	}
	for i := 0; i < n; i++ {
		for j := 1; j < n-i+1; j++ {
			if a[i][i+j-1] {
				dp[i+j] = min(dp[i+j], dp[i]+c[j])
			}
		}
	}
	fmt.Println(dp[n])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
