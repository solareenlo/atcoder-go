package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	initMod()

	var N, M int
	fmt.Fscan(in, &N, &M)
	fs := make(Poly, N+1)
	for i := 1; i < M; i++ {
		fs[i] = 1
	}
	gs := make(Poly, N+1)
	copy(gs, fs)
	gs[0] += 1
	gs[M] += 1
	polyWork0 = make([]int, LIM_POLY)
	polyWork1 = make([]int, LIM_POLY)
	polyWork2 = make([]int, LIM_POLY)
	polyWork3 = make([]int, LIM_POLY)
	gs = gs.div(Poly{1}.plus(fs), N+1)
	ans := 0
	ans = (ans + Poly{1}.plus(fs).log(N + 1)[N]) % MOD
	ans = (ans + gs.log(N + 1)[N]) % MOD
	hs := Poly{0, 1}
	for m := 1; m < N+1; m <<= 1 {
		mm := m << 1
		tmp0 := hs.mulPoly(hs)
		tmp1 := mulInt(M, hs.powInt(M+1, mm)).mulPoly(Poly{1, MOD - 1})
		tmp2 := mulInt((M - 1), hs.powInt(M+2, mm)).mulPoly(Poly{1, MOD - 1})
		tmp3 := mulInt(2, hs).mulPoly(Poly{0, 1})
		numer := tmp0.minus(tmp1).plus(tmp2).plus(Poly{0, 1}).minus(tmp3)
		tmp4 := mulInt((M + 1), hs.powInt(M, mm))
		tmp5 := mulInt(M, hs.powInt(M+1, mm))
		tmp6 := Poly{1}.minus(tmp4).plus(tmp5)
		denom := tmp6.mulPoly(Poly{1, MOD - 1})
		hs = numer.mod(mm).div(denom, mm)
	}

	hs = hs[1:]
	hs = hs.powMint(mint((MOD - N)), N+1)

	for i := 1; i < N; i++ {
		ans = (ans + (N*int(inv[N-i])%MOD)*((MOD-(int(inv[N])*hs[N-i]%MOD))%MOD)%MOD) % MOD
	}
	fmt.Println(ans * N % MOD)
}

type Poly []int

func (p Poly) at(k int) int {
	if 0 <= k && k < len(p) {
		return p[k]
	}
	return 0
}

func (l Poly) plus(r Poly) Poly {
	res := make(Poly, len(l))
	copy(res, l)
	if len(l) < len(r) {
		resizePoly(&res, len(r))
	}
	for i := 0; i < len(r); i++ {
		res[i] = (res[i] + r[i]) % MOD
	}
	return res
}

func (l Poly) minus(r Poly) Poly {
	res := make(Poly, max(len(l), len(r)))
	copy(res, l)
	for i := 0; i < len(r); i++ {
		res[i] = (res[i] - r[i] + MOD) % MOD
	}
	return res
}

func (l Poly) mulPoly(r Poly) Poly {
	res := make(Poly, len(l))
	copy(res, l)
	if len(l) == 0 || len(r) == 0 {
		return res
	}
	nt := len(l)
	nf := len(r)
	n := 1
	for ; n < nt+nf-1; n <<= 1 {
	}
	resizePoly(&res, n)
	fft(res, n)
	for i := 0; i < nf; i++ {
		polyWork0[i] = r[i]
	}
	for i := 0; i < n-nf; i++ {
		polyWork0[i+nf] = 0
	}
	fft(polyWork0, n)
	for i := 0; i < n; i++ {
		res[i] = res[i] * polyWork0[i] % MOD
	}
	invFft(res, n)
	resizePoly(&res, nt+nf-1)
	return res
}

func mulInt(n int, l Poly) Poly {
	res := make(Poly, len(l))
	copy(res, l)
	for i := range res {
		res[i] = res[i] * n % MOD
	}
	return res
}

const LIM_POLY = 1 << 20

var polyWork0, polyWork1, polyWork2, polyWork3 []int

