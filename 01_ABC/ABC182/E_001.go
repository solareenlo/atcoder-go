package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	di         = []int{-1, 0, 1, 0}
	dj         = []int{0, -1, 0, 1}
	h, w, n, m int
	light      = [1505][1505]bool{}
	wall       = [1505][1505]bool{}
	ok         = [1505][1505]bool{}
	visited    = [1505][1505]bool{}
	memo       = [1505][1505]bool{}
)

func f(v, i, j int) bool {
	if i < 0 || j < 0 || i >= h || j >= w {
		return false
	}
	if wall[i][j] {
		return false
	}
	if light[i][j] {
		return true
	}
	if visited[i][j] {
		return memo[i][j]
	}
	visited[i][j] = true
	memo[i][j] = f(v, i+di[v], j+dj[v])
	return memo[i][j]
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &h, &w, &n, &m)
	for i := 0; i < n; i++ {
		var r, c int
		fmt.Fscan(in, &r, &c)
		r--
		c--
		light[r][c] = true
	}

	for i := 0; i < m; i++ {
		var r, c int
		fmt.Fscan(in, &r, &c)
		r--
		c--
		wall[r][c] = true
	}

	for k := 0; k < 4; k++ {
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				visited[i][j] = false
			}
		}
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				if f(k, i, j) {
					ok[i][j] = true
				}
			}
		}
	}

	res := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if ok[i][j] {
				res++
			}
		}
	}
	fmt.Println(res)
}
