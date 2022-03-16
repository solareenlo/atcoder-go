package main

import (
	"bufio"
	"fmt"
	"os"
)

type tuple struct{ x, y, z int }

var (
	vis = make([]tuple, 0)
)

func update(i, j, k int) {
	vis = append(vis, tuple{i, j, k})
	vis = append(vis, tuple{j, i, k})
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, 6006)
	for i := 1; i <= 3*n; i++ {
		fmt.Fscan(in, &a[i])
	}

	const N = 2005
	rowMax := [N]int{}
	for i := 0; i < N; i++ {
		rowMax[i] = -1 << 60
	}
	dp := [N][N]int{}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			dp[i][j] = -1 << 60
		}
	}

	zero := 0
	dp[a[1]][a[2]] = zero
	rowMax[a[1]] = zero
	rowMax[a[2]] = zero
	allMax := 0
	for i := 3; i <= 3*n; i += 3 {
		if a[i] == a[i+1] && a[i] == a[i+2] {
			zero--
			continue
		}
		for k := 0; k < 3; k++ {
			a[i], a[i+1] = a[i+1], a[i]
			a[i+1], a[i+2] = a[i+2], a[i+1]
			update(a[i], a[i+1], max(allMax, dp[a[i+2]][a[i+2]]+1))
			for j := 1; j <= n; j++ {
				if a[i+1] == a[i+2] {
					update(j, a[i], dp[j][a[i+1]]+1)
				}
				update(j, a[i], rowMax[j])
			}
		}
		for _, j := range vis {
			e := j.x
			f := j.y
			g := j.z
			dp[e][f] = max(dp[e][f], g)
			rowMax[e] = max(rowMax[e], g)
			allMax = max(allMax, g)
		}
		vis = vis[:0]
	}
	fmt.Println(allMax - zero)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