func (p Poly) div(fs Poly, n int) Poly {
	if n == 1 {
		return Poly{int(mint(p.at(0)).div(mint(fs[0])))}
	}
	m := 1 << (31 - clz(uint32(n-1)))
	gs := fs.inv(m)
	resizePoly(&gs, m<<1)
	fft(gs, m<<1)
	for i := 0; i < min(m, len(p)); i++ {
		polyWork0[i] = p[i]
	}
	for i := 0; i < (m<<1)-min(m, len(p)); i++ {
		polyWork0[i+min(m, len(p))] = 0
	}
	fft(polyWork0, m<<1)
	for i := 0; i < m<<1; i++ {
		polyWork0[i] = polyWork0[i] * gs[i] % MOD
	}
	invFft(polyWork0, m<<1)
	hs := make(Poly, n)
	for i := 0; i < m; i++ {
		hs[i] = polyWork0[i]
	}
	for i := 0; i < m; i++ {
		polyWork0[i+m] = 0
	}
	fft(polyWork0, m<<1)
	for i := 0; i < min(m<<1, len(fs)); i++ {
		polyWork1[i] = fs[i]
	}
	for i := 0; i < (m<<1)-min(m<<1, len(fs)); i++ {
		polyWork1[i+min(m<<1, len(fs))] = 0
	}
	fft(polyWork1, m<<1)
	for i := 0; i < m<<1; i++ {
		polyWork0[i] = polyWork0[i] * polyWork1[i] % MOD
	}
	invFft(polyWork0, m<<1)
	for i := 0; i < m; i++ {
		polyWork0[i] = 0
	}
	for i, i0 := m, min(m<<1, len(p)); i < i0; i++ {
		polyWork0[i] = (polyWork0[i] - p[i] + MOD) % MOD
	}
	fft(polyWork0, m<<1)
	for i := 0; i < m<<1; i++ {
		polyWork0[i] = polyWork0[i] * gs[i] % MOD
	}
	invFft(polyWork0, m<<1)
	for i := m; i < n; i++ {
		hs[i] = (MOD - polyWork0[i]) % MOD
	}
	return hs
}

func (p Poly) inv(n int) Poly {
	fs := make(Poly, n)
	fs[0] = int(mint(p[0]).inv())
	for m := 1; m < n; m <<= 1 {
		for i := 0; i < min(m<<1, len(p)); i++ {
			polyWork0[i] = p[i]
		}
		for i := 0; i < ((m << 1) - min(m<<1, len(p))); i++ {
			polyWork0[i+min(m<<1, len(p))] = 0
		}
		fft(polyWork0, m<<1)
		for i := 0; i < min(m<<1, n); i++ {
			polyWork1[i] = fs[i]
		}
		for i := 0; i < (m<<1)-min(m<<1, n); i++ {
			polyWork1[i+min(m<<1, n)] = 0
		}
		fft(polyWork1, m<<1)
		for i := 0; i < m<<1; i++ {
			polyWork0[i] = polyWork0[i] * polyWork1[i] % MOD
		}
		invFft(polyWork0, m<<1)
		for i := 0; i < m; i++ {
			polyWork0[i] = 0
		}
		fft(polyWork0, m<<1)
		for i := 0; i < m<<1; i++ {
			polyWork0[i] = polyWork0[i] * polyWork1[i] % MOD
		}
		invFft(polyWork0, m<<1)
		for i, i0 := m, min(m<<1, n); i < i0; i++ {
			fs[i] = (MOD - polyWork0[i]) % MOD
		}
	}
	return fs
}

func (p Poly) ord() int {
	for i := 0; i < len(p); i++ {
		if p[i] != 0 {
			return i
		}
	}
	return -1
}

func (p Poly) mod(n int) Poly {
	size := min(n, len(p))
	res := make(Poly, size)
	copy(res, p[:size])
	return res
}

