package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

const mod = 998244353

func ksm(a, b int) int {
	ans := 1
	for b > 0 {
		if b&1 != 0 {
			ans *= a
			ans %= mod
		}
		b >>= 1
		a *= a
		a %= mod
	}
	return ans
}
func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 17

	var cntr, cntb, f, fac [1<<N + 1]int
	fac[0] = 1
	for i := 1; i <= 1<<N; i++ {
		fac[i] = (fac[i-1]) * i % mod
	}

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 0; i < m; i++ {
		var x int
		fmt.Fscan(in, &x)
		cntr[1<<(x-1)]++
	}
	for i := 0; i < m; i++ {
		var x int
		fmt.Fscan(in, &x)
		cntb[1<<(x-1)]++
	}
	for i := 1; i < 1<<n; i <<= 1 {
		for j := 0; j < 1<<n; j++ {
			if i&j != 0 {
				cntr[j] += cntr[j^i]
				cntb[j] += cntb[j^i]
			}
		}
	}
	for i := 0; i < 1<<n; i++ {
		if cntr[i] == cntb[i] {
			f[i] = fac[cntr[i]]
		}
	}

	ans := 0
	for i := 1; i < 1<<n; i++ {
		if cntr[i] == cntb[i] {
			for j := (i - 1) & i; j != 0 && clz(uint32(j)) == clz(uint32(i)); j = (j - 1) & i {
				if cntr[j] == cntb[j] {
					f[i] = (f[i] - (f[j])*fac[cntr[i^j]]%mod + mod) % mod
				}
			}
			ans = (ans + f[i]*(fac[m-cntr[i]])) % mod
		}
	}
	fmt.Println(ans * ksm(fac[m], mod-2) % mod)
}

func clz(x uint32) int {
	return bits.LeadingZeros32(x)
}
