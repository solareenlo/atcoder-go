package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 200200

type Edge struct {
	y, w, pre int
}

var E [N << 1]Edge
var sz, lst, w, pos, son [N]int

func dfs(x int) {
	sz[x] = 1
	for i := lst[x]; i > 0; i = E[i].pre {
		y := E[i].y
		if sz[y] == 0 {
			w[y] = E[i].w
			pos[i>>1] = y
			dfs(y)
			sz[x] += sz[y]
			if sz[y] > sz[son[x]] {
				son[x] = y
			}
		}
	}
}

var n2 int
var dfn, top, dep, b, fa [N]int

func dfs2(x, tp, h int) {
	n2++
	dfn[x] = n2
	top[x] = tp
	h++
	dep[x] = h
	b[n2] += w[x]
	b[n2+sz[x]] -= w[x]
	if son[x] == 0 {
		return
	}
	dfs2(son[x], tp, h)
	for i := lst[x]; i > 0; i = E[i].pre {
		y := E[i].y
		if top[y] == 0 {
			fa[y] = x
			dfs2(y, y, h)
		}
	}
}

var n int

func upd(i, x int) {
	for i <= n {
		b[i] += x
		i += (i & -i)
	}
}

func qry(i int) int {
	s := 0
	i = dfn[i]
	for i > 0 {
		s += b[i]
		i &= i - 1
	}
	return s
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &n)
	m := 1
	for i := 1; i < n; i++ {
		var x, y, tw int
		fmt.Fscan(in, &x, &y, &tw)
		m++
		E[m] = Edge{y, tw, lst[x]}
		lst[x] = m
		m++
		E[m] = Edge{x, tw, lst[y]}
		lst[y] = m
	}
	dfs(1)
	dfs2(1, 1, 0)
	for i := 1; i <= n; i++ {
		x := i + (i & -i)
		if x <= n {
			b[x] += b[i]
		}
	}
	var q int
	fmt.Fscan(in, &q)
	for q > 0 {
		q--
		var opt, x, y int
		fmt.Fscan(in, &opt, &x, &y)
		if opt == 1 {
			x = pos[x]
			upd(dfn[x], y-w[x])
			upd(dfn[x]+sz[x], w[x]-y)
			w[x] = y
		} else {
			s := qry(x) + qry(y)
			tx := top[x]
			ty := top[y]
			for tx != ty {
				if dep[tx] < dep[ty] {
					tx, ty = ty, tx
					x, y = y, x
				}
				x = fa[tx]
				tx = top[x]
			}
			if dep[x] < dep[y] {
				fmt.Fprintln(out, s-(qry(x)<<1))
			} else {
				fmt.Fprintln(out, s-(qry(y)<<1))
			}
		}
	}
}
