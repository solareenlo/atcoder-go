package main

import "fmt"

const MOD = 998244353
const SIZE = 1 << 21
const inv2 = (MOD + 1) / 2
const _G = 3

var fac, ifac, inv [SIZE]int

func Init() {
	fac[0] = 1
	ifac[0] = 1
	inv[1] = 1
	for i := 2; i < SIZE; i++ {
		inv[i] = inv[MOD%i] * (MOD - MOD/i) % MOD
	}
	for i := 1; i < SIZE; i++ {
		fac[i] = fac[i-1] * i % MOD
		ifac[i] = ifac[i-1] * inv[i] % MOD
	}
}

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		n /= 2
	}
	return res
}

func nCrMod(n, r int) int {
	if n < r {
		return 0
	}
	if n == r {
		return 1
	}
	if n < 0 || r < 0 {
		return 0
	}
	return (fac[n] * ifac[r] % MOD) * ifac[n-r] % MOD
}

var rt [SIZE]int
var Lim int

func Pinit(x int) {
	for Lim = 1; Lim <= x; Lim <<= 1 {

	}
	for i := 1; i < Lim; i <<= 1 {
		sG := powMod(_G, (MOD-1)/(i<<1))
		rt[i] = 1
		for j := i + 1; j < i*2; j++ {
			rt[j] = rt[j-1] * sG % MOD
		}
	}
}

type poly struct {
	a []int
}

func (p poly) size() int {
	return len(p.a)
}

func (p poly) val(x int) int {
	if x < 0 || p.size() <= x {
		return 0
	}
	return p.a[x]
}

func (p *poly) clear() {
	p.a = make([]int, 0)
}

func (p *poly) initInt(n int) {
	p.a = make([]int, n)
}

func (p *poly) initVec(o []int) {
	p.a = make([]int, len(o))
	copy(p.a, o)
}

func (p *poly) initPoly(o poly) {
	p.a = make([]int, o.size())
	copy(p.a, o.a)
}

func (p *poly) resize(x int) poly {
	var res poly
	res.initPoly(*p)
	resize(&res.a, x)
	return res
}

func (p *poly) dif() {
	n := p.size()
	for l := n >> 1; l >= 1; l >>= 1 {
		for j := 0; j < n; j += l << 1 {
			for k, w := 0, 0; k < l; k, w = k+1, w+1 {
				x := p.a[j+k]
				y := p.a[j+k+l]
				p.a[j+k] = (x + y) % MOD
				p.a[j+k+l] = rt[l+w] * ((x - y + MOD) % MOD) % MOD
			}
		}
	}
}

func (p *poly) dit() {
	n := p.size()
	for i := 2; i <= n; i <<= 1 {
		for j, l := 0, (i >> 1); j < n; j += i {
			for k, w := 0, 0; k < l; k, w = k+1, w+1 {
				pa := p.a[j+k]
				pb := p.a[j+k+l] * rt[l+w] % MOD
				p.a[j+k] = (pa + pb) % MOD
				p.a[j+k+l] = (pa - pb + MOD) % MOD
			}
		}
	}
	reverseOrderInt(p.a[1:])
	for i, iv := 0, powMod(n, MOD-2); i < n; i++ {
		p.a[i] = p.a[i] * iv % MOD
	}
}

func mulPoly(aa, bb poly) poly {
	if aa.size() == 0 || bb.size() == 0 {
		var res poly
		return res
	}
	var A, B poly
	A.initPoly(aa)
	B.initPoly(bb)
	var lim int
	all := A.size() + B.size() - 1
	for lim = 1; lim < all; lim <<= 1 {

	}
	A = A.resize(lim)
	B = B.resize(lim)
	A.dif()
	B.dif()
	for i := 0; i < lim; i++ {
		A.a[i] = A.a[i] * B.a[i] % MOD
	}
	A.dit()
	A = A.resize(all)
	return A
}

func (p *poly) Inv() poly {
	var res poly
	res.initInt(1)
	res.a[0] = powMod(p.a[0], MOD-2)
	for m := 1; m < p.size(); m <<= 1 {
		pn := m << 1
		var f poly
		f.initPoly(res)
		f = f.resize(pn)
		var g poly
		g.initInt(pn)
		for i := 0; i < pn; i++ {
			g.a[i] = p.val(i)
		}
		f.dif()
		g.dif()
		for i := 0; i < pn; i++ {
			g.a[i] = f.a[i] * g.a[i] % MOD
		}
		g.dit()
		for i := 0; i < m; i++ {
			g.a[i] = 0
		}
		g.dif()
		for i := 0; i < pn; i++ {
			g.a[i] = f.a[i] * g.a[i] % MOD
		}
		g.dit()
		res = res.resize(pn)
		for i := m; i < min(pn, p.size()); i++ {
			res.a[i] = (MOD - g.a[i]) % MOD
		}
	}
	res = res.resize(p.size())
	return res
}

func (p *poly) Shift(x int) poly {
	var zm poly
	zm.initInt(p.size() + x)
	for i := 0; i < p.size(); i++ {
		zm.a[i+x] = p.a[i]
	}
	return zm
}

func mulInt(aa poly, bb int) poly {
	var res poly
	res.initInt(aa.size())
	for i := 0; i < aa.size(); i++ {
		res.a[i] = aa.a[i] * bb % MOD
	}
	return res
}

func plusPoly(aa, bb poly) poly {
	var res poly
	res.initInt(max(aa.size(), bb.size()))
	for i := 0; i < res.size(); i++ {
		res.a[i] = (aa.val(i) + bb.val(i)) % MOD
	}
	return res
}

