package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const EPS = 1e-9

type Point complex128
type Line [2]Point
type Polygon []Point

func sgn(x float64) int {
	if x < -EPS {
		return -1
	}
	if x > +EPS {
		return 1
	}
	return 0
}

func det(x, y Point) float64 {
	return real(x)*imag(y) - imag(x)*real(y)
}

func cross_point(l1, l2 Line) Point {
	u := l1[1] - l1[0]
	v := l2[0] - l2[1]
	c := l2[0] - l1[0]
	return l2[0] - Point(complex(det(u, c)/det(u, v), 0))*v
}

func convex_cut(convex Polygon, l Line) Polygon {
	la := l[0]
	lb := l[1]
	var res Polygon
	sz := len(convex)
	for i := 0; i < sz; i++ {
		j := i + 1
		if j == sz {
			j -= sz
		}
		a := convex[i]
		b := convex[j]
		da := sgn(det(lb-la, a-la))
		if da >= 0 {
			res = append(res, a)
		}
		db := sgn(det(lb-la, b-la))
		if da*db < 0 {
			res = append(res, cross_point(l, Line{a, b}))
		}
	}
	return res
}

func signed_area(poly Polygon) float64 {
	res := 0.0
	n := len(poly)
	for i := 0; i < n; i++ {
		res += det(poly[i], poly[(i+1)%n])
	}
	return res / 2
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	s := make(Polygon, n)
	t := make(Polygon, m)
	for i := 0; i < n; i++ {
		var x, y float64
		fmt.Fscan(in, &x, &y)
		s[i] = Point(complex(x, y))
	}

	for i := 0; i < m; i++ {
		var x, y float64
		fmt.Fscan(in, &x, &y)
		t[i] = Point(complex(x, y))
	}

	for i := 0; i < m; i++ {
		var line Line
		line[0] = t[i]
		line[1] = t[(i+1)%m]
		s = convex_cut(s, line)
	}

	fmt.Println(math.Abs(signed_area(s)))
}
