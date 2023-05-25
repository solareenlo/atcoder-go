package main

import (
	"bufio"
	"fmt"
	"os"
)

var suf, val, tl [2010]int
var es [2010][]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	for i := 1; i < n; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u++
		v++
		es[u] = append(es[u], v)
		es[v] = append(es[v], u)
	}

	ans := 0
	for i := 1; i <= n; i++ {
		for j := 1; j < n+1; j++ {
			suf[j] = 0
		}
		for j := 1; j < n+1; j++ {
			tl[j] = j
		}
		dfs(i, 0)
		r := i
		cnt := n
		cur := 0
		for !empty(r) {
			u := r
			r = pop(r)
			for u > 0 {
				cur += cnt * (val[u])
				cnt--
				u = suf[u]
			}
		}
		ans = max(ans, cur)
	}
	fmt.Println(ans + n - 1)
}

func dfs(u, fa int) {
	for _, v := range es[u] {
		if v != fa {
			dfs(v, u)
		}
	}
	if fa != 0 {
		val[u] = len(es[u]) - 2
	} else {
		val[u] = len(es[u])
	}
	nw := I{val[u], 1}
	cur := 0
	for _, v := range es[u] {
		if v != fa {
			cur = merge(cur, v)
		}
	}
	for !empty(cur) && nw.lessThan(top(cur)) {
		nw = nw.plus(top(cur))
		suf[tl[u]] = cur
		tl[u] = tl[cur]
		cur = pop(cur)
	}
	z[u].v = nw
	z[u].l = 0
	z[u].r = 0
	z[u].dis = 1
	merge(u, cur)
}

type I struct {
	sum, cnt int
}

func (a I) lessThan(b I) bool {
	return a.sum*b.cnt <= a.cnt*b.sum
}

func (a I) plus(b I) I {
	return I{a.sum + b.sum, a.cnt + b.cnt}
}

var z [2010]_z

type _z struct {
	v    I
	l, r int
	dis  int
}

func top(u int) I { return z[u].v }

func pu(u int) {
	if z[z[u].l].dis < z[z[u].r].dis {
		z[u].l, z[u].r = z[u].r, z[u].l
	}
	z[u].dis = z[z[u].r].dis + 1
}

func merge(u, v int) int {
	if u == 0 || v == 0 {
		return u | v
	}
	if z[v].v.lessThan(z[u].v) {
		z[u].r = merge(z[u].r, v)
		pu(u)
		return u
	} else {
		z[v].r = merge(z[v].r, u)
		pu(v)
		return v
	}
}

func pop(u int) int {
	return merge(z[u].l, z[u].r)
}

func empty(u int) bool { return u == 0 }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
