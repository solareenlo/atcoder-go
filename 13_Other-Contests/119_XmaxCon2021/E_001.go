package main

import "fmt"

func main() {
	const MOD = 998244353

	var k int
	fmt.Scan(&k)
	x := 1
	y := 0
	z := (MOD - 1) % MOD
	w := 1
	a := (MOD - 1) % MOD
	b := 1
	c := 1
	s, t := 1, 0
	for i := 0; k > 0; i++ {
		x, z = z, x
		y, w = w, y
		x = (x - s*z%MOD + MOD) % MOD
		y = (y - s*w%MOD + MOD) % MOD
		if i%3 == 1 {
			t += 2
			s = t
		} else {
			s = 1
		}
		s = min(s, k)
		k = k - s
		a = (a - s*x%MOD + MOD) % MOD
		b = (b - s*y%MOD + MOD) % MOD
		if i&1 != 0 {
			c = (c - s*y%MOD + MOD) % MOD
		} else {
			c = (c + s*y%MOD) % MOD
		}
	}
	fmt.Println((c-1+MOD)%MOD, a, b)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
