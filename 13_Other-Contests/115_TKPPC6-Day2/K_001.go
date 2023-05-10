package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	var dacm divide_and_conquer_method
	dacm.v = make([]int, n)
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(in, &x)
		dacm.v[i] = x
	}

	var p Poly = dacm.calc()
	x := make([]int, m+1)
	for i := 0; i < m+1; i++ {
		x[i] = i
	}

	var t Tree
	t.init(x)
	sub := t.multi_eval(p)
	ans := 1
	for i := 1; i < m+1; i++ {
		fmt.Println(ans)
		ans = ans * sub[i] % MOD
		pw := mint(PowMod(i, n, MOD))
		ans = ans * int(pw.inv()) % MOD
	}
}

type divide_and_conquer_method struct {
	v []int
}

func (dacm divide_and_conquer_method) dfs(l, r int) []int {
	if r-l <= 1 {
		vv := make([]int, 2)
		vv[0] = dacm.v[l]
		vv[1] = 1
		return vv
	}
	c := (l + r) / 2
	vl := dacm.dfs(l, c)
	vr := dacm.dfs(c, r)
	return con.Convolve(vl, vr)
}

func (dacm divide_and_conquer_method) calc() []int { return dacm.dfs(0, len(dacm.v)) }

type Poly []int

func mulPoly(a, b Poly) Poly {
	return con.Convolve(a, b)
}

func mid_prod(a, b Poly) Poly {
	if min(len(b), len(a)-len(b)+1) <= 60 {
		res := make(Poly, len(a)-len(b)+1)
		for i := 0; i < len(res); i++ {
			res[i] = inner_product(b, a[i:], 0)
		}
		return res
	}
	n := 1 << log2(2*len(a)-1)
	fa := make([]int, n)
	fb := make([]int, n)
	copy(fa, a)
	for i := 0; i < len(b); i++ {
		fb[i] = b[len(b)-1-i]
	}
	con2.butterfly(fa)
	con2.butterfly(fb)
	for i := 0; i < n; i++ {
		fa[i] = fa[i] * fb[i] % MOD
	}
	con2.butterflyInv(fa)
	resize(&fa, len(a))
	fa = fa[len(b)-1:]
	div(fa, n)
	return fa
}

type Tree struct {
	m int
	f []Poly
}

func (t *Tree) init(x []int) {
	t.m = len(x)
	t.f = make([]Poly, 2<<log2(2*t.m-1))
	for i := 0; i < t.size(); i++ {
		if i < t.m {
			t.f[t.size()+i] = Poly{x[i]}
		} else {
			t.f[t.size()+i] = Poly{0}
		}
	}
	for i := t.size() - 1; i >= 1; i-- {
		resizePoly(&(t.f[i]), 2*len(t.f[2*i]))
		t.f[i] = plus(t.f[i], plus(t.f[2*i], t.f[2*i+1]))
		prod := mulPoly(t.f[2*i], t.f[2*i+1])
		for j := 0; j < len(prod); j++ {
			t.f[i][j+1] = (t.f[i][j+1] - prod[j] + MOD) % MOD
		}
	}
}

func (t Tree) size() int { return len(t.f) / 2 }

func (t Tree) get(i int) Poly {
	res := make(Poly, len(t.f[i])+1)
	for j := 0; j < len(res); j++ {
		if j != 0 {
			res[j] = (MOD - t.f[i][j-1]) % MOD
		} else {
			res[j] = 1
		}
	}
	return res
}

func (tree Tree) multi_eval(a Poly) []int {
	n := len(a)
	resizePoly(&a, 2*n-1)
	t := make([]Poly, 2*tree.size())
	den := tree.get(1)
	resizePoly(&den, n)
	t[1] = mid_prod(a, inv(den))
	resizePoly(&t[1], tree.size())
	for i := 1; i < tree.size(); i++ {
		t[2*i] = mid_prod(t[i], tree.get(2*i+1))
		t[2*i+1] = mid_prod(t[i], tree.get(2*i))
	}
	res := make([]int, tree.m)
	for i := 0; i < tree.m; i++ {
		res[i] = t[tree.size()+i][0]
	}
	return res
}

func plus(a, b []int) []int {
	res := make([]int, max(len(a), len(b)))
	copy(res, a)
	for i := range b {
		res[i] = (res[i] + b[i]) % MOD
	}
	return res
}

func minus(a, b []int) []int {
	res := make([]int, max(len(a), len(b)))
	copy(res, a)
	for i := range b {
		res[i] = (res[i] - b[i] + MOD) % MOD
	}
	return res
}

func div(a []int, b int) {
	b = int(mint(b).inv())
	for i := range a {
		a[i] = a[i] * b % MOD
	}
}

