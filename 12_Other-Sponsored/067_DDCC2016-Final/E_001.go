package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type pair struct {
	x, y int
}

const MX = 110000
const MX2 = 262144

var g [MX][]pair
var conv, v, fi, dep, dn, last [MX]int
var eu [MX * 2]int
var segtree [MX * 5]int
var num, ind int

func dfs(a, b int) {
	v[num] = a
	num++
	conv[a] = num - 1
	eu[ind] = num - 1
	ind++
	fi[conv[a]] = ind - 1
	for i := 0; i < len(g[a]); i++ {
		if b == g[a][i].x {
			continue
		}
		dep[g[a][i].x] = dep[a] + g[a][i].y
		dn[g[a][i].x] = dn[a] + 1
		dfs(g[a][i].x, a)
		eu[ind] = conv[a]
		ind++
	}
	last[conv[a]] = num - 1
}

func query(a, b, c, d, e int) int {
	if d < a || b < c {
		return 999999999
	}
	if c <= a && b <= d {
		return segtree[e]
	}
	return min(query(a, (a+b)/2, c, d, e*2), query((a+b)/2+1, b, c, d, e*2+1))
}

func update(a, b int) {
	a += MX2
	for a != 0 {
		segtree[a] = min(segtree[a], b)
		a /= 2
	}
}

var st, bit [MX2][]int

func add(a, b, c int) {
	for ; b < len(bit[a]); b |= b + 1 {
		bit[a][b] += c
	}
}

func sum(a, b, c int) int {
	if b != 0 {
		return sum(a, 0, c) - sum(a, 0, b-1)
	}
	ret := 0
	for ; c >= 0; c = (c & (c + 1)) - 1 {
		ret += bit[a][c]
	}
	return ret
}

func add1(a, b, c, d, e, f int) {
	if d < a || b < c {
		return
	}
	if c <= a && b <= d {
		st[e] = append(st[e], f)
		return
	}
	add1(a, (a+b)/2, c, d, e*2, f)
	add1((a+b)/2+1, b, c, d, e*2+1, f)
}

func add2(a, b, c, d, e, f, h int) {
	if d < a || b < c {
		return
	}
	if c <= a && b <= d {
		at := lowerBound(st[e], f)
		add(e, at, h)
		return
	}
	add2(a, (a+b)/2, c, d, e*2, f, h)
	add2((a+b)/2+1, b, c, d, e*2+1, f, h)
}

var qu [110000][4]int

func geta(a int) int {
	b := conv[a]
	b += 131072
	ret := 0
	for b != 0 {
		if len(st[b]) == 0 {
			b /= 2
			continue
		}
		at := upperBound(st[b], dn[a]) - 1
		if at >= 0 {
			ret += sum(b, 0, at)
		}
		b /= 2
	}
	return ret
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var a int
	fmt.Fscan(in, &a)
	for i := 0; i < a-1; i++ {
		var p, q, r int
		fmt.Fscan(in, &p, &q, &r)
		p--
		q--
		g[p] = append(g[p], pair{q, r})
		g[q] = append(g[q], pair{p, r})
	}
	dfs(0, -1)
	for i := 0; i < 1<<19; i++ {
		segtree[i] = 999999999
	}
	for i := 0; i < ind; i++ {
		update(i, eu[i])
	}
	var b int
	fmt.Fscan(in, &b)
	for i := 0; i < b; i++ {
		fmt.Fscan(in, &qu[i][0])
		if qu[i][0] == 1 {
			fmt.Fscan(in, &qu[i][1], &qu[i][2], &qu[i][3])
		} else {
			fmt.Fscan(in, &qu[i][1], &qu[i][2])
			qu[i][2]--
		}
		qu[i][1]--
	}
	for i := 0; i < b; i++ {
		if qu[i][0] == 1 {
			add1(0, 131071, conv[qu[i][1]], last[conv[qu[i][1]]], 1, dn[qu[i][1]]+qu[i][2])
		}
	}
	for i := 0; i < MX2; i++ {
		sort.Ints(st[i])
	}
	for i := 0; i < MX2; i++ {
		bit[i] = make([]int, len(st[i]))
	}

	for i := 0; i < b; i++ {
		if qu[i][0] == 1 {
			add2(0, 131071, conv[qu[i][1]], last[conv[qu[i][1]]], 1, dn[qu[i][1]]+qu[i][2], qu[i][3])
		} else {
			f1 := conv[qu[i][1]]
			f2 := conv[qu[i][2]]
			h1 := fi[f1]
			h2 := fi[f2]
			if h1 > h2 {
				h1, h2 = h2, h1
			}
			lca := v[query(0, 262143, h1, h2, 1)]
			ret := dep[qu[i][1]] + dep[qu[i][2]] - dep[lca]*2
			ret += geta(qu[i][1])
			ret += geta(qu[i][2])
			ret -= geta(lca) * 2

			fmt.Println(ret)
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}
