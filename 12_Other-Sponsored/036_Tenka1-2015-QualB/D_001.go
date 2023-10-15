package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	hen := "ODIZEhsqLBG"
	to := []string{"0", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	var d, n int
	fmt.Fscan(in, &d, &n)

	s := make([]string, 10000)
	for i := 0; i < n; i++ {
		var A string
		fmt.Fscan(in, &A)
		A = reverseString(A)
		B := ""
		for c := range A {
			for j := 0; j < 11; j++ {
				if A[c] == hen[j] {
					B += to[j]
				}
			}
		}
		s[i] = B
	}

	tmp := s[:n]
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i]+tmp[j] < tmp[j]+tmp[i]
	})

	dp := make([]string, 201)
	for i := 0; i < n; i++ {
		for j := d; j >= len(s[i]); j-- {
			k := j - len(s[i])
			if len(dp[k]) == k {
				dp[j] = max(dp[j], s[i]+dp[k])
			}
		}
	}

	ans := ""
	for i := 1; i <= d; i++ {
		if len(dp[i]) == 0 {
			continue
		}
		if len(ans) == 0 {
			ans = dp[i]
			continue
		}
		if dp[i][0] != '0' {
			ans = dp[i]
		} else if ans[0] == '0' && dp[i] > ans {
			ans = dp[i]
		}
	}

	if ans[0] == '0' {
		for len(ans) > 1 && ans[len(ans)-1] == '0' {
			ans = ans[:len(ans)-1]
		}
		if len(ans) > 1 {
			ans = "0." + ans[1:]
		}
	}

	fmt.Println(ans)
}

func reverseString(s string) string {
	res := []rune(s)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

func max(a, b string) string {
	if len(a) > len(b) || (len(a) == len(b)) && a > b {
		return a
	}
	return b
}
