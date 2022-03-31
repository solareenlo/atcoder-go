package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	const N = 300005
	a := make([]int, N)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	dp := make([]int, N)
	dp[0] = 1
	c := make([]int, N)
	c[0] = 1

	const mod = 998244353
	type node struct{ l, val int }
	mx := make([]node, N)
	mn := make([]node, N)
	hx, hn := 0, 0
	cx := make([]int, N)
	cn := make([]int, N)
	for i := 1; i <= n; i++ {
		lx, ln := i, i
		for hx != 0 && a[i] >= mx[hx].val {
			lx = mx[hx].l
			hx--
		}
		hx++
		mx[hx] = node{lx, a[i]}
		tmp := 0
		if lx >= 2 {
			tmp = c[lx-2]
		}
		cx[hx] = (cx[hx-1] + (c[i-1]-tmp+mod)*a[i]) % mod
		for hn != 0 && a[i] <= mn[hn].val {
			ln = mn[hn].l
			hn--
		}
		hn++
		mn[hn] = node{ln, a[i]}
		tmp = 0
		if ln >= 2 {
			tmp = c[ln-2]
		}
		cn[hn] = (cn[hn-1] + (c[i-1]-tmp+mod)*a[i]) % mod
		dp[i] = (cx[hx] - cn[hn] + mod) % mod
		c[i] = (c[i-1] + dp[i]) % mod
	}
	fmt.Println(dp[n])
}
