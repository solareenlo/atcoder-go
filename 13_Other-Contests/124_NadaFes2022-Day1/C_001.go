package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = int(1e18)

	type pair struct {
		x, y int
	}
	var S [2000]string
	var dist, lim [2000][2000]int
	var d [5]int = [5]int{0, 1, 0, -1}

	var H, W, K int
	fmt.Fscan(in, &H, &W, &K)
	Q := make([]pair, 0)
	for i := 0; i < H; i++ {
		fmt.Fscan(in, &S[i])
		for j := 0; j < W; j++ {
			dist[i][j] = INF
			lim[i][j] = INF
			if S[i][j] == '@' {
				Q = append(Q, pair{i, j})
				lim[i][j] = 0
			}
		}
	}
	for len(Q) > 0 {
		x := Q[0].x
		y := Q[0].y
		Q = Q[1:]
		for r := 0; r < 4; r++ {
			tx := x + d[r]
			ty := y + d[r+1]
			if tx < 0 || ty < 0 || tx >= H || ty >= W || S[tx][ty] == '#' || lim[tx][ty] <= lim[x][y]+1 {
				continue
			}
			lim[tx][ty] = lim[x][y] + 1
			Q = append(Q, pair{tx, ty})
		}
	}
	Q = append(Q, pair{0, 0})
	dist[0][0] = 0
	for len(Q) > 0 {
		x := Q[0].x
		y := Q[0].y
		Q = Q[1:]
		for r := 0; r < 4; r++ {
			tx := x + d[r]
			ty := y + d[r+1]
			if tx < 0 || ty < 0 || tx >= H || ty >= W || S[tx][ty] == '#' || dist[tx][ty] <= dist[x][y]+1 {
				continue
			}
			if dist[x][y]+1 >= K*lim[tx][ty] {
				continue
			}
			dist[tx][ty] = dist[x][y] + 1
			Q = append(Q, pair{tx, ty})
		}
	}
	if dist[H-1][W-1] < INF {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
