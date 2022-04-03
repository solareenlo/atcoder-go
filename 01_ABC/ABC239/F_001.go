package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const maxn = 200010

var fa = make([]int, maxn)

func Find(p int) int {
	if p == fa[p] {
		return p
	}
	fa[p] = Find(fa[p])
	return fa[p]
}

func Merge(x, y int) bool {
	x = Find(x)
	y = Find(y)
	if x == y {
		return false
	}
	fa[x] = y
	return true
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	d := make([]int, n+1)
	u := 0
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &d[i])
		u += d[i]
		fa[i] = i
	}

	if u != ((n - 1) << 1) {
		fmt.Fprintln(out, -1)
		return
	}

	for i := 1; i <= m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		if Merge(u, v) {
			d[u]--
			d[v]--
			if d[u] < 0 || d[v] < 0 {
				fmt.Fprintln(out, -1)
				return
			}
		} else {
			fmt.Fprintln(out, -1)
			return
		}
	}

	s := make([]int, n+1)
	dg := make([]int, n+1)
	nxt := make([]int, n+1)
	h := make([]int, n+1)
	j := 0
	for i := 1; i <= n; i++ {
		u = Find(i)
		if s[u] == 0 {
			j++
			s[u] = j
		}
		if d[i] != 0 {
			dg[s[u]] += d[i]
			nxt[i] = h[s[u]]
			h[s[u]] = i
		}
	}

	type node struct{ ind, idx int }
	N := make([]node, j+1)
	t := 0
	for i := 1; i <= j; i++ {
		N[i] = node{dg[i], i}
	}
	sort.Slice(N, func(i, j int) bool {
		return N[i].ind < N[j].ind
	})
	N = append(N, node{0, 0})

	stk := make([]int, j+1)
	var i int
	for i = 1; i <= j; i++ {
		if N[i].ind == 1 {
			t++
			stk[t] = N[i].idx
		} else {
			break
		}
	}

	t1 := 0
	a1 := make([]int, n+1)
	a2 := make([]int, n+1)
	for t > 0 {
		u = stk[t]
		t--
		if N[i].idx == 0 {
			continue
		}
		t1++
		a1[t1] = h[u]
		a2[t1] = h[N[i].idx]
		d[a2[t1]]--
		if d[a2[t1]] == 0 {
			h[N[i].idx] = nxt[a2[t1]]
		}
		N[i].ind--
		if N[i].ind == 1 {
			t++
			stk[t] = N[i].idx
			i++
		}
	}

	a1[t1+1] = h[stk[1]]
	if stk[2] != 0 {
		t1++
		a2[t1] = h[stk[2]]
	}
	if t1+m != n-1 {
		fmt.Fprintln(out, -1)
	} else {
		for i := 1; i <= t1; i++ {
			fmt.Fprintln(out, a1[i], a2[i])
		}
	}
}
