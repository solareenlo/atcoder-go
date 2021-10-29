package main

import "fmt"

var (
	dx    []int = []int{1, 0, -1, 0}
	dy    []int = []int{0, 1, 0, -1}
	H, W  int
	field []string
	seen  [30][30]int
)

func dfs(h, w, k int) {
	if h < 0 || H <= h || w < 0 || W <= w {
		return
	}
	if field[h][w] == '#' {
		return
	}
	if seen[h][w] <= k {
		return
	} else {
		seen[h][w] = k
	}

	for i := 0; i < 4; i++ {
		dfs(h+dx[i], w+dy[i], k+1)
	}
}

func memset(a *[30][30]int, val int) {
	for i := range a {
		for j := range a[i] {
			(*a)[i][j] = val
		}
	}
}

func main() {
	fmt.Scan(&H, &W)

	field = make([]string, H)
	for i := range field {
		fmt.Scan(&field[i])
	}

	maxi := 0
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if field[i][j] == '.' {
				memset(&seen, 1<<60)
				dfs(i, j, 0)
				for h := 0; h < H; h++ {
					for w := 0; w < W; w++ {
						if seen[h][w] != seen[29][29] {
							maxi = max(maxi, seen[h][w])
						}
					}
				}
			}
		}
	}
	fmt.Println(maxi)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
