package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	sz  int
	dat = make([]int, 800800)
)

func upd(l, r, x int) {
	l += sz
	r += sz
	for l <= r {
		dat[l] = min(dat[l], x)
		dat[r] = min(dat[r], x)
		l = (l + 1) >> 1
		r = (r - 1) >> 1
	}
}

func qry(x int) int {
	r := 1 << 60
	x += sz
	for x != 0 {
		r = min(r, dat[x])
		x >>= 1
	}
	return r
}

func main() {
	in := bufio.NewReader(os.Stdin)

	for i := range dat {
		dat[i] = 1 << 60
	}

	var n int
	fmt.Fscan(in, &n)

	sz = 1
	for sz < n {
		sz <<= 1
	}
	w := make([]int, n+1)
	s := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &w[i])
		s[i] = s[i-1] + 1 - w[i]
	}

	var q int
	fmt.Fscan(in, &q)
	g := make([][]int, 200200)
	for i := 0; i < q; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		g[b] = append(g[b], a)
	}

	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		upd(i-1, i-1, dp[i-1]-s[i-1])
		for _, j := range g[i] {
			upd(j, i, qry(j-1))
		}
		dp[i] = min(dp[i-1]+w[i], qry(i)+s[i])
	}
	fmt.Println(dp[n])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
