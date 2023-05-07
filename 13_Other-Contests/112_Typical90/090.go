package main

import (
	"fmt"
	"math/bits"
)

func main() {
	initMod()

	var n, k int
	fmt.Scan(&n, &k)
	var f, X Poly
	f.init(0, 1)
	X.init(1, 1)
	var tmp Poly
	tmp.init(0, 1)
	for h, i := k, 0; h > 0; h = i {
		m := k/h + 1
		i = k / m
		X[1] = mint(h - i)
		tmp2 := mul(f.minus(), X)
		tmp1 := plus(tmp2, tmp)
		f = mul(f, tmp1.inv(m))
		resize(&f, m)
	}
	X[1] = 1
	fmt.Println(plus(mul(f.minus(), X), tmp).div_bm(f, n))
}

type Poly []mint

func (p *Poly) init(n int, a mint) {
	*p = make([]mint, n+1)
	(*p)[n] = a
}

func (p *Poly) initVec(v []mint) {
	*p = make([]mint, len(v))
	copy(*p, v)
}

func resize(a *Poly, n int) {
	if len(*a) > n {
		*a = (*a)[:n]
	} else {
		n = n - len(*a)
		for i := 0; i < n; i++ {
			*a = append(*a, 0)
		}
	}
}

func plus(l, r Poly) Poly {
	maxi := max(len(l), len(r))
	var ret Poly = make([]mint, maxi)
	copy(ret, l)
	for i := range r {
		ret[i] = (ret[i] + r[i]) % Mod
	}
	return ret
}

func (p Poly) minus() Poly {
	ret := make([]mint, len(p))
	for i := range ret {
		ret[i] = (Mod - p[i]) % Mod
	}
	return ret
}

func (p *Poly) fft() {
	con2.butterfly(*p)
}

func (p *Poly) ifft() {
	con2.butterflyInv(*p)
}

func mul(l, r Poly) Poly {
	return con1.Convolve(l, r)
}

func dfft(f *Poly, z mint) {
	var f2 Poly
	f2 = make(Poly, len(*f))
	copy(f2, *f)
	f.fft()
	dfft1(f, f2, z)
}

func dfft1(f *Poly, f2 Poly, z mint) {
	d := len(*f)
	x := mint(1)
	for i := 0; i < d; i++ {
		f2[i] = f2[i] * x % Mod
		x = x * z % Mod
	}
	f2.fft()
	*f = append(*f, f2...)
}

func (p Poly) inv(m int) Poly {
	var h Poly
	h.init(m-1, 0)
	h[0] = p[0].inv()
	z := mint(ROOT)
	for d, c := 1, 1; d < m; d, c = d+d, c+1 {
		var f Poly
		f.init(d+d-1, 0)
		for i := 0; i < min(d+d, len(p)); i++ {
			f[i] = p[i]
		}
		f.fft()
		var g Poly
		g.initVec(h[:d])
		id := mint(Mod >> c)
		dfft(&g, z.pow(id))
		for i := 0; i < d+d; i++ {
			f[i] = f[i] * g[i] % Mod
		}
		f.ifft()
		for i := 0; i < d; i++ {
			f[i] = 0
		}
		f.fft()
		for i := 0; i < d+d; i++ {
			f[i] = f[i] * g[i] % Mod
		}
		f.ifft()
		x := id
		x = x * x % Mod
		for i := d; i < min(d+d, m); i++ {
			h[i] = ((Mod - f[i]) % Mod) * x % Mod
		}
	}
	return h
}

