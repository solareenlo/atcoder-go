package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const NN = 101000

var n int
var nd [4 * NN]node

var a, b, vv [NN]int
var v []int

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		v = append(v, a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &b[i])
	}
	sort.Ints(v)
	v = Unique(v)
	m := len(v)
	build(1, 0, m-1)
	for i := 0; i < n; i++ {
		r := LowerBound(v, a[i])
		vr := queryv(1, 0, m-1, r)
		s := query(1, 0, m-1, vr+b[i])
		if r < s {
			modify(1, 0, m-1, r, s-1, tag{0, 0, vr + b[i]})
		}
		if r > 0 {
			modify(1, 0, m-1, 0, r-1, tag{1, 1, -a[i] + b[i]})
		}
	}
	fmt.Println(nd[1].rval)
}

func build(p, l, r int) {
	nd[p].t = tag{1, 0, 0}
	nd[p].rh = v[r]
	if l == r {
	} else {
		md := (l + r) >> 1
		build(p+p, l, md)
		build(p+p+1, md+1, r)
		upd(p)
	}
}

func upd(p int) {
	nd[p].rval = nd[p+p+1].rval
}

func queryv(p, l, r, x int) int {
	if l == r {
		return nd[p].rval
	} else {
		push(p)
		md := (l + r) >> 1
		if x <= md {
			return queryv(p+p, l, md, x)
		}
		return queryv(p+p+1, md+1, r, x)
	}
}

func push(p int) {
	setf(p+p, nd[p].t)
	setf(p+p+1, nd[p].t)
	nd[p].t = tag{1, 0, 0}
}

func setf(p int, v tag) {
	nd[p].t = plus(nd[p].t, v)
	if v.p1 == 1 {
		nd[p].rval = nd[p].rval + v.p2*nd[p].rh + v.q2
	} else {
		nd[p].rval = v.p2*nd[p].rh + v.q2
	}
}

func query(p, l, r, val int) int {
	if val > nd[p].rval {
		return r + 1
	}
	if l == r {
		return r
	} else {
		push(p)
		md := (l + r) >> 1
		if nd[p+p].rval >= val {
			return query(p+p, l, md, val)
		} else {
			return query(p+p+1, md+1, r, val)
		}
	}
}

func modify(p, l, r, tl, tr int, v tag) {
	if tl > tr {
		return
	}
	if tl == l && tr == r {
		setf(p, v)
		return
	} else {
		push(p)
		md := (l + r) >> 1
		if tr <= md {
			modify(p+p, l, md, tl, tr, v)
		} else if tl > md {
			modify(p+p+1, md+1, r, tl, tr, v)
		} else {
			modify(p+p, l, md, tl, md, v)
			modify(p+p+1, md+1, r, md+1, tr, v)
		}
		upd(p)
	}
}

type tag struct {
	p1, p2, q2 int
}

func plus(a, b tag) tag {
	return tag{a.p1 * b.p1, a.p2*b.p1 + b.p2, a.q2*b.p1 + b.q2}
}

type node struct {
	rval, rh int
	t        tag
}

func Unique(a []int) []int {
	occurred := map[int]bool{}
	result := []int{}
	for i := range a {
		if occurred[a[i]] != true {
			occurred[a[i]] = true
			result = append(result, a[i])
		}
	}
	n := len(result)
	for i := 0; i < n; i++ {
		a[i] = result[i]
	}
	return result
}

func LowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}