func mul(a, b []int) []int {
	res := con.Convolve(a, b)
	resize(&res, max(len(a), len(b)))
	return res
}

func inv(a []int) []int {
	var res []int = []int{int(mint(a[0]).inv())}
	for len(res) < len(a) {
		resize(&res, min(2*len(res), len(a)))
		res = mul(res, minus([]int{2}, mul(a[:len(res)], res)))
	}
	return res
}

const MOD = 998244353

func PowMod(a, e, mod int) int {
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

var con = NewConvolution(MOD, ROOT)

// 特殊な剰余と原始根
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
	root[rank2] = PowMod(primroot, (mod-1)>>rank2, mod)
	iroot[rank2] = PowMod(root[rank2], mod-2, mod)
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
	iz := PowMod(n, mod-2, mod)
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

var con2 = NewConvolution2(MOD, ROOT)

type Convolution2 struct {
	mod, primroot, rank2                      int
	root, iroot, rate2, irate2, rate3, irate3 []int
}

func NewConvolution2(mod, primroot int) *Convolution2 {
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
	root[rank2] = PowMod(primroot, (mod-1)>>rank2, MOD)
	iroot[rank2] = PowMod(root[rank2], mod-2, MOD)
	for i := rank2 - 1; i >= 0; i-- {
		root[i] = root[i+1] * root[i+1] % mod
		iroot[i] = iroot[i+1] * iroot[i+1] % mod
	}
	prod, iprod := 1, 1
	for i := 0; i <= int(rank2-2); i++ {
		rate2[i] = root[i+2] * prod % mod
		irate2[i] = iroot[i+2] * iprod % mod
		prod = prod * iroot[i+2] % mod
		iprod = iprod * root[i+2] % mod
	}
	prod, iprod = 1, 1
	for i := 0; i <= int(rank2-3); i++ {
		rate3[i] = root[i+3] * prod % mod
		irate3[i] = iroot[i+3] * iprod % mod
		prod = prod * iroot[i+3] % mod
		iprod = iprod * root[i+3] % mod
	}
	return &Convolution2{mod, primroot, rank2, root, iroot, rate2, irate2, rate3, irate3}
}

func (q *Convolution2) butterfly(a []int) {
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
					a1 := a[i+offset+p] * rot % mod
					a2 := a[i+offset+2*p] * rot2 % mod
					a3 := a[i+offset+3*p] * rot3 % mod
					a1na3imag := (a1 + mod2 - a3) % mod * imag % mod
					na2 := (mod2 - a2 + mod) % mod
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

func (q *Convolution2) butterflyInv(a []int) {
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
					a[i+offset+p] = ((l - r + mod) % mod) * irot % mod
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
					a2na3iimag := ((a2 - a3 + mod) % mod) * iimag % mod
					a[i+offset] = ((a0+a1)%mod + (a2+a3)%mod) % mod
					a[i+offset+p] = ((a0 - a1 + mod + a2na3iimag) % mod) * irot % mod
					a[i+offset+2*p] = ((a0 + a1 - a2 - a3 + mod + mod) % mod) * irot2 % mod
					a[i+offset+3*p] = ((a0-a1+mod)%mod + (mod-a2na3iimag)%mod) * irot3 % mod
				}
				if s+1 != (1 << (len - 2)) {
					irot = irot * q.irate3[bits.TrailingZeros(^uint(s))] % mod
				}
			}
			len -= 2
		}
	}
}

func log2(n int) int {
	var k int
	for k = 0; n != 0; n >>= 1 {
		k++
	}
	return k - 1
}

type mint int

func (m mint) pow(p int) mint {
	return POWMOD(m, p)
}

func (m mint) inv() mint {
	return invMod(m)
}

func (m mint) div(n mint) mint {
	return divMod(m, n)
}

func POWMOD(a mint, n int) mint {
	res := mint(1)
	for n > 0 {
		if n%2 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		n /= 2
	}
	return res
}

func invMod(a mint) mint {
	return POWMOD(a, MOD-2)
}

func divMod(a, b mint) mint {
	ret := a * modInv(b)
	ret %= MOD
	return ret
}

func modInv(a mint) mint {
	b, u, v := mint(MOD), mint(1), mint(0)
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

func resizePoly(a *Poly, n int) {
	if len(*a) > n {
		*a = (*a)[:n]
	} else {
		n = n - len(*a)
		for i := 0; i < n; i++ {
			*a = append(*a, 0)
		}
	}
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func inner_product(a, b []int, init int) int {
	for i := 0; i < len(a); i++ {
		init = (init + a[i]*b[i]%MOD) % MOD
	}
	return init
}
