package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)

	a := make([]int, n)
	mp := map[int]int{}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		mp[a[i]] ^= 1
	}

	t := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(in, &t[i])
		mp[t[i]] ^= 1
	}

	e := 0
	c := make([]int, 5005)
	for u, v := range mp {
		if v != 0 {
			c[e] = u
			e++
		}
	}

	d := n - e
	if d < 0 || d%2 == 1 {
		fmt.Println(-1)
		return
	}
	sort.Ints(a)
	tmp := c[:e]
	sort.Ints(tmp)

	dp := [5005][5005]int{}
	for i := 1; i <= n; i++ {
		for j := 0; j <= d && j <= i; j += 2 {
			tmp1 := 1 << 60
			if i > 1 && j > 1 {
				tmp1 = dp[i-2][j-2] + abs(a[i-1]-a[i-2])
			}
			tmp2 := 1 << 60
			if i-j <= e && i-1-j >= 0 {
				tmp2 = dp[i-1][j] + abs(a[i-1]-c[i-1-j])
			}
			dp[i][j] = min(tmp1, tmp2)
		}
	}
	fmt.Println(dp[n][d])
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
