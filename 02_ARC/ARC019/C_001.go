package main

import "fmt"

type P struct{ x, y, k int }

var (
	inf     int = 1 << 60
	dx          = []int{0, 1, 0, -1}
	dy          = []int{1, 0, -1, 0}
	R, C, K int
	dist1   = make([][][]int, 0)
	dist2   = make([][][]int, 0)
	dist3   = make([][][]int, 0)
	G       = [50]string{}
)

func dijkstra(dist [][][]int, r, c int) {
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			for k := 0; k < 101; k++ {
				dist[i][j][k] = inf
			}
		}
	}
	dist[r][c][0] = 0
	Q := make([]P, 0)
	Q = append(Q, P{r, c, 0})
	for len(Q) > 0 {
		x := Q[0].x
		y := Q[0].y
		k := Q[0].k
		Q = Q[1:]
		for i := 0; i < 4; i++ {
			X := x + dx[i]
			Y := y + dy[i]
			if 0 <= X && X < R && 0 <= Y && Y < C && G[X][Y] != 'T' {
				l := k
				if G[X][Y] == 'E' {
					l++
				}
				if l <= K {
					if dist[X][Y][l] > dist[x][y][k]+1 {
						Q = append(Q, P{X, Y, l})
					}
					dist[X][Y][l] = min(dist[X][Y][l], dist[x][y][k]+1)
				}
			}
		}
	}
}

func main() {
	fmt.Scan(&R, &C, &K)
	for i := 0; i < R; i++ {
		fmt.Scan(&G[i])
	}

	dist1 = make([][][]int, 50)
	dist2 = make([][][]int, 50)
	dist3 = make([][][]int, 50)
	for i := range dist1 {
		dist1[i] = make([][]int, 50)
		for j := range dist1[i] {
			dist1[i][j] = make([]int, 101)
		}
	}
	for i := range dist2 {
		dist2[i] = make([][]int, 50)
		for j := range dist2[i] {
			dist2[i][j] = make([]int, 101)
		}
	}
	for i := range dist3 {
		dist3[i] = make([][]int, 50)
		for j := range dist3[i] {
			dist3[i][j] = make([]int, 101)
		}
	}

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if G[i][j] == 'S' {
				dijkstra(dist1, i, j)
			}
			if G[i][j] == 'C' {
				dijkstra(dist2, i, j)
			}
			if G[i][j] == 'G' {
				dijkstra(dist3, i, j)
			}
		}
	}
	ans := inf
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			for k := 0; k < K; k++ {
				dist1[i][j][k+1] = min(dist1[i][j][k+1], dist1[i][j][k])
				dist2[i][j][k+1] = min(dist2[i][j][k+1], dist2[i][j][k])
				dist3[i][j][k+1] = min(dist3[i][j][k+1], dist3[i][j][k])
			}
			for k := 0; k < K+1; k++ {
				for l := 0; k+l < K+1; l++ {
					ans = min(ans, dist1[i][j][k]+dist2[i][j][l]*2+dist3[i][j][K-k-l])
					if G[i][j] == 'E' {
						if k < K && l < K {
							ans = min(ans, dist1[i][j][k+1]+dist2[i][j][l+1]*2+dist3[i][j][K-k-l])
						}
						if l < K && k+l > 0 {
							ans = min(ans, dist1[i][j][k]+dist2[i][j][l+1]*2+dist3[i][j][K-k-l+1])
						}
						if k+l > 0 && k < K {
							ans = min(ans, dist1[i][j][k+1]+dist2[i][j][l]*2+dist3[i][j][K-k-l+1])
						}
					}
				}
			}
		}
	}

	if ans == inf {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
