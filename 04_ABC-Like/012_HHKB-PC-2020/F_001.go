package main

import (
	"fmt"
	"sort"
)

func integrate(coef []int, lo, hi int) int {
	pow_lo := 1
	pow_hi := 1
	ret := 0
	for i := 0; i < len(coef); i++ {
		pow_lo *= lo
		pow_lo %= mod
		pow_hi *= hi
		pow_hi %= mod
		ret += divMod(coef[i]*(pow_hi-pow_lo+mod)%mod, i+1)
		ret %= mod
	}
	return ret
}

func main() {
	var n int
	fmt.Scan(&n)

	type pair struct{ x, y int }
	ev := make([]pair, n*2)
	left := make([]int, n)
	right := make([]int, n)
	denom := 1
	for i := 0; i < n; i++ {
		var l, r int
		fmt.Scan(&l, &r)
		ev[i*2] = pair{l, i + 1}
		ev[i*2+1] = pair{r, -i - 1}
		left[i] = l
		right[i] = r
		denom *= r - l
		denom %= mod
		denom *= i + 2
		denom %= mod
	}
	sort.Slice(ev, func(i, j int) bool {
		return ev[i].x < ev[j].x
	})

	poly := make([]int, 0)
	poly = append(poly, 1)
	ans := 0
	for i := n*2 - 1; i >= 0; i-- {
		if ev[i].y > 0 {
			ans += ev[i].x
			ans %= mod
			break
		}
		idx := -ev[i].y - 1
		a := divMod(1, right[idx]-left[idx])
		b := (-left[idx]*a + mod) % mod
		poly = append(poly, poly[len(poly)-1]*a%mod)
		for i := len(poly) - 2; i > 0; i-- {
			poly[i] = poly[i-1]*a%mod + poly[i]*b%mod
			poly[i] %= mod
		}
		poly[0] *= b
		poly[0] %= mod
		lo := ev[i-1].x
		hi := ev[i].x
		ans += (hi - lo - integrate(poly, lo, hi) + mod) % mod
	}
	fmt.Println(ans * denom % mod)
}

const mod = 1000000007

func divMod(a, b int) int {
	ret := a * modInv(b)
	ret %= mod
	return ret
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
