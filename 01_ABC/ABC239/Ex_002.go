package main

import (
	"fmt"
	"math"
)

func main() {
	const S = 180180

	var n, m int
	fmt.Scan(&n, &m)
	rev := powMod(n-1, mod-2)
	g := make([]int, S*2+1)
	f := make([]int, S*2+1)
	for i := 1; i < S+1; i++ {
		g[i] = (g[i] + g[i-1]) % mod
		f[i] = (g[i] + n) * rev % mod
		if f[i] < 0 {
			f[i] += mod
		}
		for j := 2; j < min(n, S/i)+1; j++ {
			g[i*j] += f[i]
			g[i*j+j] -= f[i]
		}
	}

	h := make([]int, S)
	LIM := m / S
	for d := m / S; d >= 1; d-- {
		M := m / d
		h[d] = n
		sqM := int(math.Floor(math.Sqrt(float64(M))))
		Go := M / 2
		en := M/(M/2) - 1
		for i := 2; i <= min(M, n); i++ {
			if i > sqM+3 {
				Go--
			} else {
				Go = M / i
			}
			if i < sqM-3 {
				en++
			} else {
				en = M / Go
			}
			if en > n {
				en = n
			}
			if d*i <= LIM {
				h[d] += (en - i + 1) * h[d*i]
			} else {
				h[d] += (en - i + 1) * f[Go]
			}
			i = en
		}
		h[d] = h[d] % mod * rev % mod
	}

	if m < S {
		fmt.Println(f[m])
	} else {
		fmt.Println(h[1])
	}
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
