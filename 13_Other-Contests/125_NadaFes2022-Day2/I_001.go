package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = int(1e9 + 7)
const Mod = int(1e9 + 6)
const Mod2 = int(5e8) + 3
const N = 200005
const mx = 200000

var n, k int
var f, jie, Jie, ni, ff [N]int
var mp [N][22]int

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &k)

	jie[0] = 1
	ni[0] = 1
	Jie[0] = 1
	for i := 1; i <= mx; i++ {
		jie[i] = jie[i-1] * i % Mod2
		Jie[i] = Jie[i-1] * i % Mod
	}
	ni[mx] = kuai(jie[mx], Mod2-2, Mod2)
	for i := mx - 1; i > 0; i-- {
		ni[i] = ni[i+1] * (i + 1) % Mod2
	}

	for i := 1; i <= n; i++ {
		var x int
		fmt.Fscan(in, &x)
		for j := 2; j*j <= x; j++ {
			if x%j == 0 {
				s := 0
				for x%j == 0 {
					x /= j
					s++
				}
				mp[j][s]++
				ff[j] = 1
			}
		}
		if x != 1 {
			mp[x][1]++
			ff[x] = 1
		}
	}

	ans := 1
	for i := 1; i <= mx; i++ {
		if ff[i] != 0 {
			s := 0
			las := 0
			for j := 20; j >= 0; j-- {
				s += mp[i][j]
				now := (solve(s) - las + Mod) % Mod
				ans = ans * kuai(i, j*now%Mod, mod) % mod
				las = (las + now) % Mod
			}
		}
	}
	fmt.Println(ans)
}

func kuai(a, b, mod int) int {
	ans := 1
	for ; b > 0; b, a = b>>1, a*a%mod {
		if (b & 1) != 0 {
			ans = ans * a % mod
		}
	}
	return ans
}

func C(n, m int) int {
	t := jie[n] * ni[m] % Mod2 * ni[n-m] % Mod2
	if (n & m) == m {
		t += Mod2
	}
	return t
}

func solve(x int) int {
	if f[x] != 0 {
		return f[x]
	}
	ans := 0
	for i, bit := 1, 1; i <= x/k; i, bit = i+1, Mod-bit {
		ans = (ans + bit*C(n/k, i)%Mod*C(n-i*k, x-i*k)%Mod) % Mod
	}
	f[x] = ans * Jie[x] % Mod * Jie[n-x] % Mod
	return f[x]
}
