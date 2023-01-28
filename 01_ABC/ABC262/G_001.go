package main

import (
	"bufio"
	"fmt"
	"os"
)

var a [55]int
var f [55][55][55][55]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	for i, _ := range f {
		for j, _ := range f[i] {
			for k, _ := range f[i][j] {
				for l, _ := range f[i][j][k] {
					f[i][j][k][l] = -1
				}
			}
		}
	}

	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	fmt.Println(dfs(1, n, 1, 50))
}

func dfs(l, r, x, y int) int {
	if l > r {
		return 0
	}
	if ^f[l][r][x][y] != 0 {
		return f[l][r][x][y]
	}
	s := dfs(l+1, r, x, y)
	if a[l] < x || a[l] > y {
		f[l][r][x][y] = s
		return f[l][r][x][y]
	}
	for i := l; i <= r; i++ {
		s = max(s, 1+dfs(l+1, i, x, a[l])+dfs(i+1, r, a[l], y))
	}
	f[l][r][x][y] = s
	return f[l][r][x][y]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