func (p Poly) log(n int) Poly {
	fs := p.mod(n)
	for i := 0; i < len(fs); i++ {
		fs[i] = fs[i] * i % MOD
	}
	fs = fs.div(p, n)
	for i := 1; i < n; i++ {
		fs[i] = fs[i] * int(inv[i]) % MOD
	}
	return fs
}

func (p Poly) powMint(a mint, n int) Poly {
	return (mulInt(int(a), p.log(n))).exp(n)
}

func (p Poly) powInt(a, n int) Poly {
	if a == 0 {
		gs := make(Poly, n)
		gs[0] = 1
		return gs
	}
	o := p.ord()
	if o == -1 || o > (n-1)/a {
		res := make(Poly, n)
		return res
	}
	b := int(mint(p[o]).inv())
	c := int(mint(p[o]).pow(a))
	ntt := min(n-a*o, len(p)-o)
	tts := make(Poly, ntt)
	for i := 0; i < ntt; i++ {
		tts[i] = b * p[o+i] % MOD
	}
	tts = tts.powMint(mint(a), n-a*o)
	gs := make(Poly, n)
	for i := 0; i < n-a*o; i++ {
		gs[a*o+i] = c * tts[i] % MOD
	}
	return gs
}

func (p Poly) exp(n int) Poly {
	if n == 1 {
		return Poly{1}
	}
	if n == 2 {
		return Poly{1, p.at(1)}
	}
	fs := make(Poly, n)
	fs[0] = 1
	polyWork1[0] = 1
	polyWork1[1] = 1
	polyWork2[0] = 1
	var m int
	for m = 1; m<<1 < n; m <<= 1 {
		for i, i0 := 0, min(m, len(p)); i < i0; i++ {
			polyWork0[i] = i * p[i] % MOD
		}
		for i := 0; i < m-min(m, len(p)); i++ {
			polyWork0[i+min(m, len(p))] = 0
		}
		fft(polyWork0, m)
		for i := 0; i < m; i++ {
			polyWork0[i] = polyWork0[i] * polyWork1[i] % MOD
		}
		invFft(polyWork0, m)
		for i := 0; i < m; i++ {
			polyWork0[i] = (polyWork0[i] - i*fs[i]%MOD + MOD) % MOD
		}
		for i := 0; i < m; i++ {
			polyWork0[i+m] = 0
		}
		fft(polyWork0, m<<1)
		for i := 0; i < m; i++ {
			polyWork3[i] = polyWork2[i]
		}
		for i := 0; i < m; i++ {
			polyWork3[i+m] = 0
		}
		fft(polyWork3, m<<1)
		for i := 0; i < m<<1; i++ {
			polyWork0[i] = polyWork0[i] * polyWork3[i] % MOD
		}
		invFft(polyWork0, m<<1)
		for i := 0; i < m; i++ {
			polyWork0[i] = polyWork0[i] * int(inv[m+i]) % MOD
		}
		for i, i0 := 0, min(m, len(p)-m); i < i0; i++ {
			polyWork0[i] = (polyWork0[i] + p[m+i]) % MOD
		}
		for i := 0; i < m; i++ {
			polyWork0[i+m] = 0
		}
		fft(polyWork0, m<<1)
		for i := 0; i < m<<1; i++ {
			polyWork0[i] = polyWork0[i] * polyWork1[i] % MOD
		}
		invFft(polyWork0, m<<1)
		for i := 0; i < m; i++ {
			fs[i+m] = polyWork0[i]
		}
		for i := 0; i < m<<1; i++ {
			polyWork1[i] = fs[i]
		}
		for i := 0; i < m<<1; i++ {
			polyWork1[i+(m<<1)] = 0
		}
		fft(polyWork1, m<<2)
		for i := 0; i < m<<1; i++ {
			polyWork0[i] = polyWork1[i] * polyWork3[i] % MOD
		}
		invFft(polyWork0, m<<1)
		for i := 0; i < m; i++ {
			polyWork0[i] = 0
		}
		fft(polyWork0, m<<1)
		for i := 0; i < m<<1; i++ {
			polyWork0[i] = polyWork0[i] * polyWork3[i] % MOD
		}
		invFft(polyWork0, m<<1)
		for i := m; i < m<<1; i++ {
			polyWork2[i] = (MOD - polyWork0[i]) % MOD
		}
	}
	for i, i0 := 0, min(m, len(p)); i < i0; i++ {
		polyWork0[i] = i * p[i] % MOD
	}
	for i := 0; i < m-min(m, len(p)); i++ {
		polyWork0[i+min(m, len(p))] = 0
	}
	fft(polyWork0, m)
	for i := 0; i < m; i++ {
		polyWork0[i] = polyWork0[i] * polyWork1[i] % MOD
	}
	invFft(polyWork0, m)
	for i := 0; i < m; i++ {
		polyWork0[i] = (polyWork0[i] - i*fs[i]%MOD + MOD) % MOD
	}
	for i := 0; i < m>>1; i++ {
		polyWork0[i+m] = polyWork0[i+(m>>1)]
	}
	for i := 0; i < m>>1; i++ {
		polyWork0[i+(m>>1)] = 0
	}
	for i := 0; i < m>>1; i++ {
		polyWork0[i+m+(m>>1)] = 0
	}
	fft(polyWork0, m)
	fft(polyWork0[m:], m)
	for i := 0; i < m>>1; i++ {
		polyWork3[i+m] = polyWork2[i+(m>>1)]
	}
	for i := 0; i < m>>1; i++ {
		polyWork3[i+m+(m>>1)] = 0
	}
	fft(polyWork3[m:], m)
	for i := 0; i < m; i++ {
		polyWork0[m+i] = ((polyWork0[i] * polyWork3[m+i] % MOD) + (polyWork0[m+i] * polyWork3[i] % MOD)) % MOD
	}
	for i := 0; i < m; i++ {
		polyWork0[i] = polyWork0[i] * polyWork3[i] % MOD
	}
	invFft(polyWork0, m)
	invFft(polyWork0[m:], m)
	for i := 0; i < m>>1; i++ {
		polyWork0[(m>>1)+i] = (polyWork0[(m>>1)+i] + polyWork0[m+i]) % MOD
	}
	for i := 0; i < m; i++ {
		polyWork0[i] = polyWork0[i] * int(inv[m+i]) % MOD
	}
	for i, i0 := 0, min(m, len(p)-m); i < i0; i++ {
		polyWork0[i] = (polyWork0[i] + p[m+i]) % MOD
	}
	for i := 0; i < m; i++ {
		polyWork0[i+m] = 0
	}
	fft(polyWork0, m<<1)
	for i := 0; i < m<<1; i++ {
		polyWork0[i] = polyWork0[i] * polyWork1[i] % MOD
	}
	invFft(polyWork0, m<<1)
	for i := 0; i < n-m; i++ {
		fs[i+m] = polyWork0[i]
	}
	return fs
}

