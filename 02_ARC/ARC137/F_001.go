package main

import "fmt"

const N = 200000
const MD = 998244353
const L = 19 /* L = ceil(log2((N + 1) * 2)) */
const N_ = (1 << L)

var vv, ff, gg [N*2 + 1]int
var ff_, gg_ [N + 1]int
var vv_ [L + 1]int
var wwu, wwv [L + 1][]int

func main() {
	aa := make([]int, N+1)
	bb := make([]int, N+1)
	cc := make([]int, N+1)
	dd := make([]int, N+1)
	ee := make([]int, N+1)
	Init()
	var n, k int
	fmt.Scan(&n, &k)
	if k == n {
		fmt.Println(1)
		return
	}
	k++
	for i := 0; i <= n-k; i++ {
		aa[i] = ((ff[k+i*2] * gg[k] % MD) * gg[i*2] % MD) * ff_[i] % MD
	}
	for i := 0; i <= n-k; i++ {
		cc[i] = (ff[k+i*2] * gg[i*2] % MD) * ff_[i] % MD
	}
	for i := 0; i <= k; i++ {
		dd[i] = ((ff[k] * gg[i] % MD) * gg[k-i] % MD) * gg[i] % MD
	}
	for i := 0; i <= n-k; i++ {
		ee[i] = gg[i*2] * ff_[i] % MD
	}
	mult(dd, ee, n+1)
	for i := 0; i <= n; i++ {
		if i <= n-k {
			dd[i] = dd[i] * ff[i*2] % MD
		} else {
			dd[i] = 0
		}
	}
	inv(dd, bb, n-k+1)
	mult(aa, bb, n-k+1)
	ans := ff_[n]
	for i := 0; i <= n-k; i++ {
		ans = (ans - aa[i]*cc[n-k-i]%MD + MD) % MD
	}
	ans = ans * gg_[n] % MD
	fmt.Println(ans)
}

func inv(aa, bb []int, n int) {
	if n == 1 {
		bb[0] = 1
		return
	}
	m := (n + 1) / 2
	aa_ := make([]int, N_)
	bb_ := make([]int, N_)
	inv(aa, bb, m)
	l_ := 0
	for 1<<l_ < n*2 {
		l_++
	}
	n_ := 1 << l_
	for i := 0; i < n_; i++ {
		aa_[i] = 0
		bb_[i] = 0
		if i < n {
			aa_[i] = aa[i]
		}
		if i < m {
			bb_[i] = bb[i]
		}
	}
	ntt(aa_, l_, 0)
	ntt(bb_, l_, 0)
	for i := 0; i < n_; i++ {
		aa_[i] = bb_[i] * (MD + 2 - aa_[i]*bb_[i]%MD) % MD
	}
	ntt(aa_, l_, 1)
	for i := 0; i < n; i++ {
		bb[i] = aa_[i] * vv_[l_] % MD
	}
}

func mult(aa, bb []int, n int) {
	aa_ := make([]int, N_)
	bb_ := make([]int, N_)
	l_ := 0
	for 1<<l_ < n*2-1 {
		l_++
	}
	n_ := 1 << l_
	for i := 0; i < n_; i++ {
		if i < n {
			aa_[i] = aa[i]
			bb_[i] = bb[i]
		} else {
			aa_[i] = 0
			bb_[i] = 0
		}
	}
	ntt(aa_, l_, 0)
	ntt(bb_, l_, 0)
	for i := 0; i < n_; i++ {
		aa_[i] = aa_[i] * bb_[i] % MD
	}
	ntt(aa_, l_, 1)
	for i := 0; i < n; i++ {
		aa[i] = aa_[i] * vv_[l_] % MD
	}
}

func ntt(aa []int, l, inverse int) {
	n := 1 << l
	for i, j := 0, 1; j < n; j++ {
		b := n >> 1
		i ^= b
		for i < b {
			b >>= 1
			i ^= b
		}
		if i < j {
			aa[i], aa[j] = aa[j], aa[i]
		}
	}
	ntt_(aa, l, inverse)
}

func ntt_(aa []int, l, inverse int) {
	if l != 0 {
		n := 1 << l
		m := n >> 1
		var ww []int
		if inverse != 0 {
			ww = wwv[l]
		} else {
			ww = wwu[l]
		}
		ntt_(aa, l-1, inverse)
		tmp := aa[m:]
		ntt_(tmp, l-1, inverse)
		for i := 0; i+m < n; i++ {
			j := i + m
			a := aa[i]
			b := aa[j] * ww[i] % MD
			aa[i] = (a + b) % MD
			aa[j] = (a - b + MD) % MD
		}
	}
}

func Init() {
	ff[0] = 1
	gg[0] = 1
	for i := 1; i <= N*2; i++ {
		if i == 1 {
			vv[i] = 1
		} else {
			vv[i] = vv[i-MD%i] * (MD/i + 1) % MD
		}
		ff[i] = ff[i-1] * i % MD
		gg[i] = gg[i-1] * vv[i] % MD
	}
	ff_[0] = 1
	gg_[0] = 1
	for i := 1; i <= N; i++ {
		ff_[i] = ff_[i-1] * (i*2 - 1) % MD
		gg_[i] = gg_[i-1] * vv[i*2-1] % MD
	}
	u := power(3, (MD-1)>>L)
	v := power(u, MD-2)
	for l := L; l > 0; l-- {
		n := 1 << l
		m := n >> 1
		vv_[l] = power(1<<l, MD-2)
		wwu[l] = make([]int, m)
		wwv[l] = make([]int, m)
		wwu[l][0] = 1
		wwv[l][0] = 1
		for i := 1; i < m; i++ {
			wwu[l][i] = wwu[l][i-1] * u % MD
			wwv[l][i] = wwv[l][i-1] * v % MD
		}
		u = u * u % MD
		v = v * v % MD
	}
	vv_[0] = 1
}

func power(a, k int) int {
	p := 1
	for k != 0 {
		if (k & 1) != 0 {
			p = p * a % MD
		}
		a = a * a % MD
		k >>= 1
	}
	return p
}
