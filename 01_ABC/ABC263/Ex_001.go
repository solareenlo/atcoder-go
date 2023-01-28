package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type node struct {
	k float64
	p int
}

const maxn = 5e4 + 50
const eps = 1e-6
const Pi = math.Pi

var n, m, k int
var p [maxn]float64
var d [maxn]float64
var aa []node
var t [maxn << 1]int
var vis [maxn]int

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &k)
	for i := 1; i <= n; i++ {
		var A, B, C float64
		fmt.Fscan(in, &A, &B, &C)
		d[i] = math.Abs(C / math.Sqrt(A*A+B*B))
		if C != 0 {
			p[i] = math.Atan2(-B*C, -A*C)
		} else {
			p[i] = math.Atan2(B, A)
		}
	}

	aa = make([]node, maxn<<1)
	var l, r float64
	l, r = 0, 3e6
	for r-l > eps {
		mid := (l + r) / 2
		if calc(mid) < k {
			l = mid
		} else {
			r = mid
		}
	}
	fmt.Println(l)
}

func calc(r float64) int {
	m = 0
	for i := 1; i <= n; i++ {
		if d[i]-r > eps {
			continue
		}
		x, y := p[i], math.Acos(d[i]/r)
		m++
		aa[m] = node{k: x + y, p: i}
		if aa[m].k > Pi {
			aa[m].k -= 2 * Pi
		}
		m++
		aa[m] = node{k: x - y, p: i}
		if aa[m].k < -Pi {
			aa[m].k += 2 * Pi
		}
	}
	tmp := aa[1 : m+1]
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].k < tmp[j].k
	})
	for i := 1; i <= m; i++ {
		t[i] = 0
	}
	for i := 1; i <= n; i++ {
		vis[i] = 0
	}
	ret := 0
	for i := 1; i <= m; i++ {
		if vis[aa[i].p] == 0 {
			vis[aa[i].p] = i
		} else {
			ret += query(vis[aa[i].p])
			add(vis[aa[i].p], 1)
			add(i, -1)
		}
	}
	return ret
}

func add(x, k int) {
	for x <= m {
		t[x] += k
		x += lowbit(x)
	}
}

func query(x int) int {
	ret := 0
	for x > 0 {
		ret += t[x]
		x -= lowbit(x)
	}
	return ret
}

func lowbit(x int) int {
	return x & -x
}