func minusPoly(aa, bb poly) poly {
	var res poly
	res.initInt(max(aa.size(), bb.size()))
	for i := 0; i < res.size(); i++ {
		res.a[i] = (aa.val(i) - bb.val(i) + MOD) % MOD
	}
	return res
}

func (p *poly) Integ() poly {
	var res poly
	if p.size() == 0 {
		res.initInt(0)
		return res
	}
	res.initInt(p.size() + 1)
	for i := 1; i <= p.size(); i++ {
		res.a[i] = p.a[i-1] * inv[i] % MOD
	}
	return res
}

func (p *poly) Deriv() poly {
	var res poly
	if p.size() == 0 {
		res.initInt(0)
		return res
	}
	res.initInt(p.size() - 1)
	for i := 1; i < p.size(); i++ {
		res.a[i-1] = p.a[i] * i % MOD
	}
	return res
}

func (p *poly) Ln() poly {
	tmpInv := p.Inv()
	tmpDeriv := p.Deriv()
	tmp := mulPoly(tmpInv, tmpDeriv)
	res := tmp.Integ()
	res = res.resize(p.size())
	return res
}

func (p *poly) Exp() poly {
	var res poly
	res.initInt(1)
	res.a[0] = 1
	var f poly
	f.initInt(0)
	for m := 1; m < p.size(); m <<= 1 {
		pn := min(m<<1, p.size())
		f = f.resize(pn)
		res = res.resize(pn)
		for i := 0; i < pn; i++ {
			f.a[i] = p.val(i)
		}
		f = minusPoly(f, res.Ln())
		f.a[0] = (f.a[0] + 1) % MOD
		res = mulPoly(res, f)
		res = res.resize(pn)
	}
	res = res.resize(p.size())
	return res
}

// x : the power % MOD;
// rx : the power % (MOD - 1)
func (p *poly) pow(x, rx int) poly {
	if rx == -1 {
		rx = x
	}
	cnt := 0
	for p.a[cnt] == 0 && cnt < p.size() {
		cnt += 1
	}

	var res poly
	res.initPoly(*p)
	for i := cnt; i < p.size(); i++ {
		res.a[i-cnt] = res.a[i]
	}
	for i := p.size() - cnt; i < p.size(); i++ {
		res.a[i] = 0
	}
	c := res.a[0]
	w := powMod(res.a[0], MOD-2)
	for i := 0; i < res.size(); i++ {
		res.a[i] = res.a[i] * w % MOD
	}
	res = res.Ln()
	for i := 0; i < res.size(); i++ {
		res.a[i] = res.a[i] * x % MOD
	}
	res = res.Exp()
	c = powMod(c, rx)
	for i := 0; i < res.size(); i++ {
		res.a[i] = res.a[i] * c % MOD
	}

	if cnt*x > p.size() {
		for i := 0; i < p.size(); i++ {
			res.a[i] = 0
		}
	} else if cnt != 0 {
		for i := p.size() - cnt*x - 1; i >= 0; i-- {
			res.a[i+cnt*x] = res.a[i]
		}
		for i := 0; i < cnt*x; i++ {
			res.a[i] = 0
		}
	}
	return res
}

func (p *poly) sqrt(rt int) poly {
	var res poly
	res.initInt(1)
	res.a[0] = rt
	var f poly
	f.initInt(0)
	for m := 1; m < p.size(); m <<= 1 {
		pn := min(m<<1, p.size())
		f = f.resize(pn)
		for i := 0; i < pn; i++ {
			f.a[i] = p.val(i)
		}
		f = plusPoly(f, mulPoly(res, res))
		f = f.resize(pn)
		res = res.resize(pn)
		res = mulPoly(f, res.Inv())
		res = res.resize(pn)
		for i := 0; i < pn; i++ {
			res.a[i] = res.a[i] * inv2 % MOD
		}
	}
	return res
}

type pairPoly struct {
	x, y poly
}

/* f / g = first */
/* f % g = second */
func divPoly(f, g poly) pairPoly {
	f = f.resize(max(f.size(), g.size()))
	reverseOrderInt(f.a)
	reverseOrderInt(g.a)
	n := f.size()
	m := g.size()
	g = g.resize(n - m + 1)
	A := g.Inv()
	var t poly
	A = mulPoly(A, f.resize(n-m+1))
	A.resize(n - m + 1)
	reverseOrderInt(A.a)
	reverseOrderInt(g.a)
	reverseOrderInt(f.a)
	t = minusPoly(f, mulPoly(A, g))
	t.resize(m - 1)
	return pairPoly{A, t}
}

func main() {
	var n int
	fmt.Scan(&n)
	Pinit(n*2 + 4)
	Init()
	var f poly
	f.initInt(n + 1)
	for i := 0; i <= n; i++ {
		f.a[i] = powMod(2, i*(i-1)/2%(MOD-1)) * ifac[i] % MOD
	}
	f = f.Ln()
	ns := f.a[n] * fac[n] % MOD

	var cur poly
	cur.initInt(n + 1)
	for i := 2; i <= n; i += 2 {
		cur.a[i] = ifac[i] * n % MOD
	}
	cur = cur.Exp()
	qwq := ((cur.a[n-1] * inv[n] % MOD) * fac[n] % MOD) * inv[n] % MOD
	fmt.Println((ns + MOD - qwq) % MOD)
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

func reverseOrderInt(a []int) {
	n := len(a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
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
