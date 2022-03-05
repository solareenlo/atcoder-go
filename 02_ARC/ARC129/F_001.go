package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n   int
	m   int
	ans int
	mi  = make([]int, N)
	rev = make([]int, N)
	a   = make([]int, N)
	A   = make([]int, N)
	B   = make([]int, N)
	g   = make([]int, N)
)

func Add(x, y int) int {
	if x+y < mod {
		return x + y
	}
	return x + y - mod
}

func Dec(x, y int) int {
	if x >= y {
		return x - y
	}
	return x - y + mod
}

func ntt(a []int, p int) {
	for i := 0; i < N; i++ {
		if i < rev[i] {
			a[i], a[rev[i]] = a[rev[i]], a[i]
		}
	}
	for i, l := 2, 0; i <= N; i <<= 1 {
		s := powMod(3, (mod-1)/i)
		if p != 0 {
			s = powMod(s, mod-2)
		}
		g[0] = 1
		for k := 1; k < (i >> 1); k++ {
			g[k] = g[k-1] * s % mod
		}
		for j := 0; j < N; j += i {
			for k := 0; k < (i >> 1); k++ {
				x := a[j+k]
				y := a[j+k+(i>>1)] * g[k] % mod
				a[j+k] = Add(x, y)
				a[j+k+(i>>1)] = Dec(x, y)
			}
		}
		l++
	}
	if p != 0 {
		s := powMod(N, mod-2)
		for i := 0; i < N; i++ {
			a[i] = a[i] * s % mod
		}
	}
}

func calc() int {
	ans := 0
	for i := range A {
		A[i] = 0
	}
	for i := 1; i < n; i++ {
		A[i] = a[i] * fac[n-i-1] % mod * fac[i+m-2] % mod
	}
	ntt(A, 0)

	for i := range B {
		B[i] = 0
	}
	for i := 1; i < min(n, m); i++ {
		B[i] = mi[i] * inv[i-1] % mod * inv[m-i-1] % mod
	}
	ntt(B, 0)
	for i := 0; i < N; i++ {
		B[i] = A[i] * B[i] % mod
	}
	ntt(B, 1)
	for i := 1; i <= n; i++ {
		ans = (ans + 4*B[i]*inv[n-i]%mod*inv[i-1]) % mod
	}

	for i := range B {
		B[i] = 0
	}
	for i := 1; i < min(n, m+1); i++ {
		B[i] = mi[i] * inv[i-1] % mod * inv[m-i] % mod
	}
	ntt(B, 0)
	for i := 0; i < N; i++ {
		B[i] = A[i] * B[i] % mod
	}
	ntt(B, 1)
	for i := 2; i <= n; i++ {
		ans = (ans + B[i]*inv[n-i]%mod*inv[i-2]) % mod
	}

	for i := range B {
		B[i] = 0
	}
	for i := 1; i < min(n, m-1); i++ {
		B[i] = mi[i] * inv[i-1] % mod * inv[m-i-2] % mod
	}
	ntt(B, 0)
	for i := 0; i < N; i++ {
		B[i] = A[i] * B[i] % mod
	}
	ntt(B, 1)
	for i := 0; i <= n; i++ {
		ans = (ans + 3*B[i]*inv[n-i]%mod*inv[i]) % mod
	}

	ans = 4 * (mod + 1) / 3 * ans % mod
	for x := 1; x <= n; x++ {
		ans = (ans + 5*nCrMod(n-1, x-1)*nCrMod(m-1, x-1)%mod*a[n]%mod + mod) % mod
	}
	for x := 2; x <= n; x++ {
		ans = (ans + nCrMod(n-1, x-1)*nCrMod(m-1, x-2)%mod*a[n]%mod + mod) % mod
	}
	for x := 1; x <= n; x++ {
		ans = (ans + 4*nCrMod(n-1, x-1)*nCrMod(m-1, x)%mod*a[n]) % mod
	}
	return ans
}

func main() {
	in := bufio.NewReader(os.Stdin)

	initMod()

	mi[0] = 1
	for i := 1; i < N; i++ {
		mi[i] = 9 * mi[i-1] % mod
	}
	for i := 0; i < N; i++ {
		rev[i] = (rev[i>>1] >> 1) + ((i & 1) * (N >> 1))
	}

	fmt.Fscan(in, &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	ans = calc()
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &a[i])
	}
	n, m = m, n
	ans = (ans + calc()) % mod
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

const mod = 998244353
const N = 1 << 19

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

func nCrMod(n, r int) int {
	if n < r || n < 0 || r < 0 {
		return 0
	}
	return fac[n] * inv[r] % mod * inv[n-r] % mod
}
