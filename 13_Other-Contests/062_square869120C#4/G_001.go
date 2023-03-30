package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N = 450000
const B = 700

type Que struct {
	tp, x, y, z int
}

var n, q, tot int
var a, fa, dep, sz, dfn, px, pd [N + 5]int
var e [N + 5][]int
var dfns int
var q2 [N + 5]Que
var q1 []Que
var x, d []int
var xs, ds int
var pl, pr [B + 5]int
var c, s [B + 5][B*2 + 5]int
var out = bufio.NewWriter(os.Stdout)

func main() {
	in := bufio.NewReader(os.Stdin)
	defer out.Flush()

	x = make([]int, B*2+5)
	d = make([]int, B+5)

	fmt.Fscan(in, &n, &q)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &fa[i], &a[i])
		if i != 0 {
			e[fa[i]] = append(e[fa[i]], i)
			dep[i] = dep[fa[i]] + 1
		} else {
			dep[i] = 1
		}
	}
	tot = n
	for i := 1; i <= q; i++ {
		var t Que
		fmt.Fscan(in, &t.tp, &t.x, &t.y)
		if t.tp < 3 {
			t.y = min(t.y, tot-dep[t.x])
		} else {
			e[t.x] = append(e[t.x], tot)
			fa[tot] = t.x
			dep[tot] = dep[t.x] + 1
			tot++
		}
		if t.tp == 1 {
			fmt.Fscan(in, &t.z)
		}
		q2[i] = t
	}
	dfs(0)
	for i := 1; i <= q; i++ {
		q1 = append(q1, q2[i])
		if i%B == 0 {
			solve()
		}
	}
	if q%B != 0 {
		solve()
	}
}

func dfs(u int) {
	sz[u] = 1
	dfns++
	dfn[u] = dfns
	for _, v := range e[u] {
		dfs(v)
		sz[u] += sz[v]
	}
}

func chk(x, y int) bool {
	return dfn[x] <= dfn[y] && dfn[y] < dfn[x]+sz[x]
}

func qyc(l, r, d int) int {
	d = pd[d]
	return c[d][r] - c[d][l-1]
}

func qys(l, r, d int) int {
	d = pd[d]
	return s[d][r] - s[d][l-1]
}

func prep() {
	m := n
	xs, ds = 0, 0
	for it := range q1 {
		if q1[it].tp == 3 {
			a[n] = q1[it].y
			q1[it].z = n
			n++
		} else {
			u := q1[it].x
			xs++
			x[xs] = dfn[u]
			xs++
			x[xs] = dfn[u] + sz[u]
			ds++
			d[ds] = dep[u] + q1[it].y + 1
		}
	}
	tmpX := x[1 : xs+1]
	sort.Ints(tmpX)
	tmpX = x[1 : xs+1]
	xs = len(unique(tmpX))
	tmpD := d[1 : ds+1]
	sort.Ints(tmpD)
	tmpD = d[1 : ds+1]
	ds = len(unique(tmpD))
	for i, j := 1, 1; i <= tot; i++ {
		px[i] = px[i-1]
		if j <= xs && i == x[j] {
			px[i] = j
			j++
		}
	}
	for i, j := 1, 1; i <= tot; i++ {
		pd[i] = pd[i-1]
		if j <= ds && i == d[j] {
			pd[i] = j
			j++
		}
	}
	for i := 0; i < ds; i++ {
		for j := range c[i] {
			c[i][j] = 0
		}
		for j := range s[i] {
			s[i][j] = 0
		}
	}
	for i := 0; i < m; i++ {
		x := px[dfn[i]]
		y := pd[dep[i]]
		c[y][x]++
		s[y][x] += a[i]
	}
	for i := 0; i < ds; i++ {
		for j := 1; j < xs; j++ {
			c[i][j] += c[i][j-1]
		}
		if i != 0 {
			for j := 0; j < xs; j++ {
				c[i][j] += c[i-1][j]
			}
		}
	}
	for i := 0; i < ds; i++ {
		for j := 1; j < xs; j++ {
			s[i][j] += s[i][j-1]
		}
		if i != 0 {
			for j := 0; j < xs; j++ {
				s[i][j] += s[i-1][j]
			}
		}
	}
}

func unique(a []int) []int {
	occurred := map[int]bool{}
	result := []int{}
	for i := range a {
		if occurred[a[i]] != true {
			occurred[a[i]] = true
			result = append(result, a[i])
		}
	}
	sort.Ints(result)
	n := len(result)
	for i := 0; i < n; i++ {
		a[i] = result[i]
	}
	return result
}

func solve() {
	m := n
	prep()
	for i := 0; i < len(q1); i++ {
		t := q1[i]
		if t.tp == 3 {
			continue
		}
		u := t.x
		td := dep[u] + t.y
		pl[i] = px[dfn[u]]
		pr[i] = px[dfn[u]+sz[u]-1]
		if t.tp == 1 {
			for j := 0; j < i; j++ {
				if q1[j].tp == 3 {
					v := q1[j].z
					if chk(u, v) && dep[v] <= td {
						a[v] += t.z
					}
				}
			}
		} else {
			ans := qys(pl[i], pr[i], td)
			for j := 0; j < i; j++ {
				if q1[j].tp == 3 {
					v := q1[j].z
					if chk(u, v) && dep[v] <= td {
						ans += a[v]
					}
				}
				if q1[j].tp == 1 {
					v := q1[j].x
					dd := min(td, dep[v]+q1[j].y)
					if chk(u, v) && dep[v] <= dd {
						ans += q1[j].z * qyc(pl[j], pr[j], dd)
					} else if chk(v, u) && dep[u] <= dd {
						ans += q1[j].z * qyc(pl[i], pr[i], dd)
					}
				}
			}
			fmt.Fprintln(out, ans)
		}
	}
	for i := 0; i <= ds; i++ {
		for j := 0; j < xs+1; j++ {
			c[i][j] = 0
		}
	}
	for i := 0; i < len(q1); i++ {
		t := q1[i]
		if t.tp != 1 {
			continue
		}
		u := t.x
		dd := pd[dep[u]+t.y]
		c[dd][pl[i]] += t.z
		c[dd][pr[i]+1] -= t.z
	}
	for i := ds; i >= 0; i-- {
		for j := 1; j <= xs; j++ {
			c[i][j] += c[i][j-1]
		}
		if i < ds {
			for j := 0; j <= xs; j++ {
				c[i][j] += c[i+1][j]
			}
		}
	}
	for i := 0; i < m; i++ {
		x := px[dfn[i]]
		y := pd[dep[i]]
		a[i] += c[y][x]
	}
	q1 = make([]Que, 0)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
