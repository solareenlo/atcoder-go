package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, A, B int
	fmt.Fscan(in, &N, &A, &B)
	r := 1000000
	used := make([]bool, r+1)
	prime := make([]bool, N)
	for i := range prime {
		prime[i] = true
	}
	for i := 2; i <= r; i++ {
		if used[i] {
			continue
		}
		rem, mod := Crt([]int{B, 0}, []int{A, i})
		if mod != 0 {
			if rem < B {
				rem += (B - rem + mod - 1) / mod * mod
			}
			if rem == i {
				rem += mod
			}
			rem -= B
			mod /= A
			rem /= A
			for rem < N {
				prime[rem] = false
				rem += mod
			}
		}
		for j := i; j <= r; j += i {
			used[j] = true
		}
	}
	fmt.Println(count(prime, true))
}

func count(slice []bool, target bool) int {
	count := 0
	for _, num := range slice {
		if num == target {
			count++
		}
	}
	return count
}

// Ref: https://github.com/monkukui/ac-library-go
// (rem, mod)
func Crt(r, m []int) (int, int) {
	if len(r) != len(m) {
		panic("")
	}
	n := len(r)
	r0, m0 := 0, 1
	for i := 0; i < n; i++ {
		if !(1 <= m[i]) {
			panic("")
		}
		r1 := SafeMod(r[i], m[i])
		m1 := m[i]
		if m0 < m1 {
			r0, r1 = r1, r0
			m0, m1 = m1, m0
		}
		if m0%m1 == 0 {
			if r0%m1 != r1 {
				return 0, 0
			}
			continue
		}
		g, im := InvGcd(m0, m1)

		u1 := m1 / g
		if (r1-r0)%g != 0 {
			return 0, 0
		}

		x := (r1 - r0) / g % u1 * im % u1

		r0 += x * m0
		m0 *= u1
		if r0 < 0 {
			r0 += m0
		}
	}
	return r0, m0
}

func SafeMod(x, m int) int {
	x %= m
	if x < 0 {
		x += m
	}
	return x
}

func InvGcd(a, b int) (int, int) {
	a = SafeMod(a, b)
	if a == 0 {
		return b, 0
	}

	s := b
	t := a
	m0, m1 := 0, 1

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
	return s, m0
}
