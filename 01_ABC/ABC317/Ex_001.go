package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

type matrix [3][3][]int

func iadd(P *[]int, Q []int) {
	n := len(Q)
	if len(*P) < n {
		resize(P, n)
	}
	for i := 0; i < n; i++ {
		(*P)[i] = ((*P)[i] + Q[i]) % MOD
	}
}

func mul(X, Y matrix) matrix {
	var Z matrix
	for i := 0; i < 3; i++ {
		for k := 0; k < 3; k++ {
			for j := 0; j < 3; j++ {
				iadd(&Z[i][j], con.Convolve(X[i][k], Y[k][j]))
			}
		}
	}
	return Z
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var A [4][50000]int

	var N, K int
	fmt.Fscan(in, &N, &K)
	for j := 0; j < 4; j++ {
		for i := 0; i < N; i++ {
			fmt.Fscan(in, &A[j][i])
		}
	}
	V := make([]matrix, 0)
	W := make([]matrix, 0)
	for i := 1; i < N; i++ {
		var X matrix
		var P, Q, R, D []int
		if i >= 0 && A[0][i] != 0 {
			P = []int{1}
		} else {
			P = []int{1, -1}
		}
		if i-1 >= 0 && A[1][i-1] != 0 {
			Q = []int{0, 1}
		} else {
			Q = []int{0}
		}
		if i-2 >= 0 && A[2][i-2] != 0 {
			R = []int{0, 1}
		} else {
			R = []int{0}
		}
		if A[3][i] != 0 {
			D = []int{1}
		} else {
			D = []int{0}
		}
		X[0][0] = con.Convolve(P, Q)
		X[0][1] = con.Convolve(P, R)
		X[1][0] = []int{1, -1}
		X[2][0] = con.Convolve(X[0][0], D)
		X[2][1] = con.Convolve(X[0][1], D)
		X[2][2] = []int{1, -1}
		V = append(V, X)
	}
	for len(V) > 1 {
		for len(V) != 0 {
			X := V[0]
			V = V[1:]
			if len(V) != 0 {
				X = mul(V[0], X)
				V = V[1:]
			}
			W = append(W, X)
		}
		V, W = W, V
	}
	X := V[0]
	F := X[0][0]
	G := X[2][0]
	H := make([]int, max(N, len(G)+1))
	if A[3][0] != 0 {
		iadd(&G, X[2][2])
	}
	H[0] = 1
	for i := 0; i < N-1; i++ {
		H[i+1] = divMod(H[i]*(MOD-(N-1-i))%MOD, (i + 1))
	}
	for i := 0; i < len(G); i++ {
		H[i+1] = (H[i+1] - G[i] + MOD) % MOD
	}
	for ; K > 0; K >>= 1 {
		E := make([]int, len(H))
		copy(E, H)
		for i := 1; i < len(E); i += 2 {
			E[i] = (MOD - E[i]) % MOD
		}
		S := con.Convolve(F, E)
		T := con.Convolve(H, E)
		F = make([]int, 0)
		H = make([]int, 0)
		for i := K & 1; i < len(S); i += 2 {
			F = append(F, S[i])
		}
		for i := 0; i < len(T); i += 2 {
			H = append(H, T[i])
		}
	}
	fmt.Println(F[0])
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

func divMod(a, b int) int {
	ret := a * modInv(b)
	ret %= MOD
	return ret
}

func modInv(a int) int {
	b, u, v := MOD, 1, 0
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
const MOD1 = 754974721
const MOD2 = 167772161
const MOD3 = 469762049

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

const ROOT = 3
const ROOT1 = 11
const ROOT2 = 3
const ROOT3 = 3

var con = NewConvolution(MOD, ROOT)
var con1 = NewConvolution(MOD1, ROOT1)
var con2 = NewConvolution(MOD2, ROOT2)
var con3 = NewConvolution(MOD3, ROOT3)

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

func ConvolveLL(a, b []int) []int {
	n := len(a)
	m := len(b)
	if n == 0 || m == 0 {
		return []int{}
	}

	MOD1 := 754974721 // 2^24
	MOD2 := 167772161 // 2^25
	MOD3 := 469762049 // 2^26
	M2M3 := MOD2 * MOD3
	M1M3 := MOD1 * MOD3
	M1M2 := MOD1 * MOD2
	M1M2M3 := MOD1 * MOD2 * MOD3

	i1 := inv_gcd(MOD2*MOD3, MOD1).y
	i2 := inv_gcd(MOD1*MOD3, MOD2).y
	i3 := inv_gcd(MOD1*MOD2, MOD3).y

	c1 := con1.Convolve(a, b)
	c2 := con2.Convolve(a, b)
	c3 := con3.Convolve(a, b)

	c := make([]int, n+m-1)
	for i := 0; i < n+m-1; i++ {
		x := 0
		x += (c1[i] * i1) % MOD1 * M2M3
		x += (c2[i] * i2) % MOD2 * M1M3
		x += (c3[i] * i3) % MOD3 * M1M2
		diff := c1[i] - safe_mod((x), (MOD1))
		if diff < 0 {
			diff += MOD1
		}
		offset := [5]int{0, 0, M1M2M3, 2 * M1M2M3, 3 * M1M2M3}
		x -= offset[diff%5]
		c[i] = x
	}

	return c
}

type pair struct{ x, y int }

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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
