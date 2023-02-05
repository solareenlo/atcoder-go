package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://atcoder.jp/contests/abc278/editorial/5210

const N = 200000
const L = 19 /* L = ceil(log2(N * 2 + 1)) */
const N_ = (1 << L)
const MD = 998244353

var pp2, pp2_, qq2, vq2, vv, ff, gg [N + 1]int
var vv_ []int
var wwu, wwv [][]int

func main() {
	in := bufio.NewReader(os.Stdin)

	ss := make([]int, N+1)
	aa := make([]int, N+1)
	bb := make([]int, N+1)
	vv_ = make([]int, L+1)
	wwu = make([][]int, L+1)
	wwv = make([][]int, L+1)
	Init()
	var n, m int
	fmt.Fscan(in, &n, &m)
	n--
	stirling(ss, n)
	p := power(2, m-1)
	q := 1
	for i := 0; i <= n && i < m; i++ {
		aa[i] = q * vq2[i] % MD * pp2_[i] % MD * pp2[i] % MD
		q = q * (p - 1) % MD
		p = p * vv[2] % MD
	}
	for i := 0; i <= n; i++ {
		bb[i] = vq2[i]
	}
	multiply(aa, bb, aa, n+1, n+1, n+1)
	ans := 0
	for i := 0; i <= n; i++ {
		tmp := 1
		if i%2 != 0 {
			tmp = MD - 1
		}
		ans += aa[n-i] * qq2[n-i] % MD * ss[n-i] % MD * tmp
		ans %= MD
	}
	fmt.Println(ans)
}

func power(a, k int) int {
	p := 1
	for k != 0 {
		if (k & 1) > 0 {
			p = p * a % MD
		}
		a = a * a % MD
		k >>= 1
	}
	return p
}

func multiply(aa, bb, cc []int, n, m, n1 int) {
	aa_ := make([]int, N_)
	bb_ := make([]int, N_)
	l := 0
	for 1<<l < n+m-1 {
		l++
	}
	n_ := 1 << l
	for i := 0; i < n_; i++ {
		if i < n {
			aa_[i] = aa[i]
		} else {
			aa_[i] = 0
		}
		if i < m {
			bb_[i] = bb[i]
		} else {
			bb_[i] = 0
		}
	}
	ntt(aa_, l, 0)
	ntt(bb_, l, 0)
	for i := 0; i < n_; i++ {
		aa_[i] = aa_[i] * bb_[i] % MD
	}
	ntt(aa_, l, 1)
	for i := 0; i < n1; i++ {
		cc[i] = aa_[i] * vv_[l] % MD
	}
}

func ntt_(aa []int, l, inverse int) {
	if l != 0 {
		n := 1 << l
		m := n >> 1
		var ww []int
		if inverse != 0 {
			ww = make([]int, len(wwv[l]))
			copy(ww, wwv[l])
		} else {
			ww = make([]int, len(wwu[l]))
			copy(ww, wwu[l])
		}
		ntt_(aa, l-1, inverse)
		tmp := aa[m:]
		ntt_(tmp, l-1, inverse)
		for i := 0; i+m < n; i++ {
			j := i + m
			a := aa[i]
			b := aa[j] * ww[i] % MD
			aa[i] = a + b
			if aa[i] >= MD {
				aa[i] -= MD
			}
			aa[j] = a - b
			if aa[j] < 0 {
				aa[j] += MD
			}
			j = i + m
		}
	}
}

func ntt(aa []int, l, inverse int) {
	n := 1 << l
	for i, j := 0, 1; j < n; j++ {
		m := n >> 1
		i ^= m
		for i < m {
			m >>= 1
			i ^= m
		}
		if i < j {
			aa[i], aa[j] = aa[j], aa[i]
		}
	}
	ntt_(aa, l, inverse)
}

func shift(aa, bb []int, n, x int) {
	aa_ := make([]int, N+1)
	bb_ := make([]int, N+1)
	cc_ := make([]int, N*2+1)
	for i := 0; i <= n; i++ {
		aa_[i] = aa[i] * ff[i] % MD
	}
	p := 1
	for i := 0; i <= n; i++ {
		bb_[n-i] = p * gg[i] % MD
		p = p * x % MD
	}
	multiply(aa_, bb_, cc_, n+1, n+1, n*2+1)
	for i := 0; i <= n; i++ {
		bb[i] = cc_[n+i] * gg[i] % MD
	}
}

func stirling(aa []int, n int) {
	bb := make([]int, N+1)
	if n == 0 {
		aa[0] = 1
		return
	}
	stirling(aa, n/2)
	shift(aa, bb, n/2+1, n/2)
	multiply(aa, bb, aa, n/2+1, n/2+1, n/2*2+1)
	if n%2 != 0 {
		for i := n; i >= 0; i-- {
			tmp := 0
			if i != 0 {
				tmp = aa[i-1]
			}
			aa[i] = (aa[i]*(n-1) + tmp) % MD
		}
	}
}

func inv(a int) int {
	if a == 1 {
		return 1
	}
	return inv(a-MD%a) * (MD/a + 1) % MD
}

func Init() {
	pp2[0], pp2_[0], qq2[0], ff[0], gg[0] = 1, 1, 1, 1, 1
	for i := 1; i <= N; i++ {
		pp2[i] = pp2[i-1] * 2 % MD
		pp2_[i] = pp2_[i-1] * pp2[i] % MD
		qq2[i] = qq2[i-1] * (pp2[i] - 1) % MD
		if i == 1 {
			vv[i] = 1
		} else {
			vv[i] = vv[i-MD%i] * (MD/i + 1) % MD
		}
		ff[i] = ff[i-1] * i % MD
		gg[i] = gg[i-1] * vv[i] % MD
	}
	vq2[N] = inv(qq2[N])
	for i := N - 1; i >= 0; i-- {
		vq2[i] = vq2[i+1] * (pp2[i+1] - 1) % MD
	}
	u := power(3, (MD-1)>>L)
	v := power(u, MD-2)
	for l := L; l > 0; l-- {
		n := 1 << l
		m := n >> 1
		vv_[l] = power(1<<l, MD-2)
		wwu[l] = make([]int, m)
		wwv[l] = make([]int, m)
		wwu[l][0], wwv[l][0] = 1, 1
		for i := 1; i < m; i++ {
			wwu[l][i] = wwu[l][i-1] * u % MD
			wwv[l][i] = wwv[l][i-1] * v % MD
		}
		u = u * u % MD
		v = v * v % MD
	}
	vv_[0] = 1
}
