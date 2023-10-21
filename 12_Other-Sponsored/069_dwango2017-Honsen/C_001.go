package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const R = 100000.0
const L = R * 10.0

var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func ret(x, y float64) bool {
	fmt.Fprintf(out, "%.10f %.10f\n", x, y)
	out.Flush()
	var s string
	fmt.Fscan(in, &s)
	if s == "found" || s == "kill" {
		os.Exit(0)
	}
	if s == "close" {
		return true
	} else {
		return false
	}
}

func ans(x, y int) {
	fmt.Fprintln(out, x, y)
	out.Flush()
	var s string
	fmt.Fscan(in, &s)
	if s == "found" || s == "kill" {
		os.Exit(0)
	}
	return
}

func main() {
	d := L / 8
	cx, cy := 0.0, 0.0
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			x := float64(i)*d + d/2
			y := float64(j)*d + d/2
			b := ret(x, y)
			if b {
				cx = x
				cy = y
				break
			}
		}
		if cx != 0 {
			break
		}
	}

	l, r := -R, cy
	for i := 0; i < 67; i++ {
		m := (l + r) / 2
		b := ret(cx, m)
		if b {
			r = m
		} else {
			l = m
		}
	}

	left := r
	l = cy
	r = L + R
	for i := 0; i < 67; i++ {
		m := (l + r) / 2
		b := ret(cx, m)
		if b {
			l = m
		} else {
			r = m
		}
	}
	right := l
	Y := (left+right)/2 + 0.1
	dis := math.Sqrt(R*R - ((right - left) / 2 * (right - left) / 2))
	X := cx + dis + 0.1
	ans(int(X), int(Y))
	X = cx - dis + 0.1
	ans(int(X), int(Y))
}
