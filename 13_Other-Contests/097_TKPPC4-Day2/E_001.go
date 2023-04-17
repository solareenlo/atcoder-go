package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const MX = 100005

type Tuple struct {
	x, y, z int
}

type Pair struct {
	x, y int
}

var p, cnt [MX]int
var l []Tuple
var qq []Pair
var am int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, q int
	fmt.Fscan(in, &n, &m, &q)
	l = make([]Tuple, m)
	for i := 0; i < m; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		l[i] = Tuple{c, a - 1, b - 1}
	}
	qq = make([]Pair, q)
	for i := 0; i < q; i++ {
		var a int
		fmt.Fscan(in, &a)
		qq[i] = Pair{a, i}
	}
	sort.Slice(l, func(i, j int) bool {
		if l[i].x == l[j].x {
			if l[i].y == l[j].y {
				return l[i].z < l[j].z
			}
			return l[i].y < l[j].y
		}
		return l[i].x < l[j].x
	})
	for i := 0; i < n; i++ {
		p[i] = -1
	}
	cnt[1] = n
	sort.Slice(qq, func(i, j int) bool {
		if qq[i].x == qq[j].x {
			return qq[i].y < qq[j].y
		}
		return qq[i].x < qq[j].x
	})

	var ans [MX]int
	am = 1
	var t int
	for i := 0; i < q; i++ {
		for ; am < qq[i].x && t < m; t++ {
			uni(l[t].y, l[t].z)
		}
		if am >= qq[i].x {
			if t-1 < 0 {
				ans[qq[i].y] = 0
			} else {
				ans[qq[i].y] = l[t-1].x
			}
		} else {
			ans[qq[i].y] = -1
		}
	}
	for i := 0; i < q; i++ {
		if ans[i] == -1 {
			fmt.Fprintln(out, "trumpet")
		} else {
			fmt.Fprintln(out, ans[i])
		}
	}
}

func uni(x, y int) {
	x = find(x)
	y = find(y)
	if x != y {
		if p[x] > p[y] {
			x, y = y, x
		}
		cnt[-p[x]]--
		cnt[-p[y]]--
		p[x] += p[y]
		cnt[-p[x]]++
		p[y] = x
		for cnt[am] == 0 {
			am++
		}
	}
}

func find(x int) int {
	if p[x] < 0 {
		return x
	}
	p[x] = find(p[x])
	return p[x]
}
