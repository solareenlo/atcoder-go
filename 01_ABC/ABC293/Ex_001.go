package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF = int(1e18)

var f [200200][3]int
var v [200200][]int
var n, m int

func Dfs(u, fa int) {
	for i := 0; i < 3; i++ {
		f[u][i] = 1
	}
	for _, i := range v[u] {
		if i != fa {
			Dfs(i, u)
			for j := 2; j >= 0; j-- {
				if j < 2 && f[u][j]+f[i][1]-1 <= m {
					f[u][j+1] = min(f[u][j+1], max(f[u][j], f[i][1]))
				}
				if f[u][j]+f[i][2] > m {
					f[u][j] = INF
				} else {
					f[u][j] = max(f[u][j], f[i][2]+1)
				}
			}
		}
	}
}

func check(k int) bool {
	m = k
	Dfs(1, 0)
	return f[1][2] <= n
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)
	for i := 1; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		v[x] = append(v[x], y)
		v[y] = append(v[y], x)
	}
	Dfs(1, 0)
	l, r := 1, n
	asw := 0
	for l <= r {
		mid := (l + r) >> 1
		if check(mid) {
			r = mid - 1
			asw = mid
		} else {
			l = mid + 1
		}
	}
	fmt.Println(asw)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
