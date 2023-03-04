package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var M int
	fmt.Scan(&N, &M)

	initMod()

	var DC func(int, int) [4][]int
	DC = func(l, r int) [4][]int {
		if r-l == 1 {
			return [4][]int{
				{1, MOD - l},
				{0, (MOD - Fac[l]) * Fac[l] % MOD},
				{0, Inv[l] * Inv[l-1] % MOD},
				{0, 1},
			}
		}
		mid := (l + r) / 2
		L := DC(l, mid)
		R := DC(mid, r)
		siz := len(L[0]) + len(R[0]) - 1
		W := 1 << ceilPow2(siz)
		inv := inv_mod(W, MOD)
		for v := range L {
			resize(&L[v], W)
			con.butterfly(L[v])
		}
		for v := range R {
			resize(&R[v], W)
			con.butterfly(R[v])
		}
		var ans [4][]int
		for i := range ans {
			ans[i] = make([]int, W)
		}
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				if (i & 1) == (j >> 1) {
					k := (i & 2) | (j & 1)
					for n := 0; n < W; n++ {
						ans[k][n] = (ans[k][n] + R[i][n]*L[j][n]%MOD) % MOD
					}
				}
			}
		}
		for v := range ans {
			con.butterflyInv(ans[v])
			resize(&ans[v], siz)
			for x := range ans[v] {
				ans[v][x] = ans[v][x] * inv % MOD
			}
		}
		return ans
	}

	tmp := DC(1, M)
	f, g := tmp[0], tmp[1]
	resize(&f, N+1)
	fmt.Fprintln(out, MOD-con.Convolve(inv_fps(f), g)[N-1]%MOD)
}

func inv_fps(a []int) []int {
	n := len(a)
	ans := []int{1}
	for m := 1; m < n; m <<= 1 {
		siz := 2 * m
		f := make([]int, min(n, siz))
		copy(f, a[:min(n, siz)])
		g := make([]int, len(ans))
		copy(g, ans)
		resize(&f, siz)
		con.butterfly(f)
		resize(&g, siz)
		con.butterfly(g)
		for i := 0; i < siz; i++ {
			f[i] = f[i] * g[i] % MOD
		}
		con.butterflyInv(f)
		f = f[m:]
		resize(&f, siz)
		con.butterfly(f)
		for i := 0; i < siz; i++ {
			f[i] = f[i] * g[i] % MOD
		}
		con.butterflyInv(f)
		iz := inv_mod(siz, MOD)
		iz = iz * (MOD - iz) % MOD
		resize(&f, m)
		for x := range f {
			f[x] = f[x] * iz % MOD
		}
		ans = append(ans, f...)
	}
	resize(&ans, n)
	return ans
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

var N int

var Fac, Inv []int

func initMod() {
	Fac = make([]int, N+1)
	Inv = make([]int, N+1)
	Fac[0] = 1
	Inv[0] = 1
	for i := int(1); i <= N; i++ {
		Fac[i] = (Fac[i-1] * i) % MOD
		Inv[i] = invMod(Fac[i])
	}
}

func invMod(a int) int {
	return PowMod(a, MOD-2, MOD)
}

type pair struct{ x, y int }

func inv_mod(x, m int) int {
	if 1 > m {
		os.Exit(1)
	}
	z := inv_gcd(x, m)
	if z.x != 1 {
		os.Exit(1)
	}
	return z.y
}

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

var con = NewConvolution(MOD, 3)

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

func (q *Convolution) butterflyInv2(a []int) {
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
	iz := PowMod(n, mod-2, MOD)
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
	q.butterflyInv2(la)
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
