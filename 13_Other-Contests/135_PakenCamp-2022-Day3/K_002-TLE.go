package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M, K int
	fmt.Fscan(in, &N, &M, &K)
	A := make([]int, N)
	B := make([]int, N+1)
	C := make([]int, M)
	D := make([]int, M+1)
	for i := range A {
		fmt.Fscan(in, &A[i])
	}
	for i := 1; i <= N; i++ {
		fmt.Fscan(in, &B[i])
	}
	for i := range C {
		fmt.Fscan(in, &C[i])
	}
	for i := 1; i <= M; i++ {
		fmt.Fscan(in, &D[i])
	}
	if K < M {
		t := make([]int, len(A))
		copy(t, A)
		u := make([]int, N+1)
		v := make([]int, len(C))
		copy(v, C)
		w := make([]int, M+1)
		u[0] = 1
		w[0] = 1
		for i := 1; i <= N; i++ {
			u[i] = (MOD - B[i]) % MOD
		}
		for i := 1; i <= M; i++ {
			w[i] = (MOD - D[i]) % MOD
		}
		t = con.Convolve(t, u)
		v = con.Convolve(v, w)
		resize(&t, N)
		resize(&v, M)
		t = con.Convolve(t, invFPS(u, K+1))
		v = con.Convolve(v, invFPS(w, K+1))
		res := 0
		for i := 0; i < K+1; i++ {
			res = (res + t[i]*v[i]%MOD) % MOD
		}
		fmt.Println(res)
		return
	}

	Q := make([]int, M+1)
	Q[0] = 1
	for i := 1; i <= M; i++ {
		Q[i] = (MOD - D[i]) % MOD
	}
	RQ := make([]int, M+1)
	for i := 0; i < M+1; i++ {
		RQ[i] = Q[M-i]
	}
	x := myPow([]int{0, 1}, K-M+1, RQ)
	P := make([]int, len(C))
	copy(P, C)
	P = con.Convolve(P, Q)
	resize(&P, M)
	tmp := invFPS(Q, 2*M-1)
	P = con.Convolve(P, tmp)

	y := make([]int, M)
	G := make([]int, 2*M-1)
	resize(&x, M)
	for i := 0; i < M; i++ {
		y[i] = x[M-1-i]
	}
	for i := 0; i <= 2*M-2; i++ {
		G[i] = P[i]
	}
	V := con.Convolve(y, G)
	E := make([]int, M)
	for t := 0; t <= M-1; t++ {
		E[t] = V[M-1+t]
	}
	reverseOrderInt(E)
	D[0] = -1
	F := make([]int, M+1)
	for i := 1; i <= M; i++ {
		F[i] = divMod((MOD-D[M-i])%MOD, D[M])
	}

	p := make([]int, len(A))
	copy(p, A)
	q := make([]int, N+1)
	r := make([]int, len(E))
	copy(r, E)
	s := make([]int, M+1)
	q[0] = 1
	s[0] = 1
	for i := 1; i <= N; i++ {
		q[i] = (MOD - B[i]) % MOD
	}
	for i := 1; i <= M; i++ {
		s[i] = (MOD - F[i]) % MOD
	}
	p = con.Convolve(p, q)
	r = con.Convolve(r, s)
	resize(&p, N)
	resize(&r, M)
	fmt.Println(BostanMori(con.Convolve(p, r), con.Convolve(q, s), K))
}

func invFPS(val []int, mx int) []int {
	if mx == -1 {
		mx = len(val)
	}
	if val[0] == 0 {
		panic("invFPS")
	}
	g := make([]int, 1)
	g[0] = divMod(1, val[0])
	now := 1
	for now < mx {
		now <<= 1
		t := make([]int, len(val))
		copy(t, val)
		resize(&t, now)
		t = con.Convolve(t, g)
		for i := range t {
			t[i] = (MOD - t[i]) % MOD
		}
		t[0] = (t[0] + 2) % MOD
		g = con.Convolve(g, t)
		resize(&g, now)
	}
	resize(&g, mx)
	return g
}

func myPow(a []int, n int, m []int) []int {
	res := make([]int, 1)
	res[0] = 1
	for n > 0 {
		if (n & 1) != 0 {
			res = modPoly(con.Convolve(res, a), m)
		}
		tmp := con.Convolve(a, a)
		a = modPoly(tmp, m)
		n >>= 1
	}
	return res
}

func modPoly(lhs, rhs []int) []int {
	res := make([]int, len(lhs))
	copy(res, lhs)
	if len(res) < len(rhs) {
		return res
	}
	tmp0 := divPoly(res, rhs)
	tmp := con.Convolve(tmp0, rhs)
	res = minusPoly(res, tmp)
	resize(&res, len(rhs)-1)
	shrink(&res)
	return res
}

