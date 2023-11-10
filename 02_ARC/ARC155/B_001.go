package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const INF = int(1e18)
const Q = 200000
const N = (Q+1)*2 + 1

var X int

func rand_() int {
	tmp := X
	X *= 3
	return tmp >> 1
}

var zz, xx, rr, ll [N]int
var u_, l_, r_ int

var cnt int = 1

func node(x int) int {
	zz[cnt] = rand_()
	xx[cnt] = x
	cnt++
	return cnt - 1
}

func split(u, x int) {
	if u == 0 {
		u_ = 0
		l_ = 0
		r_ = 0
		return
	}
	if xx[u] < x {
		split(rr[u], x)
		rr[u] = l_
		l_ = u
	} else if xx[u] > x {
		split(ll[u], x)
		ll[u] = r_
		r_ = u
	} else {
		u_ = u
		l_ = ll[u]
		r_ = rr[u]
		ll[u] = 0
		rr[u] = 0
	}
}

func merge(u, v int) int {
	if u == 0 {
		return v
	}
	if v == 0 {
		return u
	}
	if zz[u] < zz[v] {
		rr[u] = merge(rr[u], v)
		return u
	} else {
		ll[v] = merge(u, ll[v])
		return v
	}
}

func tr_add(x int) {
	split(u_, x)
	u_ = merge(merge(l_, node(x)), r_)
}

func tr_floor(x int) int {
	u := u_
	x_ := -1
	for u != 0 {
		if xx[u] <= x {
			x_ = xx[u]
			u = rr[u]
		} else {
			u = ll[u]
		}
	}
	return x_
}

func tr_ceil(x int) int {
	u := u_
	x_ := INF
	for u != 0 {
		if xx[u] >= x {
			x_ = xx[u]
			u = ll[u]
		} else {
			u = rr[u]
		}
	}
	return x_
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var q, a, b int
	fmt.Fscan(in, &q, &a, &b)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	X = r.Int() | 1
	tr_add(a - b)
	tr_add(a + b)
	for q > 0 {
		q--
		var t, p, q, d int
		fmt.Fscan(in, &t, &a, &b)
		if t == 1 {
			tr_add(a - b)
			tr_add(a + b)
		} else {
			if tr_ceil(a) <= b {
				fmt.Fprintln(out, 0)
			} else {
				p = tr_floor(a)
				q = tr_ceil(a)
				d = INF
				if p != -1 {
					d = min(d, a-p)
				}
				if q != INF {
					d = min(d, q-b)
				}
				fmt.Fprintln(out, d)
			}
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
