package main

import (
	"bufio"
	"fmt"
	"os"
)

type edge struct {
	A, B, C int
}

type pair struct {
	x, y int
}

type Pair struct {
	x pair
	y int
}

var pr, sz, ans [1 << 17]int
var K, sum int

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(in, &N, &M, &K)
	for i := 0; i < N; i++ {
		pr[i] = i
		sz[i] = 1
	}
	E := make([]edge, M)
	for i := 0; i < M; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		E[i] = edge{a - 1, b - 1, c}
	}
	var Q int
	fmt.Fscan(in, &Q)
	qs := make([]pair, Q)
	for i := 0; i < Q; i++ {
		var d int
		fmt.Fscan(in, &d)
		qs[i] = pair{d, i}
	}
	Dfs(E, qs, 29)
	for i := 0; i < Q; i++ {
		fmt.Println(ans[i] / 2)
	}
}

func Dfs(E []edge, qs []pair, k int) {
	if k < 0 {
		for _, q := range qs {
			ans[q.y] = sum
		}
		return
	}
	qs0 := make([]pair, 0)
	qs1 := make([]pair, 0)
	for _, q := range qs {
		if ((q.x >> k) & 1) != 0 {
			qs1 = append(qs1, q)
		} else {
			qs0 = append(qs0, q)
		}
	}
	if len(qs0) != 0 {
		nE := make([]edge, 0)
		use := 0
		for _, e := range E {
			if (e.C >> k & 1) == (K >> k & 1) {
				nE = append(nE, e)
			} else if (K >> k & 1) != 0 {
				unite(e.A, e.B)
				use++
			}
		}
		Dfs(nE, qs0, k-1)
		for use > 0 {
			use--
			undo()
		}
	}
	if len(qs1) != 0 {
		nE := make([]edge, 0)
		use := 0
		for _, e := range E {
			if (e.C >> k & 1) != (K >> k & 1) {
				nE = append(nE, e)
			} else if ((e.C >> k) & 1) != 0 {
				unite(e.A, e.B)
				use++
			}
		}
		Dfs(nE, qs1, k-1)
		for use > 0 {
			use--
			undo()
		}
	}
}

func find(u int) int {
	if u == pr[u] {
		return u
	}
	return find(pr[u])
}

var pre []Pair

func unite(u, v int) {
	u = find(u)
	v = find(v)
	if u == v {
		pre = append(pre, Pair{pair{-1, -1}, -1})
	} else {
		if sz[u] < sz[v] {
			u, v = v, u
		}
		pre = append(pre, Pair{pair{u, v}, pr[v]})
		sum -= sz[u] * (sz[u] - 1)
		sum -= sz[v] * (sz[v] - 1)
		sz[u] += sz[v]
		pr[v] = u
		sum += sz[u] * (sz[u] - 1)
	}
}

func undo() {
	u := pre[len(pre)-1].x.x
	v := pre[len(pre)-1].x.y
	p := pre[len(pre)-1].y
	pre = pre[:len(pre)-1]
	if p >= 0 {
		pr[v] = p
		sum -= sz[u] * (sz[u] - 1)
		sz[u] -= sz[v]
		sum += sz[u] * (sz[u] - 1)
		sum += sz[v] * (sz[v] - 1)
	}
}
