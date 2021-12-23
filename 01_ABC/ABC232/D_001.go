package main

import "fmt"

var (
	h, w int
	a    = [105]string{}
	vis  = [105][105]bool{}
)

func dfs(x, y int) int {
	if x > h-1 || y > w-1 || vis[x][y] {
		return 0
	}
	vis[x][y] = true
	if a[x][y] == '#' {
		return 0
	}
	return max(dfs(x+1, y), dfs(x, y+1)) + 1
}
func main() {
	fmt.Scan(&h, &w)
	for i := 0; i < h; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Println(dfs(0, 0))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
