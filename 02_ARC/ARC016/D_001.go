package main

import "fmt"

func main() {
	var n, m, h int
	fmt.Scan(&n, &m, &h)

	G := make([][]int, 100)
	for i := 0; i < m; i++ {
		var f, t int
		fmt.Scan(&f, &t)
		G[f-1] = append(G[f-1], t-1)
	}

	D := make([]int, 100)
	for i := 0; i < n; i++ {
		fmt.Scan(&D[i])
	}

	dp := [100][100]float64{}
	ng, ok := 10000000.0, 0.0
	for z := 0; z < 100; z++ {
		x := (ng + ok) / 2.0
		for i := n - 1; i >= 0; i-- {
			m := len(G[i])
			for j := 0; j < h; j++ {
				if i == n-1 {
					dp[i][j] = 0
					continue
				}
				if m == 0 {
					dp[i][j] = x + float64(j)
					continue
				}
				t := 1.0
				for _, v := range G[i] {
					c := D[v] + j
					if c < h {
						t += dp[v][c] / float64(m)
					} else {
						t += x + float64(j)
					}
				}
				dp[i][j] = min(t, x+float64(j))
			}
		}
		if x <= dp[0][0] {
			ok = x
		} else {
			ng = x
		}
	}

	if ok > 1e6 {
		fmt.Println(-1)
	} else {
		fmt.Println(ok)
	}
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
