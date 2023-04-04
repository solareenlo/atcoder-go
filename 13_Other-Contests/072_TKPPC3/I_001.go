package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const MAX = 400002
const MAX_LOG = 20
const LLONG_MAX = 9223372036854775807

type pair struct {
	x, y int
}

type P struct {
	x int
	y pair
}

var p, l []int
var n, m, ord, CUR_COL int
var pr, Len, star, en, dep, dist, col, im, outt [MAX]int
var v [][]int
var node, G [MAX][]int
var lca LCA
var add [MAX][]P

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &n, &m)
	pr[0] = -1
	for i := 1; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		a--
		pr[i] = a
		p = append(p, a)
	}
	for i := 1; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		l = append(l, a)
		Len[i] = a
	}
	v = make([][]int, n)
	for i := 0; i < len(p); i++ {
		v[p[i]] = append(v[p[i]], i+1)
	}
	lca.init(v, 0)
	dfs(0, 0)
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		a--
		node[a] = append(node[a], i)
		col[i] = a
	}
	for i := 0; i < m; i++ {
		node[i] = append(node[i], 0)
		CUR_COL = i
		sort.Slice(node[i], func(a, b int) bool {
			return star[node[i][a]] < star[node[i][b]]
		})
		for j := 0; j+1 < len(node[i]); j++ {
			if node[i][j] == node[i][j+1] {
				continue
			}
			tmp := lca.lca(node[i][j], node[i][j+1])
			node[i] = append(node[i], tmp)
		}
		sort.Slice(node[i], func(a, b int) bool {
			return star[node[i][a]] < star[node[i][b]]
		})
		node[i] = unique(node[i])
		for _, el := range node[i] {
			G[el] = make([]int, 0)
			dist[el] = LLONG_MAX
		}
		stk := make([]int, 0)
		for _, el := range node[i] {
			for len(stk) != 0 && !(star[stk[len(stk)-1]] <= star[el] && en[el] <= en[stk[len(stk)-1]]) {
				stk = stk[:len(stk)-1]
			}
			if len(stk) != 0 {
				G[stk[len(stk)-1]] = append(G[stk[len(stk)-1]], el)
			}
			stk = append(stk, el)
		}
		dfs2(0, LLONG_MAX)
		dfs2(0, LLONG_MAX)
		dfs3(0, -1)
	}
	dfs5(0)
	dfs4(0, -1, 0)
	for i := 0; i < n; i++ {
		fmt.Fprintln(out, outt[i])
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
	// sort.Ints(result)
	n := len(result)
	for i := 0; i < n; i++ {
		a[i] = result[i]
	}
	return result
}

func dfs(b, d int) {
	d += Len[b]
	star[b] = ord
	ord++
	dep[b] = d
	for _, Go := range v[b] {
		dfs(Go, d)
	}
	en[b] = ord
}

func dfs2(b, cur int) int {
	dist[b] = min(dist[b], cur)
	if col[b] == CUR_COL {
		dist[b] = 0
	}
	for _, Go := range G[b] {
		Len := dep[Go] - dep[b]
		nex := dist[b]
		if nex != LLONG_MAX {
			nex += Len
		}
		nex = dfs2(Go, nex)
		if nex != LLONG_MAX {
			nex += Len
		}
		dist[b] = min(dist[b], nex)
	}
	return dist[b]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func dfs3(b, pr int) {
	if b == 0 {
		add[b] = append(add[b], P{CUR_COL, pair{0, dist[b]}})
	}
	for _, Go := range G[b] {
		dfs3(Go, b)
	}
	if pr == -1 {
		return
	}
	up := b
	for i := MAX_LOG - 1; i >= 0; i-- {
		if lca.lcc[i][up] == -1 {
			continue
		}
		if dep[lca.lcc[i][up]] < dep[pr] {
			continue
		}
		inter := lca.lcc[i][up]
		from_dw := dist[b] + dep[b] - dep[inter]
		from_up := dist[pr] + dep[inter] - dep[pr]
		if from_up >= from_dw {
			up = inter
		}
	}
	add[up] = append(add[up], P{CUR_COL, pair{dist[pr] + dep[up] - dep[pr], dist[b] + dep[b] - dep[up]}})
	im[b] -= 2
	im[up] += 2
}

func dfs4(b, pr, cur_ans int) {
	if pr != -1 {
		cur_ans += (im[b] + m) * (dep[b] - dep[pr])
	}
	for i := 0; i < len(add[b]); i++ {
		cur_ans -= add[b][i].y.x
		cur_ans += add[b][i].y.y
	}
	outt[b] = cur_ans
	for _, Go := range v[b] {
		dfs4(Go, b, cur_ans)
	}
}

func dfs5(b int) {
	for _, Go := range v[b] {
		dfs5(Go)
		im[b] += im[Go]
	}
	return
}

type LCA struct {
	g       [][]int
	MAX_LOG int
	lcc     [][]int
	dep     []int
	myr     []int
	flag_r  int
}

func (l *LCA) init2() {
	n := len(l.g)
	for i := 0; i+1 < l.MAX_LOG; i++ {
		for j := 0; j < n; j++ {
			if l.lcc[i][j] == -1 {
				l.lcc[i+1][j] = -1
				continue
			}
			l.lcc[i+1][j] = l.lcc[i][l.lcc[i][j]]
		}
	}
}

func (l *LCA) lca(a, b int) int {
	if l.myr[a] != l.myr[b] {
		return -1
	}
	if l.dep[a] < l.dep[b] {
		a, b = b, a
	}
	for i := 0; i < l.MAX_LOG; i++ {
		if (((l.dep[a] - l.dep[b]) >> i) & 1) != 0 {
			a = l.lcc[i][a]
		}
	}
	if a == b {
		return a
	}
	for i := l.MAX_LOG - 1; i >= 0; i-- {
		if l.lcc[i][a] != l.lcc[i][b] {
			a = l.lcc[i][a]
			b = l.lcc[i][b]
		}
	}
	return l.lcc[0][a]
}

func (l *LCA) dfs(b, pr, d int) {
	for _, i := range l.g[b] {
		if i == pr {
			continue
		}
		l.dfs(i, b, d+1)
	}
	l.dep[b] = d
	l.lcc[0][b] = pr
	l.myr[b] = l.flag_r
}

func (l *LCA) init(tree [][]int, root int) {
	l.MAX_LOG = 20
	l.g = make([][]int, len(tree))
	for i := range l.g {
		l.g[i] = make([]int, len(tree[i]))
		copy(l.g[i], tree[i])
	}
	resize(&l.myr, len(tree), -1)
	l.flag_r = root
	resizeVV(&l.lcc, l.MAX_LOG, len(l.g))
	resize(&l.dep, len(l.g), 0)
	l.dfs(root, -1, 0)
	l.init2()
}

func resizeVV(a *[][]int, n int, val int) {
	if len(*a) > n {
		*a = (*a)[:n]
	} else {
		n = n - len(*a)
		for i := 0; i < n; i++ {
			tmp := make([]int, val)
			*a = append(*a, tmp)
		}
	}
}

func resize(a *[]int, n int, val int) {
	if len(*a) > n {
		*a = (*a)[:n]
	} else {
		n = n - len(*a)
		for i := 0; i < n; i++ {
			*a = append(*a, val)
		}
	}
}