func minusPoly(lhs, rhs []int) []int {
	res := make([]int, max(len(lhs), len(rhs)))
	copy(res, lhs)
	for i := 0; i < len(rhs); i++ {
		res[i] = (res[i] - rhs[i] + MOD) % MOD
	}
	shrink(&res)
	return res
}

func divPoly(lhs, tmp []int) []int {
	res := make([]int, len(lhs))
	copy(res, lhs)
	rhs := make([]int, len(tmp))
	copy(rhs, tmp)
	if len(res) < len(rhs) {
		resize(&res, 0)
		return res
	}
	rsiz := len(res) - len(rhs) + 1
	reverseOrderInt(res)
	reverseOrderInt(rhs)
	resize(&res, rsiz)
	rhs = invPoly(rhs, rsiz)
	res = con.Convolve(res, rhs)
	resize(&res, rsiz)
	reverseOrderInt(res)
	return res
}

func invPoly(val []int, mx int) []int {
	g := make([]int, 1)
	g[0] = divMod(1, val[0])
	now := 1
	for now < mx {
		now <<= 1
		t := make([]int, len(val))
		copy(t, val)
		resize(&t, now)
		t = con.Convolve(t, g)
		for i := range t {
			t[i] = (MOD - t[i]) % MOD
		}
		t[0] += 2
		g = con.Convolve(g, t)
		resize(&g, now)
	}
	resize(&g, mx)
	return g
}

func shrink(val *[]int) {
	for i := len(*val) - 1; i > 0; i-- {
		if (*val)[i] == 0 {
			*val = (*val)[:len(*val)-1]
		} else {
			return
		}
	}
}

func BostanMori(f, g []int, k int) int {
	for k > 0 {
		if k%2 == 0 {
			f = Even(con.Convolve(f, Minus(g)))
		} else {
			f = Odd(con.Convolve(f, Minus(g)))
		}
		g = Even(con.Convolve(g, Minus(g)))
		k /= 2
	}
	return divMod(f[0], g[0])
}

func Even(f []int) []int {
	g := make([]int, (len(f)+1)/2)
	for i := 0; i < (len(f)+1)/2; i++ {
		g[i] = f[i*2]
	}
	return g
}

func Odd(f []int) []int {
	g := make([]int, len(f)/2)
	for i := 0; i < len(f)/2; i++ {
		g[i] = f[i*2+1]
	}
	return g
}

func Minus(f []int) []int {
	res := make([]int, len(f))
	copy(res, f)
	for i := 1; i < len(res); i += 2 {
		res[i] = (MOD - res[i]) % MOD
	}
	return res
}

const MOD = 998244353
const MOD1 = 754974721
const MOD2 = 167772161
const MOD3 = 469762049

func powMod(a, e, mod int) int {
	res, m := 1, a
	for e > 0 {
		if e&1 != 0 {
			res = res * m % mod
		}
		m = m * m % mod
		e >>= 1
	}
	return res
}

const ROOT = 3
const ROOT1 = 11
const ROOT2 = 3
const ROOT3 = 3

var con = NewConvolution(MOD, ROOT)
var con1 = NewConvolution(MOD1, ROOT1)
var con2 = NewConvolution(MOD2, ROOT2)
var con3 = NewConvolution(MOD3, ROOT3)

// 特殊な剰余と原始根
// (2, 1)
// (167772161, 3)
// (469762049, 3)
// (754974721, 11)
// (924844033, 5)
// (998244353, 3)
// (1012924417, 5)
// (167772161, 3)
// (469762049, 3)
// (1224736769, 3)
// (1107296257, 10)

type Convolution struct {
	mod, primroot, rank2                      int
	root, iroot, rate2, irate2, rate3, irate3 []int
}

