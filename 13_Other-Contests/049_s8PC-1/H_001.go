package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, d, e int
	fmt.Scan(&n, &d, &e)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	type Info struct {
		diff0 int
		diff1 int
		loss  int
	}

	type pair struct {
		x int
		y int
	}

	type BInfo struct {
		x        int
		y        int
		y_comped int
	}

	var zen_rekkyo func(int, int) []Info
	zen_rekkyo = func(l, r int) []Info {
		cur := make([]Info, 1)
		cur[0] = Info{0, 0, 0}
		for i := l; i < r; i++ {
			next := make([]Info, 0)
			for _, j := range cur {
				next = append(next, Info{j.diff0 + a[i], j.diff1, j.loss})
				next = append(next, Info{j.diff0 - a[i], j.diff1 + a[i], j.loss})
				next = append(next, Info{j.diff0, j.diff1 - a[i], j.loss})
				if j.loss < e {
					next = append(next, Info{j.diff0, j.diff1, j.loss + 1})
				}
			}
			cur, next = next, cur
		}
		return cur
	}

	res := 0
	var solve func([]pair, []BInfo, []int)
	solve = func(a []pair, b []BInfo, b_comp []int) {
		head := 0
		tail := 0
		m := len(b)
		tree := NewSegTree(m)
		for _, i := range a {
			for head < m && b[head].x <= d-i.x {
				tree.add(b[head].y_comped, 1)
				head++
			}
			for tail < m && b[tail].x < -d-i.x {
				tree.add(b[tail].y_comped, -1)
				tail++
			}
			res += tree.sum(lowerBound(b_comp, -d-i.y), lowerBound(b_comp, d-i.y+1))
		}
	}

	half := n / 2
	r0 := make([][]pair, e+1)
	r1 := make([][]BInfo, e+1)
	r1_comp := make([][]int, e+1)
	for _, i := range zen_rekkyo(0, half) {
		r0[i.loss] = append(r0[i.loss], pair{i.diff0, i.diff1})
	}
	for _, i := range zen_rekkyo(half, n) {
		r1[i.loss] = append(r1[i.loss], BInfo{i.diff0, i.diff1, 0})
	}
	for _, tmp := range r0 {
		sort.Slice(tmp, func(i, j int) bool {
			if tmp[i].x == tmp[j].x {
				return tmp[i].y > tmp[j].y
			}
			return tmp[i].x > tmp[j].x
		})
	}
	for i := 0; i <= e; i++ {
		sort.Slice(r1[i], func(a, b int) bool {
			return r1[i][a].y < r1[i][b].y
		})
		for j := 0; j < len(r1[i]); {
			k := j
			for ; k < len(r1[i]) && r1[i][k].y == r1[i][j].y; k++ {

			}
			for l := j; l < k; l++ {
				r1[i][l].y_comped = j
			}
			j = k
		}
		r1_comp[i] = make([]int, len(r1[i]))
		for j := 0; j < len(r1[i]); j++ {
			r1_comp[i][j] = r1[i][j].y
		}
		sort.Slice(r1[i], func(a, b int) bool {
			return r1[i][a].x < r1[i][b].x
		})
	}
	for i := 0; i <= e; i++ {
		for j := 0; j <= e-i; j++ {
			solve(r0[i], r1[j], r1_comp[j])
		}
	}
	fmt.Println(res)
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

type SegTree struct {
	n    int
	data []int
}

func NewSegTree(n_ int) *SegTree {
	seg := new(SegTree)
	seg.n = 1
	for seg.n < n_ {
		seg.n <<= 1
	}
	seg.data = make([]int, seg.n<<1)
	return seg
}

func (seg *SegTree) add(i, val int) {
	for i += seg.n; i > 0; i >>= 1 {
		seg.data[i] += val
	}
}

func (seg *SegTree) sum(l, r int) int {
	res := 0
	for l, r = l+seg.n, r+seg.n; l < r; l, r = l>>1, r>>1 {
		if (r & 1) != 0 {
			r--
			res += seg.data[r]
		}
		if (l & 1) != 0 {
			res += seg.data[l]
			l++
		}
	}
	return res
}
