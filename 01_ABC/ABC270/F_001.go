package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type edge struct {
	u, v, w int
}

const N = 200005
const M = 600005

var n, m int
var e []edge
var f [N]int

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &m)
	k := 0
	e = make([]edge, M)
	for i := 1; i <= n; i++ {
		var x int
		fmt.Fscan(in, &x)
		e[k] = edge{n + 1, i, x}
		k++
	}
	for i := 1; i <= n; i++ {
		var x int
		fmt.Fscan(in, &x)
		e[k] = edge{n + 2, i, x}
		k++
	}
	for m > 0 {
		var u, v, w int
		fmt.Fscan(in, &u, &v, &w)
		e[k] = edge{u, v, w}
		k++
		m--
	}
	tmp := e[:k]
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].w < tmp[j].w
	})
	fmt.Println(min(krus(0, 0), krus(0, 1), krus(1, 0), krus(1, 1)))
}

func find(x int) int {
	if x == f[x] {
		return x
	}
	f[x] = find(f[x])
	return f[x]
}

func krus(x, y int) int {
	s, cnt := 0, 0
	for i := 1; i <= n+2; i++ {
		f[i] = i
	}
	for i := range e {
		u := e[i].u
		v := e[i].v
		w := e[i].w
		if x == 0 && (u == n+1 || v == n+1) {
			continue
		}
		if y == 0 && (u == n+2 || v == n+2) {
			continue
		}
		u = find(u)
		v = find(v)
		if u != v {
			f[u] = v
			s += w
			cnt++
		}
	}
	if cnt == n-1+x+y {
		return s
	}
	return 1e18
}

func min(a ...int) int {
	res := a[0]
	for i := range a {
		if res > a[i] {
			res = a[i]
		}
	}
	return res
}
