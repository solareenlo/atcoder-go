package main

import (
	"bufio"
	"fmt"
	"os"
)

const M = 6000005
const mx = 300000
const mod = 1000000007

var (
	t  int
	z  int
	t1 = [M]int{}
	t2 = [M]int{}
	lc = [M]int{}
	rc = [M]int{}
	tg = [M]int{}
)

func pu(k int) {
	t1[k] = (t1[lc[k]] + t1[rc[k]]) % mod
	t2[k] = (t2[lc[k]] + t2[rc[k]]) % mod
}

func pt(k, v int) {
	t1[k] = t1[k] * v % mod
	t2[k] = t2[k] * v % mod * v % mod
	tg[k] = tg[k] * v % mod
}

func pd(k int) {
	if tg[k] == 1 {
		return
	}
	if lc[k] != 0 {
		pt(lc[k], tg[k])
	}
	if rc[k] != 0 {
		pt(rc[k], tg[k])
	}
	tg[k] = 1
}

func mg(x, y, l, r int) int {
	if x == 0 || y == 0 {
		return x ^ y
	}
	if l == r {
		t1[x] = (t1[x] + t1[y]) % mod
		t2[x] = t1[x] * t1[x] % mod
		return x
	}
	pd(x)
	pd(y)
	z = (z + t1[rc[x]]*t1[lc[y]]%mod) % mod
	mid := (l + r) >> 1
	lc[x] = mg(lc[x], lc[y], l, mid)
	rc[x] = mg(rc[x], rc[y], mid+1, r)
	pu(x)
	return x
}

func ins(k *int, l, r, v int) {
	if *k == 0 {
		t++
		*k = t
		tg[*k] = 1
	}
	if l == r {
		t1[*k] = 1
		t2[*k] = 1
		return
	}
	mid := (l + r) >> 1
	if v <= mid {
		ins(&lc[*k], l, mid, v)
	} else {
		ins(&rc[*k], mid+1, r, v)
	}
	pu(*k)
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	st := make([]int, n+1)
	ans := make([]int, n+1)
	tp := 0
	for i := 1; i <= n; i++ {
		var a int
		fmt.Fscan(in, &a)
		if a > 0 {
			tp++
			ans[tp] = 0
			st[tp] = 0
			ins(&st[tp], 1, mx, a)
		} else if a < 0 {
			a = -a
			k := st[tp]
			ans[tp] = (ans[tp]*a + (t1[k]*t1[k]%mod-t2[k]%mod+mod)%mod*(mod+1)/2%mod*(a*(a-1)/2%mod)%mod) % mod
			pt(k, a)
		} else {
			z = 0
			st[tp-1] = mg(st[tp-1], st[tp], 1, mx)
			tp--
			ans[tp] = (ans[tp] + ans[tp+1]) % mod
			ans[tp] = (ans[tp] + z) % mod
		}
	}
	fmt.Println(ans[tp])
}
