package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1005
const M = 200005

type edge struct{ to, nxt, num int }

var (
	tot int
	e   = make([]edge, M<<1)
	h   = make([]int, N)
	fa  = make([]int, N)
	p   = make([]int, N)
	ans = make([]int, 0)
)

func add(u, v, n int) {
	tot++
	e[tot].to = v
	e[tot].num = n
	e[tot].nxt = h[u]
	h[u] = tot
}

func find(x int) int {
	if x == fa[x] {
		return x
	}
	fa[x] = find(fa[x])
	return fa[x]
}

func dfs(u, fa, f int) bool {
	if u == f {
		return true
	}
	for i := h[u]; i > 0; i = e[i].nxt {
		v := e[i].to
		if v == fa {
			continue
		}
		if dfs(v, u, f) {
			ans = append(ans, e[i].num)
			p[u], p[v] = p[v], p[u]
			return true
		}
	}
	return false
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &p[i])
		fa[i] = i
	}

	var m int
	fmt.Fscan(in, &m)
	d := make([]int, N)
	for i := 1; i <= m; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		xx := find(x)
		yy := find(y)
		if xx == yy {
			continue
		}
		fa[xx] = yy
		d[x]++
		d[y]++
		add(x, y, i)
		add(y, x, i)
	}

	q := make([]int, 0)
	for i := 1; i <= n; i++ {
		if d[i] == 1 {
			q = append(q, i)
		}
	}

	for len(q) != 0 {
		u := q[0]
		q = q[1:]
		t := 0
		for i := 1; i <= n; i++ {
			if p[i] == u {
				t = i
			}
		}
		if !dfs(u, 0, t) {
			fmt.Fprintln(out, -1)
			return
		}
		for i := h[u]; i > 0; i = e[i].nxt {
			v := e[i].to
			d[v]--
			if d[v] == 1 {
				q = append(q, v)
			}
		}
	}

	fmt.Fprintln(out, len(ans))
	for i := range ans {
		fmt.Fprint(out, ans[i], " ")
	}
}
