package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = 1000000007
	const C = 20000

	type P struct {
		x, y int
	}

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	var dp [50001][2]int
	for i := 0; i < 50001; i++ {
		dp[i][0] = INF
		dp[i][1] = INF
	}

	que := make([]P, 0)
	dp[C][0] = 0
	que = append(que, P{C, 0})
	for len(que) > 0 {
		now := que[0]
		que = que[1:]
		pos := now.x
		cnt := now.y
		for i := 0; i < n; i++ {
			npos := pos
			if cnt%2 == 0 {
				npos += 2 * a[i]
			} else {
				npos -= 2 * a[i]
			}
			var tmp int
			if cnt == 0 {
				tmp = 1
			} else {
				tmp = 0
			}
			if 0 <= npos && npos < 50000 && dp[npos][tmp] >= INF {
				dp[npos][tmp] = dp[pos][cnt] + 1
				tmp := 0
				if cnt == 0 {
					tmp = 1
				}
				que = append(que, P{npos, tmp})
			}
		}
	}

	var q int
	fmt.Fscan(in, &q)
	for i := 0; i < q; i++ {
		var s, t int
		fmt.Fscan(in, &s, &t)
		res := min(dp[t-s+C][0], dp[t+s+C][1])
		if res >= INF {
			fmt.Println(-1)
		} else {
			fmt.Println(res)
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
