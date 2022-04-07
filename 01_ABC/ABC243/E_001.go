package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	const N = 310
	dis := [N][N]int{}
	for i := range dis {
		for j := range dis[i] {
			dis[i][j] = 1 << 60
		}
	}

	g := [N][N]int{}
	for i := 1; i <= m; i++ {
		var x, y, z int
		fmt.Fscan(in, &x, &y, &z)
		g[x][y] = 1
		g[y][x] = 1
		dis[y][x] = min(dis[x][y], z)
		dis[x][y] = dis[y][x]
	}

	ans := 0
	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				if dis[i][j] >= dis[i][k]+dis[k][j] {
					dis[i][j] = dis[i][k] + dis[k][j]
					ans += g[i][j]
					g[i][j] = 0
				}
			}
		}
	}

	fmt.Println(ans / 2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
