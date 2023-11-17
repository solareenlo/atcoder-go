package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for t > 0 {
		t--
		var n, a, b, c int
		fmt.Fscan(in, &n, &a, &b, &c)
		g := gcd(gcd(a, b), c)
		a /= g
		b /= g
		c /= g
		x, y := n, n
		if gcd(b, c) > 1 {
			h := gcd(b, c)
			x /= h
			b /= h
			c /= h
		}
		e := a * inv_mod(b, c) % c
		fmt.Fprintln(out, floor_sum(x+1, c, e, y)-floor_sum(x+1, c, e, c-1)+x+1)
	}
}

func floor_sum(n, m, a, b int) int {
	ans := 0
	if a >= m {
		ans += (n - 1) * n * (a / m) / 2
		a %= m
	}
	if b >= m {
		ans += n * (b / m)
		b %= m
	}

	yMax := (a*n + b) / m
	xMax := (yMax*m - b)
	if yMax == 0 {
		return ans
	}
	ans += (n - (xMax+a-1)/a) * yMax
	ans += floor_sum(yMax, a, m, (a-xMax%a)%a)
	return ans
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
