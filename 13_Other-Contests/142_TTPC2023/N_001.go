package main

import (
	"fmt"
	"math/bits"
)

type fps []int
type Matrix [][]int
type Mat [][]int
type MatrixFps [][]fps

func mulmat(a, b Mat) Mat {
	ha := len(a)
	wa := len(a[0])
	hb := len(b)
	wb := len(b[0])
	if wa != hb {
		panic("wa != hb")
	}
	c := make(Mat, ha)
	for i := range c {
		c[i] = make([]int, wb)
	}
	for i := 0; i < ha; i++ {
		for k := 0; k < wa; k++ {
			for j := 0; j < wb; j++ {
				c[i][j] = (c[i][j] + a[i][k]*b[k][j]%MOD) % MOD
			}
		}
	}
	return c
}

func get_Mat(n int) Mat {
	res := make(Mat, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	return res
}

const m0 = 167772161
const m1 = 469762049
const m2 = 754974721
const mint0 = m0
const mint1 = m1
const mint2 = m2

var r01 = inv_gcd(m0, m1).y
var r02 = inv_gcd(m0, m2).y
var r12 = inv_gcd(m1, m2).y

var r02r12 = r02 * r12 % m2

const w1 = m0
const w2 = m0 * m1

func mul(a, b []int, mod int) []int {
	s := make([]int, len(a))
	t := make([]int, len(b))
	for i := 0; i < len(a); i++ {
		s[i] = a[i] % mod
	}
	for i := 0; i < len(b); i++ {
		t[i] = b[i] % mod
	}
	switch mod {
	case m0:
		return con2.Convolve(s, t)
	case m1:
		return con3.Convolve(s, t)
	case m2:
		return con1.Convolve(s, t)
	}
	return []int{}
}

func multiplyMod(s, t []int, mod int) []int {
	d0 := mul(s, t, m0)
	d1 := mul(s, t, m1)
	d2 := mul(s, t, m2)
	n := len(d0)
	ret := make([]int, n)
	W1 := w1 % mod
	W2 := w2 % mod
	for i := 0; i < n; i++ {
		n1 := d1[i]
		n2 := d2[i]
		a := d0[i]
		b := (n1 + m1 - a) * r01 % m1
		c := ((n2+m2-a)*r02r12 + (m2-b)*r12) % m2
		ret[i] = (a + b*W1 + c*W2) % mod
	}
	return ret
}

func multiply(a, b []int) []int {
	if len(a) == 0 && len(b) == 0 {
		return []int{}
	}
	if min(len(a), len(b)) < 128 {
		ret := make([]int, len(a)+len(b)-1)
		for i := 0; i < len(a); i++ {
			for j := 0; j < len(b); j++ {
				ret[i+j] = (ret[i+j] + a[i]*b[j]%MOD) % MOD
			}
		}
		return ret
	}
	s := make([]int, len(a))
	t := make([]int, len(b))
	copy(s, a)
	copy(t, b)
	u := multiplyMod(s, t, MOD)
	ret := make([]int, len(u))
	copy(ret, u)
	return ret
}

func eval(f fps, x int) int {
	res := 0
	p := 1
	for _, c := range f {
		res = (res + c*p%MOD) % MOD
		p = p * x % MOD
	}
	return res
}

type binomial struct {
	facts, ifacts []int
}

func (b *binomial) init() {
	Len := 300000
	b.facts = make([]int, Len+1)
	b.ifacts = make([]int, Len+1)
	for i := 0; i < Len+1; i++ {
		b.facts[i] = 1
		b.ifacts[i] = 1
	}
	for i := 1; i <= Len; i++ {
		b.facts[i] = b.facts[i-1] * i % MOD
	}
	b.ifacts[Len] = invMod(b.facts[Len], MOD)
	for i := Len - 1; i >= 1; i-- {
		b.ifacts[i] = b.ifacts[i+1] * (i + 1) % MOD
	}
}

func (b binomial) fact(n int) int  { return b.facts[n] }
func (b binomial) ifact(n int) int { return b.ifacts[n] }
func (b binomial) inv(n int) int   { return b.ifacts[n] * b.facts[n-1] % MOD }

func sample_point_shift(y fps, t, m int) fps {
	T := t % MOD
	k := len(y) - 1
	if T <= k {
		ret := make(fps, m)
		ptr := 0
		for i := T; i <= k && ptr < m; i++ {
			ret[ptr] = y[i]
			ptr++
		}
		if k+1 < T+m {
			suf := sample_point_shift(y, k+1, m-ptr)
			for i := k + 1; i < T+m; i++ {
				ret[ptr] = suf[i-(k+1)]
				ptr++
			}
		}
		return ret
	}
	if T+m > MOD {
		pref := sample_point_shift(y, T, MOD-T)
		suf := sample_point_shift(y, 0, m-len(pref))
		pref = append(pref, suf...)
		return pref
	}
	d := make(fps, k+1)
	var bnm binomial
	bnm.init()
	for i := 0; i <= k; i++ {
		d[i] = (bnm.ifact(i) * bnm.ifact(k-i) % MOD) * y[i] % MOD
		if ((k - i) & 1) != 0 {
			d[i] = MOD - d[i]
		}
	}
	fact := make([]int, m+k+1)
	fact[0] = 1
	for i := 0; i < m+k; i++ {
		fact[i+1] = fact[i] * (T - k + i) % MOD
	}
	h := make(fps, m+k)
	h[m+k-1] = invMod(fact[m+k], MOD)
	for i := m + k - 1; i >= 1; i-- {
		h[i-1] = h[i] * (T - k + i) % MOD
	}
	for i := 0; i < m+k; i++ {
		h[i] = h[i] * fact[i] % MOD
	}
	dh := multiply(d, h)
	ret := make(fps, m)
	cur := T
	for i := 1; i <= k; i++ {
		cur = cur * (T - i) % MOD
	}
	for i := 0; i < m; i++ {
		ret[i] = cur * dh[k+i] % MOD
		cur = cur * (T + i + 1) % MOD
		cur = cur * h[i] % MOD
	}
	return ret
}

func polynomial_matrix_product(m MatrixFps, k int) Mat {
	var shift func([]Mat, int) []Mat
	shift = func(G []Mat, x int) []Mat {
		d := len(G)
		n := len(G[0])
		H := make([]Mat, d)
		for i := range H {
			H[i] = get_Mat(n)
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				g := make(fps, d)
				for l := 0; l < d; l++ {
					g[l] = G[l][i][j]
				}
				h := sample_point_shift(g, x, d)
				for l := 0; l < d; l++ {
					H[l][i][j] = h[l]
				}
			}
		}
		return H
	}
	n := len(m)
	deg := 1
	for _, X := range m {
		for _, x := range X {
			deg = max(deg, len(x)-1)
		}
	}
	for (deg & (deg - 1)) != 0 {
		deg++
	}
	G := make([]Mat, deg+1)
	v := 1
	for deg*v*v < k {
		v *= 2
	}
	iv := invMod(v, MOD)
	for i := 0; i < len(G); i++ {
		x := v * i % MOD
		mat := get_Mat(n)
		for j := 0; j < n; j++ {
			for l := 0; l < n; l++ {
				mat[j][l] = eval(m[j][l], x)
			}
		}
		G[i] = mat
	}
	for w := 1; w != v; w <<= 1 {
		W := w % MOD
		G1 := shift(G, W*iv%MOD)
		G2 := shift(G, ((((W*deg%MOD)*v%MOD)+v)%MOD)*iv%MOD)
		G3 := shift(G, (((((W*deg%MOD)*v%MOD)+v)%MOD+W)%MOD)*iv%MOD)
		for i := 0; i <= w*deg; i++ {
			G[i] = mulmat(G1[i], G[i])
			G2[i] = mulmat(G3[i], G2[i])
		}
		G = append(G, G2[:len(G2)-1]...)
	}
	res := get_Mat(n)
	for i := 0; i < n; i++ {
		res[i][i] = 1
	}
	i := 0
	for i+v <= k {
		res = mulmat(G[i/v], res)
		i += v
	}
	for i < k {
		mat := get_Mat(n)
		for j := 0; j < n; j++ {
			for l := 0; l < n; l++ {
				mat[j][l] = eval(m[j][l], i)
			}
		}
		res = mulmat(mat, res)
		i++
	}
	return res
}

