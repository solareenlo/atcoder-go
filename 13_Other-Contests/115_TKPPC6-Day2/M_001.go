package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	type pair struct {
		x, y int
	}

	var N int
	fmt.Fscan(in, &N)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		var l, r int
		fmt.Fscan(in, &l, &r)
		A[i] = r - l + 1
	}

	var G LiChaoTree
	G.init(N)

	var Q int
	fmt.Fscan(in, &Q)
	queries := make([][]pair, N)
	for i := 0; i < Q; i++ {
		var l, r int
		fmt.Fscan(in, &l, &r)
		queries[r-1] = append(queries[r-1], pair{l - 1, i})
	}

	ans := make([]int, Q)
	for i := 0; i < N; i++ {
		G.addLine(Func{i + 1, A[i]})
		for _, q := range queries[i] {
			res := G.minVal(q.x)
			ans[q.y] = -(res.b / res.a)
		}
	}
	for i := 0; i < Q; i++ {
		fmt.Fprintln(out, ans[i])
	}
}

type Fraq struct {
	a, b int
}

func lessThan(l, r Fraq) bool {
	if l.b == 0 && r.b == 0 {
		return l.b < r.b
	}
	return l.a*r.b < l.b*r.a
}

type Func struct {
	i, a int
}

func (f Func) op(x int) Fraq {
	return Fraq{x - f.i, f.a}
}

var infFunc Func = Func{-100, 1000000000000}

type LiChaoTree struct {
	N       int
	V       []Func
	visited []bool
}

func (lct *LiChaoTree) init(n int) {
	lct.N = 1
	for lct.N < n {
		lct.N *= 2
	}
	lct.V = make([]Func, lct.N*2)
	for i := range lct.V {
		lct.V[i] = infFunc
	}
	lct.visited = make([]bool, lct.N*2)
}

func (lct *LiChaoTree) addSegment(l, r int, f Func) {
	if l >= r {
		return
	}
	var dfs func(int, Func, int, int)
	dfs = func(i int, f Func, a, b int) {
		lct.visited[i] = true
		if i >= len(lct.V) {
			return
		}
		if r <= a || b <= l {
			return
		}
		m := (a + b) / 2
		if !(l <= a && b <= r) {
			dfs(i*2, f, a, m)
			dfs(i*2+1, f, m, b)
			return
		}
		greatf_l := !lessThan(lct.V[i].op(a), f.op(a))
		greatf_r := !lessThan(lct.V[i].op(b-1), f.op(b-1))
		if !greatf_l && !greatf_r {
			return
		}
		if greatf_l && greatf_r {
			lct.V[i] = f
			return
		}
		if a+1 == b {
			return
		}
		if !lessThan(lct.V[i].op(m), f.op(m)) {
			lct.V[i], f = f, lct.V[i]
			greatf_l, greatf_r = greatf_r, greatf_l
		}
		if greatf_l {
			dfs(i*2, f, a, m)
		} else {
			dfs(i*2+1, f, m, b)
		}
	}
	dfs(1, f, 0, lct.N)
}

func (lct *LiChaoTree) addLine(f Func) {
	lct.addSegment(0, lct.N, f)
}

func (lct LiChaoTree) minVal(p int) Fraq {
	i := 1
	res := infFunc.op(p)
	l := 0
	r := lct.N
	for i < len(lct.V) {
		if !lct.visited[i] {
			break
		}
		if !lessThan(res, lct.V[i].op(p)) {
			res = lct.V[i].op(p)
		}
		m := (l + r) / 2
		if p < m {
			i = i * 2
			r = m
		} else {
			i = i*2 + 1
			l = m
		}
	}
	return res
}

func (lct LiChaoTree) minFunc(p int) Func {
	i := 1
	res := infFunc
	l := 0
	r := lct.N
	for i < len(lct.V) {
		if !lct.visited[i] {
			break
		}
		if lessThan(lct.V[i].op(p), res.op(p)) {
			res = lct.V[i]
		}
		m := (l + r) / 2
		if p < m {
			i = i * 2
			r = m
		} else {
			i = i*2 + 1
			l = m
		}
	}
	return res
}
