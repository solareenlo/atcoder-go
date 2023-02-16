package main

import (
	"bufio"
	"fmt"
	"os"
)

var a [1000][1000]byte
var f [1000][1000]bool
var t byte
var n, m, s int

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &m)
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(in, &s)
		for j := 0; j < m; j++ {
			a[i][j] = s[j]
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if a[i][j] == a[i+1][j] && a[i][j] == a[i+2][j] {
				f[i+2][j] = true
				f[i+1][j] = true
				f[i][j] = true
			}
			if a[i][j] == a[i][j+1] && a[i][j] == a[i][j+2] {
				f[i][j+2] = true
				f[i][j+1] = true
				f[i][j] = true
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if f[i][j] {
				t = a[i][j]
				s++
				dfs(i, j)
			}
		}
	}
	fmt.Println(s)
}

func dfs(x, y int) {
	if x < 0 || y < 0 || x >= n || y >= m || !f[x][y] || a[x][y] != t {
		return
	}
	f[x][y] = false
	dfs(x-1, y)
	dfs(x+1, y)
	dfs(x, y-1)
	dfs(x, y+1)
}
