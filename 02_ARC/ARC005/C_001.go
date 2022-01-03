package main

import "fmt"

var (
	h, w int
	v    = [505][505]int{}
	dy   = [4]int{0, 0, 1, -1}
	dx   = [4]int{1, -1, 0, 0}
	flag bool
	m    = [505]string{}
)

func dfs(y, x, f int) {
	if y < 0 || y >= h || x < 0 || x >= w || (v[y][x] != 0 && f >= v[y][x]) {
		return
	}
	if m[y][x] == 'g' {
		flag = true
		return
	}
	if m[y][x] == '#' {
		if f == 3 {
			return
		}
		f++
	}
	v[y][x] = f
	for i := 0; i < 4; i++ {
		dfs(y+dy[i], x+dx[i], f)
	}
}

func main() {
	fmt.Scan(&h, &w)

	var a, b int
	for i := 0; i < h; i++ {
		fmt.Scan(&m[i])
		for j := 0; j < w; j++ {
			if m[i][j] == 's' {
				a = i
				b = j
			}
		}
	}

	dfs(a, b, 1)

	if flag {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
