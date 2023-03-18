package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var t int
	fmt.Fscan(in, &t)
	var a string
	fmt.Fscan(in, &a)
	l := make([]int, t)
	r := make([]int, t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &l[i], &r[i])
		l[i]--
	}

	INIT := make([]Info, t)
	qrys := make([]int, t)
	for i := range qrys {
		qrys[i] = i
	}

	val := make([]Info, len(a))
	var solve func(int, int, []int)
	solve = func(sl, sr int, qrys []int) {
		if sr-sl == 1 {
			for _, i := range qrys {
				INIT[i] = NewInfo2(int(a[sl] - '0'))
			}
			return
		}
		for i := sl; i < sr; i++ {
			val[i] = NewInfo2(int(a[i] - '0'))
		}
		m := (sl + sr) / 2
		for i := m - 2; i >= sl; i-- {
			val[i] = merge(val[i], val[i+1])
		}
		for i := m + 1; i < sr; i++ {
			val[i] = merge(val[i-1], val[i])
		}
		ql := make([]int, 0)
		qr := make([]int, 0)
		for _, i := range qrys {
			if r[i] <= m {
				ql = append(ql, i)
			} else if l[i] >= m {
				qr = append(qr, i)
			} else {
				INIT[i] = merge(val[l[i]], val[r[i]-1])
			}
		}
		solve(sl, m, ql)
		solve(m, sr, qr)
	}
	solve(0, len(a), qrys)

	var q int
	fmt.Fscan(in, &q)

	tree := make([]*Tree, q+1)
	tree[0] = build(0, t, INIT)

	for i := 0; i < q; i++ {
		var o int
		fmt.Fscan(in, &o)
		if o == 0 {
			var x int
			fmt.Fscan(in, &x)
			x--
			v := find(tree[i], t, x)
			fmt.Println(v.info.ans)
			tree[i+1] = tree[i]
		} else if o == 1 {
			var k int
			fmt.Fscan(in, &k)
			tree[i+1] = tree[k]
		} else {
			var a, b int
			fmt.Fscan(in, &a, &b)
			a--
			b--
			u := find(tree[i], t, a)
			v := find(tree[i], t, b)
			res := merge(u.info, v.info)
			tree[i+1] = tree[i]
			if u != v {
				if u.siz < v.siz {
					u, v = v, u
				}
				tree[i+1] = modify(tree[i+1], 0, t, v.f, Info{}, u.f, 0)
				tree[i+1] = modify(tree[i+1], 0, t, u.f, res, u.f, u.siz+v.siz)
			}
		}
	}
}

type Info struct {
	lenl int
	lenr int
	len  int
	ans  int
}

func NewInfo1() *Info {
	return &Info{0, 0, 0, 0}
}

func NewInfo2(x int) Info {
	tmp := 0
	if x == 0 {
		tmp = 1
	}
	return Info{tmp, tmp, 1, tmp}
}

func merge(a, b Info) Info {
	c := Info{}
	if a.len == a.lenl {
		c.lenl = a.len + b.lenl
	} else {
		c.lenl = a.lenl
	}
	if b.len == b.lenr {
		c.lenr = a.lenr + b.len
	} else {
		c.lenr = b.lenr
	}
	c.len = a.len + b.len
	c.ans = MAX(a.ans, b.ans, a.lenr+b.lenl)
	return c
}

func MAX(nums ...int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > res {
			res = nums[i]
		}
	}
	return res
}

type Tree struct {
	info Info
	l, r *Tree
	f    int
	siz  int
}

func build(l, r int, init []Info) *Tree {
	t := &Tree{}
	if r-l == 1 {
		t.info = init[l]
		t.f = l
		t.siz = 1
	} else {
		m := (l + r) / 2
		t.l = build(l, m, init)
		t.r = build(m, r, init)
	}
	return t
}

func GET(t *Tree, l, r, x int) *Tree {
	if r-l == 1 {
		return t
	}
	m := (l + r) / 2
	if x < m {
		return GET(t.l, l, m, x)
	}
	return GET(t.r, m, r, x)
}

func find(t *Tree, n, x int) *Tree {
	for {
		v := GET(t, 0, n, x)
		if v.f == x {
			return v
		}
		x = v.f
	}
}

func modify(t *Tree, l, r, p int, ni Info, nf, ns int) *Tree {
	x := &Tree{}
	x.l = t.l
	x.r = t.r
	if r-l == 1 {
		x.info = ni
		x.f = nf
		x.siz = ns
	} else {
		m := (l + r) / 2
		if p < m {
			x.l = modify(t.l, l, m, p, ni, nf, ns)
		} else {
			x.r = modify(t.r, m, r, p, ni, nf, ns)
		}
	}
	return x
}
