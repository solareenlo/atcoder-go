package main

import (
	"fmt"
	"math"
)

var t, p [105]float64

func main() {
	var n int
	fmt.Scan(&n)
	p[0] = 1
	for i := 0; i < n; i++ {
		p[i+1] = p[i] * float64(i+1)
	}
	fmt.Println(f(n))
}

func f(n int) float64 {
	if n == 1 {
		return 0
	}
	if t[n] != 0 {
		return t[n]
	}

	m := math.Pow(float64(3), float64(n))

	res := m
	for a := 0; a < n; a++ {
		for b := 0; b <= n-a; b++ {
			c := n - a - b
			if a == n || b == n || c == n {
				continue
			}
			if a == b && b == c && c == a {
				continue
			}
			k := min(a, b, c)
			if k == 0 {
				k = max(min(a, b), min(b, c), min(c, a))
			}
			res += f(k) * p[n] / (p[a] * p[b] * p[c])
		}
	}

	m -= 3
	if n%3 == 0 {
		k := n / 3
		m -= p[n] / (p[k] * p[k] * p[k])
	}
	t[n] = res / m
	return t[n]
}

func min(a ...int) int {
	res := a[0]
	for i := range a {
		if res > a[i] {
			res = a[i]
		}
	}
	return res
}

func max(a ...int) int {
	res := a[0]
	for i := range a {
		if res < a[i] {
			res = a[i]
		}
	}
	return res
}
