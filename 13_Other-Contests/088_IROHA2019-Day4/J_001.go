package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 100005

var xp [N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var vis [N]bool
	var pr, f, g [N]int
	m := 0
	for i, li := 2, int(1e5); i <= li; i++ {
		if !vis[i] {
			m++
			pr[m] = i
			f[i] = 1
			g[i] = i & 1
			xp[i] = i
		}
		for j := 1; j <= m; j++ {
			u := i * pr[j]
			if u > li {
				break
			}
			xp[u] = xp[i]
			f[u] = f[i] + 1
			g[u] = g[i] + (pr[j] & 1)
			vis[u] = true
			if (i % pr[j]) == 0 {
				break
			}
		}
	}
	var n int
	fmt.Fscan(in, &n)
	var sa, sb [N]int
	for i := 1; i <= n; i++ {
		var o1, o2, w, h int
		fmt.Fscan(in, &o1, &o2, &w, &h)
		sa[i] = sa[i-1]
		sb[i] = sb[i-1]
		if o1 != 0 && o2 != 0 {
			tmp := 0
			if (w & h & 1) == 0 {
				tmp = 1
			}
			sb[i] ^= (g[w] ^ g[h]) + tmp
		}
		if o1 == 0 && o2 == 0 {
			if f[w] > f[h] {
				sa[i] += Calc(f[w]-f[h], w)
			} else {
				sa[i] += -Calc(f[h]-f[w], h)
			}
		}
		if o1 != 0 && o2 == 0 {
			sa[i] += Calc(f[w], w) * h
			tmp := 0
			if (h & 1) == 0 {
				tmp = 1
			}
			sb[i] ^= g[h] + tmp
		}
		if o1 == 0 && o2 != 0 {
			sa[i] -= Calc(f[h], h) * w
			tmp := 0
			if (w & 1) == 0 {
				tmp = 1
			}
			sb[i] ^= g[w] + tmp
		}
	}
	var q int
	fmt.Fscan(in, &q)
	for i := 1; i <= q; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		if (sa[v] ^ sa[u-1]) != 0 {
			if sa[v] > sa[u-1] {
				fmt.Fprintln(out, "Yes")
			} else {
				fmt.Fprintln(out, "No")
			}
		} else {
			if (sb[v] ^ sb[u-1]) != 0 {
				fmt.Fprintln(out, "Yes")
			} else {
				fmt.Fprintln(out, "No")
			}
		}
	}
}

func Calc(u, v int) int {
	sm := 0
	mul := 1
	for u > 0 {
		u--
		sm += mul
		mul *= xp[v]
		v /= xp[v]
	}
	return sm
}
