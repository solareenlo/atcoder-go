package main

import "fmt"

var dx [4]int = [4]int{0, -1, 0, 1}
var dy [4]int = [4]int{1, 0, -1, 0}
var c [16]string
var m [16][16][4][16][16]int
var h, w int

func dfs(i, j, k, a, b int) int {
	if m[i][j][k][a][b] != 0 {
		return m[i][j][k][a][b]
	}
	if i == h-2 && j == w-2 && a == 0 && b == 0 {
		m[i][j][k][a][b] = 1
		return 1
	}
	k2 := k
	if c[i+dx[k2]][j+dy[k2]] != '#' && dfs(i+dx[k2], j+dy[k2], k2, a, b) == 1 {
		m[i][j][k][a][b] = 1
		return 1
	}
	k2 = (k + 1) % 4
	if a > 0 && c[i+dx[k2]][j+dy[k2]] != '#' && dfs(i+dx[k2], j+dy[k2], k2, a-1, b) == 1 {
		m[i][j][k][a][b] = 1
		return 1
	}
	k2 = (k + 3) % 4
	if b > 0 && c[i+dx[k2]][j+dy[k2]] != '#' && dfs(i+dx[k2], j+dy[k2], k2, a, b-1) == 1 {
		m[i][j][k][a][b] = 1
		return 1
	}
	m[i][j][k][a][b] = -1
	return -1
}

func main() {
	var a, b int
	fmt.Scan(&a, &b, &h, &w)
	for i := 0; i < h; i++ {
		fmt.Scan(&c[i])
	}
	if dfs(1, 1, 3, a, b) == 1 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
