package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const maxn = 200010
const mod = 998244353

var L, R, X, V, F, S [maxn]int
var pos, lim, col, pw, ipw [maxn]int
var tot int
var tr [maxn << 2]int
var Q, P [][]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, q int
	fmt.Fscan(in, &n, &m, &q)

	for i := 1; i <= q; i++ {
		fmt.Fscan(in, &L[i], &R[i], &X[i])
		V[i] = X[i]
	}
	V[q+1] = m
	tmp := V[1 : q+2]
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] < tmp[j]
	})
	tmp2 := V[1 : q+2]
	tot = len(unique(tmp2))
	build(1, 1, n)
	Q = make([][]int, maxn)
	for i := 1; i <= q; i++ {
		tmp := V[1 : tot+1]
		X[i] = lowerBound(tmp, X[i]) + 1
		Q[X[i]] = append(Q[X[i]], i)
		upd(1, 1, n, L[i], R[i], X[i])
	}
	bl(1, 1, n, tot)
	P = make([][]int, maxn)
	for i := 1; i <= n; i++ {
		P[col[i]] = append(P[col[i]], i)
	}
	ans := 1
	for i := 1; i <= tot; i++ {
		tmp := solve(i)
		ans = ans * tmp % mod
	}
	fmt.Println(ans)
}

func unique(a []int) []int {
	occurred := map[int]bool{}
	result := []int{}
	for i := range a {
		if occurred[a[i]] != true {
			occurred[a[i]] = true
			result = append(result, a[i])
		}
	}
	sort.Ints(result)
	n := len(result)
	for i := 0; i < n; i++ {
		a[i] = result[i]
	}
	return result
}

func build(h, l, r int) {
	tr[h] = tot
	if l == r {
		return
	}
	mid := (l + r) >> 1
	build(h<<1, l, mid)
	build(h<<1|1, mid+1, r)
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}

func upd(h, l, r, x, y, z int) {
	if l > y || r < x {
		return
	}
	if l >= x && r <= y {
		tr[h] = min(tr[h], z)
		return
	}
	mid := (l + r) >> 1
	upd(h<<1, l, mid, x, y, z)
	upd(h<<1|1, mid+1, r, x, y, z)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func bl(h, l, r, x int) {
	x = min(x, tr[h])
	if l == r {
		col[l] = x
		return
	}
	mid := (l + r) >> 1
	bl(h<<1, l, mid, x)
	bl((h<<1)|1, mid+1, r, x)
}

func solve(c int) int {
	if len(P[c]) == 0 && len(Q[c]) == 0 {
		return 1
	}
	num := 0
	for _, x := range P[c] {
		num++
		pos[num] = x
	}
	for i := 0; i <= num+1; i++ {
		lim[i] = -1
		F[i], S[i] = 0, 0
	}
	for _, x := range Q[c] {
		tmp := pos[1 : num+1]
		l := lowerBound(tmp, L[x]) + 1
		r := upperBound(tmp, R[x]) + 1
		lim[r] = max(lim[r], l-1)
		if l == r {
			return 0
		}
	}
	pw[0], ipw[0] = 1, 1
	ix := ksm(V[c], mod-2)
	now := lim[0]
	for i := 1; i <= num; i++ {
		pw[i] = pw[i-1] * V[c] % mod
	}
	for i := 1; i <= num; i++ {
		ipw[i] = ipw[i-1] * ix % mod
	}
	for i := 1; i <= num; i++ {
		now = max(now, lim[i])
		if now == -1 {
			F[i] = pw[i-1]
		}
		F[i] = (F[i] + (S[i-1]-S[max(0, now)]+mod)*pw[i-1]) % mod
		S[i] = (S[i-1] + F[i]*ipw[i]) % mod
	}
	now = max(now, lim[num+1])
	res := (S[num] - S[max(0, now)] + mod) * pw[num] % mod
	if now == -1 {
		res = (res + pw[num]) % mod
	}
	return res
}
func ksm(x, y int) int {
	res := 1
	for y > 0 {
		if (y & 1) != 0 {
			res = res * x % mod
		}
		x = x * x % mod
		y >>= 1
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
