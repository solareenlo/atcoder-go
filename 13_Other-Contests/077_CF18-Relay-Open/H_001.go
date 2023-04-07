package main

import (
	"fmt"
	"math/bits"
)

func main() {
	const mod = 1000000007

	var fact, inv, invfact [1515151]int
	fact[0] = 1
	inv[0] = 1
	invfact[0] = 1

	var n, m, k int
	fmt.Scan(&n, &m, &k)

	all := 1 << n
	fact[0] = 1
	fact[1] = 1
	invfact[0] = 1
	invfact[1] = 1
	inv[1] = 1
	for i := 2; i < all+3; i++ {
		fact[i] = fact[i-1] * i % mod
		inv[i] = mod - inv[mod%i]*(mod/i)%mod
		invfact[i] = invfact[i-1] * inv[i] % mod
	}

	var cnt [1 << 20]int
	for bit := 0; bit < all; bit++ {
		if bit > all-m {
			continue
		}
		cnt[all-bit-1] = fact[all-m] * invfact[all-m-bit] % mod
		cnt[all-bit-1] = cnt[all-bit-1] * fact[all-bit-1] % mod
	}

	for i := 0; i < n; i++ {
		for j := 0; j < all; j++ {
			if ((1 << i) & j) != 0 {
				cnt[j] = (cnt[j] - cnt[j^(1<<i)] + mod) % mod
			}
		}
	}

	ans := 0
	for i := 0; i < all; i++ {
		if bits.OnesCount(uint(i)) > k {
			continue
		}
		ans = (ans + cnt[i]) % mod
	}
	fmt.Println(ans * all % mod)
}
