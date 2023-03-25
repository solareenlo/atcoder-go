package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func dot(a, b complex128) float64 {
	return real(cmplx.Conj(a) * b)
}

func cross(a, b complex128) float64 {
	return imag(cmplx.Conj(a) * b)
}

type L struct {
	a, b complex128
}

func is_ll(s, t L) complex128 {
	sv, tv := s.b-s.a, t.b-t.a
	return s.a + div(mul(sv, cross(tv, t.a-s.a)), cross(tv, sv))
}

func mul(a complex128, d float64) complex128 {
	return complex(real(a)*d, imag(a)*d)
}

func div(a complex128, d float64) complex128 {
	return complex(real(a)/d, imag(a)/d)
}

func main() {
	var p [6]complex128
	for i := 0; i < 6; i++ {
		var x, y float64
		fmt.Scan(&x, &y)
		p[i] = complex(x, y)
	}
	if math.Abs(cross(p[0]-p[1], p[3]-p[4])) > 1e-9 ||
		math.Abs(cross(p[1]-p[2], p[4]-p[5])) > 1e-9 ||
		math.Abs(cross(p[2]-p[3], p[5]-p[0])) > 1e-9 {
		fmt.Println(-1)
		return
	}
	a := is_ll(L{p[0], p[1]}, L{p[2], p[3]})
	b := is_ll(L{p[2], p[3]}, L{p[4], p[5]})
	c := is_ll(L{p[4], p[5]}, L{p[0], p[1]})
	x := (cmplx.Abs(a-b)*cmplx.Abs(a-b) + cmplx.Abs(b-c)*cmplx.Abs(b-c) - cmplx.Abs(c-a)*cmplx.Abs(c-a)) / 2
	y := (cmplx.Abs(b-c)*cmplx.Abs(b-c) + cmplx.Abs(c-a)*cmplx.Abs(c-a) - cmplx.Abs(a-b)*cmplx.Abs(a-b)) / 2
	z := (cmplx.Abs(c-a)*cmplx.Abs(c-a) + cmplx.Abs(a-b)*cmplx.Abs(a-b) - cmplx.Abs(b-c)*cmplx.Abs(b-c)) / 2
	if x < 1e-9 || y < 1e-9 || z < 1e-9 {
		fmt.Println(-1)
		return
	}
	res := math.Sqrt(math.Max(0, x*y*z))
	res *= cmplx.Abs(a-p[3]) / cmplx.Abs(a-b)
	res *= cmplx.Abs(b-p[5]) / cmplx.Abs(b-c)
	res *= cmplx.Abs(c-p[1]) / cmplx.Abs(c-a)
	fmt.Printf("%.12f\n", res)
}
