package main

import "fmt"

var (
	h, w, black int
	dx          [4]int          = [4]int{1, -1, 0, 0}
	dy          [4]int          = [4]int{0, 0, 1, -1}
	s           [1001]string    = [1001]string{}
	vis         [1001][1001]int = [1001][1001]int{}
)

func dfs(x, y, black int) {
	if black >= vis[x][y] {
		return
	}
	vis[x][y] = black
	for i := 0; i < 4; i++ {
		nx, ny := x+dx[i], y+dy[i]
		if 0 <= nx && nx < h && 0 <= ny && ny < w && s[nx][ny] == '.' {
			dfs(nx, ny, black+1)
		}
	}
}

func main() {
	fmt.Scan(&h, &w)
	for i := 0; i < h; i++ {
		fmt.Scan(&s[i])
		for j := 0; j < w; j++ {
			vis[i][j] = int(1e9)
			if s[i][j] == '#' {
				black++
			}
		}
	}

	dfs(0, 0, 1)
	if res := h*w - black - vis[h-1][w-1]; res < 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(res)
	}
}
