package main

import (
	"fmt"
	"math"
)

func rotate(vec complex128, ang float64) complex128 {
	x := real(vec)
	y := imag(vec)
	return complex(x*math.Cos(ang)-y*math.Sin(ang), x*math.Sin(ang)+y*math.Cos(ang))
}

func main() {
	var n float64
	fmt.Scan(&n)

	var x, y float64
	fmt.Scan(&x, &y)
	p0 := complex(x, y)
	fmt.Scan(&x, &y)
	p1 := complex(x, y)

	center := (p0 + p1) / 2.0
	res := rotate(p0-center, 2.0*math.Pi/n) + center
	fmt.Printf("%0.9f %0.9f\n", real(res), imag(res))
}
