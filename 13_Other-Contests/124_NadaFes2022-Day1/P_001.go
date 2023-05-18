package main

import (
	"fmt"
	"math/bits"
)

const FFT_MAX = 23

var FFT_ROOTS [FFT_MAX + 1]int = [FFT_MAX + 1]int{1, 998244352, 911660635, 372528824, 929031873, 452798380, 922799308, 781712469, 476477967, 166035806, 258648936, 584193783, 63912897, 350007156, 666702199, 968855178, 629671588, 24514907, 996173970, 363395222, 565042129, 733596141, 267099868, 15311432}
var FFT_RATIOS [FFT_MAX]int = [FFT_MAX]int{911660635, 509520358, 369330050, 332049552, 983190778, 123842337, 238493703, 975955924, 603855026, 856644456, 131300601, 842657263, 730768835, 942482514, 806263778, 151565301, 510815449, 503497456, 743006876, 741047443, 56250497, 867605899}
var INV_FFT_RATIOS [FFT_MAX]int = [FFT_MAX]int{86583718, 372528824, 373294451, 645684063, 112220581, 692852209, 155456985, 797128860, 90816748, 860285882, 927414960, 354738543, 109331171, 293255632, 535113200, 308540755, 121186627, 608385704, 438932459, 359477183, 824071951, 103369235}
var N, M int

func main() {
	fmt.Scan(&N, &M)
	nums := make([]int, 2*M)

	InitMod()

	for k := 1; k < M; k++ {
		m := M - (N-1)*k
		if m > 0 {
			nums[M-k] = nCrMOD(m+N-1, N)
		}
	}

	polyWork0 = make([]int, SIZE)
	polyWork1 = make([]int, SIZE)
	res := solve(0, N+1)
	ks := make([]int, M)
	for k := 0; k < M; k++ {
		ks[k] = k
	}
	var tree SubproductTree
	tree.init(ks)
	ys := tree.multiEval(res.y)
	for k := 0; k < M; k++ {
		nums[M+k] = ys[k]
	}

	ans := 0
	for k := -M + 1; k < M; k++ {
		ans = (ans + (((nums[M+k]-nums[M+k-1]+MOD)%MOD)*k)%MOD) % MOD
	}
	fmt.Println(ans)
}

func solve(l, r int) pairPoly {
	if l+1 == r {
		i := l
		return pairPoly{Poly{(M + (N - i)) % MOD, MOD - 1}, Poly{nCrMOD(N, i) * INVF[N-i] % MOD}}
	} else {
		mid := (l + r) / 2
		resL := solve(l, mid)
		resR := solve(mid, r)
		gs := make(Poly, (mid-l)+len(resR.y))
		for i := 0; i < len(resR.y); i++ {
			gs[(mid-l)+i] = resR.y[i]
		}
		return pairPoly{mul(resL.x, resR.x), plus(mul(resL.y, resR.x), gs)}
	}
}

var polyWork0, polyWork1 []int

type Poly []int

type pairPoly struct {
	x, y Poly
}

func plus(l, r Poly) Poly {
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

func mul(l, r Poly) Poly {
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

func (p Poly) inv(n int) Poly {
	res := make(Poly, n)
	res[0] = INVMOD(p[0])
	for m := 1; m < n; m <<= 1 {
		for i := 0; i < min(m<<1, len(p)); i++ {
			polyWork0[i] = p[i]
		}
		for i := 0; i < (m<<1)-min(m<<1, len(p)); i++ {
			polyWork0[i+min(m<<1, len(p))] = 0
		}
		fft(polyWork0, m<<1)
		for i := 0; i < min(m<<1, n); i++ {
			polyWork1[i] = res[i]
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
			res[i] = (MOD - polyWork0[i]) % MOD
		}
	}
	return res
}

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
			prod = prod * FFT_RATIOS[ctz(h)] % MOD
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
				prod = prod * FFT_RATIOS[ctz(h)] % MOD
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
				prod = prod * FFT_RATIOS[ctz(h)] % MOD
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
			prod = prod * INV_FFT_RATIOS[ctz(h)] % MOD
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
			prod = prod * INV_FFT_RATIOS[ctz(h)] % MOD
		}
	}
	if m < n {
		for i := 0; i < m; i++ {
			y := as[i] + MOD2 - as[i+m]
			as[i] += as[i+m]
			as[i+m] = y
		}
	}
	invN := INVMOD(n)
	for i := 0; i < n; i++ {
		as[i] = as[i] * invN % MOD
	}
}

type SubproductTree struct {
	logN, n, nn int
	xs, buf     []int
	gss         [][]int
	all         Poly
}

