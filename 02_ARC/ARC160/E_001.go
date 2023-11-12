package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	const maxn = 200005

	var sta, w [maxn]int

	G := make([][]int, maxn)

	top := 0

	var dfs func(int, int)
	dfs = func(x, fa int) {
		if len(G[x]) == 1 {
			top++
			sta[top] = x
		}
		for _, V := range G[x] {
			if V^fa != 0 {
				dfs(V, x)
			}
		}
	}

	var solve func()
	solve = func() {
		var n int
		fmt.Fscan(in, &n)
		for i := 1; i <= n; i++ {
			fmt.Fscan(in, &w[i])
			G[i] = make([]int, 0)
		}
		for i := 1; i < n; i++ {
			var x, y int
			fmt.Fscan(in, &x, &y)
			G[x] = append(G[x], y)
			G[y] = append(G[y], x)
		}
		top = 0
		dfs(1, 0)
		fmt.Fprintln(out, (top+1)/2)
		p, mn, pos := n, 0, 0
		if top&1 != 0 {
			for pos == 0 {
				w[mn] = 1e9 + 1
				mn = 0
				for i := 1; i <= n; i++ {
					if w[i] < w[mn] {
						mn = i
					}
				}
				for _, V := range G[mn] {
					tmp := top
					dfs(V, mn)
					if top > tmp+1 {
						pos = sta[top]
					}
					top = tmp
				}
			}
			fmt.Fprintln(out, mn, pos)
			for i := 1; i <= top; i++ {
				if sta[i] == pos {
					p = i
				}
			}
		}
		for i := 1; i <= top/2; i++ {
			var tmp0 = 0
			if i >= p {
				tmp0 = 1
			}
			var tmp1 = 0
			if i+top/2 >= p {
				tmp1 = 1
			}
			fmt.Fprintln(out, sta[i+tmp0], sta[i+top/2+tmp1])
		}
	}

	var T int
	fmt.Fscan(in, &T)
	for T > 0 {
		T--
		solve()
	}
}
