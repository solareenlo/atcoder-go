package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const SIZE = 100005

	type P struct {
		x, y int
	}
	type PP struct {
		x, y P
	}

	var n int
	fmt.Fscan(in, &n)
	mx := 0
	vec := make([]string, 0)
	Rank := make([][]int, 0)
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(in, &s)
		l := len(s)
		vec = append(vec, s)
		r := make([]int, 0)
		for j := 0; j < l; j++ {
			r = append(r, int(s[j]-'a'))
		}
		Rank = append(Rank, r)
		mx = max(mx, l)
	}
	if n == 1 {
		fmt.Println(mx)
		return
	}
	tmp := make([]PP, SIZE)
	for k := 1; k <= mx; k <<= 1 {
		sz := 0
		for i := 0; i < n; i++ {
			for j := 0; j < len(Rank[i]); j++ {
				if j+k < len(Rank[i]) {
					tmp[sz] = PP{P{Rank[i][j], Rank[i][j+k]}, P{i, j}}
				} else {
					tmp[sz] = PP{P{Rank[i][j], -1}, P{i, j}}
				}
				sz++
			}
		}
		tmp1 := tmp[:sz]
		sort.Slice(tmp1, func(i, j int) bool {
			if tmp1[i].x == tmp1[j].x {
				if tmp1[i].y.x == tmp1[j].y.x {
					return tmp1[i].y.y < tmp1[j].y.y
				}
				return tmp1[i].y.x < tmp1[j].y.x
			}
			if tmp1[i].x.x == tmp1[j].x.x {
				return tmp1[i].x.y < tmp1[j].x.y
			}
			return tmp1[i].x.x < tmp1[j].x.x
		})
		for i := 0; i < sz; {
			if 2*k <= mx {
				p := tmp[i]
				s := i
				for ; i < sz && tmp[i].x == p.x; i++ {
					q := tmp[i].y
					Rank[q.x][q.y] = s
				}
			} else {
				q := tmp[i].y
				Rank[q.x][q.y] = i
				i++
			}
		}
	}
	sz := 0
	sa := make([]P, SIZE)
	for i := 0; i < n; i++ {
		for j := 0; j < len(Rank[i]); j++ {
			sa[Rank[i][j]] = P{i, j}
			sz++
		}
	}
	lcp := make([]int, SIZE)
	bf := make([]int, SIZE)
	for i := 0; i < n; i++ {
		H := 0
		for j := 0; j < len(Rank[i]); j++ {
			if H > 0 {
				H--
			}
			if Rank[i][j] == sz-1 {
				continue
			}
			ti := sa[Rank[i][j]+1].x
			tj := sa[Rank[i][j]+1].y
			for j+H < len(Rank[i]) && tj+H < len(Rank[ti]) && vec[i][j+H] == vec[ti][tj+H] {
				H++
			}
			lcp[Rank[i][j]] = H
		}
		bf[i] = -1
	}
	lcp[sz-1] = 0
	ret := 0
	now := 0
	st := make([]P, SIZE)
	seg := NewSegTree(sz + 2)
	for i := 0; i < sz; i++ {
		last := i
		c := sa[i].x
		if bf[c] != -1 {
			seg.add(bf[c], -1)
		}
		seg.add(i, 1)
		bf[c] = i
		for now > 0 && st[now-1].x >= lcp[i] {
			now--
			p := st[now]
			last = p.y
			ret = max(ret, seg.get(last, i+1)*p.x)
		}
		if lcp[i] > 0 {
			st[now] = P{lcp[i], last}
			now++
		}
	}
	for i := 0; i < n; i++ {
		ret = max(ret, len(Rank[i]))
	}
	fmt.Println(ret)
}

const BT = 1024 * 128 * 2

type segTree struct {
	seg []int
	mum int
}

func NewSegTree(n int) *segTree {
	s := new(segTree)
	s.seg = make([]int, BT)
	s.mum = 1
	for s.mum < n {
		s.mum <<= 1
	}
	for i := 0; i < s.mum*2; i++ {
		s.seg[i] = 0
	}
	return s
}

func (s *segTree) add(k, x int) {
	k += s.mum - 1
	s.seg[k] += x
	for k > 0 {
		k = (k - 1) / 2
		s.seg[k] = s.seg[k*2+1] + s.seg[k*2+2]
	}
}

func (s segTree) get1(a, b, k, l, r int) int {
	if r <= a || b <= l {
		return 0
	}
	if a <= l && r <= b {
		return s.seg[k]
	} else {
		vl := s.get1(a, b, k*2+1, l, (l+r)/2)
		vr := s.get1(a, b, k*2+2, (l+r)/2, r)
		return vl + vr
	}
}

func (s segTree) get(a, b int) int {
	return s.get1(a, b, 0, 0, s.mum)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
