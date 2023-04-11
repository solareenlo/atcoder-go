package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var s, t string
	fmt.Fscan(in, &s, &t)
	m := len(s)
	n := len(t)
	s = " " + s
	t = " " + t
	var dp [3005][3005]int
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i] == t[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	var z string
	for m > 0 && n > 0 {
		if s[m] == t[n] {
			z += string(s[m])
			m--
			n--
		} else if dp[m-1][n] > dp[m][n-1] {
			m--
		} else {
			n--
		}
	}
	fmt.Println(reverseString(z))
}

func reverseString(s string) string {
	res := []rune(s)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
