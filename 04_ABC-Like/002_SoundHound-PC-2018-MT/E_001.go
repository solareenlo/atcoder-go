package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct{ to, val int }

const N = 233333

var (
	e   = [N][]pair{}
	vis = [N][2]bool{}
	f   = [N][2]int{}
	flg = 1<<60 + 1
)

func dfs(x, op int) {
	vis[x][op] = true
	for _, y := range e[x] {
		if vis[y.to][op^1] {
			if f[y.to][op^1] != y.val-f[x][op] {
				flg *= 0
			}
		} else {
			f[y.to][op^1] = y.val - f[x][op]
			dfs(y.to, op^1)
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	for i := 0; i < m; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		e[a] = append(e[a], pair{b, c})
		e[b] = append(e[b], pair{a, c})
	}

	dfs(1, 0)

	l, r := 0, 1<<60
	for i := 1; i <= n; i++ {
		if vis[i][0] {
			l = max(l, 1-f[i][0])
		}
		if vis[i][1] {
			r = min(r, f[i][1]-1)
		}
		if vis[i][0] && vis[i][1] {
			if f[i][1] > f[i][0] && (f[i][1]-f[i][0])%2 == 0 {
				flg &= 1
			} else {
				flg &= 0
			}
		}
	}
	fmt.Println(max(min(flg, r-l+1), 0))
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
