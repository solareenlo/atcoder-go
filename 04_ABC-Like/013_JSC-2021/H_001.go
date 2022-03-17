package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 400005

var (
	vis = make([]int, N)
	a   = make([]int, N)
	id  = make([]int, N)
	sum = make([]int, N)
	w   = make([]int, N)
	dep = make([]int, N)
	lst = make([]int, N)
	nxt = make([]int, N)
	rt  = make([]int, N)
	d   = make([]int, N)
	c   = make([]int, N)
	num = make([]int, N<<2)
	s   = make([]int, N<<2)
	fa  = [N][21]int{}
	cnt int
	ans int
)

func cir() {
	u := 1
	for vis[u] == 0 {
		vis[u] = 1
		u = a[u]
	}
	for i := u; vis[i] >= 0; i = a[i] {
		vis[i] = -1
		cnt++
		c[cnt] = i
		id[i] = cnt
		sum[cnt] = sum[cnt-1] + w[i]
	}
}

func DFS(x, f int) {
	fa[x][0] = f
	dep[x] = dep[f] + 1
	for i := 1; i <= 19; i++ {
		fa[x][i] = fa[fa[x][i-1]][i-1]
	}
	for i := lst[x]; i > 0; i = nxt[i] {
		if rt[i] != 0 {
			continue
		}
		rt[i] = rt[x]
		DFS(i, x)
	}
}

func lca(x, y int) int {
	if dep[x] < dep[y] {
		x, y = y, x
	}
	k := dep[x] - dep[y]
	for i := 0; i <= 19; i++ {
		if k&(1<<i) != 0 {
			x = fa[x][i]
		}
	}
	if x == y {
		return x
	}
	for i := 19; i >= 0; i-- {
		if fa[x][i]^fa[y][i] != 0 {
			x = fa[x][i]
			y = fa[y][i]
		}
	}
	return fa[x][0]
}

func update(p, l, r int) {
	if num[p] != 0 {
		s[p] = sum[r] - sum[l-1]
	} else {
		s[p] = s[p<<1] + s[p<<1|1]
	}
}

func modify(p, l, r, x, y, k int) {
	if y < x {
		return
	}
	if x <= l && y >= r {
		num[p] += k
		update(p, l, r)
		return
	}
	mid := (l + r) >> 1
	if x <= mid && y >= l {
		modify(p<<1, l, mid, x, y, k)
	}
	if y > mid && x <= r {
		modify(p<<1|1, mid+1, r, x, y, k)
	}
	update(p, l, r)
}

func gett(x int) {
	for i := lst[x]; i > 0; i = nxt[i] {
		if rt[i] != rt[x] {
			continue
		}
		gett(i)
		d[x] += d[i]
	}
	tmp := 0
	if d[x] > 0 && rt[x] != x {
		tmp = 1
	}
	ans += tmp * w[x]
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i], &w[i])
	}
	cir()
	for i := 1; i <= n; i++ {
		if vis[i] != -1 {
			nxt[i] = lst[a[i]]
			lst[a[i]] = i
		}
	}
	for i := 1; i <= cnt; i++ {
		rt[c[i]] = c[i]
		DFS(c[i], 0)
	}

	v := make([][]int, N)
	for i := 1; i <= m; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		if rt[x] == rt[y] {
			l := lca(x, y)
			d[x]++
			d[y]++
			d[l] -= 2
			continue
		}
		d[x]++
		d[y]++
		x = rt[x]
		y = rt[y]
		d[x]--
		d[y]--
		v[x] = append(v[x], y)
		v[y] = append(v[y], x)
		if id[x] > id[y] {
			x, y = y, x
		}
		modify(1, 1, cnt, id[x], id[y]-1, 1)
	}
	mn := s[1]
	for i := 1; i <= cnt; i++ {
		gett(c[i])
	}
	for i := 1; i < cnt; i++ {
		sz := len(v[c[i]])
		for j := 0; j < sz; j++ {
			y := v[c[i]][j]
			if id[c[i]] < id[y] {
				modify(1, 1, cnt, id[c[i]], id[y]-1, -1)
				modify(1, 1, cnt, id[y], cnt, 1)
				modify(1, 1, cnt, 1, id[c[i]]-1, 1)
			} else {
				modify(1, 1, cnt, id[c[i]], cnt, -1)
				modify(1, 1, cnt, 1, id[y]-1, -1)
				modify(1, 1, cnt, id[y], id[c[i]]-1, 1)
			}
		}
		mn = min(mn, s[1])
	}
	fmt.Println(ans + mn)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
