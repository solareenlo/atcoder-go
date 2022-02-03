package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N = 200005
const M = 1000005

type edge struct{ v, nx int }
type node struct{ pos, id int }

var (
	e    = make([]edge, M)
	a    = make([]node, N<<1)
	n    int
	ne   int
	ans  int
	cntn int
	f    = make([]int, N)
	id   = make([]int, M)
	tot  int
	top  int
	cnt  int
	dfn  = make([]int, N)
	low  = make([]int, N)
	st   = make([]int, N)
	arr  = make([]int, N)
	vis  = make([]bool, N)
)

func read(u, v int) {
	ne++
	e[ne].v = v
	e[ne].nx = f[u]
	f[u] = ne
}

func build(k, l, r int) {
	cntn++
	id[k] = cntn
	if l == r {
		read(id[k], (a[l].id+n)%(2*n))
		return
	}
	mid := (l + r) >> 1
	build(k<<1, l, mid)
	build(k<<1|1, mid+1, r)
	read(id[k], id[k<<1])
	read(id[k], id[k<<1|1])
}

func getlk(k, l, r, x, y, z int) {
	if y < x {
		return
	}
	mid := (l + r) >> 1
	if l == x && r == y {
		read(z, id[k])
	} else if y <= mid {
		getlk(k<<1, l, mid, x, y, z)
	} else if x > mid {
		getlk(k<<1|1, mid+1, r, x, y, z)
	} else {
		getlk(k<<1, l, mid, x, mid, z)
		getlk(k<<1|1, mid+1, r, mid+1, y, z)
	}
}

func Tarjan(u int) {
	tot++
	dfn[u] = tot
	low[u] = tot
	vis[u] = true
	top++
	st[top] = u
	for i := f[u]; i > 0; i = e[i].nx {
		v := e[i].v
		if dfn[v] == 0 {
			Tarjan(v)
			low[u] = min(low[u], low[v])
		} else if vis[v] {
			low[u] = min(low[u], dfn[v])
		}
	}
	if dfn[u] == low[u] {
		cnt++
		for st[top] != u {
			arr[st[top]] = cnt
			vis[st[top]] = false
			top--
		}
		arr[st[top]] = cnt
		vis[st[top]] = false
		top--
	}
}

func check(x int) bool {
	top = 0
	ne = 0
	tot = 0
	cnt = 0
	cntn = 2 * n
	for i := range f {
		f[i] = 0
		dfn[i] = 0
		low[i] = 0
	}
	build(1, 1, cntn)
	for i := 1; i <= 2*n; i++ {
		l := upperBound(a[1:2*n+1], node{a[i].pos - x, 0}) + 1
		r := upperBound(a[1:2*n+1], node{a[i].pos + x - 1, 0})
		getlk(1, 1, 2*n, l, i-1, a[i].id)
		getlk(1, 1, 2*n, i+1, r, a[i].id)
	}
	for i := 1; i <= 2*n; i++ {
		if dfn[i] == 0 {
			Tarjan(i)
		}
	}
	for i := 1; i <= n; i++ {
		if arr[i] == arr[i+n] {
			return false
		}
	}
	return true
}

func upperBound(a []node, x node) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i].pos > x.pos
	})
	return idx
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i].pos, &a[i+n].pos)
		a[i].id = i
		a[i+n].id = i + n
	}
	b := a[1 : 2*n+1]
	sort.Slice(b, func(i, j int) bool {
		return b[i].pos < b[j].pos
	})
	l := 0
	r := a[n+n].pos - a[1].pos + 1
	for l <= r {
		mid := (l + r) >> 1
		if check(mid) {
			l = mid + 1
			ans = mid
		} else {
			r = mid - 1
		}
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