func (p Poly) div_bm(G Poly, n int) mint {
	d, c := 1, 1
	for ; d < len(p); d += d {
		c++
	}
	if m := len(G); d < m {
		k := m - d + 1
		H := mul(G, p.inv(k))
		if n < k {
			return H[n]
		}
		H = mul(p, H[:k])
		for i := k; i < m; i++ {
			if i < len(H) {
				G[i-k] = (G[i] - H[i] + Mod) % Mod
			} else {
				G[i-k] = G[i]
			}
		}
		G[d-1] = 0
		n -= k
	}
	resize(&G, d)
	var H Poly
	H.init(d-1, 0)
	copy(H, p)
	var g, h Poly
	g.init(len(G)-1, 0)
	copy(g, G)
	h.init(len(H)-1, 0)
	copy(h, H)
	g.fft()
	h.fft()
	rev := make([]int, d)
	for i := 1; i < d; i++ {
		if (i & 1) != 0 {
			rev[i] = rev[i^1] | d
		} else {
			rev[i] = rev[i>>1] >> 1
		}
	}
	z := mint(ROOT)
	id := mint(Mod >> c)
	i2 := mint(Mod>>1 | 1)
	x := Mod - (id << 1)
	z = z.pow(id)
	iz := z.inv()
	for ; n >= d; n >>= 1 {
		dfft1(&g, G, z)
		dfft1(&h, H, z)
		if (n & 1) != 0 {
			y := i2
			for _, i := range rev {
				G[i>>1] = ((g[i]*h[i|1]%Mod - h[i]*g[i|1]%Mod + Mod) % Mod) * y % Mod
				y = y * iz % Mod
			}
		} else {
			for i := 0; i < d; i++ {
				G[i] = ((g[i<<1]*h[i<<1|1]%Mod + h[i<<1]*g[i<<1|1]%Mod) % Mod) * i2 % Mod
			}
		}
		for i := 0; i < d; i++ {
			H[i] = h[i<<1] * h[i<<1|1] % Mod
		}
		g = make([]mint, len(G))
		copy(g, G)
		h = make([]mint, len(H))
		copy(h, H)
		G.ifft()
		H.ifft()
		for i := 0; i < d; i++ {
			G[i] = G[i] * x % Mod
			H[i] = H[i] * x % Mod
		}
	}
	H = H.inv(n + 1)
	r := mint(0)
	for i := 0; i < n+1; i++ {
		r = (r + G[i]*H[n-i]%Mod) % Mod
	}
	return r
}

const Mod = 998244353

const ROOT = 3

var con1 = NewConvolution1(Mod, ROOT)

// 特殊な剰余と原始根
// (924844033, 5)
// (998244353, 3)
// (1012924417, 5)
// (167772161, 3)
// (469762049, 3)
// (1224736769, 3)
// (1107296257, 10)

type Convolution1 struct {
	mod, primroot, rank2                      mint
	root, iroot, rate2, irate2, rate3, irate3 []mint
}

func NewConvolution1(mod, primroot mint) *Convolution1 {
	rank2 := mint(bits.TrailingZeros(uint(mod - 1)))
	if rank2 < 3 {
		panic("Panic!")
	}
	root := make([]mint, rank2+1)
	iroot := make([]mint, rank2+1)
	rate2 := make([]mint, rank2-2+1)
	irate2 := make([]mint, rank2-2+1)
	rate3 := make([]mint, rank2-3+1)
	irate3 := make([]mint, rank2-3+1)
	root[rank2] = PowMod(primroot, (mod-1)>>rank2)
	iroot[rank2] = PowMod(root[rank2], mod-2)
	for i := rank2 - 1; i >= 0; i-- {
		root[i] = root[i+1] * root[i+1] % mod
		iroot[i] = iroot[i+1] * iroot[i+1] % mod
	}
	prod, iprod := mint(1), mint(1)
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
	return &Convolution1{mod, primroot, rank2, root, iroot, rate2, irate2, rate3, irate3}
}

func CeilPow2(n int) int {
	x := 0
	for 1<<x < n {
		x++
	}
	return x
}

