package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	const N = 1 << 17

	type frac struct {
		a, b int
	}

	type node struct {
		l, r int
		ans  frac
	}

	var cmp func(frac, frac) bool
	cmp = func(x, y frac) bool {
		return x.a*y.b < y.a*x.b
	}

	var max func(a, b frac) frac
	max = func(a, b frac) frac {
		if !cmp(a, b) {
			return a
		}
		return b
	}

	var n, q int
	fmt.Fscan(in, &n, &q)
	var a [N]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		a[i] += a[i-1]
	}

	tr := make([]node, N*2)
	for i := range tr {
		tr[i].l = -1
		tr[i].r = -1
		tr[i].ans.a = 0
		tr[i].ans.b = 1
	}

	for i := 0; i < q; i++ {
		var op int
		fmt.Fscan(in, &op)
		if op == 1 {
			var x int
			fmt.Fscan(in, &x)
			if tr[N+x].l == -1 {
				tr[N+x].l = x
				tr[N+x].r = x
			} else {
				tr[N+x].l = -1
				tr[N+x].r = -1
			}
			for i := (N + x) >> 1; i > 0; i >>= 1 {
				tr[i].ans = max(tr[i*2].ans, tr[i*2+1].ans)
				if tr[i*2].l == -1 {
					tr[i].l = tr[i*2+1].l
				} else {
					tr[i].l = tr[i*2].l
				}
				if tr[i*2+1].r == -1 {
					tr[i].r = tr[i*2].r
				} else {
					tr[i].r = tr[i*2+1].r
				}
				if tr[i*2].r != -1 && tr[i*2+1].l != -1 {
					tr[i].ans = max(tr[i].ans, frac{a[tr[i*2+1].l] - a[tr[i*2].r], tr[i*2+1].l - tr[i*2].r})
				}
			}
		} else {
			g := gcd(tr[1].ans.a, tr[1].ans.b)
			fmt.Fprintln(out, tr[1].ans.a/g, tr[1].ans.b/g)
		}
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