func kth_term_of_p_recursive(a []int, fs []fps, k int) int {
	deg := len(fs) - 1
	if deg < 1 {
		panic("deg < 1")
	}
	m := make(MatrixFps, deg)
	for i := range m {
		m[i] = make([]fps, deg)
	}
	denom := make(MatrixFps, 1)
	denom[0] = make([]fps, 1)
	for i := 0; i < deg; i++ {
		m[0][i] = fs[i+1]
		for x := range m[0][i] {
			m[0][i][x] = MOD - m[0][i][x]
		}
	}
	for i := 1; i < deg; i++ {
		m[i][i-1] = fs[0]
	}
	denom[0][0] = fs[0]
	a0 := make(Mat, deg)
	for i := range a0 {
		a0[i] = make([]int, deg)
	}
	for i := 0; i < deg; i++ {
		a0[i][0] = a[deg-1-i]
	}
	res := mulmat(polynomial_matrix_product(m, k-deg+1), a0)[0][0]
	res = divMod(res, polynomial_matrix_product(denom, k-deg+1)[0][0])
	return res
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	MOD = m
	con = NewConvolution(MOD, primitive_root(MOD))
	f2 := fps{15, 8, 1}
	f1 := fps{MOD - 159, MOD - 124, MOD - 17}
	f0 := fps{216, 468, 72}
	a := []int{1, 4, 28, 224, 1888, 16320, 143040, 1264128}
	fs := []fps{f2, f1, f0}
	fmt.Println(kth_term_of_p_recursive(a, fs, n))
}

