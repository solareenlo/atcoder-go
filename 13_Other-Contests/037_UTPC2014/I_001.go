package main

import (
	"bufio"
	"fmt"
	"os"
)

const MAXN = 100011

var N, Q int
var eA, eB, eC [MAXN]int
var G [MAXN][]int
var ord, pos []int
var subsz, depth [MAXN]int
var par [MAXN][20]int
var X *SegmentTree

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &N)
	for i := range G {
		G[i] = make([]int, 0)
	}
	for i := 0; i < N-1; i++ {
		fmt.Fscan(in, &eA[i], &eB[i], &eC[i])
		eA[i]--
		eB[i]--
		G[eA[i]] = append(G[eA[i]], eB[i])
		G[eB[i]] = append(G[eB[i]], eA[i])
	}

	ord = make([]int, 0)
	pos = make([]int, 0)
	DFS()

	X = NewSegmentTree(N)

	for i := 0; i < N-1; i++ {
		X.add(pos[getV(i)], 0, eC[i], i)
	}
	X.add(0, 0, 0, -9)

	fmt.Fscan(in, &Q)
	for j := 0; j < Q; j++ {
		var op int
		fmt.Fscan(in, &op)
		if op == 1 {
			var I, D int
			fmt.Fscan(in, &I, &D)
			I--
			v := getV(I)
			X.add(pos[v], 0, D, 0)
		} else {
			var I int
			fmt.Fscan(in, &I)
			I--
			v := getV(I)
			R := findroot(I)
			if v == R {
				fmt.Println(-1)
				continue
			}
			s := X.min_s1(pos[R]+1, pos[R]+subsz[R])
			fmt.Println(s.id + 1)
			v = getV(s.id)
			X.cut1(pos[v], pos[v]+subsz[v], 1)
		}
	}
}

func DFS() {
	V := make([]int, 0)
	P := make([]int, 0)
	V = append(V, 0)
	P = append(P, 0)
	for len(V) > 0 {
		v := V[len(V)-1]
		V = V[:len(V)-1]
		p := P[len(P)-1]
		P = P[:len(P)-1]
		ord = append(ord, v)
		for _, e := range G[v] {
			if e == p {
				continue
			}
			V = append(V, e)
			P = append(P, v)
			depth[e] = depth[v] + 1
			par[e][0] = v
		}
	}

	ord = append(ord, 0)
	for i := N; i >= 0; i-- {
		v := ord[i]
		subsz[v] = 1
		for _, e := range G[v] {
			if depth[e] == depth[v]+1 {
				subsz[v] += subsz[e]
			}
		}
	}
	for i := 0; i < 19; i++ {
		for v := 0; v < N; v++ {
			par[v][i+1] = par[par[v][i]][i]
		}
	}

	pos = make([]int, N)
	for i := 0; i < N; i++ {
		pos[i] = -1
	}
	for i := 0; i < N; i++ {
		pos[ord[i]] = i
	}
}

func getV(I int) int {
	if depth[eA[I]] < depth[eB[I]] {
		return eB[I]
	}
	return eA[I]
}

func findroot(I int) int {
	v := getV(I)
	r := v
	kill := X.min_s1(pos[v], pos[v]+1).kill
	for i := 19; i >= 0; i-- {
		p := par[r][i]
		if X.min_s1(pos[p], pos[p]+1).kill == kill {
			r = p
		}
	}
	return r
}

type SegmentTree struct {
	n, m int
	data []Seg
	all  []int
}

func NewSegmentTree(n int) *SegmentTree {
	seg := new(SegmentTree)
	seg.n = n
	seg.m = 1
	for seg.m < n {
		seg.m *= 2
	}
	seg.all = make([]int, seg.m*2)
	seg.data = make([]Seg, seg.m*2)
	return seg
}

func (seg *SegmentTree) cut1(x, y, v int) {
	seg.cut(x, y, 1, 0, seg.m, v)
}

func (seg *SegmentTree) cut(x, y, k, l, r, v int) {
	if x <= l && r <= y {
		seg.all[k] += v
		return
	} else if x < r && l < y {
		seg.cut(x, y, k*2, l, (l+r)/2, v)
		seg.cut(x, y, k*2+1, (l+r)/2, r, v)
		seg.data[k] = Min(seg.data[k*2].KK(seg.all[k*2]), seg.data[k*2+1].KK(seg.all[k*2+1]))
	}
}

func (seg *SegmentTree) add(x, k, d, i int) {
	x += seg.m
	seg.data[x].kill += k
	seg.data[x].cost += d
	seg.data[x].id += i
	for x := x / 2; x > 0; x /= 2 {
		seg.data[x] = Min(seg.data[x*2].KK(seg.all[x*2]), seg.data[x*2+1].KK(seg.all[x*2+1]))
	}
}

func (seg SegmentTree) get(x int) Seg {
	return seg.min_s1(x, x+1)
}

func (seg SegmentTree) min_s1(x, y int) Seg {
	return seg.min_s(x, y, 1, 0, seg.m)
}

func (seg SegmentTree) min_s(x, y, k, l, r int) Seg {
	if r <= x || y <= l {
		return Seg{1 << 29, 0, 0}
	}
	if x <= l && r <= y {
		return seg.data[k].KK(seg.all[k])
	}
	return Min(seg.min_s(x, y, k*2, l, (l+r)/2), seg.min_s(x, y, k*2+1, (l+r)/2, r)).KK(seg.all[k])
}

type Seg struct {
	kill int
	cost int
	id   int
}

func (x Seg) lessThan(y Seg) bool {
	if x.kill != y.kill {
		return x.kill < y.kill
	}
	if x.cost != y.cost {
		return x.cost < y.cost
	}
	return x.id < y.id
}

func (x Seg) KK(a int) Seg {
	return Seg{x.kill + a, x.cost, x.id}
}

func Min(a, b Seg) Seg {
	if a.lessThan(b) {
		return a
	}
	return b
}
