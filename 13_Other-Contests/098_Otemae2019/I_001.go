package main

import (
	"bufio"
	"fmt"
	"os"
)

const MX = 500005
const MOD = 998244353

var n, ans int
var G [MX][]int
var ring []int
var rr [][]int
var ac []int = []int{0, 0, 0}
var used [MX]bool
var rt int = -1

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		G[a] = append(G[a], b)
		G[b] = append(G[b], a)
	}
	dfs1(0, 0)
	Len := len(ring)
	for _, t := range ring {
		rr = append(rr, dfs2(t))
	}
	for i := Len/2 + 1; i < Len; i++ {
		ac = merge(prop(ac, 1), rr[i], false)
	}
	for i := 0; i < Len; i++ {
		ac = merge(prop(ac, 1), rr[i], true)
		tmp := prop(rr[(i+Len/2+1)%Len], Len-Len/2-1)
		for j := 0; j < 3; j++ {
			ac[j] = (ac[j] + MOD - tmp[j]) % MOD
		}
	}
	if Len%2 == 0 {
		for i := 0; i < Len/2; i++ {
			merge(rr[i], prop(rr[i+Len/2], Len/2), true)
		}
	}
	fmt.Println(ans)
}

func dfs1(u, p int) bool {
	used[u] = true
	for _, nx := range G[u] {
		if nx != p && (used[nx] || dfs1(nx, u)) {
			if rt == -1 {
				rt = nx
			}
			ring = append(ring, u)
			return rt != u
		} else if nx == p {
			p = -1
		}
	}
	used[u] = false
	return used[u]
}

func dfs2(u int) []int {
	used[u] = true
	ret := []int{0, 0, 1}
	for _, nx := range G[u] {
		if !used[nx] {
			ret = merge(ret, prop(dfs2(nx), 1), true)
		}
	}
	return ret
}

func merge(a, b []int, c bool) []int {
	ret := []int{0, 0, 0}
	if c {
		ans = (ans + a[0]*b[2] + b[0]*a[2] + 2*a[1]*b[1]) % MOD
	}
	for i := 0; i < 3; i++ {
		ret[i] = (a[i] + b[i]) % MOD
	}
	return ret
}

func prop(t []int, k int) []int {
	return []int{(t[0] + t[1]*2*k + t[2]*k%MOD*k) % MOD, (t[1] + t[2]*k) % MOD, t[2]}
}
