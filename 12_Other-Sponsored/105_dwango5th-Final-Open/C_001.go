package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N = 100005
const INF = 1047483647

type tuple struct {
	l, r, x int
}

type node struct {
	t, x, y int
}

type str struct {
	mn, mnn, mn2, mnn2 int
}

var a []node
var w [N]str
var t2 [N * 2]str
var n int
var f [N]int

func min(a, b str) str {
	var c str
	if a.mn < b.mn {
		c.mn = a.mn
		c.mnn = a.mnn
		if b.mnn == a.mnn {
			if a.mn2 < b.mn2 {
				c.mn2 = a.mn2
				c.mnn2 = a.mnn2
			} else {
				c.mn2 = b.mn2
				c.mnn2 = b.mnn2
			}
			return c
		}
		if b.mn < a.mn2 {
			c.mn2 = b.mn
			c.mnn2 = b.mnn
		} else {
			c.mn2 = a.mn2
			c.mnn2 = a.mnn2
		}
	} else {
		c.mn = b.mn
		c.mnn = b.mnn
		if b.mnn == a.mnn {
			if a.mn2 < b.mn2 {
				c.mn2 = a.mn2
				c.mnn2 = a.mnn2
			} else {
				c.mn2 = b.mn2
				c.mnn2 = b.mnn2
			}
			return c
		}
		if a.mn < b.mn2 {
			c.mn2 = a.mn
			c.mnn2 = a.mnn
		} else {
			c.mn2 = b.mn2
			c.mnn2 = b.mnn2
		}
	}
	return c
}

var sta [32]int
var top int
var ans str

func modify2(i, x, y int) {
	t := str{x, y, INF, 0}
	for i <= 2*n {
		t2[i] = min(t2[i], t)
		i += -i & i
	}
}

func Query2(i int) {
	for i > 0 {
		ans = min(ans, t2[i])
		i -= -i & i
	}
}

func Find(x int) int {
	if f[x] != 0 {
		f[x] = Find(f[x])
		return f[x]
	}
	return x
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var al, ar, c [N]int

	res := 0
	fmt.Fscan(in, &n)
	l := make([]int, n+1)
	r := make([]int, n+1)
	a = make([]node, 2*n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &l[i], &r[i])
		a[2*i-1] = node{l[i], i, 0}
		a[2*i] = node{r[i], i, 1}
	}
	sort.Slice(a, func(i, j int) bool {
		if a[i].t == a[j].t {
			return a[i].y > a[j].y
		}
		return a[i].t < a[j].t
	})
	for i := 1; i <= 2*n; i++ {
		if a[i].y == 0 {
			al[a[i].x] = i
		} else {
			ar[a[i].x] = i
		}
	}
	q := make([]tuple, n+1)
	p := make([]tuple, n+1)
	for i := 1; i <= n; i++ {
		q[i] = tuple{al[i], ar[i], i}
		p[i] = tuple{al[i], ar[i], i}
	}
	sort.Slice(p, func(i, j int) bool {
		return p[i].l < p[j].l
	})
	sort.Slice(q, func(i, j int) bool {
		return q[i].r < q[j].r
	})
	las := 0
	for {
		for i := 1; i <= n; i++ {
			w[i] = str{INF, 0, INF, 0}
			c[i] = Find(i)
		}
		for i := 1; i <= n*2; i++ {
			t2[i] = str{INF, 0, INF, 0}
		}
		for i := 1; i <= n; i++ {
			if p[i].l+1 < p[i].r {
				ans = str{INF, 0, INF, 0}
				Query2(2*n - (p[i].l + 1) + 1)
				ans.mn -= l[p[i].x]
				ans.mn2 -= l[p[i].x]
				w[c[p[i].x]] = min(w[c[p[i].x]], ans)
			}
			modify2(2*n-p[i].r+1, r[p[i].x], c[p[i].x])
		}
		for i := 1; i <= n*2; i++ {
			t2[i] = str{INF, 0, INF, 0}
		}
		for i := n; i >= 1; i-- {
			if p[i].l+1 < p[i].r {
				ans = str{INF, 0, INF, 0}
				Query2(p[i].r - 1)
				w[c[p[i].x]] = min(w[c[p[i].x]], ans)
			}
			modify2(p[i].r, r[p[i].x]-l[p[i].x], c[p[i].x])
		}
		for i := 1; i <= n*2; i++ {
			t2[i] = str{INF, 0, INF, 0}
		}
		for i := 1; i <= n; i++ {
			if p[i].r < 2*n {
				ans = str{INF, 0, INF, 0}
				Query2(2*n - (p[i].r + 1) + 1)
				ans.mn += r[p[i].x] - l[p[i].x]
				ans.mn2 += r[p[i].x] - l[p[i].x]
				w[c[p[i].x]] = min(w[c[p[i].x]], ans)
			}
			modify2(2*n-p[i].r+1, 0, c[p[i].x])
		}
		for i := 1; i <= n*2; i++ {
			t2[i] = str{INF, 0, INF, 0}
		}
		for i := n; i >= 1; i-- {
			if q[i].l+1 < q[i].r {
				ans = str{INF, 0, INF, 0}
				Query2(q[i].r - 1)
				ans.mn += r[q[i].x]
				ans.mn2 += r[q[i].x]
				w[c[q[i].x]] = min(w[c[q[i].x]], ans)
			}
			modify2(q[i].l, -l[q[i].x], c[q[i].x])
		}
		for i := 1; i <= n; i++ {
			if Find(i) == i {
				u := w[i].mnn
				v := w[i].mn
				if w[i].mnn == i {
					u = w[i].mnn2
					v = w[i].mn2
				}
				if u != 0 {
					if i != Find(u) {
						f[i] = Find(u)
						res += v
					}
				}
			}
		}
		s := 0
		for i := 1; i <= n; i++ {
			if Find(i) == i {
				s++
			}
		}
		if s == 1 {
			break
		}
		if s == las {
			fmt.Println(-1)
			return
		}
		las = s
	}
	fmt.Println(res)
}