func (q *Convolution1) butterfly(a []mint) {
	mod := q.mod
	n := len(a)
	h := CeilPow2(n)
	len := 0
	for len < h {
		if h-len == 1 {
			p := 1 << (h - len - 1)
			rot := mint(1)
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
			rot := mint(1)
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

func (q *Convolution1) butterflyInv(a []mint) {
	mod := q.mod
	n := len(a)
	h := CeilPow2(n)
	len := h
	for len > 0 {
		if len == 1 {
			p := 1 << (h - len)
			irot := mint(1)
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
			irot := mint(1)
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
	iz := PowMod(mint(n), mod-2)
	for i := 0; i < n; i++ {
		a[i] = a[i] * iz % mod
	}
}

func (q *Convolution1) convolveFFT(a []mint, b []mint) []mint {
	mod := q.mod
	finalsz := len(a) + len(b) - 1
	z := 1
	for z < finalsz {
		z *= 2
	}
	lena, lenb := len(a), len(b)
	la := make([]mint, z)
	lb := make([]mint, z)
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

func (q *Convolution1) ConvolutionNaive(a []mint, b []mint) []mint {
	mod := q.mod
	n := len(a)
	m := len(b)
	ans := make([]mint, n+m-1)
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

func (q *Convolution1) Convolve(a, b []mint) []mint {
	n := len(a)
	m := len(b)
	if n == 0 || m == 0 {
		return []mint{}
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

var con2 = NewConvolution2(Mod, ROOT)

type Convolution2 struct {
	mod, primroot, rank2                      mint
	root, iroot, rate2, irate2, rate3, irate3 []mint
}

func NewConvolution2(mod, primroot mint) *Convolution2 {
	rank2 := mint(bits.TrailingZeros(uint(mod - 1)))
	if rank2 < 3 {
		panic("Panic!")
	}
	root := make([]mint, rank2+1)
	iroot := make([]mint, rank2+1)
	rate2 := make([]mint, rank2-2+1)
	irate2 := make([]mint, rank2-2+1)
	rate3 := make([]mint, rank2-3+1)
	irate3 := make([]mint, rank2-3+1)
	root[rank2] = PowMod(primroot, (mod-1)>>rank2)
	iroot[rank2] = PowMod(root[rank2], mod-2)
	for i := rank2 - 1; i >= 0; i-- {
		root[i] = root[i+1] * root[i+1] % mod
		iroot[i] = iroot[i+1] * iroot[i+1] % mod
	}
	prod, iprod := mint(1), mint(1)
	for i := 0; i <= int(rank2-2); i++ {
		rate2[i] = root[i+2] * prod % mod
		irate2[i] = iroot[i+2] * iprod % mod
		prod = prod * iroot[i+2] % mod
		iprod = iprod * root[i+2] % mod
	}
	prod, iprod = mint(1), mint(1)
	for i := 0; i <= int(rank2-3); i++ {
		rate3[i] = root[i+3] * prod % mod
		irate3[i] = iroot[i+3] * iprod % mod
		prod = prod * iroot[i+3] % mod
		iprod = iprod * root[i+3] % mod
	}
	return &Convolution2{mod, primroot, rank2, root, iroot, rate2, irate2, rate3, irate3}
}

func (q *Convolution2) butterfly(a []mint) {
	mod := q.mod
	n := len(a)
	h := CeilPow2(n)
	len := 0
	for len < h {
		if h-len == 1 {
			p := 1 << (h - len - 1)
			rot := mint(1)
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
			rot := mint(1)
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

func (q *Convolution2) butterflyInv(a []mint) {
	mod := q.mod
	n := len(a)
	h := CeilPow2(n)
	len := h
	for len > 0 {
		if len == 1 {
			p := 1 << (h - len)
			irot := mint(1)
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
			irot := mint(1)
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

type mint int

func (m mint) pow(p mint) mint {
	return PowMod(m, p)
}

func (m mint) inv() mint {
	return invMod(m)
}

func (m mint) div(n mint) mint {
	return DivMod(m, n)
}

const MOD = 998244353
const VMAX = 100005

var fact, invf [VMAX]mint

func initMod() {
	fact[0] = 1
	invf[0] = 1
	for i := mint(1); i < VMAX; i++ {
		fact[i] = (fact[i-1] * i) % MOD
		invf[i] = invMod(fact[i])
	}
}

func PowMod(a, n mint) mint {
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
	return PowMod(a, MOD-2)
}

func DivMod(a, b mint) mint {
	ret := a * ModInv(b)
	ret %= MOD
	return ret
}

func ModInv(a mint) mint {
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
