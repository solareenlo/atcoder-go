package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var X, Y, P, Q int64
		fmt.Scan(&X, &Y, &P, &Q)

		res := int64(1 << 62)
		for t1 := X; t1 < X+Y; t1++ {
			t, lcm := Crt([]int64{t1, P}, []int64{(X + Y) * 2, P + Q})
			if lcm != 0 {
				res = min(res, t)
			}
		}

		for t2 := P; t2 < P+Q; t2++ {
			t, lcm := Crt([]int64{X, t2}, []int64{(X + Y) * 2, P + Q})
			if lcm != 0 {
				res = min(res, t)
			}
		}

		if res == 1<<62 {
			fmt.Println("infinity")
		} else {
			fmt.Println(res)
		}
	}
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

// Ref: https://github.com/monkukui/ac-library-go
// (rem, mod)
func Crt(r, m []int64) (int64, int64) {
	if len(r) != len(m) {
		panic("")
	}
	n := len(r)
	r0 := int64(0)
	m0 := int64(1)
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

func SafeMod(x, m int64) int64 {
	x %= m
	if x < 0 {
		x += m
	}
	return x
}

func InvGcd(a, b int64) (int64, int64) {
	a = SafeMod(a, b)
	if a == 0 {
		return b, 0
	}

	s := b
	t := a
	m0 := int64(0)
	m1 := int64(1)

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