const FFT_MAX = 23
const MOD2 = 2 * MOD

var FFT_RATIOS [FFT_MAX]int = [FFT_MAX]int{911660635, 509520358, 369330050, 332049552, 983190778, 123842337, 238493703, 975955924, 603855026, 856644456, 131300601, 842657263, 730768835, 942482514, 806263778, 151565301, 510815449, 503497456, 743006876, 741047443, 56250497, 867605899}
var INV_FFT_RATIOS [FFT_MAX]int = [FFT_MAX]int{86583718, 372528824, 373294451, 645684063, 112220581, 692852209, 155456985, 797128860, 90816748, 860285882, 927414960, 354738543, 109331171, 293255632, 535113200, 308540755, 121186627, 608385704, 438932459, 359477183, 824071951, 103369235}

func fft(as []int, n int) {
	m := n
	m = m >> 1
	if m > 0 {
		for i := 0; i < m; i++ {
			x := as[i+m]
			as[i+m] = as[i] + MOD - x
			as[i] += x
		}
	}
	m = m >> 1
	if m > 0 {
		prod := 1
		for h, i0 := 0, 0; i0 < n; i0 += (m << 1) {
			for i := i0; i < i0+m; i++ {
				x := (prod * as[i+m]) % MOD
				as[i+m] = as[i] + MOD - x
				as[i] += x
			}
			h++
			prod = prod * FFT_RATIOS[ctz(uint32(h))] % MOD
		}
	}
	for m > 0 {
		m = m >> 1
		if m > 0 {
			prod := 1
			for h, i0 := 0, 0; i0 < n; i0 += (m << 1) {
				for i := i0; i < i0+m; i++ {
					x := (prod * as[i+m]) % MOD
					as[i+m] = as[i] + MOD - x
					as[i] += x
				}
				h++
				prod = prod * FFT_RATIOS[ctz(uint32(h))] % MOD
			}
		}
		m = m >> 1
		if m > 0 {
			prod := 1
			for h, i0 := 0, 0; i0 < n; i0 += (m << 1) {
				for i := i0; i < i0+m; i++ {
					x := (prod * as[i+m]) % MOD
					if as[i] >= MOD2 {
						as[i] = as[i] - MOD2
					} else {
						as[i] = as[i]
					}
					as[i+m] = as[i] + MOD - x
					as[i] += x
				}
				h++
				prod = prod * FFT_RATIOS[ctz(uint32(h))] % MOD
			}
		}
	}
	for i := 0; i < n; i++ {
		if as[i] >= MOD2 {
			as[i] = as[i] - MOD2
		} else {
			as[i] = as[i]
		}
		if as[i] >= MOD {
			as[i] = as[i] - MOD
		} else {
			as[i] = as[i]
		}
	}
}

