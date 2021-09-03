package main

import "fmt"

var (
	n  int
	s  [20]int     = [20]int{}
	dp [20][20]int = [20][20]int{}
)

func dfs(now, sum int, lim bool) int {
	if !lim && dp[now][sum] != -1 {
		return dp[now][sum]
	}
	if now == 0 {
		return sum
	}
	res := 0
	to := s[now]
	if !lim {
		to = 9
	}
	for i := 0; i <= to; i++ {
		a := 0
		if i == 1 {
			a = 1
		}
		res += dfs(now-1, sum+a, lim && (i == s[now]))
	}
	if !lim {
		dp[now][sum] = res
	}
	return res
}

func main() {
	fmt.Scan(&n)
	cnt := 0
	for ; n > 0; n /= 10 {
		cnt++
		s[cnt] = n % 10
	}
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	fmt.Println(dfs(cnt, 0, true))
}
