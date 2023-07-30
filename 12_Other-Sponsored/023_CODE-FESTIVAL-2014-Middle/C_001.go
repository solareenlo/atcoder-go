package main

import "fmt"

var n float64
var cnt int

func main() {
	fmt.Scan(&n, &cnt)
	fmt.Println(op(cnt))
}

func op(k int) float64 {
	if k == 1 {
		return n
	}
	x := op(k >> 1)
	y := 1.0 - x
	a := x * y * 2.0
	b := x*x + y*y
	if (k & 1) != 0 {
		a = a*(1.0-n) + b*n
	}
	return a
}