func NewConvolution(mod, primroot int) *Convolution {
	rank2 := bits.TrailingZeros(uint(mod - 1))
	if rank2 < 3 {
		panic("Panic!")
	}
	root := make([]int, rank2+1)
	iroot := make([]int, rank2+1)
	rate2 := make([]int, rank2-2+1)
	irate2 := make([]int, rank2-2+1)
	rate3 := make([]int, rank2-3+1)
	irate3 := make([]int, rank2-3+1)
	root[rank2] = powMod(primroot, (mod-1)>>rank2, mod)
	iroot[rank2] = powMod(root[rank2], mod-2, mod)
	for i := rank2 - 1; i >= 0; i-- {
		root[i] = root[i+1] * root[i+1] % mod
		iroot[i] = iroot[i+1] * iroot[i+1] % mod
	}
	prod, iprod := 1, 1
	for i := 0; i <= rank2-2; i++ {
		rate2[i] = root[i+2] * prod % mod
		irate2[i] = iroot[i+2] * iprod % mod
		prod = prod * iroot[i+2] % mod
		iprod = iprod * root[i+2] % mod
	}
	prod, iprod = 1, 1
	for i := 0; i <= rank2-3; i++ {
		rate3[i] = root[i+3] * prod % mod
		irate3[i] = iroot[i+3] * iprod % mod
		prod = prod * iroot[i+3] % mod
		iprod = iprod * root[i+3] % mod
	}
	return &Convolution{mod, primroot, rank2, root, iroot, rate2, irate2, rate3, irate3}
}

func ceilPow2(n int) int {
	x := 0
	for 1<<x < n {
		x++
	}
	return x
}

func (q *Convolution) butterfly(a []int) {
	mod := q.mod
	n := len(a)
	h := ceilPow2(n)
	len := 0
	for len < h {
		if h-len == 1 {
			p := 1 << (h - len - 1)
			rot := 1
			for s := 0; s < (1 << len); s++ {
				offset := s << (h - len)
				for i := 0; i < p; i++ {
					l := a[i+offset]
					r := a[i+offset+p] * rot % mod
					a[i+offset] = (l + r) % mod
					a[i+offset+p] = (l - r + mod) % mod
				}
				if s+1 != (1 << len) {
					rot = rot * q.rate2[bits.TrailingZeros(^uint(s))] % mod
				}
			}
			len++
		} else {
			p := 1 << (h - len - 2)
			rot := 1
			imag := q.root[2]
			for s := 0; s < (1 << len); s++ {
				rot2 := rot * rot % mod
				rot3 := rot2 * rot % mod
				offset := s << (h - len)
				for i := 0; i < p; i++ {
					mod2 := mod * mod
					a0 := a[i+offset]
					a1 := a[i+offset+p] * rot
					a2 := a[i+offset+2*p] * rot2
					a3 := a[i+offset+3*p] * rot3
					a1na3imag := (a1 + mod2 - a3) % mod * imag
					na2 := mod2 - a2
					a[i+offset] = (a0 + a2 + a1 + a3) % mod
					a[i+offset+p] = (a0 + a2 + (2*mod2 - a1 - a3)) % mod
					a[i+offset+2*p] = (a0 + na2 + a1na3imag) % mod
					a[i+offset+3*p] = (a0 + na2 + (mod2 - a1na3imag)) % mod
				}
				if s+1 != (1 << len) {
					rot = rot * q.rate3[bits.TrailingZeros(^uint(s))] % mod
				}
			}
			len += 2
		}
	}
}

func (q *Convolution) butterflyInv(a []int) {
	mod := q.mod
	n := len(a)
	h := ceilPow2(n)
	len := h
	for len > 0 {
		if len == 1 {
			p := 1 << (h - len)
			irot := 1
			for s := 0; s < (1 << (len - 1)); s++ {
				offset := s << (h - len + 1)
				for i := 0; i < p; i++ {
					l := a[i+offset]
					r := a[i+offset+p]
					a[i+offset] = (l + r) % mod
					a[i+offset+p] = (l - r + mod) % mod
				}
				if s+1 != (1 << (len - 1)) {
					irot = irot * q.irate2[bits.TrailingZeros(^uint(s))] % mod
				}
			}
			len--
		} else {
			p := 1 << (h - len)
			irot := 1
			iimag := q.iroot[2]
			for s := 0; s < (1 << (len - 2)); s++ {
				irot2 := irot * irot % mod
				irot3 := irot2 * irot % mod
				offset := s << (h - len + 2)
				for i := 0; i < p; i++ {
					a0 := a[i+offset]
					a1 := a[i+offset+p]
					a2 := a[i+offset+2*p]
					a3 := a[i+offset+3*p]
					a2na3iimag := (mod + a2 - a3) * iimag % mod
					a[i+offset] = (a0 + a1 + a2 + a3) % mod
					a[i+offset+p] = (a0 + (mod - a1) + a2na3iimag) * irot % mod
					a[i+offset+2*p] = (a0 + a1 + (mod - a2) + (mod - a3)) * irot2 % mod
					a[i+offset+3*p] = (a0 + (mod - a1) + (mod - a2na3iimag)) * irot3 % mod
				}
				if s+1 != (1 << (len - 2)) {
					irot = irot * q.irate3[bits.TrailingZeros(^uint(s))] % mod
				}
			}
			len -= 2
		}
	}
	iz := powMod(n, mod-2, mod)
	for i := 0; i < n; i++ {
		a[i] = a[i] * iz % mod
	}
}

