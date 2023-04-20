package main

import (
	"bufio"
	"fmt"
	"os"
)

const MAX_N = 1 << 20
const MAX_B = 1 << 17

var N int
var L, R, D, I, ANS [1 << 18]int
var E, X [MAX_B * 2][]int
var Z LiChaoTree

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &N)
	for i := 1; i <= N; i++ {
		fmt.Fscan(in, &L[i], &R[i], &D[i], &I[i])
		if R[i] == -1 {
			R[i] = MAX_B - 1
		}
	}
	for i := 1; i <= N; i++ {
		add(L[i], R[i]+1, i, 0, MAX_B, 1)
	}
	for i := 1; i <= N; i++ {
		E[L[i]] = append(E[L[i]], i)
	}
	Z.init()
	Dfs(0, MAX_B, 1, 0)
	for i := 1; i <= N; i++ {
		fmt.Fprintln(out, ANS[i])
	}
}

func Dfs(a, b, u, dep int) {
	for i := 0; i < len(X[u]); i++ {
		pos := X[u][i]
		Z.add_edge(D[pos], ANS[pos], dep)
	}
	if b-a == 1 {
		for i := 0; i < len(E[a+1]); i++ {
			pos := E[a+1][i]
			ANS[pos] = min(Z.getmin(D[pos]), I[pos])
		}
	}
	if b-a >= 2 {
		Dfs(a, (a+b)>>1, u*2, dep+1)
		Dfs((a+b)>>1, b, u*2+1, dep+1)
	}
	Z.all_erase(dep)
}

func add(l, r, id, a, b, u int) {
	if l <= a && b <= r {
		X[u] = append(X[u], id)
		return
	}
	if r <= a || b <= l {
		return
	}

	add(l, r, id, a, (a+b)>>1, u*2)
	add(l, r, id, (a+b)>>1, b, u*2+1)
}

type Node struct {
	a, b int // y = ax + b
}

type Pair struct {
	a int
	b Node
}

type LiChaoTree struct {
	F []Node
	G [25][]Pair
}

func (lct *LiChaoTree) init() {
	lct.F = make([]Node, MAX_N*2)
	for i := 0; i < MAX_N*2; i++ {
		lct.F[i] = Node{0, 1 << 60}
	}
}

func (lct *LiChaoTree) all_erase(dep int) {
	for i := len(lct.G[dep]) - 1; i >= 0; i-- {
		lct.F[lct.G[dep][i].a] = lct.G[dep][i].b
	}
	lct.G[dep] = make([]Pair, 0)
}

func (lct *LiChaoTree) add_edge(a, b, dep int) {
	// 座標 (a, b) が頂点
	E := Node{-2 * a, b + a*a}
	cl := 0
	cr := MAX_N
	u := 1
	for {
		cm := ((cl + cr) >> 1)
		dl := ((cl*E.a + E.b) < (cl*lct.F[u].a + lct.F[u].b))
		dm := ((cm*E.a + E.b) < (cm*lct.F[u].a + lct.F[u].b))
		dr := ((cr*E.a + E.b) < (cr*lct.F[u].a + lct.F[u].b))
		if dl == true && dr == true {
			lct.G[dep] = append(lct.G[dep], Pair{u, lct.F[u]})
			lct.F[u] = E
			break
		} else if dl == false && dr == false {
			// 何も起こらない
			break
		} else if dm == true {
			lct.G[dep] = append(lct.G[dep], Pair{u, lct.F[u]})
			lct.F[u], E = E, lct.F[u]
		}

		r := cr - cl
		if dr != dm {
			u = u*2 + 1
			cl = cm
		} else {
			u = u * 2
			cr = cm
		}

		if r == 1 {
			break
		}
	}
}

func (lct LiChaoTree) getmin(pos int) int {
	u := pos + MAX_N
	minx := (1 << 60)
	for u >= 1 {
		minx = min(minx, pos*lct.F[u].a+lct.F[u].b)
		u >>= 1
	}
	return pos*pos + minx
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
