package main

import (
	"bufio"
	"fmt"
	"os"
)

type p struct{ v, w int }

const NUM = 400_044

var (
	G    = make([][]p, NUM)
	idx  int
	ok   = make([]int, NUM)
	dfn  = make([]int, NUM)
	ccno = make([]int, NUM)
	sz   = make([]int, NUM)
	cnt  int
)

func dfs(u, fa int) int {
	idx++
	lowu := idx
	dfn[u] = idx
	for i := range G[u] {
		v := G[u][i].v
		w := G[u][i].w
		if dfn[v] == 0 {
			ccno[w] = cnt
			lowv := dfs(v, u)
			lowu = min(lowu, lowv)
			sz[u] += sz[v] + 1
			if lowv > dfn[u] {
				ok[w] = ^sz[v] & 1
			} else {
				ok[w] = 1
			}
		} else if v^fa != 0 && dfn[v] < dfn[u] {
			ccno[w] = cnt
			ok[w] = 1
			lowu = min(lowu, dfn[v])
			sz[v]++
		}
	}
	return lowu
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	for i := 1; i <= n*2+1; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		G[u] = append(G[u], p{n*2 + 1 + v, i})
		G[n*2+1+v] = append(G[n*2+1+v], p{u, i})
	}

	cc := 0
	for i := 1; i <= n*4+2; i++ {
		if dfn[i] == 0 {
			cnt++
			dfs(i, 0)
			if sz[i]&1 != 0 {
				if cc != 0 {
					cc = -1
				} else {
					cc = cnt
				}
			}
		}
	}

	for i := 1; i <= n*2+1; i++ {
		if ^cc != 0 && ccno[i] == cc && ok[i] != 0 {
			fmt.Fprintln(out, "OK")
		} else {
			fmt.Fprintln(out, "NG")
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
