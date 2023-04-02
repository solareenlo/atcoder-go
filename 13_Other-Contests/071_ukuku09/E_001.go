package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math/bits"
	"os"
)

var F, I [2 << 17]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	var C [2 << 17]int
	for i := 0; i < n; i++ {
		var p int
		fmt.Fscan(in, &p)
		p--
		C[p] = 1
	}
	F[0] = 1
	for i := 0; i < 2*n; i++ {
		F[i+1] = F[i] * (i + 1) % mod
	}
	I[2*n] = invMod(F[2*n])
	for i := 2*n - 1; i >= 0; i-- {
		I[i] = I[i+1] * (i + 1) % mod
	}
	Q := &Heap{}
	C[2*n] = C[2*n-1]
	c := 0
	for i := 0; i < 2*n; i++ {
		c++
		if C[i] == C[i+1] {
			V := make([]int, c/2+1)
			for k := 0; k*2 <= c; k++ {
				V[k] = binom(c-k, k)
			}
			heap.Push(Q, V)
			c = 0
		}
	}
	for Q.Len() >= 2 {
		A := heap.Pop(Q).([]int)
		B := heap.Pop(Q).([]int)
		heap.Push(Q, con.Convolve(A, B))
	}
	ans := 0
	V := heap.Pop(Q).([]int)
	resize(&V, n+1)
	for i := 0; i <= n; i++ {
		if i%2 == 0 {
			ans = (ans + (F[n-i] * V[i] % mod)) % mod
		} else {
			ans = (ans - (F[n-i] * V[i] % mod) + mod) % mod
		}
	}
	fmt.Println(ans)
}

func binom(n, k int) int { return (F[n] * I[k] % mod) * I[n-k] % mod }

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

const mod = 924844033

func PowMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		n /= 2
	}
	return res
}

func invMod(a int) int {
	return PowMod(a, mod-2)
}

type Heap [][]int

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return len(h[i]) < len(h[j]) }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.([]int)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
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

const root = 5

var con = NewConvolution(mod, root)

// 特殊な剰余と原始根
// (924844033, 5)
// (998244353, 3)
// (1012924417, 5)
// (167772161, 3)
// (469762049, 3)
// (1224736769, 3)

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
