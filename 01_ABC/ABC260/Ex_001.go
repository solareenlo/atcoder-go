package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const maxn = 2.5e5 + 50
const maxm = 5.3e5
const mod = 998244353
const g = 3
const invg = (mod + 1) / 3

var fac, ifac, s [maxn]int
var rev [maxm]int
var c, F, G, H, T1, T2, T3 []int
var N int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)
	m++
	Init(n)
	c = make([]int, maxn)
	for i := 1; i <= n; i++ {
		var a int
		fmt.Fscan(in, &a)
		c[a]++
	}
	tmp := c[1 : n+1]
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] > tmp[j]
	})

	t := n
	for c[t] == 0 {
		t--
	}
	for i := 1; i <= t; i++ {
		s[i] = s[i-1] + c[i]
	}

	F = make([]int, maxn)
	G = make([]int, maxm)
	H = make([]int, maxm)
	T1 = make([]int, maxm)
	T2 = make([]int, maxm)
	T3 = make([]int, maxm)
	calc1(1, t)
	for i := 0; i < n; i++ {
		F[i] = (F[i] * fac[n-i] % mod) * fac[i] % mod
	}
	for i := 0; i < n; i++ {
		if (i & 1) != 0 {
			G[n-1-i] = (-ifac[i] + mod) % mod
		} else {
			G[n-1-i] = ifac[i]
		}
	}
	Mul(F, G, G, n, n)
	for i := 0; i < n; i++ {
		F[n-1-i] = G[n-1+i] * ifac[i] % mod
	}
	calc2(0, n-1)
	for i := n + 1; i < m; i++ {
		G[i] = 0
	}
	get_Inv(G, H, m)
	Mul(F, H, H, m, m)
	for i := 1; i < m; i++ {
		fmt.Fprintf(out, "%d ", H[i])
	}
}

func Init(n int) {
	fac[0] = 1
	for i := 1; i <= n; i++ {
		fac[i] = fac[i-1] * i % mod
	}
	ifac[n] = qpow(fac[n], mod-2)
	for i := n - 1; i >= 0; i-- {
		ifac[i] = ifac[i+1] * (i + 1) % mod
	}
}

func qpow(x, k int) int {
	d := 1
	for k > 0 {
		if (k & 1) != 0 {
			d = d * x % mod
		}
		x = x * x % mod
		k >>= 1
	}
	return d
}

func calc1(l, r int) {
	if l == r {
		for i := 0; i < c[l]; i++ {
			F[s[l-1]+i] = (fac[c[l]] * ifac[c[l]-i] % mod) * C(c[l]-1, i) % mod
		}
		return
	}
	mid := (l + r) >> 1
	calc1(l, mid)
	calc1(mid+1, r)
	tmp1 := F[s[l-1]:]
	tmp2 := F[s[mid]:]
	tmp3 := F[s[l-1]:]
	Mul(tmp1, tmp2, tmp3, s[mid]-s[l-1], s[r]-s[mid])
	F[s[r]-1] = 0
}

func C(n, m int) int {
	return (fac[n] * ifac[m] % mod) * ifac[n-m] % mod
}

func Mul(F, G, H []int, n, m int) {
	NTT_init(n + m)
	for i := 0; i < n; i++ {
		T1[i] = F[i]
	}
	for i := n; i < N; i++ {
		T1[i] = 0
	}
	for i := 0; i < m; i++ {
		T2[i] = G[i]
	}
	for i := m; i < N; i++ {
		T2[i] = 0
	}
	NTT(T1, N, 1)
	NTT(T2, N, 1)
	for i := 0; i < N; i++ {
		T1[i] = T1[i] * T2[i] % mod
	}
	NTT(T1, N, -1)
	for i := 0; i < n+m-1; i++ {
		H[i] = T1[i]
	}
}

func NTT_init(n int) {
	N = 1
	c := 0
	for N < n {
		N <<= 1
		c++
	}
	for i := 0; i < N; i++ {
		rev[i] = (rev[i>>1] >> 1) | ((i & 1) << (c - 1))
	}
}

func NTT(f []int, n, op int) {
	for i := 0; i < n; i++ {
		if i < rev[i] {
			f[i], f[rev[i]] = f[rev[i]], f[i]
		}
	}
	for i := 1; i < n; i <<= 1 {
		var tmp int
		if op == 1 {
			tmp = g
		} else {
			tmp = invg
		}
		w1 := qpow(tmp, (mod-1)/(i<<1))
		for j := 0; j < n; j += (i << 1) {
			w := 1
			for k := j; k < j+i; k++ {
				t1 := f[k]
				t2 := w * f[k+i] % mod
				f[k] = (t1 + t2) % mod
				f[k+i] = (t1 - t2 + mod) % mod
				w = w * w1 % mod
			}
		}
	}
	if op == -1 {
		inv := qpow(n, mod-2)
		for i := 0; i < n; i++ {
			f[i] = f[i] * inv % mod
		}
	}
}

func calc2(l, r int) {
	if l == r {
		G[l<<1] = 1
		G[l<<1|1] = (-l + mod) % mod
		return
	}
	mid := (l + r) >> 1
	calc2(l, mid)
	calc2(mid+1, r)
	tmp1 := F[mid+1:]
	tmp2 := G[l<<1:]
	Mul(tmp1, tmp2, T3, r-mid, mid-l+2)
	tmp3 := F[l:]
	tmp4 := G[(mid+1)<<1:]
	tmp5 := F[l:]
	Mul(tmp3, tmp4, tmp5, mid-l+1, r-mid+1)
	for i := l; i <= r; i++ {
		F[i] = (F[i] + T3[i-l]) % mod
	}
	tmp6 := G[l<<1:]
	tmp7 := G[(mid+1)<<1:]
	tmp8 := G[l<<1:]
	Mul(tmp6, tmp7, tmp8, mid-l+2, r-mid+1)
}

func get_Inv(F, G []int, n int) {
	if n == 1 {
		G[0] = qpow(F[0], mod-2)
		return
	}
	get_Inv(F, G, (n+1)>>1)
	NTT_init(n << 1)
	for i := 0; i < n; i++ {
		T1[i] = F[i]
	}
	for i := n; i < N; i++ {
		T1[i] = 0
	}
	NTT(T1, N, 1)
	NTT(G, N, 1)
	for i := 0; i < N; i++ {
		G[i] = G[i] * ((2 - G[i]*T1[i]%mod) + mod) % mod
	}
	NTT(G, N, -1)
	for i := n; i < N; i++ {
		G[i] = 0
	}
}
