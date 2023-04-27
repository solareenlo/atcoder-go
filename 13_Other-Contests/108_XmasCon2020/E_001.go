package main

import (
	"fmt"
	"math/bits"
)

func main() {
	var n, a int
	fmt.Scan(&n, &a)

	initMod()
	initFmt()

	wei := make(Poly, n+1)
	for i := 0; i < n+1; i++ {
		wei[i] = divMod(fact[n+i]*invf[i]%MOD, PowMod(2, i))
	}
	tmp := make(Poly, n+1)
	if a == 1 {
		tmp[0] = 1
	} else {
		for i := 0; i < n+1; i++ {
			tmp[i] = binom(a-2, i)
		}
	}
	wei = wei.mul(tmp, false)
	for i := 0; i < n+1; i++ {
		wei[i] = wei[i] * invf[n-i] % MOD
	}
	f := make(Poly, n+1)
	for i := 1; i <= n; i += 2 {
		f[i] = invf[i]
	}
	g := f.pow(n+1, a)
	res := make([]mint, n)
	for d := a; d <= n; d += 2 {
		res[d-a] = g[d] * wei[n-d] % MOD * fact[d] % MOD
	}
	last := mint(2).pow(n).inv() * invf[n] % MOD
	for i := 0; i < n; i++ {
		res[i] = res[i] * last % MOD
		if (i / 2 % 2) != 0 {
			res[i] = (MOD - res[i]) % MOD
		}
	}
	for _, i := range res {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

const MOD2 = MOD * 2
const L = 30

var g, ig, p2 [L]mint

func initFmt() {
	for i := 0; i < L; i++ {
		w := (MOD - mint(3).pow(((MOD-1)>>(i+2))*3)) % MOD
		g[i] = w
		ig[i] = w.inv()
		p2[i] = mint(1 << i).inv()
	}
}

func inplace_fmt(f Poly, inv bool) {
	n := len(f)
	if !inv {
		var b int = n
		b >>= 1
		if b != 0 { // input:[0,mod)
			for i := 0; i < b; i++ {
				x := f[i+b]
				f[i+b] = (f[i] - x + MOD) % MOD
				f[i] = (f[i] + x) % MOD
			}
		}
		b >>= 1
		if b > 0 { // input:[0,mod*2)
			p := mint(1)
			var k int = 0
			for i := 0; i < n; i += b * 2 {
				for j := i; j < i+b; j++ {
					x := f[j+b] * p % MOD
					f[j+b] = (f[j] - x + MOD) % MOD
					f[j] = (f[j] + x) % MOD
				}
				k++
				p = p * g[ctz(k)] % MOD
			}
		}
		for b != 0 {
			b >>= 1
			if b > 0 { // input:[0,mod*3)
				p := mint(1)
				var k int = 0
				for i := 0; i < n; i += b * 2 {
					for j := i; j < i+b; j++ {
						x := f[j+b] * p % MOD
						f[j+b] = (f[j] - x + MOD) % MOD
						f[j] = (f[j] + x) % MOD
					}
					k++
					p = p * g[ctz(k)] % MOD
				}
			}
			b >>= 1
			if b > 0 { // input:[0,mod*4)
				p := mint(1)
				var k int = 0
				for i := 0; i < n; i += b * 2 {
					for j := i; j < i+b; j++ {
						x := (f[j+b] * p) % MOD
						f[j+b] = (f[j] - x + MOD) % MOD
						f[j] = (f[j] + x) % MOD
					}
					k++
					p = p * g[ctz(k)] % MOD
				}
			}
		}
	} else {
		b := 1
		if b < n/2 { // input:[0,mod)
			p := mint(1)
			var k int = 0
			for i := 0; i < n; i += b * 2 {
				for j := i; j < i+b; j++ {
					x := (f[j] - f[j+b] + MOD) % MOD
					f[j] = (f[j] + f[j+b]) % MOD
					f[j+b] = x * p % MOD
				}
				k++
				p = p * ig[ctz(k)] % MOD
			}
			b <<= 1
		}
		for ; b < n/2; b <<= 1 {
			p := mint(1)
			var k int = 0
			for i := 0; i < n; i += b * 2 {
				for j := i; j < i+b/2; j++ { // input:[0,mod*2)
					x := (f[j] + MOD2 - f[j+b]) % MOD
					f[j] = (f[j] + f[j+b]) % MOD
					f[j+b] = x * p % MOD
				}
				for j := i + b/2; j < i+b; j++ { // input:[0,mod)
					x := (f[j] - f[j+b] + MOD) % MOD
					f[j] = (f[j] + f[j+b]) % MOD
					f[j+b] = x * p % MOD
				}
				k++
				p = p * ig[ctz(k)] % MOD
			}
		}
		if b < n { // input:[0,mod*2)
			for i := 0; i < b; i++ {
				x := f[i+b]
				f[i+b] = (f[i] + MOD2 - x) % MOD
				f[i] = (f[i] + x) % MOD
			}
		}
		z := p2[log2(n)]
		for i := 0; i < n; i++ {
			f[i] = f[i] * z % MOD
		}
	}
}

type Poly []mint

func plus(l, r Poly) Poly {
	res := make(Poly, max(len(l), len(r)))
	copy(res, l)
	for i := 0; i < len(r); i++ {
		res[i] = (res[i] + r[i]) % MOD
	}
	return res
}

func minus(l, r Poly) Poly {
	res := make(Poly, max(len(l), len(r)))
	copy(res, l)
	for i := 0; i < len(r); i++ {
		res[i] = (res[i] - r[i] + MOD) % MOD
	}
	return res
}

func (p Poly) mulInt(t mint) Poly {
	res := make(Poly, len(p))
	for i := range p {
		res[i] = p[i] * t % MOD
	}
	return res
}

func (x Poly) mul(y Poly, same bool) Poly {
	n := len(x) + len(y) - 1
	s := 1
	for s < n {
		s = s * 2
	}
	res := make(Poly, s)
	copy(res, x)
	inplace_fmt(res, false)
	if !same {
		z := make(Poly, s)
		copy(z, y)
		inplace_fmt(z, false)
		for i := 0; i < s; i++ {
			res[i] = res[i] * z[i] % MOD
		}
	} else {
		for i := 0; i < s; i++ {
			res[i] = res[i] * res[i] % MOD
		}
	}
	inplace_fmt(res, true)
	resize(&res, n)
	return res
}

func (P Poly) pow(s, p int) Poly {
	n := len(P)
	z := 0
	for z < n && P[z] == 0 {
		z++
	}
	if z*p >= s {
		return make(Poly, s)
	}
	var c mint = P[z]
	var cinv mint = c.inv()
	var d mint = c.pow(p)
	var t int = s - z*p
	x := make(Poly, t)
	for i := z; i < min(z+t, n); i++ {
		x[i-z] = P[i] * cinv % MOD
	}
	x = x.log(t)
	for i := 0; i < t; i++ {
		x[i] = x[i] * mint(p) % MOD
	}
	x = x.exp(t)
	for i := 0; i < t; i++ {
		x[i] = x[i] * d % MOD
	}
	y := make(Poly, s)
	for i := 0; i < t; i++ {
		y[z*p+i] = x[i]
	}
	return y
}

func (p Poly) log(s int) Poly {
	if s == 1 {
		return Poly{0}
	}
	return p.low(s).dif().mul(p.inv(s-1), false).low(s - 1).inte()
}

func (p Poly) low(s int) Poly {
	res := make(Poly, s)
	for i := 0; i < min(s, len(p)); i++ {
		res[i] = p[i]
	}
	return res
}

func (p Poly) dif() Poly {
	if len(p) == 1 {
		return Poly{0}
	} else {
		res := make(Poly, len(p)-1)
		for i := 0; i < len(res); i++ {
			res[i] = p[i+1] * mint(i+1) % MOD
		}
		return res
	}
}

func (p Poly) inv(s int) Poly {
	res := make(Poly, s)
	res[0] = mint(1).div(p[0])
	for n := 1; n < s; n *= 2 {
		f := p.low(2 * n)
		inplace_fmt(f, false)
		g := res.low(2 * n)
		inplace_fmt(g, false)
		for i := 0; i < 2*n; i++ {
			f[i] = f[i] * g[i] % MOD
		}
		inplace_fmt(f, true)
		for i := 0; i < n; i++ {
			f[i] = 0
		}
		inplace_fmt(f, false)
		for i := 0; i < 2*n; i++ {
			f[i] = f[i] * g[i] % MOD
		}
		inplace_fmt(f, true)
		for i := n; i < min(2*n, s); i++ {
			res[i] = (MOD - f[i]) % MOD
		}
	}
	return res
}

func (p Poly) inte() Poly {
	res := make(Poly, len(p)+1)
	for i := 0; i < len(p); i++ {
		res[i+1] = p[i] * invs[i+1] % MOD
	}
	return res
}

func (p Poly) exp(s int) Poly {
	f := Poly{1}
	g := Poly{1}
	for n := 1; ; n *= 2 {
		if n >= s {
			break
		}
		g = minus(g.mulInt(2), (g.mul(g, true).mul(f, false)).low(n))
		q := p.low(n).dif()
		q = plus(q, g.mul(minus(f.dif(), f.mul(q, false)), false).low(2*n-1))
		f = plus(f, f.mul(minus(p.low(2*n), q.inte()), false).low(2*n))
	}
	return f.low(s)
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

const MOD = 998244353
const VMAX = (1 << 21) + 10

var fact, invf, invs [VMAX]mint

func initMod() {
	fact[0] = 1
	invf[0] = 1
	for i := mint(1); i < VMAX; i++ {
		fact[i] = (fact[i-1] * i) % MOD
		invf[i] = invMod(fact[i])
	}
	for i := VMAX - 1; i >= 1; i-- {
		invs[i] = invf[i] * fact[i-1] % MOD
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

func binom(a, b int) mint {
	return fact[a+b] * invf[a] % MOD * invf[b] % MOD
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

func ctz(x int) int {
	return bits.TrailingZeros64(uint64(x))
}

func log2(n int) int {
	var k int
	for k = 0; n != 0; n >>= 1 {
		k++
	}
	return k - 1
}
