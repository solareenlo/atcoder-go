package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = (1 << 18)

var (
	A   = make([]int, N)
	mi  = make([]int, N)
	a   = make([]int, N)
	b   = make([]int, N)
	c   = make([]int, N)
	rev = make([]int, N)
	n   int
)

func ntt(a []int, p int) {
	for i := 0; i < N; i++ {
		if i < rev[i] {
			a[i], a[rev[i]] = a[rev[i]], a[i]
		}
	}
	for i := 2; i <= N; i <<= 1 {
		s := powMod(3, (mod-1)/i)
		if p != 0 {
			s = powMod(s, mod-2)
		}
		for j := 0; j < N; j += i {
			for k, ss := 0, 1; k < (i >> 1); k++ {
				x := a[j+k]
				y := ss * a[j+k+(i>>1)] % mod
				a[j+k] = (x + y) % mod
				a[j+k+(i>>1)] = (x - y + mod) % mod
				ss = ss * s % mod
			}
		}
	}
	if p != 0 {
		s := powMod(N, mod-2)
		for i := 0; i < N; i++ {
			a[i] = a[i] * s % mod
		}
	}
}

func mul() {
	ntt(a, 0)
	ntt(b, 0)
	for i := 0; i < N; i++ {
		c[i] = a[i] * b[i] % mod
	}
	ntt(c, 1)
}

func calc() {
	for i := range a {
		a[i] = 0
	}
	for i := range b {
		b[i] = 0
	}
	for i := 0; i <= n; i++ {
		a[i] = A[i] * fac[i] % mod
		b[i] = inv[n-i]
	}
	mul()
	for i := 0; i <= n; i++ {
		A[n-i] = mi[i] * inv[i] % mod * c[n+i] % mod
	}
	for i := range a {
		a[i] = 0
	}
	for i := range b {
		b[i] = 0
	}
	for i := 0; i <= n; i++ {
		a[i] = A[i] * fac[i] % mod
		b[i] = inv[n-i]
		if (n-i)&1 != 0 {
			b[i] = mod - b[i]
		}
	}
	mul()
	for i := 0; i <= n; i++ {
		A[i] = inv[i] * c[n+i] % mod
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	initMod()
	mi[0] = 1
	for i := 0; i < N; i++ {
		rev[i] = (rev[i>>1] >> 1) + ((i & 1) * (N >> 1))
	}
	for i := 1; i < N; i++ {
		mi[i] = (mi[i-1] << 1) % mod
	}

	var k int
	fmt.Fscan(in, &n, &k)
	for i := 0; i <= n; i++ {
		fmt.Fscan(in, &A[i])
	}
	calc()
	for i := 0; i <= n; i++ {
		A[i] = powMod((2*i-n+mod)%mod, k) * A[i] % mod
	}
	calc()
	C := powMod(1e9*powMod(n, k)%mod*mi[n]%mod, mod-2)
	for i := 0; i <= n; i++ {
		fmt.Fprint(out, C*A[i]%mod, " ")
	}
}

const mod = 998244353

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
