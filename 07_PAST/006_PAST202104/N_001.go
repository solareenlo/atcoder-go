package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, H int
	fmt.Fscan(in, &N, &H)

	type pair struct{ x, y int }
	AB := make([]pair, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &AB[i].x, &AB[i].y)
	}
	sort.Slice(AB, func(i, j int) bool {
		return AB[i].y*AB[j].x < AB[i].x*AB[j].y
	})

	dp := make([]int, 1<<17)
	for i := 0; i < N; i++ {
		A := AB[i].x
		B := AB[i].y
		for h := 1; h <= H; h++ {
			id := max(h-B, 0)
			dp[id] = max(dp[id], dp[h]+h*A)
		}
	}

	ans := 0
	for i := 0; i <= H; i++ {
		ans = max(ans, dp[i])
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
