package main

import "fmt"

func main() {
	var l, a, b, m int
	fmt.Scan(&l, &a, &b, &m)
	a -= b

	d := [19]int{}
	f := [19]int{}
	t := 0
	for i := 1; i < 19; i++ {
		t = (t+1)*10 - 1
		d[i] = max(0, min(l, (t-a)/b))
		f[i] = (a + b*d[i]) % m
	}

	for i := 18; i > 0; i-- {
		d[i] -= d[i-1]
	}

	p := [19]int{}
	p[0] = 1
	for i := 1; i < 19; i++ {
		p[i] = p[i-1] * 10 % m
	}

	res := 0
	t = 1
	for i := 18; i > 0; i-- {
		s, u, v := f[i], p[i], b%m
		for d[i] > 0 {
			if d[i]%2 != 0 {
				res = (res + s*t%m) % m
				t = t * u % m
				s = (s + m - v%m) % m
			}
			s = (s + (s+m-v%m)%m*u%m) % m
			v = v * (u + 1) % m * 2 % m
			u = u * u % m
			d[i] /= 2
		}
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
