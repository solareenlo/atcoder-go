package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const o = 18
const leng = 1 << 18

var n int
var a []int
var g, w, r [leng]int
var iv, fac, ifac [leng]int
var up, L int
var v [][]int
var f []int
var X, Y []int

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)
	Init()
	prep(n)
	a = make([]int, leng)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	tmp := a[1 : n+1]
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] < tmp[j]
	})
	for i := 1; i <= n; i++ {
		a[i] = mod - a[i] - 2
	}

	v = make([][]int, 19)
	f = make([]int, leng)
	for i := 2; i <= lg(n)+1; i++ {
		leng := 1 << i
		for j := 0; j < leng; j++ {
			f[j] = iv[j]
		}
		pre(leng - 1)
		ntt(f, leng, false)
		v[i] = make([]int, leng)
		for j := 0; j < leng; j++ {
			v[i][j] = f[j]
		}
		for i := 0; i < leng; i++ {
			f[i] = 0
		}
	}
	X = make([]int, leng)
	Y = make([]int, leng)
	solve(1, n, f)

	g[0] = (powMod(2, n) - 1) * ifac[n] % mod
	for i := 1; i <= n; i++ {
		x := ifac[i] * ifac[n-i] % mod
		if (i & 1) == 0 {
			x = mod - x
		}
		g[i] = (x - g[i-1] + mod) % mod
	}

	ans := 0
	for i := 0; i <= n; i++ {
		x := ifac[i] * ifac[n-i] % mod
		if (i & 1) != 0 {
			x = mod - x
		}
		ans += ((x + g[i]) % mod) * f[i] % mod
		ans %= mod
	}
	fmt.Println(ans * n % mod)
}

func Init() {
	w1 := powMod(3, (mod-1)>>o)
	w[leng>>1] = 1
	for i := (leng >> 1) + 1; i != leng; i++ {
		w[i] = w[i-1] * w1 % mod
	}
	for i := (leng >> 1) - 1; i > 0; i-- {
		w[i] = w[i<<1]
	}
	for i := 1; i != leng; i++ {
		r[i] = (r[i>>1] >> 1) | ((i & 1) << (o - 1))
	}
}

func prep(n int) {
	fac[0] = 1
	for i := 1; i <= n; i++ {
		fac[i] = fac[i-1] * i % mod
	}
	ifac[n] = powMod(fac[n], mod-2)
	for i := n - 1; i != -1; i-- {
		ifac[i] = ifac[i+1] * (i + 1) % mod
	}
	for i := 1; i <= n; i++ {
		iv[i] = ifac[i] * fac[i-1] % mod
	}
}

func pre(n int) int {
	L = lg(n) + 1
	up = 1 << L
	return up
}

var t [leng]int

func ntt(a []int, n int, op bool) {
	for i := 0; i < n; i += 2 {
		x := a[r[i]>>(o-L)]
		y := a[r[i+1]>>(o-L)]
		t[i] = (x + y) % mod
		t[i+1] = (x + mod - y) % mod
	}
	for l := 2; l < n; l <<= 1 {
		k := w[l:]
		for idxF := 0; idxF < n; idxF += l {
			for idxK := 0; idxK < l; idxK++ {
				x := t[idxF]
				y := t[idxF+l] * k[idxK] % mod
				t[idxF+l] = (x + mod - y) % mod
				t[idxF] = (t[idxF] + y) % mod
				idxF++
			}
		}
	}
	if op {
		x := mod - ((mod - 1) >> L)
		for i := 0; i < n; i++ {
			a[i] = t[i] * x % mod
		}
		tmp := a[1:n]
		tmp = reverseOrderInt(tmp)
		for i := 0; i < n-1; i++ {
			a[1+i] = tmp[i]
		}
	} else {
		for i := 0; i < n; i++ {
			a[i] = t[i]
		}
	}
}

func reverseOrderInt(a []int) []int {
	n := len(a)
	res := make([]int, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func solve(l, r int, f []int) {
	if l == r {
		f[0] = a[l] + l
		f[1] = a[l] + l - 1
		return
	}

	mid := (l + r) >> 1
	l1 := mid - l + 1
	l2 := r - mid
	solve(l, mid, f)
	solve(mid+1, r, f[l1+1:])

	for i := 0; i <= l1; i++ {
		v := ifac[i] * ifac[l1-i] % mod
		if (l1-i)&1 != 0 {
			v = mod - v
		}
		X[i] = f[i] * v % mod
	}
	for i := 0; i <= l2; i++ {
		v := ifac[i] * ifac[l2-i] % mod
		if (l2-i)&1 != 0 {
			v = mod - v
		}
		Y[i] = f[l1+1+i] * v % mod
	}

	up := pre(r - l + 1)
	_l := L
	ntt(X, up, false)
	ntt(Y, up, false)
	for i := 0; i < up; i++ {
		X[i] = X[i] * v[_l][i] % mod
		Y[i] = Y[i] * v[_l][i] % mod
	}
	ntt(X, up, true)
	ntt(Y, up, true)

	for i := 0; i < (l1 + 1); i++ {
		X[i] = f[i]
	}
	for i := 0; i < (l2 + 1); i++ {
		Y[i] = f[l1+1+i]
	}
	for i := l1 + 1; i <= r-l+1; i++ {
		X[i] = (X[i] * fac[i] % mod) * ifac[i-l1-1] % mod
	}
	for i := l2 + 1; i <= r-l+1; i++ {
		Y[i] = (Y[i] * fac[i] % mod) * ifac[i-l2-1] % mod
	}
	for i := 0; i <= r-l+1; i++ {
		f[i] = X[i] * Y[i] % mod
	}
	f[r-l+2] = 0

	for i := 0; i < up; i++ {
		X[i] = 0
		Y[i] = 0
	}
}

func lg(__n int) int {
	var __k int
	for __k = 0; __n != 0; __n >>= 1 {
		__k++
	}
	return __k - 1
}

const mod = 998244353

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		n /= 2
	}
	return res
}
