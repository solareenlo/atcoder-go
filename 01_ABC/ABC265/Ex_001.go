package main

import (
	"fmt"
	"math/bits"
)

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	p := 1
	for p < h*(w-2)+1 {
		p <<= 1
	}

	const L = 5
	n := p << L
	f := make([]int, n)

	var get func(int, int) int
	get = func(x, d int) int {
		return (x * p) | d
	}

	for i := 0; i < w; i++ {
		for j := i + 1; j < w; j++ {
			f[get(0, i+j-1)]++
			f[get(j-i-1, w-2)]++
		}
	}

	c := NewConvolution(mod, 3)

	var multi_dim_fft func(bool)
	multi_dim_fft = func(inv bool) {
		for k := 0; k != n; k += p {
			g := make([]int, p)
			for i := range g {
				g[i] = f[k+i]
			}
			if inv {
				c.butterflyInv(g)
			} else {
				c.butterfly(g)
			}
			for i := range g {
				f[k+i] = g[i]
			}
		}
		for k := p; k != n; k <<= 1 {
			for l := 0; l+2*k <= n; l += 2 * k {
				for i := 0; i < k; i++ {
					u := f[l+i]
					v := f[l+k+i]
					f[l+i] = (u + v) % mod
					f[l+k+i] = (u - v + mod) % mod
				}
			}
		}
	}

	multi_dim_fft(false)
	for i := range f {
		f[i] = powMod(f[i], h, mod)
	}
	multi_dim_fft(true)

	z := invMod(n)
	pos := powMod(w*(w-1), h, mod)
	zero := 0
	for x := 0; x < 1<<L; x++ {
		if x != 0 {
			zero += (f[get(x, h*(w-2))] * z) % mod
			zero %= mod
		}
		pos -= (f[get(x, h*(w-2))] * z) % mod
		pos = (pos + mod) % mod
	}
	fmt.Println((divMod(pos, 2) + zero) % mod)
}

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

const mod = 998244353

func invMod(a int) int {
	return powMod(a, mod-2, mod)
}

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

var con = NewConvolution(mod, 3)

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