func (q *Convolution) convolveFFT(a []int, b []int) []int {
	mod := q.mod
	finalsz := len(a) + len(b) - 1
	z := 1
	for z < finalsz {
		z *= 2
	}
	lena, lenb := len(a), len(b)
	la := make([]int, z)
	lb := make([]int, z)
	for i := 0; i < lena; i++ {
		la[i] = a[i]
	}
	for i := 0; i < lenb; i++ {
		lb[i] = b[i]
	}
	q.butterfly(la)
	q.butterfly(lb)
	for i := 0; i < z; i++ {
		la[i] *= lb[i]
		la[i] %= mod
	}
	q.butterflyInv(la)
	return la[:finalsz]
}

func (q *Convolution) ConvolutionNaive(a []int, b []int) []int {
	mod := q.mod
	n := len(a)
	m := len(b)
	ans := make([]int, n+m-1)
	if n < m {
		for j := 0; j < m; j++ {
			for i := 0; i < n; i++ {
				ans[i+j] += a[i] * b[j] % mod
				ans[i+j] %= mod
			}
		}
	} else {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				ans[i+j] += a[i] * b[j] % mod
				ans[i+j] %= mod
			}
		}
	}
	return ans
}

func (q *Convolution) Convolve(a []int, b []int) []int {
	n := len(a)
	m := len(b)
	if n == 0 || m == 0 {
		return []int{}
	}
	if m < n {
		n = m
	}
	if n <= 60 {
		return q.ConvolutionNaive(a, b)
	} else {
		return q.convolveFFT(a, b)
	}
}

func ConvolveLL(a, b []int) []int {
	n := len(a)
	m := len(b)
	if n == 0 || m == 0 {
		return []int{}
	}

	MOD1 := 754974721 // 2^24
	MOD2 := 167772161 // 2^25
	MOD3 := 469762049 // 2^26
	M2M3 := MOD2 * MOD3
	M1M3 := MOD1 * MOD3
	M1M2 := MOD1 * MOD2
	M1M2M3 := MOD1 * MOD2 * MOD3

	i1 := inv_gcd(MOD2*MOD3, MOD1).y
	i2 := inv_gcd(MOD1*MOD3, MOD2).y
	i3 := inv_gcd(MOD1*MOD2, MOD3).y

	c1 := con1.Convolve(a, b)
	c2 := con2.Convolve(a, b)
	c3 := con3.Convolve(a, b)

	c := make([]int, n+m-1)
	for i := 0; i < n+m-1; i++ {
		x := 0
		x += (c1[i] * i1) % MOD1 * M2M3
		x += (c2[i] * i2) % MOD2 * M1M3
		x += (c3[i] * i3) % MOD3 * M1M2
		diff := c1[i] - safe_mod((x), (MOD1))
		if diff < 0 {
			diff += MOD1
		}
		offset := [5]int{0, 0, M1M2M3, 2 * M1M2M3, 3 * M1M2M3}
		x -= offset[diff%5]
		c[i] = x
	}

	return c
}

type pair struct{ x, y int }

func inv_gcd(a, b int) pair {
	a = safe_mod(a, b)
	if a == 0 {
		return pair{b, 0}
	}
	s := b
	t := a
	m0 := 0
	m1 := 1
	for t > 0 {
		u := s / t
		s -= t * u
		m0 -= m1 * u
		tmp := s
		s = t
		t = tmp
		tmp = m0
		m0 = m1
		m1 = tmp
	}
	if m0 < 0 {
		m0 += b / s
	}
	return pair{s, m0}
}

func safe_mod(x, m int) int {
	x %= m
	if x < 0 {
		x += m
	}
	return x
}

func resize(a *[]int, n int) {
	if len(*a) > n {
		*a = (*a)[:n]
	} else {
		n = n - len(*a)
		for i := 0; i < n; i++ {
			*a = append(*a, 0)
		}
	}
}

const mod = 998244353

func divMod(a, b int) int {
	ret := a * modInv(b)
	ret %= mod
	return ret
}

func modInv(a int) int {
	b, u, v := mod, 1, 0
	for b != 0 {
		t := a / b
		a -= t * b
		a, b = b, a
		u -= t * v
		u, v = v, u
	}
	u %= mod
	if u < 0 {
		u += mod
	}
	return u
}

func reverseOrderInt(a []int) {
	n := len(a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
