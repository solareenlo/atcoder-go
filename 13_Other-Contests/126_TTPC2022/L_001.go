package main

import (
	"fmt"
	"math/bits"
)

func main() {
	initMod()

	var n, m int
	fmt.Scan(&n, &m)

	X := make([]mint, 1<<20)
	for k := 0; k < m+1; k++ {
		X[k] = (((F[m] * I[k] % MOD) * I[m-k] % MOD) * I[m-k] % MOD) * mint(1-k%2*2+MOD) % MOD
	}
	con2.butterfly(X)
	for x := range X {
		X[x] = X[x].pow(n)
	}
	con2.butterflyInv(X)
	ans := mint(0)
	for k := 0; k < n*m+1; k++ {
		ans = (ans + X[k]*F[n*m-k]%MOD) % MOD
	}
	ans = ans * F[m].pow(n).div(1<<20) % MOD
	fmt.Println(ans)
}

const MOD = 998244353

func powMod(a, e, mod mint) mint {
	res, m := mint(1), a
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

var con2 = NewConvolution2(MOD, ROOT)

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
	root[rank2] = powMod(primroot, (mod-1)>>rank2, MOD)
	iroot[rank2] = powMod(root[rank2], mod-2, MOD)
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
	h := ceilPow2(n)
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
	h := ceilPow2(n)
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

func ceilPow2(n int) int {
	x := 0
	for (1 << uint(x)) < n {
		x++
	}
	return x
}

type mint int

func (m mint) pow(p int) mint {
	return PowMod(m, p)
}

func (m mint) inv() mint {
	return invMod(m)
}

func (m mint) div(n mint) mint {
	return divMod(m, n)
}

const VMAX = 1 << 20

var F, I [VMAX]mint

func initMod() {
	F[0] = 1
	I[0] = 1
	for i := mint(1); i < VMAX; i++ {
		F[i] = (F[i-1] * i) % MOD
		I[i] = invMod(F[i])
	}
}

func PowMod(a mint, n int) mint {
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