var MOD int

// const MOD = 998244353
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

func primitive_root(m int) int {
	if m == 2 {
		return 1
	}
	if m == 167772161 {
		return 3
	}
	if m == 469762049 {
		return 3
	}
	if m == 754974721 {
		return 11
	}
	if m == 998244353 {
		return 3
	}
	var divs [20]int
	divs[0] = 2
	cnt := 1
	x := (m - 1) / 2
	for x%2 == 0 {
		x /= 2
	}
	for i := 3; i*i <= x; i += 2 {
		if x%i == 0 {
			divs[cnt] = i
			cnt++
			for x%i == 0 {
				x /= i
			}
		}
	}
	if x > 1 {
		divs[cnt] = x
		cnt++
	}
	for g := 2; ; g++ {
		ok := true
		for i := 0; i < cnt; i++ {
			if pow_mod(g, (m-1)/divs[i], m) == 1 {
				ok = false
				break
			}
		}
		if ok {
			return g
		}
	}
}

func pow_mod(x, n, m int) int {
	if m == 1 {
		return 0
	}
	_m := uint(m)
	r := uint(1)
	y := uint(safe_mod(x, m))
	for n > 0 {
		if n&1 != 0 {
			r = (r * y) % _m
		}
		y = (y * y) % _m
		n >>= 1
	}
	return int(r)
}

const ROOT = 3
const ROOT1 = 11
const ROOT2 = 3
const ROOT3 = 3

// var con = NewConvolution(MOD, ROOT)
var con *Convolution
var con1 = NewConvolution(MOD1, primitive_root(MOD1))
var con2 = NewConvolution(MOD2, primitive_root(MOD2))
var con3 = NewConvolution(MOD3, primitive_root(MOD2))

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
	rank2 := countr_zero(uint(mod - 1))
	if rank2 < 3 {
		rank2 = 3
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

func countr_zero(n uint) int {
	x := 0
	for (n & (1 << x)) == 0 {
		x++
	}
	return x
}

func bit_ceil(n uint) uint {
	x := uint(1)
	for x < uint(n) {
		x *= 2
	}
	return x
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
					a1na3imag := ((a1 + mod2 - a3) % mod) * imag
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
				ans[i+j] = (ans[i+j] + a[i]*b[j]%mod) % mod
			}
		}
	} else {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				ans[i+j] = (ans[i+j] + a[i]*b[j]%mod) % mod
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func invMod(a, mod int) int {
	return powMod(a, mod-2, mod)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func divMod(a, b int) int {
	ret := a * modInv(b)
	ret %= MOD
	return ret
}

func modInv(a int) int {
	b, u, v := MOD, 1, 0
	for b != 0 {
		t := a / b
		a -= t * b
		a, b = b, a
		u -= t * v
		u, v = v, u
	}
	u %= MOD
	if u < 0 {
		u += MOD
	}
	return u
}
