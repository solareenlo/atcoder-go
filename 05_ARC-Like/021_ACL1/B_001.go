package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Scan(&n)

	n *= 2
	mn := 1 << 60
	for i := 1; i*i <= n; i++ {
		if n%i == 0 && gcd(i, n/i) == 1 {
			a := inv_mod(i, n/i) * i
			tmp := 0
			if a < 2 {
				tmp = 1
			}
			mn = min(mn, a+n*tmp)
			a = inv_mod(n/i, i) * (n / i)
			mn = min(mn, a+n*tmp)
		}
	}
	fmt.Println(mn - 1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

type pair struct{ x, y int }

func inv_mod(x, m int) int {
	if 1 > m {
		os.Exit(1)
	}
	z := inv_gcd(x, m)
	if z.x != 1 {
		os.Exit(1)
	}
	return z.y
}

func inv_gcd(a, b int) pair {
	a = safe_mod(a, b)
	if a == 0 {
		return pair{b, 0}
	}
	s := b
	t := a
	m0 := 0
	m1 := 1
	for t > 0 {
		u := s / t
		s -= t * u
		m0 -= m1 * u
		tmp := s
		s = t
		t = tmp
		tmp = m0
		m0 = m1
		m1 = tmp
	}
	if m0 < 0 {
		m0 += b / s
	}
	return pair{s, m0}
}

func safe_mod(x, m int) int {
	x %= m
	if x < 0 {
		x += m
	}
	return x
}
