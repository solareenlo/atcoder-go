package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 200005

type pair struct{ x, y int }

var (
	gph  = make([][]pair, N)
	vis  = make([]bool, N)
	val  = make([]int, N)
	l, r int
	n    int
)

func dfs(x, v int) {
	if vis[x] {
		if val[x] != v {
			fmt.Println("No")
			os.Exit(0)
		}
		return
	}
	vis[x] = true
	val[x] = v
	if x <= n {
		l = max(l, v)
	} else {
		r = min(r, v)
	}
	for _, i := range gph[x] {
		dfs(i.x, v+i.y)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var m, k int
	fmt.Fscan(in, &n, &m, &k)
	for i := 0; i < k; i++ {
		var x, y, v int
		fmt.Fscan(in, &x, &y, &v)
		gph[x] = append(gph[x], pair{y + n, v})
		gph[y+n] = append(gph[y+n], pair{x, -v})
	}

	for i := 1; i <= n+m; i++ {
		if !vis[i] {
			l = -1 << 60
			r = 1 << 60
			dfs(i, 0)
			if l > r {
				fmt.Println("No")
				return
			}
		}
	}
	fmt.Println("Yes")
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
