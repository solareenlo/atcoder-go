package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const MOD = 1000000007

type node struct {
	l, m, r int
	b       bool
}

var dat [800000]node
var L [800000]int
var a [500000]int
var s []string
var N int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	var t string
	fmt.Fscan(in, &t)
	t += strings.Repeat(" ", 500000-len(t))
	s = strings.Split(t, "")

	Init(n)

	var q int
	fmt.Fscan(in, &q)
	for i := 0; i < q; i++ {
		var t, x, y int
		fmt.Fscan(in, &t, &x, &y)
		if t == 1 {
			x--
			a[x] = y
			dat[x+N-1].m = y
			update(x)
		}
		if t == 2 {
			x--
			if s[x] == "+" {
				s[x] = "*"
			} else {
				s[x] = "+"
			}
			update(x)
		}
		if t == 3 {
			x--
			res := query(x, y, 0, 0, N)
			fmt.Println((res.l + res.m + res.r) % MOD)
		}
	}
}

func Init(n int) {
	N = 1
	for N < n {
		N <<= 1
	}
	for i := n - 1; i < N; i++ {
		s[i] = "+"
	}
	for i := 2*N - 2; i >= 0; i-- {
		if i >= N-1 {
			dat[i] = node{0, a[i-N+1], 0, true}
			L[i] = i - N + 1
		} else {
			L[i] = L[i*2+1]
			dat[i] = merge(dat[i*2+1], dat[i*2+2], s[L[i*2+2]-1])
		}
	}
}

func update(k int) {
	k += N - 1
	for k != 0 {
		k = (k - 1) / 2
		dat[k] = merge(dat[k*2+1], dat[k*2+2], s[L[k*2+2]-1])
	}
}

func merge(l, r node, c string) node {
	if l.l == -1 {
		return r
	}
	if r.l == -1 {
		return l
	}
	if c == "+" {
		if l.b && r.b {
			return node{l.m, 0, r.m, false}
		}
		if l.b {
			return node{l.m, (r.l + r.m) % MOD, r.r, false}
		}
		if r.b {
			return node{l.l, (l.m + l.r) % MOD, r.m, false}
		}
		return node{l.l, (l.m + l.r + r.l + r.m) % MOD, r.r, false}
	} else {
		if l.b && r.b {
			return node{0, l.m * r.m % MOD, 0, true}
		}
		if l.b {
			return node{l.m * r.l % MOD, r.m, r.r, false}
		}
		if r.b {
			return node{l.l, l.m, l.r * r.m % MOD, false}
		}
		return node{l.l, (l.m + l.r*r.l + r.m) % MOD, r.r, false}
	}
}

func query(a, b, k, l, r int) node {
	if b <= l || r <= a {
		return node{-1, 0, 0, false}
	}
	if a <= l && r <= b {
		return dat[k]
	}
	lb := query(a, b, k*2+1, l, (l+r)/2)
	rb := query(a, b, k*2+2, (l+r)/2, r)
	return merge(lb, rb, s[(l+r)/2-1])
}
