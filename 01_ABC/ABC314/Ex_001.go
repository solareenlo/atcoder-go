package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func PointToSegDist(x, y, x1, y1, x2, y2 float64) float64 {
	cross := (x2-x1)*(x-x1) + (y2-y1)*(y-y1)
	if cross <= 0 {
		return math.Sqrt((x-x1)*(x-x1) + (y-y1)*(y-y1))
	}
	d2 := (x2-x1)*(x2-x1) + (y2-y1)*(y2-y1)
	if cross >= d2 {
		return math.Sqrt((x-x2)*(x-x2) + (y-y2)*(y-y2))
	}
	r := cross / d2
	px := x1 + (x2-x1)*r
	py := y1 + (y2-y1)*r
	return math.Sqrt((x-px)*(x-px) + (y-py)*(y-py))
}

var n int
var xa, xb, ya, yb [110]float64

func g(x, y float64) float64 {
	mx := 0.0
	for i := 1; i <= n; i++ {
		mx = math.Max(mx, PointToSegDist(x, y, xa[i], ya[i], xb[i], yb[i]))
	}
	return mx
}

func f(x float64) float64 {
	l, r := 0.0, 1e3
	cnt := 100
	for cnt > 0 {
		cnt--
		p1 := l + (r-l)/3
		p2 := r - (r-l)/3
		if g(x, p1) > g(x, p2) {
			l = p1
		} else {
			r = p2
		}
	}
	return g(x, l)
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &xa[i], &ya[i], &xb[i], &yb[i])
	}
	l, r := 0.0, 1e3
	cnt := 100
	for cnt > 0 {
		cnt--
		p1 := l + (r-l)/3
		p2 := r - (r-l)/3
		if f(p1) > f(p2) {
			l = p1
		} else {
			r = p2
		}
	}
	fmt.Println(f(l))
}
