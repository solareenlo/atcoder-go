package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

type Edge struct {
	to, id int
}

type Graph [][]Edge

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(in, &N, &M)
	A := make([]int, M)
	B := make([]int, M)
	G := make(Graph, N)
	for i := 0; i < M; i++ {
		fmt.Fscan(in, &A[i], &B[i])
		A[i]--
		B[i]--
		G[A[i]] = append(G[A[i]], Edge{B[i], i})
		G[B[i]] = append(G[B[i]], Edge{A[i], i})
	}

	var TECC TwoEdgeConnectedComponents
	TECC.init(G)
	C := TECC.comp
	N = TECC.sz
	G = make([][]Edge, N)
	for i := 0; i < M; i++ {
		if C[A[i]] == C[B[i]] {
			continue
		}
		G[C[A[i]]] = append(G[C[A[i]]], Edge{C[B[i]], 0})
		G[C[B[i]]] = append(G[C[B[i]]], Edge{C[A[i]], 0})
	}
	W := make([]int, N)
	for _, i := range C {
		W[i]++
	}

	var tree TreeCentroid
	tree.init(G, W)
	g := tree.centroid

	var TDP TreeDP
	TDP.init(G, W, g)
	ans := TDP.ans
	for _, i := range W {
		ans += i * i
	}

	n := 1
	for n < len(G[g]) {
		n *= 2
	}
	X := make([][]int, 2*n-1)
	for i := range X {
		X[i] = []int{1}
	}
	sum := 0
	for i := 0; i < len(G[g]); i++ {
		v := G[g][i].to
		resize(&X[i+n-1], TDP.sz[v]+1)
		X[i+n-1][0] = 1
		X[i+n-1][TDP.sz[v]] = 1
		sum += TDP.sz[v]
	}
	for i := n - 2; i >= 0; i-- {
		X[i] = con.Convolve(X[i*2+1], X[i*2+2])
	}
	for i := sum / 2; i >= 0; i-- {
		if X[0][i] != 0 {
			ans += i * (sum - i)
			break
		}
	}
	fmt.Println(ans)
}

func deepCopyGraph(dst *Graph, src Graph) {
	*dst = make(Graph, len(src))
	for i := range src {
		(*dst)[i] = make([]Edge, len(src[i]))
		copy((*dst)[i], src[i])
	}
}

type TwoEdgeConnectedComponents struct {
	G      Graph
	bridge []bool
	sz     int
	comp   []int
}

func (tecc *TwoEdgeConnectedComponents) init(G Graph) {
	deepCopyGraph(&tecc.G, G)
	N := len(G)
	M := 0
	for i := 0; i < N; i++ {
		M += len(G[i])
	}
	M /= 2
	tecc.bridge = make([]bool, M)
	var lowlink LowLink
	lowlink.init(G)
	for _, b := range lowlink.bridge {
		tecc.bridge[b] = true
	}
	tecc.comp = make([]int, N)
	for i := range tecc.comp {
		tecc.comp[i] = -1
	}
	tecc.sz = 0
	for i := 0; i < N; i++ {
		if tecc.comp[i] == -1 {
			tecc.dfs(i, tecc.sz)
			tecc.sz++
		}
	}
}

func (tecc *TwoEdgeConnectedComponents) dfs(v, id int) {
	tecc.comp[v] = id
	for _, e := range tecc.G[v] {
		if tecc.comp[e.to] != -1 {
			continue
		}
		if !tecc.bridge[e.id] {
			tecc.dfs(e.to, id)
		}
	}
}

type LowLink struct {
	G        Graph
	bridge   []int
	ord, low []int
}

func (ll *LowLink) init(G Graph) {
	deepCopyGraph(&ll.G, G)
	N := len(G)
	ll.ord = make([]int, N)
	ll.low = make([]int, N)
	for i := 0; i < N; i++ {
		ll.ord[i] = -1
		ll.low[i] = -1
	}
	ll.dfs(0, -1, 0)
}

func (ll *LowLink) dfs(v, p, id int) int {
	ll.ord[v] = id
	id++
	ll.low[v] = ll.ord[v]
	for _, e := range ll.G[v] {
		if e.to == p {
			continue
		}
		if ll.ord[e.to] == -1 {
			id = ll.dfs(e.to, v, id)
			ll.low[v] = min(ll.low[v], ll.low[e.to])
			if ll.ord[v] < ll.low[e.to] {
				ll.bridge = append(ll.bridge, e.id)
			}
		} else {
			ll.low[v] = min(ll.low[v], ll.low[e.to])
		}
	}
	return id
}

type TreeCentroid struct {
	G        Graph
	sz       []int
	N        int
	centroid int
}

func (tree *TreeCentroid) init(G Graph, W []int) {
	deepCopyGraph(&tree.G, G)
	tree.sz = make([]int, len(W))
	copy(tree.sz, W)
	tree.N = 0
	for _, i := range W {
		tree.N += i
	}
	tree.dfs(0, -1)
}

func (tree *TreeCentroid) dfs(v, p int) int {
	isCentroid := true
	for _, e := range tree.G[v] {
		if e.to == p {
			continue
		}
		tree.sz[v] += tree.dfs(e.to, v)
		if tree.sz[e.to] > tree.N/2 {
			isCentroid = false
		}
	}
	if tree.N-tree.sz[v] > tree.N/2 {
		isCentroid = false
	}
	if isCentroid {
		tree.centroid = v
	}
	return tree.sz[v]
}

type TreeDP struct {
	G   Graph
	sz  []int
	ans int
}

func (tree *TreeDP) init(G Graph, W []int, root int) {
	deepCopyGraph(&tree.G, G)
	tree.sz = make([]int, len(W))
	copy(tree.sz, W)
	tree.ans = 0
	tree.dfs(root, -1)
}

func (tree *TreeDP) dfs(v, p int) int {
	w := tree.sz[v]
	for _, e := range tree.G[v] {
		if e.to == p {
			continue
		}
		tree.ans += w * tree.dfs(e.to, v)
		tree.sz[v] += tree.sz[e.to]
	}
	return tree.sz[v]
}

const MOD = 998244353

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

var con = NewConvolution(MOD, ROOT)

// 特殊な剰余と原始根
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
