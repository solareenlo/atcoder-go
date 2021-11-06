package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, h, m float64
	fmt.Scan(&a, &b, &h, &m)

	th := float64(h*60+m) / 720.0 * 2.0 * math.Pi
	tm := float64(m) / 60.0 * 2.0 * math.Pi
	xh := math.Cos(th) * a
	yh := math.Sin(th) * a
	xm := math.Cos(tm) * b
	ym := math.Sin(tm) * b
	dist := math.Hypot(xh-xm, yh-ym)

	fmt.Println(dist)
}
