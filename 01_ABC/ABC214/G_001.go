package main

import (
	"bufio"
	"fmt"
	"os"
)

func Mod(x int) int {
	if x >= mod {
		return x - mod
	}
	return x
}

func main() {
	in := bufio.NewReader(os.Stdin)

	initMod()

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	b := make([]int, n+1)
	to := make([]int, N)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &b[i])
		to[a[i]] = b[i]
	}

	f := make([]int, N)
	f[0] = 1
	var cnt int
	sum := 0
	vis := make([]bool, N)
	for i := 1; i <= n; i++ {
		if !vis[i] {
			cnt = 0
			for u := i; !vis[u]; u = to[u] {
				vis[u] = true
				cnt++
			}
			for j := sum; j >= 0; j-- {
				f[j+1] = (f[j+1] + f[j]*cnt*cnt) % mod
				for k := 2; k <= cnt; k++ {
					f[j+k] = (f[j+k] + f[j]*(nCrMod(cnt+k-1, k<<1)+nCrMod(cnt+k, k<<1))) % mod
				}
				if cnt > 1 {
					f[j] = Mod(f[j] << 1)
				}
			}
			sum += cnt
		}
	}

	ans := 0
	for i := 0; i <= n; i++ {
		if i&1 != 0 {
			ans = Mod(ans + mod - f[n-i]*fact[n-i]%mod)
		} else {
			ans = (ans + f[n-i]*fact[n-i]) % mod
		}
	}
	fmt.Println(ans)
}

const mod = 1000000007
const N = 6006

var fact, invf [N]int

func initMod() {
	fact[0] = 1
	invf[0] = 1
	for i := int(1); i < N; i++ {
		fact[i] = (fact[i-1] * i) % mod
		invf[i] = invMod(fact[i])
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
	return fact[n] * invf[r] % mod * invf[n-r] % mod
}
