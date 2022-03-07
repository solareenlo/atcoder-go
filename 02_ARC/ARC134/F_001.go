package main

import "fmt"

const g = 3
const invg = 332748118

var (
	maxn int
	rev  = make([]int, N)
)

func NTT(a []int, flag bool) {
	for i := 0; i < maxn; i++ {
		if rev[i] < i {
			a[i], a[rev[i]] = a[rev[i]], a[i]
		}
	}
	for i := 1; i < maxn; i <<= 1 {
		var wm int
		if flag {
			wm = powMod(g, (mod-1)/(i<<1))
		} else {
			wm = powMod(invg, (mod-1)/(i<<1))
		}
		for j := 0; j < maxn; j += (i << 1) {
			w := 1
			for k := j; k < j+i; k++ {
				x := a[k]
				y := w * a[k+i] % mod
				a[k] = (x + y) % mod
				a[k+i] = (x - y + mod) % mod
				w = w * wm % mod
			}
		}
	}
	if !flag {
		inv := powMod(maxn, mod-2)
		for i := 0; i < maxn; i++ {
			a[i] = a[i] * inv % mod
		}
	}
}

func INV(a, b []int, n int) {
	if n == 1 {
		b[0] = powMod(a[0], mod-2)
		return
	}
	c := make([]int, N)
	INV(a, b, (n+1)/2)
	for i := 0; i < n; i++ {
		c[i] = a[i]
	}
	l := 0
	maxn = 1
	for maxn < (n << 1) {
		maxn <<= 1
		l++
	}
	for i := 0; i < maxn; i++ {
		rev[i] = ((rev[i>>1] >> 1) | ((i & 1) << (l - 1)))
	}
	NTT(b, true)
	NTT(c, true)
	for i := 0; i < maxn; i++ {
		b[i] = b[i] * (2 - b[i]*c[i]%mod + mod) % mod
	}
	NTT(b, false)
	for i := n; i < maxn; i++ {
		b[i] = 0
	}
	for i := 0; i < maxn; i++ {
		c[i] = 0
	}
}

func main() {
	initMod()

	var n, W int
	fmt.Scan(&n, &W)

	ans := 0
	F := make([]int, N)
	for i, p := 2, 1; i <= n; i++ {
		F[i] = mod - inv[i]*(i-1)%mod*p%mod
		p = p * W % mod
	}
	F[0] = 1
	G := make([]int, N)
	INV(F, G, n+1)
	for i, p := 0, 1; i <= n; i++ {
		F[i] = p * inv[i] % mod
		p = p * W % mod
	}
	for i := 0; i <= n; i++ {
		ans = (ans + G[i]*F[n-i]) % mod
	}
	fmt.Println(ans * fac[n] % mod)
}

const mod = 998244353
const N = 600600

var fac, inv [N]int

func initMod() {
	fac[0] = 1
	inv[0] = 1
	for i := int(1); i < N; i++ {
		fac[i] = (fac[i-1] * i) % mod
		inv[i] = invMod(fac[i])
	}
}

func powMod(a, n int) int {
	res := int(1)
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
	return powMod(a, mod-2)
}