func (tree *SubproductTree) init(xs_ []int) {
	tree.n = len(xs_)
	tree.logN = 0
	tree.nn = 1
	for tree.nn < tree.n {
		tree.logN++
		tree.nn <<= 1
	}
	tree.xs = make([]int, tree.nn)
	for i := 0; i < len(xs_); i++ {
		tree.xs[i] = xs_[i]
	}
	tree.buf = make([]int, (tree.logN+1)*(tree.nn<<1))
	tree.gss = make([][]int, tree.nn<<1)
	for h := 0; h <= tree.logN; h++ {
		for u := 1 << h; u < 1<<(h+1); u++ {
			tree.gss[u] = tree.buf[(h*(tree.nn<<1) + ((u - (1 << h)) << (tree.logN - h + 1))):]
		}
	}
	for i := 0; i < tree.nn; i++ {
		tree.gss[tree.nn+i][0] = (MOD-tree.xs[i])%MOD + 1
		tree.gss[tree.nn+i][1] = (MOD-tree.xs[i])%MOD - 1
	}
	if tree.nn == 1 {
		tree.gss[1][1] += 2
	}
	for h := tree.logN - 1; h >= 0; h-- {
		m := 1 << (tree.logN - h)
		for u := 1<<(h+1) - 1; u >= 1<<h; u-- {
			for i := 0; i < m; i++ {
				tree.gss[u][i] = tree.gss[u<<1][i] * tree.gss[u<<1|1][i] % MOD
			}
			for i := 0; i < m; i++ {
				tree.gss[u][i+m] = tree.gss[u][i]
			}
			invFft(tree.gss[u][m:], m)
			if h > 0 {
				tree.gss[u][m] -= 2
				a := FFT_ROOTS[tree.logN-h+1]
				aa := 1
				for i := m; i < m<<1; i++ {
					tree.gss[u][i] = tree.gss[u][i] * aa % MOD
					aa = aa * a % MOD
				}
				fft(tree.gss[u][m:], m)
			}
		}
	}
	resizePoly(&tree.all, tree.nn+1)
	tree.all[0] = 1
	for i := 1; i < tree.nn; i++ {
		tree.all[i] = tree.gss[1][tree.nn+tree.nn-i]
	}
	tree.all[tree.nn] = tree.gss[1][tree.nn] - 1
}

func (tree *SubproductTree) multiEval(fs Poly) []int {
	work0 := make([]int, tree.nn)
	work1 := make([]int, tree.nn)
	work2 := make([]int, tree.nn)
	m := max(len(fs), 1)
	invAll := tree.all.inv(m)
	invAll = reverseOrderPoly(invAll)
	var mm int
	for mm = 1; mm < m-1+tree.nn; mm <<= 1 {
	}
	resizePoly(&invAll, mm)
	fft(invAll, len(invAll))
	ffs := make([]int, mm)
	copy(ffs, fs)
	fft(ffs, len(ffs))
	for i := 0; i < mm; i++ {
		ffs[i] = ffs[i] * invAll[i] % MOD
	}
	invFft(ffs, len(ffs))
	if (tree.logN & 1) != 0 {
		for i := 0; i < tree.nn; i++ {
			work1[i] = ffs[i+m-1]
		}
	} else {
		for i := 0; i < tree.nn; i++ {
			work0[i] = ffs[i+m-1]
		}
	}
	for h := 0; h < tree.logN; h++ {
		m := 1 << (tree.logN - h)
		for u := 1 << h; u < 1<<(h+1); u++ {
			var hs, hs0 []int
			if ((tree.logN - h) & 1) != 0 {
				hs = work1[(u-(1<<h))<<(tree.logN-h):]
				hs0 = work0[(u-(1<<h))<<(tree.logN-h):]
			} else {
				hs = work0[(u-(1<<h))<<(tree.logN-h):]
				hs0 = work1[(u-(1<<h))<<(tree.logN-h):]
			}
			hs1 := hs0[m>>1:]
			fft(hs, m)
			for i := 0; i < m; i++ {
				work2[i] = tree.gss[(u<<1)|1][i] * hs[i] % MOD
			}
			invFft(work2, m)
			for i := 0; i < m>>1; i++ {
				hs0[i] = work2[i+(m>>1)]
			}
			for i := 0; i < m; i++ {
				work2[i] = tree.gss[u<<1][i] * hs[i] % MOD
			}
			invFft(work2, m)
			for i := 0; i < m>>1; i++ {
				hs1[i] = work2[i+(m>>1)]
			}
		}
	}
	resize(&work0, tree.n)
	return work0
}

const MOD = 998244353
const MOD2 = 998244353 * 2
const SIZE = 1 << 20

var FACT, INVF [SIZE]int

func InitMod() {
	FACT[0] = 1
	INVF[0] = 1
	for i := int(1); i < SIZE; i++ {
		FACT[i] = (FACT[i-1] * i) % MOD
		INVF[i] = INVMOD(FACT[i])
	}
}

func POWMOD(a, n int) int {
	res := int(1)
	for n > 0 {
		if n%2 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		n /= 2
	}
	return res
}

func INVMOD(a int) int {
	return POWMOD(a, MOD-2)
}

func nCrMOD(n, r int) int {
	if n < r || n < 0 || r < 0 {
		return 0
	}
	return FACT[n] * INVF[r] % MOD * INVF[n-r] % MOD
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

func ctz(x int) int {
	return bits.TrailingZeros32(uint32(x))
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

func reverseOrderPoly(a Poly) Poly {
	n := len(a)
	res := make(Poly, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
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
