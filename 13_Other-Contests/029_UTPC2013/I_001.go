package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const INF = 2100000000

type pair struct {
	x, y int
}

var child [][]int
var sz, ret, ord, A []int
var seg *segTree

func main() {
	in := bufio.NewReader(os.Stdin)

	const MAXN = 100100

	var N int
	fmt.Fscan(in, &N)
	A = make([]int, MAXN)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
	}

	sorter := make([]pair, 0)
	for i := 0; i < N; i++ {
		sorter = append(sorter, pair{A[i], i})
	}
	sort.Slice(sorter, func(i, j int) bool {
		if sorter[i].x == sorter[j].x {
			return sorter[i].y < sorter[j].y
		}
		return sorter[i].x < sorter[j].x
	})

	ord = make([]int, MAXN)
	for i := 0; i < N; i++ {
		ord[sorter[i].y] = i
	}

	root := make([]int, MAXN)
	child = make([][]int, MAXN)
	for i := range child {
		child[i] = make([]int, 0)
	}
	for i := 0; i < N; i++ {
		root[i] = -1
	}
	for i := 0; i < N-1; i++ {
		var s, t int
		fmt.Fscan(in, &s, &t)
		root[t] = s
		child[s] = append(child[s], t)
	}

	r := -1
	for i := 0; i < N; i++ {
		if root[i] == -1 {
			r = i
			break
		}
	}

	sz = make([]int, MAXN)
	calc_size(r)

	ret = make([]int, MAXN)
	seg = newSegTree()
	solve(r)

	for i := 0; i < N; i++ {
		fmt.Println(ret[i])
	}
}

func calc_size(p int) int {
	tmp := 1
	for i := 0; i < len(child[p]); i++ {
		tmp += calc_size(child[p][i])
	}
	sz[p] = tmp
	return sz[p]
}

func solve(p int) {
	cs := make([]pair, 0)
	for i := 0; i < len(child[p]); i++ {
		cs = append(cs, pair{sz[child[p][i]], child[p][i]})
	}
	sort.Slice(cs, func(i, j int) bool {
		if cs[i].x == cs[j].x {
			return cs[i].y < cs[j].y
		}
		return cs[i].x < cs[j].x
	})
	cs = reverseOrder(cs)

	ret[p] = find_nearest(p)

	for i := 1; i < len(cs); i++ {
		append_rec(cs[i].y)
	}

	for i := 0; i < len(cs); i++ {
		if i != 0 {
			reset_rec(cs[i].y)
		}
		solve(cs[i].y)
	}

	seg.set(ord[p], A[p])
}

func reverseOrder(a []pair) []pair {
	n := len(a)
	res := make([]pair, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func find_nearest(p int) int {
	c1 := seg.max_query(0, ord[p])
	c2 := seg.min_query(ord[p]+1, seg.N)

	if c1 == -1 && c2 == INF {
		return -1
	}
	if c1 == -1 {
		return c2
	}
	if c2 == INF {
		return c1
	}
	if (A[p] - c1) > (c2 - A[p]) {
		return c2
	}
	return c1
}

func append_rec(p int) {
	seg.set(ord[p], A[p])
	for i := 0; i < len(child[p]); i++ {
		append_rec(child[p][i])
	}
}

func reset_rec(p int) {
	seg.reset(ord[p])
	for i := 0; i < len(child[p]); i++ {
		reset_rec(child[p][i])
	}
}

const DEPTH = 17

type segTree struct {
	mxq []int
	mnq []int
	N   int
}

func newSegTree() *segTree {
	seg := new(segTree)
	seg.N = 1 << DEPTH
	seg.mxq = make([]int, 2*seg.N)
	seg.mnq = make([]int, 2*seg.N)
	for i := 0; i < 2*seg.N; i++ {
		seg.mxq[i] = -1
		seg.mnq[i] = INF
	}
	return seg
}

func (seg *segTree) set(p, v int) {
	p += seg.N
	seg.mxq[p] = v
	seg.mnq[p] = v
	p >>= 1
	for p != 0 {
		seg.mxq[p] = max(seg.mxq[p*2], seg.mxq[p*2+1])
		seg.mnq[p] = min(seg.mnq[p*2], seg.mnq[p*2+1])
		p >>= 1
	}
}

func (seg *segTree) reset(p int) {
	p += seg.N
	seg.mxq[p] = -1
	seg.mnq[p] = INF
	p >>= 1
	for p != 0 {
		seg.mxq[p] = max(seg.mxq[p*2], seg.mxq[p*2+1])
		seg.mnq[p] = min(seg.mnq[p*2], seg.mnq[p*2+1])
		p >>= 1
	}
}

func (seg segTree) max_query(left, right int) int {
	left += seg.N
	right += seg.N
	ret := -1
	for left < right {
		if (left & 1) != 0 {
			ret = max(ret, seg.mxq[left])
			left++
		}
		if (right & 1) != 0 {
			right--
			ret = max(ret, seg.mxq[right])
		}
		left >>= 1
		right >>= 1
	}
	return ret
}

func (seg segTree) min_query(left, right int) int {
	left += seg.N
	right += seg.N
	ret := INF
	for left < right {
		if (left & 1) != 0 {
			ret = min(ret, seg.mnq[left])
			left++
		}
		if (right & 1) != 0 {
			right--
			ret = min(ret, seg.mnq[right])
		}
		left >>= 1
		right >>= 1
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
