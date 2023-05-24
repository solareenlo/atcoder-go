package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type pair struct {
	x, y int
}

var n, k, p int
var dp [1 << 17]pair
var ar [1 << 17]int
var A []int

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &k)
	A = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &A[i])
	}
	sort.Ints(A)
	ok := -(1 << 40)
	ng := 1 << 40
	x := 0
	for ng-ok > 1 {
		x = (ok + ng) >> 1
		if f(x).y <= k {
			ok = x
		} else {
			ng = x
		}
	}
	fmt.Println(f(ok).x + ok*k)
}

func f(x int) pair {
	p = x
	for i := 0; i < n+1; i++ {
		dp[i] = pair{1 << 60, 0}
		ar[i] = 0
	}
	dp[0] = pair{0, 0}
	check(n, 0)
	solve(0, n)
	return dp[n]
}

func (l pair) lessThan(r pair) bool {
	if l.x == r.x {
		return l.y < r.y
	}
	return l.x < r.x
}

func check(i, j int) {
	t := pair{dp[j].x + A[i-1]*(i-j) - p, dp[j].y + 1}
	if t.lessThan(dp[i]) {
		dp[i] = t
		ar[i] = j
	}
}

func solve(u, d int) {
	if d-u == 1 {
		return
	}
	m := (u + d) >> 1
	for j := ar[u]; j <= ar[d]; j++ {
		check(m, j)
	}
	solve(u, m)
	for j := u + 1; j <= m; j++ {
		check(d, j)
	}
	solve(m, d)
}