func invFft(as []int, n int) {
	m := 1
	if m < n>>1 {
		prod := 1
		for h, i0 := 0, 0; i0 < n; i0 += (m << 1) {
			for i := i0; i < i0+m; i++ {
				y := as[i] + MOD - as[i+m]
				as[i] += as[i+m]
				as[i+m] = (prod * y) % MOD
			}
			h++
			prod = prod * INV_FFT_RATIOS[ctz(uint32(h))] % MOD
		}
		m <<= 1
	}
	for ; m < n>>1; m <<= 1 {
		prod := 1
		for h, i0 := 0, 0; i0 < n; i0 += (m << 1) {
			for i := i0; i < i0+(m>>1); i++ {
				y := as[i] + MOD2 - as[i+m]
				as[i] += as[i+m]
				if as[i] >= MOD2 {
					as[i] = as[i] - MOD2
				} else {
					as[i] = as[i]
				}
				as[i+m] = (prod * y) % MOD
			}
			for i := i0 + (m >> 1); i < i0+m; i++ {
				y := as[i] + MOD - as[i+m]
				as[i] += as[i+m]
				as[i+m] = (prod * y) % MOD
			}
			h++
			prod = prod * INV_FFT_RATIOS[ctz(uint32(h))] % MOD
		}
	}
	if m < n {
		for i := 0; i < m; i++ {
			y := as[i] + MOD2 - as[i+m]
			as[i] += as[i+m]
			as[i+m] = y
		}
	}
	invN := int(invMod(mint(n)))
	for i := 0; i < n; i++ {
		as[i] = as[i] * invN % MOD
	}
}

func clz(x uint32) int {
	return bits.LeadingZeros32(x)
}

func ctz(x uint32) int {
	return bits.TrailingZeros32(x)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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

const MOD = 998244353
const LIM_INV = 1 << 20

var fact, invf, inv [LIM_INV]mint

func initMod() {
	inv[1] = 1
	for i := 2; i < LIM_INV; i++ {
		inv[i] = (MOD - (mint(MOD/i) * inv[MOD%i] % MOD)) % MOD
	}
	fact[0] = 1
	invf[0] = 1
	for i := mint(1); i < LIM_INV; i++ {
		fact[i] = (fact[i-1] * i) % MOD
		invf[i] = invMod(fact[i])
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
