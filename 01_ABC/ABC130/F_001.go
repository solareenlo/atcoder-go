package main

import "fmt"

var (
	n  int
	x  [100005]float64
	y  [100005]float64
	dx [100005]float64
	dy [100005]float64
)

func check(a float64) float64 {
	miniX, miniY := 1e18, 1e18
	maxX, maxY := -1e18, -1e18
	for i := 0; i < n; i++ {
		xx := x[i] + a*dx[i]
		yy := y[i] + a*dy[i]
		miniX = min(miniX, xx)
		maxX = max(maxX, xx)
		miniY = min(miniY, yy)
		maxY = max(maxY, yy)
	}
	return (maxX - miniX) * (maxY - miniY)
}

func main() {
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var d string
		fmt.Scan(&x[i], &y[i], &d)
		switch d {
		case "D":
			dy[i] = -1
		case "U":
			dy[i] = 1
		case "R":
			dx[i] = 1
		case "L":
			dx[i] = -1
		}
	}

	l, r, res := 0.0, 1e9, 1e18
	for i := 0; i < 200; i++ {
		ll := (l*2 + r) / 3.0
		rr := (r*2 + l) / 3.0
		lll := check(ll)
		rrr := check(rr)
		res = min(res, lll)
		res = min(res, rrr)
		if lll < rrr {
			r = rr
		} else {
			l = ll
		}
	}
	fmt.Println(res)
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
