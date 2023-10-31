package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"sort"
)

const MOD = 998244353

var inv, fact, invfact []int

func mod_build(n int) {
	fact = make([]int, n+1)
	inv = make([]int, n+1)
	invfact = make([]int, n+1)
	fact[0] = 1
	inv[0] = 1
	invfact[0] = 1
	inv[1] = 1
	for i := 0; i < n; i++ {
		fact[i+1] = fact[i] * (i + 1) % MOD
		if i > 0 {
			inv[i+1] = MOD - inv[MOD%(i+1)]*(MOD/(i+1))%MOD
		}
		invfact[i+1] = invfact[i] * inv[i+1] % MOD
	}
}

func comb(n, k int) int {
	if n < 0 || k < 0 || k > n {
		return 0
	}
	return (fact[n] * invfact[n-k] % MOD) * invfact[k] % MOD
}

type NumberTheoreticTransform struct {
	rev, rts             []int
	base, max_base, root int
}

func (ntt *NumberTheoreticTransform) init() {
	ntt.base = 1
	ntt.rev = []int{0, 1}
	ntt.rts = []int{0, 1}
	tmp := MOD - 1
	ntt.max_base = 0
	for tmp%2 == 0 {
		tmp >>= 1
		ntt.max_base++
	}
	ntt.root = 2
	for ntt.mod_pow(ntt.root, (MOD-1)>>1) == 1 {
		ntt.root++
	}
	ntt.root = ntt.mod_pow(ntt.root, (MOD-1)>>ntt.max_base)
}

func (ntt *NumberTheoreticTransform) mod_pow(x, n int) int {
	ret := 1
	for n > 0 {
		if (n & 1) != 0 {
			ret = ntt.mul(ret, x)
		}
		x = ntt.mul(x, x)
		n >>= 1
	}
	return ret
}

func (ntt *NumberTheoreticTransform) inverse(x int) int {
	return ntt.mod_pow(x, MOD-2)
}

func (ntt NumberTheoreticTransform) add(x, y int) int {
	x = (x + y) % MOD
	return x
}

func (ntt NumberTheoreticTransform) mul(a, b int) int {
	return a * b % MOD
}

func (ntt *NumberTheoreticTransform) ensure_base(nbase int) {
	if nbase <= ntt.base {
		return
	}
	resize(&ntt.rev, 1<<nbase)
	resize(&ntt.rts, 1<<nbase)
	for i := 0; i < 1<<nbase; i++ {
		ntt.rev[i] = (ntt.rev[i>>1] >> 1) + ((i & 1) << (nbase - 1))
	}
	for ntt.base < nbase {
		z := ntt.mod_pow(ntt.root, 1<<(ntt.max_base-1-ntt.base))
		for i := 1 << (ntt.base - 1); i < (1 << ntt.base); i++ {
			ntt.rts[i<<1] = ntt.rts[i]
			ntt.rts[(i<<1)+1] = ntt.mul(ntt.rts[i], z)
		}
		ntt.base++
	}
}

func (ntt *NumberTheoreticTransform) ntt(a *[]int) {
	n := len(*a)
	zeros := ctz(n)
	ntt.ensure_base(zeros)
	shift := ntt.base - zeros
	for i := 0; i < n; i++ {
		if i < (ntt.rev[i] >> shift) {
			(*a)[i], (*a)[ntt.rev[i]>>shift] = (*a)[ntt.rev[i]>>shift], (*a)[i]
		}
	}
	for k := 1; k < n; k <<= 1 {
		for i := 0; i < n; i += 2 * k {
			for j := 0; j < k; j++ {
				z := ntt.mul((*a)[i+j+k], ntt.rts[j+k])
				(*a)[i+j+k] = ntt.add((*a)[i+j], MOD-z)
				(*a)[i+j] = ntt.add((*a)[i+j], z)
			}
		}
	}
}

func (ntt *NumberTheoreticTransform) multiply(a, b []int) []int {
	A := make([]int, len(a))
	copy(A, a)
	B := make([]int, len(b))
	copy(B, b)
	need := len(a) + len(b) - 1
	nbase := 1
	for (1 << nbase) < need {
		nbase++
	}
	ntt.ensure_base(nbase)
	sz := 1 << nbase
	resize(&A, sz)
	resize(&B, sz)
	ntt.ntt(&A)
	ntt.ntt(&B)
	inv_sz := ntt.inverse(sz)
	for i := 0; i < sz; i++ {
		A[i] = ntt.mul(A[i], ntt.mul(B[i], inv_sz))
	}
	reverseOrderInt(A[1:])
	ntt.ntt(&A)
	resize(&A, need)
	return A
}

var v [101010][]int
var sz [101010]int
var used [101010]bool

func szdfs(x, p int) {
	sz[x] = 1
	for _, to := range v[x] {
		if to == p || used[to] {
			continue
		}
		szdfs(to, x)
		sz[x] += sz[to]
	}
}

func centroid(x, p, total int) int {
	for _, to := range v[x] {
		if to == p || used[to] {
			continue
		}
		if 2*sz[to] > total {
			return centroid(to, x, total)
		}
	}
	return x
}

func dfs(x, p, h int, res *[]int) {
	if len(*res) == h {
		*res = append(*res, 1)
	} else {
		(*res)[h]++
	}
	if used[x] {
		return
	}
	for _, to := range v[x] {
		if to == p {
			continue
		}
		dfs(to, x, h+1, res)
	}
}

var ntt NumberTheoreticTransform
var ans int

func solve(x, p int) {
	szdfs(x, -1)
	if sz[x] == 1 {
		return
	}
	x = centroid(x, -1, sz[x])
	szdfs(x, -1)
	sort.Slice(v[x], func(s, t int) bool {
		return sz[v[x][s]] < sz[v[x][t]]
	})
	ret := []int{0}
	sub := []int{0}
	for _, to := range v[x] {
		res := make([]int, 0)
		dfs(to, x, 0, &res)
		mul := ntt.multiply(res, sub)
		if len(ret) < len(mul) {
			ret, mul = mul, ret
		}
		for i := 0; i < len(mul); i++ {
			ret[i] = (ret[i] + mul[i]) % MOD
		}
		if len(sub) < len(res) {
			sub, res = res, sub
		}
		for i := 0; i < len(res); i++ {
			sub[i] = (sub[i] + res[i]) % MOD
		}
	}
	for i := 1; i < len(ret); i++ {
		ans = (ans + ((ret[i]*inv[i+2]%MOD)*inv[i+1]%MOD)*4%MOD) % MOD
	}
	used[x] = true
	for _, to := range v[x] {
		if to == p || used[to] {
			continue
		}
		solve(to, x)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	mod_build(1234567)
	ntt.init()

	var n int
	fmt.Fscan(in, &n)
	for i := 0; i < n-1; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		x--
		y--
		v[x] = append(v[x], y)
		v[y] = append(v[y], x)
	}
	ans = n - 1
	for i := 0; i < n; i++ {
		ans = (ans + comb(len(v[i]), 2)) % MOD
	}
	solve(0, -1)
	fmt.Println(ans * fact[n-1] % MOD)
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

func ctz(x int) int {
	return bits.TrailingZeros32(uint32(x))
}

func reverseOrderInt(a []int) {
	n := len(a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
