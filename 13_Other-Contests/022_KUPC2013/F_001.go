package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = int(1e18)

	type pair struct {
		x, y int
	}

	var N, s int
	fmt.Fscan(in, &N, &s)
	s--

	l := make([]int, N)
	r := make([]int, N)
	R := make([]pair, 0)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &l[i], &r[i])
		R = append(R, pair{r[i], i})
	}
	sort.Slice(R, func(i, j int) bool {
		if R[i].x == R[j].x {
			return R[i].y < R[j].y
		}
		return R[i].x < R[j].x
	})

	dis := make([][]int, N)
	for i := range dis {
		dis[i] = make([]int, N)
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Fscan(in, &dis[i][j])
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			for k := 0; k < N; k++ {
				dis[j][k] = min(dis[j][k], dis[j][i]+dis[i][k])
			}
		}
	}

	dp := make([]int, N)
	for i := range dp {
		dp[i] = -INF
	}
	for i := 0; i < N; i++ {
		t := r[i] - max(l[i], dis[s][i])
		if t < 0 {
			continue
		}
		dp[i] = t
	}

	for i := 0; i < N; i++ {
		i0 := R[i].y
		for j := 0; j < i; j++ {
			i1 := R[j].y
			if dp[i1] == -INF {
				continue
			}
			D := r[i0] - r[i1] - dis[i1][i0]
			D = min(D, r[i0]-l[i0])
			if D < 0 {
				continue
			}
			dp[i0] = max(dp[i0], dp[i1]+D)
		}
	}

	ans := 0
	for i := 0; i < N; i++ {
		ans = max(ans, dp[i])
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
