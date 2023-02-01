package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

type pair struct {
	P, Q poly
}

type poly struct {
	poly []int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	par := make([]int, n+1)
	sons := make([][]int, n+1)
	for i := 2; i <= n; i++ {
		fmt.Fscan(in, &par[i])
		sons[par[i]] = append(sons[par[i]], i)
	}

	siz := make([]int, n+1)
	h := make([]int, n+1)
	for i := n; i >= 1; i-- {
		siz[i] = 1
		for _, j := range sons[i] {
			siz[i] += siz[j]
		}
		for _, j := range sons[i] {
			if siz[j] > siz[h[i]] {
				h[i] = j
			}
		}
	}

	var polys []poly
	var solve func(int, int) pair
	solve = func(L, R int) pair {
		if L == R {
			var tmp poly
			tmp.poly = make([]int, 0)
			tmp.poly = append(tmp.poly, 1)
			return pair{polys[L], tmp}
		}
		M := (L + R) / 2
		tmp1 := solve(L, M)
		PL, QL := tmp1.P, tmp1.Q
		tmp2 := solve(M+1, R)
		PR, QR := tmp2.P, tmp2.Q
		return pair{mul(PL, PR), add(QL, mul(PL, QR))}
	}

	f := make([]poly, n+1)
	for i := n; i >= 1; i-- {
		if h[par[i]] != i {
			polys = make([]poly, 0)
			for j := i; j > 0; j = h[j] {
				qolys := make([]poly, 0)
				for _, k := range sons[j] {
					if k != h[j] {
						qolys = append(qolys, f[k])
					}
				}
				if len(qolys) == 0 {
					var tmp poly
					tmp.poly = make([]int, 0)
					tmp.poly = append(tmp.poly, 1)
					qolys = append(qolys, tmp)
				}
				for len(qolys) >= 2 {
					P := qolys[0]
					qolys = qolys[1:]
					Q := qolys[0]
					qolys = qolys[1:]
					qolys = append(qolys, mul(P, Q))
				}
				polys = append(polys, qolys[0])
			}
			tmp := solve(0, len(polys)-1)
			P, Q := tmp.P, tmp.Q
			f[i] = add(P, mov(Q, 1))
		}
	}

	for k := 1; k <= n; k++ {
		if k < len(f[1].poly) {
			fmt.Fprintln(out, f[1].poly[k])
		} else {
			fmt.Fprintln(out, 0)
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func mov(F poly, mov int) poly {
	n := len(F.poly)
	var G poly
	G.poly = make([]int, max(0, n+mov))
	for i := 0; i < n; i++ {
		G.poly[i+mov] = F.poly[i]
	}
	return G
}

func add(a, b poly) poly {
	n := len(a.poly)
	m := len(b.poly)
	var c poly
	c.poly = make([]int, max(n, m))
	for i := 0; i < n; i++ {
		c.poly[i] += a.poly[i]
		c.poly[i] %= mod
	}
	for i := 0; i < m; i++ {
		c.poly[i] += b.poly[i]
		c.poly[i] %= mod
	}
	return c
}

func mul(a, b poly) poly {
	aa, bb := a.poly, b.poly
	tmp := con.Convolve(aa, bb)
	var res poly
	res.poly = make([]int, len(tmp))
	for i := range tmp {
		res.poly[i] = tmp[i]
	}
	return res
}

const mod = 998244353

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
