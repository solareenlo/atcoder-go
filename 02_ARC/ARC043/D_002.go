package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	p := make([]int, m)
	s := 0
	for i := range p {
		fmt.Scan(&p[i])
		s += p[i]
	}
	sort.Sort(sort.Reverse(sort.IntSlice(p)))

	var f func(x int) int
	f = func(x int) int {
		return x * (s - x)
	}

	dp := make([]int, s+1)
	for i := range dp {
		dp[i] = -1 << 60
	}
	dp[0] = 0

	cur := 0
	for _, x := range p {
		cur += x
		for i := s; i >= x; i-- {
			dp[i] = max(dp[i]+f(cur-i), dp[i-x]+f(i))
		}
	}

	ans := 0
	for i := 0; i < s+1; i++ {
		res := dp[i]
		res += (n - m - 1) * f(i)
		ans = max(ans, res)
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
