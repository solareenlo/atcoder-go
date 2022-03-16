package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math/bits"
	"os"
)

const SIZE = 250005
const MOD = 998244353

var (
	inv  = [SIZE]int{}
	fac  = [SIZE]int{}
	finv = [SIZE]int{}
)

func MAKE() {
	fac[0] = 1
	fac[1] = 1
	finv[0] = 1
	finv[1] = 1
	inv[1] = 1
	for i := 2; i < SIZE; i++ {
		inv[i] = MOD - inv[MOD%i]*(MOD/i)%MOD
		fac[i] = fac[i-1] * i % MOD
		finv[i] = finv[i-1] * inv[i] % MOD
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

func invMod(a int) int {
	return POWMOD(a, MOD-2)
}

func hop(L int) []int {
	if L == 0 {
		ret := make([]int, 1)
		ret[0] = 1
		return ret
	}
	vh := hop(L / 2)
	ret := con.Convolve(vh, vh)
	if L%2 == 1 {
		coef := make([]int, 4)
		coef[0] = 0
		coef[1] = 1
		coef[2] = 1
		coef[3] = 1
		ret = con.Convolve(ret, coef)
	}
	return ret
}

func ctoi(c byte) int {
	if c == 'k' {
		return 0
	}
	if c == 'y' {
		return 2
	}
	if c == 'n' {
		return 4
	}
	return 5
}

func prod(poly [][]int) []int {
	que := &Heap{}
	for i := 0; i < len(poly); i++ {
		heap.Push(que, pair{len(poly[i]) - 1, i})
	}
	for que.Len() > 1 {
		p := (*que)[0]
		heap.Pop(que)
		q := (*que)[0]
		heap.Pop(que)
		s := p.y
		t := q.y
		poly[s] = con.Convolve(poly[s], poly[t])
		heap.Push(que, pair{len(poly[s]) - 1, s})
	}
	r := (*que)[0].y
	return poly[r]
}

func main() {
	in := bufio.NewReader(os.Stdin)

	MAKE()
	var N int
	var S string
	fmt.Fscan(in, &N, &S)
	n := len(S)
	rep := "keyence"
	coord := 0
	poly := make([][]int, 0)
	for i := 0; i < n; i++ {
		f := i
		for ; i < n && S[i] == 'e'; i++ {
		}
		to := 5
		if len(S) > i {
			to = ctoi(S[i])
		}
		if i == n {
			to = 0
		}
		L := i - f
		vc := hop(L)
		cnt := 0
		loop := 0
		if i == n-1 {
			loop = 1
		}
		a := vc[0]
		b := 0
		c := 0
		mul := make([]int, 0)
		z := 7
		for j := coord; ; j, z = j+1, z-1 {
			if L == 0 && z == 0 {
				break
			}
			if L > 0 && !(cnt < len(vc) || a > 0 || b > 0 || c > 0) {
				break
			}
			if j == 7 {
				j = 0
				loop++
			}
			if rep[j] == 'e' {
				cnt++
				if L > 0 {
					c = b
					b = a
					if cnt < len(vc) {
						a = vc[cnt]
					} else {
						a = 0
					}
				}
			} else if j == to {
				for len(mul) <= loop {
					mul = append(mul, 0)
				}
				mul[loop] = (a + b + c) % MOD
			}
		}
		coord = (to + 1) % 7
		poly = append(poly, mul)
	}
	coef := prod(poly)
	N %= MOD
	ans := 0
	vc := finv[n]
	for i := 1; i <= n; i++ {
		vc = vc * (N + i) % MOD
	}
	for i := 1; i < len(coef); i++ {
		f := N + 1 - i
		s := N + n - i + 1
		if s != 0 {
			vc = vc * f % MOD * invMod(s) % MOD
		} else {
			vc = finv[n]
			for j := f; j < s; j++ {
				vc = vc * j % MOD
			}
		}
		ans += vc * coef[i]
		ans %= MOD
	}
	fmt.Println(ans)
}

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

type Heap []pair

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i].x < h[j].x }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(pair)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type pair struct{ x, y int }
