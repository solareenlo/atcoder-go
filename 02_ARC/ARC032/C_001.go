package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n)
	b := make([]int, n)
	amax := 0
	start := make([][]int, 1000010)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i], &b[i])
		start[a[i]] = append(start[a[i]], i)
		amax = max(amax, a[i])
	}

	type P struct{ x, y int }
	dp := [1000010]P{}
	dp[amax] = P{1, start[amax][0]}
	for i := amax - 1; i >= 0; i-- {
		dp[i] = dp[i+1]
		for _, p := range start[i] {
			if dp[i].x < dp[b[p]].x+1 {
				dp[i] = P{dp[b[p]].x + 1, p}
			} else if dp[i].x == dp[b[p]].x+1 && dp[i].y > p {
				dp[i] = P{dp[b[p]].x + 1, p}
			}
		}
	}

	fmt.Fprintln(out, dp[0].x)
	cur := dp[0].y
	for i := dp[0].x; i > 0; i-- {
		fmt.Fprint(out, cur+1)
		if i > 1 {
			fmt.Fprint(out, " ")
		}
		cur = dp[b[cur]].y
	}
	fmt.Fprintln(out)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
