package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	n = abs(n)

	cnt := map[int]int{}
	for i := 2; i*i <= n; i++ {
		if n%i != 0 {
			continue
		}
		for n%i == 0 {
			n /= i
			cnt[i]++
		}
	}
	if n > 1 {
		cnt[n]++
	}

	keys := make([]int, 0, len(cnt))
	for k := range cnt {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	fact := make([]int, 200200)
	fact[0] = 1
	for i := 1; i < 200200; i++ {
		fact[i] = fact[i-1] * i % mod
	}

	ans := powMod(2, m-1)
	for _, x := range keys {
		ans *= divMod(divMod(fact[m+cnt[x]-1], fact[m-1]), fact[cnt[x]])
		ans %= mod
	}
	fmt.Println(ans)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

const mod = 1000000007

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		n /= 2
	}
	return res
}

func modInv(a int) int {
	b, u, v := mod, 1, 0
	for b != 0 {
		t := a / b
		a -= t * b
		a, b = b, a
		u -= t * v
		u, v = v, u
	}
	u %= mod
	if u < 0 {
		u += mod
	}
	return u
}

func divMod(a, b int) int {
	ret := a * modInv(b)
	ret %= mod
	return ret
}
